package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/fatih/color"
)

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func TestDisplayCodes(t *testing.T) {
	lights := make([]Dashlight, 0)
	lights = append(lights, Dashlight{
		Name:        "foo",
		Glyph:       "X",
		Diagnostic:  "",
		Color:       color.New(),
		UnsetString: "unset foo",
	})
	var b bytes.Buffer
	displayClearCodes(&b, &lights)
	if b.String() != "unset foo\n" {
		t.Error("Expected 'unset foo\n', got ", b.String())
	}
	lights = append(lights, Dashlight{
		Name:        "bar",
		Glyph:       "Y",
		Diagnostic:  "",
		Color:       color.New(),
		UnsetString: "unset bar",
	})
	b.Reset()
	displayClearCodes(&b, &lights)
	//	fmt.Printf("output: %s\n", "X"+b.String()+"X")
	if b.String() != "unset foo\nunset bar\n" {
		t.Error("Expected 'unset foo\nunset bar\n', got ", b.String())
	}
}

func TestParseDashlightFromEnv(t *testing.T) {
	lights := make([]Dashlight, 0)
	// missing namespace prefix...
	parseDashlightFromEnv(&lights, "FOO_2112_BGWHITE=foo")
	if 0 != len(lights) {
		t.Error("Expected length of 0, got ", len(lights))
	}
	// missing utf8 hex string...
	parseDashlightFromEnv(&lights, "DASHLIGHT_FOO=foo")
	if 0 != len(lights) {
		t.Error("Expected length of 0, got ", len(lights))
	}
	// invalid utf8 hex strings...
	parseDashlightFromEnv(&lights, "DASHLIGHT_FOO_ZZDA9=")
	if 0 != len(lights) {
		t.Error("Expected length of 0, got ", len(lights))
	}
	parseDashlightFromEnv(&lights, "DASHLIGHT_FOO_X=")
	if 0 != len(lights) {
		t.Error("Expected length of 0, got ", len(lights))
	}
	// invalid colormap codes are ignored...
	parseDashlightFromEnv(&lights, "DASHLIGHT_NOCODETEST_0021_NOTACODE=")
	if 1 != len(lights) {
		t.Error("Expected length of 1, got ", len(lights))
	}
	parseDashlightFromEnv(&lights, "DASHLIGHT_VALIDCODETEST_0021_BGWHITE=")
	if 2 != len(lights) {
		t.Error("Expected length of 2, got ", len(lights))
	}
	light := lights[1]
	if light.Name != "VALIDCODETEST" {
		t.Error("Expected Name of 'VALIDCODETEST', got ", light.Name)
	}
	if light.Glyph != "!" {
		t.Error("Expected Glyph of '!', got ", light.Glyph)
	}
	if light.Diagnostic != "No diagnostic info provided." {
		t.Error("Expected default diagnostic string, got ", light.Diagnostic)
	}
	if "*color.Color" != typeof(light.Color) {
		t.Error("Expected color to be type *color.Color, got ", typeof(light.Color))
	}
	if light.UnsetString != "unset DASHLIGHT_VALIDCODETEST_0021_BGWHITE" {
		t.Error("Expected valid unset string, got ", light.UnsetString)
	}
}

func TestDisplayColorList(t *testing.T) {
	var b bytes.Buffer
	listLen := len(colorMap)
	displayColorList(&b)
	commaCount := strings.Count(b.String(), ",")
	if commaCount != listLen-1 {
		t.Errorf("Expected %d commas in colorlist, got %d", listLen-1, commaCount)
	}
	// color attributes are listed in UPPER CASE...
	if !strings.Contains(b.String(), "BGWHITE") {
		t.Error("Expected to see string 'BGWHITE' in: ", b.String())
	}
}
