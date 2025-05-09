package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Fahmi36/api/media"
	"github.com/Fahmi36/examples/assets"
	"github.com/Fahmi36/imagekit-go"
)

var ctx = context.Background()

func main() {
	var err error
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	file := assets.UploadFile(ik, "data/nature.jpg")

	var api = ik.Media
	// Add tags
	tagsResp, err := api.UpdateFile(ctx, file.FileId, media.UpdateFileParam{
		Tags: []string{"natural", "mountains", "scene", "day"},
	})
	fmt.Println(tagsResp, err)

	// Remove tags
	remTagResp, err := api.RemoveTags(ctx, media.TagsParam{
		FileIds: []string{file.FileId},
		Tags:    []string{"scene", "day"},
	})
	log.Println(remTagResp, err)

	// Remove AI tags
	remTagResp, err = api.RemoveAITags(ctx, media.AITagsParam{
		FileIds: []string{file.FileId},
		AITags:  []string{"x", "y"}, // replace with real AI tags
	})
	log.Println(remTagResp, err)
}
