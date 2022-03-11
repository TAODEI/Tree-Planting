package main

import (
	"TreePlanting/config"
	"TreePlanting/model"
	"TreePlanting/router"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

// @title TreePlanting API
// @version 1.0.0
// @description 植树节API
// @termsOfService http://swagger.io/terrms/
// @contact.name TAODEI
// @contact.email tao_dei@qq.com
// @host tree-planting.muxixyz.com:30002
// @BasePath /api
// @Schemes http

func main() {
	err := config.Init("./conf/config.yaml", "")
	if err != nil {
		panic(err)
	}

	dbMap := viper.GetStringMapString("db")
	addr := os.Getenv("MYSQL_ADDR")
	if addr == "" {
		addr = dbMap["addr"]
	}
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", dbMap["username"], dbMap["password"], addr, dbMap["name"])

	model.DB, err = gorm.Open("mysql", dbConfig)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router.Router(r)

	port := viper.GetString("port")
	r.Run(port)
	defer model.DB.Close()
}
