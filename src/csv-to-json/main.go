package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type candidate struct {
	name    string
	gender  string
	birth   string
	job     string
	company string
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func CsvToMap(filename string) ([]candidate, error) {
	// init variables
	var (
		candidates []candidate
		err        error
	)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return candidates, err
	}
	defer Close(f)
	// read buffer
	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')

	for err == nil {
		s := strings.Split(line, ",")

		if s != nil {
			candidates = append(candidates, candidate{
				name:    s[0],
				gender:  s[1],
				birth:   s[2],
				job:     s[3],
				company: s[4],
			})
		}
		line, err = r.ReadString('\n')
	}

	if err != io.EOF {
		fmt.Println(err)
		return candidates, err
	}

	return candidates, nil
}

func main() {
	fmt.Println("hihi")
	result, err := CsvToMap("data.csv")
	if result != nil {
		fmt.Println("======result=======")
		fmt.Println(result)
	}
	if err != nil {
		fmt.Println("=======err=========")
		fmt.Println(err)
	}
}
