package server

import (
	"fmt"
	"godbms/src/command"
	"godbms/src/command/kvstorage"
	"strings"
)

func CommandProcessor(wrapper *WorkerWrapper, commandStr string) {
	var cmd command.Command

	switch {
	case strings.HasPrefix(strings.ToLower(commandStr), "kv"):
		cmd = kvstorage.NewCommand(wrapper.GetStorages(), commandStr)
	}

	fmt.Fprint(wrapper.GetConnection(), cmd.Process())
}
