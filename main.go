package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")

	user.Age = readInt("How old are you?")

	user.FavouriteNumber = readFloat("What is your favourite number?")

	fmt.Println(fmt.Sprintf("Your name is %s and you are %d years old. Your favourite unmber is %.2f",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
	))
}

func prompt() {
	fmt.Print("->")
}

func readString(s string) string {
	fmt.Println(s)

	prompt()

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}

func readInt(s string) int {
	for {

		fmt.Println((s))

		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {

		fmt.Println((s))

		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}
