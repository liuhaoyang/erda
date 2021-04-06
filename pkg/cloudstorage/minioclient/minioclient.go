// Copyright 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package minioclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/minio/minio-go"
	"github.com/pkg/errors"
)

type MinioClient struct {
	endpoint  string
	accessKey string
	secretKey string
	client    *minio.Client
}

func New(endpoint, accessKey, secretKey string) (*MinioClient, error) {
	secure := strings.HasPrefix(endpoint, "https")
	minioclient, err := minio.New(endpoint, accessKey, secretKey, secure)
	if err != nil {
		return nil, err
	}

	client := MinioClient{
		endpoint:  endpoint,
		accessKey: accessKey,
		secretKey: secretKey,
		client:    minioclient,
	}

	if err := client.HealthCheck(); err != nil {
		return nil, err
	}
	return &client, nil
}

func (c *MinioClient) UploadFile(bucketName, objectName, file string) (string, error) {
	contentType, err := detectFileType(file)
	if err != nil {
		return "", errors.Wrapf(err, "upload file with bucketName=%s, objectName=%s, file=%s", bucketName, objectName, file)
	}
	if _, err := c.client.FPutObject(bucketName, objectName, file, minio.PutObjectOptions{ContentType: contentType}); err != nil {
		return "", err
	}

	url, err := c.GetFileUrl(bucketName, objectName)
	if err != nil {
		return "", errors.Wrapf(err, "get url")
	}
	return url, nil
}

func (c *MinioClient) DownloadFile(bucketName, objectName string) ([]byte, error) {
	obj, err := c.client.GetObject(bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	defer func() {
		if obj != nil {
			obj.Close()
		}
	}()

	data, err := ioutil.ReadAll(obj)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *MinioClient) GetFileUrl(bucketName, objectName string) (string, error) {
	info, err := c.client.StatObject(bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return "", errors.Wrapf(err, "get bk=%s file=%s url %v", bucketName, objectName, info)
	}
	if info.Err != nil {
		return "", errors.Wrapf(err, "get bk=%s file=%s info", bucketName, objectName)
	}
	return strings.Join([]string{c.endpoint, bucketName, objectName}, "/"), nil
}

func (c *MinioClient) HealthCheck() error {
	if _, err := c.client.BucketExists("bucket"); err != nil {
		return err
	}
	return nil
}

func detectFileType(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil {
		return "", err
	}
	fmt.Println(http.DetectContentType(buffer[:n]))
	return http.DetectContentType(buffer[:n]), nil
}
