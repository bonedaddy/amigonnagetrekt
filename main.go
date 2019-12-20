package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	dirPath = "contracts"
)

func main() {
	files, err := ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	start := time.Now()
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		ff, err := os.Open(dirPath + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		data, err := ioutil.ReadAll(ff)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		ff.Close()
		_ = data
	}
	end := time.Now()
	diff := end.Sub(start)
	fmt.Println("time :", diff.Milliseconds())
}

// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}