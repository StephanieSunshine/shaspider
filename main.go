package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of shaspider [options] <filename>:\n")
		flag.PrintDefaults()
	}

	blocksizePtr := flag.Int64("blocksize", 1073741824, "Blocksize in bytes")

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(flag.Args()[0])

	if err != nil {
		fmt.Println("cannot able to read the file", err)
		return
	} // UPDATE: close after checking error

	defer f.Close() //Do not forget to close the file

	r := bufio.NewReader(f)

	// 1M = 1048576 or 1024^2
	// 1G = 1073741824 or 1024^3
	buf := make([]byte, *blocksizePtr) //the chunk size

	var blockList []string

	for {

    //buf = buf[:0] // zero but keep memory
		n, err := r.Read(buf) //loading chunk into buffer
		buf = buf[:n]

		if n == 0 {
			if err != nil {
				//fmt.Println(err)
				break
			}
			if err == io.EOF {
				break
			}
			// return err
		}

		h := sha256.New()
		h.Write(buf)
		blockList = append(blockList, fmt.Sprintf("%x", h.Sum(nil)))
	}

	j, err := json.Marshal(blockList)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}
}
