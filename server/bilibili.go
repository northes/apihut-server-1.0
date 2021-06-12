package server

import (
	"apihut-server/model"
	"errors"
	"math"
	"strings"
)

var (
	ErrBvID = errors.New("BV号非法")
)

func Av2Bv(id *model.Av2Bv) (*model.Av2Bv, error) {
	var rid model.Av2Bv
	if len(id.Bid) != 0 {
		// 验证
		nbid := strings.Trim(id.Bid, "BV")
		if len(nbid) != 10 {
			return nil, ErrBvID
		}
		rid.Aid = bv2av(id.Bid)
		rid.Bid = id.Bid
	} else {
		rid.Bid = av2bv(id.Aid)
		rid.Aid = id.Aid
	}
	return &rid, nil
}

var table = "fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF"
var tr = map[string]int{}
var s = []int{11, 10, 3, 8, 4, 6}
var xor = 177451812
var add = 8728348608

func init() {
	tableByte := []byte(table)
	for i := 0; i < 58; i++ {
		tr[string(tableByte[i])] = i
	}
}

// bv转av
func bv2av(bv string) int {
	var r int
	arr := []rune(bv)

	for i := 0; i < 6; i++ {
		r += tr[string(arr[s[i]])] * int(math.Pow(float64(58), float64(i)))
	}
	return (r - add) ^ xor
}

// av转bv
func av2bv(av int) string {
	x := (av ^ xor) + add
	r := []string{"B", "V", "1", " ", " ", "4", " ", "1", " ", "7", " ", " "}
	for i := 0; i < 6; i++ {
		r[s[i]] = string(table[int(math.Floor(float64(x/int(math.Pow(float64(58), float64(i))))))%58])
	}
	var result string
	for i := 0; i < 12; i++ {
		result += r[i]
	}
	return result
}
