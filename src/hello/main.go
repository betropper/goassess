package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	"strings"
	"time"
)

func PrintPlain(message string) {
	fmt.Println(c.Clear + "What's up, " + message + "?")
}

func PrintMulti(message string) {
	for {
		fmt.Print(c.Clear)
		fmt.Println(c.Multi("ayy lmao, " + message + "."))
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	message := "world"
	option := ""
	if len(os.Args) == 2 {
		if strings.HasPrefix(os.Args[1], "-") {
			option = os.Args[1]
		} else {
			message = os.Args[1]
		}

	} else if len(os.Args) > 2 {
		option = os.Args[1]
		message = os.Args[2]
	}

	if option == "-m" {
		PrintMulti(message)
	} else {
		PrintPlain(message)
	}

}
