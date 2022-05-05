package utils

import (
	"fmt"
	"os"
)

func PrintAndDie(err error) {
	fmt.Printf("init failed, err:%s \n", err)
	os.Exit(-1)
}
