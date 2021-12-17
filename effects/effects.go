package effects

import(
	"time"
	"math/rand"
	"strings"
	"fmt"
	inc "stuntman/include")

func Randbin() {

	rand.Seed(time.Now().UnixNano())

	j := inc.Cnt

	for i := 0; i < j; i++ {
		inc.Arrl[i] = rand.Intn(2) //Generate a number between 1 and 0 for every num in the range of the width flag and add it to the array to be printed
	}
}

func Binary() {
	for  {
		Randbin()


		fmt.Println(strings.Trim(fmt.Sprint(inc.Arrl), "[]")) //Remove the stupid brackets when you print an array and print it
		time.Sleep(time.Duration(inc.Speed)* time.Millisecond) //Sleep the amount of time specified in the speed flag
			for i := 0; i < inc.Vgap; i++ {

				fmt.Println() //Print an empty line the amount of times specified in the vgap flag

			}
		

	}

}

func Texteff(){

	for {
		RandString()
	}

	

}


func RandString(){

	rand.Seed(time.Now().UnixNano())

	j := inc.Cnt

	var Strdrop = make([]rune, j)

	for i := range Strdrop{

		Strdrop[i] = inc.Ascii[rand.Intn(len(inc.Ascii))]
		


	}
	fmt.Println(string(Strdrop))
	time.Sleep(time.Duration(inc.Speed)* time.Millisecond) //Sleep the amount of time specified in the speed flag
		for i := 0; i < inc.Vgap; i++ {

			fmt.Println() //Print an empty line the amount of times specified in the vgap flag

		}
	

}