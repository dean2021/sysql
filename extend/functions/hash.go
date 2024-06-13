package functions

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"math"
	"os"
	"path/filepath"
)

func Md5File(path string) string {
	return hashFile(path, md5.New())
}

func Sha1File(path string) string {
	return hashFile(path, sha1.New())
}

func Sha256File(path string) string {
	return hashFile(path, sha256.New())
}

func hashFile(path string, hash hash.Hash) string {
	path, err := filepath.EvalSymlinks(path)
	if err != nil {
		return ""
	}
	const fileChunk = 1024 * 32
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return ""
	}
	if info.IsDir() {
		return ""
	}
	fileSize := info.Size()
	blocks := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	for i := uint64(0); i < blocks; i++ {
		blockSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		buf := make([]byte, blockSize)
		_, err = f.Read(buf)
		if err != nil {
			return ""
		}
		_, err = io.WriteString(hash, string(buf))
		if err != nil {
			return ""
		}
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Md5(str interface{}) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%v", str)))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1(str interface{}) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%v", str)))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256(str interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", str)))
	return hex.EncodeToString(h.Sum(nil))
}
