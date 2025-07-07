package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"path/filepath"
	"pearviewer/client/grpc"
	pb "pearviewer/generated"
	"strings"
)

var selectedNode = ""
var data *pb.ListDirRes

func refreshData(username string) *pb.ListDirRes {
	rootPathName := grpc.GetRootPath(username)
	return grpc.ListDir("", rootPathName.PathName)
}

func homePage(w fyne.Window, username string) {

	rootPathName := grpc.GetRootPath(username)
	data = grpc.ListDir("", rootPathName.PathName)
	selectedNode = rootPathName.PathName

	// Dir Tree
	tree := createTree()
	tree.OnSelected = func(id string) {
		selectedNode = id
		tree.Refresh()
		log.Info("Selected node: %s", selectedNode)
	}
	tree.OnUnselected = func(id string) {
		selectedNode = rootPathName.PathName
		tree.Refresh()
		log.Info("Unselected node:")
	}

	// Button bar
	gridButton := createGridButton(tree, username)

	logOff := widget.NewButton("Log Off", func() {
		login(w)
	})
	content := container.NewBorder(
		gridButton, // top
		logOff,     // bottom
		nil,        // left
		nil,        // right
		tree,       // center
	)

	w.SetContent(content)
}

func createUser(w fyne.Window) {
	title := canvas.NewText("Create your account", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 32

	username := widget.NewEntry()
	username.SetPlaceHolder("Username")
	usernameWrapped := NewEntryWrapper(username, fyne.NewSize(300, 40))

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	passwordWrapped := NewEntryWrapper(password, fyne.NewSize(300, 40))

	submit := widget.NewButton("Submit", func() {
		_, err := grpc.CreateUser(username.Text, password.Text)
		if err != nil {
			log.Debug("create user error: %v", err)
		} else {
			log.Info("create user success")
			login(w)
		}

	})
	submit.Importance = widget.HighImportance

	back := widget.NewButton("Back", func() {
		log.Info("back")
		login(w)
	})

	form := container.NewVBox(
		title,
		usernameWrapped,
		passwordWrapped,
		submit,
		back,
	)
	centered := container.NewCenter(form)

	w.SetContent(centered)
}

func dataNameToUserName(dirName string) string {
	lastIndex := strings.LastIndex(dirName, "/")
	return dirName[lastIndex+1:]
}

func pathNameToDirName(pathName string) string {
	lastIndex := strings.LastIndex(pathName, "/")
	return pathName[:lastIndex]
}

func createRootNode() []widget.TreeNodeID {
	nodes := make([]widget.TreeNodeID, 0)

	for _, dir := range data.Dir.Dir {
		nodes = append(nodes, dir.FullName)
	}

	for _, file := range data.Dir.File {
		nodes = append(nodes, file.FullName)
	}
	return nodes
}

func getBranch(data *pb.Dir, dirName string) *pb.Dir {
	for _, dir := range data.Dir {
		if dirName == dir.FullName {
			return dir
		}
		return getBranch(dir, dirName)
	}
	return nil
}

func createGridButton(tree *widget.Tree, username string) *fyne.Container {
	addDirInput := widget.NewEntry()
	addDirInput.SetPlaceHolder("Name your directory")
	addDirButton := widget.NewButton("Create dir", func() {
		grpc.CreateDir(addDirInput.Text, selectedNode)
		log.Info("Create dir")
		data = refreshData(username)
		tree.Refresh()
	})
	addDir := container.NewVBox(addDirInput, addDirButton)

	uploadFileInput := widget.NewEntry()
	uploadFileInput.SetPlaceHolder("Upload file")
	uploadFileButton := widget.NewButton("Upload File", func() {
		grpc.UploadFile(uploadFileInput.Text, selectedNode)
		log.Info("Upload file")
		data = refreshData(username)
		tree.Refresh()
	})
	uploadFile := container.NewVBox(uploadFileInput, uploadFileButton)

	downloadFileInput := widget.NewEntry()
	downloadFileInput.SetPlaceHolder("Download file")
	downloadFileButton := widget.NewButton("Download File", func() {
		grpc.DownloadFile(filepath.Base(selectedNode), pathNameToDirName(selectedNode), downloadFileInput.Text)
		log.Info("Download file")
		data = refreshData(username)
		tree.Refresh()
	})
	downloadFile := container.NewVBox(downloadFileInput, downloadFileButton)
	return container.New(layout.NewGridLayout(3), addDir, uploadFile, downloadFile)
}

func createTree() *widget.Tree {
	return widget.NewTree(
		// ChildUIDs
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			// Create Root
			// Else if create sub dir
			// Else create file
			if id == "" {
				return createRootNode()
			} else if !strings.Contains(id, ".") {
				nodes := make([]widget.TreeNodeID, 0)

				branch := getBranch(data.Dir, id)

				if branch != nil {
					for _, file := range branch.File {
						if strings.Contains(file.FullName, id) {
							nodes = append(nodes, file.FullName)
						}
					}

					for _, dir := range branch.Dir {
						if strings.Contains(dir.FullName, id) {
							nodes = append(nodes, dir.FullName)
						}
					}
				}
				return nodes
			} else {
				nodes := make([]widget.TreeNodeID, 0)
				nodes = append(nodes, id)
				return nodes
			}
			return []string{}
		},
		// IsBranch
		func(id widget.TreeNodeID) bool {
			if !strings.Contains(id, ".") {
				return true
			}
			return false
		},
		// CreateNode
		func(branch bool) fyne.CanvasObject {
			if branch {
				icon := widget.NewIcon(theme.FolderIcon())
				label := widget.NewLabel("Dir")
				return container.NewHBox(icon, label)
			}
			icon := widget.NewIcon(theme.FileIcon())
			label := widget.NewLabel("File")
			label.Importance = widget.LowImportance
			return container.NewHBox(icon, label)
		},
		// UpdateNode
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			icon := widget.NewIcon(nil)
			label := o.(*fyne.Container).Objects[1].(*widget.Label)
			text := id
			text = filepath.Base(text)
			label.Importance = widget.LowImportance
			label.SetText(text)

			if branch {
				icon.SetResource(theme.FolderIcon())
			} else {
				icon.SetResource(theme.FileIcon())
			}
			o.(*fyne.Container).Objects[0] = icon

			if id == selectedNode {
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.Refresh()
			} else {
				label.TextStyle = fyne.TextStyle{} // normal
				label.Refresh()
			}

			/*o.(*widget.Label).SetText(text)
			o.(*widget.Label).Importance = widget.LowImportance*/
		})
}
