package main

import (
	"go-spam-detectors/common"
)

func main() {
	classifier := NewBayes1Classifier()
	common.Summary(classifier)
}
