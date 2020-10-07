package main

import "fmt"

func main() {
	uaString := "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.7 (KHTML, like Gecko) RockMelt/0.8.36.78 Chrome/7.0.517.44 Safari/534.7"

	ua := NewUAParser(uaString)
	ua.Parse()

	fmt.Printf("%+v\n", ua.OS)
	fmt.Printf("%+v\n", ua.Browser)
	fmt.Printf("%+v\n", ua.Device)
	fmt.Printf("%+v\n", ua.clients)
}
