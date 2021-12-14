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
			},
		},
	}
	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			got, err := ReadCSVFile(test.filename)
			if err != nil {
				t.Errorf("received error reading CSV file: %s", err)
			}
			assertString(t, "original filename", got.OriginalFilename, test.want.OriginalFilename)
			assertString(t, "title", got.Title, test.want.Title)
			assertString(t, "model", got.Model, test.want.Model)
			assertString(t, "s/n", got.SerialNum, test.want.SerialNum)
			assertFloat64(t, "center freq", got.CenterFreq, test.want.CenterFreq, 0.01)
			assertFloat64(t, "span", got.Span, test.want.Span, 0.01)
			assertFloat64(t, "rbw", got.RBW, test.want.RBW, 0.01)
			assertFloat64(t, "vbw", got.VBW, test.want.VBW, 0.01)
			assertFloat64(t, "ref level", got.RefLevel, test.want.RefLevel, 0.0000001)
			assertFloat64(t, "sweep time", got.SweepTime, test.want.SweepTime, 0.00000001)
			assertInt(t, "num points", got.NumPoints, test.want.NumPoints)
		})
	}
}

func assertInt(t *testing.T, label string, got, want int) {
	if got != want {
		t.Errorf("\ngot %s  = %d\nwant %s = %d", label, got, label, want)
	}
}

func assertFloat64(t *testing.T, label string, got, want, tolerance float64) {
	if diff := math.Abs(want - got); diff >= tolerance {
		t.Errorf("\ngot %s  = %#v \t\nwant %s = %#v", label, got, label, want)
	}
}

func assertBool(t *testing.T, label string, got, want bool) {
	if got != want {
		t.Errorf("\ngot = %t %s\nwant = %t", got, label, want)
	}
}

func assertString(t *testing.T, label string, got, want string) {
	if got != want {
		t.Errorf("\ngot  = `%#v` for %s\nwant = `%#v`", got, label, want)
	}
}
