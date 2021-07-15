package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Crpyto-ticker")
	reader := bufio.NewReader(os.Stdin)
	getData(reader)
}

func getData(reader *bufio.Reader) {
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("An error occurred, " + err.Error())
	}
	text = strings.Replace(text, "\n", "", -1)
	println(text)
}
