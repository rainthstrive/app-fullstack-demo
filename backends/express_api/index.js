const express = require("express");
const app = express();
const PORT = 8080;
const db = require("./database");

app.use(express.json());

app.get("/prog_langs", async (_, res) => {
  try {
    const results = await db
      .promise()
      .query(`SELECT id, name, rel_date, auth, comp FROM prog_langs`);
    res.status(200).send(results[0]);
  } catch (err) {
    console.log(err);
  }
});

app.get("/prog_langs/:id", async (req, res) => {
  const { id } = req.params;
  if (!id) {
    res.status(418).send({ message: "Datos insuficientes" });
  }
  try {
    const results = await db
      .promise()
      .query(
        `SELECT id, name, rel_date, auth, comp FROM prog_langs WHERE id = ${id}`
      );
    const row = results[0][0];
    res.status(200).send(row);
  } catch (err) {
    console.log(err);
  }
});

app.put("/prog_langs/:id", async (req, res) => {
  const { id } = req.params;
  const { name, rel_date, auth, comp } = req.body;
  if (!id || !name || !rel_date || !auth || !comp) {
    res.status(418).send({ message: "Datos insuficientes" });
  }
  try {
    const results = await db.promise().query(
      `UPDATE prog_langs 
        SET 
            name = '${name}', 
            rel_date = ${rel_date}, 
            auth = '${auth}', 
            comp = '${comp}' 
        WHERE id = ${id}`
    );
    const info = results[0].info;
    res.status(200).send({
      message: info,
    });
  } catch (err) {
    console.log(err);
  }
});

app.delete("/prog_langs/:id", async (req, res) => {
  const { id } = req.params;
  if (!id) {
    res.status(418).send({ message: "Datos insuficientes" });
  }
  try {
    await db.promise().query(
      `DELETE FROM prog_langs 
        WHERE id = ${id}`
    );
    res.status(200).send({
        message: "Recurso eliminado",
    });
  } catch (err) {
    console.log(err);
  }
});

app.post("/prog_langs", async (req, res) => {
  const { name, rel_date, auth, comp } = req.body;
  if (!name || !rel_date || !auth || !comp) {
    res.status(418).send({ message: "Datos insuficientes" });
  } else {
    try {
      await db.promise().query(
        `INSERT INTO prog_langs 
                (name, rel_date, auth, comp) 
                VALUES
                ('${name}', ${rel_date}, '${auth}', '${comp}')`
      );
      res.status(201).send({
        message: "Recurso creado",
      });
    } catch (err) {
      console.log(err);
    }
  }
});

app.listen(PORT, () =>
  console.log(`ExpressJS API viva en http://localhost:${PORT}`)
);
