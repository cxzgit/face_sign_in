package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var (
	apiKey    = "BVMzzN97sykvpiz26prgT0oKK9Fc4EAAqHQdhLPJeJJcBMEGhdR3IYpO"
	secretKey = "K9Fc4EAAqHQdhLPJeJJcBMEGhdR3IYpO"
)

// getAccessToken 获取百度云API访问令牌
func getAccessToken() (string, error) {
	tokenURL := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", apiKey, secretKey)
	resp, err := http.Post(tokenURL, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if accessToken, ok := data["access_token"].(string); ok {
		return accessToken, nil
	}

	return "", errors.New("failed to get access token")
}

// FaceVerify 人脸比对/验证
// 这里简化为单张人脸的活体检测与质量判断，实际业务中可能需要与库中底图比对
func FaceVerify(imageBase64 string) (bool, error) {
	accessToken, err := getAccessToken()
	if err != nil {
		return false, err
	}

	apiURL := "https://aip.baidubce.com/rest/2.0/face/v3/faceverify"
	params := url.Values{}
	params.Add("access_token", accessToken)

	// Base64解码以检查图片大小等（可选）
	_, err = base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return false, errors.New("invalid base64 image")
	}

	// 构造请求体
	requestBody := fmt.Sprintf(`[{"image": "%s", "image_type": "BASE64", "face_field": "quality,liveness"}]`, imageBase64)

	resp, err := http.Post(apiURL+"?"+params.Encode(), "application/json", bytes.NewReader([]byte(requestBody)))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	// 解析百度云返回结果
	if errorCode, ok := result["error_code"].(float64); ok && errorCode != 0 {
		errorMsg := result["error_msg"].(string)
		return false, fmt.Errorf("face verification failed: %s", errorMsg)
	}

	// 简单校验：这里只检查是否成功检测到人脸
	if result["result"] != nil {
		return true, nil
	}

	return false, errors.New("no face detected or verification failed")
}

// init 检查环境变量中是否设置了百度云的Key
func init() {
	if os.Getenv("BAIDU_API_KEY") != "" {
		apiKey = os.Getenv("BAIDU_API_KEY")
	}
	if os.Getenv("BAIDU_SECRET_KEY") != "" {
		secretKey = os.Getenv("BAIDU_SECRET_KEY")
	}
}
