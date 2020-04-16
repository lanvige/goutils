package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// DeepCopy DeepCopy
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// DeepCopyByJSON DeepCopyByJSON
func DeepCopyByJSON(dst, src interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}

	return json.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)

}
