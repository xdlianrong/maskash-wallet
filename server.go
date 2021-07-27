package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"time"
	"wallet/controllers"
	"wallet/model"
	"wallet/utils"
)

var (
	app       = cli.NewApp()
	baseFlags = []cli.Flag{
		utils.PortFlag,
		utils.ExchangeIPFlag,
		utils.EthIPFlag,
		utils.RegulatorIPFlag,
	}
)

func init() {
	app.Name = "wallet"
	app.Usage = "user buy/exchange from here"
	app.Action = wallet
	app.Flags = append(app.Flags, baseFlags...)
}
func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func wallet(ctx *cli.Context) {
	model.Ethurl = "http://" + ctx.String("ethIP") + ":8545"
	controllers.ExchangeURL = "http://" + ctx.String("exchangeIP") + ":1323/"
	controllers.RegulatorURL = "http://" + ctx.String("regulatorIP") + ":1423/"
	_ = startNetwork(ctx)
}
func startNetwork(ctx *cli.Context) error {
	e := echo.New()
	// 跨域请求配置
	port := ctx.String("port")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderOrigin, echo.HeaderContentType},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST},
		AllowCredentials: true,    // 允许 cookie
		MaxAge:           43200})) // 预检结果能保留 12h
	// 一组路由
	g := e.Group("/wallet")
	{
		g.POST("/register", controllers.Register)     //注册
		g.POST("/buycoin", controllers.Buycoin)       //购币
		g.POST("/exchange", controllers.ExchangeCoin) //转账
		g.POST("/receive", controllers.Receive)       //收款
	}
	// 网页的静态文件
	// 启动服务，平滑关闭
	go func() {
		if err := e.Start(":" + port); err != nil {
			e.Logger.Fatal("Fail to start with error:%v", err)
		}
	}()
	fmt.Println("服务启动成功")
	// 监听停止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	// 留 5s 处理已经接受的请求，然后关闭服务
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx2); err != nil {
		e.Logger.Fatal("Fail to shutdown with error", err)
		return err
	}
	return nil
}
