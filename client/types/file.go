package types

type File struct {
	name        string
	oldPathName string
	newPathName string
}

func (f *File) Name() string {
	return f.name
}

func (f *File) OldPathName() string {
	return f.oldPathName
}

func (f *File) NewPathName() string {
	return f.newPathName
}

func NewFile(name string, oldPathName string, newPathName string) *File {
	file := &File{name: name, oldPathName: oldPathName, newPathName: newPathName}
	return file
}

func (f *File) SetOldPathName(oldPathName string) {
	f.oldPathName = oldPathName
}

func (f *File) SetNewPathName(newPathName string) {
	f.newPathName = newPathName
}

func (f *File) SetName(name string) {
	f.name = name
}
