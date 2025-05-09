package main

import (
	"context"
	"log"

	"github.com/Fahmi36/imagekit-go"
	"github.com/Fahmi36/imagekit-go/api/media"
	"github.com/Fahmi36/imagekit-go/examples/assets"
)

var ctx = context.Background()

func main() {
	var err error
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	file := assets.UploadFile(ik, "data/nature.jpg")

	log.Println(file.Url)

	var param = media.PurgeCacheParam{
		Url: file.Url,
	}

	response, err := ik.Media.PurgeCache(ctx, param)
	log.Println(response, err)

	statusResp, err := ik.Media.PurgeCacheStatus(ctx, response.Data.RequestId)

	log.Println(statusResp.Data.Status, err)
}
