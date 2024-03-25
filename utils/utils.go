package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func GenHashValue(s string) string {
	// 创建一个新的 SHA256 哈希器实例
	hasher := sha256.New()
	// 将字符串写入到哈希器
	hasher.Write([]byte(s))
	// 计算哈希值并获得字节切片结果
	sha := hasher.Sum(nil)
	// 将字节切片结果转换成16进制表示的字符串
	hashValue := hex.EncodeToString(sha)
	return hashValue
}

func BackupFile(path string) error {
	// 打开原始文件
	originalFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer originalFile.Close()

	// 备份文件（copy 一份，加上后缀 .back）
	backupPath := path + ".back"
	// 循环，直到不存在
	for i := 1; ; i++ {
		if _, err := os.Stat(backupPath); os.IsNotExist(err) {
			break
		}
		backupPath = path + fmt.Sprintf(".back.%d", i)
	}

	// 创建新的文件
	file, err := os.Create(backupPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 复制内容到备份文件
	_, err = io.Copy(file, originalFile)

	return err
}
