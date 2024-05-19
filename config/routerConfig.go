package config

import (
	"log"
	"net/http"
	"user-service/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)
func routerConfig() *gin.Engine{
    r := gin.Default()
	http.ListenAndServe(viper.GetString("base_path")+":"+viper.GetString("port"), r)
	err := r.Run(viper.GetString("base_path")+":"+viper.GetString("port"))
	if err!=nil{
        log.Fatal(err)
    }

    router.UserRouter(r)
	
	return r
}



