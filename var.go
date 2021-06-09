package main

import "flag"

var (
	Mode      = flag.String("mode", "upload", "类型 [upload|download]")
	FolderId  = flag.String("folder", "", "Google Drive Folder ID(upload 可选参数)")
	File      = flag.String("path", "", "文件路径(upload 必选参数)")
	FileId    = flag.String("file", "", "Google Drive File ID(download 必选参数)")
	UserClint = flag.String("config", "token.json", "配置文件路径")
	Share     = flag.Bool("share", false, "公开分享该文件(upload 可选参数)")
)
