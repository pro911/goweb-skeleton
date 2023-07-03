package dal

import (
	"fmt"
	"github.com/pro911/golang-demo/internal/pkg/dal/query"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func MustInitMySQL() {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_DATABASE"),
	)

	fmt.Println(dsn)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("fatal error db init: %w", err))
	}

	if viper.GetBool("APP_DEBUG") {
		db = db.Debug()
	}

	fmt.Println("mysql initialized")
}

func DB() *gorm.DB {
	return db
}

func GetQuery() *query.Query {
	return query.Use(db)
}
