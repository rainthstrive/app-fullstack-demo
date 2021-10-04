package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
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
	jwtCustomClaims struct {
		Name  string `json:"name"`
		Admin bool   `json:"admin"`
		jwt.StandardClaims
	}
)

//----------
// Handlers
//----------

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("gobhgb76/&Jngnghn"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

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
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	var res prog_langs
	// Encuentra fila con llave primaria ingresada en la variable id 
	db.First(&res, id)

	var params = new(prog_langs)
	params.Name = c.FormValue("name")
	rel_date, _ := strconv.Atoi(c.FormValue("rel_date"))
	params.RelDate = rel_date
	params.Auth = c.FormValue("auth")
	params.Comp = c.FormValue("comp")

	if(len(params.Name) > 0 && params.Name != ""){
		res.Name = params.Name
	}
	if(params.RelDate> 0){
		res.RelDate = params.RelDate
	}
	if(len(params.Auth) > 0 && params.Auth != ""){
		res.Auth = params.Auth
	}
	if(len(params.Comp) > 0 && params.Comp != ""){
		res.Comp = params.Comp
	}
	db.Save(&res)
	return c.JSON(http.StatusOK, res)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	var res prog_langs
	db.Where("id = ?", id).Delete(&res)
	return c.NoContent(http.StatusNoContent)
}

func getAllLangs(c echo.Context) error {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}
	var res []prog_langs
	// Encuentra fila con llave primaria ingresada en la variable id 
	db.Find(&res)
	return c.JSON(http.StatusOK, res)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200", "http://localhost:3100"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	  })) 

	// Routes
	e.GET("/prog_langs", getAllLangs)
	e.GET("/prog_langs/:id", getLang)
	// Login route
	e.POST("/login", login)

	// Restricted group
	r := e.Group("/restricted")

	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("gobhgb76/&Jngnghn"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.POST("/prog_langs", createLang)
	r.PUT("/prog_langs/:id", updateUser)
	r.DELETE("/prog_langs/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start("localhost:1323"))
}
