package view

import (
	"fmt"

	"github.com/alexeykirinyuk/putman/application"
)

type TreeView struct {
	collectionService *application.CollectionService
}

func (c TreeView) GetName() string {
	return "tree"
}

func (c TreeView) GetHelp() string {
	return "HELP: tree\nEXAMPLE: tree"
}

func (c TreeView) Handle(_ []string) error {
	tree, err := c.collectionService.GetTree()
	if err != nil {
		return err
	}

	for _, col := range tree {
		printCollection(col)
	}

	return nil
}

func printCollection(col application.CollectionDto) {
	fmt.Printf(" - col [%s] %s", col.ID, col.Name)
	fmt.Println()

	for _, req := range col.Requests {
		printRequest(2, req)
	}

	for _, fol := range col.Folders {
		printFolder(2, fol)
	}
}

func printSpaces(count int) {
	for i := 0; i < count; i++ {
		fmt.Print(" ")
	}
}

func printFolder(spaces int, fol application.FolderDto) {
	printSpaces(spaces)
	fmt.Printf(" - fol %s", fol.Name)
	fmt.Println()

	for _, req := range fol.Requests {
		printRequest(spaces+2, req)
	}
	for _, childFol := range fol.Folders {
		printFolder(spaces+2, childFol)
	}
}

func printRequest(spaces int, req application.RequestDto) {
	printSpaces(spaces)
	fmt.Printf(" - req [%s] %s", req.ID, req.Name)
	fmt.Println()
}
