package main

import (
	"flag"
	"fmt"
	"go-md-image-transfer-helper/config"
	"go-md-image-transfer-helper/modify"
	"go-md-image-transfer-helper/upload"
	"go-md-image-transfer-helper/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	c      string // config.json 的路径
	w      string // 工作目录
	m      bool   // 是否修改文本文件
	u      bool   // 是否上传图片文件
	r      bool   // 是否删除图片文件
	f      bool   // 是否强制删除
	b      bool   // 禁用备份文本文件
	hashed bool   // 是否修改上传文件名为其哈希值
)

var (
	uploadIMGFile upload.UploadMethod // 上传方式
	fileExt       string              // 待修改文本文件的后缀
)

func init() {
	// 初始化 flags
	flag.StringVar(&c, "c", "./config.json", "Path to configuration file")
	flag.StringVar(&w, "w", ".", "Root working directory")
	flag.BoolVar(&m, "m", false, "Modify text files")
	flag.BoolVar(&u, "u", false, "Upload image files")
	flag.BoolVar(&r, "r", false, "Delete image files")
	flag.BoolVar(&f, "f", false, "Force delete image files without confirmation")
	flag.BoolVar(&b, "b", false, "Disable backup of text files")
	flag.BoolVar(&hashed, "hashed", false, "Rename uploaded files to their hash value")

	flag.Parse()

	if w == "." {
		var err error
		w, err = os.Getwd()
		if err != nil {
			log.Fatalf("get work dir failed: %v\n", err)
		}
	}

	config.InitConfig(c)

	method := viper.GetString("upload.method")
	switch method {
	case "qiniu":
		upload.InitQiniu()
		uploadIMGFile = upload.UploadFileByQiniu
	default:
		log.Fatal("invalid upload method")
	}

	fileExt = ".md" // 仅支持 markdown 文件
}

func main() {
	err := WalkDir(w, viper.GetString("modify.domain"))
	if err != nil {
		log.Fatalf("fatal error: %v\n", err)
		return
	}
}

func WalkDir(dirPath string, domain string) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		log.Printf("path: %v\n", path)

		if !info.IsDir() {
			switch filepath.Ext(path) {
			// 如果是图片类型
			case ".png", ".jpg", ".jpeg", ".gif":
				if u {
					fileName := filepath.Base(path)
					if hashed {
						fileName = utils.GenHashValue(path)
					}
					err = uploadIMGFile(path, fileName)
					if err != nil {
						return err
					}
				}
				if r {
					remove := f
					if !f {
						// 询问是否需要删除
						log.Printf("remove file %v? (y/n): ", path)
						var s string
						fmt.Scanln(&s)
						if s == "y" {
							remove = true
						}
					}
					// 删除图片文件
					if remove {
						err = os.Remove(path)
					}
				}

			// 如果是目标文本文件类型
			case fileExt:
				if !m {
					break
				}
				if !b {
					// 备份
					log.Printf("Backing up %v ...", path)
					err = utils.BackupFile(path)
					if err != nil {
						return err
					}
				}
				err = modify.UpdateFileContent(path, domain, hashed)
			}

			if err != nil {
				return err
			}
		}
		return nil
	})
}
