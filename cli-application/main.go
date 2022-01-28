package main // All go applications start with a package.

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() { // keyboard shortcut = fm<tab>
	path := flag.String("path", "myapp.log", "The path to the log file.")
	level := flag.String("level", "ERROR", "The error level to search for.")

	flag.Parse()

	// vv Notice the asterix!
	f, err := os.Open(*path) // Saving added import automatically. Huzzah!
	if err != nil {
		log.Fatal(err)
	}

	// defer, after the main function exists will close the file automatically
	defer f.Close()

	// reading contents of the file.
	r := bufio.NewReader(f)

	for { // infinite for loop
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
