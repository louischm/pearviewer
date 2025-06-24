package types

// ReturnCode enum
const (
	Success     int32 = 0
	Fail        int32 = -1
	ServerError int32 = 500
)

// Tree types
var id int32 = 0

type Dir struct {
	name        string
	oldPathName string
	newPathName string
	id          int32
	files       []*File
	children    []*Dir
}

func (d *Dir) NewPathName() string {
	return d.newPathName
}

func (d *Dir) OldPathName() string {
	return d.oldPathName
}

func NewDir(name string, oldPathName string, newPathName string) *Dir {
	dir := &Dir{name: name, oldPathName: oldPathName, newPathName: newPathName, id: id}
	id++
	return dir
}

func (d *Dir) SetFiles(files []*File) {
	d.files = files
}

func (d *Dir) SetChildren(children []*Dir) {
	d.children = children
}

func (d *Dir) Files() []*File {
	return d.files
}

func (d *Dir) Children() []*Dir {
	return d.children
}

func (d *Dir) Name() string {
	return d.name
}

func (d *Dir) SetName(name string) {
	d.name = name
}

func (d *Dir) SetOldPathName(oldPathName string) {
	d.oldPathName = oldPathName
}

func (d *Dir) SetNewPathName(newPathName string) {
	d.newPathName = newPathName
}

func (d *Dir) SetId(id int32) {
	d.id = id
}

type File struct {
	name        string
	oldPathName string
	newPathName string
	id          int32
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
	file := &File{name: name, oldPathName: oldPathName, newPathName: newPathName, id: id}
	id++
	return file
}

func (f *File) SetOldPathName(oldPathName string) {
	f.oldPathName = oldPathName
}

func (f *File) SetNewPathName(newPathName string) {
	f.newPathName = newPathName
}

func (f *File) SetId(id int32) {
	f.id = id
}

func (f *File) SetName(name string) {
	f.name = name
}
