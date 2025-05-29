package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"oop-go/src/storage
)

// RunCLI запускает консольку для выбора
func RunCLI() {
	reader := bufio.NewReader(os.Stdin)
	// Выбор типа устройства
	fmt.Println("Выберите тип устройства:")
	fmt.Println("1 — Флешка")
	fmt.Println("2 — Облако")
	fmt.Print("Введите номер: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var dev storage.StorageDevice
	// Инициализация выбранного устройства
	switch input {
	case "1":
		dev = &storage.FlashDrive{BaseStorage: storage.BaseStorage{Capacity: 1000}}
	case "2":
		dev = &storage.CloudDrive{BaseStorage: storage.BaseStorage{Capacity: 500}}
	default:
		fmt.Println("Неверный выбор")
		return
	}
	// Главное меню
	for {
		fmt.Println("\nМеню:")
		fmt.Println("1 — Добавить файл")
		fmt.Println("2 — Удалить файл")
		fmt.Println("3 — Показать файлы")
		fmt.Println("4 — Показать свободное место")
		fmt.Println("5 — Форматировать устройство")
		fmt.Println("0 — Выход")
		fmt.Print("Выбор: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		// Добавление файла
		case "1":
			fmt.Print("Введите имя файла: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Введите размер файла (МБ): ")
			sizeStr, _ := reader.ReadString('\n')
			sizeStr = strings.TrimSpace(sizeStr)
			size, err := strconv.Atoi(sizeStr)
			if err != nil || size <= 0 {
				fmt.Println("Некорректный размер")
				continue
			}

			err = dev.AddFile(storage.File{Name: name, Size: size})
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Файл добавлен")
			}
			// Удаление файла
		case "2":
			fmt.Print("Введите имя файла для удаления: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			err := dev.DeleteFile(name)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Файл удалён")
			}
			// Показать список файлов
		case "3":
			dev.ListFiles()

			// Показать свободное место
		case "4":
			fmt.Printf("Свободное место: %d МБ\n", dev.GetFreeSpace())

			// Форматировать устройство
		case "5":
			dev.Format()

			// Выход из программы
		case "0":
			fmt.Println("Выход.")
			return
			// Обработка неправильного ввода
		default:
			fmt.Println("Неверный ввод")
		}
	}
}
