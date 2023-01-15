package services

import (
	"flag"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type GenImageService struct{}

func NewGenImageService() *GenImageService {
	return &GenImageService{}
}

func (s *GenImageService) GenImage(text string) image.Image {
	flag.Parse()

	dc := gg.NewContext(1200, 630)

	dc.SetColor(color.RGBA{255, 255, 255, 255})
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.Fill()

	frame, err := gg.LoadImage("services/src/frame.png")
	if err != nil {
		panic(errors.Wrap(err, "load background image"))
	}
	frame = imaging.Fill(frame, dc.Width(), dc.Height(), imaging.Center, imaging.Lanczos)
	dc.DrawImage(frame, 0, 0)

	face, err := fetchPostScriptFontFace("services/font/NotoSansJP-Medium.otf")
	if err != nil {
		log.Fatalln(err)
	}
	dc.SetFontFace(face)
	dc.SetColor(color.Black)

	drawStringMultiLine(dc, []rune(text), 100.0)

	return dc.Image()
}

// PostScript アウトラインのフォント読み込み
func fetchPostScriptFontFace(fontfile string) (font.Face, error) {
	ftBinary, err := os.ReadFile(fontfile)
	if err != nil {
		return nil, err
	}

	ft, err := opentype.Parse(ftBinary)
	if err != nil {
		return nil, err
	}

	opt := opentype.FaceOptions{
		Size:    80.0,
		DPI:     72.0,
		Hinting: font.HintingNone,
	}

	face, _ := opentype.NewFace(ft, &opt)

	return face, nil
}

func drawStringMultiLine(dc *gg.Context, text []rune, margin float64) {
	width := float64(dc.Width())

	drewStrCnt := 0
	drewHeight := 40.0
	log.Println(string(text))
	for i := 0; i < len(text); i++ {
		candidateStr := text[drewStrCnt : i+1]
		candidateWidth, h := dc.MeasureString(string(candidateStr))
		if candidateWidth <= width-2*margin {
			// まだ右に余白がある場合
			continue
		}

		dc.DrawStringAnchored(string(text[drewStrCnt:i]), margin, drewHeight, 0, 1)
		drewStrCnt = i
		drewHeight += h
	}
	dc.DrawStringAnchored(string(text[drewStrCnt:]), margin, drewHeight, 0, 1)
}
