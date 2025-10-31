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
		// اگر مدل‌های دیگری داشته باشیممی توانیم اینجا اضافه کنید
		// {"products", Models.Product{}},
		// {"orders", Models.Order{}},
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
