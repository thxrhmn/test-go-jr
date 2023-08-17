package database

import (
	"fmt"
	"todo-list-thxrhmn/models"
	"todo-list-thxrhmn/pkg/postgresql"
)

func RunMigration() {
	err := postgresql.DB.AutoMigrate(
		&models.Todo{},
		&models.SubTodo{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}

	fmt.Println("Migration success")
}
