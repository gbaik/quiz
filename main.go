/*
	- Add default timer for program
	- Add flag to accept a time
*/

package main

import (
	"flag"
	"os"
	"encoding/csv"
	"fmt"
	"log"
	"time"
	"io"
	"strings"
	"bufio"
)

const seconds = 5

func main() {
	wordPtr := flag.String("csv", "problems.csv", "A CSV file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*wordPtr)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	problemCounter := 1;
	correctAnswerCounter := 0;

	c := make(chan int)
	c2 := make(chan string)

	go processCSVFile(r, problemCounter, correctAnswerCounter, c, c2)

	for currentProblem := range c {
		record := <- c2

		fmt.Printf("Problem #%v: %v = ", currentProblem, record)
	}
}

func NewTimer(seconds int, action func()) *time.Timer {
	timer := time.NewTimer(time.Second * time.Duration(seconds))

	go func() {
		<-timer.C
		action()
	}()

	return timer
}
 
func processCSVFile(r *csv.Reader, problemCounter int, correctAnswerCounter int, c chan int, c2 chan string) {
	NewTimer(seconds, func() {
		fmt.Printf("You scored %v out of %v", correctAnswerCounter, problemCounter)

		close(c)
	})

	for {
		record, err := r.Read()

		if err == io.EOF {
			problemCounter--

			fmt.Printf("You scored %v out of %v", correctAnswerCounter, problemCounter)
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		c <- problemCounter
		c2 <- record[0]

		reader := bufio.NewReader(os.Stdin)
		rawText, _ := reader.ReadString('\n')
		text := strings.Replace(rawText, "\n", "", -1)

		if strings.Compare(record[1], text) == 0 {
      correctAnswerCounter++
		}
		
		problemCounter++
	}

	close(c)
}