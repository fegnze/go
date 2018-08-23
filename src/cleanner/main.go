package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func clearDir(path string, do func(path string)) {
	var pathArr []string
	i := 0
	err := filepath.Walk(path, func(path1 string, f os.FileInfo, err0 error) error {
		if f == nil || err0 != nil {
			return err0
		}

		if path != path1 {
			pathArr = append(pathArr, path1)
			i++
		}
		return nil
	})

	if err == nil {
		if do == nil {
			do = func(path string) {
				fi, err0 := os.Stat(path)
				if path != "./" && !os.IsNotExist(err0) {
					if time.Now().Unix()-fi.ModTime().Unix() > 60*60*24*2 {
						fmt.Println("删除 ... ", path, fi.ModTime().Unix())
						os.RemoveAll(path)
					}
				}
			}
		}
		for _, s := range pathArr {
			do(s)
		}
		fmt.Println("结束")
	} else {
		fmt.Println(err)
	}

}

func main() {
	clearDir("./", nil)
}
