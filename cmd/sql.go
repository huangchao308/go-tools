package cmd

import (
	"log"

	"github.com/huangchao308/go-tools/internal/sql2struct"
	"github.com/spf13/cobra"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql 转换和处理",
	Long:  "sql 转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql 转换成 struct",
	Long:  "sql 转换成 struct",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("connect to mysql failed: %s", err.Error())
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("get columns failed: %s", err.Error())
		}
		template := sql2struct.NewStructTemplate()
		structColumns := template.AssemblyColumns(columns)
		err = sql2struct.NewStructTemplate().Generate(tableName, structColumns)
		if err != nil {
			log.Fatalf("generate struct failed: %s", err.Error())
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "u", "root", "username")
	sql2structCmd.Flags().StringVarP(&password, "password", "p", "123456", "password")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "c", "utf8mb4", "charset")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "d", "mysql", "dbType")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "n", "", "dbName")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "t", "", "tableName")
}
