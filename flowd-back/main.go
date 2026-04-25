package main

// 导入 gin 框架
import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
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

// initDB 初始化数据库，自动创建表（如果不存在）
func initDB() error {
	dbPath := GetCurPath() + "/flow.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	// 测试连接
	if err = db.Ping(); err != nil {
		return err
	}

	// 创建 flow 表（如果不存在）
	createTableSQL := `CREATE TABLE IF NOT EXISTS flow (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		dir TEXT NOT NULL,
		success INTEGER DEFAULT 0,
		failure INTEGER DEFAULT 0
	);`
	_, err = db.Exec(createTableSQL)
	return err
}

func main() {
	// 启动时初始化数据库，确保表和文件存在
	if err := initDB(); err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

	r := gin.Default()

	// 配置 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(_ string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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
