package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"
)

// database instance

var db *sql.DB

//database setting

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "Briancale20122003"
	dbname   = "mercaditogamer"
)

// products struct
type Prouduct struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Section     string  `json:"section"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

type Products struct {
	Products []Prouduct `json:"products"`
}

// conect function

func Connect() error {
	var err error

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))

	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func main() {

	//connect database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/products", func(c *fiber.Ctx) error {
		//get data products
		rows, err := db.Query("SELECT * FROM products")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer rows.Close()
		result := Products{}

		for rows.Next() {
			product := Prouduct{}
			if err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Section, &product.Image, &product.Price); err != nil {
				return err
			}

			result.Products = append(result.Products, product)
		}

		return c.JSON(result)
	})

	app.Listen(":3000")

	fmt.Println("server on port: 300")

}
