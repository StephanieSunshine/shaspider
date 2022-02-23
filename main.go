package main
//
// Shaspider -- 2022 Stephanie Sunshine -- MIT License
//
import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
	"sync"
	"runtime"
)

func checkerror(title string, err error, fatal ...bool) {
  if err != nil {
    if len(fatal) > 0 {
      if fatal[0] {
        fmt.Fprintln(os.Stderr, title)
        panic(err)
      }else{
        fmt.Fprintln(os.Stderr, title, err)
        return
      }
    } else {
      // assume fatal is true
      fmt.Fprintln(os.Stderr, title)
      panic(err)
    }
  }
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of shaspider [options] <filename>:\n")
		flag.PrintDefaults()
	}
	// currently 2g
	blocksizePtr := flag.Uint64("blocksize", 2147483648, "Blocksize in bytes")
	corecountPtr := flag.Uint("corecount", uint(runtime.NumCPU()), "Cores to use")
	// progressPtr := flag.Bool("progress", false, "show progress tick")

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(flag.Args()[0])

  checkerror("Error opening file:", err)

	fi, err := f.Stat()

  checkerror("Error stating file:", err)

	f.Close()

	var parcels uint64 = uint64(fi.Size()) / *blocksizePtr

	if parcels == 0 {
		panic("File must be bigger than blocksize")
	}

  // indexes to workers workers can easily share a map and mutexes to make collating easier
	indtoworkers := make(chan uint64, *corecountPtr*3)

  // waitgroup so we can wait for the threads to complete
	var wg sync.WaitGroup

  // mutex for the map
	var rm sync.Mutex

  // the result map itself (index:hash)
	results := make(map[uint64]string)

  // start our threads
	for w:=uint(0); w<*corecountPtr; w++ {
		wg.Add(1)
		go work(flag.Args()[0], indtoworkers, *blocksizePtr, results, &rm, &wg)
	}

  // this is as simple as it looks, the magic trick here is only one child will receive a sent message, we are simply sending a list of parcel chunks to process so that the threads may attack it
	for p:=uint64(0); p<=parcels; p++ {
		indtoworkers <- p
	}

  // let the workeres catch up, we can watch the channel's length to see how many messages are still waiting to process, when we hit zero there is nothing left to process

  sle, _ := time.ParseDuration("300ms")

	for len(indtoworkers) > 0 {
		time.Sleep(sle)
	}

  // close the channel now that the buffer is done
	close(indtoworkers)

  // wait for children to finish processing
	wg.Wait()
  res, _ := json.Marshal(results)

	fmt.Println(string(res))
}

// thread main function
func work(fn string, in chan uint64, blocksize uint64, results map[uint64]string, rm *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	f, err := os.Open(fn)
  checkerror("Error opening file:", err)
	defer f.Close()
	//buf := make([]byte, blocksize)
	sha := sha256.New()
	for i := range(in) {
		//f.ReadAt(buf, int64(i*blocksize))
		sha.Reset()
		f.Seek(int64(i*blocksize), 0)
    // this should only happen when something really bad happened under us
    if _, err := io.CopyN(sha, f, int64(blocksize)); err != nil && err != io.EOF {
			panic(err)
		}
		checksum := fmt.Sprintf("%x", sha.Sum(nil))
		rm.Lock()
		results[i] = checksum
		rm.Unlock()
	}
}
