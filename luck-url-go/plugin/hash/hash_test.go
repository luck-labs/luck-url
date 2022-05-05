package hash

import (
	"context"
	"fmt"
	"github.com/luck-labs/luck-url/plugin/id"
	"testing"
)

/**
 * @brief 生成hashID
 */
func TestGetId(t *testing.T) {
	hd := NewData()
	hd.Salt = ""
	hd.MinLength = 6
	h, _ := NewWithData(hd)
	e, _ := h.Encode([]int{45, 434, 1313, 99})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
}

/**
 * @brief 生成hashID
 */
func TestGetIdV2(t *testing.T) {
	ctx := context.Background()
	id.Init()
	hd := NewData()
	hd.Salt = ""
	hd.MinLength = 6
	for i := 0; i < 1000; i++ {
		id := id.GetNextId(ctx)
		fmt.Println(id)
		h, _ := NewWithData(hd)
		e, _ := h.EncodeInt64([]int64{id})
		fmt.Println(e)
		d, _ := h.DecodeWithError(e)
		fmt.Println(d)
	}
}
