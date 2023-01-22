package main

import (
	"bufio"
	"fmt"
	f "fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var pl = f.Println

func main() {
	// pl("Hello Go")
	// pl("What is your name> ")
	// reader := bufio.NewReader(os.Stdin)
	// name, err := reader.ReadString('\n')
	// if err == nil {
	// 	pl("Hello", name)
	// } else {
	// 	log.Fatal(err)
	// }

	// %d : integer
	// %c : character
	// %f : float
	// %t : boolean
	// %s : string
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value
	// fmt.Printf("%9f\n", 3.14)
	// fmt.Printf("%.2f\n", 3.141592)
	// fmt.Printf("%.9f\n", 3.141592)
	// sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	// pl(sp1)

	// fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)

	// pl("What is your age?")
	// reader := bufio.NewReader(os.Stdin)
	// age, err := reader.ReadString('\n')
	// if err == nil {
	// 	age = strings.TrimSpace(age)
	// 	iAge, err := strconv.Atoi(age)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if iAge < 5 {
	// 		fmt.Println("Too young for school")
	// 	}
	// 	if iAge == 5 {
	// 		fmt.Println("Go to Kindergarten")
	// 	}

	// 	if (iAge > 5) && (iAge <= 17) {
	// 		grade := iAge - 5
	// 		fmt.Printf("Go to grade  %d \n", grade)
	// 	}

	// 	if iAge > 17 {
	// 		fmt.Println("Go to college")
	// 	}

	// } else {
	// 	log.Fatal(err)
	// }

	// pl("Enter number 1")
	// reader := bufio.NewReader(os.Stdin)
	// num1, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// numString := strings.TrimSpace(num1)
	// number1, err := strconv.Atoi(numString)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pl("Enter number 2")
	// num2, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// num2String := strings.TrimSpace(num2)
	// number2, err := strconv.Atoi(num2String)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("The sum of %d and %d is %d", number1, number2, number1+number2)

	// for x := 1; x < 5; x++ {
	// 	pl(x)
	// }

	// fx := 0
	// for fx <= 5 {
	// 	pl(fx)
	// 	fx++
	// }

	// aNums := []int{1, 2, 3}
	// for _, num := range aNums {
	// 	pl(num)
	// }

	// xVal := 1
	// for true {
	// 	if xVal == 5 {
	// 		break
	// 	}
	// 	pl(xVal)
	// 	xVal++
	// }

	reader := bufio.NewReader(os.Stdin)
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1

	for true {
		fmt.Println("Guess a number between 1 and 50")
		pl("Random number is : ", randNum)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strConv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			fmt.Println("Guess something lower")
		}

		if iGuess < randNum {
			fmt.Println("Guess something higher")
		}
		if iGuess == randNum {
			fmt.Println("YOU guessed it!")
			break
		}
	}

}
