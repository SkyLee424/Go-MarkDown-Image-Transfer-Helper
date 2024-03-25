package upload

import (
	"context"
	"log"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
)

var (
	uploadToken string
	cfg         storage.Config
)

func InitQiniu() {
	accessKey := viper.GetString("upload.qiniu.access_key")
	secretKey := viper.GetString("upload.qiniu.secret_key")
	bucket := viper.GetString("upload.qiniu.bucket")

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := auth.New(accessKey, secretKey)
	uploadToken = putPolicy.UploadToken(mac)

	cfg = storage.Config{}

	// 空间对应的机房
	zone := viper.GetString("upload.qiniu.zone")
	switch zone {
	case "z0":
		cfg.Zone = &storage.Zone_z0
	case "z1":
		cfg.Zone = &storage.Zone_z1
	case "z2":
		cfg.Zone = &storage.Zone_z2
	case "na0":
		cfg.Zone = &storage.Zone_na0
	case "as0":
		cfg.Zone = &storage.Zone_as0
	case "cn-east-2":
		cfg.Zone = &storage.ZoneHuadongZheJiang2
	default:
		log.Fatalf("Qiniu Cloud: Invalid region: %s, cur available regions are: \nz0(East China-Zhejiang)\nz1(North China-Hebei)\nz2(South China-Guangdong)\nna0(North America-Los Angeles)\nas0(Asia Pacific-Singapore)\ncn-east-2(East China-Zhejiang-2)", zone)
	}

	// 是否使用https域名
	cfg.UseHTTPS = viper.GetBool("upload.qiniu.use_https")
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = viper.GetBool("upload.qiniu.use_cdn_domains")
}

func UploadFileByQiniu(path, fileName string) error {
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 使用文件名作为 key
	key := fileName

	log.Printf("Uploading file: %v...\n", path)
	err := formUploader.PutFile(context.Background(), &ret, uploadToken, key, path, nil)
	if err != nil {
		return err
	}

	return nil
}
