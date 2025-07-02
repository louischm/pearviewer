package pages

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"pearviewer/client/types"
)

func AddHomePage(pg *core.Pages, user *types.User) {
	pg.AddPage("home", func(pg *core.Pages) {
		core.NewText(pg).SetText("Welcome, " + user.Username + "!").SetType(core.TextHeadlineSmall)
		core.NewButton(pg).SetText("Sign out").OnClick(func(e events.Event) {
			*user = types.User{}
			pg.Open("sign-in")
		})
	})
}
