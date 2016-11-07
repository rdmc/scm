package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	//"net"
	//"sync"

	"github.com/rdmc/mac"
)

func main() {
	fmt.Println("Show Cable Modem test")

	flag.Parse()
	m, err := mac.ParseMAC(flag.Arg(0))
	if err != nil {
		log.Fatal("You must provide a valid mac addr")
	}

	fmt.Println("looking for", m)

	c := &CMTS{
		name:   "cm02ac01",
		addr:   "10.212.128.1",
		prompt: "#",
	}

	if err := c.Connect(); err != nil {
		log.Fatal("UNEBLE TO CONNECT:", err)
	}
	fmt.Println("!#connect ok")
	s, err := c.CreateSession()
	if err != nil {
		log.Fatal("UNABLE TO CREATE SESSION:", err)
	}
	fmt.Println("!#create session ok")
	res, _ := s.Command("show cable modem " + m.CiscoString())
	fmt.Println(strings.Join(res, "\n"))
	fmt.Printf("len res = %d\n", len(res))

	fmt.Println("!#send commmand ok")

	s.Close()
	fmt.Println("!#Session close ok")
	c.Close()
	fmt.Println("!#Client close ok")

	fmt.Println("!#Bye, Bye")

	fmt.Println("That's All Folks!")
}
