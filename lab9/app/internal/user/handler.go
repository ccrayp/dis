package user

import (
	"app/utils"
	"fmt"
	"math"
)

type SysUserHandler struct {
	service *SysUserService
}

func NewSysUserHandler(service *SysUserService) *SysUserHandler {
	return &SysUserHandler{
		service: service,
	}
}

func (h *SysUserHandler) Name() string {
	return "Системный пользователь"
}

func (h *SysUserHandler) Menu() {
	var choice int

	for {
		utils.Clean()
		fmt.Print("МЕНЮ \"Системный пользователь\"\n\n")
		fmt.Println("1. Просмотр по страницам")
		fmt.Println("2. Добавить")
		fmt.Println("3. Изменить")
		fmt.Println("4. Удалить")
		fmt.Println("5. Проверить пароль")
		fmt.Println("6. Выход")
		fmt.Print(">>> ")

		fmt.Scan(&choice)

		if choice <= 0 || choice > 6 {
			continue
		} else if choice == 6 {
			break
		}

		switch choice {
		case 1:
			h.ShowMenu()
		case 2:
			h.CreateMenu()
		case 3:
			h.UpdateMenu()
		case 4:
			h.DeleteMenu()
		case 5:
			h.CheckPasswordMenu()
		}
	}
}

func (h *SysUserHandler) ShowMenu() {
	choice := 0
	page := 1

	for {
		q, _ := h.service.GetAmountQuantity()
		pages := int(math.Ceil(float64(q) / 10.0))

		switch choice {
		case 1:
			if page > 1 {
				page--
			}
		case 2:
			if page < pages {
				page++
			}
		case 3:
			return
		}

		utils.Clean()
		fmt.Print("Просмотр системных пользователей\n\n")
		fmt.Print("1. Предыдущая страница\n2. Следующая страница\n3. Выход\n\n")
		fmt.Printf("Общее число пользователей: %d\n", q)
		fmt.Printf("Страница %d из %d\n", page, pages)

		users, err := h.service.GetByLimitOffset(10, (page-1)*10)
		if err != nil {
		}

		fmt.Println()
		for _, user := range users {
			fmt.Printf("%d. Идентификатор персоны=%d hash=%s\n", user.Id, user.PersonId, user.PasswordHash)
		}

		fmt.Print("\n>>> ")
		fmt.Scan(&choice)

		if choice < 1 || choice > 3 {
			choice = 0
		}
	}
}

func (h *SysUserHandler) CreateMenu() {
	var personID int
	var password string

	for {
		utils.Clean()

		fmt.Print("Добавление системного пользователя\n\n")
		fmt.Print("Введите id персоны\n>>> ")
		fmt.Scan(&personID)
		fmt.Print("Введите пароль\n>>> ")
		fmt.Scan(&password)

		id, err := h.service.Create(uint(personID), password)
		if err != nil {
			fmt.Printf("\nОшибка создания записи: %s", err.Error())
			utils.Pause()
			continue
		}

		fmt.Printf("\nСистемный пользователь успешно добавлен, id новой записи: %d", id)
		utils.Pause()
		return
	}
}

func (h *SysUserHandler) UpdateMenu() {
	var id int
	var password string

	for {
		utils.Clean()

		fmt.Print("Обновление данных системного пользователя\n\n")
		fmt.Print("Введите идентификатор пользователя для поиска\n>>> ")
		fmt.Scan(&id)

		user, err := h.service.GetById(uint(id))
		if err != nil {
			fmt.Printf("\nОшибка поиска записи: %s", err.Error())
			utils.Pause()
			continue
		}

		fmt.Printf("\nВведите id персоны (%d)\n>>> ", user.PersonId)
		fmt.Scan(&user.PersonId)

		fmt.Printf("Введите новый пароль\n>>> ")
		fmt.Scan(&password)

		if err := h.service.Update(user, password); err != nil {
			fmt.Printf("\nОшибка обновления записи: %s", err.Error())
			utils.Pause()
			continue
		}

		fmt.Printf("\nДанные системного пользователя успешно обновлены")
		utils.Pause()
		return
	}
}

func (h *SysUserHandler) DeleteMenu() {
	var id int

	for {
		utils.Clean()

		fmt.Print("Удаление системного пользователя\n\n")
		fmt.Print("Введите идентификатор пользователя для поиска\n>>> ")
		fmt.Scan(&id)

		user, err := h.service.GetById(uint(id))
		if err != nil {
			fmt.Printf("\nОшибка поиска записи: %s", err.Error())
			utils.Pause()
			continue
		}

		if err := h.service.Delete(user.Id); err != nil {
			fmt.Printf("\nОшибка удаления записи: %s", err.Error())
			utils.Pause()
			continue
		}

		fmt.Printf("\nЗапись успешно удалена")
		utils.Pause()
		return
	}
}

func (h *SysUserHandler) CheckPasswordMenu() {
	var id int
	var password string

	for {
		utils.Clean()

		fmt.Print("Проверка пароля системного пользователя\n\n")
		fmt.Print("Введите идентификатор пользователя\n>>> ")
		fmt.Scan(&id)

		fmt.Print("Введите пароль\n>>> ")
		fmt.Scan(&password)

		ok, err := h.service.CheckPassword(uint(id), password)
		if err != nil {
			fmt.Printf("\nОшибка проверки пароля: %s", err.Error())
			utils.Pause()
			continue
		}

		if ok {
			fmt.Printf("\nПароль верный")
		} else {
			fmt.Printf("\nПароль неверный")
		}

		utils.Pause()
		return
	}
}
