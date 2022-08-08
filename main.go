package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/astaxie/beego/session/redis"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"remind-relay/controllers"
	"strings"
)

func main() {
	//加载本地配置
	err := loadLocalConfig()
	if err != nil {
		return
	}

	//跨域支持
	beego.InsertFilter("/*", beego.BeforeRouter, corsFunc)

	//添加路由
	AddRouter()

	//加载session配置
	loadSessionConfig(err)

	//启动
	beego.Run()
}

//AddRouter 注册路由
func AddRouter() {
	beego.Router("/out/captcha", &controllers.CaptchaController{})
}

//加载session配置
func loadSessionConfig(err error) {
	sessionProviderConfig := "127.0.0.1:6379,1000"
	if err != nil {
		fmt.Print(fmt.Sprintf("启动失败：%s", err.Error()), "\n")
		return
	}
	fmt.Println("redis配置：", sessionProviderConfig)
	fmt.Println("SessionOn:", beego.BConfig.WebConfig.Session.SessionOn)
	fmt.Println("SessionProvider:", beego.BConfig.WebConfig.Session.SessionProvider)
	beego.BConfig.WebConfig.Session.SessionProviderConfig = sessionProviderConfig
}

//加载本地配置
func loadLocalConfig() error {
	fmt.Print("启动参数：", os.Args[0], "\n")
	executeFileName, err := filepath.Abs(os.Args[0])
	if err != nil {
		fmt.Println(fmt.Sprintf("启动失败：%s", err.Error()))
		return err
	}
	rootDir := filepath.Dir(executeFileName)
	dirs, err := ioutil.ReadDir(rootDir)
	if err != nil {
		fmt.Println(fmt.Sprintf("启动失败：%s", err.Error()))
		return err
	}
	hasConfigs := false
	for _, dir := range dirs {
		if dir.IsDir() && strings.Contains(dir.Name(), "configs") {
			hasConfigs = true
		}
	}
	if !hasConfigs {
		rootDir = filepath.Dir(rootDir)
	}
	fmt.Println("rootDir：", rootDir)
	err = beego.LoadAppConfig("ini", fmt.Sprintf("%s/configs/com-config.conf", rootDir))
	if err != nil {
		fmt.Println(fmt.Sprintf("读取配置文件失败：%s", err.Error()))
		return err
	}
	return nil
}

//跨域options
var success = []byte("SUPPORT OPTIONS")

//跨域回调
var corsFunc = func(ctx *context.Context) {
	origin := ctx.Input.Header("Origin")
	ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,DELETE,POST,GET,PUT,PATCH")
	ctx.Output.Header("Access-Control-Max-Age", "3600")
	ctx.Output.Header("Access-Control-Allow-Headers", "X-Custom-Header,accept,Content-Type,Access-Token")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	ctx.Output.Header("Access-Control-Allow-Origin", origin)
	if ctx.Input.Method() == http.MethodOptions {
		// options请求，返回200
		ctx.Output.SetStatus(http.StatusOK)
		_ = ctx.Output.Body(success)
	}
}
