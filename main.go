package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	//"net"
	//"sync"

	"github.com/rdmc/mac"
)

func main() {
	fmt.Println("Show Cable Modem test")
	ticker := time.Tick(200 * time.Millisecond)
	whell := []byte("\\|/-")
	whellCnt := 0
	whellPtr := flag.Bool("whell", false, "rotating whell")

	flag.Parse()
	m, err := mac.ParseMAC(flag.Arg(0))
	if err != nil {
		log.Fatal("You must provide a valid mac addr")
	}

	fmt.Println("looking for", m)

	c := &CMTS{
		//name:   "pdl1cmts002",
		//addr:   "10.212.128.1",
		name:   config.cmtsAddr,
		addr:   config.cmtsAddr,
		prompt: "#",
	}

	if err := c.Connect(); err != nil {
		log.Fatal("UNEBLE TO CONNECT:", err)
	}
	defer c.Close()
	s, err := c.CreateSession()
	if err != nil {
		log.Fatal("UNABLE TO CREATE SESSION:", err)
	}
	defer s.Close()

	startTime := time.Now()
	last_state := ""
main_loop:
	for {
		select {
		case <-ticker:
			// pass
		}

		res, _ := s.Command("show cable modem " + m.CiscoString())

		state, line, err := parseSCM(res)

		if err != nil {
			log.Println("Error:", err)
			break main_loop
		}
		//fmt.Printf("CMTS:%s\nModem: %s\nState: %s\nline: %q\n", c.name, m, state, line)
		if state != last_state {
			last_state = state
			elapsedTime := time.Since(startTime)
			desc, _ := cm_status(state)
			fmt.Printf("\r                                                                                     ")
			fmt.Printf("\r@%s, Elapsed: %s, State: %s,\n\t%s\n", time.Now().Format("2006-01-02 15:04:05.000"), elapsedTime, state, desc)
		}
		if *whellPtr {
			whellCnt++
			whellCnt %= 4
			fmt.Printf("\r%s %c", line, whell[whellCnt])
		}
	}
	fmt.Println("That's All Folks!")
}

func loop() {
	ticker := time.Tick(200 * time.Millisecond)
	for {
		select {
		case <-ticker:
			// pass
		}

	}
}
