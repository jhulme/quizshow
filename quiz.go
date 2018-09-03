package main

import (
	"encoding/csv"
	//"flags"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var signal chan bool
	var scoresheet []bool

	signal = make(chan bool)
	reader := loadProblems("files/problems.csv")
	go roundTimer(signal)

	fmt.Println("Welcome to Quiz Show!")

	for {
		record, err := reader.Read()
		if err == io.EOF {
			close(signal)
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		stop := <-signal
		scoresheet = append(scoresheet, printQuestion(record[0], record[1]))
		if stop == false {
			fmt.Println("Time Up!\n")
			break
		}
	}
	answ, quest := calcScore(scoresheet)
	fmt.Printf("You scored: %d/%d\n", answ, quest)
}

func roundTimer(signal chan bool) {
	time.NewTimer(10 * time.Second)
	// go func() {
	// 	for t := range ticker.C {
	// 		fmt.Println(t)
	// 	}
	// }()
	// time.After(10 * time.Second)
	// ticker.Stop()
	signal <- false
}

func loadProblems(filepath string) *csv.Reader {
	data, err := ioutil.ReadFile(filepath)
	checkErr(err)

	reader := csv.NewReader(strings.NewReader(string(data)))

	return reader
}

func printQuestion(question, answer string) (result bool) {
	fmt.Printf("What is %s?\n", question)
	input := bufio.NewReader(os.Stdin)
	text, _ := input.ReadString('\n')

	if strings.TrimRight(text, "\n") == answer {
		fmt.Print("Correct Answer!\n")
		return true
	} else {
		fmt.Printf("Incorrect Answer! Expecting %s\n", answer)
		return false
	}
}

func calcScore(scores []bool) (score, questions int) {
	//use array to log either correct or incorrect - count number of corrects at end
	var correct []bool
	questions = len(scores)
	for _, v := range scores {
		if v == true {
			correct = append(correct, v)
		}
	}
	score = len(correct)
	return score, questions
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
