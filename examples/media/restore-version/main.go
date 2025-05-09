package main

import (
	"context"
	"log"

	"github.com/Fahmi36/imagekit-go"
	"github.com/Fahmi36/imagekit-go/api/media"
)

var ctx = context.Background()

func main() {
	var err error
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	// replace FileId and VersionId with real values
	resp, err := ik.Media.RestoreVersion(ctx, media.FileVersionsParam{
		FileId:    "",
		VersionId: "",
	})
	log.Println(resp, err)
}
