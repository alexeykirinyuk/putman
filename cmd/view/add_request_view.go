package view

import (
	"fmt"

	"github.com/alexeykirinyuk/putman/application"
)

type AddRequestView struct {
	requestService *application.RequstService
}

func (c AddRequestView) GetName() string {
	return "add-request"
}

func (c AddRequestView) GetHelp() string {
	return `HELP: add-request <collection-name> <request-name> <method> <url>
	OUTPUT: request-id
	EXAMPLE: add-collection google-apis google-get get https://google.com
	EXAMPLE OUTPUT: b4ea857e-01bf-4b5e-93eb-2b2bf75e1cb3`
}

func (c AddRequestView) Handle(args []string) error {
	if argsLen := len(args); argsLen != 4 {
		return fmt.Errorf("Not valid args count. Expected - 4, Found - %d", argsLen)
	}

	id, err := c.requestService.Create(args[0], args[1], args[2], args[3])
	if err != nil {
		return err
	}

	fmt.Println(id)
	return nil
}
