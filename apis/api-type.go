package apis

// struct 中的字段名要注意大小写，小写的字段名在 json 序列化时会被忽略
type ResponseBody struct {
	Retcode int                    `json:"retcode"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
