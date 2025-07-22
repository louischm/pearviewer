package ui

import (
	"bufio"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/louischm/pkg/utils"
	"image/color"
	"os"
	"path/filepath"
	"pearviewer/client/grpc"
	pb "pearviewer/generated"
	"runtime"
	"strconv"
	"strings"
)

var rootPathName *pb.GetRootPathRes
var selectedNode = ""
var data *pb.ListDirRes
var loadFileText *canvas.Text
var loadDirText *canvas.Text
var loadFileBar *widget.ProgressBar
var loadDirBar *widget.ProgressBar
var searchBar *widget.Entry

func refreshData(username string) *pb.ListDirRes {
	rootPathName = grpc.GetRootPath(username)
	if searchBar.Text == "" {
		return grpc.ListDir("", rootPathName.PathName)
	} else {
		return grpc.SearchFile(searchBar.Text, "", rootPathName.PathName)
	}

}

func createTopContainer(text string) *fyne.Container {
	title := canvas.NewText(text, color.NRGBA{255, 255, 255, 255})
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: false, Italic: false}
	title.TextSize = 16

	searchBar = widget.NewEntry()
	searchBar.SetPlaceHolder("Search")

	topContainer := container.NewVBox(title, searchBar)
	return topContainer
}

func homePage(w fyne.Window, username string) {
	// Top container (title, searchBar)
	topContainer := createTopContainer("Pearviewer")

	rootPathName = grpc.GetRootPath(username)
	data = grpc.ListDir("", rootPathName.PathName)
	selectedNode = rootPathName.PathName

	// Dir Server Tree
	treeServer := createTreeServer()
	// Button bar
	gridServerButton := createServerButton(treeServer, username)
	delBtn := gridServerButton.Objects[4].(*fyne.Container).Objects[2].(*widget.Button)
	dwnBtn := gridServerButton.Objects[4].(*fyne.Container).Objects[3].(*widget.Button)
	loadDirText = gridServerButton.Objects[0].(*canvas.Text)
	loadDirBar = gridServerButton.Objects[1].(*widget.ProgressBar)
	loadFileText = gridServerButton.Objects[2].(*canvas.Text)
	loadFileBar = gridServerButton.Objects[3].(*widget.ProgressBar)

	serverContent := container.NewBorder(
		topContainer,     // top
		gridServerButton, // bottom | Loader et bulle info
		nil,              // left
		nil,              // right
		treeServer,       // center
	)

	treeServer.OnSelected = func(id string) {
		if id == "" {
			selectedNode = rootPathName.PathName
			delBtn.Disable()
			dwnBtn.Disable()
		} else {
			selectedNode = id
			delBtn.Enable()
			dwnBtn.Enable()
		}
		treeServer.Refresh()
		log.Info("Selected node: %s", selectedNode)
	}

	searchBar.OnChanged = func(s string) {
		data = refreshData(username)
		treeServer.Refresh()
	}

	// Tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Server", serverContent),
		container.NewTabItem("My files", widget.NewLabel("My files")),
	)
	w.SetContent(tabs)
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
		dirRec := getBranch(dir, dirName)
		if dirRec != nil {
			return dirRec
		}
	}
	return nil
}

func createServerButton(tree *widget.Tree, username string) *fyne.Container {
	// Loader
	fileText := canvas.NewText("", color.White)
	fileText.Alignment = fyne.TextAlignLeading
	fileText.TextSize = 14
	dirText := canvas.NewText("", color.White)
	dirText.Alignment = fyne.TextAlignLeading
	dirText.TextSize = 14
	progressDir := widget.NewProgressBar()
	progressFile := widget.NewProgressBar()
	progressDir.Hide()
	progressFile.Hide()

	// Button
	addDir := createAddDirContainer(tree, username)
	rename := createRenameContainer(tree, username)
	deleteBtn := createDeleteContainer(tree, username)
	downLoad := createDownloadContainer(tree, username)
	deselect := createDeselectContainer(tree)
	gridLayout := container.New(layout.NewGridLayout(5), addDir, rename, deleteBtn, downLoad, deselect)

	/*progress.Hide()*/
	return container.NewVBox(dirText, progressDir, fileText, progressFile, gridLayout)
}

