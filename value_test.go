package m2rcache

import "testing"

func TestArrayFullValue(t *testing.T) {
	var test = make([]interface{}, 0)
	var result = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fullValue(&test, nil, result)

	t.Log(test)
}
func TestMapFullValue(t *testing.T) {
	var test map[string]interface{}
	var result = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var fields = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	testFullValue(&test, fields, result, t)

	//	t.Log(test)
}

type testS1 struct {
	A int
	B int8
	C int16
	D int32
	E int64
	F uint
	G uint8
	H uint16
	I uint32
	J uint64
	K float32
	M float64
	N string
}

func TestStructFullValue(t *testing.T) {
	var test map[string]interface{}
	var ts1 testS1
	var r = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	var r1 = []interface{}{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"}
	var rI16 = []interface{}{int16(1), int16(2), int16(3), int16(4), int16(5), int16(6), int16(7), int16(8), int16(9), int16(10), int16(11), int16(12), int16(13)}
	var rI8 = []interface{}{int8(1), int8(2), int8(3), int8(4), int8(5), int8(6), int8(7), int8(8), int8(9), int8(10), int8(11), int8(12), int8(13)}
	var rI32 = []interface{}{int32(1), int32(2), int32(3), int32(4), int32(5), int32(6), int32(7), int32(8), int32(9), int32(10), int32(11), int32(12), int32(13)}
	var rI64 = []interface{}{int64(1), int64(2), int64(3), int64(4), int64(5), int64(6), int64(7), int64(8), int64(9), int64(10), int64(11), int64(12), int64(13)}
	var rU = []interface{}{uint(1), uint(2), uint(3), uint(4), uint(5), uint(6), uint(7), uint(8), uint(9), uint(10), uint(11), uint(12), uint(13)}
	var rU8 = []interface{}{uint8(1), uint8(2), uint8(3), uint8(4), uint8(5), uint8(6), uint8(7), uint8(8), uint8(9), uint8(10), uint8(11), uint8(12), uint8(13)}
	var rU16 = []interface{}{uint16(1), uint16(2), uint16(3), uint16(4), uint16(5), uint16(6), uint16(7), uint16(8), uint16(9), uint16(10), uint16(11), uint16(12), uint16(13)}
	var rU32 = []interface{}{uint32(1), uint32(2), uint32(3), uint32(4), uint32(5), uint32(6), uint32(7), uint32(8), uint32(9), uint32(10), uint32(11), uint32(12), uint32(13)}
	var rU64 = []interface{}{uint64(1), uint64(2), uint64(3), uint64(4), uint64(5), uint64(6), uint64(7), uint64(8), uint64(9), uint64(10), uint64(11), uint64(12), uint64(13)}
	var rF32 = []interface{}{float32(1), float32(2), float32(3), float32(4), float32(5), float32(6), float32(7), float32(8), float32(9), float32(10), float32(11), float32(12), float32(13)}
	var rF64 = []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10), float64(11), float64(12), float64(13)}
	var fields = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "M", "N"}
	var rArr = [][]interface{}{r, r1, rI8, rI16, rI32, rI64, rU, rU8, rU16, rU32, rU64, rF32, rF64}
	for _, v := range rArr {
		testFullValue(&test, fields, v, t)
		testFullValue(&ts1, fields, v, t)
	}

}
func testFullValue(dst interface{}, fields []string, result []interface{}, t *testing.T) {
	t.Log("填充源：", result)
	t.Log("填充字段：", fields)
	fullValue(dst, fields, result)

	t.Log("填充结果：", dst)
}
