package utils

import (
	"fmt"
	"io"
	"os"
	"time"
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

	info, err := src.Stat()
	if err != nil {
		return fmt.Errorf("no se pudo obtener el tamaño de la ISO: %w", err)
	}
	totalSize := info.Size()

	buf := make([]byte, 32*1024*1024) // 32 MB
	var written int64
	start := time.Now()

	for {
		nr, err := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[:nr])
			if ew != nil {
				return fmt.Errorf("error al escribir: %w", ew)
			}
			written += int64(nw)

			// Calcular porcentaje y tiempo estimado
			percent := float64(written) * 100 / float64(totalSize)
			elapsed := time.Since(start)
			var eta time.Duration
			if percent > 0 {
				eta = time.Duration(float64(elapsed) * (100.0/percent - 1))
			}

			fmt.Printf("\r📦 Transferencia: %.2f%% | ⏳ Tiempo restante estimado: %s ", percent, eta.Truncate(time.Second))
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error al leer: %w", err)
		}
	}
	fmt.Println("\n🧘 Sellando los últimos bloques del Espíritu Máquina...")

	// Timeout de sync
	done := make(chan error, 1)
	go func() {
		done <- dst.Sync()
	}()

	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("falló al sincronizar los bloques: %w", err)
		}
	case <-time.After(20 * time.Second):
		return fmt.Errorf("⏳ el ritual de sellado está tardando demasiado")
	}

	fmt.Printf("✅ Transferencia completa. Se escribieron %.2f MB en %s\n",
		float64(written)/1024.0/1024.0,
		time.Since(start).Truncate(time.Second))

	return nil
}
