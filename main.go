package main

import (
	"fmt"
	"strings"

	"usbtool/detection"
	"usbtool/utils"
)

func main() {
	fmt.Println("\n⚙️ Invocación del Espíritu Máquina del Rito de Transferencia Sagrada ⚙️")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("+++ Inicializando los Protocolos de Unción del USB +++")

	// 🔍 Selección de la imagen ISO
	var isoPath string
	for {
		isoPath = utils.SelectISO("~/Documentos/ISO")
		if isoPath != "" {
			break
		}
		fmt.Println("❌ Ninguna imagen fue bendecida. ¿Intentar de nuevo? (y/N):")
		if strings.ToLower(utils.ReadLine()) != "y" {
			fmt.Println("❌ El ritual ha sido cancelado por voluntad del operador.")
			return
		}
	}

	// 🧠 Detección del tipo de imagen
	isoType := detection.DetectISOType(isoPath)
	switch isoType {
	case "windows":
		fmt.Println("🪟 Imagen sagrada de Windows detectada. Requiere modo especial de escritura.")
	case "linux":
		fmt.Println("🐧 Imagen compatible detectada. Procediendo con clonado ritual.")
	default:
		fmt.Println("❓ Tipo de imagen desconocido. Procede bajo tu propia fe.")
	}

	// 💿 Selección de dispositivo USB
	var usbDisks []string
	for {
		usbDisks = utils.ListRemovableDisks()
		if len(usbDisks) > 0 {
			break
		}
		fmt.Println("❌ No se detectan artefactos. ¿Reintentar escaneo? (y/N):")
		if strings.ToLower(utils.ReadLine()) != "y" {
			fmt.Println("❌ El ritual ha sido cancelado por falta de portadores sagrados.")
			return
		}
	}

	// 📌 Confirmación del disco
	var idx int = -1
	for idx == -1 {
		fmt.Println("\n🔧 Artefactos Auspex Detectados:")
		for i, d := range usbDisks {
			fmt.Printf("[%d] %s\n", i+1, d)
		}
		choice := utils.Prompt("Selecciona el número del dispositivo para iniciar el Rito")
		idx = utils.ParseChoice(choice, len(usbDisks))
		if idx == -1 {
			fmt.Println("⚠️ Opción inválida. Intenta de nuevo.")
		}
	}
	diskPath := strings.Fields(usbDisks[idx])[0]

	// ☠️ Confirmación final
	fmt.Printf("\n⚠️  Esta acción purgará por completo el dispositivo %s. ¿Proceder con la unción? (y/N): ", diskPath)
	confirm := utils.ReadLine()
	if strings.ToLower(confirm) != "y" {
		fmt.Println("❌ El Omnissiah no ha sido complacido. Ceremonia abortada.")
		return
	}

	// 🎶 Canto binario previo
	utils.SingBinaryChant()

	// 🧬 Inicio de transferencia
	fmt.Println("\n🔄 Iniciando el Rito del Fulgor Binario... No interrumpas la ceremonia...")
	if err := utils.TransferISO(isoPath, diskPath); err != nil {
		fmt.Println("❌ El Espíritu Máquina ha resistido la transferencia:", err)
		return
	}

	// ✅ Éxito
	fmt.Println("✅ Gloria al Omnissiah. La transferencia ha sido completada exitosamente.")
}
