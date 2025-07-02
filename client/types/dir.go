package types

type Dir struct {
	name        string
	oldPathName string
	newPathName string
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
	dir := &Dir{name: name, oldPathName: oldPathName, newPathName: newPathName}
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
