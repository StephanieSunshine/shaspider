<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/StephanieSunshine/shaspider">
    <img src="images/logo.svg" alt="Logo" width="320" height="320">
  </a>

  <h3 align="center">Shaspider</h3>

  <p align="center">
    An awesome new way to figure out which part of a file is bad using checksums!
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple example steps. Your result will be either a json array of sha256 checksums split by `blocksize` or an error message.

### Prerequisites

1. [Golang](https://golang.org)

### Installation

1. Clone repo `git clone https://github.com/StephanieSunshine/shaspider`
2. Install `go install`

<!-- USAGE EXAMPLES -->
## Usage
```
Usage of shaspider [options] <filename>:
  -blocksize int
    	Blocksize in bytes (default 1073741824)
```

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->
## License

Distributed under the MIT License. See [LICENSE](https://github.com/StephanieSunshine/shaspider/blob/master/LICENSE) for more information.

## Acknowledgements

* [Best Readme Template](https://github.com/othneildrew/Best-README-Template)
* [Spider Logo](https://thenounproject.com/search/?q=spider&i=2013883)

<!-- CONTACT -->
## Contact

* Project Link: [https://github.com/StephanieSunshine/shaspider](https://github.com/StephanieSunshine/shaspider)
* Email: [Gmail](mailto:ponyosunshine)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/StephanieSunshine/shaspider
[contributors-url]: https://github.com/StephanieSunshine/shaspider/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/StephanieSunshine/shaspider
[forks-url]: https://github.com/StephanieSunshine/shaspider/network/members
[stars-shield]: https://img.shields.io/github/stars/StephanieSunshine/shaspider
[stars-url]: https://github.com/StephanieSunshine/shaspider/stargazers
[issues-shield]: https://img.shields.io/github/issues/StephanieSunshine/shaspider
[issues-url]: https://github.com/StephanieSunshine/shaspider/issues
[license-shield]: https://img.shields.io/github/license/StephanieSunshine/shaspider
[license-url]: https://github.com/StephanieSunshine/shaspider/blob/master/LICENSE
