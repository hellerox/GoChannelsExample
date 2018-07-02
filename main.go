package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://particl.io",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//2 - Aquí la recibes
	//Esto también es un Blocking call
	for l := range c { //nunca terminará, si colocamos el sleep aqui sería como atrasar a todos las demás rutinas que lleguen
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // Esto es una lambda function / function literal que creamos para no agregar directamente time.Sleep en la función checkLink
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "migh be down!")
		c <- link //1 - Aquí le envias la info al channel
		return
	}
	fmt.Println(link, "is OK!")
	c <- link
}
