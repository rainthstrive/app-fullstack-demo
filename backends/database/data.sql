-- Nota: Mostrará errores en VSCode, pero es código valido en MySQL

-- Creamos la base de datos
CREATE DATABASE `demodatabase`;

-- Empezamos a usarla en MySQL
USE `demodatabase`;

-- Tabla sobre lenguajes de programación
CREATE TABLE IF NOT EXISTS `demodatabase`.`prog_langs` (
  `id` INT NOT NULL AUTO_INCREMENT, -- ID de la tabla 
  `name` VARCHAR(50) NULL, -- Nombre del lenguaje de programación
  `rel_date` INT(11) NULL, -- Fecha de lanzamiento en formato Unix
  `auth` VARCHAR(100) NULL, -- Autor del lenguaje
  `comp` VARCHAR(50) NULL, -- Compañía a la que pertenece
  PRIMARY KEY (`id`));

-- Insertamos dos filas

-- Primera fila
INSERT INTO `demodatabase`.`prog_langs`
(`name`,
`rel_date`,
`auth`,
`comp`)
VALUES (
'Javascript',
818100000,
'Brendan Eich',
'Netscape');

-- Segunda fila
INSERT INTO `demodatabase`.`prog_langs`
(`name`,
`rel_date`,
`auth`,
`comp`)
VALUES (
'Go',
1257876000,
'Robert Griesemer',
'Google');
