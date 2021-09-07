package main

import (
	"go-spam-detector/common"
)

func main() {
	classifier := NewBayesClassifier()
	common.Summary(classifier)
}
