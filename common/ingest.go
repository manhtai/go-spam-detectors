package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type errList []error

func (err errList) Error() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Errors Found:\n")
	for _, e := range err {
		fmt.Fprintf(&buf, "\t%v\n", e)
	}
	return buf.String()
}

func LoadSamples(typ string) ([]Sample, []Sample) {
	total, err := ingest(typ)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total loaded: %d\n", len(total))
	shuffle(total)
	cvStart := len(total) - len(total)/3

	return total[:cvStart], total[cvStart:]
}

func ingest(typ string) (samples []Sample, err error) {
	switch typ {
	case "bare", "lemm", "lemm_stop", "stop":
	default:
		return nil, errors.Errorf("Expected only \"bare\", \"lemm\", \"lemm_stop\" or \"stop\"")
	}

	var errs errList
	start, end := 0, 11

	for i := start; i < end; i++ {
		matches, err := filepath.Glob(fmt.Sprintf("./lingspam_public/%s/part%d/*.txt", typ, i))
		if err != nil {
			errs = append(errs, err)
			continue
		}

		for _, match := range matches {
			str, err := ingestOneFile(match)
			if err != nil {
				errs = append(errs, errors.WithMessage(err, match))
				continue
			}

			if strings.Contains(match, "spmsg") { // is spam
				samples = append(samples, Sample{str, Spam})
			} else { // is ham
				samples = append(samples, Sample{str, Ham})
			}
		}
	}
	if errs != nil {
		err = errs
	}
	return
}

func ingestOneFile(abspath string) ([]string, error) {
	bs, err := ioutil.ReadFile(abspath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bs), " "), nil
}
