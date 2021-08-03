package git

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/starkers/ggg/pkg/logger"

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
	commonPrefixList := []string{
		"http://",
		"https://",
	}

	keybasePrefixList := []string{
		"keybase://",
	}

	for _, prefix := range keybasePrefixList {
		if strings.HasPrefix(input, prefix) {
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
			return fmt.Sprintf("keybase/%s/%s", hostname, rightResult)
		}
	}

	for _, prefix := range commonPrefixList {
		if strings.HasPrefix(input, prefix) {
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
	}

	gitRegex := ".*@.*"
	match, _ := regexp.MatchString(gitRegex, input)
	if match {
		slicedColon := strings.Split(input, ":")
		if len(slicedColon) > 1 {
			hostname = slicedColon[0]
			x := strings.Split(hostname, "@")

			rightBlob := slicedColon[(len(slicedColon) - 1)]
			if strings.HasSuffix(rightBlob, ".git") {
				rightResult = strings.Replace(rightBlob, ".git", "", 1)
			} else {
				rightResult = rightBlob
			}
			return fmt.Sprintf("%s/%s", x[1], rightResult)
		}
	}
	return "unknown"
}
