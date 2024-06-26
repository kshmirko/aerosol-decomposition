package plots

import (
	"log"

	"gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const CANVAS_X int = 10
const CANVAS_Y int = 10

func PlotY(y, yh utlis.Vector, xlab, ylab, title, fname string) error {
	p := plot.New()
	p.Title.Text = title
	p.Y.Scale = plot.LogScale{}
	p.Y.Tick.Marker = plot.LogTicks{Prec: 3}
	p.X.Label.Text = xlab
	p.Y.Label.Text = ylab

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(3)

	ptsy := make(plotter.XYs, 0, len(y))
	ptsyh := make(plotter.XYs, 0, len(yh))
	for i := range y {
		ptsyh = append(ptsyh, plotter.XY{
			X: float64(i),
			Y: yh[i],
		})
		ptsy = append(ptsy, plotter.XY{
			X: float64(i),
			Y: y[i],
		})
	}

	line, err := plotter.NewLine(ptsy)

	if err != nil {
		log.Fatal(err)
	}

	scatter, err := plotter.NewScatter(ptsyh)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(line, scatter)
	p.Legend.Add("measured", line)
	p.Legend.Add("reconstructed", scatter)

	err = p.Save(10*vg.Centimeter, 10*vg.Centimeter, fname)
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func Scatter(y, yh utlis.Vector, xlab, ylab, title, fname string) error {
	p := plot.New()
	p.Title.Text = title
	p.Y.Scale = plot.LogScale{}
	p.Y.Tick.Marker = plot.LogTicks{Prec: 3}
	p.X.Scale = plot.LogScale{}
	p.X.Tick.Marker = plot.LogTicks{Prec: 3}

	p.X.Label.Text = xlab
	p.Y.Label.Text = ylab

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(3)

	ptsy := make(plotter.XYs, 0, len(y))

	for i := range y {
		ptsy = append(ptsy, plotter.XY{
			X: y[i],
			Y: yh[i],
		})
	}

	scatter, err := plotter.NewScatter(ptsy)
	if err != nil {
		log.Fatal(err)
	}

	line, err := plotter.NewLine(plotter.XYs{
		plotter.XY{
			X: 0.0000001,
			Y: 0.0000001,
		},
		plotter.XY{
			X: 0.5,
			Y: 0.5,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	p.Add(scatter, line)

	err = p.Save(font.Length(CANVAS_X)*vg.Centimeter, font.Length(CANVAS_Y)*vg.Centimeter, fname)
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func PlotXY(x, y utlis.Vector, xlab, ylab, title, fname string) error {

	p := plot.New()
	p.Title.Text = title
	p.X.Scale = plot.LogScale{}
	p.X.Tick.Marker = plot.LogTicks{Prec: 3}
	p.X.Label.Text = xlab
	p.Y.Label.Text = ylab
	p.Add(plotter.NewGrid())
	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(3)

	ptsy := make(plotter.XYs, 0, len(y))
	for i := range y {
		ptsy = append(ptsy, plotter.XY{
			X: x[i],
			Y: y[i],
		})
	}

	line, err := plotter.NewLine(ptsy)

	if err != nil {
		log.Fatal(err)
	}

	scatter, err := plotter.NewScatter(ptsy)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(line, scatter)

	err = p.Save(10*vg.Centimeter, 10*vg.Centimeter, fname)
	if err != nil {
		log.Panic(err)
	}
	return nil
}
