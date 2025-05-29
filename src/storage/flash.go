package storage

import "errors"

// структура для флешки
type FlashDrive struct {
	BaseStorage
}

// функция для добавления файла
func (f *FlashDrive) AddFile(file File) error {
	if f.GetFreeSpace() < file.Size {
		return errors.New("недостаточно места")
	}
	f.Files = append(f.Files, file)
	return nil
}

// функция для удаления  файла
func (f *FlashDrive) DeleteFile(fileName string) error {
	for i, file := range f.Files {
		if file.Name == fileName {
			f.Files = append(f.Files[:i], f.Files[i+1:]...)
			return nil
		}
	}
	return errors.New("файл не найден")
}
