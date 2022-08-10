package library

var TextScenes = map[string]map[string]string{
	"1000": {
		"desc":  "帮助",
		"value": "about",
	},
	"1001": {
		"desc":  "支付协议",
		"value": "pay_agreement",
	},
}

var UploadScenes = map[string]map[string]string{
	"1000": {
		"desc":  "ppt封面",
		"value": "ppt_cover",
	},
	"1001": {
		"desc":  "ppt详情",
		"value": "ppt_detail",
	},
	"1003": {
		"desc":  "ppt文件",
		"value": "ppt_file",
	},
	"1004": {
		"desc":  "头像",
		"value": "head_img",
	},
	"1005": {
		"desc":  "客服二维码",
		"value": "customer_service_wx",
	},
}

var scenes = make(map[string]map[string]string)

func init() {
	scenes["text"] = parseScene(TextScenes)
	scenes["upload"] = parseScene(UploadScenes)
}

func parseScene(scene map[string]map[string]string) map[string]string {
	res := make(map[string]string, len(scene))
	for idx, v := range scene {
		res[v["value"]] = idx
	}
	return res
}

// key => value
func GetScenes(mark string) map[string]string {
	return scenes[mark]
}

// key => value
func GetSceneAll() map[string]map[string]string {
	return scenes
}

// key => value
func GetUploadSceneInfo(scene string) map[string]string {
	return UploadScenes[scene]
}

func GetTextSceneInfo(scene string) map[string]string {
	return TextScenes[scene]
}
