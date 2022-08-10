package library

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/putyy/ai-share/config"
	"io/ioutil"
	"mvdan.cc/xurls/v2"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetImgUrl(src string) string {
	if src == "" {
		return config.App.DefaultImg
	}
	if strings.HasPrefix(src, "http") {
		return src
	}
	return config.Qiniu.ImageDomain + "/" + src
}

func GetFileUrl(src string) string {
	if src == "" {
		return config.App.DefaultImg
	}
	if strings.HasPrefix(src, "http") {
		return src
	}
	return config.Qiniu.MediaDomain + "/" + src
}

func GetSaveUrl(src string) string {
	if src == "" {
		return ""
	}

	p, err := url.Parse(src)
	if err != nil {
		return ""
	}

	return strings.TrimLeft(p.Path, "/")
}

func En100(i int) int {
	return i * 100
}

func De100(i int) int {
	return i / 100
}

// get请求
func HttpGet(url string, ua string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

// 提取文字中的url
func ParseGetUrl(url string) (string, error) {
	rxRelaxed := xurls.Strict()
	src := rxRelaxed.FindAllString(url, -1)
	if len(src) == 0 {
		return "", errors.New("无效地址")
	}
	return src[0], nil
}

func MakeFileName(scene string) (string, error) {
	sceneInfo := GetUploadSceneInfo(scene)
	v, ok := sceneInfo["value"]
	if ok == false {
		return "", errors.New("无效scene配置")
	}
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	s := time.Now().Format("20060102") + fmt.Sprintf("%x", b)
	return "upload/" + v + "/" + s, nil
}

func MakeOrderNo(uid int) (string, error) {
	b := make([]byte, 3)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	suid := fmt.Sprintf("%07s%s", strconv.Itoa(uid), "")[:7]
	s := time.Now().Format("20060102") + fmt.Sprintf("%x", b) + suid
	return s, nil
}
