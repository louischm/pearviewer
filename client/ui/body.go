package ui

import (
	"cogentcore.org/core/core"
	"github.com/louischm/pkg/logger"
	"pearviewer/client/types"
	"pearviewer/client/ui/pages"
)

var log = logger.NewLog()

func CreateBody() {
	b := core.NewBody()

	user := &types.User{}
	pg := core.NewPages(b)
	pages.AddSignInPage(pg, user)
	pages.AddHomePage(pg, user)

	b.RunMainWindow()
}
