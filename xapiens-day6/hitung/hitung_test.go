package hitung

import "testing"

var (
	expectedResultVolume float64 = 125
	expectedResultLuas   float64 = 25
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Hasil volume : %.2f ", VolumeKubus(5))

	//5 pangkat 3 hasilnya 125
	if VolumeKubus(5) != expectedResultVolume {
		t.Errorf("Wrong, hasil seharusnya : %.2f ", expectedResultVolume)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Hasil luas : %.2f ", LuasPersegi(5))

	//5 pangkat 3 hasilnya 125
	if LuasPersegi(5) != expectedResultLuas {
		t.Errorf("Wrong, hasil seharusnya : %.2f ", expectedResultLuas)
	}
}

// ngga recommend, karena jika salah satu ada yang fail maka function testing akan di set fail
// func TestHitungVolumeLuas(t *testing.T) {
// 	t.Logf("Hasil volume : %.2f ", VolumeKubus(5))

// 	//5 pangkat 3 hasilnya 125
// 	if VolumeKubus(5) != expectedResultVolume {
// 		t.Errorf("Wrong, hasil seharusnya : %.2f ", expectedResultVolume)
// 	}

// 	t.Logf("Hasil luas : %.2f ", LuasPersegi(5))

// 	//5 pangkat 3 hasilnya 125
// 	if LuasPersegi(5) != expectedResultLuas {
// 		t.Errorf("Wrong, hasil seharusnya : %.2f ", expectedResultLuas)
// 	}
// }
