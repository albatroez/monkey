package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(MONKEY_FACE)
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", usr.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
