package controllers

//ControllerInterface 控制器interface
type ControllerInterface interface {
	SetData(data interface{})
	ServeJSON(encoding ...bool)
}

//SetErrorReturn 设置错误返回
func SetErrorReturn(c ControllerInterface, errCode int64, err error) {
	c.SetData(map[string]interface{}{
		"status": errCode,
		"msg":    err.Error(),
	},
	)
	c.ServeJSON()
}
