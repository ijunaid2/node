/*
 * Copyright (C) 2021 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package random

import (
	"math/rand"
	"testing"
)

type mockRandomSource int64

func (s *mockRandomSource) Seed(seed int64) {
	*s = mockRandomSource(seed)
}

func (s *mockRandomSource) Int63() int64 {
	return int64(*s)
}

type mockRandomSource64 int64

func (s *mockRandomSource64) Seed(seed int64) {
	*s = mockRandomSource64(seed)
}

func (s *mockRandomSource64) Int63() int64 {
	return int64(*s)
}

func (s *mockRandomSource64) Uint64() uint64 {
	return uint64(*s)
}

func TestVirtualConstructor(t *testing.T) {
	src := new(mockRandomSource)
	src64 := new(mockRandomSource64)

	cSrc := NewConcurrentRandomSource(src)
	cSrc64 := NewConcurrentRandomSource(src64)

	// Generated by fair dice roll
	var seed int64 = 4
	cSrc.Seed(seed)
	cSrc64.Seed(seed)

	_, ok := cSrc.(rand.Source64)
	if ok {
		t.Errorf("%T exposes unexpected interface", cSrc)
	}

	cSrc64Extended, ok := cSrc64.(rand.Source64)
	if !ok {
		t.Errorf("%T exposes unexpected interface", cSrc64)
	}

	if cSrc.Int63() != seed {
		t.Errorf("Wrapped %T returned unexpected value", cSrc)
	}

	if cSrc64.Int63() != seed {
		t.Errorf("Wrapped %T returned unexpected value", cSrc)
	}

	if cSrc64Extended.Uint64() != uint64(seed) {
		t.Errorf("Wrapped %T returned unexpected value", cSrc64Extended)
	}
}