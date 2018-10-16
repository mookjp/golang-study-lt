// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 287.

//!+main

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	format := parseFlag()

	if err := convert(format, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func parseFlag() string {
	jpgFlg := flag.Bool("jpg", false, "convert to jpg")
	pngFlg := flag.Bool("png", false, "convert to png")
	gifFlg := flag.Bool("gif", false, "convert to gif")

	flag.Parse()

	if !*jpgFlg && !*pngFlg && !*gifFlg {
		log.Fatal("Specify -jpg -png or -gif as output format")
	}
	if *jpgFlg && *pngFlg ||
		*pngFlg && *gifFlg ||
		*jpgFlg && *gifFlg {
		log.Fatal("Multiple flags are not supported")
	}

	var format string
	if *jpgFlg {
		format = "jpeg"
	}
	if *pngFlg {
		format = "png"
	}
	if *gifFlg {
		format = "gif"
	}
	return format
}

func convert(format string, in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	var encodeErr error
	switch format {
	case "jpeg":
		encodeErr = jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif":
		encodeErr = gif.Encode(out, img, &gif.Options{})
	case "png":
		encodeErr = png.Encode(out, img)
	}
	return encodeErr
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
