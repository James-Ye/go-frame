package pipe

import (
	"fmt"
	"net"
	"time"

	"github.com/James-Ye/go-frame/logger"
	gw "github.com/Microsoft/go-winio"
)

type CloseWriter interface {
	CloseWrite() error
}

type callbackFunc func(uintptr, string, interface{})

func Receive(pipeFile string, control chan bool, listenerDone chan bool, pfun callbackFunc) {
	c := gw.PipeConfig{
		MessageMode:      true,  // Use message mode so that CloseWrite() is supported
		InputBufferSize:  65536, // Use 64KB buffers to improve performance
		OutputBufferSize: 65536,
	}
	l, err := gw.ListenPipe(pipeFile, &c)
	if err != nil {
		logger.Trace("Listen fail")
	}
	defer l.Close()

	go func() {
		for {
			// server echo
			conn, e := l.Accept()
			logger.Trace("Accept recieve information")
			if e != nil {
				logger.Trace("Accept recieve information fail")
				break
			}
			defer conn.Close()

			bytes := make([]byte, 2000)
			_, e = conn.Read(bytes)
			if e != nil {
				fmt.Println("Read error")
				fmt.Println(e)
				break
			}

			command := string(bytes[:])
			pfun(0, command, conn)
		}

		close(listenerDone)
	}()

	<-listenerDone
	close(control)
}

func WriteBack(conn interface{}, message string) bool {
	if conn == nil {
		return false
	}

	client := conn.(net.Conn)
	if _, err := client.Write(([]byte)(message)); err == nil {
		return true
	}
	return false
}

func Send(pipeFile string, message []byte) {
	timeout := 1 * time.Second
	client, err := gw.DialPipe(pipeFile, &timeout)
	if err != nil {
		fmt.Println("DialPipe error")
		fmt.Println(err)
	}
	defer client.Close()

	n, err := client.Write(message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Writed")

	message_len := len(message)
	if n != message_len {
		fmt.Printf("expected %d bytes, send %v\n", message_len, n)
	}
	client.(CloseWriter).CloseWrite()
}
