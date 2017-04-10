package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Dashlight struct {
	Name       string
	Glyph      string
	Diagnostic string
	Color      *color.Color
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

func init() {
	diagMode = flag.Bool("diag", false, "display diagnostic information, if provided.")
	listColorMode = flag.Bool("listcolors", false, "show supported color attributes.")
}

func displayColorList() {
	keys := make([]string, 0)
	for k := range colorMap {
		keys = append(keys, k)
	}
	sizeKeys := len(keys)
	sort.Strings(keys)
	fmt.Println("Supported color attributes:")
	for i, attrib := range keys {
		fmt.Printf("%s", attrib)
		if i < sizeKeys-1 {
			fmt.Printf("%s", ", ")
		}
	}
	fmt.Println()
}

func main() {
	flag.Parse()
	if *listColorMode {
		displayColorList()
		os.Exit(0)
	}

	var lights []Dashlight

	for _, env := range os.Environ() {
		parseDashlightFromEnv(&lights, env)
	}

	for _, light := range lights {
		light.Color.Printf("%s", light.Glyph)
		fmt.Printf("%s", "  ")
	}
	if len(lights) > 0 {
		fmt.Println()
	}
	if *diagMode {
		fmt.Printf("\n-------- Diagnostics --------\n")
		for _, light := range lights {
			light.Color.Printf("%s", light.Glyph)
			fmt.Printf("  : %s - %s\n", light.Name, light.Diagnostic)
		}
	}
}

func parseDashlightFromEnv(lights *[]Dashlight, env string) {
	kv := strings.Split(env, "=")
	dashvar := kv[0]
	diagnostic := kv[1]
	if strings.Contains(dashvar, "DASHLIGHT_") {
		if diagnostic == "" {
			diagnostic = "No diagnostic info provided."
		}
		elements := strings.Split(dashvar, "_")
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
		*lights = append(*lights, Dashlight{
			Name:       name,
			Glyph:      glyph,
			Diagnostic: diagnostic,
			Color:      dashColor,
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
