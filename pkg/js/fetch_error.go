package js

import (
	"encoding/json"
	"fmt"
	"strings"
)

func fetchHTTPError(status int, statusText, contentType string, body []byte) error {
	message := fmt.Sprintf("HTTP %d: %s", status, statusText)
	if detail := fetchErrorDetail(contentType, body); detail != "" {
		message += ": " + detail
	}
	return fmt.Errorf("%s", message)
}

func fetchErrorDetail(contentType string, body []byte) string {
	text := strings.TrimSpace(string(body))
	if text == "" {
		return ""
	}

	if isJSONContentType(contentType) {
		if message := fetchJSONErrorMessage(body); message != "" {
			return message
		}
	}

	if isJSONContentType(contentType) || isTextContentType(contentType) {
		return text
	}

	return ""
}

func fetchJSONErrorMessage(body []byte) string {
	var value any
	if err := json.Unmarshal(body, &value); err != nil {
		return ""
	}
	return fetchJSONMessage(value)
}

func fetchJSONMessage(value any) string {
	switch value := value.(type) {
	case string:
		return strings.TrimSpace(value)
	case []any:
		for _, item := range value {
			if message := fetchJSONMessage(item); message != "" {
				return message
			}
		}
	case map[string]any:
		for _, key := range []string{"message", "error_description", "error", "detail"} {
			if message := fetchJSONMessage(value[key]); message != "" {
				return message
			}
		}
		encoded, err := json.Marshal(value)
		if err == nil {
			return strings.TrimSpace(string(encoded))
		}
	}

	return ""
}

func isJSONContentType(contentType string) bool {
	contentType = normaliseContentType(contentType)
	return contentType == "application/json" || strings.HasSuffix(contentType, "+json")
}

func isTextContentType(contentType string) bool {
	contentType = normaliseContentType(contentType)
	return strings.HasPrefix(contentType, "text/")
}

func normaliseContentType(contentType string) string {
	contentType = strings.TrimSpace(strings.ToLower(contentType))
	if index := strings.Index(contentType, ";"); index >= 0 {
		contentType = strings.TrimSpace(contentType[:index])
	}
	return contentType
}

func headerValue(headers map[string]string, key string) string {
	for headerKey, value := range headers {
		if strings.EqualFold(headerKey, key) {
			return value
		}
	}
	return ""
}
