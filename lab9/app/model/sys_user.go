package model

// Структура описывающая модель данных пользователя
type SysUser struct {
	Id           uint
	PersonId     uint
	PasswordHash string
}

// Метод получения имени таблицы пользователя
func (SysUser) TableName() string {
	return "user"
}

// Метод получения идентификатора пользователя
func (s SysUser) GetId() uint {
	return s.Id
}
