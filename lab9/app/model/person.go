package model

// Структура описывающая модель данных персоны
type Person struct {
	Id    uint
	Name  string
	Age   uint
	Phone string
}

// Метод получения имени таблицы перосны
func (Person) TableName() string {
	return "person"
}

// Метод получения идентификатора персоны
func (p Person) GetId() uint {
	return p.Id
}
