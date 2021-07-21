package utils

import (
	"log"
	"mini-bank/entities"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func DB_open() *pg.DB {
	var db *pg.DB = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123456789",
		Addr:     "localhost:5432",
		Database: "mini_bank",
	})

	if db == nil {
		log.Println("Database connection failed!")
		os.Exit(100)
	}

	log.Println("Database connection successfull")
	return db
}

func DB_close(db *pg.DB) {
	err := db.Close()
	if err != nil {
		log.Println("Some problem occured while closing connection")
		os.Exit(100)
	}
	log.Println("Database connection closed successfully.")
}

func DB_create(db *pg.DB) error {
	models := []interface{}{
		(*entities.User)(nil),
		(*entities.Account)(nil),
		(*entities.Transaction)(nil),
		(*entities.AccountHolder)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func ErrorResponse(c *gin.Context, err error) bool {
	if err != nil {
		log.Printf("%v", err.Error())
		var res = make(map[string]string)
		res["message"] = "some error occured"
		c.JSON(400, res)
		return true
	}
	return false
}
