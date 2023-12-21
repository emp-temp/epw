package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var helpMessage = `
Usage:
	epw [options] [length]

META OPTIONS
	-h show list of command-line options

GENERATION OPTIONS
	-a use ascii_letters (default)
	-d use digits
	-p use punctuation

This command can be compounded



EXAMPLE:
	ecm 12       -> LtwKoUHbqKJS 
	ecm 12       -> LtwKoUHbqKJS
	ecm -d 12    -> 151899596572 
	ecm -p 12    -> '{{+[]_"]-*% 
	ecm -ad 12   -> W04odpa4uynz
	ecm -dp 12   -> .\67@*@;/554
	ecm -ap 12   -> M]Izy@sp!SAi
	ecm -a -d 12 -> FAotPiLXawv3
`

var defaultPassChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	password := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	options := map[rune]bool{
		'a': false,
		'd': false,
		'p': false,
	}

	optionsString := map[rune]string{
		'a': "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		'd': "0123456789",
		'p': "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~",
	}

	if len(os.Args) == 1 {
		return
	}

	var positionArgs []string
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			for _, opt := range arg[1:] {
				if _, exist := options[opt]; exist {
					options[opt] = true
				} else if opt == 'h' {
					fmt.Print(helpMessage)
					return
				} else {
					fmt.Fprintf(os.Stderr, "Unknown option: %c\n", opt)
				}
			}
		} else {
			positionArgs = append(positionArgs, arg)
		}
	}

	if len(positionArgs) > 1 {
		fmt.Fprintf(os.Stderr, "Too many positons arguments: %d\n", len(positionArgs))
	}

	passwordLength, err := strconv.Atoi(positionArgs[0])
	if err != nil || passwordLength < 0 {
		fmt.Fprintf(os.Stderr, "Count argument can only be a positive number: %v, type: %T\n", passwordLength, passwordLength)
	}

	passChars := ""
	for k, isEnableOption := range options {
		if isEnableOption {
			passChars = passChars + optionsString[k]
		}
	}
	if len(passChars) == 0 {
		passChars = defaultPassChars
	}

	for i := 0; i < passwordLength; i++ {
		n := r.Intn(len(passChars))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failure to generate random numbers")
		}
		password = password + string(passChars[n])
	}

	fmt.Println(password)
}
