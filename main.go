package main

import (
	"TreePlanting/config"
	"TreePlanting/model"
	"TreePlanting/router"

	"github.com/gin-gonic/gin"
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

	model.InitDB()

	r := gin.Default()
	router.Router(r)

	port := viper.GetString("port")
	r.Run(port)
	defer model.DB.Close()
}
