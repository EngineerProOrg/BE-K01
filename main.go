package main

import (
	"errors"
	"github.com/EngineerProOrg/BE-K01/internal/sqlclient"
	"github.com/EngineerProOrg/BE-K01/pkg/repository"
	"github.com/caarlos0/env"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"time"
)

type Config struct {
	Dir string `env:"CONFIG_DIR" envDefault:"configs/config.json"`
	DB  bool
}

var config Config
var sqlClient sqlclient.ISqlClientConn

func init() {
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = loc

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}
	viper.SetConfigFile(config.Dir)
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
		panic(err)
	}

	cfg := Config{
		Dir: config.Dir,
		DB:  viper.GetBool(`main.db`),
	}

	if cfg.DB {
		sqlClientConfig := sqlclient.SqlConfig{
			Driver:       "mysql",
			Host:         viper.GetString(`db.host`),
			Database:     viper.GetString(`db.database`),
			Username:     viper.GetString(`db.username`),
			Password:     viper.GetString(`db.password`),
			Port:         viper.GetInt(`db.port`),
			DialTimeout:  20,
			ReadTimeout:  30,
			WriteTimeout: 30,
			Timeout:      30,
			PoolSize:     10,
			MaxIdleConns: 10,
			MaxOpenConns: 10,
		}
		sqlClient = sqlclient.NewSqlClient(sqlClientConfig)
		if sqlClient == nil {
			panic(errors.New("sqlClient is nil"))
		} else {
			CreateOrUpdateGradeView(sqlClient.GetDB())
		}
	}

}

func CreateOrUpdateGradeView(db *gorm.DB) error {
	query := `
		CREATE OR REPLACE VIEW Grade AS
		SELECT e.class_id, e.stud_id,
		       CASE e.grade
		           WHEN 'A' THEN 10
		           WHEN 'B' THEN 8
		           WHEN 'C' THEN 6
		           WHEN 'D' THEN 4
		           WHEN 'E' THEN 2
		           WHEN 'F' THEN 0
		           ELSE 0
		       END
		    AS Grade
		FROM enrolls e
	`
	err := db.Exec(query).Error
	return err
}

func main() {
	stuRepo := repository.NewStudentRepository(sqlClient.GetDB())
	courses, err := stuRepo.GetAvgGradeOfCourse()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(courses)
}
