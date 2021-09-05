package v1

import (
	"geek/04/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type GetArticleRequest struct {
	Id string `json:"id,omitempty"`
}

// User 结构体定义
type UpdateUserRequest struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type User struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type UpdateUserReply struct {
	User *User `json:"User,omitempty"`
}

type GetUserReply struct {
	User *User `json:"User,omitempty"`
}

type UserServiceHTTPServer interface {
	//GetArticle(context.Context, *GetArticleRequest) (*GetUserReply, error)
	//UpdateArticle(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)


	//GetArticle(*GetArticleRequest) (*GetUserReply, error)
	//UpdateArticle(*UpdateUserRequest) (*UpdateUserReply, error)
}

func (u *UpdateUserRequest) GetName() string {
	if u != nil {
		return u.Name
	}
	return ""
}

func GetUserHttp(c *gin.Context) {
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

func PostUserHttp(c *gin.Context) {
	// 初始化user struct
	u := UpdateUserRequest{}
	// 通过ShouldBind函数，将请求参数绑定到struct对象， 处理json请求代码是一样的。
	// 如果是post请求则根据Content-Type判断，接收的是json数据，还是普通的http请求参数
	if err := c.ShouldBind(&u); err != nil {
		c.String(http.StatusFound, "Query Error")
		return
	}

	//var us UserServiceHTTPServer
	//var ctx context.Context
	//ur,err := us.UpdateArticle(&u)
	var s service.UserService
	ur,err := s.UpdateArticle(c,&u)
	if err != nil {
		c.String(http.StatusFound, "Query Error")
		return
	}
	// 绑定成功， 打印请求参数
	log.Println(u.Name)
	log.Println(u.Email)

	c.JSON(http.StatusOK, gin.H{"params": ur})
}
