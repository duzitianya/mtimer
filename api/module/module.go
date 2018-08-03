package module

import (
	"log"
	"mtimer/db/mysql"
)

type User struct {
	//其中json标签用来绑定content-type=application/json时的参数名称
	//form标签用来绑定x-www-form-urlencoded时的参数名称
	Username string `json:"uname" form:"un" binding:"required"`
	Password string `json:"pwds" form:"pwd" binding:"required"`
	Age int `json:"age" form:"age"`
}

type Person struct {
	Id int64
	UserName string
	Password string
	FirstName string
	LastName string
}

func (p *Person)GetPerson(username, password string) error{
	db := mysql.GetDB()

	err := db.QueryRow("SELECT * FROM person where username=? AND pwd=?", username, password).Scan(&p.Id, &p.Password, &p.UserName, &p.FirstName, &p.LastName)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}