package user

import (
	"app/internal"
	"app/internal/logger"
	"app/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type SysUserService struct {
	repository *internal.Repository[model.SysUser]
	logger     *logger.Logger
}

func NewSysUserService(repository *internal.Repository[model.SysUser], logger *logger.Logger) *SysUserService {
	return &SysUserService{
		repository: repository,
		logger:     logger,
	}
}

func (s *SysUserService) GetAmountQuantity() (int, error) {
	return s.repository.GetAmountQuantity()
}

func (s *SysUserService) GetByLimitOffset(limit, offset int) ([]model.SysUser, error) {
	return s.repository.GetByLimitOffset(limit, offset)
}

func (s *SysUserService) GetById(id uint) (*model.SysUser, error) {
	return s.repository.GetById(id)
}

func (s *SysUserService) Create(personID uint, password string) (uint, error) {
	if personID == 0 {
		s.logger.MakeLog("service", "CREATE", "user", "invalid person_id", logger.Warning)
		return 0, fmt.Errorf("Некорректный id персоны")
	}
	if len(password) < 4 {
		s.logger.MakeLog("service", "CREATE", "user", "password have to contain at least 4 chars", logger.Warning)
		return 0, fmt.Errorf("Пароль должен содержать минимум 4 символа")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	return s.repository.Create(model.SysUser{
		PersonId:     personID,
		PasswordHash: string(hash),
	})
}

func (s *SysUserService) Update(user *model.SysUser, password string) error {
	if user.PersonId == 0 {
		s.logger.MakeLog("service", "UPDATE", "user", "invalid person_id", logger.Warning)
		return fmt.Errorf("Некорректный id персоны")
	}
	if len(password) < 4 {
		s.logger.MakeLog("service", "UPDATE", "user", "password have to contain at least 4 chars", logger.Warning)
		return fmt.Errorf("Пароль должен содержать минимум 4 символа")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hash)

	return s.repository.Update(user)
}

func (s *SysUserService) Delete(id uint) error {
	return s.repository.Delete(id)
}

func (s *SysUserService) CheckPassword(id uint, password string) (bool, error) {
	user, err := s.repository.GetById(id)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		s.logger.MakeLog("service", "CHECK_PASSWORD", "password", "passwords do not mathc", logger.Warning)
	}

	return err == nil, nil
}
