package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"math"
	"os"
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
	const fileChunk = 8192 // we settle for 8KB
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()
	// calculate the file size
	info, err := file.Stat()
	if err != nil {
		return ""
	}
	fileSize := info.Size()
	blocks := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	for i := uint64(0); i < blocks; i++ {
		blockSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		buf := make([]byte, blockSize)
		_, err = file.Read(buf)
		if err != nil {
			return ""
		}
		_, err = io.WriteString(hash, string(buf)) // append into the hash
		if err != nil {
			return ""
		}
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
