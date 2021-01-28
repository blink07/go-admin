package service

import "go-admin/conf/settings"

func GetExcelFullUrl(name string) string {
	println("AAAAAAA"+GetExcelPath())
	return GetExcelPath() + name
}

func GetExcelPath() string {
	return settings.FileSettings.BasePath+settings.ServerSetting.ExcelDir
}

func GetExcelFullPath() string {
	println(settings.FileSettings.PrefixPath + settings.FileSettings.BasePath + GetExcelPath())
	return settings.FileSettings.PrefixPath + settings.FileSettings.BasePath + GetExcelPath()
}

