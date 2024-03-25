package modify

import (
	"fmt"
	"go-md-image-transfer-helper/utils"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func UpdateFileContent(filePath, domain string, hashed bool) error {
	log.Printf("Updating file: %v\n", filePath)
	domainPattern := regexp.MustCompile(`^http(s?)://`)
	imagePattern := regexp.MustCompile(`!\[(.*?)\]\((.*?\.(?:png|jpg|jpeg|gif))\)`)

	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %v: %v\n", filePath, err)
	}

	// Replace local files path with the full URL to the image hosted on domain
	newContent := imagePattern.ReplaceAllStringFunc(string(content), func(submatch string) string {
		matches := imagePattern.FindStringSubmatch(submatch)
		if len(matches) == 3 {
			altText := matches[1]
			imagePath := matches[2]

			// 检查这个图片是否已经是一个 url
			if domainPattern.MatchString(imagePath) {
				// 已经是一个 url，不需要修改
				return submatch
			}

			if hashed {
				basePath := filepath.Dir(filePath) // 获取 md 文件所在目录
				imagePath = basePath + "/" + imagePath
				// 以完整相对路径作为文件名，计算哈希值
				imagePath = utils.GenHashValue(imagePath)
			} else {
				// 仅保留 image 的原始文件名
				imagePath = filepath.Base(imagePath)
			}

			url := fmt.Sprintf("%s/%s", domain, imagePath)
			return fmt.Sprintf(`![%s](%s)`, altText, url)
		}
		return submatch
	})

	// 更新 markdown
	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
