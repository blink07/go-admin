package service

import (
	"fmt"
	"go-admin/api/utils/hash"
	"go-admin/conf/settings"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func FullImagePath(name string) string {
	return settings.FileSettings.PrefixPath + settings.FileSettings.BasePath + settings.FileSettings.ImagePath +name
}

// 获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)  // 取后缀
	println(ext)
	fileName := strings.TrimSuffix(name, ext)
	fileName = hash.EncodeMD5(fileName)

	return fileName + ext
}


// 获取存储图片存取路径
func GetImagePath() string {
	return settings.FileSettings.BasePath + settings.FileSettings.ImagePath
}

// 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// 检查图片后缀
func CheckImageExt(fileName string) bool {
	ext:=GetExt(fileName)

	for _, allowExt := range settings.FileSettings.ImageAllowExts{
		println("<<<<<<<<<<<<<1111")
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

// 检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := GetSize(f)
	if err != nil {
		println(">>>>>>>>>>>>>",err.Error())
		return false
	}
	println(settings.FileSettings.ImageMaxSize> size)
	return size <= settings.FileSettings.ImageMaxSize
}

// 检查文件是否有权限
func CheckPermission(src string) bool {
	_,err := os.Stat(src)
	return os.IsPermission(err)
}

// 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err!= nil {
		return err
	}
	return nil
}

// 检查文件是否存在
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// 如果不存在则创建文件夹
func IsNotExistMkdir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err:=MkDir(src); err!=nil {
			return err
		}
	}
	return nil
}

// 检查图片路径和权限
func CheckImage(src string) error {
	dir, err := os.Getwd()

	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = IsNotExistMkdir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}