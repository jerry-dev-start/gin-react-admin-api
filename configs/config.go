package configs

type Config struct {
	Server *Server `mapstructure:"server"`
	Mysql  *Mysql  `mapstructure:"mysql"`
}

type Server struct {
	Host         *string `mapstructure:"host"`
	Port         *int    `mapstructure:"port"`
	DbType       *string `mapstructure:"db-type"`
	RouterPrefix *string `mapstructure:"prefix"`
}

func (s *Server) GetDbType() string {
	if s.DbType == nil {
		return "mysql"
	}
	return *s.DbType
}
