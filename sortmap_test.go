package sortmap

import (
	"testing"
)

func TestSortMap(t *testing.T) {
	strMap := New[string, any]()
	// number
	strMap.Set("number", 3)
	v, _ := strMap.Get("number")
	if v.(int) != 3 {
		t.Error("Set number")
	}
	// string
	strMap.Set("string", "x")
	v, _ = strMap.Get("string")
	if v.(string) != "x" {
		t.Error("Set string")
	}
	// string slice
	strMap.Set("strings", []string{
		"t",
		"u",
	})
	v, _ = strMap.Get("strings")
	if v.([]string)[0] != "t" {
		t.Error("Set strings first index")
	}
	if v.([]string)[1] != "u" {
		t.Error("Set strings second index")
	}
	// mixed slice
	strMap.Set("mixed", []any{
		1,
		"1",
	})
	v, _ = strMap.Get("mixed")
	if v.([]any)[0].(int) != 1 {
		t.Error("Set mixed int")
	}
	if v.([]any)[1].(string) != "1" {
		t.Error("Set mixed string")
	}
	// overriding existing key
	strMap.Set("number", 4)
	v, _ = strMap.Get("number")
	if v.(int) != 4 {
		t.Error("Override existing key")
	}
	// Keys method
	keys := strMap.Keys()
	expectedKeys := []string{
		"number",
		"string",
		"strings",
		"mixed",
	}
	for i, key := range keys {
		if key != expectedKeys[i] {
			t.Error("Keys method", key, "!=", expectedKeys[i])
		}
	}
	for i, key := range expectedKeys {
		if key != expectedKeys[i] {
			t.Error("Keys method", key, "!=", expectedKeys[i])
		}
	}
	// delete
	strMap.Delete("strings")
	strMap.Delete("not a key being used")
	if len(strMap.Keys()) != 3 {
		t.Error("Delete method")
	}
	_, ok := strMap.Get("strings")
	if ok {
		t.Error("Delete did not remove 'strings' key")
	}

	intMap := New[int, string]()
	intMap.Set(2, "a")
	v, _ = intMap.Get(2)
	if v != "a" {
		t.Error("Set 1")
	}

	intMap.Set(1, "b")
	v, _ = intMap.Get(1)
	if v != "b" {
		t.Error("Set 2")
	}

	intKeys := intMap.Keys()
	expectedIntKeys := []int{
		2, 1,
	}
	for i, key := range intKeys {
		if key != expectedIntKeys[i] {
			t.Error("Keys method", key, "!=", expectedKeys[i])
		}
	}
}
