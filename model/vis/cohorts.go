package vis

import (
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/mlange-42/arche/ecs"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
)

type Cohorts struct {
	drawer imdraw.IMDraw
	scale  float64
}

// Initialize the drawer.
func (c *Cohorts) Initialize(w *ecs.World, win *opengl.Window) {
	c.drawer = *imdraw.New(nil)
	c.scale = calcScaleCorrection()
}

// Update the drawer.
func (c *Cohorts) Update(w *ecs.World) {
}

// UpdateInputs handles input events of the previous frame update.
func (c *Cohorts) UpdateInputs(w *ecs.World, win *opengl.Window) {}

// Draw the drawer.
func (c *Cohorts) Draw(w *ecs.World, win *opengl.Window) {
	/*width := win.Canvas().Bounds().W()
	height := win.Canvas().Bounds().H()

	bins := 100*/
}

func calcScaleCorrection() float64 {
	width := 100.0
	c := vgimg.New(vg.Points(width), vg.Points(width))
	img := c.Image()
	return width / float64(img.Bounds().Dx())
}
