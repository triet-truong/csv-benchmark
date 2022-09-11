package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

const testFilePath string = "data.csv"

func BenchmarkReadCsvChan(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for r := range readCsvChan[Model](testFilePath) {
			logrus.Debugln(r)
		}
		p.Next()
	})
}

func BenchmarkReadCsvChanOld(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for r := range readCsvChanOld[Model](testFilePath) {
			logrus.Debugln(r)
		}
		p.Next()
	})

}
