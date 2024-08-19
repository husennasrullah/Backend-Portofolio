package config

type DevelopmentConfig struct {
	Configuration
	Server struct {
		Protocol string `json:"protocol"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
	} `json:"server"`
	Postgresql struct {
		Address       string `json:"address"`
		DefaultSchema string `json:"default_schema"`
	} `json:"postgresql"`
	Redis struct {
		Host                   string `json:"host"`
		Port                   int    `json:"port"`
		Db                     int    `json:"db"`
		Password               string `json:"password"`
		Timeout                int    `json:"timeout"`
		RequestVolumeThreshold int    `json:"request_volume_threshold"`
		SleepWindow            int    `json:"sleep_window"`
		ErrorPercentThreshold  int    `json:"error_percent_threshold"`
		MaxConcurrentRequests  int    `json:"max_concurrent_requests"`
	} `json:"redis"`
}

func (input DevelopmentConfig) GetServerProtocol() string {
	return input.Server.Protocol
}
func (input DevelopmentConfig) GetServerHost() string {
	return input.Server.Host
}
func (input DevelopmentConfig) GetServerPort() int {
	return input.Server.Port
}
func (input DevelopmentConfig) GetPostgreSQLAddress() string {
	return input.Postgresql.Address
}
func (input DevelopmentConfig) GetPostgreSQLDefaultSchema() string {
	return input.Postgresql.DefaultSchema
}
func (input DevelopmentConfig) GetRedisHost() string {
	return input.Redis.Host
}
func (input DevelopmentConfig) GetRedisPort() int {
	return input.Redis.Port
}
func (input DevelopmentConfig) GetRedisDB() int {
	return input.Redis.Db
}
func (input DevelopmentConfig) GetRedisPassword() string {
	return input.Redis.Password
}
func (input DevelopmentConfig) GetRedisTimeout() int {
	return input.Redis.Timeout
}
func (input DevelopmentConfig) GetRedisRequestVolumeThreshold() int {
	return input.Redis.RequestVolumeThreshold
}
func (input DevelopmentConfig) GetRedisSleepWindow() int {
	return input.Redis.SleepWindow
}
func (input DevelopmentConfig) GetRedisErrorPercentThreshold() int {
	return input.Redis.ErrorPercentThreshold
}
func (input DevelopmentConfig) GetRedisMaxConcurrentRequests() int {
	return input.Redis.MaxConcurrentRequests
}
