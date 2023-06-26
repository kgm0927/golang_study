package structandinterface

type FileSystem interface {
	Rename(oldpath, newpath string) error
	Remove(name string) error
}
type OSFileSystem struct {
}

func (fs OSFileSystem) Rename(oldpath, newpath string) {

}
