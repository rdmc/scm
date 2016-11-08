package main

import (
	"fmt"

	//"net"
	"strings"
	//"sync"
	"io"

	"golang.org/x/crypto/ssh"
)

type CMTS struct {
	name    string
	addr    string
	prompt  string
	logging string
	conn    *ssh.Client
	// open time
	// last time
}

type Session struct {
	//mu sync.mutex
	in   io.WriteCloser
	out  io.Reader
	cmts *CMTS
	*ssh.Session
}

func (c *CMTS) Connect() error {
	var err error
	sshConfig := &ssh.ClientConfig{
		User: config.cmtsUsername,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.cmtsPassword),
		},
	}
	sshConfig.Ciphers = append(sshConfig.Ciphers, "aes128-cbc") // bad ciscio, bad !!!!
	c.conn, err = ssh.Dial("tcp", c.addr+":22", sshConfig)
	if err != nil {
		//		log.Println("error dialing", err)
	}
	return err
}

func (c *CMTS) Close() error {
	return c.conn.Close()
}

func (c *CMTS) CreateSession() (*Session, error) {
	var err error
	var s Session

	s.Session, err = c.conn.NewSession()
	if err != nil {
		//		log.Println("error new session", err)
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.OCRNL:         0,
		ssh.TTY_OP_OSPEED: 9600,
		ssh.TTY_OP_ISPEED: 9600,
	}

	s.in, err = s.Session.StdinPipe()
	s.out, err = s.Session.StdoutPipe()

	err = s.Session.RequestPty("vt100", 0, 2000, modes)
	if err != nil {
		//		log.Println("error session request tty", err)
		return nil, err
	}

	err = s.Session.Shell()
	if err != nil {
		//		log.Println("error session shell", err)
		return nil, err
	}

	s.cmts = c
	s.readnl() // discar banner

	return &s, nil
}

func (s *Session) Close() error {
	return s.Session.Close()
}

func (s *Session) readnl() (string, error) {
	// TODO need a time out ....
	buf := make([]byte, 16*1024)
	res := ""
	for {
		n, err := s.out.Read(buf)
		if err != nil {
			return "", err
		}
		res += string(buf[:n])
		if strings.Contains(res, s.cmts.prompt) {
			break
		}
	}
	res = strings.Replace(res, "\r", "", -1)
	return res, nil
}

func (s *Session) Command(cmd string) ([]string, error) {
	// TODO need a timeout/cancel
	_, err := s.in.Write([]byte(cmd + "\n"))

	if err != nil {
		//log.Println("error command write", err)
		return nil, err
	}

	res, err := s.readnl()
	if err != nil {
		//log.Println("error command read", err)
		return nil, err
	}

	out := strings.Split(res, "\n")
	// remove 1st line, the command itself, and last line, the prompt
	if len(out) < 2 {
		return nil, fmt.Errorf("CMTS response too short")
	}
	out = out[1 : len(out)-1]

	return out, nil
}

/*
func scm(t *telnet.Conn, mac string) (string, string, error) {

	if len(res) != 5 {
		return "", "", fmt.Errorf("show cable modem: WTF???")
	}

	// replace "!" in response exp: "...online(pt)    4172!-4.50  4422"
	f := strings.Fields(strings.Replace(res[3], "!", " ", -1))

	if len(f) != 9 {
		for k, v := range f {
			fmt.Println(k, v)
		}

		return "", "", fmt.Errorf("show cable modem FIELDS: WTF???")
	}

	return f[3], res[3], nil
	//
        //
	//                                                                                    D
	//   MAC Address    IP Address     I/F           MAC           Prim RxPwr  Timing Num I
	//                                               State         Sid  (dBmv) Offset CPE P
	//   7085.c6dd.cd57 10.1.1


*/
