package configs

type StudentManagerConfig struct {
	DB struct {
		Addr string `json:"addr"`
	} `json:"db"`
	Redis struct {
		Addr string `json:"addr"`
	} `json:"redis"`
}
