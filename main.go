package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	fetchData(text)
}

func fetchData(resource string) {
	r, err := http.Get("https://jsonplaceholder.typicode.com/" + resource + "/1")
	if err != nil {
		fmt.Println("An error occurred" + err.Error())
	}
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(data))

}
