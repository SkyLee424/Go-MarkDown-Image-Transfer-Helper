# Go-MarkDown-Image-Transfer-Helper

## Introduction

English | [简体中文](./README.zh-CN.md)

This tool is a Go program designed for migrating image files from local Markdown directories to an image hosting service.

It is suitable for: **batch migration of local image files to an image hosting service and updating local Markdown files to match new image host URLs.**

> The purpose of writing this tool was: initially, I stored files locally when taking notes, but later on, when wanting to share these notes, I had to synchronize and share the image files with others separately, which was very inconvenient.
> 
> Additionally, if you want to upload notes to other platforms (like WeChat Official Accounts, Zhihu Columns, etc.), you have to manually upload the images each time.
> 
> Therefore, I wrote a tool that would migrate local image files to an image hosting service and update local Markdown files to match the new image host URLs. Thus, when sharing notes or uploading them to other platforms, only the Markdown files need to be shared, which is very convenient.

## Features

- **Batch** migration of local image files to an image hosting service
- Update local Markdown files to match new image host URLs
- Customizable image host URL format

## Basic Usage

### Preparation

First, clone this repository locally:

```zsh
git clone https://github.com/SkyLee424/Go-MarkDown-Image-Transfer-Helper.git
```

Next, enter the project directory and install dependencies:

```zsh
cd Go-MarkDown-Image-Transfer-Helper
go mod tidy
```

After that, you need to configure the `config.json` file, which contains the image hosting configuration:

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

- **domain**: Custom image host URL prefix that will be added to the front of every image link, for example, `http://savomwqu6.hn-bkt.clouddn.com`
- **method**: The method of upload, for now, only [Qiniu Cloud](https://portal.qiniu.com/) image hosting is supported

Finally, compile the file:

```zsh
go build
```

### Usage

#### Help Information

Execute the following command in the project directory:

```zsh
./go-md-image-transfer-helper -h
```

You can view the help information:

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

#### Batch Upload Images to Image Hosting

Execute the following command in the project directory:

```zsh
./go-md-image-transfer-helper -u -w .
```

- **-u**: Batch upload of image files
- **-w**: Root working directory, i.e., project directory

#### Batch Modify Image Links

Execute the following command in the project directory:

```zsh
./go-md-image-transfer-helper -m -w .
```

- **-m**: Modify text files
- **-w**: Root working directory, i.e., project directory

By default, markdown files are backed up in the original directory. If you do not need to backup, you can add the **-b** parameter:

#### Batch Delete Images

Execute the following command in the project directory:

```zsh
./go-md-image-transfer-helper -r -w .
```

- **-r**: Delete image files
- **-w**: Root working directory, i.e., project directory

### Example

Let's take a look at the current directory:

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

Assume I want to upload the images in Note1 to the image host and modify the `note.md` image links to match the image host, then I can execute the following command:

```zsh
./go-md-image-transfer-helper -m -u -w ./Note-Test/Notes/Note1
```

The output is as follows:

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

As a result, a `note.md.back` file is generated in the `./Note-Test/Notes/Note1` directory, and the image links in the `note.md` file have been modified:

```zsh
Sky_Lee@SkyLeeMacBook-Pro Go-MarkDown-Image-Transfer-Helper % cat Note-Test/Notes/Note1/note.md
# Test Notes

![](http://savomwqu6.hn-bkt.clouddn.com/2024-03-25-08-39-000.png)

![](http://savomwqu6.hn-bkt.clouddn.com/2024-03-25-08-40-000.png)
```

Moreover, the images have been successfully uploaded to the image host:

![image](http://images.blogs.skylee.top/2024-03-25-10-51-56-248.png)

If you have multiple images with the same name, you can use the `-hashed` parameter, so the hash value of the image will be used as the image name to avoid duplicates:

```zsh
./transfer-image-helper -m -u -w ./Note-Test/Notes/Note1 -hashed
```

If you want to delete the images after uploading, you can use the `-r` parameter. Furthermore, you can use the `-f` parameter to force delete the images (**not recommended**).

In addition, if you **do not want to back up** the original markdown files, you can use the `-b` parameter (**not recommended**).

> Note: When using Go-MarkDown-Image-Transfer-Helper for the first time, it is recommended to first operate on the `Test-Notes` files to get familiar with the basic steps.

## Incorporating Other Image Hosts

If your image hosting service is not Qiniu Cloud but another platform, you will need to manually modify the source code. Specifically:

- Add an API in the upload package, referring to the method format in `upload/qiniu.go`.
- Modify `main.go`:
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
    log.Fatal("Unsupported upload method")
  }
  ```

You are welcome to contribute a PR to this project, assisting others in the migration of their local images to an image hosting service!

---

ToDo:

- [ ] Introduce parallel upload functionality