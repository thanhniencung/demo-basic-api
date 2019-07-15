package main

import (
	"basic-api/db"
	"basic-api/router"
	"github.com/labstack/echo"
)

func main()  {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "ryan",
		Password: "postgres",
		DbName:   "chapi",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	router.Router(e, sql)
	e.Logger.Fatal(e.Start(":3000"))
}
