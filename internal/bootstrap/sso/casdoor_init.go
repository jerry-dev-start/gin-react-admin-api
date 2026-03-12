package sso

import (
	"apis/configs"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func InitCasDoor(c *configs.CasdoorConfig) {
	endpoint := c.GetEndpoint()
	clientID := c.GetClientID()
	clientSecret := c.GetClientSecret()
	organizationName := c.GetOrganizationName()
	applicationName := c.GetApplicationName()
	certificate := ""

	casdoorsdk.InitConfig(endpoint, clientID, clientSecret, certificate, organizationName, applicationName)
}
