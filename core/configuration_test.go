package core

import (
	"strings"
	"testing"
)

func TestConfiguration_Read(t *testing.T) {
	type fields struct {
		AccessKey  string
		SecretKey  string
		Region     string
		BucketName string
		Directory  string
	}
	tests := []struct {
		name       string
		wantConfig fields
		json       string
		wantN      int
		wantErr    bool
	}{
		{
			"empty test",
			fields{
				AccessKey:  "",
				SecretKey:  "",
				Region:     "",
				BucketName: "",
				Directory:  "",
			},
			"{}",
			0,
			false,
		},
		{
			"error test",
			fields{
				AccessKey:  "",
				SecretKey:  "",
				Region:     "",
				BucketName: "",
				Directory:  "",
			},
			"",
			0,
			true,
		},
		{
			"good test",
			fields{
				AccessKey:  "testAccessKey",
				SecretKey:  "testSecretKey",
				Region:     "testRegion",
				BucketName: "testBucket",
				Directory:  "testDirectory",
			},
			"{\"access_key\":\"testAccessKey\",\"secret_key\":\"testSecretKey\"," +
				"\"bucket\":\"testBucket\",\"directory\":\"testDirectory\",\"region\":\"testRegion\"}",
			5,
			false,
		},
		{
			"good test",
			fields{
				AccessKey:  "testAccessKey",
				SecretKey:  "testSecretKey",
				Region:     "testRegion",
				BucketName: "testBucket",
				Directory:  "testDirectory",
			},
			"{\"secret_key\":\"testSecretKey\"," +
				"\"bucket\":\"testBucket\",\"directory\":\"testDirectory\",\"region\":\"testRegion\"}",
			4,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Configuration{}
			gotN, err := c.Read([]byte(tt.json))
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Read() gotN = %v, want %v", gotN, tt.wantN)
			}
			if c.AccessKey != tt.wantConfig.AccessKey &&
				c.SecretKey != tt.wantConfig.SecretKey &&
				c.Region != tt.wantConfig.Region &&
				c.BucketName != tt.wantConfig.BucketName &&
				c.Directory != tt.wantConfig.Directory {
				t.Errorf("Read() struct = %v, want %v", c, tt.wantConfig)
			}
		})
	}
}

func TestConfiguration_ReadFromReader(t *testing.T) {
	type fields struct {
		AccessKey  string
		SecretKey  string
		Region     string
		BucketName string
		Directory  string
	}
	tests := []struct {
		name       string
		wantConfig fields
		json       string
		wantErr    bool
	}{
		{
			"empty test",
			fields{
				AccessKey:  "",
				SecretKey:  "",
				Region:     "",
				BucketName: "",
				Directory:  "",
			},
			"{}",
			false,
		},
		{
			"error test",
			fields{
				AccessKey:  "",
				SecretKey:  "",
				Region:     "",
				BucketName: "",
				Directory:  "",
			},
			"",
			true,
		},
		{
			"good test",
			fields{
				AccessKey:  "testAccessKey",
				SecretKey:  "testSecretKey",
				Region:     "testRegion",
				BucketName: "testBucket",
				Directory:  "testDirectory",
			},
			"{\"access_key\":\"testAccessKey\",\"secret_key\":\"testSecretKey\"," +
				"\"bucket\":\"testBucket\",\"directory\":\"testDirectory\",\"region\":\"testRegion\"}",
			false,
		},
		{
			"good test",
			fields{
				AccessKey:  "testAccessKey",
				SecretKey:  "testSecretKey",
				Region:     "testRegion",
				BucketName: "testBucket",
				Directory:  "testDirectory",
			},
			"{\"secret_key\":\"testSecretKey\"," +
				"\"bucket\":\"testBucket\",\"directory\":\"testDirectory\",\"region\":\"testRegion\"}",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Configuration{}
			reader := strings.NewReader(tt.json)
			if err := c.ReadFromReader(reader); (err != nil) != tt.wantErr {
				t.Errorf("ReadFromReader() error = %v, wantErr %v", err, tt.wantErr)
			}
			if c.AccessKey != tt.wantConfig.AccessKey &&
				c.SecretKey != tt.wantConfig.SecretKey &&
				c.Region != tt.wantConfig.Region &&
				c.BucketName != tt.wantConfig.BucketName &&
				c.Directory != tt.wantConfig.Directory {
				t.Errorf("Read() struct = %v, want %v", c, tt.wantConfig)
			}
		})
	}
}
