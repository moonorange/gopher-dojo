package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"moonorange/converter"
)

func main() {
	var (
		srcDir  = flag.String("src", ".", "Directory of images that you want to convert")
		dstDir  = flag.String("dst", "./dst", "Destination of converted images")
		fromExt = flag.String("from", "jpg", "Extension before conversion")
		toExt   = flag.String("to", "png", "Extension after conversion")
	)

	// To parse the command line into the defined flags.
	flag.Parse()

	filepath.Walk(*srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if info.IsDir() {
			return nil
		}
		err = converter.Do(*srcDir, *dstDir, *fromExt, *toExt, path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return nil
	})
}
