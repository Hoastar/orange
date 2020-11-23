/*
@Time : 2020/10/28 下午10:06
@Author : hoastar
@File : migrate
@Software: GoLand
*/

package migrate

import (
	"fmt"
	"github.com/hoastar/orange/database"
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/models/gorm"
	"github.com/hoastar/orange/models/system"

	"github.com/hoastar/orange/pkg/logger"
	config2 "github.com/hoastar/orange/tools/config"

	"github.com/spf13/cobra"
)

var (
	config string
	mode string
	StartCmd = &cobra.Command{
		Use: "init",
		Short: "initialize the database",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func run() {
	usage := `start init`
	fmt.Println(usage)

	//1. 读取配置文件
	config2.ConfigSetup(config)
	//2. 初始化数据库链接
	database.Setup()
	//3. 数据导入
	_ = migrateModel()
	logger.Info("数据库结构初始话成功！")
	//4. 数据初始话完成
	if err := system.InitDb(); err != nil {
		logger.Fatalf("数据库基础数据初始化失败，%v", err)
	}
}

func migrateModel() error {
	if config2.DatabaseConfig.Dbtype == "mysql" {
		orm.Eloquent = orm.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}

	return gorm.AutoMigrate(orm.Eloquent)
}




































