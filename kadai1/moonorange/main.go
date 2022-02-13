package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"moonorange/converter"
)

func errHandle(err error) {
	if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
	}
}

func main() {
	var (
		srcDir  = flag.String("src", ".", "Directory of images that you want to convert")
		dstDir  = flag.String("dst", "./dst", "Destination of converted images")
		fromExt = flag.String("from", "jpg", "Extension before conversion")
		toExt   = flag.String("to", "png", "Extension after conversion")
	)

	// To parse the command line into the defined flags.
	flag.Parse()
	_, err := converter.ChkAvailFmt(*fromExt, *toExt)
	errHandle(err)

	err = filepath.Walk(*srcDir, func(path string, info os.FileInfo, err error) error {
		errHandle(err)

		if info.IsDir() {
			return nil
		}
		// Open file and decode it to image.Image
		file, err := os.Open(path)
		errHandle(err)
		defer file.Close()

		img, format, err := image.Decode(file)
		if converter.AvailFmt[format] && err == nil {
			err = converter.Do(*dstDir, *fromExt, *toExt, path, format, img)
			// If you want to unwrap original error, uncomment the following line
			// fmt.Println(errors.Unwrap(err))
			errHandle(err)
		}
		return nil
	})
	errHandle(err)
}
