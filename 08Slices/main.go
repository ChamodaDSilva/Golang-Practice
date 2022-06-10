package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to vieo on slices")

	var fruitlist = []string{"apple", "peach"}
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "mango")
	fmt.Println(fruitlist)

	fruitlist=append(fruitlist[1:] )//can have fruitlist[:3] too
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "pineapple","tomatoe")
	fmt.Println(fruitlist)

	fruitlist=append(fruitlist[1:3] )
	fmt.Println(fruitlist)


	///////////////////////////////

	highScores :=make([]int,4)
	highScores[0]=234
	highScores[1]=212
	highScores[2]=231
	highScores[3]=214

	highScores = append(highScores, 121,1324,12131,1212)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

}
