package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
type product struct {
	id      int
	model   string
	company string
	price   int
}

func main() {

	connStr := "user=postgres dbname=subscriptionsDB password=5597 sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rest := fiber.New()
	//Получение всех записей
	rest.Get("/", func(c fiber.Ctx) error { // ← добавили error
		rows, err := db.Query("select * from Products")
		if err != nil {
			panic(err)
		}
		products := []product{}

		//Добавление из базы данных в
		products = append(products, product{0, "ddf", "dfdf", 100})
		if rows.Next() {
			// 2. Теперь сканируем
			var p product

			err = rows.Scan(&p.id, &p.model, &p.company, &p.price)
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		}

		result := fmt.Sprintln(products[0].model)
		return c.SendString(result)
	})

	//Отправка одной записи
	rest.Post("/", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	rest.Listen(":3000")
	//Изменение записи

	//Удаление записи

	//Подсчет суммы

	/*db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.id, p.model, p.company, p.price)
	}*/

}
