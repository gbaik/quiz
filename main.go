package main

import (
	"flag"
	"os"
	"encoding/csv"
	"fmt"
	"log"
	"io"
	"strings"
	"bufio"
)

func main() {
	wordPtr := flag.String("csv", "problems.csv", "Please enter the CSV to parse")
	flag.Parse()g

	file, err := os.Open(*wordPtr)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	r.Comma = ';'
	problemCounter := -1;
	correctAnswerCounter := 0;

	for {
		problemCounter++

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

		if strings.Compare(record[1], text) == 0 {
      correctAnswerCounter++
    }
  }

  fmt.Printf("You scored %v out of %v", correctAnswerCounter, problemCounter)
}
