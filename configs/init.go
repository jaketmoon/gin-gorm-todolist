package configs

import "gin/internal/global/log"

// 1.从yaml中读取数据NewSettings 2.从viper中获取能用的数据readSection

func Init() {
	err := SetUpSettings()
	if err != nil {
		log.SugarLogger.Error(err)
		return
	}
}
