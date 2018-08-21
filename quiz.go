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
)

func main() {
	loadProblems("files/problems.csv")
}

func loadProblems(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	checkErr(err)

	//fmt.Printf("DATA: %v\n", string(data))

	reader := csv.NewReader(strings.NewReader(string(data)))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(record[0])
		printQuestion(record[0], record[1])
	}

}

func printQuestion(question, answer string) {
	fmt.Printf("What is %s?\n", question)
	input := bufio.NewReader(os.Stdin)
	text, _ := input.ReadString('\n')

	if strings.Compare(text, answer) == 0 {
		fmt.Print("Correct Answer!")
	} else {
		fmt.Print("Incorrect Answer!")
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
