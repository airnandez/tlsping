package tlsping

import (
	"fmt"
	"math"
)

// PingResults contains summary statistics of the measured connection
// times
type PingResult struct {
	// Number of measurements summarized in this result
	Count int

	// Minimum and maximum observed connection times, in seconds
	Min, Max float64

	// Average and standard deviation of the observed connection
	// times, in seconds
	Avg, Std float64
}

// summarize summarizes the measurements of time durations given as
// argument. The argument values and the returned values are understood
// in seconds
func summarize(durations []float64) PingResult {
	min, max := math.MaxFloat64, math.SmallestNonzeroFloat64
	sum := float64(0)
	for _, d := range durations {
		sum += d
		if d < min {
			min = d
		}
		if d > max {
			max = d
		}
	}
	n := float64(len(durations))
	avg := sum / n
	std := float64(0)
	for _, d := range durations {
		dev := d - avg
		std += dev * dev
	}
	return PingResult{
		Count: len(durations),
		Min:   min,
		Max:   max,
		Avg:   avg,
		Std:   math.Sqrt(std / n),
	}
}

func (r *PingResult) MinStr() string {
	return secsToString(r.Min)
}

func (r *PingResult) MaxStr() string {
	return secsToString(r.Max)
}

func (r *PingResult) AvgStr() string {
	return secsToString(r.Avg)
}

func (r *PingResult) StdStr() string {
	return secsToString(r.Std)
}

func secsToString(secs float64) string {
	if secs >= 1.0 {
		// unit is seconds
		return fmt.Sprintf("%.2fs", secs)
	}
	secs = secs * 1000.0
	if secs >= 1.0 {
		// unit is milliseconds
		return fmt.Sprintf("%.2fms", secs)
	}
	secs = secs * 1000.0
	if secs >= 1.0 {
		// unit is µ seconds
		return fmt.Sprintf("%.2fµs", secs)
	}
	// unit is nano seconds
	secs = secs * 1000.0
	return fmt.Sprintf("%.2fns", secs)
}
