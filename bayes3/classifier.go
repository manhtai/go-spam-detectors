package main

import (
	"github.com/navossoc/bayesian"
	"go-spam-detectors/common"
)

type bayesClassifier struct {
	classifier *bayesian.Classifier
}

const (
	Spam bayesian.Class = "Spam"
	Ham  bayesian.Class = "Ham"
)

func NewBayes3Classifier() *bayesClassifier {
	classifier := bayesian.NewClassifierTfIdf(Spam, Ham)
	return &bayesClassifier{
		classifier: classifier,
	}
}

func fromClass(s common.Class) bayesian.Class {
	if s == common.Spam {
		return Spam
	}
	return Ham
}

func (b *bayesClassifier) Train(samples []common.Sample) {
	for _, sample := range samples {
		b.classifier.Learn(sample.Content, fromClass(sample.Class))
	}
	b.classifier.ConvertTermsFreqToTfIdf()
}

func (b *bayesClassifier) Predict(s []string) common.Class {
	_, likely, _ := b.classifier.LogScores(s)
	return []common.Class{common.Spam, common.Ham}[likely]
}
