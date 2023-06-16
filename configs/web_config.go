package configs

type WebConfig struct {
	Port                int `yaml:"port"`
	AuthenticateAndPost struct {
		Hosts []string `yaml:"hosts"`
	} `yaml:"authenticate_and_post"`
	Newsfeed struct {
		Hosts []string `yaml:"hosts"`
	} `yaml:"newsfeed"`
}
