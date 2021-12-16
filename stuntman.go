package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)
var cnt int
var vgap int
var speed int
var clr string
var binar bool
var arrl []int

var colr string

func initflags(){
	flag.IntVar(&cnt, "width", 50, "The width of the cascade in rows")
	flag.IntVar(&vgap, "vertgap", 0, "Gap between lines")
	flag.IntVar(&speed, "speed", 0, "Time between lines in milliseconds")
	flag.StringVar(&clr, "color", "nil", "The color of the cascade")
	flag.BoolVar(&binar, "binar", false, "Display random binary code")
	


}

func randbin(cnt int) {

	rand.Seed(time.Now().UnixNano())

	j := *&cnt

	for i := 0; i < j; i++ {
		arrl[i] = rand.Intn(2) //Generate a number between 1 and 0 for every num in the range of the width flag and add it to the array to be printed
	}
}

func binary() {
	for true {
		randbin(cnt)

		fmt.Println(strings.Trim(fmt.Sprint(arrl), "[]")) //Remove the stupid brackets when you print an array and print it
		time.Sleep(time.Duration(speed)* time.Millisecond) //Sleep the amount of time specified in the speed flag
		if vgap < 0 {
			for i := 0; i < vgap; i++ {

				fmt.Println() //Print an empty line the amount of times specified in the vgap flag

			}
		}

	}

}

func main() {
	


	initflags()
	flag.Parse()

	switch clr {

	case "red":
		colr = Red

	case "green":
		colr = Green

	case "blue":
		colr = Blue

	case "purple":
		colr = Purple

	case "yellow":
		colr = Yellow

	case "cyan":
		colr = Cyan

	case "white":
		colr = White

		default: colr = Reset
	}

	arrl = make([]int, cnt) //Make array of non constant length

	fmt.Println(colr)

	if binar {
		binary()
	} else {
		fmt.Println("No mode selected, QUITTING!!")
		os.Exit(1)

	}

}
