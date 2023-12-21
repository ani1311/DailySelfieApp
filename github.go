package main

import (
	"context"
	"time"

	"fortio.org/log"
	"github.com/google/go-github/v57/github"
)

func UploadFile(name string, data []byte) error {
	client := github.NewClient(nil).WithAuthToken("")

	opts := github.RepositoryContentFileOptions{
		Message: github.String("Selfie: " + time.Now().Format("2006-01-02")),
		Content: data,
		Branch:  github.String("temp"),
	}

	fileResp, resp, err := client.Repositories.CreateFile(context.TODO(), repoOwner, repoName, "data/"+name, &opts)
	if err != nil {
		return err
	}

	log.Infof("UploadFile: %v %v", fileResp, resp.StatusCode)
	return nil
}
