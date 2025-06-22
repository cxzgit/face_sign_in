package globals

type AppCode int

// 自定义状态码：StatusOK = 2000，区别于 http.StatusOK = 200
const (
	StatusOK                  AppCode = 2000 // 成功
	StatusBadRequest          AppCode = 4000 // 请求语法错误或无效参数
	StatusInternalServerError AppCode = 5000 // 服务器内部错误
	StatusUnauthorized        AppCode = 4001 // RFC 9110, 15.5.2
)
