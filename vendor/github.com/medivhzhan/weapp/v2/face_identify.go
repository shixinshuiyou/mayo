package weapp

const (
	apiFaceIdentify = "/cityservice/face/identify/getinfo"
)

// FaceIdentifyResponse 人脸识别结果返回
type FaceIdentifyResponse struct {
	CommonError
	Result          int    `json:"identify_ret"`       // 认证结果
	Time            uint32 `json:"identify_time"`      // 认证时间
	Data            string `json:"validate_data"`      // 用户读的数字（如是读数字）
	OpenID          string `json:"openid"`             // 用户openid
	UserIDKey       string `json:"user_id_key"`        // 用于后台交户表示用户姓名、身份证的凭证
	FinishTime      uint32 `json:"finish_time"`        // 认证结束时间
	IDCardNumberMD5 string `json:"id_card_number_md5"` // 身份证号的md5（最后一位X为大写）
	NameUTF8MD5     string `json:"name_utf8_md5"`      // 姓名MD5
}

// FaceIdentify 获取人脸识别结果
//
// token 微信 access_token
// key 小程序 verify_result
func FaceIdentify(token, key string) (*FaceIdentifyResponse, error) {
	api := baseURL + apiFaceIdentify
	return faceIdentify(api, token, key)
}

func faceIdentify(api, token, key string) (*FaceIdentifyResponse, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"verify_result": key,
	}

	res := new(FaceIdentifyResponse)
	err = postJSON(api, params, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
