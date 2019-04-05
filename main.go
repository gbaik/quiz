/*

- Parse CSV file (problems.csv) and return on command
- Take input from user
- Check to see if answer is correct next to CSV, and update score
- Return next question in CSV immediately
- At end of quiz return how many correct questions 
	- Invalid answers = incorrect
- Be able to take flag for the filename

*/

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

		fmt.Printf("Problem #%v: %v \n", problemCounter, record[0]);

		problemCounter++
	}
}
