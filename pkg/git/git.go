package git

import (
	"errors"
	"fmt"
	"github.com/starkers/ggg/pkg/logger"
	"strings"

	"github.com/mitchellh/go-homedir"
)

func FigureUnixDiskPath(localPath string, url string) (string, error) {

	if url == "" {
		err := errors.New("blank git URL received, please provide a URL argument")
		return "", err
	}

	expandedPath, err := homedir.Expand(localPath)
	if err != nil {
		logger.Bad(err)
	}
	subDir := mungeUrl(url)

	result := fmt.Sprintf("%s/%s", expandedPath, subDir)
	return result, nil
}

func mungeUrl(input string) string {
	rightResult := ""
	hostname := ""
	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		slicedHTTP := strings.Split(input, "/")
		slashCount := len(slicedHTTP)
		hostname = slicedHTTP[2]
		rightBlobSliced := slicedHTTP[3:slashCount]
		rightBlob := strings.Join(rightBlobSliced, "/")
		if strings.HasSuffix(rightBlob, ".git") {
			rightResult = strings.Replace(rightBlob, ".git", "", 1)
		} else {
			rightResult = rightBlob
		}
		return fmt.Sprintf("%s/%s", hostname, rightResult)
	}
	if strings.HasPrefix(input, "git@") {
		slicedColon := strings.Split(input, ":")
		if len(slicedColon) > 1 {
			hostname = slicedColon[0]
			hostname = strings.Replace(hostname, "git@", "", 1)
			rightBlob := slicedColon[(len(slicedColon) - 1)]
			if strings.HasSuffix(rightBlob, ".git") {
				rightResult = strings.Replace(rightBlob, ".git", "", 1)
			} else {
				rightResult = rightBlob
			}
			return fmt.Sprintf("%s/%s", hostname, rightResult)
		}
	}
	return "unknown"
}
