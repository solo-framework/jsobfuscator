package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"gitlab.com/naicoi92/obfuscator"
)

func main() {

	level := flag.String("conlevelfig", "", "Level of obfuscation: low (default), medium, high")
	flag.Parse()

	var oLevel obfuscator.ObfuscationLevel

	if *level != "" {
		switch *level {
		case "low":
			oLevel = obfuscator.LowLevel
		case "medium":
			oLevel = obfuscator.MediumLevel
		case "high":
			oLevel = obfuscator.HighLevel
		default:
			oLevel = obfuscator.LowLevel
		}
	}

	// read from stdin

	scan := bufio.NewScanner(os.Stdin)
	sb := strings.Builder{}

	for scan.Scan() {
		if err := scan.Err(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		sb.Write(scan.Bytes())
	}

	jsCode := sb.String()
	if "" == jsCode {

		fmt.Println("no data")
		os.Exit(1)
	}

	// Initialize the obfuscator
	obf, err := obfuscator.NewObfuscator()
	if err != nil {
		panic(err)
	}

	obf.Level = oLevel //obfuscator.MediumLevel

	// Perform obfuscation
	obfuscatedCode, err := obf.Obfuscate(jsCode)
	if err != nil {
		panic(err)
	}

	// Print the obfuscated code
	fmt.Println(obfuscatedCode)
}
