package configs

import "github.com/spf13/viper" // 引入viper包，用于处理配置文件

// ReadSection 方法将配置文件中的特定部分解码并绑定到传入的结构体v中。
func (s *Setting) ReadSection(name string, v interface{}) error {
	err := s.vp.UnmarshalKey(name, v) // 使用viper的UnmarshalKey方法读取指定的键，并将其解码到结构体v中
	if err != nil {
		return err // 如果解码过程中出现错误，则返回错误
	}

	return nil // 如果没有错误，则返回nil表示成功
}

// NewSetting 创建并返回一个包含viper实例的Setting对象，有错误则返回错误。
func NewSetting() (*Setting, error) {
	vp := viper.New()             // 使用viper库创建一个新的viper实例
	vp.SetConfigName("config")    // 设置配置文件的名字，不包含文件后缀
	vp.AddConfigPath("./configs") // 添加配置文件的搜索路径，这里是相对路径下的configs目录
	vp.AddConfigPath(".")         // 添加另一个配置文件的搜索路径，这里是当前工作目录
	vp.SetConfigType("yaml")      // 设置配置文件的类型，这里是YAML

	err := vp.ReadInConfig() // 读取和解析配置文件
	if err != nil {
		return nil, err // 如果读取配置文件出错，则返回错误
	}
	return &Setting{vp}, nil // 如果没有错误，则返回一个包含viper实例的Setting对象
}

// SetUpSettings 用于初始化全局配置，通过调用NewSetting来读取配置文件，
// 并将配置文件中的特定部分映射到全局变量。
func SetUpSettings() error {
	setting, err := NewSetting() // 创建Setting实例，该实例包含了读取的配置
	if err != nil {
		return err // 如果在创建Setting实例时遇到错误，则返回该错误
	}
	// 以下代码可以针对配置文件中的多个不同部分进行处理

	// 读取配置文件中的"Database"部分，并将其内容赋值给DbSettings全局变量
	err1 := setting.ReadSection("Database", &DbSettings)
	if err1 != nil {
		return err1 // 如果读取"Database"部分失败，则返回错误
	}
	err2 := setting.ReadSection("jwt", &JwtSettings)
	if err2 != nil {
		return err2 // 如果读取"jwt"部分失败，则返回错误

	}
	// 如果所有配置部分都成功读取，则返回nil
	return nil
}
