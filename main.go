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
	ticker := time.Tick(100 * time.Millisecond)
	wheel := []byte("\\|/-")
	wheelCnt := 0
	wheelPtr := flag.Bool("wheel", false, "rotating wheel")

	flag.Parse()
	m, err := mac.ParseMAC(flag.Arg(0))
	if err != nil {
		log.Fatal("You must provide a valid mac addr")
	}

	c := &CMTS{
		//name:   "pdl1cmts002",
		//addr:   "10.212.128.1",
		name:   config.cmtsAddr,
		addr:   config.cmtsAddr,
		prompt: config.cmtsPrompt,
	}

	fmt.Println("looking for", m, "in", c.addr)

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
	lastState := ""
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
		if state != lastState {
			lastState = state
			elapsedTime := time.Since(startTime)
			desc, _ := cmStatus(state)
			fmt.Printf("\r                                                                                     ")
			fmt.Printf("\r@%s, Elapsed: %s, State: %s,\n\t%s\n", time.Now().Format("2006-01-02 15:04:05.000"), elapsedTime, state, desc)
		}
		if *wheelPtr {
			wheelCnt++
			wheelCnt %= 4
			fmt.Printf("\r%s %c", line, wheel[wheelCnt])
		}
	}
	fmt.Println("That's All Folks!")
}
