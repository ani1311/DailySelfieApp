package main

import (
	"context"
	"crypto/sha1"
	"time"

	"fortio.org/log"
	"github.com/google/go-github/v57/github"
)

func UploadFile(data []byte) error {
	client := github.NewClient(nil).WithAuthToken("")

	hasher := sha1.New()
	hasher.Write(data)
	dateNow := time.Now().Format("2006-01-02")
	path := time.Now().Format("2006/01")
	name := time.Now().Format("02") + ".png"

	opts := github.RepositoryContentFileOptions{
		Message: github.String("Selfie: " + dateNow),
		Content: data,
		Branch:  github.String("selfies"),
		SHA:     github.String(dateNow),
	}

	fileResp, resp, err := client.Repositories.CreateFile(context.TODO(), repoOwner, repoName, path+"/"+name, &opts)
	if err != nil {
		log.Errf("Error uploading file: %v", err)
		return err
	}

	log.Infof("UploadFile: %v %v", fileResp, resp.StatusCode)
	return nil
}
