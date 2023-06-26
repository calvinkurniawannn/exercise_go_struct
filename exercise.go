// Buatlah class mobil, mobil tersebut memiliki maksimal 4 roda, 
// semua roda tersebut bisa diganti dengan berbagai 
// jenis ban (ban karet, ban kayu, ban besi). Mobil memiliki 2 pintu yang sama, 
// namun berbeda sifat. Pintu kanan ketika diketuk berbunyi “Knock” dan ketika 
// dibuka berbunyi “beep”, sedangkan bunyi pintu kiri kebalikan dari bunyi 
// pintu kanan. Lakukan penggantian roda pada mobil tersebut dan coba ketuk 
// dan buka kedua pintunya. 
//Buatlah repo github dengan akses publik, dan copy kan link repo	
package main
import "fmt"

type Mobil struct{
	Roda [4]string
	PintuKiri Pintu
	PintuKanan Pintu
}
type Pintu struct{
	BunyiKetuk string
	BunyiBuka string
}

func main(){
	mobil := Mobil{
		Roda:[4]string{"ban karet","ban karet","ban karet","ban karet"},
		PintuKiri: Pintu{
			BunyiKetuk:"Beep",
			BunyiBuka:"Knock",
		},
		PintuKanan: Pintu{
			BunyiKetuk:"Knock",
			BunyiBuka:"Beep",
		},
	}

	mobil.GantiRoda("ban kayu")

	fmt.Println("Jenis Ban Rodanya adalah ",mobil.Roda)
	fmt.Println("Suara Pintu Kiri : ",mobil.PintuKiri)
	fmt.Println("Suara Pintu Kanan : ",mobil.PintuKanan)
}

func (m *Mobil) GantiRoda(jenisBan string) {
	for i := 0; i < len(m.Roda); i++ {
		m.Roda[i] = jenisBan
	}
}