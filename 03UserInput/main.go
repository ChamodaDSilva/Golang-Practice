package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || error ok syntax
	input, _ := reader.ReadString('\n')// mulin thiyena variable ekata initailize wenne, error ekk awoth ignore krnn _ danne, error eka mokkd balathaki ethanata varable ekk dala print krla onenm
	//input, err := reader.ReadString('\n')


	fmt.Println("Thanks for reading, " + input)
	fmt.Printf("type of input %T ", input)

}
