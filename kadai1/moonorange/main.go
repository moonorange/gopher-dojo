package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

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

func keys(m map[string] bool) []string {
	ks := []string{}
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

func chkAvailFmt(fromExt, toExt string) error {
	var err error
	if !availFmt[fromExt] || !availFmt[toExt]{
		err =  fmt.Errorf("invalid file extension. supported format is " + strings.Join(keys(availFmt), " "))
	}
	if fromExt == toExt {
		err = fmt.Errorf("from and to flag value should be different from each other. from:%s to:%s", fromExt, toExt)
	}
	return err
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
		if availFmt[format] && err == nil {
			err = converter.Do(*dstDir, *fromExt, *toExt, path, format, img)
			errHandle(err)
		}
		return nil
	})
	errHandle(err)
}
