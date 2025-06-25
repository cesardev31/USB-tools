package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func SelectISO(initialDir string) string {
	dir := ExpandHome(initialDir)

	for {
		files, _ := os.ReadDir(dir)
		var isos []string
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".iso") {
				isos = append(isos, filepath.Join(dir, f.Name()))
			}
		}

		if len(isos) == 0 {
			fmt.Println("⚠️ No se han encontrado imágenes santificadas en", dir)
			fmt.Println("¿Deseas introducir manualmente la ruta? (y/N):")
			if strings.ToLower(ReadLine()) != "y" {
				return ""
			}
			fmt.Print("Ingresa la ruta al patrón de carga sacra: ")
			dir = ReadLine()
			dir = ExpandHome(dir)
			continue
		}

		fmt.Println("\n📜 Imágenes de transferencia descubiertas:")
		for i, iso := range isos {
			fmt.Printf("[%d] %s\n", i+1, iso)
		}

		choice := Prompt("Elige el número correspondiente al patrón de carga sacra")
		idx := ParseChoice(choice, len(isos))
		if idx != -1 {
			return isos[idx]
		}
		fmt.Println("⚠️ Opción inválida. Intenta de nuevo.")
	}
}
