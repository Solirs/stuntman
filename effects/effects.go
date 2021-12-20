package effects

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	inc "github.com/Solirs/stuntman/include"
)

//Randbin: DEPRECATED
/*func Randbin() {

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
}*/

//Binary displays the Binary effect
func Binary() {

	inc.Charlist = inc.Binary_list

	for m := 0; m < inc.Spaces; m++ {

		inc.Charlist = append(inc.Charlist, " ")

	}
	for {
		RandString()

	}

}

//Texteff loops the Randstring function
func Texteff() {

	if inc.Loweronly {
		inc.Charlist = inc.ASCIILower
	} else if inc.Upperonly {
		inc.Charlist = inc.ASCIIUpper
	}else{
		inc.Charlist = inc.ASCII
	}

	for m := 0; m < inc.Spaces; m++ {

		inc.Charlist = append(inc.Charlist, " ")

	}

	for {
		RandString()
	}

}

//RandString generates a line according to inc.Charlist and displays it, its sort of the engine of stuntman
func RandString() {

	rand.Seed(time.Now().UnixNano())

	j := inc.Cnt

	var Strdrop = make([]string, j)

	for i := range Strdrop {

		Strdrop[i] = inc.Charlist[rand.Intn(len(inc.Charlist))]

	}
	fmt.Println(strings.Trim(fmt.Sprint(Strdrop), "[]"))
	time.Sleep(time.Duration(inc.Speed) * time.Millisecond) //Sleep the amount of time specified in the speed flag
	for i := 0; i < inc.Vgap; i++ {

		fmt.Println() //Print an empty line the amount of times specified in the vgap flag

	}

}

func Init_Custom() {

	inc.Charlist = strings.Split(inc.Custom_string, ",")

	for m := 0; m < inc.Spaces; m++ {

		inc.Charlist = append(inc.Charlist, " ")

	}

	for {
		RandString()
	}

}
