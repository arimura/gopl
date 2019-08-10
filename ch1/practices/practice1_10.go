package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	filePath = "./fetched_urls/"
)

func check(e error) {
	if e != nil {
		panic(2)
	}
}

func main() {
	//get options
	options := os.Args[1:]

	//make base dir
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, os.ModePerm)
	}

	for _, url := range options {
		tmp := strings.Split(url, "://")[1]
		p := filePath + strings.Split(tmp, "?")[0]

		if _, err := os.Stat(p); os.IsNotExist(err) {
			os.Mkdir(p, os.ModePerm)
		}

		write(p+"/1st", fetch(url))
		write(p+"/2nd", fetch(url))
	}
}

func fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		bodyString := string(bodyBytes)

		return bodyString
	}
	return ""
}

func write(path, content string) {
	f, err1 := os.Create(path)
	if err1 != nil {
		panic(err1)
	}

	w := bufio.NewWriter(f)
	_, err2 := w.WriteString(content)
	if err2 != nil {
		log.Fatal(err2)
		return
	}
	w.Flush()
}
