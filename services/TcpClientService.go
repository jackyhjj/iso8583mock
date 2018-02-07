package services

import (
	"os"
	"fmt"
	"io/ioutil"
	"net"
)

type TcpClientService struct {
}

func (s *TcpClientService) StartClient() {
	service := "localhost:3333"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	s.CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	s.CheckError(err)

	_, err = conn.Write([]byte("test message"))
	s.CheckError(err)

	//result, err := readFully(conn)
	result, err := ioutil.ReadAll(conn)
	s.CheckError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func (s *TcpClientService) CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
