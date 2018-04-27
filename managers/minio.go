package managers

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"errors"

	"github.com/minio/minio-go"
	"github.com/open-fightcoder/oj-web/common/g"
	. "github.com/open-fightcoder/oj-web/common/store"
)

func GetCode(name string) (string, error) {
	cfg := g.Conf()
	resp, err := http.Get("http://xupt1.fightcoder.com:9001/" + cfg.Minio.CodeBucket + "/" + name)
	if err != nil {
		return "", errors.New("获取失败")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("获取失败")
	}
	return string(body), nil
}

func GetNameByPath(path string) string {
	strs := strings.Split(path, "/")
	return strs[len(strs)-1]
}

func GetImgName(userId int64, picType string) string {
	str := strconv.FormatInt(userId, 10)
	return str + "." + picType
}

func GetCodeName() string {
	timestamp := time.Now().UnixNano() / 1000000

	str := strconv.FormatInt(timestamp, 10)
	return str + ".txt"
}

func SaveImg(reader io.Reader, userId int64, picType string) (string, error) {
	cfg := g.Conf()
	str := GetImgName(userId, picType)
	_, err := MinioClient.PutObject(cfg.Minio.ImgBucket, str, reader, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return "", errors.New("存储失败")
	}
	return str, nil
}

func SaveCode(code string) (string, error) {
	cfg := g.Conf()
	str := GetCodeName()
	_, err := MinioClient.PutObject(cfg.Minio.CodeBucket, str, strings.NewReader(code), -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return "", errors.New("存储失败")
	}
	return str, nil
}

func RemoveCode(name string) error {
	cfg := g.Conf()
	err := MinioClient.RemoveObject(cfg.Minio.CodeBucket, name)
	if err != nil {
		return errors.New("删除失败")
	}
	return nil
}

func Download(objectName, filePath string) error {
	cfg := g.Conf()
	err := MinioClient.FGetObject(cfg.Minio.CodeBucket, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return errors.New("下载失败")
	}
	return nil
}
