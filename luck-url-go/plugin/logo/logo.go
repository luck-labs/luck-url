package logo

import (
	"fmt"
	"io/ioutil"
)

/**
 * @brief ASCII LOGO https://patorjk.com/ ANSI Shadow
 */
func Init() {
	fileName := "plugin/logo/logo.txt"
	logoStrByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	logoStr := string(logoStrByte)
	fmt.Println(logoStr)
}
