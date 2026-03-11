package configs

type Config struct {
	Server        *Server        `mapstructure:"server"`
	Mysql         *Mysql         `mapstructure:"mysql"`
	CasdoorConfig *CasdoorConfig `mapstructure:"casdoor-config"`
}

type Server struct {
	Host         *string `mapstructure:"host"`
	Port         *int    `mapstructure:"port"`
	DbType       *string `mapstructure:"db-type"`
	RouterPrefix *string `mapstructure:"prefix"`
	Model        *string `mapstructure:"model"`
	OpenSso      *bool   `mapstructure:"open-sso"`
}

func (s *Server) GetOpenSso() bool {
	if s.OpenSso == nil {
		return false
	}
	return *s.OpenSso
}

func (s *Server) GetModel() string {
	if s.Model == nil {
		return "debug"
	}
	return *s.Model
}

func (s *Server) GetDbType() string {
	if s.DbType == nil {
		return "mysql"
	}
	return *s.DbType
}

func (s *Server) GetHost() string {
	if s.Host == nil {
		return "127.0.0.1"
	}
	return *s.Host
}

func (s *Server) GetPort() int {
	if s.Port == nil {
		return 8888
	}
	return *s.Port
}

func (s *Server) GetRouterPrefix() string {
	if s.RouterPrefix == nil {
		return ""
	}
	return *s.RouterPrefix
}
