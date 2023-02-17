package business

import (
	"fiber1/model"

	"github.com/jmoiron/sqlx"
	"github.com/kokizzu/gotro/S"
)

type Guest struct {
	Db *sqlx.DB
}

type Guest_LoginIn struct {
	CommonIn
	Email, Password string
}

type Guest_LoginOut struct {
	CommonOut
	User *model.User
}

const Gues_LoginPath = `/guest/login`

func (g *Guest) login(in *Guest_LoginIn) (out Guest_LoginOut) {
	if len(in.Email) < 3 {
		out.SetError(400, `invalid email`)
		return
	}

	user := model.NewUser(g.Db)
	if !user.FindByEmail(in.Email) {
		out.SetError(400,`user not found`)
		return
	}

	out.SetCookie = S.RandomCB63(32)
	// TODO: insert to session table (redis)

	out.User = user.Clean()
	return

}