package apis

// PrintHelloworld 打印 helloworld
func PrintHelloworld() ResponseBody {
	return ResponseBody{
		Retcode: 0,
		Message: "helloworld",
		Data:    map[string]interface{}{},
	}
}
