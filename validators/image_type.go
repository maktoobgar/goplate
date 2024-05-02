package validators

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func ImageType(contentTypes ...string) func(input interface{}) bool {
	return func(input interface{}) bool {
		content, _ := input.(string)
		fileByte, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			return false
		}

		fileContentType := http.DetectContentType(fileByte)
		for _, contentType := range contentTypes {
			if strings.HasPrefix(fileContentType, contentType) {
				return true
			}
		}
		return false
	}
}
