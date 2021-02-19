package common

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// StringInSlice is a helper function which iterates through a slice and checks for a string
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// StringInSliceIgnoreCase is a helper function which iterates through a slice and checks for a string
func StringInSliceIgnoreCase(a string, list []string) bool {
	for _, b := range list {
		if strings.EqualFold(a, b) {
			return true
		}
	}
	return false
}

// Int64InSlice is a helper function which iterates through a slice and checks for an int64
func Int64InSlice(a int64, list []int64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// StringsMinus returns all UNIQUE strings that are in a but not in b.
// Returned list does NOT preserve original order in slice.
func StringsMinus(a, b []string) []string {
	aMap := make(map[string]bool)
	for _, e := range a {
		aMap[e] = true
	}

	bMap := make(map[string]bool)
	for _, e := range b {
		bMap[e] = true
	}

	for bKey := range bMap {
		delete(aMap, bKey)
	}

	result := make([]string, 0)

	for aKey := range aMap {
		result = append(result, aKey)
	}

	return result
}

// StringToInt64Slice converts a slice of strings to a slice of integers
// throwing an error if the input contains invalid integers.
func StringToInt64Slice(strings []string) ([]int64, error) {
	nums := make([]int64, len(strings))
	for idx, s := range strings {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not convert: %s to int. error: %+v", s, err))
		}

		nums[idx] = i
	}

	return nums, nil
}

// Int64ToStringSlice converts a slice of integers to a slice of strings
func Int64ToStringSlice(ints []int64) []string {
	nums := make([]string, len(ints))
	for idx, i := range ints {
		s := strconv.FormatInt(i, 10)
		nums[idx] = s
	}
	return nums
}

// FilterStrings filters out empty strings from a list
func FilterStrings(strings []string) []string {
	result := make([]string, 0, len(strings))
	for _, s := range strings {
		if s != "" {
			result = append(result, s)
		}
	}

	return result
}

// IndexOfStringInSlice returns the index of a string in a slice
// similar to Array.indexOf in Java / Javascript
//
// It returns an error if the string is not found.
func IndexOfStringInSlice(haystack []string, needle string) (int, error) {
	for i, val := range haystack {
		if val == needle {
			return i, nil
		}
	}

	return -1, errors.New(fmt.Sprintf("%s not found in slice", needle))
}

// IntInSlice returns true if the given item
// exists in the given string array, or false otherwise.
func IntInSlice(arr []int, input int) bool {
	for _, cur := range arr {
		if cur == input {
			return true
		}
	}
	return false
}
