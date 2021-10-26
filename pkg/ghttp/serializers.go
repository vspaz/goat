package ghttp

import (
	"bytes"
	"encoding/json"
)

func toByteBuffer(headers map[string]string, body interface{}) *bytes.Buffer {
	if isJson(headers) {
		encodedPayload, _ := json.Marshal(body)
		return bytes.NewBuffer(encodedPayload)
	}

	return bytes.NewBuffer([]byte{})
}

func FromJson(deserializable []byte, deserialized interface{}) error {
	return json.Unmarshal(deserializable, deserialized)
}

func toJson(serializable interface{}) []byte {
	encodedMessage, _ := json.Marshal(serializable)
	return encodedMessage
}
