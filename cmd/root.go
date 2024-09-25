package main

import (
	"Aterminal/cmd/start"
	"Aterminal/config"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// 假设 config.ConfigCmd 和 start.StartCmd 已定义并导入

var rootCmd = &cobra.Command{
	Use: "sun",
	// 如果没有输入命令，禁止自动输出 Usage 信息
	Run: func(cmd *cobra.Command, args []string) {
		// 当没有输入子命令时，可以自定义处理

	},
}

func Execute() {

	// 关闭 Cobra 的默认错误和 usage 输出
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	// 添加子命令
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(start.StartCmd)

	// 执行 rootCmd，处理命令行参数
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// 创建一个新的扫描器，用于读取用户输入
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// 读取用户输入
		if scanner.Scan() {
			input := scanner.Text()

			// 处理用户输入
			switch strings.TrimSpace(input) {
			case "exit", "quit":
				// 用户输入 'exit' 或 'quit'，程序退出循环
				fmt.Println("Exiting...")
				os.Exit(0)
			default:
				args := strings.Fields(input)
				if len(args) > 0 {
					rootCmd.SetArgs(args)
					if err := rootCmd.Execute(); err != nil {
						fmt.Fprintln(os.Stderr, err)
					}
				}
			}
		} else {
			break
		}
	}

	// 检查是否有错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

func main() {
	// 启动CLI程序
	Execute()
}
