# Go-MarkDown-Image-Transfer-Helper

## 介绍

本工具是一个用于将 Markdown 中的图片文件从本地迁移到图床的 Go 程序。

适用于：**批量迁移本地图片文件到图床，并修改本地 MarkDown 文件以匹配新的图床链接。**

> 写这个工具的目的是：一开始做笔记都是将文件存储到本地的，但是后面想要分享笔记的时候，又要将图片文件同步分享给他人，非常麻烦
>
> 此外，如果要上传笔记到其它平台（如微信公众号、知乎专栏等），也需要手动上传图片
>
> 于是，我写了一个工具，可以将本地图片文件迁移到图床，并修改本地 MarkDown 文件以匹配新的图床链接。这样，分享笔记，或者上传笔记到其它平台的时候就只需要分享 MarkDown 文件即可，非常方便。

## 功能

- 将本地图片文件 **批量** 迁移到图床
- 修改本地 MarkDown 文件以匹配新的图床链接
- 自定义图床链接格式

## 基本使用

### 准备

首先克隆本仓库到本地：

```zsh
git clone https://github.com/Sky_Lee424/Go-MarkDown-Image-Transfer-Helper.git
```

然后，进入项目目录，安装依赖：

```zsh
cd Go-MarkDown-Image-Transfer-Helper
go mod tidy
```

接下来，需要配置 `config.json` 文件，即图床相关配置：

```json
{
  "modify": {
    "domain": "your-domain"
  },
  "upload": {
    "method": "upload-method",
    "qiniu": {
      "access_key": "your-access-key",
      "secret_key": "your-secret-key",
      "bucket": "your-bucket",
      "zone": "your-img-zone",
      "use_https": false,
      "use_cdn_domains": false
    }
  }
}
```

