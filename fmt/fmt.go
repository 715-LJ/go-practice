package main

import (
	"fmt"
)

func main() {

	a, b := "name", "os"
	//函数用于打印一行文本，并自动在最后添加一个换行符。它会将每个参数之间以空格分隔，并在最后添加一个换行符。如果不传入参数，则只打印一个空行
	fmt.Println("lijiaos" + " " + a + " " + b)
	//函数用于格式化输出，并根据指定的格式字符串将数据打印到标准输出。它类似于 C 语言中的 printf 函数。通过占位符 % 可以指定要打印的数据类型和格式。
	//例如，%s：用于字符串类型的数据。
	//%d：用于十进制整数类型的数据。
	//%f：用于浮点数类型的数据。
	//%t：用于布尔类型的数据。
	//%v：用于通用类型的数据，会自动选择合适的格式进行输出。
	//可以在格式字符串中使用多个占位符，并通过逗号分隔的方式传入相应的参数
	fmt.Printf("%s %v \n", a, b)
}
