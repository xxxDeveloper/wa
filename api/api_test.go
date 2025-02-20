// 版权 @2022 凹语言 作者。保留所有权利。

package api_test

import (
	"fmt"

	"github.com/wa-lang/wa/api"
)

func ExampleFormatCode() {
	s, _ := api.FormatCode("hello.wa", "fn add(a i32, b i32) i32 {return a+b}")
	fmt.Println(s)

	// Output:
	// fn add(a: i32, b: i32) => i32 { return a + b }
}
