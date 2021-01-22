package user

import (
	pg "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-singleton/internal/psql"
	"github.com/jeffotoni/project.go.standard/project-web/standard.libray/crud.user.singleton/internal/zerolog"
)

//Delete ..
func Delete(id string) error {
	Db := pg.Connect()
	dell := `DELETE FROM users WHERE id=$1`
	_, err := Db.Exec(dell, id)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"delete.user.go",
			17,
			"localhost",
			"Repo Delete",
			err.Error())
		return err
	}
	return nil
}
