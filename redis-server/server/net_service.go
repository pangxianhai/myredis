package server

import (
    "log"
    "net"
    "reids-server/data/config"
    "strconv"
)

type ClientInfo struct {
    //标记
    flag uint8
    //客户端ID
    id uint32
    //客户端地址
    address string
    // tcp 连接通道
    conn net.Conn
}

type NetService struct {
    config     config.Config
    clientList []ClientInfo
}

func (netService *NetService) StartTcpServer() {
    netListen, err := net.Listen("tcp", ":"+strconv.Itoa(netService.config.GetServerPort()))
    if err != nil {
        log.Fatalln("server start failed", err)
    }
    netService.clientList = make([]ClientInfo, 0)
    for {
        conn, err := netListen.Accept()
        if err != nil {
            continue
        }
        clientInfo := ClientInfo{
            flag:    0,
            id:      0,
            address: conn.RemoteAddr().String(),
            conn:    conn,
        }
        netService.clientList = append(netService.clientList, clientInfo)
        go netService.acceptTcpMessage(clientInfo)
    }
}

func (netService *NetService) acceptTcpMessage(clientInfo ClientInfo) {

}
