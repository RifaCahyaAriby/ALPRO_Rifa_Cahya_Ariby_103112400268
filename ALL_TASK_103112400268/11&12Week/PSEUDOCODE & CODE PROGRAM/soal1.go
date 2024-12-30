package main

import "fmt"

func main() {
	var username, password string
	failedAttempts := 0 // ini 0 karena belum melakukan perulangan

	for {
		fmt.Print("Masukan username dan password: ") // input
		fmt.Scan(&username)
		fmt.Scan(&password)

		if username == "admin" && password == "admin" { // kondisi benar
			fmt.Println(failedAttempts)
			break
		} else { // kondisi sebaliknya
			failedAttempts++ // akan terus berulang jika salah
		}
	}
	fmt.Println("LoginBerhasil")
}


// PROGRAM LoginSistem

// DEKLARASI
    // Mendeklarasikan variabel untuk menyimpan username dan password
    //username, password : STRING
    // Mendeklarasikan variabel untuk menghitung jumlah percobaan gagal
    //failedAttempts : INTEGER

// MULAI
    // Inisialisasi jumlah percobaan gagal
   // failedAttempts = 0

    // Loop untuk meminta input username dan password
   // SELAMA TRUE DO
        // Tampilkan pesan untuk meminta input username dan password
      //  TAMPILKAN "Masukkan username dan password: "
        // Baca input username dari pengguna
      //  BACA username
        // Baca input password dari pengguna
      //  BACA password

        // Cek apakah username dan password benar
       // JIKA username == "admin" DAN password == "admin" THEN
            // Tampilkan jumlah percobaan gagal jika login berhasil
         //   TAMPILKAN failedAttempts 
            // Keluar dari loop jika login berhasil
         //   BREAK 
       // ELSE
            // Jika login gagal, increment jumlah percobaan gagal
         //   failedAttempts = failedAttempts + 1
       // ENDIF
   // END DO

    // Tampilkan pesan bahwa login berhasil setelah keluar dari loop
   // TAMPILKAN "Login Berhasil"

// SELESAI
