package storage

import "fmt"

// структура для хранилища
type BaseStorage struct {
	Capacity int    // Общая ёмкость устройства
	Files    []File // Список файлов, сохранённых на устройстве

}

// функция для  вывода списка файлов, находящихся в хранилище.
func (b *BaseStorage) ListFiles() {
	if len(b.Files) == 0 {
		fmt.Println("Устройство пусто")
		return
	}
	fmt.Println("Содержимое:")
	for _, f := range b.Files {
		fmt.Printf("- %s (%d МБ)\n", f.Name, f.Size)
	}
}

func (b *BaseStorage) Format() {
	b.Files = []File{}
	fmt.Println("Устройство отформатировано")
}

// функция для форматирования
func (b *BaseStorage) GetFreeSpace() int {
	used := 0
	for _, f := range b.Files {
		used += f.Size
	}
	return b.Capacity - used
}