func createRenameContainer(tree *widget.Tree, username string) *fyne.Container {
	renameInput := widget.NewEntry()
	renameInput.SetPlaceHolder("Rename selection")
	renameButton := widget.NewButton("Rename", func() {
		oldName := filepath.Base(selectedNode)
		pathName := pathNameToDirName(selectedNode)
		if strings.Contains(renameInput.Text, ".") {
			grpc.RenameFile(oldName, renameInput.Text, pathName)
			log.Info("rename file success")
			setLoadText("Renamed file " + oldName + " to " + renameInput.Text)
		} else {
			grpc.RenameDir(oldName, renameInput.Text, pathName)
			log.Info("rename dir success")
			setLoadText("Renamed directory " + oldName + " to " + renameInput.Text)
		}
		data = refreshData(username)
		tree.Refresh()
	})
	renameButton.Disable()
	renameInput.OnChanged = func(s string) {
		if s == "" || selectedNode == rootPathName.PathName {
			renameButton.Disable()
		} else if strings.Contains(s, ".") && strings.Contains(selectedNode, ".") ||
			!strings.Contains(s, ".") && !strings.Contains(selectedNode, ".") {
			renameButton.Enable()
		} else {
			renameButton.Disable()
		}
	}
	return container.NewVBox(renameButton, renameInput)
}

func createDeleteContainer(tree *widget.Tree, username string) *widget.Button {
	deleteButton := widget.NewButton("Delete", func() {
		name := filepath.Base(selectedNode)
		pathName := pathNameToDirName(selectedNode)
		if strings.Contains(selectedNode, ".") {
			grpc.DeleteFile(name, pathName)
			log.Info("delete file success")
			setLoadText("File " + name + " deleted")
		} else {
			grpc.DeleteDir(name, pathName)
			log.Info("delete dir success")
			setLoadText("Directory " + name + " deleted")
		}
		data = refreshData(username)
		tree.Refresh()
	})
	deleteButton.Disable()
	return deleteButton
}

func createAddDirContainer(tree *widget.Tree, username string) *fyne.Container {
	addDirInput := widget.NewEntry()
	addDirInput.SetPlaceHolder("Name your directory")
	addDirButton := widget.NewButton("Create dir", func() {
		grpc.CreateDir(addDirInput.Text, selectedNode)
		log.Info("Create dir")
		setLoadText("Directory " + addDirInput.Text + " created")
		data = refreshData(username)
		tree.Refresh()
	})
	addDirButton.Disable()

	addDirInput.OnChanged = func(s string) {
		if s == "" || strings.Contains(s, ".") {
			addDirButton.Disable()
		} else {
			addDirButton.Enable()
		}
	}
	return container.NewVBox(addDirButton, addDirInput)
}

func createUploadContainer(tree *widget.Tree, username string) *fyne.Container {
	uploadFileInput := widget.NewEntry()
	uploadFileInput.SetPlaceHolder("Upload selected file")
	uploadFileButton := widget.NewButton("Upload File", func() {
		grpc.UploadFile(uploadFileInput.Text, selectedNode)
		log.Info("Upload file")
		data = refreshData(username)
		tree.Refresh()
	})
	return container.NewVBox(uploadFileButton, uploadFileInput)
}

func resetLoadBars() {
	loadFileText.Text = ""
	loadDirText.Text = ""
	loadFileBar.Hide()
	loadDirBar.Hide()
	loadFileText.Hide()
	loadDirText.Hide()
}

func createDownloadContainer(tree *widget.Tree, username string) *widget.Button {
	downloadButton := widget.NewButton("Download", func() {

		resetLoadBars()
		fileName := filepath.Base(selectedNode)
		pathName := pathNameToDirName(selectedNode)

		if strings.Contains(selectedNode, ".") {
			done := make(chan bool)
			loadFileText.Show()
			loadFileBar.Show()
			go DownloadFile(fileName, pathName, getDownloadPathName(), done)
			<-done
		} else {
			loadFileText.Show()
			loadFileBar.Show()
			loadDirText.Show()
			loadDirBar.Show()
			go downloadDir(fileName, pathName)
		}
		data = refreshData(username)
		tree.Refresh()
	})
	downloadButton.Disable()
	return downloadButton
}

func DownloadFile(fileName, sourcePathName, destPathName string, doneFile chan bool) {
	done := make(chan bool, 1)
	maxSize, err := grpc.GetFileSize(fileName, sourcePathName)
	if err != nil {
		log.Debug("Get file size error: " + err.Error())
		return
	}
	loadFileBar.Min = 1
	loadFileBar.Max = float64(maxSize)
	size := make(chan float64, maxSize/1000+1)

	go grpc.DownloadFile(fileName, sourcePathName, destPathName, done, size)

	for j := 1; j <= int(maxSize/1000)+1; j++ {
		val := <-size
		fyne.Do(func() {
			loadFileBar.SetValue(val)
			log.Info("FILE SIZE %d", loadFileBar.Value)

		})
	}

	<-done
	strValue := strconv.Itoa(int(loadFileBar.Max))
	log.Info("Download file success")
	loadFileText.Text = "File " + fileName + " downloaded: " + strValue + " bytes"
	doneFile <- true
}

