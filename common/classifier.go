package common

import (
	"math"
	"math/rand"
)

const Tiny = 0.0000001

type Class byte

const (
	Ham Class = iota
	Spam
	MAX_CLASS
)

func (c Class) String() string {
	switch c {
	case Ham:
		return "Ham"
	case Spam:
		return "Spam"
	default:
		panic("HELP")
	}
}

// Sample is a tuple representing a classification example
type Sample struct {
	Document []string
	Class
}

func shuffle(a []Sample) {
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func Argmax(a [MAX_CLASS]float64) Class {
	max := math.Inf(-1)
	var maxClass Class
	for i := Ham; i < MAX_CLASS; i++ {
		score := a[i]
		if score > max {
			maxClass = i
			max = score
		}
	}
	return maxClass
}

type Doc []int

func (d Doc) IDs() []int { return d }

type Classifier interface {
	Train(ex []Sample)
	Predict(s []string) Class
	Unseens(s []string) int
}
