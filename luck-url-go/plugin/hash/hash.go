package hash

import (
	"context"
	"fmt"
	"github.com/luck-labs/luck-url/common/consts/url_const"
	"github.com/luck-labs/luck-url/plugin/id"
)

func EncodeHashStr() (string, int64) {
	ctx := context.Background()
	hd := NewData()
	hd.Salt = url_const.DefaultSalt
	hd.MinLength = 5
	id := id.GetNextId(ctx)
	fmt.Println(id)
	h, _ := NewWithData(hd)
	e, _ := h.EncodeInt64([]int64{id})
	return e, id
}

func DecodeHashStr(str string) int64 {
	hd := NewData()
	hd.Salt = url_const.DefaultSalt
	hd.MinLength = 5
	h, _ := NewWithData(hd)
	d, _ := h.DecodeInt64WithError(str)
	fmt.Println(d)
	return d[0]
}
