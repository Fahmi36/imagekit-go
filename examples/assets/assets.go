package assets

import (
	"context"
	"embed"
	"log"
	"path/filepath"

	"github.com/Fahmi36/api/uploader"
	"github.com/Fahmi36/imagekit-go"
)

//go:embed data
var Fs embed.FS
var ctx = context.Background()

func UploadFile(ik *imagekit.ImageKit, path string) uploader.UploadResult {
	var err error

	file, err := Fs.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, fileName := filepath.Split(path)

	resp, err := ik.Uploader.Upload(ctx, file, uploader.UploadParam{
		FileName: fileName,
	})

	return resp.Data
}
