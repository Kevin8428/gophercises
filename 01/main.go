package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("one.csv")
	r := csv.NewReader(bufio.NewReader(file))
	reader := bufio.NewReader(os.Stdin)
	correct := 0
	incorrect := 0
	for {
		line, error := r.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println("err")
			log.Fatal(error)
		}
		question := line[0]
		answer := line[1]
		fmt.Print(question + ":")
		userAnswer, _ := reader.ReadString('\n')
		if strings.TrimSpace(answer) == strings.TrimSpace(userAnswer) {
			correct++
		} else {
			incorrect++
		}
	}
	fmt.Println("correct: ", correct)
	fmt.Println("incorrect: ", incorrect)
}
