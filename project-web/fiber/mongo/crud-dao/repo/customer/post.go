package rcustomer

import (
	"go.mongodb.org/mongo-driver/mongo"

	mcustomer "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/models/customer"
)

func Post(db *mongo.Client, customer mcustomer.Customer) error {
	//ainda sendo feito

	//Those ones that are not here on the insert, are default fields
	sqlexec := `insert 
	into ad_customer(
		imp_ip,
		imp_user_create, 
		imp_user_update,
		imp_nome 
	)values($1,$2,$3,$4)`
	_, err := db.Exec(sqlexec, //TIPOS DE CADA CAMPO
		customer.ImpIp,         //string
		customer.ImpUserCreate, //uuid
		customer.ImpUserUpdate, //uuid
		customer.ImpNome,       //string
	)
	if err != nil {
		return err
	}
	return nil

}
