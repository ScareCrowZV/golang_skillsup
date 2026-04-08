package main

import "fmt"

type hashmap struct {
	m []string
}

func (h *hashmap) Set(key, val string) {
	h.m[h.hashstr(key)] = val
}

func (h *hashmap) Get(key string) (value string, ok bool) {

	if h.m[h.hashstr(key)] == "" {
		return "", false
	} else {
		return h.m[h.hashstr(key)], true
	}
}

func (h *hashmap) Delete(key string) {
	h.m[h.hashstr(key)] = ""
}

func (h *hashmap) hashstr(val string) (res uint64) {

	for i, v := range val {
		res += (uint64(v) * uint64(i))
	}

	return res % 1000
}

func main() {
	var hm hashmap
	hm.m = make([]string, 1000)
	hm.Set("test", "test")
	fmt.Println(hm.Get("test"))
	hm.Delete("test")
	fmt.Println(hm.Get("test"))
}
