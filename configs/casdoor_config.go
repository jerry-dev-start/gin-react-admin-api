package configs

type CasdoorConfig struct {
	Endpoint         *string `mapstructure:"endpoint"`
	ClientID         *string `mapstructure:"client-id"`
	ClientSecret     *string `mapstructure:"client-secret"`
	OrganizationName *string `mapstructure:"organization-name"`
	ApplicationName  *string `mapstructure:"application-name"`
}

func (c *CasdoorConfig) GetClientSecret() string {
	if c.ClientSecret != nil {
		return *c.ClientSecret
	}
	return ""
}

func (c *CasdoorConfig) GetOrganizationName() string {
	if c.OrganizationName != nil {
		return *c.OrganizationName
	}
	return ""
}

func (c *CasdoorConfig) GetApplicationName() string {
	if c.ApplicationName != nil {
		return *c.ApplicationName
	}
	return ""
}

func (c *CasdoorConfig) GetEndpoint() string {
	if c.Endpoint != nil {
		return *c.Endpoint
	}
	return ""
}

func (c *CasdoorConfig) GetClientID() string {
	if c.ClientID != nil {
		return *c.ClientID
	}
	return ""
}
