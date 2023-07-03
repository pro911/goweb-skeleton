// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "用来刷新生成gorm gen的相关数据",
	Long: `加入数据表结构发生了变化,这里需要重新根据数据表结构生成新的gorm gen代码生成:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: genGorm,
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func genGorm(cmd *cobra.Command, args []string) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/pkg/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	//dsn
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_DATABASE"),
	)

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	g.UseDB(db) // reuse your gorm db

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	}

	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(
		g.GenerateModel("user"),
		g.GenerateModel("team"),
		g.GenerateModel("team_place"),
		g.GenerateModel("project"),
		g.GenerateModel("project_user"),
		g.GenerateModel("project_user"),
		g.GenerateModel("task"),
		g.GenerateModel("task_response"),
	)

	g.Execute()
}
