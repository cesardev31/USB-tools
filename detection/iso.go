package detection

import (
	"os"
	"strings"
)

func DetectISOType(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	header := make([]byte, 512*64) // 32 KB
	file.Read(header)

	data := strings.ToLower(string(header))
	switch {
	case strings.Contains(data, "efi") && strings.Contains(data, "bootmgr"):
		return "windows"
	case strings.Contains(data, "casper") || strings.Contains(data, "live"):
		return "linux"
	default:
		return "unknown"
	}
}
