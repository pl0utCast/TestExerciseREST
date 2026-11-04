package main

import (
	_ "TestExerciseREST/docs"
	"database/sql"
	"fmt"

	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

// @title Swaggo Example API
// @version 1.0
// @description this is sample
// @license.name Apache 2.0
func main() {

	connStr := "user=postgres dbname=subscriptions password=1234 sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rest := fiber.New()

	rest.Get("/swagger/*", swaggo.HandlerDefault)

	//Получение всех записей
	rest.Get("/getName", func(c fiber.Ctx) error {
		rows, err := db.Query("select * from products")
		if err != nil {
			panic(err)
		}
		products := []product{}

		for rows.Next() {
			// 2. Теперь сканируем
			var p product

			err = rows.Scan(&p.id, &p.model, &p.company, &p.price)
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
			products = append(products, p)
		}

		result := fmt.Sprintln(products)
		return c.SendString(result)
	})

	//Получение одной записи
	rest.Get("/getName/:id", func(c fiber.Ctx) error {
		rows, err := db.Query("select * from products where id=$1", c.Params("id"))
		if err != nil {
			panic(err)
		}
		var p product
		rows.Next()
		err = rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		result := fmt.Sprintln(p)
		return c.SendString(result)
	})

	//Добавление записи
	rest.Post("/PostNew/:id/:name/:desc/:price", func(c fiber.Ctx) error {
		_, err := db.Query("insert into products(id, model, company, price) values($1, $2, $3, $4)",
			c.Params("id"), c.Params("name"), c.Params("desc"), c.Params("price"))
		if err != nil {
			panic(err)
		}
		return c.SendString("Added Successfully")
	})

	//Удаление записи
	rest.Delete("/deleteById/:id", func(c fiber.Ctx) error {
		_, err = db.Query("Delete from products where id=$1", c.Params("id"))
		if err != nil {
			panic(err)
		}
		return c.SendString("Delete success")
	})

	//Изменение записи
	rest.Put("/PutById/:id/:name/:desc/:price", func(c fiber.Ctx) error {
		_, err = db.Query("Update products set model = $2, company = $3, price = $4 where id = $1",
			c.Params("id"), c.Params("name"), c.Params("desc"), c.Params("price"))
		if err != nil {
			panic(err)
		}
		return c.SendString("Update success")
	})

	rest.Listen(":3000")
}

// // @Param        id   path      string  true  "Account ID"
// @Router /getName [get]
func GetName(ctx fiber.Ctx) error {
	return ctx.SendString("")
}

// @Param id path integer true "product ID"
// @Router /getName/{id} [get]
func GetNameByID(ctx fiber.Ctx) error {
	return ctx.SendString("")
}

// @Param id path integer true "product ID"
// @Param name path string true "name"
// @Param description path string true "description"
// @Param price path integer true "price"
// @Router /PostNew/{id}/{name}/{description}/{price} [post]
func PostNew(ctx fiber.Ctx) error {
	return ctx.SendString("")
}

// @Param id path integer true "product ID"
// @Param name path string true "name"
// @Param description path string true "description"
// @Param price path integer true "price"
// @Router /PutById/{id}/{name}/{description}/{price} [put]
func PutById(ctx fiber.Ctx) error {
	return ctx.SendString("")
}

// @Param id path integer true "product Id"
// @Router /DeleteById/{id} [delete]
func DeleteById(ctx fiber.Ctx) error {
	return ctx.SendString("")
}
