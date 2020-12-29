package net

import (
	"bufio"
	"io"
	"log"
	"net"
	"redis-cli/biz/config"
	"redis-common/bytes"
	"redis-common/proto/message"
	"sync"
)

type Call struct {
	Args  *message.Message
	Reply *message.Message
	Error error
	Done  chan *Call
}

type Client struct {
	reader        *bufio.Reader
	writer        *bufio.Writer
	mutex         sync.Mutex
	seq           uint64
	pending       map[uint64]*Call
	closing       bool
	ServerAddress string
}

func New() (*Client, error) {
	var err error
	conn, err := net.Dial("tcp", config.ServerAddress())
	if err != nil {
		return nil, err
	}
	client := new(Client)
	client.reader = bufio.NewReader(conn)
	client.writer = bufio.NewWriter(conn)
	client.pending = make(map[uint64]*Call, 0)
	client.ServerAddress = config.ServerAddress()
	go client.input()
	return client, nil
}

func (client *Client) Send(data []byte) (response []byte, err error) {
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
	response = c.Reply.Content
	return
	//var client *rpc.Client
	//client.Call("UserServiceImpl", nil, nil)
}

func (client *Client) write(msg *message.Message) (err error) {
	buf := message.ToPacket(msg)

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
			if err == io.EOF {
				log.Fatalln("断开服务器连接")
			}
			log.Println("读取服务器数据失败", err)
			continue
		}
		msgLen := bytes.ToUint64(lenBuf)
		msgBuf := make([]byte, msgLen)
		_, err = io.ReadFull(client.reader, msgBuf)
		if err != nil {
			if err == io.EOF {
				log.Fatalln("断开服务器连接")
			}
			log.Println("读取服务器数据失败", err)
			continue
		}
		msg := message.FromByte(msgBuf)

		client.mutex.Lock()
		call := client.pending[msg.Id]
		delete(client.pending, msg.Id)
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
