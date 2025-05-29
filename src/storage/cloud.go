package storage

import (
	"errors"
	"fmt"
)

// структура для облака в нее встраиваем структуру BaseStorage(наследование)
type CloudDrive struct {
	BaseStorage
}

// функуия для добавления файла
func (c *CloudDrive) AddFile(file File) error {
	c.Files = append(c.Files, file)
	fmt.Println("Файл загружен в облако")
	return nil
}

// функуия для удаления файла
func (c *CloudDrive) DeleteFile(fileName string) error {
	for i, file := range c.Files {
		if file.Name == fileName {
			c.Files = append(c.Files[:i], c.Files[i+1:]...)
			return nil
		}
	}
	return errors.New("файл не найден в облаке")
}
