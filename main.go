// buat struct, inisiasi,
// var dengan tipe struct diatas, bisa ubah nilai inisiasi.
package main

import "fmt"

type Kendaraan struct{
	Nama string
	Kecepatan int
}
type JarakTempuh interface {
	hitungJarak()
}
func (obj Kendaraan) hitungJarak(bensin int) int{
	result := bensin*obj.Kecepatan
	return result
}
func cariTerEfisien(bensin int, arr []Kendaraan) string {
	var result string
	var jarakTempuhTerbesar int
	for _, v := range arr {
		JarakTempuh := v.hitungJarak(bensin)
		if jarakTempuhTerbesar <=  JarakTempuh{
			jarakTempuhTerbesar = JarakTempuh
			result = v.Nama
		}
	}
	return result
}
func main() {
	arr := []Kendaraan{{"mobil", 1},{"motor", 3},{"bajaj",4}}
	bensin := 10
	fmt.Println("Kendaraan Paling Efisien :", cariTerEfisien(bensin,arr))
}