package cmd

import (
	"fmt"
	"io"
	"redis-common/proto/response"
	"redis-common/proto/str"
	"strconv"
	"strings"
)

func init() {
	set := new(Set)
	Register(set)
	setex := new(SetEx)
	Register(setex)
}

type Set struct {
}

//set 命令 set key value [EX|PX KEEPTTL] [NX|XX]
func (set *Set) HandleInput(args []string) ([]byte, error) {
	if len(args) < 2 {
		return nil, ArgsNumErr(set.Name())
	}
	req := new(str.SetReq)
	req.Key = args[0]
	req.Value = args[1]
	for i := 3; i < len(args); i++ {
		if len(args[i]) == 0 {
			continue
		}
		v := strings.ToLower(args[i])
		if v == "ex" || v == "px" {
			req.Expx = v
			if i == len(args)-1 {
				return nil, SyntaxErr
			}
			timeout, err := strconv.Atoi(args[i+1])
			if err != nil {
				return nil, SyntaxErr
			}
			req.Timeout = int32(timeout)
		} else if v == "nx" || v == "xx" {
			req.Nxxx = v
		}
	}
	return str.SetReqToByte(req), nil
}
func (set *Set) HandleResult(res *response.Response, writer io.Writer) {
	if res == nil {
		return
	}
	if res.Code == response.Success {
		_, _ = fmt.Fprintln(writer, "OK")
	} else {
		_, _ = fmt.Fprintln(writer, res.Msg)
	}
}

func (set *Set) Name() string {
	return "set"
}

type SetEx struct {
}

func (setEx *SetEx) HandleInput(args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, ArgsNumErr(setEx.Name())
	}
	timeout, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, SyntaxErr
	}
	req := new(str.SetReq)
	req.Key = args[0]
	req.Timeout = int32(timeout)
	req.Value = args[2]
	return str.SetReqToByte(req), nil
}

func (setEx *SetEx) HandleResult(res *response.Response, writer io.Writer) {
	if res == nil {
		return
	}
	if res.Code == response.Success {
		_, _ = fmt.Fprintln(writer, "OK")
	} else {
		_, _ = fmt.Fprintln(writer, res.Msg)
	}
}

func (setEx *SetEx) Name() string {
	return "setex"
}
