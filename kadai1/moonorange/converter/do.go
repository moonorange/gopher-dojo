package converter

import (
	"fmt"
)


func Do(srcDir, dstDir, fromExt, toExt, path, format string) error {
	fmt.Println(srcDir, dstDir, fromExt, toExt, path, format)
	return nil
}
