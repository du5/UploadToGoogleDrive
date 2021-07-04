package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
)

func Up(srv *drive.Service) {
	f, err := os.Open(*File)
	if err != nil {
		log.Panicln(err.Error())
	}
	info, _ := f.Stat()
	total := FileSizeFormat(info.Size(), false)
	getRate := MeasureTransferRate()
	parents := []string{}
	if *FolderId != "" {
		parents = append(parents, *FolderId)
	}
	upf := srv.Files.Create(&drive.File{Name: info.Name(), Parents: parents})
	upfc := upf.SupportsAllDrives(true).Media(f)
	log.Printf("Uploading %s .", info.Name())
	if upFile, err := upfc.ProgressUpdater(func(current, _ int64) {
		fmt.Printf("Uploaded at %s, %s/%s @ %.2f%%\r", getRate(current), FileSizeFormat(current, false), total, float64(current*100)/float64(info.Size()))
	}).Do(); err == nil {
		log.Printf("https://drive.google.com/file/d/%s/view", upFile.Id)
		if *Share {
			_, _ = srv.Permissions.Create(upFile.Id, &drive.Permission{
				Role: "reader",
				Type: "anyone",
			}).Do()
		} else {
			prl, _ := srv.Permissions.List(upFile.Id).Do()
			for _, v := range prl.Permissions {
				_ = srv.Permissions.Delete(upFile.Id, v.Id).Do()
			}
		}
		log.Println("Done.")
	} else {
		log.Println(err.Error())
	}
}
func Down(srv *drive.Service) {
	dlf := srv.Files.Get(*FileId).SupportsAllDrives(true)
	f, err := dlf.Do()
	if err != nil {
		log.Fatalf("Can not download file.")
	}
	if f.MimeType == "application/vnd.google-apps.folder" {
		log.Fatalf("Does not support folders.")
	}
	log.Printf("Downloading %s .", f.Name)

	out, _ := os.Create(f.Name)

	resp, err := dlf.Download()
	if err != nil {
		log.Fatalf("Download Fatal, %s", err.Error())
	}
	defer resp.Body.Close()
	if n, err := io.Copy(out, resp.Body); err == nil {
		log.Printf("Done. File size: %s.", FileSizeFormat(n, false))
	} else {
		log.Println(err)
	}
}
