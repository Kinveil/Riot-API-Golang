package patch

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Kinveil/Riot-API-Golang/constants/region"
)

type ShortPatch float32

func (v ShortPatch) String() string {
	return fmt.Sprintf("%.1f", float32(v))
}

// FromString creates a ShortPatch from a string like "13.2"
func (v *ShortPatch) FromString(s string) error {
	val, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return fmt.Errorf("invalid short patch format: %s", s)
	}
	*v = ShortPatch(val)
	return nil
}

type Patch string

func (v Patch) String() string {
	return string(v)
}

// FromString creates a Patch from a string like "13.2.1"
func (v *Patch) FromString(s string) error {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return fmt.Errorf("invalid patch format: %s", s)
	}

	// Validate each part is a number
	for _, part := range parts {
		if _, err := strconv.Atoi(part); err != nil {
			return fmt.Errorf("invalid patch format: %s", s)
		}
	}

	*v = Patch(s)
	return nil
}

// ShortPatch returns the short version of the patch (removes patch component)
func (v Patch) ShortPatch() ShortPatch {
	parts := strings.Split(string(v), ".")
	if len(parts) < 2 {
		return ShortPatch(0)
	}

	shortStr := parts[0] + "." + parts[1]
	val, err := strconv.ParseFloat(shortStr, 32)
	if err != nil {
		return ShortPatch(0)
	}

	return ShortPatch(val)
}

// Compare returns -1 if v < other, 0 if v == other, 1 if v > other
func (v Patch) Compare(other Patch) int {
	vParts := strings.Split(string(v), ".")
	otherParts := strings.Split(string(other), ".")

	for i := 0; i < 3; i++ {
		vNum, _ := strconv.Atoi(vParts[i])
		otherNum, _ := strconv.Atoi(otherParts[i])

		if vNum < otherNum {
			return -1
		}
		if vNum > otherNum {
			return 1
		}
	}

	return 0
}

// Compare returns -1 if v < other, 0 if v == other, 1 if v > other
func (v ShortPatch) Compare(other ShortPatch) int {
	if v < other {
		return -1
	}
	if v > other {
		return 1
	}
	return 0
}

type PatchWithStartTime struct {
	Patch     ShortPatch
	StartTime time.Time
	Shifts    map[region.Region]int
}

func (v *PatchWithStartTime) GetRegionStartDate(region region.Region) time.Time {
	// Add the shift to the start date
	shift, ok := v.Shifts[region]
	if !ok {
		return v.StartTime
	}

	// Shift is in seconds
	return v.StartTime.Add(time.Duration(shift) * time.Second)
}

// Helper functions for creating patches from strings
func NewPatchFromString(s string) (Patch, error) {
	var p Patch
	err := p.FromString(s)
	return p, err
}

func NewShortPatchFromString(s string) (ShortPatch, error) {
	var p ShortPatch
	err := p.FromString(s)
	return p, err
}
