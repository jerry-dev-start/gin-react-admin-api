package system

import (
	"apis/internal/handler/system/req"
	"apis/internal/handler/system/res"
	systemlogic "apis/internal/logic/system_logic"
	"apis/model/common/response"
	"apis/pkg"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userLogin *systemlogic.UserLogic
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userLogin: systemlogic.NewUserLogic(),
	}
}

func (a *AuthHandler) Login(gin *gin.Context) {
	var loginReq req.LoginReq
	if err := gin.ShouldBindJSON(&loginReq); err != nil {
		response.FailWithMessage("参数邦定失败", gin)
		return
	}
	userInfo, err := a.userLogin.GetUserInfoByUsername(loginReq.Username)
	if err != nil {
		response.FailWithMessage(err.Error(), gin)
		return
	}
	if userInfo == nil {
		response.FailWithMessage("当前用户不存在", gin)
		return
	}
	if hash := pkg.CheckPasswordHash(loginReq.Password, *userInfo.PasswordHash); !hash {
		response.FailWithMessage("账号或密码错误", gin)
		return
	}
	token, err := pkg.GenerateToken(userInfo.ID, *userInfo.Username)
	if err != nil {
		response.FailWithMessage("系统繁忙，生成令牌失败", gin)
	}
	response.OkWithData(res.LoginRes{Token: token}, gin)
}
