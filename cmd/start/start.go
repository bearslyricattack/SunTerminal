package start

import (
	"Aterminal/ai/chatgpt"
	"Aterminal/config"
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "begin ai command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("进入交互模式，输入命令，执行后附加操作")
		interactiveMode()
	},
}

func init() {
	StartCmd.PersistentFlags().StringVar(&config.Path, "path", "/", "config config path")
}

func interactiveMode() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")

		// 读取用户输入
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// 如果用户输入 "exit"，退出交互模式
		if input == "exit" {
			fmt.Println("退出交互模式")
			os.Exit(0)
			break
		}

		// 执行用户命令
		res, _ := executeCommand(input)
		fmt.Println(res)
		fmt.Println("---------------------------")
		//命令返回结果给ai分析
		res, err := chatgpt.Query(config.Path, res+"用中文分析一下这个命令的结果")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}

//func executeCommand(command string, args ...string) (string, error) {
//	// 创建命令对象
//	cmd := exec.Command(command, args...)
//
//	// 获取命令的输出
//	output, err := cmd.CombinedOutput()
//
//	// 忽略错误，直接返回输出
//	if err != nil {
//		fmt.Printf("执行命令时出错: %v，但继续执行\n", err)
//	}
//
//	// 返回输出，忽略错误
//	return string(output), err
//}

func executeCommand(input string) (string, error) {
	// 将用户输入的命令分割成命令和参数
	args := strings.Fields(input)
	if len(args) == 0 {
		return "", fmt.Errorf("no command provided")
	}

	// 使用 exec.Command 执行命令
	cmd := exec.Command(args[0], args[1:]...)

	// 获取命令的输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
