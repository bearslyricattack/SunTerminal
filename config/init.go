package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Path string

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "config config path",
	Run: func(cmd *cobra.Command, args []string) {
		err := createYMLFile(Path)
		if err != nil {
			println("create config error:", err.Error())
			log.Fatal(err)
		}
		println("create config success!")
	},
}

func init() {
	ConfigCmd.PersistentFlags().StringVar(&Path, "path", "/", "config config path")
}

func createSampleConfig() Config {
	return Config{
		Model: []Model{
			{
				Name: "",
				Path: "",
				Key:  "",
				Type: "",
			},
		},
		Context: "",
	}
}

// 创建yml文件
func createYMLFile(path string) error {
	// 检查路径是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", path)
	}

	// 检查文件是否存在
	ymlFilePath := path + "/" + "config.yml"
	fmt.Println("写入路径为：" + ymlFilePath)
	if _, err := os.Stat(ymlFilePath); err == nil {
		return fmt.Errorf("file already exists: %s", path)
	}

	// 创建示例配置对象
	config := createSampleConfig()

	// 将配置对象编码为 YAML 格式
	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Failed to marshal config to YAML: %v", err)
	}

	// 将 YAML 数据写入文件
	file, err := os.Create(ymlFilePath)
	if err != nil {
		log.Fatalf("Failed to create YAML config: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Failed to close config: %v", err)
		}
	}()

	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("Failed to write to YAML config: %v", err)
	}
	fmt.Println("YAML config created successfully")

	//修改path的位置
	Path = path

	return nil
}

// GetModelInfo 根据配置文件中的内容，获取模型相关配置
func GetModelInfo(path string) (*Model, error) {
	//获取配置文件
	file, err := os.ReadFile(path + "/" + "config.yml")
	if err != nil {
		log.Fatalf("Failed to read config.yml: %v", err)
	}
	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	//获取context name
	contextName := config.Context
	//根据context，获取匹配的model位置
	for _, model := range config.Model {
		if model.Name == contextName {
			return &model, nil
		}
	}
	//如果没有找到，返回错误
	return nil, fmt.Errorf("context %s not found", contextName)
}
