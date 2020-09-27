package controllers

import (
	"Beego01/db_mysql"
	"Beego01/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post(){
	fmt.Println(r == nil)
	fmt.Println(r.Ctx == nil)
	fmt.Println(r.Ctx.Request == nil)
	var user models.User

	//err := r.ParseForm(&user)
	//if err != nil{
	//
	//}
	//解析前段提交的用户注册的信息
	//直接调用保存数据的一个函数，并判断保存后的结果
	dataBytes, err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil{
		r.Ctx.WriteString("数据接收请求失败")
		return
	}
	err = json.Unmarshal(dataBytes,&user)
	if err != nil{
		r.Ctx.WriteString("数据解析失败")
		return
		}
		//一切正常，将用户信息保存到数据库中去
		fmt.Println("姓名",user.Name)
	    fmt.Println("生日",user.Birthday)
	    fmt.Println("地址",user.Address)
	    fmt.Println("昵称",user.Nick)
	    fmt.Println("密码",user.Password)
	//r.Ctx.WriteString("数据解析成功")
		//3、将解析到的用户数据
		id,err := db_mysql.InsertUser(user)
		if err != nil{
			fmt.Println(err.Error())
			r.Ctx.WriteString("用户保存失败")
			return
		}
		fmt.Println(id)
		result := models.ResponseResult{
			Code:   0,
			Message: "保存成功",
			Data:    nil,
		}
		r.Data["json"] = &result
		r.ServeJSON()
		//r.Ctx.WriteString("恭喜，用户保存成功")


	//	fmt.Println("姓名：",user.Name)
	//fmt.Println("地址：",user.Address)
	//r.Ctx.WriteString("数据解析成功，并接收。")
}