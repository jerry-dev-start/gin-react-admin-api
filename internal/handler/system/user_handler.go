package system

import (
	"apis/internal/handler/system/req"
	systemlogic "apis/internal/logic/system_logic"
	"apis/model/common/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userLogic *systemlogic.UserLogic
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userLogic: systemlogic.NewUserLogic(),
	}
}

// SaveUserInfo 保存用户信息
func (u *UserHandler) SaveUserInfo(r *gin.Context) {
	var userReq req.UserReq
	if err := r.ShouldBindJSON(&userReq); err != nil {
		response.FailWithMessage("参数获取失败", r)
		return
	}
	if err := u.userLogic.SaveUserInfo(&userReq); err != nil {
		response.FailWithMessage("用户保存失败", r)
		return
	}
}
