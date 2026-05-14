package main

import (
	"fmt"
	"errors"
	"math/rand/v2"
)

type Stack struct {
	elements []any
	typeName string
	isSorted bool
}

func NewStack(values ...any) (*Stack, error) {
	if len(values) < 2 {
		return nil, errors.New("no values")
	}

	typeName, ok := values[0].(string)
	if !ok {
		return nil, errors.New("first must be string")
	}

	switch typeName {
	case "int":
		for _, v := range values[1:] {
			if _, ok := v.(int); !ok {
				return nil, errors.New("value is not int")
			}
		}

	case "float":
		for _, v := range values[1:] {
			if _, ok := v.(float64); !ok {
				return nil, errors.New("value is not float")
			}
		}

	case "string":
		for _, v := range values[1:] {
			if _, ok := v.(string); !ok {
				return nil, errors.New("value is not string")
			}
		}

	case "rune":
		for _, v := range values[1:] {
			if _, ok := v.(rune); !ok {
				return nil, errors.New("value is not rune")
			}
		}

	default:
		return nil, errors.New("wrong type")
	}


	return &Stack{
		elements: values[1:],
		typeName: typeName,
		isSorted: false,
	}, nil
}

func NewSortedStack(values ...any) (*Stack, error) {
	s, err := NewStack(values...)
    if err != nil {
        return nil, err
    }

    elements := values[1:]
    for i := 0; i <len(elements)-1; i++ {
        switch s.typeName {
        case "int":
            if elements[i].(int) > elements[i+1].(int) {
                return nil, errors.New("not sorted")
            }
        case "float":
            if elements[i].(float64) > elements[i+1].(float64) {
                return nil, errors.New("not sorted")
            }
        case "string":
            if elements[i].(string) > elements[i+1].(string) {
                return nil, errors.New("not sorted")
            }
        case "rune":
            if elements[i].(rune) > elements[i+1].(rune) {
                return nil, errors.New("not sorted")
            }
        }
    }

    s.isSorted = true
    return s, nil
}

func (s *Stack) Push(value any) error {
	if fmt.Sprintf("%T", value) != s.typeName{
                return errors.New("wrong type")
	} else {
		s.elements = append(s.elements, value)
	}
	return nil
}

func (s *Stack) Pop() error {
	if len(s.elements) < 1{
        return errors.New("already empty")
    }
    s.elements = s.elements[:len(s.elements)-1]
	return nil
}

func (s *Stack) Size() int {
    size := len(s.elements)
	return size
}

func randomStackSize() int {
	N := 100
	randRange := 100
	i := 0
	sizesArr := []int{}

	for i < N {
		randomNumber := rand.IntN(randRange)
		s, err := NewSortedStack("int", randomNumber)
		if err != nil {
			fmt.Println(err)
		}

		for j := 0; j < N-1; j++ {
            nextNum := rand.IntN(randRange)
            lastIdx := len(s.elements) - 1
            if val, ok := s.elements[lastIdx].(int); ok {
                if nextNum > val {
                    s.Push(nextNum)
                }
            }
        }
		// fmt.Println(s)
		sizesArr = append(sizesArr, s.Size())
		i++
	}
	
	sizeSum := 0
	for i = 0; i < N; i++ {
		sizeSum += sizesArr[i]
	}

	avgSize := sizeSum/ N
	return avgSize
}

// func bracketCheck(input string) bool {
// 	pairs := map[rune]rune {')': '(', ']': '[', '}': '{', '>': '<'}
	
// 	return true
// }

func main(){
	s, err := NewStack("int", 2)
	if err == nil {
		fmt.Println(s)
	} else{
		fmt.Println(err)
	}

	sSorted, err := NewSortedStack("int", 2, 5)
	if err == nil {
		fmt.Println(sSorted)
	} else{
		fmt.Println(err)
	}

	err = sSorted.Push(7)
	if err == nil {
		fmt.Println(sSorted)
	} else{
		fmt.Println(err)
	}

	err = s.Pop()
	if err == nil {
		fmt.Println(s)
	} else{
		fmt.Println(err)
	}

	err = s.Pop()
	if err == nil {
		fmt.Println(s)
	} else{
		fmt.Println(err)
	}

	stackSize := sSorted.Size()
	fmt.Println("Size of sorted stack is ", stackSize)

	avgStackSize := randomStackSize()
	fmt.Println("Average random sorted stack size: ", avgStackSize)

	// brackets := "(()())"
}