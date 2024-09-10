package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	// Atur port serial
	c := &serial.Config{
		Name: "/dev/ttyUSB0", // Ganti dengan nama port yang terdeteksi di Raspberry Pi
		Baud: 57600,          // Kecepatan komunikasi, biasanya 9600
		//ReadTimeout: time.Second * 5,
	}

	var idCards []string

	// Buka koneksi ke port
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatalf("Gagal membuka port serial: %v", err)
	}
	defer s.Close()

	fmt.Println("Membaca data RFID dari Card Teck CT i809...")

	buf := make([]byte, 18)

	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Printf("Gagal membaca data: %v", err)
			continue
		}

		if n > 0 {
			if !contains(idCards, tagManipulasi) {
				idCards = append(idCards, tagManipulasi)

			// Konversi hasil baca ke string hex
				tag := hex.EncodeToString(buf)

				// Hapus 8 digit pertama dan 2 digit terakhir
				if len(tag) >= 20 {
					tagManipulasi := tag[8:32]
					timestamp := time.Now().Format("2006-01-02 15:04:05")
					fmt.Printf("Waktu: %s, RFID Tag (manipulasi): %s\n", timestamp, tagManipulasi)
					time.Sleep(5 * time.Second)
				} else {
					fmt.Println("Data tidak lengkap, tidak dapat dimanipulasi.")
				}
			
			}
		}
	}
}