func downloadDir(dirName, pathName string) {
	listDir := grpc.ListDir(dirName, pathName)
	fileNumber, err := grpc.GetFileNumber(dirName, pathName)
	if err != nil {
		loadDirText.Text = err.Error()
		return
	}
	log.Info("Number of files to download: %d", fileNumber.Number)

	loadDirBar.Min = 1
	loadDirBar.Max = float64(fileNumber.Number)
	fileDwnld := make(chan float64, fileNumber.Number)

	go downloadFilesInDir(listDir.Dir, pathName, getDownloadPathName(), fileDwnld, 0)

	for j := 1; j <= int(fileNumber.Number); j++ {
		val := <-fileDwnld
		fyne.Do(func() {
			loadDirBar.SetValue(val)
		})
	}
	strValue := strconv.Itoa(int(loadDirBar.Max))
	log.Info("Download dir success")
	loadDirText.Text = "Directory " + dirName + " downloaded: " + strValue + "/" +
		strconv.Itoa(int(fileNumber.Number))
}

func downloadFilesInDir(dir *pb.Dir, sourcePathName, destPathName string, fileDwnld chan float64, fileNumber int) {
	grpc.CreateSourceDir(dir, destPathName)
	sourceName := utils.Joins(sourcePathName, dir.DirName)
	destName := utils.Joins(destPathName, dir.DirName)

	for _, file := range dir.GetFile() {
		done := make(chan bool)
		go DownloadFile(file.Name, sourceName, destName, done)
		fileNumber++
		<-done
		fileDwnld <- float64(fileNumber)
	}

	for _, child := range dir.GetDir() {
		downloadFilesInDir(child, sourceName, destName, fileDwnld, fileNumber)
	}
}

func createDeselectContainer(tree *widget.Tree) *widget.Button {
	return widget.NewButton("Deselect", func() {
		tree.OnSelected("")
	})
}

func getDownloadPathName() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	switch runtime.GOOS {
	case "windows":
		return filepath.Join(home, "Downloads")
	case "linux":
		return getLinuxDownloadDir()
	default:
		return fyne.CurrentApp().Storage().RootURI().String()
	}
}

func getLinuxDownloadDir() string {
	file, err := os.Open(os.ExpandEnv("$HOME/.config/user-dirs.dirs"))
	if err != nil {
		return filepath.Join(os.Getenv("HOME"), "Downloads")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "XDG_DOWNLOAD_DIR") {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				path := strings.Trim(parts[1], "\"")
				path = strings.ReplaceAll(path, "$HOME", os.Getenv("HOME"))
				return path
			}
		}
	}

	return filepath.Join(os.Getenv("HOME"), "Downloads")
}

func isIdSearched(search, id string) bool {
	if search == "" {
		return true
	} else if isChildInBranch(id, search, data.Dir) {
		return true
	}
	return false
}

func isChildInBranch(id, search string, data *pb.Dir) bool {
	for _, file := range data.File {
		if strings.Contains(file.Name, search) {
			return true
		}
	}

	if id == "" {
		for _, dir := range data.Dir {
			return isChildInBranch(dir.FullName, search, dir)
		}
	} else {
		branch := getBranch(data, id)

		if branch != nil {
			for _, dir := range branch.Dir {
				return isChildInBranch(dir.FullName, search, dir)
			}
		}
	}
	return false
}

func createTreeServer() *widget.Tree {
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
			bg := canvas.NewRectangle(color.Transparent)

			if branch {
				icon := widget.NewIcon(theme.FolderIcon())
				label := widget.NewLabel("Dir")
				label.Importance = widget.SuccessImportance
				content := container.NewHBox(icon, label)
				return container.NewStack(bg, content)
			}

			icon := widget.NewIcon(theme.FileIcon())
			label := widget.NewLabel("File")
			label.Importance = widget.SuccessImportance
			content := container.NewHBox(icon, label)
			return container.NewStack(bg, content)
		},
		// UpdateNode
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {

			bg := o.(*fyne.Container).Objects[0].(*canvas.Rectangle)
			content := o.(*fyne.Container).Objects[1].(*fyne.Container)
			label := content.Objects[1].(*widget.Label)

			text := id
			text = filepath.Base(text)
			label.Importance = widget.SuccessImportance
			label.SetText(text)

			if id == selectedNode {
				label.TextStyle = fyne.TextStyle{Bold: true}
				bg.FillColor = color.RGBA{30, 144, 255, 255} // Bleu sélection
			} else {
				label.TextStyle = fyne.TextStyle{}
				bg.FillColor = color.Transparent
			}
			bg.Refresh()
			label.Refresh()
		})
}

func setLoadText(text string) {
	loadFileText.Text = text
}
