package system

import (
	"apis/internal/handler/system/req"
	"apis/model/common/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// SaveUserInfo 保存用户信息
func (u *UserHandler) SaveUserInfo(r *gin.Context) {
	var userReq req.UserReq
	if err := r.ShouldBindJSON(&userReq); err != nil {
		response.FailWithMessage("参数获取失败", r)
		return
	}

}
