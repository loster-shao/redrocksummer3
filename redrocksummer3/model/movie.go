package model

import (
	_ "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"redrocksummer3/tool"
)

type Movie struct{
	gorm.Model
	Name  string
	Num   int
	Where string
}

type Orders struct {
	gorm.Model
	Name  string
	Pname string
	Num   int
	Where string
}

func Buy(body string)  {
	order, _, num,  position :=tool.Chuli(body)

	//查询改电影
	var m Movie
	DB.Where("name=?",order).First(&m)

	//设置电影
	movie := &Movie{
		Name:     order,
		Num:      m.Num-num,
		Where:    position,
	}
	DB.Model(&movie).Update("name", order)
}

func Order(body string)  {
	order, person, num, position :=tool.Chuli(body)
	dingdan := &Orders{
		Name:  order,
		Pname: person,
		Num:   num,
		Where: position,
	}
	DB.Create(dingdan)
}
