package user

import (
	"github.com/jeffotoni/project.go.standard/project-web/standard.libray/crud.user.singleton/internal/zerolog"
	pg "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/internal/psql"
)

//List ..
func List(rid string) (name, lastname, id string) {
	Db := pg.Connect()
	row := Db.QueryRow(`SELECT * FROM users WHERE id=$1`, rid)
	err := row.Scan(&name, &lastname, &id)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"user.go",
			13,
			"api.crud-dao.com.br",
			"Repo List user",
			err.Error())
		return "", "", ""
	}
	if len(name) > 0 {
		return name, lastname, id
	}
	return "", "", ""
}
