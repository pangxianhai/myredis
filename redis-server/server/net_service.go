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
    id  uint32
    //客户端地址
    address string
    // tcp 连接通道
    conn net.Conn
}

type NetService struct {
    config     config.Config
    clientList []ClientInfo
}

// StartTCPServer 初始化 tcp 连接
func (netService *NetService) StartTCPServer() {
    //监听 tcp 端口
    netListen, err := net.Listen("tcp", ":"+strconv.Itoa(netService.config.GetServerPort()))
    //如果监听 tcp 失败 则启动失败 退出
    if err != nil {
        log.Fatalln("server start failed", err)
    }
    netService.clientList = make([]ClientInfo, 0)
    //循环接受每个客户端的连接
    for {
        conn, err := netListen.Accept()
        if err != nil {
            continue
        }
        //组装客户端信息
        clientInfo := ClientInfo{
            flag:    0,
            id:      0,
            address: conn.RemoteAddr().String(),
            conn:    conn,
        }
        netService.clientList = append(netService.clientList, clientInfo)
        //启一个协程 服务每个客户端
        go netService.acceptTCPMessagego(clientInfo)
    }
}

func (netService *NetService) acceptTCPMessagego(clientInfo ClientInfo) {
    buf := make([]byte, 1024)
    clientInfo.conn.Read(buf)

}
