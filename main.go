package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jszwec/csvutil"
	log "github.com/sirupsen/logrus"
)

type IRecord interface {
	Model
}

func readCsvChan[T IRecord](fileName string) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		f, err := os.Open(fileName)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		dec, err := csvutil.NewDecoder(csv.NewReader(f))
		if err != nil {
			log.Fatalln(err)
		}
		for i := 0; true; i++ {
			var r T
			if err := dec.Decode(&r); err == io.EOF {
				break
			}
			out <- r
		}
	}()
	return out
}

func readCsvChanOld[T IRecord](fileName string) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		records, err := readCsvFile[T](fileName)
		if err != nil {
			log.Fatalln(err)
		}

		for _, r := range records {
			out <- r
		}
	}()
	return out
}

func readCsvFile[T IRecord](name string) (res []T, err error) {
	if name == "" {
		return res, errors.New("empty input file")
	}
	b, err := os.ReadFile(name)
	if err != nil {
		return res, err
	}
	fileBytes, err := ioutil.ReadAll(bytes.NewReader(b))
	if err != nil {
		return res, err
	}

	err = gocsv.UnmarshalBytes(fileBytes, &res)
	return res, err
}
