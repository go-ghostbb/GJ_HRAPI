package main

import (
	"fmt"
	gbstr "ghostbb.io/gb/text/gb_str"
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gorm"
	"hrapi/internal/types"
	"hrapi/internal/types/method"
)

func main() {
	const sqlServerDSN = "sqlserver://sa:Ab@589095@localhost:1433?database=hrms&encrypt=disable"
	var connectDB = func(dsn string) *gorm.DB {
		db, err := gorm.Open(sqlserver.Open(dsn))
		if err != nil {
			panic(fmt.Errorf("connect db fail: %w", err))
		}
		return db
	}

	// 指定生成程式碼的具體相對目錄（相對於當前檔案），默認為：./query
	// 默認生成需要使用WithContext之後才可以查詢的程式碼，但可以通過設置gen.WithoutContext禁用該模式
	config := gen.Config{
		// 默認會在 OutPath 目錄生成 CRUD 程式碼，並且同目錄下生成 model 套件
		// 所以 OutPath 最終 package 不能設置為 model，在有數據庫表同步的情況下會產生衝突
		// 若一定要使用可以通過 ModelPkgPath 單獨指定 model package 的名稱
		OutPath: "./internal/query",
		/*ModelPkgPath: "cmd/gen/model",*/

		// gen.WithoutContext：禁用 WithContext 模式
		// gen.WithDefaultQuery：生成一個全局 Query 對象 Q
		// gen.WithQueryInterface：生成 Query 接口
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,

		FieldWithIndexTag: true, // generate with gorm index tag
		FieldWithTypeTag:  true, // generate with gorm column type tag
	}

	// 解決特殊字串問題。
	// 首字母必須大寫，"-"全部轉為"_"，將"_"移除並將後面字母轉大寫
	config.WithModelNameStrategy(func(tableName string) (modelName string) {
		return gbstr.CaseCamel(tableName)
		//tableName = strings.ReplaceAll(tableName, "-", "_")
		//tableName = strings.ToUpper(tableName[:1]) + tableName[1:]
		//for strings.Contains(tableName, "_") {
		//	index := strings.Index(tableName, "_")
		//	tableName = tableName[:index] + strings.ToUpper(tableName[index+1:index+2]) + tableName[index+2:]
		//}
		//return tableName
	})

	g := gen.NewGenerator(config)

	// 通常複用項目中已有的 SQL 連接配置 db(*gorm.DB)
	// 非必需，但如果需要複用連接時的 gorm.Config 或需要連接數據庫同步表信息則必須設置
	g.UseDB(connectDB(sqlServerDSN))

	// 從連接的數據庫為所有表生成 Model 結構體和 CRUD 程式碼
	// 也可以手動指定需要生成程式碼的數據表
	g.ApplyBasic(types.Basic()...)
	g.ApplyBasic(types.M2M()...)
	g.ApplyInterface(func(method.Menu) {}, types.Menu{})
	g.ApplyInterface(func(method.Employee) {}, types.Employee{})
	g.ApplyInterface(func(method.WorkSchedule) {}, types.WorkSchedule{})
	g.ApplyInterface(func(method.CheckInStatus) {}, types.CheckInStatus{})
	g.ApplyInterface(func(method.SalaryAddItem) {}, types.SalaryAddItem{})
	g.ApplyInterface(func(method.SalaryReduceItem) {}, types.SalaryReduceItem{})
	g.ApplyInterface(func(method.LeaveRequestForm) {}, types.LeaveRequestForm{})
	g.ApplyInterface(func(method.Department) {}, types.Department{})
	g.ApplyInterface(func(method.LeaveCorrect) {}, types.LeaveCorrect{})

	// 執行並生成程式碼
	g.Execute()
}
