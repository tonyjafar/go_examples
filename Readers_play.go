package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("C:/Users/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for x := 0; ; x++ {
		l, err := r.Peek(x)
		if err == io.EOF {
			if strings.Contains(string(l), "Testing") {
				new := strings.Replace(string(l), "Testing", "NewString", 10000)
				n, err := os.Create("C:/Users/test3.txt")
				if err != nil {
					log.Fatal(err)
				}
				defer n.Close()
				io.Copy(n, strings.NewReader(new))
			}
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open("C:/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	n, err := os.Create("C:/Users/test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer n.Close()
	io.Copy(n, bufio.NewReader(f))
}
