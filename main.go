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
	wordPtr := flag.String("csv", "problems.csv", "A CSV file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*wordPtr)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	problemCounter := -1;
	correctAnswerCounter := 0;

	
	for {
		problemCounter++

		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

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
