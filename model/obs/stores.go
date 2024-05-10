package obs

import (
	"fmt"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type Stores struct {
	stores      *res.Stores
	data        []float64
	energyHoney float64

	PollenFactor int
}

func (o *Stores) Initialize(w *ecs.World) {
	o.stores = ecs.GetResource[res.Stores](w)
	o.energyHoney = ecs.GetResource[res.EnergyParams](w).EnergyHoney
	o.data = make([]float64, len(o.Header()))

	if o.PollenFactor <= 0 {
		o.PollenFactor = 1
	}
}
func (o *Stores) Update(w *ecs.World) {}
func (o *Stores) Header() []string {
	if o.PollenFactor == 1 {
		return []string{"Honey", "Pollen", "DecentHoney", "IdealPollen"}
	}
	return []string{
		"Honey",
		fmt.Sprintf("Pollen x%d", o.PollenFactor),
		"DecentHoney",
		fmt.Sprintf("IdealPollen x%d", o.PollenFactor),
	}
}
func (o *Stores) Values(w *ecs.World) []float64 {
	o.data[0] = 0.001 * o.stores.Honey / o.energyHoney
	o.data[1] = 0.001 * o.stores.Pollen * float64(o.PollenFactor)
	o.data[2] = 0.001 * o.stores.DecentHoney / o.energyHoney
	o.data[3] = 0.001 * o.stores.IdealPollen * float64(o.PollenFactor)

	return o.data
}
