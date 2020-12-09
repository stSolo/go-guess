package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//get now date int format
	now := time.Now().Unix()
	rand.Seed(now)

	//generate random num
	randNum := rand.Intn(100) + 1

	//create buffer reader
	reader := bufio.NewReader(os.Stdin)

	success := false

	fmt.Println("I have random number. Can you guess it?")

	//lifecycle loop
	for lives := 0; lives < 10; lives++ {
		fmt.Println("CHOOSE YOU NUMBER")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		//parse int from string
		input = strings.TrimSpace(input)
		userNum, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}


		//check user number
		if userNum == randNum {
			success = true
			fmt.Println("YOU WIN")
			break
		} else if userNum > randNum {
			fmt.Println("Your number is greater than random number. Pls try again")
		} else {
			fmt.Println("Your number is lesser than random number. Pls try again")
		}

		fmt.Println(userNum)
	}

	if !success {
		fmt.Println("YOU LOSE")
	}

}
