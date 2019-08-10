package main

import (
	"com.phh/generator/config"
	"com.phh/generator/dao"
	"com.phh/generator/service"
	"com.phh/generator/web/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"github.com/rs/zerolog/log"
)

func main() {
	//加载配置
	cfg := config.Cfg()
	//创建iris服务
	app := iris.New()
	app.Logger().SetLevel("debug")
	// 加载模版文件
	tmpl := iris.HTML("./web/templates", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)
	//静态资源目录
	app.StaticWeb(cfg.StaticPath, "./web/res")
	app.Use(recover.New())
	app.Use(func(ctx context.Context) {
		log.Info().Str("url", ctx.RequestPath(false)).Msg("")
		ctx.ViewData("ctxPath", cfg.ContextPath)
		ctx.Next()
	})
	//错误处理
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})

	dao := dao.NewGeneratorDao()
	service := service.NewGeneratorService(dao)

	//root context-path
	//root := app.Party(cfg.ContextPath)
	root := mvc.New(app.Party(cfg.ContextPath))
	{
		//可以为controller 注入需要的对象
		root.Register(service)
		//controller 处理器注册
		root.Handle(new(controller.IndexController))
	}
	//启动应用
	app.Run(
		// 启动端口
		iris.Addr(cfg.ServerAddress),
		//iris.WithoutVersionChecker,
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// 优化
		iris.WithOptimizations,
	)

}
