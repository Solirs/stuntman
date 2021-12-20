[![Go Report Card](https://goreportcard.com/badge/github.com/Solirs/stuntman)](https://goreportcard.com/report/github.com/Solirs/stuntman)
# stuntman v1.0.0
a simple program written in go to make you look like a hackerman with cool terminal effects



## Demo

##### Obviously speed isn't visible because those are images and not videos

`stuntman -text -spaces 100 -loweronly -width 90 -speed 35 -color green`

![alt text](https://github.com/Solirs/stuntman/blob/main/ressources/Demo-txt.png?raw=true)


`stuntman -binar -color red -spaces 10 -speed 20 -width 90`

![alt text](https://github.com/Solirs/stuntman/blob/main/ressources/Demo-bin-2.png?raw=true)

`stuntman -custom 0,1,2,3,4,5,6,7,8,9 -color green -speed 55 -spaces 1 -width 100`

![alt text](https://github.com/Solirs/stuntman/blob/main/ressources/Demo-custom.png?raw=true)

## Installation

### Build from source

`git clone https://github.com/Solirs/stuntman`

`cd stuntman`

`sudo make install`

### Precompiled binary

Download the release tarball in the releases.

`tar -xvf stuntman_release_*.*.*.tar.gz`

`cd stuntman_release_*.*.*`

`sudo make install`

### Compatibility

Should work on any Unix-like operating system, especially if you build from source.

I personally tested stuntman on the following:

  - FreeBSD 13.0-Release (Working)
  - Arch Linux (Working)
  - Fedora 35 (Working)


### Uninstall

`cd STUNTMAN_DIRECTORY`

`sudo make uninstall`

## Use guide

#### -h

##### Displays the help menu

#### -color

##### The color of the text, can be cyan, blue, red, green, purple, or white.

#### -speed

##### The delay between every new lines in milliseconds

#### -linegap

##### Gap between every new line

#### -text

##### Generate ascii text

##### -loweronly

###### Only generate lowercase characters and special symbols

##### -upperonly

###### Same as loweronly but for upper

#### -binar

##### Generate ones and zeroe's

#### -custom foo,bar,eggs

##### Generate lines from a custom set of given strings separed by ",".

#### -width

##### The width of the effect in rows 
