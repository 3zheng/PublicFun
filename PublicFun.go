package PublicFun

import (
	"math/rand"
)

//公共函数

/******************************************************************
几个关键的ascii值 48='0', 57='9', 65='A', 90='Z', 97='a', 122='z'
函数作用：创建随机的ascii字符串(需要自己预设随机种子)
参数说明：scope	创建字符串选择的字符范围
		length	创建的字符串的长度

例子：见PublicFun_test.go
******************************************************************/
func CreateRandAsciiString(scope []byte, length int) string {
	if length < 0 {
		return ""
	}

	scopeLen := len(scope)
	out := make([]byte, length)
	for i := 0; i < length; i++ {
		num := rand.Intn(scopeLen)
		out[i] = scope[num]
	}

	return string(out)
}
