package main

import (
	"go-spam-detector/common"
)

func main() {
	classifier := NewBayes1Classifier()
	common.Summary(classifier)
}
