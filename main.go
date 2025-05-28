//package main
//
//import (
//	"bufio"
//	"errors"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//// ---------- Абстракция ----------
//// Интерфейс описывает поведение хранилища, не раскрывая реализацию.
//// Любое устройство (флешка, облако) может реализовать этот интерфейс.
//type StorageDevice interface {
//	AddFile(file File) error
//	DeleteFile(fileName string) error
//	GetFreeSpace() int
//	ListFiles()
//	Format()
//}
//
//// ---------- Инкапсуляция ----------
//// Структура файла инкапсулирует имя и размер.
//type File struct {
//	Name string
//	Size int
//}
//
//// ---------- Наследование ----------
//// Базовая структура, содержащая общее поведение.
//// Используется как встраиваемая часть (композиция = наследование в Go).
//type BaseStorage struct {
//	Capacity int
//	Files    []File
//}
//
//// Методы, общие для всех типов устройств.
//func (b *BaseStorage) ListFiles() {
//	if len(b.Files) == 0 {
//		fmt.Println("Устройство пусто")
//		return
//	}
//	fmt.Println("Содержимое:")
//	for _, f := range b.Files {
//		fmt.Printf("- %s (%d МБ)\n", f.Name, f.Size)
//	}
//}
//
//func (b *BaseStorage) Format() {
//	b.Files = []File{}
//	fmt.Println("Устройство отформатировано")
//}
//
//func (b *BaseStorage) GetFreeSpace() int {
//	used := 0
//	for _, f := range b.Files {
//		used += f.Size
//	}
//	return b.Capacity - used
//}
//
//// ---------- FlashDrive (Флешка) ----------
//// Наследует BaseStorage — повторно использует логику и данные.
//type FlashDrive struct {
//	BaseStorage
//}
//
//// ---------- Полиморфизм ----------
//// Реализует интерфейс StorageDevice — можно использовать как абстрактное хранилище.
//func (f *FlashDrive) AddFile(file File) error {
//	if f.GetFreeSpace() < file.Size {
//		return errors.New("недостаточно места")
//	}
//	f.Files = append(f.Files, file)
//	return nil
//}
//
//func (f *FlashDrive) DeleteFile(fileName string) error {
//	for i, file := range f.Files {
//		if file.Name == fileName {
//			f.Files = append(f.Files[:i], f.Files[i+1:]...)
//			return nil
//		}
//	}
//	return errors.New("файл не найден")
//}
//
//// ---------- CloudDrive (Облако) ----------
//// Также наследует поведение от BaseStorage.
//type CloudDrive struct {
//	BaseStorage
//}
//
//// Реализация интерфейса StorageDevice (другое поведение).
//func (c *CloudDrive) AddFile(file File) error {
//	c.Files = append(c.Files, file)
//	fmt.Println("Файл загружен в облако")
//	return nil
//}
//
//func (c *CloudDrive) DeleteFile(fileName string) error {
//	for i, file := range c.Files {
//		if file.Name == fileName {
//			c.Files = append(c.Files[:i], c.Files[i+1:]...)
//			return nil
//		}
//	}
//	return errors.New("файл не найден в облаке")
//}
//
//// ---------- Главная программа с вводом ----------
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//
//	fmt.Println("Выберите тип устройства:")
//	fmt.Println("1 — Флешка")
//	fmt.Println("2 — Облако")
//	fmt.Print("Введите номер: ")
//	input, _ := reader.ReadString('\n')
//	input = strings.TrimSpace(input)
//
//	var storage StorageDevice // ---------- Полиморфизм ----------
//	// Одна переменная для двух разных реализаций интерфейса.
//
//	switch input {
//	case "1":
//		// Используем FlashDrive
//		storage = &FlashDrive{BaseStorage{Capacity: 1000}}
//	case "2":
//		// Используем CloudDrive
//		storage = &CloudDrive{BaseStorage{Capacity: 500}}
//	default:
//		fmt.Println("Неверный выбор")
//		return
//	}
//
//	for {
//		fmt.Println("\nМеню:")
//		fmt.Println("1 — Добавить файл")
//		fmt.Println("2 — Удалить файл")
//		fmt.Println("3 — Показать файлы")
//		fmt.Println("4 — Показать свободное место")
//		fmt.Println("5 — Форматировать устройство")
//		fmt.Println("0 — Выход")
//		fmt.Print("Выбор: ")
//
//		choice, _ := reader.ReadString('\n')
//		choice = strings.TrimSpace(choice)
//
//		switch choice {
//		case "1":
//			fmt.Print("Введите имя файла: ")
//			name, _ := reader.ReadString('\n')
//			name = strings.TrimSpace(name)
//
//			fmt.Print("Введите размер файла (МБ): ")
//			sizeStr, _ := reader.ReadString('\n')
//			sizeStr = strings.TrimSpace(sizeStr)
//			size, err := strconv.Atoi(sizeStr)
//			if err != nil || size <= 0 {
//				fmt.Println("Некорректный размер")
//				continue
//			}
//
//			err = storage.AddFile(File{Name: name, Size: size})
//			if err != nil {
//				fmt.Println("Ошибка:", err)
//			} else {
//				fmt.Println("Файл добавлен")
//			}
//
//		case "2":
//			fmt.Print("Введите имя файла для удаления: ")
//			name, _ := reader.ReadString('\n')
//			name = strings.TrimSpace(name)
//
//			err := storage.DeleteFile(name)
//			if err != nil {
//				fmt.Println("Ошибка:", err)
//			} else {
//				fmt.Println("Файл удалён")
//			}
//
//		case "3":
//			storage.ListFiles()
//
//		case "4":
//			fmt.Printf("Свободное место: %d МБ\n", storage.GetFreeSpace())
//
//		case "5":
//			storage.Format()
//
//		case "0":
//			fmt.Println("Выход.")
//			return
//
//		default:
//			fmt.Println("Неверный ввод")
//		}
//	}
//}

package main

import (
	"oop-go/ui"
)

func main() {
	ui.RunCLI()
}
