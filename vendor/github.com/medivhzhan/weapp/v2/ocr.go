package weapp

const (
	apiBankcard        = "/cv/ocr/bankcard"
	apiVehicleLicense  = "/cv/ocr/driving"
	apiDrivingLicense  = "/cv/ocr/drivinglicense"
	apiIDCard          = "/cv/ocr/idcard"
	apiBusinessLicense = "/cv/ocr/bizlicense"
	apiPrintedText     = "/cv/ocr/comm"
)

// RecognizeMode 图片识别模式
type RecognizeMode = string

// 所有图片识别模式
const (
	RecognizeModePhoto RecognizeMode = "photo" // 拍照模式
	RecognizeModeScan  RecognizeMode = "scan"  // 扫描模式
)

// BankCardResponse 识别银行卡返回数据
type BankCardResponse struct {
	CommonError
	Number string `json:"number"` // 银行卡号
}

// BankCardByURL 通过URL识别银行卡
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// url 要检测的图片 url，传这个则不用传 img 参数。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func BankCardByURL(token, cardURL string, mode RecognizeMode) (*BankCardResponse, error) {
	api := baseURL + apiBankcard
	return bankCardByURL(api, token, cardURL, mode)
}

func bankCardByURL(api, token, cardURL string, mode RecognizeMode) (*BankCardResponse, error) {
	res := new(BankCardResponse)
	err := ocrByURL(api, token, cardURL, mode, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// BankCard 通过文件识别银行卡
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// img form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传递 img_url。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func BankCard(token, filename string, mode RecognizeMode) (*BankCardResponse, error) {
	api := baseURL + apiBankcard
	return bankCard(api, token, filename, mode)
}

func bankCard(api, token, filename string, mode RecognizeMode) (*BankCardResponse, error) {
	res := new(BankCardResponse)
	err := ocrByFile(api, token, filename, mode, res)
	if err != nil {
		return nil, err
	}

	return res, err
}

// CardType 卡片方向
type CardType = string

// 所有卡片方向
const (
	CardTypeFront = "Front" // 正面
	CardTypeBack  = "Back"  // 背面
)

// CardResponse 识别卡片返回数据
type CardResponse struct {
	CommonError
	Type      CardType `json:"type"`       // 正面或背面，Front / Back
	ValidDate string   `json:"valid_date"` // 有效期
}

// DrivingLicenseResponse 识别行驶证返回数据
type DrivingLicenseResponse struct {
	CommonError
	IDNum        string `json:"id_num"`        // 证号
	Name         string `json:"name"`          // 姓名
	Nationality  string `json:"nationality"`   // 国家
	Sex          string `json:"sex"`           // 性别
	Address      string `json:"address"`       // 地址
	BirthDate    string `json:"birth_date"`    // 出生日期
	IssueDate    string `json:"issue_date"`    // 初次领证日期
	CarClass     string `json:"car_class"`     // 准驾车型
	ValidFrom    string `json:"valid_from"`    // 有效期限起始日
	ValidTo      string `json:"valid_to"`      // 有效期限终止日
	OfficialSeal string `json:"official_seal"` // 印章文构
}

// DriverLicenseByURL 通过URL识别行驶证
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// url 要检测的图片 url，传这个则不用传 img 参数。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func DriverLicenseByURL(token, licenseURL string) (*DrivingLicenseResponse, error) {
	api := baseURL + apiDrivingLicense
	return driverLicenseByURL(api, token, licenseURL)
}

