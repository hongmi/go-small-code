package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	//"time"
	"errors"
	"sort"
	"sync"
)

type result struct {
	path string
	sum  string
	err  error
}

func main() {

	dirPath := os.Args[1]

	m, err := MD5All(dirPath)

	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	var paths []string
	for k, _ := range m {
		paths = append(paths, k)
	}

	sort.Strings(paths)
	for _, p := range paths {
		fmt.Printf("%s %s\n", m[p], p)
	}
	//	time.Sleep(time.Second)

}

func MD5All(root string) (map[string]string, error) {

	m := make(map[string]string)

	done := make(chan struct{})

	paths, errc := WalkFiles(done, root)

	rc := digestFiles(done, paths)

	for v := range rc {
		if v.err != nil {
			//fmt.Printf("digest file error %v", err)
			return nil, v.err
		}
		m[v.path] = v.sum
	}

	if err := <-errc; err != nil {
		//fmt.Printf("%s\n", err)
		return nil, err
	}

	return m, nil
}

func WalkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		defer close(paths)

		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled!")
			}
			return nil
		})
	}()

	return paths, errc
}

func digestFiles(done <-chan struct{}, paths <-chan string) <-chan result {
	kvc := make(chan result)

	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for path := range paths {
				data, err := ioutil.ReadFile(path)

				if err != nil {
					return
				}

				md5bytes := md5.Sum(data)
				md5string := hex.EncodeToString(md5bytes[:])

				select {
				case kvc <- result{path, md5string, err}:
				case <-done:
					return
				}
			}

		}()
	}

	go func() {
		wg.Wait()
		close(kvc)
	}()

	return kvc
}
