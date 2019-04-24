package utils

import (
	"encoding/base64"
	"io/ioutil"
	"bytes"
	"log"
)

func Base64Encode(stream *bytes.Buffer) (string, error) {

	readBuf, err := ioutil.ReadAll(stream)

	if err != nil {
		log.Println("Could not read buffer", err)
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(readBuf)

	return str, nil

}
