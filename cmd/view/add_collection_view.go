package view

import (
	"fmt"

	"github.com/alexeykirinyuk/putman/application"
)

type AddCollectionView struct {
	collectionService *application.CollectionService
}

func (c AddCollectionView) GetName() string {
	return "add-collection"
}

func (c AddCollectionView) GetHelp() string {
	return "HELP: add-collection <collection-name>\nEXAMPLE: add-collection google-apis"
}

func (c AddCollectionView) Handle(args []string) error {
	if argsLen := len(args); argsLen != 1 {
		return fmt.Errorf("Not valid args count. Expected - 1, Found - %d", argsLen)
	}

	id, err := c.collectionService.Create(args[0])
	if err != nil {
		return err
	}

	fmt.Println(id)
	return nil
}
