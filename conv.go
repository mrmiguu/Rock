package rock

import (
	"encoding/binary"
	"strconv"
)

func int2bytes(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func bytes2int(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func bool2bytes(b bool) []byte {
	return strconv.AppendBool(nil, b)
}

func bytes2bool(b []byte) bool {
	tf, _ := strconv.ParseBool(string(b))
	return tf
}
