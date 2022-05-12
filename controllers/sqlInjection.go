package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type SqlInjectionVuln1Controller struct {
	beego.Controller
}

type SqlInjectionVuln2Controller struct {
	beego.Controller
}

type SqlInjectionSafe1Controller struct {
	beego.Controller
}
type SqlInjectionSafe2Controller struct {
	beego.Controller
}

func (c *SqlInjectionVuln1Controller) Get() {
	id := c.GetString("id")
	source := "root:password@tcp(127.0.0.1:3306)/goseccode"
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionVuln2Controller) Get() {
	username := c.GetString("username")
	source := "root:password@tcp(127.0.0.1:3306)/goseccode"
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStr := fmt.Sprintf("select * from user where username=\"%s\"", username)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionSafe1Controller) Get() {
	id, err := c.GetInt("id", 1)
	if err != nil {
		panic(err)
	}
	source := "root:password@tcp(127.0.0.1:3306)/goseccode"
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStr := "select * from user where id=?"
	user := models.User{}
	err = db.QueryRow(sqlStr, id).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionSafe2Controller) Get() {
	username := c.GetString("username")
	source := "root:password@tcp(127.0.0.1:3306)/goseccode"
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStr := "select * from user where username=?"
	user := models.User{}
	err = db.QueryRow(sqlStr, username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}
