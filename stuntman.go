package main

import (
	"flag"
	"fmt"
	"os"
	eff "stuntman/effects"
	inc "stuntman/include")



func initflags(){
	flag.IntVar(&inc.Cnt, "width", 50, "The width of the cascade in rows")
	flag.IntVar(&inc.Vgap, "vertgap", 0, "Gap between lines")
	flag.IntVar(&inc.Speed, "speed", 0, "Time between lines in milliseconds")
	flag.StringVar(&inc.Clr, "color", "nil", "The color of the cascade")
	flag.BoolVar(&inc.Binar, "binar", false, "Display random Binary code")
	flag.BoolVar(&inc.Text, "text", false, "Display random text characters")
	flag.BoolVar(&inc.V, "v", false, "Display version and additonal info")
	


}

func Version(){
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

	if inc.V{
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

		default: inc.Colr = inc.Reset
	}

	inc.Arrl = make([]int, inc.Cnt) //Make array of non constant length

	fmt.Println(inc.Colr)

	if inc.Binar {
		eff.Binary()
	}else if inc.Text{
		eff.Texteff()
	} else {
		fmt.Println("No mode selected, QUITTING!!")
		os.Exit(1)

	}

}
