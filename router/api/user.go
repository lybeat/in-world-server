package api

import (
	"in-world-server/model"
	"in-world-server/pkg/e"
	"in-world-server/pkg/net"
	"in-world-server/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type RegisterUserForm struct {
	Username string `form:"username" valid:"Required;MaxSize(40)"`
	Password string `form:"password" valid:"Required;MaxSize(20)"`
}

func Register(c *gin.Context) {
	var (
		netG = net.Gin{c}
		form RegisterUserForm
	)

	httpCode, errCode := net.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		netG.Response(httpCode, errCode, nil)
		return
	}

	user := model.User{
		Username: form.Username,
		Password: form.Password,
	}
	if err := user.Register(); err != nil {
		netG.Response(http.StatusInternalServerError, e.ERROR_UESR_REGISTER_FAIL, nil)
		return
	}

	netG.Response(http.StatusOK, e.SUCCESS, nil)
}

func Login(c *gin.Context) {
	netG := net.Gin{c}

	username := c.PostForm("username")
	password := c.PostForm("password")
	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空")
	valid.Required(password, "password").Message("密码不能为空")

	if valid.HasErrors() {
		net.MarkErrors(valid.Errors)
		netG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	user, err := model.Login(username, password)
	if err != nil {
		netG.Response(http.StatusOK, e.ERROR_USER_GET_FAIL, nil)
	}

	netG.Response(http.StatusOK, e.SUCCESS, user)
}

func GetUsers(c *gin.Context) {
	netG := net.Gin{c}
	valid := validation.Validation{}

	if valid.HasErrors() {
		net.MarkErrors(valid.Errors)
		netG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	total, err := model.GetUserTotal()
	if err != nil {
		netG.Response(http.StatusOK, e.ERROR_USER_GET_FAIL, nil)
		return
	}

	users, err := model.GetUsers(util.GetPage(c))
	if err != nil {
		netG.Response(http.StatusInternalServerError, e.ERROR_USERS_GET_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = users
	data["total"] = total

	netG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetUser(c *gin.Context) {
	netG := net.Gin{c}

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		net.MarkErrors(valid.Errors)
		netG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	exist, err := model.ExistUserByID(id)
	if err != nil {
		netG.Response(http.StatusOK, e.ERROR_USER_GET_FAIL, nil)
		return
	}
	if !exist {
		netG.Response(http.StatusOK, e.ERROR_USER_NOT_EXIST, nil)
		return
	}

	user, err := model.GetUser(id)
	if err != nil {
		netG.Response(http.StatusOK, e.ERROR_USER_GET_FAIL, nil)
	}

	netG.Response(http.StatusOK, e.SUCCESS, user)
}
