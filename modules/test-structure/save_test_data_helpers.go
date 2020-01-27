package test_structure

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

// GetSaveTestDataPath gets unique temporary directory path for each test
func GetSaveTestDataPath(t *testing.T) string {
	name := t.Name()
	bytes := []byte(name)
	md5 := md5.Sum(bytes)
	hash := hex.EncodeToString(md5[:])
	return filepath.Join(os.TempDir(), hash)
}
