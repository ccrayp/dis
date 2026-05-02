package model

type SysUser struct {
	Id           uint
	PersonId     uint
	PasswordHash string
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (s SysUser) GetId() uint {
	return s.Id
}
