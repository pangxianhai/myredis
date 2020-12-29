package cmd

import (
	"fmt"
	"io"
	"redis-common/proto/response"
	"redis-common/proto/str"
)

func init() {
	get := new(Get)
	Register(get)
}

type Get struct {
}

func (get *Get) HandleInput(args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, ArgsNumErr(get.Name())
	}
	req := new(str.GetReq)
	req.Key = args[0]
	return str.GetReqToByte(req), nil
}

func (get *Get) HandleResult(res *response.Response, writer io.Writer) {
	if res == nil {
		return
	}
	getRes := str.GetResFromByte(res.Data)
	if len(getRes.Value) > 0 {
		_, _ = fmt.Fprintln(writer, getRes.Value)
	} else {
		_, _ = fmt.Fprintln(writer, "nil")
	}
}

func (get *Get) Name() string {
	return "get"
}
