package core

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Configuration struct {
	AccessKey  string `json:"access_key" env:"AWS_ACCESS_KEY"`
	SecretKey  string `json:"secret_key" env:"AWS_SECRET_KEY"`
	Region     string `json:"region" env:"AWS_REGION"`
	BucketName string `json:"bucket" env:"AWS_S3_BUCKET"`
	Directory  string `json:"directory" env:"WATCHED_DIRECTORY"`
}

func (c *Configuration) filledFieldsNumber() (n int) {
	if c.AccessKey != "" {
		n++
	}
	if c.SecretKey != "" {
		n++
	}
	if c.Region != "" {
		n++
	}
	if c.BucketName != "" {
		n++
	}
	if c.Directory != "" {
		n++
	}
	return
}

func (c *Configuration) ReadFromReader(reader io.Reader) (err error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	_, err = c.Read(data)
	return
}

func (c *Configuration) Read(byte []byte) (n int, err error) {
	err = json.Unmarshal(byte, c)
	if err != nil {
		return
	}

	return c.filledFieldsNumber(), nil
}
