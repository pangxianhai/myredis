package net

import (
    "bufio"
    "io"
    "log"
    "net"
    "redis-cli/biz/config"
    "redis-common/bytes"
    "redis-common/message"
    "sync"
)

type Call struct {
    Args  *message.Message
    Reply *message.Message
    Error error
    Done  chan *Call
}

type Client struct {
    reader  *bufio.Reader
    writer  *bufio.Writer
    mutex   sync.Mutex
    seq     uint64
    pending map[uint64]*Call
    closing bool
}

var client *Client

func Init() error {
    var err error
    conn, err := net.Dial("tcp", config.ServerAddress())
    if err != nil {
        return err
    }
    client = new(Client)
    client.reader = bufio.NewReader(conn)
    client.writer = bufio.NewWriter(conn)
    go client.input()
    return nil
}

func Send(data string) (result string, err error) {
    call := new(Call)
    call.Done = make(chan *Call, 1)

    client.mutex.Lock()
    seq := client.seq
    client.seq++
    client.pending[seq] = call
    client.mutex.Unlock()

    call.Args = message.New(seq, data)
    err = client.write(call.Args)
    if err != nil {
        return
    }
    c := <-call.Done
    result = c.Reply.Content
    return
    //var client *rpc.Client
    //client.Call("UserServiceImpl", nil, nil)
}

func (client *Client) write(msg *message.Message) (err error) {
    msgBuf, err := message.ToByte(msg)
    if err != nil {
        return
    }
    msgLen := uint64(len(msgBuf))
    lenBuf, err := bytes.FromUint64(msgLen)
    if err != nil {
        return
    }

    buf := make([]byte, 0)
    buf = append(buf, lenBuf...)
    buf = append(buf, msgBuf...)

    client.mutex.Lock()
    defer client.mutex.Unlock()
    _, err = client.writer.Write(buf)
    if err != nil {
        return
    }
    err = client.writer.Flush()
    return
}

func (client *Client) input() {
    for {
        lenBuf := make([]byte, 8)
        _, err := io.ReadFull(client.reader, lenBuf)
        if err != nil {
            continue
        }
        msgLen, err := bytes.ToUint64(lenBuf)
        if err != nil {
            continue
        }
        msgBuf := make([]byte, msgLen)
        _, err = io.ReadFull(client.reader, msgBuf)
        msg, err := message.FromByte(msgBuf)

        client.mutex.Lock()
        call := client.pending[msg.ID]
        delete(client.pending, msg.ID)
        client.mutex.Unlock()
        if call != nil {
            call.Reply = msg
            call.done()
        }
    }
}

func (call *Call) done() {
    select {
    case call.Done <- call:
        //ok
    default:
        log.Print(" net server discarding Call reply due to insufficient Done chan capacity")
    }
}
