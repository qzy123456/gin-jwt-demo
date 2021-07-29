package xstr

import (
	"encoding/binary"
	"math"
	"strconv"
)

// translate value(int/uint/float/string/[]byte) to float64
func Ftval(value interface{}) (ft float64, err error) {
	if value == nil {
		return
	}

	switch value.(type) {
	case float64:
		ft = value.(float64)
	case float32:
		ft = float64(value.(float32))
	case int:
		ft = float64(value.(int))
	case uint:
		ft = float64(value.(uint))
	case int8:
		ft = float64(value.(int8))
	case uint8:
		ft = float64(value.(uint8))
	case int16:
		ft = float64(value.(int16))
	case uint16:
		ft = float64(value.(uint16))
	case int32:
		ft = float64(value.(int32))
	case uint32:
		ft = float64(value.(uint32))
	case int64:
		ft = float64(value.(int64))
	case uint64:
		ft = float64(value.(uint64))
	case string:
		ft, err = strconv.ParseFloat(value.(string), 64)
	case []byte:
		bits := binary.LittleEndian.Uint64(value.([]byte))
		ft = math.Float64frombits(bits)
	}
	return
}

// round( v1 / v2, i), v1/v2(int/uint/float/string/[]byte), default: Inf is 0
func FtDivideRound(v1, v2 interface{}, i int) (res float64, err error) {
	var f1, f2 float64
	f1, err = Ftval(v1)
	if err != nil {
		return
	}
	f2, err = Ftval(v2)
	if err != nil {
		return
	}

	if f1 == 0 || f2 == 0 {
		return
	}

	res = math.Trunc((f1/f2)*math.Pow10(i)) / math.Pow10(i)
	return
}