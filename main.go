package main

import (
)

import (
	"os"
	"encoding/csv"
	"fmt"
	"log"
	"io"
	"strings"
	"bufio"
)

func main() {
	file, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	r.Comma = ';'
	problemCounter := 1;

	for {
		rawRecord, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		record := strings.Split(rawRecord[0], ",")

		fmt.Printf("Problem #%v: %v = ", problemCounter, record[0])

		reader := bufio.NewReader(os.Stdin)
		rawText, _ := reader.ReadString('\n')
		text := strings.Replace(rawText, "\n", "", -1)

		fmt.Println(text)

		problemCounter++
	}
}
