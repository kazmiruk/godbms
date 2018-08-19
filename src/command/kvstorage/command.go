package kvstorage

import (
	"godbms/src/storages"
	"godbms/src/structures/kvstorage"
	"strconv"
	"strings"
)

type KeyValueCommand struct {
	storages   *storages.Storages
	commandStr string
}

func NewCommand(storages *storages.Storages, commandStr string) (storage *KeyValueCommand) {
	return &KeyValueCommand{commandStr: commandStr, storages: storages}
}

func (command *KeyValueCommand) Process() (response string) {
	commandParts := strings.Split(command.commandStr, " ")
	response = ""
	storage := command.storages.GetStorage(commandParts[1])

	switch strings.ToLower(commandParts[2]) {
	case "create":
		command.storages.CreateStorage("kv", commandParts[1])
		return "KV storage " + commandParts[1] + " created"
	case "get":
		kvObject := storage.Get(commandParts[3]).(kvstorage.KeyValueObject)
		return kvObject.Value
	case "set":
		exp, _ := strconv.Atoi(commandParts[4])

		storage.Set(commandParts[3], strings.Join(commandParts[5:], " "), int64(exp))
	case "del":
		storage.Delete(commandParts[3])
	}

	return response
}
