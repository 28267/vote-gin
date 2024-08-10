package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

type Search struct {
	Cid  int    `json:"cid"`
	Name string `json:"name"`
}

func (o OrderController) GetList(c *gin.Context) {
	//1. POST方式获取信息
	//cid := c.PostForm("cid")
	//name := c.DefaultPostForm("name", "lisi")
	//2. JSON方式
	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)
	// if err == nil {
	// 	ReturnSuccess(c, 0, "success", param, 1)
	// 	return
	// } else {
	// 	ReturnError(c, 4401, gin.H{"err": err})
	// }
	//3. 结构体方式
	search := &Search{}
	err := c.BindJSON(search)
	if err == nil {
		ReturnSuccess(c, 0, "success", search, 1)
		return
	} else {
		ReturnError(c, 4401, gin.H{"err": err})
	}

}
