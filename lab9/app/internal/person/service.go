package person

import (
	"app/internal"
	"app/internal/logger"
	"app/model"
	"fmt"
	"regexp"
)

type PersonService struct {
	repository *internal.Repository[model.Person]
	logger     *logger.Logger
}

func NewPersonService(repository *internal.Repository[model.Person], logger *logger.Logger) *PersonService {
	return &PersonService{
		repository: repository,
		logger:     logger,
	}
}

func (s *PersonService) GetAmountQuantity() (int, error) {
	quantity, err := s.repository.GetAmountQuantity()
	if err != nil {
		return 0, err
	}

	return quantity, nil
}

func (s *PersonService) GetByLimitOffset(limit, offset int) ([]model.Person, error) {
	persons, err := s.repository.GetByLimitOffset(limit, offset)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (s *PersonService) GetById(id uint) (*model.Person, error) {
	person, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return person, nil
}

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
		return 0, nil
	}

	return id, nil
}

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

func (s *PersonService) Delete(id uint) error {
	return s.repository.Delete(id)
}
