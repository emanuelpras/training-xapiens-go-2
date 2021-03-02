package hitung

import "math"

func VolumeKubus(angka float64) (volume float64) {
	//rumus volume kubus
	// volume = angka * angka * angka
	volume = math.Pow(angka, 3) // artinya value dari angka dipangkat 3
	return
}

func LuasPersegi(angka float64) (luas float64) {
	luas = angka * angka
	return
}
