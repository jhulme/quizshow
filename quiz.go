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
	var scoresheet []bool

	reader := loadProblems("files/problems.csv")
	timer := time.NewTimer(10 * time.Second)

	fmt.Println("Welcome to Quiz Show!")

	for {
		select {
		case <-timer.C:
			answ, quest := calcScore(scoresheet)
			fmt.Printf("Time Expired. You scored: %d/%d\n", answ, quest)
			return
		default:
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			scoresheet = append(scoresheet, printQuestion(record[0], record[1]))
		}
	}
	answ, quest := calcScore(scoresheet)
	fmt.Printf("You scored: %d/%d\n", answ, quest)
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
