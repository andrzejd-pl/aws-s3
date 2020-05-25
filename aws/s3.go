package aws

import (
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/s3"
	"io"
)

const kindConnection = "s3"

type ClientInterface interface {
	Upload(content io.Reader, size int64, name, directory string) (stow.Item, error)
}

type client struct {
	config     stow.ConfigMap
	location   stow.Location
	bucketName string
}

func NewClient(accessKeyId, secretKey, region, bucket string) ClientInterface {
	return &client{
		config: stow.ConfigMap{
			s3.ConfigAccessKeyID: accessKeyId,
			s3.ConfigSecretKey:   secretKey,
			s3.ConfigRegion:      region,
		},
		bucketName: bucket,
	}
}

func (c *client) openConnection() (err error) {
	c.location, err = stow.Dial(kindConnection, c.config)
	return err
}

func (c *client) Upload(content io.Reader, size int64, name, directory string) (stow.Item, error) {
	err := c.openConnection()
	if err != nil {
		return nil, err
	}
	defer c.location.Close()

	container, err := c.location.Container(c.bucketName)
	if err != nil {
		return nil, err
	}
	return container.Put(directory+"/"+name, content, size, nil)
}
