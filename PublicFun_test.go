package PublicFun

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandString(t *testing.T) {
	rand.Seed(time.Now().Unix())
	scope := make([]byte, 26)
	//生成四位大写字母组成的字符串
	for i := 0; i < 26; i++ {
		scope[i] = byte(65 + i)
	}
	t.Log("字符串选择范围为:", string(scope))

	var str string
	for i := 0; i < 500000; i++ {
		str = CreateRandAsciiString(scope, 4)
	}
	t.Log("生成的字符串为：", str)
}
