package random

import "testing"

func TestRandomString(t *testing.T) {
	// Тест 1: Проверяем корректность длины
	t.Run("correct length", func(t *testing.T) {
		length := 16
		str, err := NewRandomString(length)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(str) != length {
			t.Errorf("expected length %d, got %d", length, len(str))
		}
	})

	// Тест 2: Проверяем генерацию уникальных строк
	t.Run("unique strings", func(t *testing.T) {
		length := 16
		set := make(map[string]struct{})
		for i := 0; i < 1000; i++ {
			str, err := NewRandomString(length)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if _, exists := set[str]; exists {
				t.Errorf("duplicate string found: %s", str)
			}
			set[str] = struct{}{}
		}
	})

	// Тест 3: Проверяем обработку ошибки при некорректной длине
	t.Run("invalid length", func(t *testing.T) {
		_, err := NewRandomString(0)
		if err == nil {
			t.Errorf("expected error for length 0, got nil")
		}
		_, err = NewRandomString(-1)
		if err == nil {
			t.Errorf("expected error for negative length, got nil")
		}
	})

	// Тест 4: Проверяем содержимое строки
	t.Run("valid characters", func(t *testing.T) {
		length := 32
		str, err := NewRandomString(length)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for _, char := range str {
			if !contains(charset, char) {
				t.Errorf("unexpected character found: %c", char)
			}
		}
	})
}

// Вспомогательная функция для проверки наличия символа в строке
func contains(set string, char rune) bool {
	for _, c := range set {
		if c == char {
			return true
		}
	}
	return false
}
