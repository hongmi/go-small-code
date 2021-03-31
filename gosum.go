package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	dirPath := os.Args[1]

	var m map[string]string

	m = make(map[string]string)

	m, err := MD5All(dirPath)
	if err != nil {
		fmt.Errorf("fail to MD5All")
		return
	}

	for k, v := range m {
		fmt.Printf("%s %s\n", k, v)
	}

}

func MD5All(root string) (map[string]string, error) {

	var m map[string]string
	m = make(map[string]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		data, err := ioutil.ReadFile(path)

		if err != nil {
			return err
		}

		md5bytes := md5.Sum(data)
		md5string := hex.EncodeToString(md5bytes[:])
		m[path] = md5string
		return nil
	})

	return m, err
}