- **domain**：自定义图床链接前缀，该前缀将会被添加到每个图片链接的前面，例如 `http://savomwqu6.hn-bkt.clouddn.com`
- **method**：上传方式，目前，仅支持 [七牛云](https://portal.qiniu.com/) 图床

最后，编译文件：

```zsh
go build
```

### 使用

#### 帮助信息

在项目目录下，执行以下命令：

```zsh
./go-md-image-transfer-helper -h
```

可以查看帮助信息：

```
Sky_Lee@SkyLeeMacBook-Pro Go-MarkDown-Image-Transfer-Helper % ./go-md-image-transfer-helper -h
Usage of ./go-md-image-transfer-helper:
  -b    Disable backup of text files
  -c string
        Path to configuration file (default "./config.json")
  -f    Force delete image files without confirmation
  -hashed
        Rename uploaded files to their hash value
  -m    Modify text files
  -r    Delete image files
  -u    Upload image files
  -w string
        Root working directory (default ".")
```

**参数说明：**

- **-b**：不备份 markdown 文件
- **-c**：配置文件路径，默认为当前目录下的 `config.json`
- **-f**：强制删除图片文件，不进行确认
- **-hashed**：将上传的图片文件重命名为其哈希值
- **-m**：修改 markdown 文件
- **-r**：删除图片文件
- **-u**：上传图片文件
- **-w**：工作根目录，即项目目录

#### 批量上传图片到图床

在项目目录下，执行以下命令：

```zsh
./go-md-image-transfer-helper -u -w .
```

- **-u**：批量上传图片文件
- **-w**：工作根目录，即项目目录

#### 批量修改图片链接

在项目目录下，执行以下命令：

```zsh
./go-md-image-transfer-helper -m -w .
```

- **-m**：修改文本文件
- **-w**：工作根目录，即项目目录

默认会备份 markdown 文件到原目录下，如果不需要备份，可以加上 **-b** 参数：

#### 批量删除图片

在项目目录下，执行以下命令：

```zsh
./go-md-image-transfer-helper -r -w .
```

- **-r**：删除图片文件
- **-w**：工作根目录，即项目目录

### 示例

先来看看当前目录的情况：

```
Sky_Lee@SkyLeeMacBook-Pro Go-MarkDown-Image-Transfer-Helper % tree
.
├── Note-Test
│   └── Notes
│       ├── Note1
│       │   ├── 2024-03-25-08-39-000.png
│       │   ├── 2024-03-25-08-40-000.png
│       │   └── note.md
│       ├── Note2
│       │   ├── images
│       │   │   ├── 2024-03-25-08-39-000.png
│       │   │   └── 2024-03-25-08-40-000.png
│       │   └── note.md
│       ├── Note3
│       │   ├── 2024-03-25-08-43-000.jpg
│       │   ├── 2024-03-25-08-44-000.jpeg
│       │   ├── 2024-03-25-08-45-000.png
│       │   ├── 2024-03-25-08-46-000.gif
│       │   └── note.md
│       └── Note4
│           ├── 2024-03-25-10-09-000.png
│           ├── 2024-03-25-10-10-000.jpg
│           ├── 2024-03-25-10-11-000.jpeg
│           ├── 2024-03-25-10-12-000.gif
│           └── note.md
├── config.json
├── go-md-image-transfer-helper
```

假设我要上传 Note1 中的图片到图床，并且修改 note.md 的图片链接以匹配图床的图片链接，那么可以执行以下命令：

```zsh
./go-md-image-transfer-helper -m -u -w ./Note-Test/Notes/Note1
```

输出如下：

```
Sky_Lee@SkyLeeMacBook-Pro Go-MarkDown-Image-Transfer-Helper % ./go-md-image-transfer-helper -m -u -w ./Note-Test/Notes/Note1
2024/03/25 10:49:01 path: ./Note-Test/Notes/Note1
2024/03/25 10:49:01 path: Note-Test/Notes/Note1/2024-03-25-08-39-000.png
2024/03/25 10:49:01 Uploading file: Note-Test/Notes/Note1/2024-03-25-08-39-000.png...
2024/03/25 10:49:02 path: Note-Test/Notes/Note1/2024-03-25-08-40-000.png
2024/03/25 10:49:02 Uploading file: Note-Test/Notes/Note1/2024-03-25-08-40-000.png...
2024/03/25 10:49:02 path: Note-Test/Notes/Note1/note.md
2024/03/25 10:49:02 Backing up Note-Test/Notes/Note1/note.md to Note-Test/Notes/Note1/note.md.back ...
2024/03/25 10:49:02 Updating file: Note-Test/Notes/Note1/note.md
```

此时，在 `./Note-Test/Notes/Note1` 目录下，会生成一个 `note.md.back` 文件，并且 `note.md` 文件中的图片链接已经被修改：

```
Sky_Lee@SkyLeeMacBook-Pro Go-MarkDown-Image-Transfer-Helper % cat Note-Test/Notes/Note1/note.md
# 测试笔记

![](http://savomwqu6.hn-bkt.clouddn.com/2024-03-25-08-39-000.png)

![](http://savomwqu6.hn-bkt.clouddn.com/2024-03-25-08-40-000.png)
```

并且，图片也成功上传到图床了：

![image](http://images.blogs.skylee.top/2024-03-25-10-51-56-248.png)

如果你有多个名称相同的图片，那么你可以使用 `-hashed` 参数，这样，图片的哈希值将作为图片的名称，避免重复：

```zsh
./transfer-image-helper -m -u -w ./Note-Test/Notes/Note1 -hashed
```

如果你想在上传完图片后，删除图片，那么可以使用 `-r` 参数，进一步地，可以使用 `-f` 参数，强制删除图片（**不建议**）

此外，如果 **不想备份** 原有 markdown 文件，可以使用 `-b` 参数（**不建议**）

> 注意：在第一次使用 Go-MarkDown-Image-Transfer-Helper 时，建议先操作 `Test-Notes` 的文件，熟悉一下基本步骤

## 引入其它图床

如果你使用的图床不是七牛云，而是其它平台，那么需要手动修改源码，具体地：

- 在 upload 包添加一个 API，方法格式参考 `upload/qiniu.go`
- 修改 `main.go`：
  ```go
  method := viper.GetString("upload.method")
  switch method {
  case "qiniu":
  	upload.InitQiniu()
  	uploadIMGFile = upload.UploadFileByQiniu
  case "your-method":
      upload.InitYourMethod()
  	uploadIMGFile = upload.UploadFileByYourMethod
  default:
  	log.Fatal("未支持的上传方式")
  }
  ```

欢迎给本项目提 PR，帮助更多人迁移自己的本地图片到图床～～

---

ToDo:

- [ ] 引入并行上传功能
