package main

import (
 "fmt"
 "strconv"
)

/**
 将 string 转换为其他基本类型
 */
func main()  {
	str := "hello"
	myBool, _ := strconv.ParseBool(str)
    fmt.Printf("mybool typs is %T, value is %v \n", myBool, myBool)

	str2 := "42"
	num, _ := strconv.ParseInt(str2, 10, 64)
    fmt.Printf("str2 type is %T, and it value is %q \n",str2, str2 )
    fmt.Printf("num type is %T, and it value is %v \n",num, num )
}