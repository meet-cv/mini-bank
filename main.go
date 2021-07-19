package main

import (
	"fmt"
	"log"
	"mini-bank/entities"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	//"github.com/go-pg/pg/v10/orm"
)

func main() {

	var db *pg.DB = db_open()
	defer db_close(db)

	router := gin.Default()
	user := router.Group("/user")
	{
		user.POST("/add", func(c *gin.Context) {
			var new_user entities.User
			c.BindJSON(&new_user)
			//validation goes here
			_, err := db.Model(&new_user).Insert()
			if err != nil {
				fmt.Printf("%v", err)
				c.String(404, "some error occured while adding")
			}
			c.String(200, "user added")
		})

		user.POST("/update", func(c *gin.Context) {
			var update_user entities.User
			c.BindJSON(&update_user)
			//validation goes here
			update_user.Updated_At = time.Now()
			_, err := db.Model(&update_user).ColumnExpr("").Insert()
			if err != nil {
				fmt.Printf("%v", err)
				c.String(404, "some error occured while updating")
			}
			c.String(200, "user updated")
		})

		user.POST("/info", func(c *gin.Context) {
			var result entities.User
			user_id := c.PostForm("user_id")
			// validation goes here
			err := db.Model(&result).Where("? = ?", pg.Ident("user_id"), user_id).Select()
			if err != nil {
				log.Printf("%v", err)
				c.String(404, "some error occured")
			} else {
				c.JSON(200, result)
			}
		})

		user.POST("/remove", func(c *gin.Context) {
			var delete_user entities.User
			user_id := c.PostForm("user_id")
			//validation goes here
			_, err := db.Model(&delete_user).Where("user_id = ?", user_id).Delete()
			if err != nil {
				fmt.Printf("%v", err)
				c.String(200, "some error occured while deleting")
			}
			c.String(200, "user deleted")
		})

		account := user.Group("/account")
		{
			account.POST("/info", func(c *gin.Context) {
				var result entities.Account
				acc_id := c.PostForm("acc_id")
				// validation goes here
				err := db.Model(&result).Where("? = ?", pg.Ident("acc_id"), acc_id).Select()
				if err != nil {
					log.Printf("%v", err)
					c.String(404, "some error occured")
				} else {
					c.JSON(200, result)
				}
			})

			account.POST("/balance", func(c *gin.Context) {
				var result entities.Account
				acc_id := c.PostForm("acc_id")
				// validation goes here
				err := db.Model(&result).Where("? = ?", pg.Ident("acc_id"), acc_id).Select()
				if err != nil {
					log.Printf("%v", err)
					c.String(404, "some error occured")
				} else {
					c.JSON(200, result)
				}
			})
		}

		transaction := user.Group("/transaction")
		{
			transaction.POST("/info", func(c *gin.Context) {
				var result entities.Transaction
				transac_id := c.PostForm("transac_id")
				// validation goes here
				err := db.Model(&result).Where("? = ?", pg.Ident("transac_id"), transac_id).Select()
				if err != nil {
					log.Printf("%v", err)
					c.String(404, "some error occured")
				} else {
					c.JSON(200, result)
				}
			})

			transaction.POST("/execute", func(c *gin.Context) {})
			transaction.POST("/all", func(c *gin.Context) {
				var result []entities.Transaction
				user_id := c.PostForm("user_id")
				// validation goes here
				err := db.Model(&result).Where("? = ?", pg.Ident("user_id"), user_id).Select()
				if err != nil {
					log.Printf("%v", err)
					c.String(404, "some error occured")
				} else {
					c.JSON(200, result)
				}
			})
		}
	}
	router.Run(":8080")
}

func db_open() *pg.DB {
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

func db_close(db *pg.DB) {
	err := db.Close()
	if err != nil {
		log.Println("Some problem occured while closing connection")
		os.Exit(100)
	}
	log.Println("Database connection closed successfully.")
}
