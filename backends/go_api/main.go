package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	prog_langs struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		RelDate int `json:"rel_date"`
		Auth string `json:"auth"`
		Comp string `json:"comp"`
	}
)

//----------
// Handlers
//----------

func createLang(c echo.Context) error {
	lang := new(prog_langs)
	lang.Name = c.FormValue("name")
	rel_date, _ := strconv.Atoi(c.FormValue("rel_date"))
	lang.RelDate = rel_date
	lang.Auth = c.FormValue("auth")
	lang.Comp = c.FormValue("comp")

	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	db.Create(&lang)

	return c.JSON(http.StatusCreated, lang)
}

func getLang(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	var res prog_langs
	// Encuentra fila con llave primaria ingresada en la variable id 
	db.First(&res, id)
	return c.JSON(http.StatusOK, res)
}

func updateUser(c echo.Context) error {
	u := new(prog_langs)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	//users[id].Name = u.Name
	return c.JSON(http.StatusOK, id)
}

func deleteUser(c echo.Context) error {
	//id, _ := strconv.Atoi(c.Param("id"))
	return c.NoContent(http.StatusNoContent)
}

func getAllLangs(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/prog_langs", getAllLangs)
	e.POST("/prog_langs", createLang)
	e.GET("/prog_langs/:id", getLang)
	e.PUT("/prog_langs/:id", updateUser)
	e.DELETE("/prog_langs/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start("localhost:1323"))
}
