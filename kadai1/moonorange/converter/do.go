package converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func Do(dstDir, fromExt, toExt, path, format string, img image.Image) error {
	if err := os.MkdirAll(dstDir, 0777); err != nil {
		return err
	}
	newfn := filepath.Join(dstDir, filepath.Base(fnWithoutExt(path) + "." + toExt))
	newf, err := os.Create(newfn)
	if err != nil {
		return err
	}

	defer func () {
		if err := newf.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	switch toExt {
	case "jpg":
		if err := jpeg.Encode(newf, img,  &jpeg.Options{Quality: 75}); err != nil {
			return err
		}
	case "png":
		err := png.Encode(newf, img)
		if err != nil {
			return err
		}
	}
	return nil
}

func fnWithoutExt(fn string) string {
	return fn[:len(fn)-len(filepath.Ext(fn))]
}
