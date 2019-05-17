package main

import "sync"

/*
 * Set
 * kind of data structure, of which the elements are non-repeated,
 * struct is thread safe
 */
type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

/*
 * generate a new set
 */
func NewSet() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}

/*
 * add an item to the set
 */
func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

/*
 * remove an item from the set
 */
func (s *Set) Remove(item interface{}) {
	s.Lock()
	s.Unlock()
	delete(s.m, item)
}

/*
 * judge whether the set contains a given item
 */
func (s *Set) Contains(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

/*
 * get the length of the set
 */
func (s *Set) Len() int {
	return len(s.List())
}

/*
 * clear the set
 */
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]bool{}
}

/*
 * judge whether the set is empty
 */
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

/*
 * transform the set into a list
 */
func (s *Set) List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	list := []interface{}{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
