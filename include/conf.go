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

//Spaces : An index to determine how many spaces will be generated when using the -text mode, the higher it is, the more spaces.

var Spaces int

//Arrl : temporary array used for effects
var Arrl []int

//Text : Displays text effct
var Text bool

//Loweronly : Only use lowercase text in -text mode.
var Loweronly bool

//Upperonly : Only use uppercase text in -text mode
var Upperonly bool

//ASCII : List of characters to use in the next effect
var ASCII = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "{", "}", "[", "]", ";", "§"}

//ASCII_Lower : ASCII but lowercase only
var ASCII_Lower = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "{", "}", "[", "]", ";", "§"}

//ASCII_Upper : ASCII but uppercase only
var ASCII_Upper = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "{", "}", "[", "]", ";", "§"}

//Colr : Unix color code of the chosen color
var Colr string

//V : Option to display version
var V bool
