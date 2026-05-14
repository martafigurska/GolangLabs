package main

import (
	"fmt"
	"errors"
)

type Iterator interface {
	Next() (any, error)
	Reset()
}



type Para struct {
	values [2]any
	state int
}

func NewPara(a, b any) *Para {
	return &Para{
		values: [2]any{a, b},
		state: 0,
	}
}

func (p *Para) Next() (any, error) {
	if p.state >= len(p.values) {
		return nil, errors.New("Finished")
	}
	val := p.values[p.state]
	p.state++
	return val, nil
}

func (p *Para) Reset() {
	p.state = 0
}



type Krotka struct {
	values []any
	state int
}

func NewKrotka(vals ...any) *Krotka {
	return &Krotka{
		values: vals,
		state: 0,
	}
}

func (k *Krotka) Next() (any, error) {
	if k.state >= len(k.values) {
		return nil, errors.New("Finished")
	}
	val := k.values[k.state]
	k.state++
	return val, nil
}

func (k *Krotka) Reset() {
	k.state = 0
}


func KrotkaIterator(k1, k2 *Krotka) func(func(*Para) bool) {
	return func(yield func(*Para) bool) {
		maxLen := len(k1.values)
		if len(k2.values) > maxLen {
			maxLen = len(k2.values)
		}
		for i:=0; i<maxLen; i++ {
			var v1, v2 any
			
			if i < len(k1.values) {
				v1 = k1.values[i]
			} else {
				v1 = nil
			}

			if i < len(k2.values) {
				v2 = k2.values[i]
			} else {
				v2 = nil
			}

			if !yield(NewPara(v1,v2)){
				return
			}
		}
	}
}



func main() {
	k1 := NewKrotka(1, "dwa", 3, "cztery", 5)
	k2 := NewKrotka(1, 2)
	fmt.Println(k1)
	fmt.Println(k2)

	fmt.Println("***************")
	for para := range KrotkaIterator(k1, k2) {
		fmt.Println(para)
	}
}