package main

import (
	"app/config"
	"app/database"
	"app/internal"
	"app/internal/logger"
	"app/internal/person"
	"app/internal/user"
	"app/model"
	"fmt"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error while loading configs: %s\n", err.Error())
	}

	db, err := database.Connect(*cfg)
	if err != nil {
		fmt.Printf("Error while connecting to database: %s\n", err.Error())
	}

	auditRepository := logger.NewAuditRepository(db)
	logger := logger.NewLogger(*auditRepository)

	personRepository := internal.NewRepository[model.Person](model.Person{}.TableName(), db, logger)
	personService := person.NewPersonService(personRepository, logger)
	personHandler := person.NewPersonHandler(personService)

	userRepository := internal.NewRepository[model.SysUser](model.SysUser{}.TableName(), db, logger)
	userService := user.NewSysUserService(userRepository, logger)
	userHandler := user.NewSysUserHandler(userService)

	handlers := []internal.Handler{
		personHandler,
		userHandler,
	}

	menu(handlers)
}

func menu(handlers []internal.Handler) {
	var choice int

	for true {
		fmt.Print("\033[2J\033[H")
		fmt.Print("МЕНЮ\n\n")

		for n, handler := range handlers {
			fmt.Printf("%d. %s\n", n+1, handler.Name())
		}

		fmt.Printf("%d. Выход\n", len(handlers)+1)
		fmt.Print(">>> ")
		fmt.Scan(&choice)

		if choice == len(handlers)+1 {
			break
		}

		handlers[choice-1].Menu()
	}
}
