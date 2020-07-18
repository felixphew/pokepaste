package pokepaste

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/blowfish"
)

var cipher *blowfish.Cipher

func init() {
	var err error
	cipher, err = blowfish.NewCipher([]byte(os.Getenv("POKEPASTE_KEY")))
	if err != nil {
		log.Fatal(err)
	}
}

func encodeID(id uint64) string {
	src := make([]byte, 8)
	binary.BigEndian.PutUint64(src, id)

	dst := make([]byte, 8)
	cipher.Encrypt(dst, src)

	return hex.EncodeToString(dst)
}

func decodeID(str string) (id uint64, err error) {
	src, err := hex.DecodeString(str)
	if err != nil {
		return
	}

	dst := make([]byte, 8)
	cipher.Decrypt(dst, src)

	id = binary.BigEndian.Uint64(dst)
	return
}

func decodeOldID(str string) (id uint64, err error) {
	id, err = strconv.ParseUint(str, 10, 64)
	if err != nil {
		return
	}

	if id >= 256 {
		id = (id * 0x7FFFFFFF) % 0x100000000
	}
	return
}
