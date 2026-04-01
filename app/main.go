package main

import (
	"bufio"
	"flag"
	"fmt"
	"jsobfuscator"
	"os"
	"strings"
)

func main() {

	level := flag.String("level", "", "Level of obfuscation: default, low, medium, high")
	flag.Parse()

	var oLevel jsobfuscator.ObfuscationLevel

	if *level != "" {
		switch *level {
		case "low":
			oLevel = jsobfuscator.LowLevel
		case "medium":
			oLevel = jsobfuscator.MediumLevel
		case "high":
			oLevel = jsobfuscator.HighLevel
		default:
			oLevel = jsobfuscator.DefaultLevel
		}
	} else {
		flag.Usage()
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
	obf, err := jsobfuscator.NewObfuscator()
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
