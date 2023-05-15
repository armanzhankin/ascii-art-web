package Hash

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
)

func GetMD5Hash(path string, b string) error {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("ERR READ FILE")
		os.Exit(0)
	}
	// || hash != shadowHash || hash != thinkertoyHash
	shadowHash := [16]byte{164, 157, 95, 203, 13, 92, 89, 178, 231, 118, 116, 170, 58, 184, 187, 177}
	thinkertoyHash := [16]byte{134, 217, 148, 116, 87, 246, 164, 26, 24, 203, 152, 66, 126, 49, 79, 248}
	correctHash := [16]byte{183, 224, 110, 127, 106, 45, 36, 216, 218, 93, 87, 211, 203, 166, 162, 199}

	hash := md5.Sum([]byte(string(dat)))

	switch b {
	case "shadow":
		if hash != shadowHash {
			return errors.New("Textfile was modified,return it back")
		}
	case "standard":
		if hash != correctHash {
			return errors.New("Textfile was modified,return it back")
		}
	case "thinkertoy":
		if hash != thinkertoyHash {
			return errors.New("Textfile was modified,return it back")
		}
	}
	return nil
}
