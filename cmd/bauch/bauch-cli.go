package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Xjs/bauch"
)

func main() {
	var (
		help  bool = false
		smile bool = true
	)

	flag.BoolVar(&help, "help", help, "sHOw HeLP")
	flag.BoolVar(&smile, "smile", smile, "SmiLE aT tHe EnD oF the SenTENCe :o)")

	flag.Parse()

	if flag.NArg() < 1 || help {
		fmt.Println("uSAgE: bauch <Message in non-Bauch-format>")
		fmt.Println("oPTiOnS:")
		flag.PrintDefaults()
		os.Exit(2)
	}

	input := strings.Join(os.Args[1:], " ")

	output := bauch.Say(input)

	if smile {
		output += " " + bauch.Smile
	}

	fmt.Println(output)
}
