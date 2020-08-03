package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.google.at")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//body := make([]byte, 99999)
	//l, _ := resp.Body.Read(body)
	//fmt.Println("length: ", l)
	//fmt.Println(string(body))

	t := testWriter{}
	//l, err := t.Write([]byte("qqwvqwqvwqqwwq"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("length:", l)
	io.Copy(t, resp.Body)
}

type testWriter struct{}

func (t testWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	fmt.Println("bytes written; ", len(p))
	return len(p), nil
}
