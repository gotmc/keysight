// Copyright (c) 2021-2022 The keysight developers. All rights reserved.
// Project site: https://github.com/gotmc/keysight
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package esa

import (
	"math"
	"testing"
)

func TestReadCSVFile(t *testing.T) {
	var tests = []struct {
		filename string
		want     Trace
	}{
		{
			filename: "./testdata/e4402b_trace924.csv",
			want: Trace{
				OriginalFilename: "C:\\TRACE924.CSV",
				Title:            "",
				Model:            "E4402B",
				SerialNum:        "MY45104598",
				CenterFreq:       34000.0,
				Span:             50000.0,
				RBW:              1000.0,
				VBW:              1000.0,
				RefLevel:         106.99,
				SweepTime:        0.085,
				NumPoints:        401,
				FreqLabel:        "",
				Trace1Label:      "Trace 1",
				Trace2Label:      "Trace 2",
				Trace3Label:      "Trace 3",
				FreqUnits:        "Hz",
				Trace1Units:      "dBuV",
				Trace2Units:      "dBuV",
				Trace3Units:      "dBuV",
				Trace1:           []float64{5.90097e+01, 5.92727e+01, 5.90557e+01},
			},
		},
		{
			filename: "./testdata/e4411b_trace080.csv",
			want: Trace{
				OriginalFilename: "A:\\TRACE080.CSV",
				Title:            "",
				Model:            "E4411B",
				SerialNum:        "MY45104634",
				CenterFreq:       750000000.0,
				Span:             500000000.0,
				RBW:              100000.0,
				VBW:              100000.0,
				RefLevel:         73.0103,
				SweepTime:        0.0644205,
				NumPoints:        401,
				FreqLabel:        "",
				Trace1Label:      "Trace 1",
				Trace2Label:      "Trace 2",
				Trace3Label:      "Trace 3",
				FreqUnits:        "Hz",
				Trace1Units:      "",
				Trace2Units:      "",
				Trace3Units:      "",
				Trace1:           []float64{3.7123, 3.3353, 3.3773},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			got, err := ReadCSVFile(test.filename)
			if err != nil {
				t.Errorf("received error reading CSV file: %s", err)
			}
			assert(t, "original filename", got.OriginalFilename, test.want.OriginalFilename)
			assert(t, "title", got.Title, test.want.Title)
			assert(t, "model", got.Model, test.want.Model)
			assert(t, "s/n", got.SerialNum, test.want.SerialNum)
			assertFloat64(t, "center freq", got.CenterFreq, test.want.CenterFreq, 0.01)
			assertFloat64(t, "span", got.Span, test.want.Span, 0.01)
			assertFloat64(t, "rbw", got.RBW, test.want.RBW, 0.01)
			assertFloat64(t, "vbw", got.VBW, test.want.VBW, 0.01)
			assertFloat64(t, "ref level", got.RefLevel, test.want.RefLevel, 0.0000001)
			assertFloat64(t, "sweep time", got.SweepTime, test.want.SweepTime, 0.00000001)
			assert(t, "num points", got.NumPoints, test.want.NumPoints)
			assert(t, "freq label", got.FreqLabel, test.want.FreqLabel)
			assert(t, "trace 1 label", got.Trace1Label, test.want.Trace1Label)
			assert(t, "trace 2 label", got.Trace2Label, test.want.Trace2Label)
			assert(t, "trace 3 label", got.Trace3Label, test.want.Trace3Label)
			assert(t, "freq units", got.FreqUnits, test.want.FreqUnits)
			assert(t, "trace 1 units", got.Trace1Units, test.want.Trace1Units)
			assert(t, "trace 2 units", got.Trace2Units, test.want.Trace2Units)
			assert(t, "trace 3 units", got.Trace3Units, test.want.Trace3Units)
			assertFloat64(t, "t1[0]", got.Trace1[0], test.want.Trace1[0], 0.00000001)
			assertFloat64(t, "t1[1]", got.Trace1[1], test.want.Trace1[1], 0.00000001)
			// Check that the frequency and trace data have the correct number of
			// data points.
			assert(t, "freq len", len(got.Frequency), test.want.NumPoints)
			assert(t, "trace 1 len", len(got.Trace1), test.want.NumPoints)
			assert(t, "trace 2 len", len(got.Trace2), test.want.NumPoints)
			assert(t, "trace 3 len", len(got.Trace3), test.want.NumPoints)
		})
	}
}

func assert(t *testing.T, label string, got, want interface{}) {
	if got != want {
		t.Errorf("\ngot  = `%#v` for %s\nwant = `%#v`", got, label, want)
	}
}

func assertFloat64(t *testing.T, label string, got, want, tolerance float64) {
	if diff := math.Abs(want - got); diff >= tolerance {
		t.Errorf("\ngot %s  = %#v \t\nwant %s = %#v", label, got, label, want)
	}
}
