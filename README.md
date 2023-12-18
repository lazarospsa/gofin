<!-- <p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p> -->

<h3 align="center">GoFin</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/lazarospsa/gofin.svg)](https://github.com/lazarospsa/gofin/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/lazarospsa/gofin.svg)](https://github.com/lazarospsa/gofin/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazarospsa/gofin)](https://goreportcard.com/report/github.com/lazarospsa/gofin)

</div>

---

<p align="center"> GoFin is a Go package for financial calculations and operations.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
<!-- - [Getting Started](#getting_started)
- [Deployment](#deployment) -->
- [Usage](#usage)
- [Built Using](#built_using)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

The `gofin` package provides a collection of financial functions for performing various calculations related to finance. This package is designed to simplify financial calculations and provide a convenient way to handle common financial operations.

The `gofin` package includes functions for calculating compound interest, present value, future value, net present value, internal rate of return, and many other financial calculations. These functions are implemented using industry-standard formulas and algorithms, ensuring accurate and reliable results.

By using the `gofin` package, developers can easily incorporate financial calculations into their applications without having to write complex formulas from scratch. This package is suitable for a wide range of financial applications, including investment analysis, loan calculations, retirement planning, and more.

### Installing

To install GoFin, use the following command:

```
go get github.com/lazarospsa/gofin
```

After that you just import it in your project and you can use the functions.

## 🎈 Usage <a name="usage"></a>

```
package main

import (
	"fmt"
	gofin "github.com/lazarospsa/gofin"
)

func main() {
	fmt.Println(gofin.FutureValue(1000, 0.05, 10))
}
```

## ⛏️ Built Using <a name = "built_using"></a>

- [Go](https://go.dev/) - Programming Language

## ✍️ Authors <a name = "authors"></a>

- [@lazarospsa](https://github.com/lazarospsa) - Idea & Initial work

See also the list of [contributors](https://github.com/lazarospsa/gofin/contributors) who participated in this project.
