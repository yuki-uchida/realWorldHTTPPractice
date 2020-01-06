package main


import "io/ioutil"
import "log"
import "net/http"
import "fmt"

func main() {
	fmt.Printf("Hello World !\n")
	resp, _ := http.Get("http://www.oreilly.co.jp/catalog/soon.xml")
	body, _ := ioutil.ReadAll(resp.Body)
	log.Print(string(body))
	resp.Body.Close()
}


