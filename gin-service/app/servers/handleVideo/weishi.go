package handleVideo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const weishiUrl = "https://h5.weishi.qq.com/webapp/json/weishi/WSH5GetPlayPage"

func requestWeiShi(rUrl, ua string) ([]byte, error) {
	req, err := http.NewRequest("GET", rUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func WeiShi(rUrl, ua string) (string, error) {
	parse, err := url.Parse(rUrl)
	if err != nil {
		return "", err
	}
	ids := parse.Query()["id"]
	if len(ids) == 0 {
		return "", errors.New("无效地址")
	}
	id := ids[0]
	type Video struct {
		Data *struct {
			Feeds []*struct {
				Video_url string
			}
		}
	}
	body, err := requestWeiShi(weishiUrl+"?feedid="+id, ua)
	if err != nil {
		return "", err
	}
	var res Video
	_ = json.Unmarshal([]byte(body), &res)
	if len(res.Data.Feeds) == 0 {
		return "", errors.New("未找到有效视频")
	}
	return res.Data.Feeds[0].Video_url, nil
}
