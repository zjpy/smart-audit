package contract

import (
	"crypto/rand"
	"testing"
)

func TestArrayMapConvert(t *testing.T) {
	var args []string
	for i := 0; i < 10; i++ {
		args = append(args, randomArg())
	}

	// 参数经过两次转换后得到的数组应该与原始的数组完全一致
	args2 := mapToArray(arrayToMap(args))

	if len(args) != len(args2) {
		t.Error("数组长度不一致")
	}
	for i := range args {
		if args[i] != args2[i] {
			t.Errorf("第%d个值出现错误，初始值为: %s，实际值为：%s",
				i, args[i], args2[i])
		}
	}
}

func randomArg() string {
	buf := make([]byte, 20)
	rand.Read(buf)
	return string(buf)
}
