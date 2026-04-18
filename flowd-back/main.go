package main

// 导入 gin 框架
import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// GetCurPath 获取二进制文件所在目录
func GetCurPath() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(exe)
}

// FlowData 表示数据库中的flow数据
type FlowData struct {
	ID      int    `json:"id"`
	Dir     string `json:"dir"`
	Success int    `json:"success"`
	Failure int    `json:"failure"`
}

func main() {
	r := gin.Default()

	r.GET("/api/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	r.GET("/api/flow", func(c *gin.Context) {
		// 获取数据库路径
		dbPath := GetCurPath() + "/flow.db"

		// 打开数据库连接
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to open database: " + err.Error(),
			})
			return
		}
		defer db.Close()

		// 查询所有flow数据
		rows, err := db.Query("SELECT id, dir, success, failure FROM flow")
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to query database: " + err.Error(),
			})
			return
		}
		defer rows.Close()

		// 解析查询结果
		var flowDataList []FlowData
		for rows.Next() {
			var data FlowData
			if err := rows.Scan(&data.ID, &data.Dir, &data.Success, &data.Failure); err != nil {
				c.JSON(500, gin.H{
					"error": "Failed to scan row: " + err.Error(),
				})
				return
			}
			flowDataList = append(flowDataList, data)
		}

		// 检查rows.Next()的错误
		if err = rows.Err(); err != nil {
			c.JSON(500, gin.H{
				"error": "Error iterating rows: " + err.Error(),
			})
			return
		}

		// 返回数据
		c.JSON(200, gin.H{
			"data": flowDataList,
		})
	})

	r.Run(":8080")
}