package main

import (
	"flag"
)


func main() {
	var (
		srcDir = flag.String("src", ".", "Directory of images that you want to convert")
		dstDir = flag.String("dst", "./dst", "Destination of converted images")
		fromExt = flag.String("from", "jpg", "Extension before conversion")
		toExt = flag.String("to", "png", "Extension after conversion")
	)

	// To parse the command line into the defined flags.
	flag.Parse()
}
