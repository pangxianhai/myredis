package net

import (
    "bufio"
    "io"
    "log"
    "net"
    "redis-common/bytes"
    "redis-common/message"
    "redis-common/result"
    "redis-server/biz/cmd"
    "redis-server/data/config"
    "sync"
)

type Client struct {
    //标记
    flag uint8
    //客户端ID
    id uint64
    //客户端地址
    address string
    // Tcp 连接通道
    reader *bufio.Reader
    writer *bufio.Writer
    conn   net.Conn
}

type Server struct {
    mutex   sync.Mutex
    seq     uint64
    clients map[uint64]*Client
}

var server *Server

func Start() error {
    //监听 tcp 端口
    netListen, err := net.Listen("tcp", ":"+config.ServerPort())
    //如果监听 tcp 失败 则启动失败 退出
    if err != nil {
        return err
    }
    server = new(Server)
    server.clients = make(map[uint64]*Client, 0)
    log.Println("服务启动完成")
    //循环接受每个客户端的连接
    for {
        conn, err := netListen.Accept()
        if err != nil {
            continue
        }
        //组装客户端信息
        client := Client{
            flag:    0,
            id:      server.seq + 1,
            address: conn.RemoteAddr().String(),
            reader:  bufio.NewReader(conn),
            writer:  bufio.NewWriter(conn),
            conn:    conn,
        }

        server.seq++
        server.clients[client.id] = &client
        //启一个协程 服务每个客户端
        go server.accept(&client)
    }
}

func (server *Server) accept(client *Client) {
    log.Println(client.address, "连接服务器")
    for {
        msg, err := server.read(client)

        if err == io.EOF {
            log.Println(client.address, "与服务器断开连接")
            server.close(client)
            break
        } else if err != nil {
            log.Println(client.address, "读取客户端数据错误", err)
            continue
        }

        res := cmd.Call(msg.Content)

        msg.Content, err = result.ToJson(res)
        if err != nil {
            log.Println("命令执行结果转成 json 字符串失败", err)
            msg.Content = ""
        }
        err = server.write(client, msg)
        if err == io.EOF {
            log.Println(client.address, "与服务器断开连接")
            server.close(client)
            break
        } else if err != nil {
            log.Println(client.address, "向客户端写数据错误", err)
        }
    }
}

func (server *Server) read(client *Client) (*message.Message, error) {
    nBuf := make([]byte, 8)
    _, err := io.ReadFull(client.reader, nBuf)
    if err != nil {
        return nil, err
    }
    n := bytes.ToUint64(nBuf)
    buf := make([]byte, n)
    _, err = io.ReadFull(client.reader, buf)
    if err != nil {
        return nil, err
    }
    msg, err := message.FromByte(buf)
    if err != nil {
        return nil, err
    }
    return msg, nil
}

func (server *Server) write(client *Client, msg *message.Message) error {
    tBuf := message.ToPacket(msg)
    _, err := client.writer.Write(tBuf)
    if err != nil {
        return err
    }
    err = client.writer.Flush()
    if err != nil {
        return err
    }
    return nil
}

func (server *Server) close(client *Client) {
    _ = client.conn.Close()
    server.mutex.Lock()
    delete(server.clients, client.id)
    server.mutex.Unlock()
}
