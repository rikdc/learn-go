package main // All go applications start with a package.

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() { // keyboard shortcut = fm<tab>
	f, err := os.Open("myapp.log") // Saving added import automatically. Huzzah!
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
		if strings.Contains(s, "ERROR") {
			fmt.Println(s)
		}
	}
}
