package util

import "crypto/sha1"

func HashBlobs(blob []byte) [20]byte {
	return sha1.Sum(blob)
}
