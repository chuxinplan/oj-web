package store

import (
	"sync"

	minio "github.com/minio/minio-go"
	"github.com/open-fightcoder/oj-web/common/g"
)

var MinioClient *minio.Client
var onceExec sync.Once

func InitMinio() {
	onceExec.Do(func() {
		cfg := g.Conf()
		var err error
		MinioClient, err = minio.New(cfg.Minio.Endpoint, cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, cfg.Minio.Secure)
		if err != nil {
			//write log
		}

	})
}
