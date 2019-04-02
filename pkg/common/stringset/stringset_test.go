// Copyright (c) 2019 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stringset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testItem = "testItem"
)

func TestStringSet_New(t *testing.T) {
	testSet := New()
	assert.NotNil(t, testSet)
}

func TestStringSet_Add(t *testing.T) {
	// Create a new StringSet
	testSet := &stringSet{
		m: make(map[string]bool),
	}
	// Add test task to StringSet
	testSet.Add(testItem)
	assert.Equal(t, true, testSet.m[testItem])
}

func TestStringSet_Contains(t *testing.T) {
	// Create a new StringSet
	testSet := &stringSet{
		m: make(map[string]bool),
	}
	assert.Equal(t, false, testSet.Contains(testItem))

	// Add test task to StringSet
	testSet.m[testItem] = true
	assert.Equal(t, true, testSet.Contains(testItem))
}

func TestStringSet_Remove(t *testing.T) {
	// Create a new StringSet
	testSet := &stringSet{
		m: make(map[string]bool),
	}
	// Add test task to StringSet
	testSet.m[testItem] = true
	assert.Equal(t, true, testSet.m[testItem])

	testSet.Remove(testItem)
	assert.Equal(t, false, testSet.m[testItem])
}

func TestStringSet_Clear(t *testing.T) {
	// Create a new StringSet
	testSet := &stringSet{
		m: make(map[string]bool),
	}
	// Add test task to StringSet
	testSet.m[testItem] = true
	assert.Equal(t, 1, len(testSet.m))

	testSet.Clear()
	assert.Equal(t, 0, len(testSet.m))
}

func TestStringSet_ToSlice(t *testing.T) {
	// Create a new StringSet
	testSet := &stringSet{
		m: make(map[string]bool),
	}

	testItems := []string{
		"testitem1",
		"testitem2",
	}
	// Add testItems to StringSet
	for _, item := range testItems {
		testSet.m[item] = true
	}
	items := testSet.ToSlice()
	assert.Len(t, items, len(testItems))
	for _, item := range items {
		assert.True(t, testSet.Contains(item))
	}
}

func TestStringSet_Intersect(t *testing.T) {
	testSet := &stringSet{
		m: make(map[string]bool),
	}

	testSet1 := &stringSet{
		m: make(map[string]bool),
	}

	testItems := []string{
		"testitem1",
		"testitem2",
	}

	// Add testItems to testSet
	for _, item := range testItems {
		testSet.m[item] = true
	}

	testItems1 := []string{
		"testitem1",
		"testitem2",
		"testitem3",
	}

	// Add testItems1 to testSet1
	for _, item := range testItems1 {
		testSet1.m[item] = true
	}

	// Intersect the two stringSet
	intersection := testSet.Intersect(testSet1)

	items := intersection.ToSlice()

	for _, item := range items {
		if item == "testitem3" {
			assert.False(t, testSet.Contains(item))
		}
		assert.True(t, testSet.Contains(item))
	}
}

func TestStringSet_Intersect_VoidIntersect(t *testing.T) {
	testSet := &stringSet{
		m: make(map[string]bool),
	}

	testSet1 := &stringSet{
		m: make(map[string]bool),
	}

	testItems := []string{
		"testitem1",
		"testitem2",
	}

	// Add testItems to testSet
	for _, item := range testItems {
		testSet.m[item] = true
	}

	testItems1 := []string{
		"testitem1",
		"testitem2",
	}

	// Add testItems1 to testSet1
	for _, item := range testItems1 {
		testSet1.m[item] = true
	}

	// Intersect the two stringSet
	intersection := testSet.Intersect(testSet1)

	items := intersection.ToSlice()

	if ok := len(items) == 0; ok {
		assert.True(t, ok)
	}
}

func TestStringSet_Intersect_VoidIntersect_Second(t *testing.T) {
	testSet := &stringSet{
		m: make(map[string]bool),
	}

	testSet1 := &stringSet{
		m: make(map[string]bool),
	}

	testItems := []string{}

	// Add testItems to testSet
	for _, item := range testItems {
		testSet.m[item] = true
	}

	testItems1 := []string{}

	// Add testItems1 to testSet1
	for _, item := range testItems1 {
		testSet1.m[item] = true
	}

	// Intersect the two stringSet
	intersection := testSet.Intersect(testSet1)

	items := intersection.ToSlice()

	if ok := len(items) == 0; ok {
		assert.True(t, ok)
	}
}

func TestStringSet_Intersect_VoidIntersect_Third(t *testing.T) {
	testSet := &stringSet{
		m: make(map[string]bool),
	}

	testSet1 := &stringSet{
		m: make(map[string]bool),
	}

	testItems := []string{
		"testitem1",
		"testitem2",
	}

	// Add testItems to testSet
	for _, item := range testItems {
		testSet.m[item] = true
	}

	testItems1 := []string{
		"testitem3",
		"testitem4",
	}

	// Add testItems1 to testSet1
	for _, item := range testItems1 {
		testSet1.m[item] = true
	}

	// Intersect the two stringSet
	intersection := testSet.Intersect(testSet1)

	items := intersection.ToSlice()

	if ok := len(items) == 0; ok {
		assert.True(t, ok)
	}
}
