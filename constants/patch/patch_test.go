package patch

import (
	"testing"
	"time"

	"github.com/Kinveil/Riot-API-Golang/constants/region"
)

func TestShortPatch_String(t *testing.T) {
	tests := []struct {
		patch    ShortPatch
		expected string
	}{
		{ShortPatch(13.2), "13.2"},
		{ShortPatch(14.1), "14.1"},
		{ShortPatch(10.0), "10.0"},
		{ShortPatch(9.24), "9.2"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.patch.String(); got != tt.expected {
				t.Errorf("ShortPatch.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestShortPatch_FromString(t *testing.T) {
	tests := []struct {
		input       string
		expected    ShortPatch
		expectError bool
	}{
		{"13.2", ShortPatch(13.2), false},
		{"14.1", ShortPatch(14.1), false},
		{"10.0", ShortPatch(10.0), false},
		{"invalid", ShortPatch(0), true},
		{"13.2.1", ShortPatch(0), true}, // Should error - wrong format for ShortPatch
		{"", ShortPatch(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var p ShortPatch
			err := p.FromString(tt.input)

			if tt.expectError && err == nil {
				t.Errorf("Expected error for input %s, but got none", tt.input)
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error for input %s: %v", tt.input, err)
			}
			if !tt.expectError && p != tt.expected {
				t.Errorf("ShortPatch.FromString(%s) = %v, want %v", tt.input, p, tt.expected)
			}
		})
	}
}

func TestShortPatch_Compare(t *testing.T) {
	tests := []struct {
		patch1   ShortPatch
		patch2   ShortPatch
		expected int
	}{
		{ShortPatch(13.2), ShortPatch(13.1), 1},
		{ShortPatch(13.1), ShortPatch(13.2), -1},
		{ShortPatch(13.2), ShortPatch(13.2), 0},
		{ShortPatch(14.0), ShortPatch(13.9), 1},
		{ShortPatch(12.5), ShortPatch(13.0), -1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.patch1.Compare(tt.patch2); got != tt.expected {
				t.Errorf("ShortPatch(%v).Compare(%v) = %v, want %v", tt.patch1, tt.patch2, got, tt.expected)
			}
		})
	}
}

func TestPatch_String(t *testing.T) {
	tests := []struct {
		patch    Patch
		expected string
	}{
		{Patch("13.2.1"), "13.2.1"},
		{Patch("14.1.0"), "14.1.0"},
		{Patch("10.0.5"), "10.0.5"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := tt.patch.String(); got != tt.expected {
				t.Errorf("Patch.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPatch_FromString(t *testing.T) {
	tests := []struct {
		input       string
		expected    Patch
		expectError bool
	}{
		{"13.2.1", Patch("13.2.1"), false},
		{"14.1.0", Patch("14.1.0"), false},
		{"10.0.5", Patch("10.0.5"), false},
		{"13.2", Patch(""), true},     // Wrong format
		{"13.2.1.4", Patch(""), true}, // Wrong format
		{"invalid", Patch(""), true},
		{"13.a.1", Patch(""), true}, // Invalid number
		{"", Patch(""), true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var p Patch
			err := p.FromString(tt.input)

			if tt.expectError && err == nil {
				t.Errorf("Expected error for input %s, but got none", tt.input)
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error for input %s: %v", tt.input, err)
			}
			if !tt.expectError && p != tt.expected {
				t.Errorf("Patch.FromString(%s) = %v, want %v", tt.input, p, tt.expected)
			}
		})
	}
}

func TestPatch_ShortPatch(t *testing.T) {
	tests := []struct {
		patch    Patch
		expected ShortPatch
	}{
		{Patch("13.2.1"), ShortPatch(13.2)},
		{Patch("14.1.0"), ShortPatch(14.1)},
		{Patch("10.0.5"), ShortPatch(10.0)},
	}

	for _, tt := range tests {
		t.Run(string(tt.patch), func(t *testing.T) {
			if got := tt.patch.ShortPatch(); got != tt.expected {
				t.Errorf("Patch(%v).ShortPatch() = %v, want %v", tt.patch, got, tt.expected)
			}
		})
	}
}

func TestPatch_Compare(t *testing.T) {
	tests := []struct {
		patch1   Patch
		patch2   Patch
		expected int
	}{
		{Patch("13.2.1"), Patch("13.2.0"), 1},
		{Patch("13.2.0"), Patch("13.2.1"), -1},
		{Patch("13.2.1"), Patch("13.2.1"), 0},
		{Patch("14.0.0"), Patch("13.9.9"), 1},
		{Patch("13.1.5"), Patch("13.2.0"), -1},
		{Patch("13.2.1"), Patch("13.1.9"), 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.patch1.Compare(tt.patch2); got != tt.expected {
				t.Errorf("Patch(%v).Compare(%v) = %v, want %v", tt.patch1, tt.patch2, got, tt.expected)
			}
		})
	}
}

func TestNewPatchFromString(t *testing.T) {
	tests := []struct {
		input       string
		expected    Patch
		expectError bool
	}{
		{"13.2.1", Patch("13.2.1"), false},
		{"invalid", Patch(""), true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := NewPatchFromString(tt.input)

			if tt.expectError && err == nil {
				t.Errorf("Expected error for input %s, but got none", tt.input)
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error for input %s: %v", tt.input, err)
			}
			if !tt.expectError && got != tt.expected {
				t.Errorf("NewPatchFromString(%s) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestNewShortPatchFromString(t *testing.T) {
	tests := []struct {
		input       string
		expected    ShortPatch
		expectError bool
	}{
		{"13.2", ShortPatch(13.2), false},
		{"invalid", ShortPatch(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := NewShortPatchFromString(tt.input)

			if tt.expectError && err == nil {
				t.Errorf("Expected error for input %s, but got none", tt.input)
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error for input %s: %v", tt.input, err)
			}
			if !tt.expectError && got != tt.expected {
				t.Errorf("NewShortPatchFromString(%s) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestPatchWithStartTime_GetRegionStartDate(t *testing.T) {
	baseTime := time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC)

	pwst := &PatchWithStartTime{
		Patch:     ShortPatch(13.2),
		StartTime: baseTime,
		Shifts: map[region.Region]int{
			region.NA1:  3600, // 1 hour shift
			region.EUW1: 7200, // 2 hour shift
		},
	}

	tests := []struct {
		region   region.Region
		expected time.Time
	}{
		{region.NA1, baseTime.Add(1 * time.Hour)},
		{region.EUW1, baseTime.Add(2 * time.Hour)},
		{region.KR, baseTime}, // No shift defined, should return base time
	}

	for _, tt := range tests {
		t.Run(string(tt.region), func(t *testing.T) {
			got := pwst.GetRegionStartDate(tt.region)
			if !got.Equal(tt.expected) {
				t.Errorf("GetRegionStartDate(%v) = %v, want %v", tt.region, got, tt.expected)
			}
		})
	}
}
