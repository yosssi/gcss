# GCSS - Pure Go CSS Preprocessor

[![wercker status](https://app.wercker.com/status/4857161fd705e6c43df492e6a33ce87f/m "wercker status")](https://app.wercker.com/project/bykey/4857161fd705e6c43df492e6a33ce87f)
[![Coverage Status](https://img.shields.io/coveralls/yosssi/gcss.svg)](https://coveralls.io/r/yosssi/gcss?branch=master)
[![GoDoc](https://godoc.org/github.com/yosssi/gcss?status.svg)](https://godoc.org/github.com/yosssi/gcss)

## Overview

GCSS is a pure Go CSS preprocessor. This is inspired by [Sass](http://sass-lang.com/) and [Stylus](http://learnboost.github.io/stylus/).

## Syntax

### Variables

```scss
$base-font: Helvetica, sans-serif
$main-color: blue

body
  font: 100% $base-font
  color: $main-color
```

### Nesting

```scss
nav
  ul
    margin: 0
    padding: 0

a
  color: blue
  &:hover
    color: red
```

### Mixins

```scss
$border-radius($radius)
  -webkit-border-radius: $radius
  -moz-border-radius: $radius
  -ms-border-radius: $radius
  border-radius: $radius

.box
  $border-radius(10px)
```

## Installation

```sh
$ go get -u github.com/yosssi/gcss/...
```

## Compile from the Command-Line

```sh
$ gcss /path/to/gcss/file
```

## Compile from Go programs

You can compile a GCSS file from Go programs by invoking `gcss.Compile` function. Please see the [GoDoc](http://godoc.org/github.com/yosssi/gcss) for the datails.

```go
pathc, errc := gcss.Compile("path_to_gcss_file")

select {
case path := <-pathc:
	http.ServeFile(w, r, path)
case err := <-errc:
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}
```

## Documentation

* [GoDoc](http://godoc.org/github.com/yosssi/gcss)

## Syntax Highlightings

* [vim-gcss](https://github.com/yosssi/vim-gcss) - Vim syntax highlighting for GCSS
