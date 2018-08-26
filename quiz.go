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
	sheet := loadProblems("files/problems.csv")
}

func roundTimer() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()
	time.Sleep(10 * time.Second)
	ticker.Stop()
	fmt.Println("Time Up!")
}

func loadProblems(filepath string) {
	var scoresheet []bool
	data, err := ioutil.ReadFile(filepath)
	checkErr(err)

	//fmt.Printf("DATA: %v\n", string(data))

	go roundTimer()
	reader := csv.NewReader(strings.NewReader(string(data)))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		scoresheet = append(scoresheet, printQuestion(record[0], record[1]))
	}
	answ, quest := calcScore(scoresheet)
	fmt.Printf("You scored: %d/%d\n", answ, quest)
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
