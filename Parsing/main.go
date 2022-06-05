package main

import (
	"fmt"
	"net"
)

func main() {
	// urlstring := "https://ether-rbbc-1:1999"
	urlstring := "34.151.69.144:1888"
	// u, err := url.Parse(urlstring)
	// if err != nil {
	// 	fmt.Println("can't parse")
	// 	return
	// }
	// fmt.Println(u)
	h, p, err := net.SplitHostPort(urlstring)
	if err != nil {
		fmt.Println("Can't split")
		fmt.Println(err)
		return
	}
	fmt.Println(h)
	fmt.Println(p)
}
