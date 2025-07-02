package pages

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"github.com/louischm/pkg/logger"
	"pearviewer/client/grpc"
	"pearviewer/client/types"
)

var log = logger.NewLog()

func AddSignInPage(pg *core.Pages, user *types.User) {
	pg.AddPage("sign-in", func(pg *core.Pages) {
		core.NewForm(pg).SetStruct(user)
		core.NewButton(pg).SetText("Sign in").OnClick(func(e events.Event) {
			res, err := grpc.SignIn(user.Username, user.Password)
			if err != nil {
				log.Debug("Failed to sign in: %v", err)
				core.MessageSnackbar(pg, "Failed to sign in: "+err.Error())
				return
			}
			if res.GetReturnCode() == types.Fail {
				log.Debug(res.Message)
				core.MessageSnackbar(pg, res.Message)
				return
			}
			log.Info("Signed in")
			pg.Open("home")
		})
	})
}
