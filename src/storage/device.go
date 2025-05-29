package storage

// интерфейс который  определяет поведение для всех типов устройств хранения данных в
// в го утиная типизация тоесть если у структуры
// есть функции одинаковые с интерфейсом то она уже реализует его
type StorageDevice interface {
	AddFile(file File) error          // AddFile добавляет файл в хранилище.
	DeleteFile(fileName string) error // DeleteFile удаляет файл из хранилища по имени
	GetFreeSpace() int                // GetFreeSpace возвращает количество свободного места
	ListFiles()                       // ListFiles выводит список всех файлов в хранилище
	Format()                          // Format очищает хранилище, удаляя все файлы
}
