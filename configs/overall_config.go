package configs

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
)

type allInOneConfig struct {
	MySQL                     mysql.Config
	Redis                     redis.Options
	AuthenticateAndPostConfig *AuthenticateAndPostConfig
	NewsfeedConfig            *NewsfeedConfig
	WebConfig                 *WebConfig
}

func getAllInOneConfig(path string) (*allInOneConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open AuthenticateAndPost config (path=%s) error: %s", path, err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read AuthenticateAndPost config (path=%s) error: %s", path, err)
	}

	allInOneConf := &allInOneConfig{}
	if err := yaml.Unmarshal(bs, allInOneConf); err != nil {
		return nil, fmt.Errorf("unmarshal AuthenticateAndPost config (path=%s) error: %s", path, err)
	}
	return allInOneConf, nil
}

func GetAuthenticateAndPostConfig(path string) (*AuthenticateAndPostConfig, error) {
	allInOneConf, err := getAllInOneConfig(path)
	if err != nil {
		return nil, err
	}
	return allInOneConf.AuthenticateAndPostConfig, nil
}

func GetNewsfeedConfig(path string) (*NewsfeedConfig, error) {
	allInOneConf, err := getAllInOneConfig(path)
	if err != nil {
		return nil, err
	}
	return allInOneConf.NewsfeedConfig, nil
}

func GetWebConfig(path string) (*WebConfig, error) {
	allInOneConf, err := getAllInOneConfig(path)
	if err != nil {
		return nil, err
	}
	return allInOneConf.WebConfig, nil
}
