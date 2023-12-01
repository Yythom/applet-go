package user

import (
	"github.com/gin-gonic/gin"
	"test/controlles/response"
	"test/services"
)

func GetUserInfo(c *gin.Context) {
	//c.Params("id") // 路径参数
	//ReturnSuccess(c, 0, "user-info", "success")
	response.ReturnError(c, 1, "error")
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

func Register(c *gin.Context) {
	// json
	params := make(map[string]string)
	err := c.BindJSON(&params)
	if err == nil {
		data := services.RegisterUser(params)
		response.ReturnSuccess(
			c,
			0,
			"注册成功",
			data,
		)
		return
	}
	response.ReturnError(c, 1, "error")
}

func Login(c *gin.Context) {
	type LoginParams struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// json
	params := make(map[string]string)
	err := c.BindJSON(&params)
	if err == nil {
		// reg
		response.ReturnSuccess(
			c,
			0,
			"登入成功",
			LoginParams{
				Username: params["username"],
				Password: params["password"],
			},
		)
		return
	}
	response.ReturnError(c, 1, "error")
}

func BindMobile(c *gin.Context) {

	//c.Params("id") // 路径参数

	// X-www-form-urlencoded
	//username := c.PostForm("username")
	//password := c.PostForm("password")

	// json
	params := make(map[string]string)
	err := c.BindJSON(&params)
	if err == nil {
		response.ReturnSuccess(c, 0, "success", "")
		return
	}
	response.ReturnError(c, 1, "error")

}
