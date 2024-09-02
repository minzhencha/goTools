package slices

import (
	"testing"
)

func TestIsSlice(t *testing.T) {
	t.Run("Test IsSlice true", func(t *testing.T) {
		if !IsSlice([]int{1, 2, 3}) {
			t.Errorf("IsSlice() = %v, want %v", false, true)
		}
	})

	t.Run("Test IsSlice false", func(t *testing.T) {
		if IsSlice(1) {
			t.Errorf("IsSlice() = %v, want %v", true, false)
		}
	})
}

func TestInSlice(t *testing.T) {
	t.Run("Test InSlice true", func(t *testing.T) {
		if !InSlice(1, []int{1, 2, 3}) {
			t.Errorf("InSlice() = %v, want %v", false, true)
		}
	})

	t.Run("Test InSlice false", func(t *testing.T) {
		if InSlice(4, []int{1, 2, 3}) {
			t.Errorf("InSlice() = %v, want %v", true, false)
		}
	})
}

func TestSliceTypesEqual(t *testing.T) {
	t.Run("Test TypesEqual true", func(t *testing.T) {
		if !TypesEqual([]int{1, 2, 3}, []int{1, 2, 3}) {
			t.Errorf("TypesEqual() = %v, want %v", false, true)
		}
	})

	t.Run("Test TypesEqual false", func(t *testing.T) {
		if TypesEqual([]int{1, 2, 3}, []string{"1", "2", "3"}) {
			t.Errorf("TypesEqual() = %v, want %v", true, false)
		}
	})
}

func TestSliceValuesEqual(t *testing.T) {
	t.Run("Test ValuesEqual true", func(t *testing.T) {
		if !ValuesEqual([]int{1, 2, 3}, []int{1, 2, 3}) {
			t.Errorf("ValuesEqual() = %v, want %v", false, true)
		}
	})

	t.Run("Test ValuesEqual false", func(t *testing.T) {
		if ValuesEqual([]int{1, 2, 3}, []int{3, 2, 1}) {
			t.Errorf("ValuesEqual() = %v, want %v", true, false)
		}
	})
}

func TestSliceValuesSame(t *testing.T) {
	t.Run("Test ValuesSame true", func(t *testing.T) {
		if !ValuesSame([]int{1, 2, 3}, []int{3, 2, 1}) {
			t.Errorf("ValuesSame() = %v, want %v", false, true)
		}
	})

	t.Run("Test ValuesSame false", func(t *testing.T) {
		if ValuesSame([]int{1, 2, 3}, []int{3, 2, 4}) {
			t.Errorf("ValuesSame() = %v, want %v", true, false)
		}
	})
}

func TestSliceContains(t *testing.T) {
	t.Run("Test Contains true", func(t *testing.T) {
		if !Contains([]int{1, 2, 3}, []int{1, 2}) {
			t.Errorf("Contains() = %v, want %v", false, true)
		}
	})

	t.Run("Test Contains false", func(t *testing.T) {
		if Contains([]int{1, 2, 3}, []int{1, 2, 4}) {
			t.Errorf("Contains() = %v, want %v", true, false)
		}
	})
}

func TestSliceMerges(t *testing.T) {
	t.Run("Test Merges true", func(t *testing.T) {
		if !ValuesEqual(Merges([]int{1, 2, 3}, []int{4, 5, 6}), []int{1, 2, 3, 4, 5, 6}) {
			t.Errorf("Merges() = %v, want %v", false, true)
		}
	})

	t.Run("Test Merges false", func(t *testing.T) {
		if !ValuesEqual(Merges([]int{1, 2, 3}, []int{1, 2, 3}), []int{1, 2, 3, 1, 2, 3}) {
			t.Errorf("Merges() = %v, want %v", false, true)
		}
	})
}

func TestSliceSub(t *testing.T) {
	t.Run("Test Sub true", func(t *testing.T) {
		if !ValuesEqual(Sub([]int{1, 2, 3}, []int{1, 2}), []int{3}) {
			t.Errorf("Sub() = %v, want %v", false, true)
		}
	})

	t.Run("Test Sub false", func(t *testing.T) {
		if !ValuesEqual(Sub([]int{1, 2, 3}, []int{1, 2, 4}), []int{3}) {
			t.Errorf("Sub() = %v, want %v", false, true)
		}
	})
}

func TestSliceDeduplicate(t *testing.T) {
	t.Run("Test Deduplicate true", func(t *testing.T) {
		if !ValuesEqual(Deduplicate([]int{1, 2, 3, 1, 2, 3}), []int{1, 2, 3}) {
			t.Errorf("Deduplicate() = %v, want %v", false, true)
		}
	})
}

func TestSliceDifferentSet(t *testing.T) {
	t.Run("Test DifferentSet true", func(t *testing.T) {
		if !ValuesEqual(DifferentSet([]int{1, 2, 3}, []int{2, 3, 4}), []int{4, 1}) {
			t.Errorf("DifferentSet() = %v, want %v", false, true)
		}
	})
}
