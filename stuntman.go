package main

import (
	"flag"
	"fmt"
	eff "github.com/Solirs/stuntman/effects"
	inc "github.com/Solirs/stuntman/include"
	"os"
)

func initflags() {
	flag.IntVar(&inc.Cnt, "width", 50, "The width of the cascade in rows")
	flag.IntVar(&inc.Vgap, "linegap", 0, "Gap between lines")
	flag.IntVar(&inc.Spaces, "spaces", 0, "An index to determine how many spaces will be generated when using the -text mode, the higher it is, the more spaces.")
	flag.IntVar(&inc.Speed, "speed", 0, "Time between lines in milliseconds")
	flag.StringVar(&inc.Clr, "color", "nil", "The color of the cascade")
	flag.StringVar(&inc.Custom_string, "custom", "nil", "Use a custom set of characters")
	flag.BoolVar(&inc.Binar, "binar", false, "Display random Binary code")
	flag.BoolVar(&inc.Text, "text", false, "Display random text characters")
	flag.BoolVar(&inc.V, "version", false, "Display version and additional info")
	flag.BoolVar(&inc.Loweronly, "loweronly", false, "Only use lowercase text in -text mode.")
	flag.BoolVar(&inc.Upperonly, "upperonly", false, "Only use uppercase text in -text mode.")

}

//Version : Displays version of stuntman
func Version() {
	fmt.Print(`


stuntman
        -Version : ALPHA-DEV
        -Author : Solirs
        -Repo : https://github.com/Solirs/stuntman



        `, "\n")
	os.Exit(0)
}

func main() {

	initflags()
	flag.Parse()

	if inc.V {
		Version()

	}

	switch inc.Clr {

	case "red":
		inc.Colr = inc.Red

	case "green":
		inc.Colr = inc.Green

	case "blue":
		inc.Colr = inc.Blue

	case "purple":
		inc.Colr = inc.Purple

	case "yellow":
		inc.Colr = inc.Yellow

	case "cyan":
		inc.Colr = inc.Cyan

	case "white":
		inc.Colr = inc.White

	default:
		inc.Colr = inc.Reset
	}

	inc.Arrl = make([]int, inc.Cnt) //Make array of non constant length

	fmt.Println(inc.Colr)

	if inc.Binar {
		eff.Binary()
	} else if inc.Text {
		eff.Text()
	} else if inc.Custom_string != "nil" {
		eff.Custom()

	} else {
		fmt.Println("No mode selected, QUITTING!!")
		os.Exit(1)

	}

	//This is only done once so spamming if else statements is fine i guess

}
