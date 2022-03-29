package repository


type FileRepository struct {
	Path string
}



func NewFileRepository(path string) *FileRepository {
	return &FileRepository{
		Path: path,
	}
}


