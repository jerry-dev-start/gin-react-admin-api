package system

import (
	"apis/internal/handler/system/req"
	systemlogic "apis/internal/logic/system_logic"
	"apis/model/common/response"

	"github.com/gin-gonic/gin"
)

type SsoHandler struct {
	ssoLogic *systemlogic.SsoLogic
}

func NewSsoHandler() *SsoHandler {
	return &SsoHandler{
		ssoLogic: systemlogic.NewSsoLogic(),
	}
}

// SsoLogin 处理单点登录请求
// @Tags SSO相关接口
// @Summary 用户通过SSO凭证进行登录验证
// @Accept json
// @Produce json
// @Param data body req.SsoLoginReq true "SSO登录所需参数（如票据、来源应用等）"
// @Success 200 {object} response.Response "登录成功，返回Token或重定向地址"
// @Router /sso/login [post]
func (s *SsoHandler) SsoLogin(gin *gin.Context) {

	var req req.SsoLoginReq
	if err := gin.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数绑定失败", gin)
		return
	}

}
