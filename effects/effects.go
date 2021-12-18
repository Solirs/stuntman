package effects

import (
	"fmt"
	inc "github.com/Solirs/stuntman/include"
	"math/rand"
	"strings"
	"time"
)

//Randbin generates random array of bits
func Randbin() {

	rand.Seed(time.Now().UnixNano())

	j := inc.Cnt

	for i := 0; i < j; i++ {
		inc.Arrl[i] = rand.Intn(2) //Generate a number between 1 and 0 for every num in the range of the width flag and add it to the array to be printed
	}
}

//Binary displays the Binary effect
func Binary() {
	for {
		Randbin()

		fmt.Println(strings.Trim(fmt.Sprint(inc.Arrl), "[]"))   //Remove the stupid brackets when you print an array and print it
		time.Sleep(time.Duration(inc.Speed) * time.Millisecond) //Sleep the amount of time specified in the speed flag
		for i := 0; i < inc.Vgap; i++ {

			fmt.Println() //Print an empty line the amount of times specified in the vgap flag

		}

	}

}

//Texteff loops the Randstring function
func Texteff() {

	if inc.Loweronly {
		inc.ASCII = inc.ASCII_Lower
	} else if inc.Upperonly {
		inc.ASCII = inc.ASCII_Upper
	}

	for m := 0; m < inc.Spaces; m++ {

		inc.ASCII = append(inc.ASCII, " ")

	}

	for {
		RandString()
	}

}

//RandString creates a random string
func RandString() {

	rand.Seed(time.Now().UnixNano())

	j := inc.Cnt

	var Strdrop = make([]string, j)

	for i := range Strdrop {

		Strdrop[i] = inc.ASCII[rand.Intn(len(inc.ASCII))]

	}
	fmt.Println(strings.Trim(fmt.Sprint(Strdrop), "[]"))
	time.Sleep(time.Duration(inc.Speed) * time.Millisecond) //Sleep the amount of time specified in the speed flag
	for i := 0; i < inc.Vgap; i++ {

		fmt.Println() //Print an empty line the amount of times specified in the vgap flag

	}

}
