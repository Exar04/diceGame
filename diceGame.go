package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) +1
	fmt.Println("The value of dice is ", diceNumber)
}