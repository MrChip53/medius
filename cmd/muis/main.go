package main

import (
	"fmt"
	"log/slog"
	"medius-server/pkg/medius"
	"medius-server/pkg/tcp"
	"net"
)

// TODO packets for muis server
//RT_MSG_SERVER_HELLO
//RT_MSG_SERVER_CONNECT_ACCEPT_TCP
//RT_MSG_SERVER_CONNECT_COMPLETE
//RT_MSG_SERVER_ECHO
//RT_MSG_CLIENT_DISCONNECT_WITH_REASON
//MediusUniverseVariableInformationResponse

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {
		var rtMessages []*medius.RTMessage

		buf := make([]byte, 8192)
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		for i := 0; i < n; i++ {
			rtMessage := medius.ReadNextRTMessage(buf[i:n])
			rtMessages = append(rtMessages, rtMessage)
		}

		for _, rtMessage := range rtMessages {
			b, err := medius.ProcessRTMessage(rtMessage)
			if err != nil {
				slog.Error("Error processing RT message: %v", err)
				continue
			}
			_, err = conn.Write(b)
		}
	}
}

func main() {
	fmt.Println("Starting Medius Universe Information Server...")
	tcpServer := tcp.NewTCPServer(handleClient)
	err := tcpServer.ListenAndServe(":10071")
	if err != nil {
		panic(err)
	}
}
