package inc

//Cnt : the width of the displayed effects
var Cnt int

//Vgap : The gap between lines
var Vgap int

//Speed : the delay between lines in milliseconds
var Speed int

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

//Custom_set : A custom set of characters to be used
var Custom_set = []string{}

//Charlist : The buffer slice that will be used to generate each line, it's purpose is to be overwritten by the slices below 
var Charlist = []string{}

//ASCII : List of characters to use when for -text. 
var ASCII = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "{", "}", "[", "]", ";", "§"}

//ASCIILower : ASCII but lowercase only
var ASCIILower = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "{", "}", "[", "]", ";", "§"}

//ASCIIUpper : ASCII but uppercase only
var ASCIIUpper = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "{", "}", "[", "]", ";", "§"}

//Binary_list an array of strings only containing 1 and 0 that can be modified according to other flags (i.e add spaces according to -spaces flag)
var Binary_list = []string{"1", "0"}

//Colr : Unix color code of the chosen color
var Colr string

//Clr : the color
var Clr string

//Custom_string : the custom character set before its split
var Custom_string string

//V : Option to display version
var V bool
