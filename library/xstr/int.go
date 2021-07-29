package xstr

import (
	"encoding/binary"
	"math"
	"strconv"
)

func Int64val(value interface{}) (i64 int64, err error) {
	if value == nil {
		return
	}

	switch value.(type) {
	case float64:
		i64 = int64(value.(float64))
	case float32:
		i64 = int64(value.(float32))
	case int:
		i64 = int64(value.(int))
	case uint:
		i64 = int64(value.(uint))
	case int8:
		i64 = int64(value.(int8))
	case uint8:
		i64 = int64(value.(uint8))
	case int16:
		i64 = int64(value.(int16))
	case uint16:
		i64 = int64(value.(uint16))
	case int32:
		i64 = int64(value.(int32))
	case uint32:
		i64 = int64(value.(uint32))
	case int64:
		i64 = value.(int64)
	case uint64:
		i64 = int64(value.(uint64))
	case string:
		i64, err = strconv.ParseInt(value.(string), 10, 64)
	case []byte:
		bits := binary.LittleEndian.Uint64(value.([]byte))
		i64 = int64(math.Float64frombits(bits))
	}
	return
}
