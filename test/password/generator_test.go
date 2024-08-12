package password

import (
	"passgenie/internal/password" // Asegúrate de que el módulo y la ruta sean correctos
	"testing"
)

func TestGenerate(t *testing.T) {
	g := password.NewGenerator()
	password, err := g.Generate(12, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(password) != 12 {
		t.Fatalf("expected password length of 12, got %d", len(password))
	}
}

func TestGenerateWithoutSymbols(t *testing.T) {
	g := password.NewGenerator()
	password, err := g.Generate(8, false)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(password) != 8 {
		t.Fatalf("expected password length of 8, got %d", len(password))
	}
}
