package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"pearviewer/client/grpc"
)

func login(w fyne.Window) {
	title := canvas.NewText("Pearviewer's Login", color.NRGBA{255, 255, 255, 255})
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: false, Italic: false}
	title.TextSize = 32

	username := widget.NewEntry()
	username.SetPlaceHolder("Username")
	usernameWrapped := NewEntryWrapper(username, fyne.NewSize(300, 40))

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	passwordWrapped := NewEntryWrapper(password, fyne.NewSize(300, 40))

	submit := widget.NewButton("Login", func() {
		_, err := grpc.SignIn(username.Text, password.Text)
		if err != nil {
			log.Debug("login fail")
		} else {
			log.Info("login success")
			homePage(w, username.Text)
		}

	})
	submit.Importance = widget.HighImportance

	newUser := widget.NewButton("Create Account", func() {
		createUser(w)
	})
	newUser.Importance = widget.LowImportance

	form := container.NewVBox(
		title,
		usernameWrapped,
		passwordWrapped,
		submit,
		newUser,
	)
	centered := container.NewCenter(form)
	w.SetContent(centered)
}
