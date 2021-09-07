package main

import (
	"github.com/chewxy/lingo/corpus"
	"github.com/go-nlp/tfidf"
	"go-spam-detector/common"
	"math"
	"sync"
)

type bayesClassifier struct {
	corpus *corpus.Corpus

	tfidfs [common.MAX_CLASS]*tfidf.TFIDF
	totals [common.MAX_CLASS]float64

	ready bool
	sync.Mutex
}

func NewBayesClassifier() common.Classifier {
	var tfidfs [common.MAX_CLASS]*tfidf.TFIDF
	for i := common.Ham; i < common.MAX_CLASS; i++ {
		tfidfs[i] = tfidf.New()
	}
	return &bayesClassifier{
		corpus: corpus.New(),
		tfidfs: tfidfs,
	}
}

func (c *bayesClassifier) Train(samples []common.Sample) {
	for _, sam := range samples {
		c.trainOne(sam)
	}
}

func (c *bayesClassifier) Postprocess() {
	c.Lock()
	if c.ready {
		return
	}

	var docs int
	for _, t := range c.tfidfs {
		docs += t.Docs
	}
	for _, t := range c.tfidfs {
		t.Docs = docs
		// t.CalculateIDF()
		for k, v := range t.TF {
			t.IDF[k] = math.Log1p(float64(t.Docs) / v)
		}
	}
	c.ready = true
	c.Unlock()
}

func (c *bayesClassifier) Score(sentence []string) (scores [common.MAX_CLASS]float64) {
	if !c.ready {
		c.Postprocess()
	}

	d := make(common.Doc, len(sentence))
	for i, word := range sentence {
		id := c.corpus.Add(word)
		d[i] = id
	}

	priors := c.priors()

	// score per class
	for i := range c.tfidfs {
		score := math.Log(priors[i])
		// likelihood
		for _, word := range sentence {
			prob := c.prob(word, common.Class(i))
			score += math.Log(prob)
		}

		scores[i] = score
	}
	return
}

func (c *bayesClassifier) Predict(sentence []string) common.Class {
	scores := c.Score(sentence)
	return common.Argmax(scores)
}

func (c *bayesClassifier) Unseens(sentence []string) (retVal int) {
	for _, word := range sentence {
		if _, ok := c.corpus.Id(word); !ok {
			retVal++
		}
	}
	return
}

func (c *bayesClassifier) trainOne(example common.Sample) {
	d := make(common.Doc, len(example.Document))
	for i, word := range example.Document {
		id := c.corpus.Add(word)
		d[i] = id
	}
	c.tfidfs[example.Class].Add(d)
	c.totals[example.Class]++
}

func (c *bayesClassifier) priors() (priors []float64) {
	priors = make([]float64, common.MAX_CLASS)
	var sum float64
	for i, total := range c.totals {
		priors[i] = total
		sum += total
	}
	for i := common.Ham; i < common.MAX_CLASS; i++ {
		priors[int(i)] /= sum
	}
	return
}

func (c *bayesClassifier) prob(word string, class common.Class) float64 {
	id, ok := c.corpus.Id(word)
	if !ok {
		return common.Tiny
	}

	freq := c.tfidfs[class].TF[id]
	idf := c.tfidfs[class].IDF[id]

	// a word may not appear at all in a class.
	if freq == 0 {
		return common.Tiny
	}

	return freq * idf / c.totals[class]
}
