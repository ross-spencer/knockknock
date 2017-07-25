// knock knock
// A programme that asks for a knock knock joke and responds to user input with derision.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var f *os.File

func init() {
	var err error
	f, err = os.Create("knockknock.log")
	checkerror(err)
}

func checkerror(err error) bool {
	if err != nil {
		log.Println("FYI there is an error!", err)
		return true
	}
	return false
}



func knock() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hey, can you please tell me a knock knock joke?", "\n", "Start by typing knock knock or type Q at anytime to exit")

	text, err := reader.ReadString('\n')
	checkerror(err)

	text = strings.TrimSpace(text)

	whatnow := time.Now().UnixNano()

	rand.Seed(whatnow)

	answers := []string{
		"That's not funny.",
		"Hmm, keep practicing.",
		"Oh, oh dear.",
		"Seriously?!",
		"That really needs some work.",
		"I see what you did there.", "Ha.",
		"Is, is that it?",
		"That's not bad.",
		"I don't get it.",
		"What's the joke?",
		"Wait, you think that's funny?!",
		"No.",
		".....",
	}

	if text == "knock knock" {
		fmt.Println("Who's there?")

		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)

		_, err := f.Write([]byte(text + "\n"))
		checkerror(err)

		if strings.ToUpper(text) == "Q" {
			os.Exit(0)
		} else {
			fmt.Println(text, "who?")
		}

		text, _ = reader.ReadString('\n')

		_, err = f.Write([]byte(text + "\n"))
		checkerror(err)		

		text = strings.TrimSpace(text)

		if strings.ToUpper(text) == "Q" {
			os.Exit(0)
		} else {
			fmt.Println(text, "?", answers[rand.Intn(len(answers))])
		}

	} else if strings.ToUpper(text) == "Q" {
		os.Exit(0)

	} else {
		fmt.Println("I don't understand!", "\n", "Let's try that again:")
		knock()

	}

	f.Close()

}

func main() {

	knock()
}
