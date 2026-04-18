/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"encoding/json"
	"flow/internal/jsons"
	"flow/pkg/color"
	ffile "flow/pkg/file"
	"os"
	"os/exec"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		curpath := ffile.GetCurPath()
		defer func() {
			if r := recover(); r != nil {
				// 任务执行失败，更新数据库
				updateDatabase(ffile.GetCWD(), false)
			}
		}()

		if len(args) < 1 {
			return
		} else if len(args) > 1 {
			color.Warning("too many args")
		}
		flowName := args[0]
		color.Info("running flow " + flowName)
		runflow := ".flow/" + flowName + ".flow"
		if !ffile.FileExists(runflow) {
			return
		}
		flowe := curpath + "/flowe"
		flowc := curpath + "/flowc"
		if (!ffile.FileExists(flowe)) || 
			(!ffile.FileExists(flowc)) {
			color.Error("binary is not installed properly")
		}
		compilation := exec.Command(flowc, runflow)
		op, err := compilation.CombinedOutput()
		if err != nil {
			updateDatabase(ffile.GetCWD(), false)
			color.Error("compilation failed")
			return
		}
		file := strings.TrimSpace(string(op))
		context, err := os.ReadFile(file)
		if err != nil {
			color.Error("read file failed")
			return
		}
		var p jsons.Plan
		err = json.Unmarshal(context, &p)
		if err != nil {
			color.Error("unmarshal failed")
			return
		}
		color.Info("executing workflow " + p.Workflow)
		steps := p.Steps
		/* DFS 检测Steps是否存在循环依赖 */
		// 构建步骤名称到步骤的映射
		stepMap := make(map[string]*jsons.Step)
		for i := range steps {
			stepMap[steps[i].Name] = &steps[i]
		}

		visited := make(map[string]bool)
		recStack := make(map[string]bool)

		var hasCycle bool
		var dfs func(stepName string) bool
		dfs = func(stepName string) bool {
			if !visited[stepName] {
				visited[stepName] = true
				recStack[stepName] = true

				step, exists := stepMap[stepName]
				if exists {
					for _, dep := range step.Deps {
						if !visited[dep] {
							if dfs(dep) {
								return true
							}
						} else if recStack[dep] {
							return true
						}
					}
				}
			}
			recStack[stepName] = false
			return false
		}

		for _, step := range steps {
			if !visited[step.Name] {
				if dfs(step.Name) {
					hasCycle = true
					break
				}
			}
		}

		if hasCycle {
			color.Error("workflow has cyclic dependencies")
			return
		}

		executed := make(map[string]bool)
		for _, step := range steps {
			executed[step.Name] = false
		}
		for len(steps) > 0 {
			step := steps[0]
			canexec := true
			for _, dep := range step.Deps {
				if !executed[dep] {
					canexec = false
					break
				}
			}
			if !canexec {
				steps = append(steps[1:], step)
				continue
			} else {
				executed[step.Name] = true
				steps = steps[1:]
				color.Info("executing step " + step.Name)
				execution := exec.Command(flowe, file, step.Name)
				err = execution.Run()
				if err != nil {
					updateDatabase(ffile.GetCWD(), false)
					color.Error("execution failed")
					return
				}
			}
		}

		// 任务执行成功，更新数据库
		updateDatabase(ffile.GetCWD(), true)
	},
}

func updateDatabase(dir string, success bool) {
	color.Info("connecting database")
	// 打开数据库连接
	db, err := sql.Open("sqlite3", 
		ffile.GetCurPath() + "/flow.db")
	if err != nil {
		color.Error("failed to open database: " + err.Error())
		return
	}
	defer db.Close()

	// 创建表（如果不存在）
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS flow (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		dir TEXT UNIQUE,
		success INTEGER DEFAULT 0,
		failure INTEGER DEFAULT 0
	)`)
	if err != nil {
		color.Error("failed to create table: " + err.Error())
		return
	}

	// 检查目录是否存在
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM flow WHERE dir = ?", dir).Scan(&count)
	if err != nil {
		color.Error("failed to query database: " + err.Error())
		return
	}

	if count > 0 {
		// 目录存在，更新数据
		if success {
			_, err = db.Exec("UPDATE flow SET success = success + 1 WHERE dir = ?", dir)
		} else {
			_, err = db.Exec("UPDATE flow SET failure = failure + 1 WHERE dir = ?", dir)
		}
	} else {
		// 目录不存在，插入数据
		if success {
			_, err = db.Exec("INSERT INTO flow (dir, success, failure) VALUES (?, 1, 0)", dir)
		} else {
			_, err = db.Exec("INSERT INTO flow (dir, success, failure) VALUES (?, 0, 1)", dir)
		}
	}

	if err != nil {
		color.Error("failed to update database: " + err.Error())
		return
	}
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
