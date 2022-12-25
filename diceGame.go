package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	diceNumber := rollDice()
	fmt.Println("The value of dice is ", diceNumber)
}

func rollDice() int{
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) +1

	return diceNumber
}