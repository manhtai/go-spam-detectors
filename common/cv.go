package common

func CrossValidation(train, test []Sample, classifier Classifier) (hams, spams, corrects, totals int) {
	classifier.Train(train)
	var unseen, totalWords int

	for _, ex := range test {
		totalWords += len(ex.Document)
		unseen += classifier.Unseens(ex.Document)
		class := classifier.Predict(ex.Document)
		if class == ex.Class {
			corrects++
		}
		switch ex.Class {
		case Ham:
			hams++
		case Spam:
			spams++
		}
		totals++
	}

	return
}
