package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in go lang")

	rand.Seed(time.Now().Unix())

	diceNumber := rand.Intn(6)+1
	fmt.Println("value of dice",diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1 you can open")
	case 2:
		fmt.Println("you can move 2 spots")
	case 3:
		fmt.Println("you can move 3 spots")
		fallthrough //meka dammama mekan yata case ekath niknm run wenwa ma case eka hri giyama
	case 4:
		fmt.Println("you can move 4 spots")

	case 5:
		fmt.Println("you can move 5 spots")
	default:
		fmt.Println("defulat runs")
	}
}
