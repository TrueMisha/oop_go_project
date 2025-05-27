package storage

type StorageDevice interface {
	AddFile(file File) error
	DeleteFile(fileName string) error
	GetFreeSpace() int
	ListFiles()
	Format()
}
