package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ceosss/crypto-ticker/helpers"
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
	apiKey := helpers.GetApiKey()
	res, _ := http.Get("https://api.nomics.com/v1/currencies/ticker?key=" + apiKey + "&ids=BTC,ETH,XRP&interval=1d,30d&convert=EUR&per-page=100&page=1")
	d, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(d))

}
