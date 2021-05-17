package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error){
	// La parte root:root representa el username:password, reemplazar por los que usen.
	// La IP 127.0.0.1:3306 representa localhost en el puerto 3306, el default en MySQL
	// Después de la diagonal va el nombre de la base de datos.
	// Para manejar las variables de time.Time en Go correctamente, se tiene que usar parseTime.
	dsn := "root:root@tcp(127.0.0.1:3306)/demodatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}