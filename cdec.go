package cdec

import "io"

func ScanFreqsFromReader(io.Reader)(FreqCount,error){

}

type FreqCount struct {
	Freqs map[string]float64
	secretvar int
}