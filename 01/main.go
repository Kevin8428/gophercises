package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func getArgs() map[string]string {
	m := map[string]string{}
	for _, v := range os.Args {
		if strings.Contains(v, "file=") {
			m["file"] = strings.SplitAfter(v, "file=")[1]
		} else if strings.Contains(v, "time=") {
			m["time"] = strings.SplitAfter(v, "time=")[1]
		}
	}
	return m
}

func main() {
	m := getArgs()
	var f string
	if val, ok := m["file"]; ok {
		f = val
	} else {
		f = "one.csv"
	}

	file, _ := os.Open(f)
	r := csv.NewReader(bufio.NewReader(file))
	reader := bufio.NewReader(os.Stdin)
	correct := 0
	incorrect := 0
	fmt.Print("hit enter to start timer")
	reader.ReadString('\n')
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("\ncorrect: ", correct)
		fmt.Println("incorrect: ", incorrect)
		os.Exit(0)
	}()
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
