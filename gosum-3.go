package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	//"time"
	"sync"
)

func main() {

	dirPath := os.Args[1]

	m := MD5All(dirPath)

	//	time.Sleep(time.Second)
	for v := range m {
		fmt.Printf("%s\n", v)
	}

}

func MD5All(root string) <-chan string {
	c := make(chan string)

	go func() {
		var wg sync.WaitGroup
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			go func() {
				wg.Add(1)
				data, err := ioutil.ReadFile(path)

				if err != nil {
					return
				}

				md5bytes := md5.Sum(data)
				md5string := hex.EncodeToString(md5bytes[:])
				c <- fmt.Sprintf("%s %s", path, md5string)
				wg.Done()
				fmt.Printf("sum path:%s\n", path)
			}()

			return nil
		})

		go func() {
			wg.Wait()
			close(c)
		}()
	}()

	return c
}
