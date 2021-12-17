package inc

//Cnt : the width of the displayed effects
var Cnt int

//Vgap : The gap between lines
var Vgap int

//Speed : the delay between lines in milliseconds
var Speed int

//Clr : the color
var Clr string

//Binar : display binary effect
var Binar bool

//Arrl : temporary array used for effects
var Arrl []int

//Text : Displays text effct
var Text bool

//ASCII : List of characters to use in the next effect
var ASCII = []rune("abcdefg;()[{}]hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRS                                                       ")

//Colr : Unix color code of the chosen color
var Colr string

//V : Option to display version
var V bool
