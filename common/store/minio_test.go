package store

import (
	"testing"

	"strings"

	"github.com/minio/minio-go"
)

func TestMinio(t *testing.T) {
	InitMinio()
	_, err := MinioClient.PutObject("ttt", "test.txt", strings.NewReader("aaaaaassssssdddd"), -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		panic(err)
	}
}
