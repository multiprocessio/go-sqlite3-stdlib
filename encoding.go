package stdlib

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

func ext_md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func ext_sha1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

func ext_sha256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func ext_sha512(s string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
}

func sha3_256(s string) string {
	return fmt.Sprintf("%x", sha3.Sum256([]byte(s)))
}

func sha3_512(s string) string {
	return fmt.Sprintf("%x", sha3.Sum512([]byte(s)))
}

func blake2b_256(s string) string {
	return fmt.Sprintf("%x", blake2b.Sum256([]byte(s)))
}

func blake2b_512(s string) string {
	return fmt.Sprintf("%x", blake2b.Sum512([]byte(s)))
}

func toBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func fromBase64(s string) string {
	d, _ := base64.StdEncoding.DecodeString(s)
	return string(d)
}

func toBase32(s string) string {
	return base32.StdEncoding.EncodeToString([]byte(s))
}

func fromBase32(s string) string {
	d, _ := base32.StdEncoding.DecodeString(s)
	return string(d)
}

var encodingFunctions = map[string]any{
	"base64":      stringy1string(toBase64),
	"from_base64": stringy1string(fromBase64),
	"base32":      stringy1string(toBase32),
	"from_base32": stringy1string(fromBase32),

	"md5":         stringy1string(ext_md5),
	"sha1":        stringy1string(ext_sha1),
	"sha256":      stringy1string(ext_sha256),
	"sha512":      stringy1string(ext_sha512),
	"sha3_256":    stringy1string(sha3_256),
	"sha3_512":    stringy1string(sha3_512),
	"blake2b_256": stringy1string(blake2b_256),
	"blake2b_512": stringy1string(blake2b_512),
}
