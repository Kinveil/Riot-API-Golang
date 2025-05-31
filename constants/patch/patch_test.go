package patch

import (
	"testing"
	"time"

	"github.com/Kinveil/Riot-API-Golang/constants/region"
)

func TestShortPatch_String(t *testing.T) {
	tests := []struct {
		name     string
		patch    ShortPatch
		expected string
	}{
		{"Basic patch", ShortPatch{Major: 13, Minor: 2}, "13.2"},
		{"Zero patch", ShortPatch{Major: 0, Minor: 0}, "0.0"},
		{"High version", ShortPatch{Major: 25, Minor: 3}, "25.3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.patch.String()
			if result != tt.expected {
				t.Errorf("ShortPatch.String() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestShortPatch_FromString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    ShortPatch
		expectError bool
	}{
		{"Valid patch", "13.2", ShortPatch{Major: 13, Minor: 2}, false},
		{"Zero patch", "0.0", ShortPatch{Major: 0, Minor: 0}, false},
		{"High version", "25.3", ShortPatch{Major: 25, Minor: 3}, false},
		{"Invalid format - too many parts", "13.2.1", ShortPatch{}, true},
		{"Invalid format - too few parts", "13", ShortPatch{}, true},
		{"Invalid format - non-numeric major", "a.2", ShortPatch{}, true},
		{"Invalid format - non-numeric minor", "13.b", ShortPatch{}, true},
		{"Empty string", "", ShortPatch{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p ShortPatch
			err := p.FromString(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("ShortPatch.FromString() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("ShortPatch.FromString() unexpected error: %v", err)
				}
				if p.Major != tt.expected.Major || p.Minor != tt.expected.Minor {
					t.Errorf("ShortPatch.FromString() = %v, want %v", p, tt.expected)
				}
			}
		})
	}
}

func TestShortPatch_Compare(t *testing.T) {
	tests := []struct {
		name     string
		patch1   ShortPatch
		patch2   ShortPatch
		expected int
	}{
		{"Equal patches", ShortPatch{13, 2}, ShortPatch{13, 2}, 0},
		{"First patch higher major", ShortPatch{14, 1}, ShortPatch{13, 9}, 1},
		{"First patch lower major", ShortPatch{13, 9}, ShortPatch{14, 1}, -1},
		{"Same major, first patch higher minor", ShortPatch{13, 3}, ShortPatch{13, 2}, 1},
		{"Same major, first patch lower minor", ShortPatch{13, 2}, ShortPatch{13, 3}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.patch1.Compare(tt.patch2)
			if result != tt.expected {
				t.Errorf("ShortPatch.Compare() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPatch_String(t *testing.T) {
	tests := []struct {
		name     string
		patch    Patch
		expected string
	}{
		{"Basic patch", Patch{Major: 13, Minor: 2, Patch: 1}, "13.2.1"},
		{"Zero patch", Patch{Major: 0, Minor: 0, Patch: 0}, "0.0.0"},
		{"High version", Patch{Major: 25, Minor: 3, Patch: 4}, "25.3.4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.patch.String()
			if result != tt.expected {
				t.Errorf("Patch.String() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPatch_FromString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    Patch
		expectError bool
	}{
		{"Valid patch", "13.2.1", Patch{Major: 13, Minor: 2, Patch: 1}, false},
		{"Zero patch", "0.0.0", Patch{Major: 0, Minor: 0, Patch: 0}, false},
		{"High version", "25.3.4", Patch{Major: 25, Minor: 3, Patch: 4}, false},
		{"Invalid format - too few parts", "13.2", Patch{}, true},
		{"Invalid format - too many parts", "13.2.1.0", Patch{}, true},
		{"Invalid major", "a.2.1", Patch{}, true},
		{"Invalid minor", "13.b.1", Patch{}, true},
		{"Invalid patch", "13.2.c", Patch{}, true},
		{"Empty string", "", Patch{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p Patch
			err := p.FromString(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("Patch.FromString() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Patch.FromString() unexpected error: %v", err)
				}
				if p.Major != tt.expected.Major || p.Minor != tt.expected.Minor || p.Patch != tt.expected.Patch {
					t.Errorf("Patch.FromString() = %v, want %v", p, tt.expected)
				}
			}
		})
	}
}

func TestPatch_ShortPatch(t *testing.T) {
	tests := []struct {
		name     string
		patch    Patch
		expected ShortPatch
	}{
		{"Basic patch", Patch{Major: 13, Minor: 2, Patch: 1}, ShortPatch{Major: 13, Minor: 2}},
		{"Zero patch", Patch{Major: 0, Minor: 0, Patch: 0}, ShortPatch{Major: 0, Minor: 0}},
		{"High version", Patch{Major: 25, Minor: 3, Patch: 4}, ShortPatch{Major: 25, Minor: 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.patch.ShortPatch()
			if result.Major != tt.expected.Major || result.Minor != tt.expected.Minor {
				t.Errorf("Patch.ShortPatch() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPatch_Compare(t *testing.T) {
	tests := []struct {
		name     string
		patch1   Patch
		patch2   Patch
		expected int
	}{
		{"Equal patches", Patch{13, 2, 1}, Patch{13, 2, 1}, 0},
		{"First patch higher major", Patch{14, 1, 0}, Patch{13, 9, 9}, 1},
		{"First patch lower major", Patch{13, 9, 9}, Patch{14, 1, 0}, -1},
		{"Same major, first patch higher minor", Patch{13, 3, 0}, Patch{13, 2, 9}, 1},
		{"Same major, first patch lower minor", Patch{13, 2, 9}, Patch{13, 3, 0}, -1},
		{"Same major/minor, first patch higher patch", Patch{13, 2, 2}, Patch{13, 2, 1}, 1},
		{"Same major/minor, first patch lower patch", Patch{13, 2, 1}, Patch{13, 2, 2}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.patch1.Compare(tt.patch2)
			if result != tt.expected {
				t.Errorf("Patch.Compare() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNewPatchFromString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    Patch
		expectError bool
	}{
		{"Valid patch", "13.2.1", Patch{Major: 13, Minor: 2, Patch: 1}, false},
		{"Invalid patch", "invalid", Patch{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewPatchFromString(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("NewPatchFromString() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("NewPatchFromString() unexpected error: %v", err)
				}
				if result.Major != tt.expected.Major || result.Minor != tt.expected.Minor || result.Patch != tt.expected.Patch {
					t.Errorf("NewPatchFromString() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestNewShortPatchFromString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    ShortPatch
		expectError bool
	}{
		{"Valid patch", "13.2", ShortPatch{Major: 13, Minor: 2}, false},
		{"Invalid patch", "invalid", ShortPatch{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewShortPatchFromString(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("NewShortPatchFromString() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("NewShortPatchFromString() unexpected error: %v", err)
				}
				if result.Major != tt.expected.Major || result.Minor != tt.expected.Minor {
					t.Errorf("NewShortPatchFromString() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestPatchWithStartTime_GetRegionStartDate(t *testing.T) {
	baseTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		patch    PatchWithStartTime
		region   region.Region
		expected time.Time
	}{
		{
			name: "Region with shift",
			patch: PatchWithStartTime{
				Patch:     ShortPatch{Major: 13, Minor: 2},
				StartTime: baseTime,
				Shifts: map[region.Region]int{
					region.NA1: 3600, // 1 hour shift
				},
			},
			region:   region.NA1,
			expected: baseTime.Add(time.Hour),
		},
		{
			name: "Region without shift",
			patch: PatchWithStartTime{
				Patch:     ShortPatch{Major: 13, Minor: 2},
				StartTime: baseTime,
				Shifts:    map[region.Region]int{},
			},
			region:   region.EUW1,
			expected: baseTime,
		},
		{
			name: "Negative shift",
			patch: PatchWithStartTime{
				Patch:     ShortPatch{Major: 13, Minor: 2},
				StartTime: baseTime,
				Shifts: map[region.Region]int{
					region.KR: -1800, // -30 minutes
				},
			},
			region:   region.KR,
			expected: baseTime.Add(-30 * time.Minute),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.patch.GetRegionStartDate(tt.region)
			if !result.Equal(tt.expected) {
				t.Errorf("PatchWithStartTime.GetRegionStartDate() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkPatch_FromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var p Patch
		p.FromString("13.2.1")
	}
}

func BenchmarkPatch_String(b *testing.B) {
	p := Patch{Major: 13, Minor: 2, Patch: 1}
	for i := 0; i < b.N; i++ {
		_ = p.String()
	}
}

func BenchmarkPatch_ShortPatch(b *testing.B) {
	p := Patch{Major: 13, Minor: 2, Patch: 1}
	for i := 0; i < b.N; i++ {
		_ = p.ShortPatch()
	}
}

func BenchmarkPatch_Compare(b *testing.B) {
	p1 := Patch{Major: 13, Minor: 2, Patch: 1}
	p2 := Patch{Major: 13, Minor: 2, Patch: 2}
	for i := 0; i < b.N; i++ {
		_ = p1.Compare(p2)
	}
}
