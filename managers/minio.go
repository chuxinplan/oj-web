package managers

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"errors"

	"time"

	"github.com/minio/minio-go"
	"github.com/open-fightcoder/oj-web/common/g"
	. "github.com/open-fightcoder/oj-web/common/store"
)

func GetSaveCode(name string) (string, error) {
	cfg := g.Conf()
	resp, err := http.Get("http://" + cfg.Minio.Endpoint + "/" + cfg.Minio.SaveCodeBucket + "/" + name)
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

func GetCode(name string) (string, error) {
	cfg := g.Conf()
	resp, err := http.Get("http://" + cfg.Minio.Endpoint + "/" + cfg.Minio.CodeBucket + "/" + name)
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
	str := strconv.FormatInt(time.Now().Unix(), 10)
	return str + ".txt"
}

func SaveCode(code string) (string, error) {
	cfg := g.Conf()
	str := GetCodeName()
	_, err := MinioClient.PutObject(cfg.Minio.SaveCodeBucket, str, strings.NewReader(code), -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return "", errors.New("存储失败")
	}
	return str, nil
}

func GetSubmitCodeName(userId int64, language string) string {
	var languageStr string
	switch language {
	case "java":
		languageStr = ".java"
		break
	case "c":
		languageStr = ".c"
		break
	case "c++":
		languageStr = ".cpp"
		break
	case "go":
		languageStr = ".go"
		break
	case "python":
		languageStr = ".py"
		break
	default:
		languageStr = ".txt"
	}
	return strconv.FormatInt(userId, 10) + "_" + strconv.FormatInt(time.Now().Unix(), 10) + languageStr
}

func SaveSubmitCode(code string, userId int64, language string) (string, error) {
	cfg := g.Conf()
	str := GetSubmitCodeName(userId, language)
	_, err := MinioClient.PutObject(cfg.Minio.CodeBucket, str, strings.NewReader(code), -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return "", errors.New(err.Error())
	}
	return str, nil
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

//func UpdateCode(path string, code string) error {
//	cfg := g.Conf()
//	_, err := MinioClient.PutObject(cfg.Minio.CodeBucket, path, strings.NewReader(code), -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
//	if err != nil {
//		return errors.New("更新失败")
//	}
//	return nil
//}

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
