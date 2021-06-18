package figlet4go

import (
	"flag"
	"fmt"
	"testing"

	"github.com/fatih/color"
)

var flag_str = flag.String("str", "golang", "input string")

func TestFonts(t *testing.T) {
	flag.Parse()
	str := *flag_str
	ascii := NewAsciiRender()

	// most simple Usage
	renderStr, _ := ascii.Render(str)
	fmt.Println(renderStr)

	// change the font color
	colors := [...]color.Attribute{
		color.FgMagenta,
		color.FgYellow,
		color.FgBlue,
		color.FgCyan,
		color.FgRed,
		color.FgWhite,
	}
	options := NewRenderOptions()
	options.FontColor = make([]color.Attribute, len(str))
	for i := range options.FontColor {
		options.FontColor[i] = colors[i%len(colors)]
	}
	renderStr, _ = ascii.RenderOpts(str, options)
	fmt.Println(renderStr)

	// change the font
	// except the default font,others need to be load from disk
	// here is the font :
	// ftp://ftp.figlet.org/pub/figlet/fonts/contributed.tar.gz
	// ftp://ftp.figlet.org/pub/figlet/fonts/international.tar.gz
	// download and extract to the disk,then specify the file path to load
	ascii.LoadFont("./fonts/ANSI Shadow.flf")
	options.FontName = "ANSI Shadow.flf"
	renderStr, _ = ascii.RenderOpts(str, options)
	fmt.Println(renderStr)
}
