package render


type RenderRet struct {
	Code int
	Msg string
}
func RenderRetData(code int, msg string) (retData RenderRet) {
	retData = RenderRet{code, msg}
	return
}

