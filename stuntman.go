package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var cnt = flag.Int("width", 50, "The width of the cascade in rows")
var clr = flag.String("color", "nil", "The color of the cascade")

var arrl []int

var colr string

func randarr(cnt int) {

	rand.Seed(time.Now().UnixNano())

	j := *&cnt

	for i := 0; i < j; i++ {
		arrl[i] = rand.Intn(2)
	}
}

func main() {
	flag.Parse()

	switch *clr{

	case "red": colr = Red

	case "green": colr = Green
		
	case "blue": colr = Blue

	case "purple": colr = Purple

	case "yellow": colr = Yellow

	case "cyan": colr = Cyan

	case "white": colr = White


		default: colr = Reset
	}

	arrl = make([]int, *cnt)

	fmt.Println(colr)

	for true {
		randarr(*cnt)

		fmt.Println(strings.Trim(fmt.Sprint(arrl), "[]"))
	}

}
