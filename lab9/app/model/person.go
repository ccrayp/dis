package model

type Person struct {
	Id    uint
	Name  string
	Age   uint
	Phone string
}

func (Person) TableName() string {
	return "person"
}

func (p Person) GetId() uint {
	return p.Id
}
