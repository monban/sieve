package main

import "testing"

func TestAdding(t *testing.T) {
	bs := NewBoolStore(8)
	bs.Set(3)
	if bs.store[0] != 8 {
		t.Log(bs)
		t.Error("bit not set")
	}
}

func TestGetting(t *testing.T) {
	bs := NewBoolStore(8)
	bs.Set(3)

	if bs.Get(3) == false {
		t.Log(bs)
		t.Fatal("bit not retreived")
	}
}

func TestLargeNumber(t *testing.T) {
	large_number := uint64(10_000)
	bs := NewBoolStore(large_number)
	bs.Set(large_number)

	if bs.Get(large_number) == false {
		t.Log(bs)
		t.Fatal("no work")
	}
}

func TestStringify(t *testing.T) {
	large_number := uint64(10_000)
	bs := NewBoolStore(large_number)
	for i := uint64(1); i < large_number; i += 2 {
		bs.Set(i)
	}
	t.Log("\n", bs)
}
