package main

import "fmt"

type Node struct {
	Key   string // should be an interface btw
	Value interface{}
	next  *Node // this is not exported at all
}

// we have a bunch of stuff
type HashMap struct {
	// we have a bunch of nodes
	elements  []*Node
	window    int
	size      int // actual size of the hashmap
	threshold float32

	// we should have a mapping to and fro
}

func NewHashMap() *HashMap {
	const window int = 100

	return &HashMap{
		elements: make([]*Node, window),
		size:     window,
		window:   window,
	}
}

// A stupid hash key implemenation
func (hashMap *HashMap) hashKey(key string) int {
	var ordinate_total int

	for i := 0; i < len(key); i++ {
		ordinate_total += int(key[i])
	}

	return ordinate_total % hashMap.size
}

func (hashMap *HashMap) Push(key string, value interface{}) {
	// check if the threshold is over .7 if so copy the underlying array over
	if hashMap.threshold >= 0.7 {
		new_element_bucket_size := hashMap.window + hashMap.size
		new_element_bucket := make([]*Node, new_element_bucket_size)
		copy(new_element_bucket, hashMap.elements)
		hashMap.elements = new_element_bucket
		hashMap.size = new_element_bucket_size

	}

	hashed_key_index := hashMap.hashKey(key)
	element_found_in_position := hashMap.elements[hashed_key_index]

	if element_found_in_position == nil {
		hashMap.elements[hashed_key_index] = &Node{
			Key:   key,
			Value: value,
		}

		return
	}

	// ( use chaining for collision resolution )
	for element_found_in_position.next != nil {
		if element_found_in_position.Key == key {
			element_found_in_position.Value = value
			goto threshold_calculation
		}

		element_found_in_position = element_found_in_position.next
	}

	element_found_in_position.next = &Node{
		Key:   key,
		Value: value,
	}

threshold_calculation:
	hashMap.threshold = float32(len(hashMap.elements)) / float32(hashMap.size)
}

func (hashMap *HashMap) Get(key string) (interface{}, error) {
	hashed_key_index := hashMap.hashKey(key)
	element_found_in_position := hashMap.elements[hashed_key_index]

	if element_found_in_position == nil {
		goto abnormal_exit
	}

	for element_found_in_position != nil {
		if element_found_in_position.Key == key {
			return element_found_in_position.Value, nil
		}

		element_found_in_position = element_found_in_position.next
	}

abnormal_exit:
	return nil, fmt.Errorf("%s does not exist", key)
}
