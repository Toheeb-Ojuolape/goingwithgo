package main

import (
	"bufio"
	"fmt"
	f "fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	pl("What is your age?")
	reader := bufio.NewReader(os.Stdin)
	age, err := reader.ReadString('\n')
	if err == nil {
		age = strings.TrimSpace(age)
		iAge, err := strconv.Atoi(age)
		if err != nil {
			log.Fatal(err)
		}
		if iAge < 5 {
			fmt.Println("Too young for school")
		}
		if iAge == 5 {
			fmt.Println("Go to Kindergarten")
		}

		if (iAge > 5) && (iAge <= 17) {
			grade := iAge - 5
			fmt.Printf("Go to grade  %d \n", grade)
		}

		if iAge > 17 {
			fmt.Println("Go to college")
		}

	} else {
		log.Fatal(err)
	}

}
