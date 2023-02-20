package testunitario

import "testing"

func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 10 {
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 10)
	}
}

func TestSumaTable(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, item.c)
		}
	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{4, 2, 4},
		{5, 2, 5},
		{2, 3, 3},
	}

	for _, item := range tabla {
		max := GetMax(item.a, item.b)

		if max != item.c {
			t.Errorf("GetMax incorrecta, tiene %d se esperaba %d", max, item.c)
		}
	}
}

func TestFib(t *testing.T) {
	tabla := []struct {
		a int
		b int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tabla {
		fibo := Fibonacci(item.a)

		if fibo != item.b {
			t.Errorf("GetMax incorrecta, tiene %d se esperaba %d", fibo, item.b)
		}
	}
}
