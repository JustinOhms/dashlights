package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type dashlight struct {
	Name        string
	Glyph       string
	Diagnostic  string
	Color       *color.Color
	UnsetString string
}

var colorMap = map[string]color.Attribute{
	"FGBLACK":      color.FgBlack,
	"FGRED":        color.FgRed,
	"FGGREEN":      color.FgGreen,
	"FGYELLOW":     color.FgYellow,
	"FGBLUE":       color.FgBlue,
	"FGMAGENTA":    color.FgMagenta,
	"FGCYAN":       color.FgCyan,
	"FGWHITE":      color.FgWhite,
	"FGHIBLACK":    color.FgHiBlack,
	"FGHIRED":      color.FgHiRed,
	"FGHIGREEN":    color.FgHiGreen,
	"FGHIYELLOW":   color.FgHiYellow,
	"FGHIBLUE":     color.FgHiBlue,
	"FGHIMAGENTA":  color.FgHiMagenta,
	"FGHICYAN":     color.FgHiCyan,
	"FGHIWHITE":    color.FgHiWhite,
	"BGBLACK":      color.BgBlack,
	"BGRED":        color.BgRed,
	"BGGREEN":      color.BgGreen,
	"BGYELLOW":     color.BgYellow,
	"BGBLUE":       color.BgBlue,
	"BGMAGENTA":    color.BgMagenta,
	"BGCYAN":       color.BgCyan,
	"BGWHITE":      color.BgWhite,
	"BGHIBLACK":    color.BgHiBlack,
	"BGHIRED":      color.BgHiRed,
	"BGHIGREEN":    color.BgHiGreen,
	"BGHIYELLOW":   color.BgHiYellow,
	"BGHIBLUE":     color.BgHiBlue,
	"BGHIMAGENTA":  color.BgHiMagenta,
	"BGHICYAN":     color.BgHiCyan,
	"BGHIWHITE":    color.BgHiWhite,
	"REVERSEVIDEO": color.ReverseVideo,
}

var diagMode *bool
var listColorMode *bool
var clearMode *bool

func init() {
	diagMode = flag.Bool("diag", false, "display diagnostic information, if provided.")
	listColorMode = flag.Bool("listcolors", false, "show supported color attributes.")
	clearMode = flag.Bool("clear", false, "eval code to clear set dashlights.")
}

func flexPrintf(w io.Writer, format string, args ...interface{}) {
	fmt.Fprintf(w, format, args...)
}

func flexPrintln(w io.Writer, line string) {
	fmt.Fprintln(w, line)
}

func displayClearCodes(w io.Writer, lights *[]dashlight) {
	for _, light := range *lights {
		flexPrintln(w, light.UnsetString)
	}
}

func displayColorList(w io.Writer) {
	keys := make([]string, 0)
	for k := range colorMap {
		keys = append(keys, k)
	}
	sizeKeys := len(keys)
	sort.Strings(keys)
	flexPrintln(w, "Supported color attributes:")
	for i, attrib := range keys {
		flexPrintf(w, "%s", attrib)
		if i < sizeKeys-1 {
			flexPrintf(w, "%s", ", ")
		}
	}
	flexPrintln(w, "")
}

var lights []dashlight

func main() {
	flag.Parse()
	parseEnviron(os.Environ(), &lights)
	display(os.Stdout, &lights)
}

func parseEnviron(environ []string, lights *[]dashlight) {
	for _, env := range environ {
		parseDashlightFromEnv(lights, env)
	}
}

func display(w io.Writer, lights *[]dashlight) {
	if *listColorMode {
		displayColorList(w)
		return
	}
	if *clearMode {
		displayClearCodes(w, lights)
		return
	}
	displayDashlights(w, lights)
	if *diagMode {
		displayDiagnostics(w, lights)
	}
}

func displayDashlights(w io.Writer, lights *[]dashlight) {
	for _, light := range *lights {
		lamp := light.Color.SprintfFunc()("%s ", light.Glyph)
		flexPrintf(w, "%s ", lamp)
	}
	if len(*lights) > 0 {
		flexPrintln(w, "")
	}
}

func displayDiagnostics(w io.Writer, lights *[]dashlight) {
	flexPrintf(w, "\n-------- Diagnostics --------\n")
	for _, light := range *lights {
		lamp := light.Color.SprintfFunc()("%s ", light.Glyph)
		flexPrintf(w, "%s: %s - %s\n", lamp, light.Name, light.Diagnostic)
	}
}

func parseDashlightFromEnv(lights *[]dashlight, env string) {
	kv := strings.Split(env, "=")
	dashvar := kv[0]
	diagnostic := kv[1]
	if strings.Contains(dashvar, "DASHLIGHT_") {
		if diagnostic == "" {
			diagnostic = "No diagnostic info provided."
		}
		elements := strings.Split(dashvar, "_")
		if len(elements) < 3 {
			// dashvars must minimally be of form: DASHLIGHT_{name}_{utf8hex}
			return
		}
		// begin shifting elements off elements slice, ignore leading DASHLIGHT_ prefix
		name, elements := elements[1], elements[2:]
		hexstr, elements := elements[0], elements[1:]
		glyph, err := utf8HexToString(string(hexstr))
		if err != nil {
			return
		}
		dashColor := color.New()
		// process any remaining elements as color additions
		for _, colorstr := range elements {
			dashColor.Add(colorMap[colorstr])
		}
		*lights = append(*lights, dashlight{
			Name:        name,
			Glyph:       glyph,
			Diagnostic:  diagnostic,
			Color:       dashColor,
			UnsetString: "unset " + dashvar,
		})
	}
}

func utf8HexToString(hex string) (string, error) {
	i, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return "", err
	}
	return string(i), nil
}
