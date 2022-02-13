package converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"moonorange/dict"
)

var AvailFmt = map[string]bool {
	"jpeg" : true,
	"jpg" : true,
	"png" : true,
}

func ChkAvailFmt(fromExt, toExt string) (bool, error) {
	var err error
	if !AvailFmt[fromExt] {
		err =  fmt.Errorf("invalid file extension for fromExt. supported format is " + strings.Join(dict.Keys(AvailFmt), " "))
		return false, err
	}

	if !AvailFmt[toExt] {
		err =  fmt.Errorf("invalid file extension for toExt. supported format is " + strings.Join(dict.Keys(AvailFmt), " "))
		return false, err
	}

	if fromExt == toExt {
		err = fmt.Errorf("from and to flag value should be different from each other. from:%s to:%s", fromExt, toExt)
		return false, err
	}
	return true, nil
}

func Do(dstDir, fromExt, toExt, path, format string, img image.Image) error {
	if err := os.MkdirAll(dstDir, 0777); err != nil {
		return &ConvError{Err: err, Code: DirCreateFail, FilePath: dstDir}
	}
	newfn := filepath.Join(dstDir, filepath.Base(fnWithoutExt(path) + "." + toExt))
	newf, err := os.Create(newfn)
	if err != nil {
		return &ConvError{Err: err, Code: FileCreateFail, FilePath: newfn}
	}

	defer func () {
		if err := newf.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	switch toExt {
	case "jpg", "jpeg":
		if err := jpeg.Encode(newf, img,  nil); err != nil {
			return &ConvError{Err: err, Code: FileEncodeFail, FilePath: newfn}
		}
		fmt.Println("Created", newfn)
	case "png":
		if err := png.Encode(newf, img); err!= nil {
			return &ConvError{Err: err, Code: FileEncodeFail, FilePath: newfn}
		}
		fmt.Println("Created", newfn)
	}
	return nil
}

func fnWithoutExt(fn string) string {
	return fn[:len(fn)-len(filepath.Ext(fn))]
}
