package main

import (
	"bufio"    //Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer)
	"fmt"
	"os" // reader in go
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)       // scanner
	fmt.Print("Enter Text:" )
	scanner.Scan()
	text := scanner.Text()
	
	//*******zeile 17 bis 25  sprachblasen******

	for k := 0; k < len(text); k++ { //for loop
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println(text)

	for k := 0; k < len(text); k++ {
		fmt.Printf("-")
	}

	printCow() // Aufrufen von Methode
}

func printCow() {

	cow := ` 

       \   ^__^
        \  (oo)\_______
           (__)\       )\/\
               ||----w |
               ||     ||  			           
			       
			   `

	fmt.Println(cow) // ausgabe in Terminal
}
