package controllers

import (
	"Beego01/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)
//匿名字段组合：
type MainController struct {
	beego.Controller // 匿名字段：
}

/**
 *方法的重写：如果一个结构体包含某个方法A，在其匿名字段中也拥有同名方法A。
 *则外部结构体的方法A对匿名字段结构体方法A的重写。
 */


func (c *MainController) Get() {
	//1、获取请求数据
	User := c.Ctx.Input.Query("User")
	Pow := c.Ctx.Input.Query("Pow")
	//2、使用固定数据进行数据校验
	//zhanglong  123456
	if User != "zhanglong" || Pow != "123456"{
		c.Ctx.ResponseWriter.Write([]byte("sorry，数据校验错误。"))
		return
	}else {
		c.Ctx.ResponseWriter.Write([]byte("恭喜，数据校验成功。"))
	}
	//1、获取name、age、sex
	//2、做数据对比
	//3、根据对比结果进行判断处理
	c.Data["Website"] = "https://www.baidu.com"
	c.Data["Email"] = "1621258931@qq.com"
	c.TplName = "index.tpl"
}

//编写一个post方法，用于处理post类型的请求
//func (c *MainController) Post(){
//	//for i:=0;i<10;i++{
//	//	fmt.Printf("第%d次打印\n",i)
//	//}
//	//1、接收（解析）psot请求的参数
//	name := c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	sex := c.Ctx.Request.FormValue("sex")
//	fmt.Println(name,age,sex)
//	//2、进行数据校验
//	if name != "张龙" && age != "18"{
//		c.Ctx.WriteString("数据校验失败。")
//		return
//	}
//	c.Ctx.WriteString("数据校验成功。")
//}

/**
 *该方法用于处理post请求
 */

func (c *MainController) Post(){
	//1、接收前端传递的数据。
	var person models.Person
	dataBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	json.Unmarshal(dataBytes,&person)
	if err != nil{
		c.Ctx.WriteString("数据接收错误，请重试")
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil{
		c.Ctx.WriteString("数据解析失败，请重试")
	}
	fmt.Println("姓名：",person.Name)
	fmt.Println("年龄：",person.Age)
	fmt.Println("性别：",person.Sex)
	c.Ctx.WriteString("数据解析成功。")
}


