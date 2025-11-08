package utils

import (
	"math/rand"
	"time"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(e T) Set[T] {
	s[e] = struct{}{}
	return s
}

func (s Set[T]) Remove(e T) Set[T] {
	delete(s, e)
	return s
}

func (s Set[T]) Has(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) Size() int {
	return len(s)
}

func Repeat(element string, count int) []string {
	repeated := make([]string, count)
	for i := range repeated {
		repeated[i] = element
	}
	return repeated
}

func Shuffle(slice []string) []string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	shuffled := make([]string, len(slice))
	copy(shuffled, slice)

	for i := len(shuffled) - 1; i > 0; i-- {
		j := r.Intn(i + 1)                                  // gen random index
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i] // then swap
	}

	return shuffled
}

func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for elem := range s {
		slice = append(slice, elem)
	}
	return slice
}

// func Pop(slice *[]string) (string, []string, error) {
// 	if len(*slice) == 0 {
// 		return "", *slice, errors.New("slice is empty")
// 	}

// 	lastIndex := len(*slice) - 1
// 	lastElement := (*slice)[lastIndex]
// 	*slice = (*slice)[:lastIndex]

// 	return lastElement, *slice, nil
// }
