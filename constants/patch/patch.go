package patch

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Kinveil/Riot-API-Golang/constants/region"
)

type ShortPatch struct {
	Major int
	Minor int
}

func (v ShortPatch) String() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

// FromString creates a ShortPatch from a string like "13.2"
func (v *ShortPatch) FromString(s string) error {
	parts := strings.Split(s, ".")
	if len(parts) != 2 {
		return fmt.Errorf("invalid short patch format: %s", s)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	v.Major = major
	v.Minor = minor
	return nil
}

type Patch struct {
	Major int
	Minor int
	Patch int
}

func (v Patch) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// FromString creates a Patch from a string like "13.2.1"
func (v *Patch) FromString(s string) error {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return fmt.Errorf("invalid patch format: %s", s)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return err
	}

	v.Major = major
	v.Minor = minor
	v.Patch = patch
	return nil
}

// ShortPatch returns the short version of the patch (removes patch component)
func (v Patch) ShortPatch() ShortPatch {
	return ShortPatch{
		Major: v.Major,
		Minor: v.Minor,
	}
}

// Compare returns -1 if v < other, 0 if v == other, 1 if v > other
func (v Patch) Compare(other Patch) int {
	if v.Major != other.Major {
		if v.Major < other.Major {
			return -1
		}
		return 1
	}

	if v.Minor != other.Minor {
		if v.Minor < other.Minor {
			return -1
		}
		return 1
	}

	if v.Patch != other.Patch {
		if v.Patch < other.Patch {
			return -1
		}
		return 1
	}

	return 0
}

// Compare returns -1 if v < other, 0 if v == other, 1 if v > other
func (v ShortPatch) Compare(other ShortPatch) int {
	if v.Major != other.Major {
		if v.Major < other.Major {
			return -1
		}
		return 1
	}

	if v.Minor != other.Minor {
		if v.Minor < other.Minor {
			return -1
		}
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
