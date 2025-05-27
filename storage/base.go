package storage

import "fmt"

type BaseStorage struct {
	Capacity int
	Files    []File
}

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

func (b *BaseStorage) GetFreeSpace() int {
	used := 0
	for _, f := range b.Files {
		used += f.Size
	}
	return b.Capacity - used
}
