// Copyright (c) 2021-2022 The keysight developers. All rights reserved.
// Project site: https://github.com/gotmc/keysight
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package esa has the ability to parse files from the Keysight/Agilent ESA
// spectrum analyzers.
package esa

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type FreqUnits string
type AmplitudeUnits string
type TimeUnits string

type Trace struct {
	Timestamp        time.Time
	OriginalFilename string
	Title            string
	Model            string
	SerialNum        string
	CenterFreq       float64
	CenterFreqUnits  FreqUnits
	Span             float64
	SpanUnits        FreqUnits
	RBW              float64
	RBWUnits         FreqUnits
	VBW              float64
	VBWUnits         FreqUnits
	RefLevel         float64
	RefLevelUnits    AmplitudeUnits
	SweepTime        float64
	SweepTimeUnits   TimeUnits
	NumPoints        int
	FrequenciesUnits string
	Trace1Units      string
	Trace2Units      string
	Trace3Units      string
	Frequencies      []float64
	Trace1           []float64
	Trace2           []float64
	Trace3           []float64
}

// ReadCSVFile reads the Keysight/Agilent ESA trace data saved in CSV format.
// It should be noted that the ESA CSV file does not meet the format described
// in RFC 4180.
func ReadCSVFile(filename string) (Trace, error) {
	trace := Trace{}
	file, err := os.Open(filename)
	if err != nil {
		return trace, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse first line, which should contain the timestamp and original
	// filename.
	scanner.Scan()
	line := scanner.Text()
	s := strings.Split(line, ",")
	if len(s) != 2 {
		return trace, fmt.Errorf("error in first line: %s", line)
	}
	trace.OriginalFilename = s[1]

	// Parse second line, which should contain the title.
	scanner.Scan()
	line = scanner.Text()
	s = strings.Split(line, ",")
	if len(s) != 2 {
		return trace, fmt.Errorf("error in Title line: %s", line)
	}
	trace.Title = s[1]

	// Parse third line, which should contain the model.
	scanner.Scan()
	line = scanner.Text()
	s = strings.Split(line, ",")
	if len(s) != 2 {
		return trace, fmt.Errorf("error in model line: %s", line)
	}
	trace.Model = s[1]

	// Parse fourth line, which should contain the serial number.
	scanner.Scan()
	line = scanner.Text()
	s = strings.Split(line, ",")
	if len(s) != 2 {
		return trace, fmt.Errorf("error in serial number line: %s", line)
	}
	trace.SerialNum = strings.TrimSuffix(s[1], "\x00")

	// Parse fifth line, which should contain the center frequency value and
	// units.
	scanner.Scan()
	line = scanner.Text()
	s = strings.Split(line, ",")
	if len(s) != 3 {
		return trace, fmt.Errorf("error in center frequency line: %s", line)
	}
	f, err := strconv.ParseFloat(s[1], 64)
	if err != nil {
		return trace, fmt.Errorf("error parsing center frequency: %s", err)
	}
	trace.CenterFreq = f
	//TODO(mdr): Parse center freq units.

	// Parse sixth line, which should contain the span value and units.
	scanner.Scan()
	line = scanner.Text()
	s = strings.Split(line, ",")
	if len(s) != 3 {
		return trace, fmt.Errorf("error in span line: %s", line)
	}
	f, err = strconv.ParseFloat(s[1], 64)
	if err != nil {
		return trace, fmt.Errorf("error parsing span: %s", err)
	}
	trace.Span = f

	// Parse seventh line, which should contain the resolution bandwidth (RBW)
	// value and units.
	scanner.Scan()
	line = scanner.Text()
	s = strings.Split(line, ",")
	if len(s) != 3 {
		return trace, fmt.Errorf("error in rbw line: %s", line)
	}
	f, err = strconv.ParseFloat(s[1], 64)
	if err != nil {
		return trace, fmt.Errorf("error parsing rbw: %s", err)
	}
	trace.RBW = f

	// Parse eighth line, which should contain the video bandwidth (vbw) value
	// and units.
	scanner.Scan()
	line = scanner.Text()
	s = strings.Split(line, ",")
	if len(s) != 3 {
		return trace, fmt.Errorf("error in vbw line: %s", line)
	}
	f, err = strconv.ParseFloat(s[1], 64)
	if err != nil {
		return trace, fmt.Errorf("error parsing vbw: %s", err)
	}
	trace.VBW = f

	// Parse ninth line, which should contain the reference level value and
	// units.
	scanner.Scan()
	line = scanner.Text()
	s = strings.Split(line, ",")
	if len(s) != 3 {
		return trace, fmt.Errorf("error in ref level line: %s", line)
	}
	f, err = strconv.ParseFloat(s[1], 64)
	if err != nil {
		return trace, fmt.Errorf("error parsing ref level: %s", err)
	}
	trace.RefLevel = f

	// Parse tenth line, which should contain the sweep time value and units.
	scanner.Scan()
	line = scanner.Text()
	s = strings.Split(line, ",")
	if len(s) != 3 {
		return trace, fmt.Errorf("error in sweep time line: %s", line)
	}
	f, err = strconv.ParseFloat(s[1], 64)
	if err != nil {
		return trace, fmt.Errorf("error parsing sweep time: %s", err)
	}
	trace.SweepTime = f

	// Parse eleventh line, which should contain the number of points.
	scanner.Scan()
	line = scanner.Text()
	s = strings.Split(line, ",")
	if len(s) != 2 {
		return trace, fmt.Errorf("error in num of points line: %s", line)
	}
	i, err := strconv.Atoi(s[1])
	if err != nil {
		return trace, fmt.Errorf("error parsing num of points: %s", err)
	}
	trace.NumPoints = i

	// Skip lines 12 and 13, which are blank.
	scanner.Scan()
	scanner.Scan()

	// Parse the remaining lines, which should now comply with RFC 4180 and be a
	// standard CSV file.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return trace, err
	}

	return trace, nil
}
