package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	//"time"
)

func main() {

	dirPath := os.Args[1]

	paths := MD5All(dirPath)

	m := digestFiles(paths)

	//	time.Sleep(time.Second)
	for v := range m {
		fmt.Printf("%s\n", v)
	}

}

func MD5All(root string) <-chan string {
	paths := make(chan string)
	go func() {
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			paths <- path

			return nil
		})
		close(paths)
	}()

	return paths
}

func digestFiles(paths <-chan string) <-chan string {

	kvc := make(chan string)

	go func() {
		for path := range paths {
			data, err := ioutil.ReadFile(path)

			if err != nil {
				return
			}

			md5bytes := md5.Sum(data)
			md5string := hex.EncodeToString(md5bytes[:])
			kvc <- fmt.Sprintf("%s %s", path, md5string)
		}
		close(kvc)
	}()
	return kvc
}
