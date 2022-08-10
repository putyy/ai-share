package handleVideo

import (
	"encoding/json"
	"errors"
	"github.com/putyy/ai-share/app/library"
	"net/http"
	"net/url"
)

const hs_url = "https://share.huoshan.com/api/item/info?item_id="

//获取真实请求地址的path
func getHSRealityUrl(rUrl, ua string) (url.Values, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", rUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Request.URL.Query(), nil
}

type HSRes struct {
	Data *struct {
		Item_info *struct {
			Cover string
			Url   string
		}
	}
}

//火山
func HuoShan(rUrl, ua string) (string, error) {
	query, err := getHSRealityUrl(rUrl, ua)
	if err != nil {
		return "", err
	}
	itemId := query["item_id"]
	if itemId == nil {
		return "", errors.New("无效地址")
	}
	body, err := library.HttpGet(hs_url+itemId[0], ua)
	if err != nil {
		return "", err
	}
	var res HSRes
	_ = json.Unmarshal(body, &res)
	return res.Data.Item_info.Url, nil
}
