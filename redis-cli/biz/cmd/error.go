package cmd

import (
	"errors"
	"fmt"
)

var SyntaxErr = errors.New("ERR syntax error")

func ArgsNumErr(cmd string) error {
	return fmt.Errorf("ERR wrong number of arguments for '%s' command", cmd)
}

func UnknownCmd(cmd string) error {
	return fmt.Errorf("ERR unknown command '%s'", cmd)
}
