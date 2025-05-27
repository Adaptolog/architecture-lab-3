package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// Цей тест перевіряє лише що пакет компілюється
	// Оскільки main() не повертає значення, ми не можемо її протестувати напряму
	t.Log("Main package compiled successfully")
}