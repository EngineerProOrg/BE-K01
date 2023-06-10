package configs

type WebConfig struct {
	AuthenticateAndPost struct {
		Hosts []string
	}
	Newsfeed struct {
		Hosts []string
	}
}
