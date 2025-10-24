package Commands

import (
	"fmt"
	"os"
	"vahid/Config"
	"vahid/Models"
)

type NameModelPair struct {
	Name  string
	Model any
}

func AutoMigrate() {
	models := []NameModelPair{
		{Name: "users", Model: Models.Users{}},
	}

	tx := Config.DB

	for _, pair := range models {
		fmt.Printf("Migrating table %s ...\n", pair.Name)
		if err := tx.AutoMigrate(pair.Model); err != nil {
			fmt.Printf("Error while auto migrating %s:\n\t%s\n", pair.Name, err.Error())
			os.Exit(2)
		}
		fmt.Printf("Table %s migrated successfully!\n", pair.Name)
	}
	fmt.Println("All migrations completed successfully!")
}
