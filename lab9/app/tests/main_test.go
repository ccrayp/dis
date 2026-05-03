package tests

import (
	"app/config"
	"app/database"
	"app/internal"
	"app/internal/logger"
	"app/internal/person"
	"app/internal/user"
	"app/model"
	"os"
	"testing"
)

type TestEnv struct {
	Config *config.Config

	DB            *database.Database
	PersonService *person.PersonService
	UserService   *user.SysUserService

	InvalidPersonService *person.PersonService
	InvalidUserService   *user.SysUserService

	Logger *logger.Logger
}

var testEnv TestEnv

func TestMain(m *testing.M) {
	cfg, err := config.LoadConfig("../config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(*cfg)
	if err != nil {
		panic(err)
	}

	auditRepository := logger.NewAuditRepository(db)
	logger := logger.NewLogger(auditRepository)

	personRepository := internal.NewRepository[model.Person](model.Person{}.TableName(), db, logger)
	personService := person.NewPersonService(personRepository, logger)

	userRepository := internal.NewRepository[model.SysUser](model.SysUser{}.TableName(), db, logger)
	userService := user.NewSysUserService(userRepository, logger)

	mockPersonRepository := &ErrorRepository[model.Person]{}
	invalidPersonService := person.NewPersonService(mockPersonRepository, logger)

	mockUserRepository := &ErrorRepository[model.SysUser]{}
	invalidUserService := user.NewSysUserService(mockUserRepository, logger)

	testEnv = TestEnv{
		Config: cfg,

		DB:            db,
		PersonService: personService,
		UserService:   userService,

		InvalidPersonService: invalidPersonService,
		InvalidUserService:   invalidUserService,

		Logger: logger,
	}

	code := m.Run()

	sqlDB, _ := db.Connection.DB()
	_ = sqlDB.Close()

	os.Exit(code)
}
