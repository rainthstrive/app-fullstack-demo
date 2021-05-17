# Fullstack Demo Apps

El objetivo de este repositorio es demostrar la conectividad entre diferentes backends y frontends.

## Backends

 - [x] Node.js (Express Js) - 100%
 - [x] Go (Echo Framework) - 100%
 - [ ] ASP.NETCore - 0%

## Frontends

 - [ ] Angular - 0%
 - [ ] React - 0%

# Instalación de Base de Datos

 1. Instalar **MySQL** y **MySQL Workbench**
 2. Generar una conexión local (recordar user y password)
 3. Ejecutar el archivo dentro de `backends/database/data.sql`

# Instalación de Backends
A continuación se explica cómo instalar las dependencias para correr los servidores de desarrollo de las APIs.

## Go API

 1. Instalar la última versión de Go: https://golang.org/
 2. Instalar la extensión Go en VS Code:
    https://marketplace.visualstudio.com/items?itemName=golang.Go
 3. Ir hacia la carpeta del proyecto: `backends/go_api`
 4. Ejecutar el siguiente comando dentro de la carpeta: `go run .`

## ExpressJs API

 1. Instalar la última versión LTS de Node.js: https://nodejs.org/en/
 2. Ir hacia la carpeta del proyecto: `backends/express_api`
 3. Ejecutar el siguiente comando dentro de la carpeta: `npm install`
 4. Y luego ejecutar el siguiente para correr el server: `node .`
