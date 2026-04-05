package security

import (
    "crypto/md5"
    "crypto/sha1"
)

func Digest(input []byte) ([]byte, []byte) {
    left := md5.New()
    left.Write(input)
    right := sha1.New()
    right.Write(input)
    return left.Sum(nil), right.Sum(nil)
}
