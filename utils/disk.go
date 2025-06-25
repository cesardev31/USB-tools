package utils

import (
	"fmt"
	"os"
	"strings"
)

func ListRemovableDisks() []string {
	entries, err := os.ReadDir("/sys/block")
	if err != nil {
		fmt.Println("❌ Error al acceder a /sys/block:", err)
		return nil
	}
	disks := []string{}
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, "sr") {
			continue
		}
		rm, _ := os.ReadFile(fmt.Sprintf("/sys/block/%s/removable", name))
		model, _ := os.ReadFile(fmt.Sprintf("/sys/block/%s/device/model", name))
		if strings.TrimSpace(string(rm)) == "1" {
			size := getSizeGiB(name)
			disks = append(disks, fmt.Sprintf("/dev/%s (%.1fG %s)", name, size, strings.TrimSpace(string(model))))
		}
	}
	return disks
}

func getSizeGiB(name string) float64 {
	data, err := os.ReadFile(fmt.Sprintf("/sys/block/%s/size", name))
	if err != nil {
		return 0
	}
	var sectors int64
	fmt.Sscanf(strings.TrimSpace(string(data)), "%d", &sectors)
	return float64(sectors*512) / (1024 * 1024 * 1024)
}
