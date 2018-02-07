package action

import (
	"github.com/nkristianto/tcp_server/services"
)

func RunTcpServer() {
	tcpServer := new(services.TcpServerService)
	tcpServer.RunServer()
}
