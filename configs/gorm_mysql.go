package configs

type Mysql struct {
	Username *string `mapstructure:"username"`
	Password *string `mapstructure:"password"`
	Host     *string `mapstructure:"host"`
	Port     *string `mapstructure:"port"`
	DbName   *string `mapstructure:"db-name"`
	Engine   *string `mapstructure:"engine"`
}
