package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048) // 2048byte씩 읽을 수 있는 버퍼
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}

		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func main() { // main 고루틴
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(counts)

	////////////////////////////////////////////////

	r, closer, err := buildGZipReader("my_data.txt.gz")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer closer()
	counts, err = countLetters(r)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(counts)

}
