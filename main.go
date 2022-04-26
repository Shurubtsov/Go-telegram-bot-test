package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.vk.com/method/wall.get?access_token=a7975a19a7975a19a7975a19a3a7eb5b2baa797a7975a19c5dc43ae78120dc1c0301231&v=5.131&domain=picturesforlovers&count=5&filter=owner")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	//log.Printf(sb)
	log.Printf("%T", sb)
}
