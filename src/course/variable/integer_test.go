package variable

import (
	"fmt"
	"math"
	"testing"
)

func TestTransType(t *testing.T) {
	// 类型转换只能在定义正确的情况下转换成功
	// 从一个取值范围较小的类型转换到一个取值范围较大的类型
	// 只有底层类型的变量之间可以进行相互转换
	a := 5.0
	b := int(a)
	fmt.Println(a, b)

	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	// 初始化一个32位整型值
	var c int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", c, c)
	// 将a变量数值转换为十六进制, 发生数值截断
	d := int16(c)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", d, d)
	// 将常量保存为float32类型
	var e float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(e))
}
