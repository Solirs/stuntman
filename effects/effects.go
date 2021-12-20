package effects

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	inc "github.com/Solirs/stuntman/include"
)

//Randbin generates random array of bits
func Randbin() {

	rand.Seed(time.Now().UnixNano())

	j := inc.Cnt

	var Strdrop = make([]string, j)

	for i := range Strdrop {

		Strdrop[i] = inc.Binary_list[rand.Intn(len(inc.Binary_list))]

	}
	fmt.Println(strings.Trim(fmt.Sprint(Strdrop), "[]"))
	time.Sleep(time.Duration(inc.Speed) * time.Millisecond) //Sleep the amount of time specified in the speed flag
	for i := 0; i < inc.Vgap; i++ {

		fmt.Println() //Print an empty line the amount of times specified in the vgap flag

	}
}

//Binary displays the Binary effect
func Binary() {

	for m := 0; m < inc.Spaces; m++ {

		inc.Binary_list = append(inc.Binary_list, " ")

	}
	for {
		Randbin()

	}

}

//Texteff loops the Randstring function
func Texteff() {

	if inc.Loweronly {
		inc.ASCII = inc.ASCIILower
	} else if inc.Upperonly {
		inc.ASCII = inc.ASCIIUpper
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
