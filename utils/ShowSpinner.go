package utils

import (
	"fmt"
	"time"
)

func ShowSpinner(done <-chan bool, mensaje string) {
	frames := []rune{'|', '/', '-', '\\'}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\r") // Borra la línea del spinner
			return
		default:
			fmt.Printf("\r%s %c", mensaje, frames[i%len(frames)])
			i++
			time.Sleep(150 * time.Millisecond)
		}
	}
}
