package configs

type Mysql struct {
	Username     *string `mapstructure:"username"`
	Password     *string `mapstructure:"password"`
	Host         *string `mapstructure:"host"`
	Port         *string `mapstructure:"port"`
	DbName       *string `mapstructure:"db-name"`
	Engine       *string `mapstructure:"engine"`
	MaxIdleConns *int    `mapstructure:"max-idle-conns"`
	MaxOpenConns *int    `mapstructure:"max-open-conns"`
}

func (m *Mysql) GetUsername() string {
	if m.Username == nil {
		return ""
	}
	return *m.Username
}

func (m *Mysql) GetPassword() string {
	if m.Password == nil {
		return ""
	}
	return *m.Password
}

func (m *Mysql) GetHost() string {
	if m.Host == nil {
		return ""
	}
	return *m.Host
}

func (m *Mysql) GetPort() string {
	if m.Port == nil {
		return ""
	}
	return *m.Port
}

func (m *Mysql) GetDbName() string {
	if m.DbName == nil {
		return ""
	}
	return *m.DbName
}

func (m *Mysql) GetEngine() string {
	if m.Engine == nil {
		return "innoDB"
	}
	return *m.Engine
}

func (m *Mysql) GetMaxIdleConns() int {
	if m.MaxIdleConns == nil {
		return 10
	}
	return *m.MaxIdleConns
}

func (m *Mysql) GetMaxOpenConns() int {
	if m.MaxOpenConns == nil {
		return 100
	}
	return *m.MaxOpenConns
}
