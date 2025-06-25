package utils

import (
	"fmt"
	"io"
	"os"
)

func TransferISO(isoPath, diskPath string) error {
	src, err := os.Open(isoPath)
	if err != nil {
		return fmt.Errorf("no se pudo abrir la imagen ISO: %w", err)
	}
	defer src.Close()

	dst, err := os.OpenFile(diskPath, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el dispositivo: %w", err)
	}
	defer dst.Close()

	buf := make([]byte, 4*1024*1024)
	var total int64
	for {
		n, err := src.Read(buf)
		if n > 0 {
			w, err := dst.Write(buf[:n])
			if err != nil {
				return fmt.Errorf("fallo al escribir en el dispositivo: %w", err)
			}
			total += int64(w)
			progress := float64(total) / float64(getFileSize(isoPath)) * 100
			fmt.Printf("\r📦 Transferencia en progreso: %.2f%%", progress)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error durante la transferencia: %w", err)
		}
	}
	done := make(chan bool)
	go ShowSpinner(done, "🧘 Sellando los últimos bloques del Espíritu Máquina...")

	err = dst.Sync()
	done <- true
	if err != nil {
		return fmt.Errorf("falló al sellar los bloques: %w", err)
	}

	fmt.Printf("✅ Transferencia completa. Se escribieron %.2f MB.\n", float64(total)/1024/1024)
	return nil
}

func getFileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 1
	}
	return info.Size()
}
