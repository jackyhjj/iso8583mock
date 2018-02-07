package action

import (

	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
	"github.com/nkristianto/tcp_server/services"
)

func RunTcpClient() {
	// graceful shutdown consumer
	var (
		signalShutdown        = make(chan os.Signal)
	)
	signal.Notify(signalShutdown,syscall.SIGTERM)
	signal.Notify(signalShutdown,syscall.SIGINT)

	go func(){
		sig := <- signalShutdown
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for 3 second to finish processing")
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}()

	client := new(services.TcpClientService)
	client.StartClient()
}
