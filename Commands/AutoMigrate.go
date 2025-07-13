package Commands

import (
	"fmt"
	"os"

	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
)

type NameModelPair struct {
	Name  string
	Model any
}

func AutoMigrate() {
	models := []NameModelPair{
		{"users", Models.User{}},
	}
	tx := Config.DB

	for _, pair := range models {
		fmt.Printf("Migrating table %s ...\n", pair.Name)
		if err := tx.AutoMigrate(&pair.Model); err != nil {
			fmt.Printf("Error while auto migrating %s, see why:\n", pair.Name)
			fmt.Printf("\t%s\n", err.Error())
			os.Exit(2)
		}
	}
}
