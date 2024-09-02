package maps

import "testing"

func TestIsMap(t *testing.T) {
	t.Run("Test IsMap true", func(t *testing.T) {
		if !IsMap(map[string]int{"a": 1, "b": 2, "c": 3}) {
			t.Errorf("IsMap() = %v, want %v", false, true)
		}
	})

	t.Run("Test IsMap false", func(t *testing.T) {
		if IsMap([]int{1, 2, 3}) {
			t.Errorf("IsMap() = %v, want %v", true, false)
		}
	})
}

func TestInMap(t *testing.T) {
	t.Run("Test InMap true", func(t *testing.T) {
		if !InMap("a", map[string]int{"a": 1, "b": 2, "c": 3}) {
			t.Errorf("InMap() = %v, want %v", false, true)
		}
	})

	t.Run("Test InMap false", func(t *testing.T) {
		if InMap("d", map[string]int{"a": 1, "b": 2, "c": 3}) {
			t.Errorf("InMap() = %v, want %v", true, false)
		}
	})
}

func TestTypesEqual(t *testing.T) {
	t.Run("Test TypesEqual true", func(t *testing.T) {
		if !TypesEqual(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 3}) {
			t.Errorf("TypesEqual() = %v, want %v", false, true)
		}
	})

	t.Run("Test TypesEqual false", func(t *testing.T) {
		if TypesEqual(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]string{"a": "1", "b": "2", "c": "3"}) {
			t.Errorf("TypesEqual() = %v, want %v", true, false)
		}
	})
}

func TestValuesEqual(t *testing.T) {
	t.Run("Test ValuesEqual true", func(t *testing.T) {
		if !ValuesEqual(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 3}) {
			t.Errorf("ValuesEqual() = %v, want %v", false, true)
		}
	})

	t.Run("Test ValuesEqual false", func(t *testing.T) {
		if ValuesEqual(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 4}) {
			t.Errorf("ValuesEqual() = %v, want %v", true, false)
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("Test Contains true", func(t *testing.T) {
		if !Contains(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2}) {
			t.Errorf("Contains() = %v, want %v", false, true)
		}
	})

	t.Run("Test Contains false", func(t *testing.T) {
		if Contains(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 4}) {
			t.Errorf("Contains() = %v, want %v", true, false)
		}
	})
}

func TestMerges(t *testing.T) {
	t.Run("Test Merges true", func(t *testing.T) {
		if !ValuesEqual(Merges(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "d": 4}), map[interface{}]interface{}{"a": 1, "b": 2, "c": 3, "d": 4}) {
			t.Errorf("Merges() = %v, want %v", false, true)
		}
	})

	t.Run("Test Merges false", func(t *testing.T) {
		if ValuesEqual(Merges(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 4}), map[interface{}]interface{}{"a": 1, "b": 2, "c": 3}) {
			t.Errorf("Merges() = %v, want %v", true, false)
		}
	})
}

func TestSub(t *testing.T) {
	t.Run("Test Sub true", func(t *testing.T) {
		if !ValuesEqual(Sub(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2}), map[interface{}]interface{}{"c": 3}) {
			t.Errorf("Sub() = %v, want %v", false, true)
		}
	})

	t.Run("Test Sub false", func(t *testing.T) {
		if len(Sub(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 3}).(map[interface{}]interface{})) != 0 {
			t.Errorf("Sub() = %v, want %v", true, false)
		}
	})
}

func TestDifferentSet(t *testing.T) {
	t.Run("Test DifferentSet true", func(t *testing.T) {
		if !ValuesEqual(DifferentSet(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "d": 4}), map[interface{}]interface{}{"c": 3, "d": 4}) {
			t.Errorf("DifferentSet() = %v, want %v", false, true)
		}
	})

	t.Run("Test DifferentSet false", func(t *testing.T) {
		if len(DifferentSet(map[string]int{"a": 1, "b": 2, "d": 4}, map[string]int{"a": 1, "b": 2, "d": 4}).(map[interface{}]interface{})) != 0 {
			t.Errorf("DifferentSet() = %v, want %v", true, false)
		}
	})
}
