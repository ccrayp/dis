package person

import (
	"app/utils"
	"fmt"
	"math"
)

type PersonHandler struct {
	service *PersonService
	t       any
}

func NewPersonHandler(service *PersonService) *PersonHandler {
	return &PersonHandler{
		service: service,
	}
}

func (h *PersonHandler) Name() string {
	return "Персона"
}

func (h *PersonHandler) Menu() {
	var choice int

	for {
		utils.Clean()
		fmt.Print("МЕНЮ \"Персона\"\n\n")
		fmt.Println("1. Просмотр по страницам")
		fmt.Println("2. Добавить")
		fmt.Println("3. Изменить")
		fmt.Println("4. Удалить")
		fmt.Println("5. Выход")
		fmt.Print(">>> ")

		fmt.Scan(&choice)

		if choice <= 0 || choice > 5 {
			continue
		} else if choice == 5 {
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
		}
	}
}

func (h *PersonHandler) ShowMenu() {
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
		fmt.Print("Просмотр персон\n\n")
		fmt.Print("1. Предыдущая страница\n2. Следующая страница\n3. Выход\n\n")
		fmt.Printf("Общее число персон: %d\n", q)
		fmt.Printf("Страница %d из %d\n", page, pages)

		persons, err := h.service.GetByLimitOffset(10, (page-1)*10)
		if err != nil {
		}

		fmt.Println()
		for _, person := range persons {
			fmt.Printf("%d. %s %d %s\n", person.Id, person.Name, person.Age, person.Phone)
		}

		fmt.Print("\n>>> ")
		fmt.Scan(&choice)

		if choice < 1 || choice > 3 {
			choice = 0
		}
	}
}

func (h *PersonHandler) CreateMenu() {
	var name, phone string
	var age int

	for {
		utils.Clean()

		fmt.Print("Добавление перосны\n\n")
		fmt.Print("Введите имя\n>>> ")
		fmt.Scanln(&name)
		fmt.Print("Введите возраст\n>>> ")
		fmt.Scan(&age)
		fmt.Print("Введите номер телефона (+7/80000000000)\n>>> ")
		fmt.Scan(&phone)

		id, err := h.service.Create(name, phone, age)
		if err != nil {
			fmt.Printf("\nОшибка создания записи: %s", err.Error())
			utils.Pause()
			continue
		}

		fmt.Printf("\nПерсона успешно добавлена, id новой записи: %d", id)
		utils.Pause()
		break
	}
}

func (h *PersonHandler) UpdateMenu() {
	var id int

	for {
		utils.Clean()

		fmt.Print("Обновление данных о персоне\n\n")
		fmt.Print("Введите идентификатор персоны для поиска\n>>> ")
		fmt.Scan(&id)

		person, err := h.service.GetById(uint(id))
		if err != nil {
			fmt.Printf("\nОшибка поиска записи: %s", err.Error())
			utils.Pause()
			continue
		}

		fmt.Printf("\nВведите имя (%s)\n>>> ", person.Name)
		fmt.Scanln(&person.Name)

		fmt.Printf("Введите возраст (%d)\n>>> ", person.Age)
		fmt.Scan(&person.Age)

		fmt.Printf("Введите номер телефона (%s)\n>>> ", person.Phone)
		fmt.Scan(&person.Phone)

		if err := h.service.Update(person); err != nil {
			fmt.Printf("\nОшибка обновления записи: %s", err.Error())
			utils.Pause()
			continue
		}

		fmt.Printf("\nДанные персоны успешно обновлены")
		utils.Pause()
		return
	}
}

func (h *PersonHandler) DeleteMenu() {
	var id int

	for {
		utils.Clean()

		fmt.Print("Удаление записи о персоне\n\n")
		fmt.Print("Введите идентификатор персоны для поиска\n>>> ")
		fmt.Scan(&id)

		person, err := h.service.GetById(uint(id))
		if err != nil {
			fmt.Printf("\nОшибка поиска записи: %s", err.Error())
			utils.Pause()
			continue
		}

		if err := h.service.Delete(person.Id); err != nil {
			fmt.Printf("\nОшибка удаления записи: %s", err.Error())
			utils.Pause()
			continue
		}

		fmt.Printf("\nЗапись успешно удалена")
		utils.Pause()
		return
	}
}
