package main

import (
	"fmt"
	"encoding/csv"
	"io"
	"os"
	"log"
	"strconv"
	"time"
	"flag"
)

func main() {
	timePtr := flag.Int("time", 10, "Timer, default 10 seconds.")
	flag.Parse()
	fmt.Println("Hello!")

	f, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := csv.NewReader(f)
	
	var questions []string
	var answers []int

	// Initalizing questions
	for{
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		questions = append(questions, record[0])
		ans, err := strconv.Atoi(record[1])
		if err != nil{
			break
		}
		answers = append(answers, ans)
	}

	correct := 0
	
	timer := time.NewTimer(time.Duration(*timePtr) * time.Second)
	go func() {
		<-timer.C
			fmt.Println("Time's up!")
			fmt.Printf("You got %d/%d correct!", correct, len(questions))
			fmt.Println("")
			os.Exit(0)
	}()
		problemLoop:	
		for i, p := range questions{
			fmt.Printf("Problem#%d: %s\n", i+1, p)

			answerCh := make(chan int)

			go func() {
				var answer int 
				fmt.Scanf("%d\n", &answer)
				answerCh <- answer
			}()

			select{
				case <- timer.C:
					fmt.Println("\nTime's up!")
					break problemLoop
				case answer := <- answerCh:
					if answer == answers[i] {
						correct++
					}
			}
	}

	fmt.Printf("\nYou got %d/%d correct!", correct, len(questions))
	fmt.Println("")
}

