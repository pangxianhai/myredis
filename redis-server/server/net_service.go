package server

import "net"

func StartTpcServer() {
    netListen, err := net.Listen("tcp", ":9988")
    for {
        conn, err := netListen.Accept()
        if err != nil {
            continue
        }
        conn.RemoteAddr().Network()
    }

}
