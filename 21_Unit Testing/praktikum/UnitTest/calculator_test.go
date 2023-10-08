package calculator

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(5, 3)
	if result != 8 {
		t.Errorf("Harusnya 5 + 3 = 8, namun hasilnya %f", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(10, 4)
	if result != 6 {
		t.Errorf("Harusnya 10 - 4 = 6, namun hasilnya %f", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(8, 2)
	if err != nil {
		t.Errorf("Pembagian seharusnya berhasil: %v", err)
	}
	if result != 4 {
		t.Errorf("Harusnya 8 / 2 = 4, namun hasilnya %f", result)
	}

	_, err = Division(5, 0)
	if err == nil {
		t.Errorf("Seharusnya terjadi error pembagian oleh nol")
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(6, 7)
	if result != 42 {
		t.Errorf("Harusnya 6 * 7 = 42, namun hasilnya %f", result)
	}
}