func driverLicenseByURL(api, token, licenseURL string) (*DrivingLicenseResponse, error) {
	res := new(DrivingLicenseResponse)
	err := ocrByURL(api, token, licenseURL, "", res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// DriverLicense 通过文件识别行驶证
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// img form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传递 img_url。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func DriverLicense(token, filename string) (*DrivingLicenseResponse, error) {
	api := baseURL + apiDrivingLicense
	return driverLicense(api, token, filename)
}

func driverLicense(api, token, filename string) (*DrivingLicenseResponse, error) {
	res := new(DrivingLicenseResponse)
	err := ocrByFile(api, token, filename, "", res)
	if err != nil {
		return nil, err
	}

	return res, err
}

// IDCardResponse 识别身份证返回数据
type IDCardResponse = CardResponse

// IDCardByURL 通过URL识别身份证
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// url 要检测的图片 url，传这个则不用传 img 参数。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func IDCardByURL(token, cardURL string, mode RecognizeMode) (*IDCardResponse, error) {
	api := baseURL + apiIDCard
	return idCardByURL(api, token, cardURL, mode)
}

func idCardByURL(api, token, cardURL string, mode RecognizeMode) (*IDCardResponse, error) {
	res := new(IDCardResponse)
	err := ocrByURL(api, token, cardURL, mode, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// IDCard 通过文件识别身份证
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// img form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传递 img_url。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func IDCard(token, filename string, mode RecognizeMode) (*IDCardResponse, error) {
	api := baseURL + apiIDCard
	return idCard(api, token, filename, mode)
}

func idCard(api, token, filename string, mode RecognizeMode) (*IDCardResponse, error) {
	res := new(IDCardResponse)
	err := ocrByFile(api, token, filename, mode, res)
	if err != nil {
		return nil, err
	}

	return res, err
}

// VehicleLicenseResponse 识别卡片返回数据
type VehicleLicenseResponse struct {
	CommonError
	VehicleType                string `json:"vehicle_type"`
	Owner                      string `json:"owner"`
	Addr                       string `json:"addr"`
	UseCharacter               string `json:"use_character"`
	Model                      string `json:"model"`
	Vin                        string `json:"vin"`
	EngineNum                  string `json:"engine_num"`
	RegisterDate               string `json:"register_date"`
	IssueDate                  string `json:"issue_date"`
	PlateNumB                  string `json:"plate_num_b"`
	Record                     string `json:"record"`
	PassengersNum              string `json:"passengers_num"`
	TotalQuality               string `json:"total_quality"`
	TotalprepareQualityQuality string `json:"totalprepare_quality_quality"`
}

// VehicleLicenseByURL 行驶证 OCR 识别
func VehicleLicenseByURL(token, cardURL string, mode RecognizeMode) (*VehicleLicenseResponse, error) {
	api := baseURL + apiVehicleLicense
	return vehicleLicenseByURL(api, token, cardURL, mode)
}

func vehicleLicenseByURL(api, token, cardURL string, mode RecognizeMode) (*VehicleLicenseResponse, error) {
	res := new(VehicleLicenseResponse)
	err := ocrByURL(api, token, cardURL, mode, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// VehicleLicense 通过文件识别行驶证
func VehicleLicense(token, filename string, mode RecognizeMode) (*VehicleLicenseResponse, error) {
	api := baseURL + apiVehicleLicense
	return vehicleLicense(api, token, filename, mode)
}

func vehicleLicense(api, token, filename string, mode RecognizeMode) (*VehicleLicenseResponse, error) {
	res := new(VehicleLicenseResponse)
	err := ocrByFile(api, token, filename, mode, res)
	if err != nil {
		return nil, err
	}

	return res, err
}

// LicensePoint 证件点
type LicensePoint struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}

// LicensePosition 证件位置
type LicensePosition struct {
	LeftTop     LicensePoint `json:"left_top"`
	RightTop    LicensePoint `json:"right_top"`
	RightBottom LicensePoint `json:"right_bottom"`
	LeftBottom  LicensePoint `json:"left_bottom"`
}

// BusinessLicenseResponse 营业执照 OCR 识别返回数据
type BusinessLicenseResponse struct {
	CommonError
	RegNum              string `json:"reg_num"`              //	注册号
	Serial              string `json:"serial"`               //	编号
	LegalRepresentative string `json:"legal_representative"` //	法定代表人姓名
	EnterpriseName      string `json:"enterprise_name"`      //	企业名称
	TypeOfOrganization  string `json:"type_of_organization"` //	组成形式
	Address             string `json:"address"`              //	经营场所/企业住所
	TypeOfEnterprise    string `json:"type_of_enterprise"`   //	公司类型
	BusinessScope       string `json:"business_scope"`       //	经营范围
	RegisteredCapital   string `json:"registered_capital"`   //	注册资本
	PaidInCapital       string `json:"paid_in_capital"`      //	实收资本
	ValidPeriod         string `json:"valid_period"`         //	营业期限
	RegisteredDate      string `json:"registered_date"`      //	注册日期/成立日期
	CertPosition        struct {
		Position LicensePosition `json:"pos"`
	} `json:"cert_position"` //	营业执照位置
	ImgSize LicensePoint `json:"img_size"` //	图片大小
}

// BusinessLicenseByURL 通过链接进行营业执照 OCR 识别
func BusinessLicenseByURL(token, cardURL string) (*BusinessLicenseResponse, error) {
	api := baseURL + apiBusinessLicense
	return businessLicenseByURL(api, token, cardURL)
}

func businessLicenseByURL(api, token, cardURL string) (*BusinessLicenseResponse, error) {
	res := new(BusinessLicenseResponse)
	err := ocrByURL(api, token, cardURL, "", res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// BusinessLicense 通过文件进行营业执照 OCR 识别
func BusinessLicense(token, filename string) (*BusinessLicenseResponse, error) {
	api := baseURL + apiBusinessLicense
	return businessLicense(api, token, filename)
}

func businessLicense(api, token, filename string) (*BusinessLicenseResponse, error) {
	res := new(BusinessLicenseResponse)
	err := ocrByFile(api, token, filename, "", res)
	if err != nil {
		return nil, err
	}

	return res, err
}

// PrintedTextResponse 通用印刷体 OCR 识别返回数据
type PrintedTextResponse struct {
	CommonError
	Items []struct {
		Text     string          `json:"text"`
		Position LicensePosition `json:"pos"`
	} `json:"items"` //	识别结果
	ImgSize LicensePoint `json:"img_size"` //	图片大小
}

// PrintedTextByURL 通过链接进行通用印刷体 OCR 识别
func PrintedTextByURL(token, cardURL string) (*PrintedTextResponse, error) {
	api := baseURL + apiPrintedText
	return printedTextByURL(api, token, cardURL)
}

func printedTextByURL(api, token, cardURL string) (*PrintedTextResponse, error) {
	res := new(PrintedTextResponse)
	err := ocrByURL(api, token, cardURL, "", res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// PrintedText 通过文件进行通用印刷体 OCR 识别
func PrintedText(token, filename string) (*PrintedTextResponse, error) {
	api := baseURL + apiPrintedText
	return printedText(api, token, filename)
}

func printedText(api, token, filename string) (*PrintedTextResponse, error) {
	res := new(PrintedTextResponse)
	err := ocrByFile(api, token, filename, "", res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func ocrByFile(api, token, filename string, mode RecognizeMode, response interface{}) error {
	queries := requestQueries{
		"access_token": token,
		"type":         mode,
	}

	url, err := encodeURL(api, queries)
	if err != nil {
		return err
	}

	if err := postFormByFile(url, "img", filename, response); err != nil {
		return err
	}

	return nil
}

func ocrByURL(api, token, cardURL string, mode RecognizeMode, response interface{}) error {
	queries := requestQueries{
		"access_token": token,
		"img_url":      cardURL,
	}

	if mode != "" {
		queries["type"] = mode
	}

	url, err := encodeURL(api, queries)
	if err != nil {
		return err
	}

	if err := postJSON(url, nil, response); err != nil {
		return err
	}

	return nil
}
