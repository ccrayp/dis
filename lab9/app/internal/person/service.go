package person

import (
	"app/internal"
	"app/internal/logger"
	"app/model"
	"fmt"
	"regexp"
)

// Сервисный класс для сущности персона
type PersonService struct {
	repository internal.IRepository[model.Person]
	logger     *logger.Logger
}

// Функция для создания нового экземпларя сервиса персон
func NewPersonService(repository internal.IRepository[model.Person], logger *logger.Logger) *PersonService {
	return &PersonService{
		repository: repository,
		logger:     logger,
	}
}

// Метод получения общего количества записей о персонах
func (s *PersonService) GetAmountQuantity() (int, error) {
	quantity, err := s.repository.GetAmountQuantity()
	if err != nil {
		return 0, err
	}

	return quantity, nil
}

// Метод получения записей о персонах с лимитом и сдвигом (по страницам)
func (s *PersonService) GetByLimitOffset(limit, offset int) ([]model.Person, error) {
	persons, err := s.repository.GetByLimitOffset(limit, offset)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

// Метод получения записи по идентификатору персоны
func (s *PersonService) GetById(id uint) (*model.Person, error) {
	person, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// Метод создания записи о персоне
func (s *PersonService) Create(name, phone string, age int) (uint, error) {
	re := regexp.MustCompile(`^(\+7|8)[0-9]{10}$`)

	if !re.MatchString(phone) {
		s.logger.MakeLog("service", "CREATE", "person", "invalid phone number", logger.Warning)
		return 0, fmt.Errorf("Некорретный номер телефона")
	}
	if age <= 0 {
		s.logger.MakeLog("service", "CREATE", "person", "invalid age", logger.Warning)
		return 0, fmt.Errorf("Некорректный возраст")
	}

	id, err := s.repository.Create(model.Person{
		Name:  name,
		Age:   uint(age),
		Phone: phone,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Метод обновления данных о персоне
func (s *PersonService) Update(person *model.Person) error {
	re := regexp.MustCompile(`^(\+7|8)[0-9]{10}$`)

	if !re.MatchString(person.Phone) {
		s.logger.MakeLog("service", "UPDATE", "person", "invalid phone number", logger.Warning)
		return fmt.Errorf("Некорректный номер телефона")
	}
	if person.Age <= 0 {
		s.logger.MakeLog("service", "UPDATE", "person", "invalid age", logger.Warning)
		return fmt.Errorf("Некорректный возраст")
	}

	if err := s.repository.Update(person); err != nil {
		return err
	}

	return nil
}

// Метод удаления записи о персоне
func (s *PersonService) Delete(id uint) error {
	if id <= 0 {
		return fmt.Errorf("Некорректрый идентификатор")
	}

	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
