package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// User 结构体定义
type UpdateUserRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (u *UpdateUserRequest) GetName() string {
	if u != nil {
		return u.Name
	}
	return ""
}

func GetUser(c *gin.Context) {
	id_s := c.Param("id")

	id, err := strconv.ParseInt(id_s, 10, 64)
	if id <= 0 || err != nil {
		// 参数不存在
		c.String(http.StatusFound, "Query Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"params": id})
	return
}

func PostUser(c *gin.Context) {
	// 初始化user struct
	u := UpdateUserRequest{}
	// 通过ShouldBind函数，将请求参数绑定到struct对象， 处理json请求代码是一样的。
	// 如果是post请求则根据Content-Type判断，接收的是json数据，还是普通的http请求参数
	if err := c.ShouldBind(&u); err != nil {
		c.String(http.StatusFound, "Query Error")
		return
	}

	// 绑定成功， 打印请求参数
	log.Println(u.Name)
	log.Println(u.Email)

	c.JSON(http.StatusOK, gin.H{"params": u})
}
