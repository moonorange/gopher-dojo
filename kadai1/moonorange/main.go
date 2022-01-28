package main

import (
	"errors"
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

var availFmt = map[string]bool {
	"jpeg" : true,
	"jpg" : true,
	"png" : true,
}

func chkAvailFmt(fromExt, toExt string) error {
	if !availFmt[fromExt] {
		err :=  fmt.Errorf("invalid input file extension %s", fromExt)
		return err
	}
	if !availFmt[toExt] {
		err :=  fmt.Errorf("invalid input file extension %s", toExt)
		return err
	}
	if fromExt == toExt {
		return errors.New("from and to ext should be different from each other")
	}
	return nil
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
	err := chkAvailFmt(*fromExt, *toExt)
	errHandle(err)

	filepath.Walk(*srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if info.IsDir() {
			return nil
		}
		// Open file and decode it to image.Image
		file, err := os.Open(path)
		errHandle(err)
		defer file.Close()

		img, format, err := image.Decode(file)
		if availFmt[format] && err == nil {
			err = converter.Do(*dstDir, *fromExt, *toExt, path, format, img)
			errHandle(err)
		}
		return nil
	})
}
