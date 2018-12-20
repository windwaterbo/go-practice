package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Candidate struct {
	Name    string
	Gender  string
	Birth   string
	Job     string
	Company string
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func CsvToStruct(filename string) ([]Candidate, error) {
	// init variables
	var (
		candidates []Candidate
		cand       Candidate
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
			cand.Name = s[0]
			cand.Gender = s[1]
			cand.Birth = s[2]
			cand.Job = s[3]
			cand.Company = s[4]
			candidates = append(candidates, cand)
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
	m, _ := CsvToStruct("data.csv")
	fmt.Println(m)
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
