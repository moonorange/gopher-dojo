package converter

import (
	"fmt"
)


func Do(srcDir, dstDir, fromExt, toExt, path string) error {
	fmt.Println(srcDir, dstDir, fromExt, toExt, path)
	return nil
}
