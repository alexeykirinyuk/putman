package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/alexeykirinyuk/putman/application"
	"github.com/alexeykirinyuk/putman/cmd/view"
	"github.com/alexeykirinyuk/putman/persistence"
	"github.com/google/uuid"
)

func main() {
	storage := persistence.NewJsonStorage("temp.json")
	col := application.NewCollectionService(storage)
	req := application.NewRequestService(storage)

	views := view.CreateViews(col)

	actions := map[string]func(args []string) error{
		"execute": func(args []string) error {
			if argsLen := len(args); argsLen != 1 {
				return fmt.Errorf("Not valid args count. Expected - 1, Found - %d", argsLen)
			}

			id, err := uuid.Parse(args[0])
			if err != nil {
				return errors.New("can't parse id")
			}

			resp, err := req.Execute(id)
			if err != nil {
				return err
			}

			fmt.Println("response:", resp)
			return nil
		},
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error when read string from console:" + err.Error())
		}

		splited := strings.Split(strings.TrimRight(str, "\n"), " ")
		command := splited[0]
		executed := false
		if command == "exit" {
			fmt.Println("Buy")
			executed = true
			break
		} else if command == "help" {
			i := 1
			for _, view := range views {
				fmt.Println(i, "-", view.GetName())
				i++
			}
			executed = true
		} else {
			for _, view := range views {
				if view.GetName() == command {
					err := view.Handle(splited[1:])
					if err != nil {
						fmt.Println(err)
					}
					executed = true
					break
				}
			}
		}

		if !executed {
			fmt.Println("Comamnd not found")
		}
	}
}
