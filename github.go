package main

import (
	"context"
	"crypto/sha1"
	"time"

	"fortio.org/log"
	"github.com/google/go-github/v57/github"
)

func UploadFile(name string, data []byte) error {
	client := github.NewClient(nil).WithAuthToken("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")

	hasher := sha1.New()
	hasher.Write(data)
	dateNow := time.Now().Format("2006-01-02")

	opts := github.RepositoryContentFileOptions{
		Message: github.String("Selfie: " + dateNow),
		Content: data,
		Branch:  github.String("temp"),
		SHA:     github.String(dateNow),
	}

	fileResp, resp, err := client.Repositories.CreateFile(context.TODO(), repoOwner, repoName, "data/"+name, &opts)
	if err != nil {
		log.Errf("Error uploading file: %v", err)
		return err
	}

	log.Infof("UploadFile: %v %v", fileResp, resp.StatusCode)
	return nil
}
