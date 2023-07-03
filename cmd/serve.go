// Package cmd /*
package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pro911/golang-demo/internal/app/router"
	"github.com/pro911/golang-demo/internal/pkg/crontab"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "这个是启动一个gin 框架的http服务",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serve(cmd *cobra.Command, args []string) {

	//启动定时任务
	go func() {
		crontab.NewCronTab()
	}()

	//创建路由引擎
	r := gin.New()

	//加载路由
	router.RegisterRouter(r)

	// 创建一个 HTTP 服务器实例
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", viper.GetInt("TASK_GO_PORT")),
		Handler: r,
	}

	// 启动 HTTP 服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v\n", err)
		}
	}()

	// 监听退出信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// 等待退出信号
	<-sigCh

	// 创建一个上下文，设置超时时间为 5 秒钟
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭 HTTP 服务器
	log.Println("Stopping HTTP server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
