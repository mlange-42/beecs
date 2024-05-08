package obs

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model/res"
)

type Stores struct {
	stores      *res.Stores
	data        []float64
	energyHoney float64
}

func (o *Stores) Initialize(w *ecs.World) {
	o.stores = ecs.GetResource[res.Stores](w)
	o.energyHoney = ecs.GetResource[res.EnergyParams](w).EnergyHoney
	o.data = make([]float64, len(o.Header()))
}
func (o *Stores) Update(w *ecs.World) {}
func (o *Stores) Header() []string {
	return []string{"Honey [kg]", "Pollen [g]", "DecentHoney [kg]", "IdealPollen [g]"}
}
func (o *Stores) Values(w *ecs.World) []float64 {
	o.data[0] = 0.001 * o.stores.Honey / o.energyHoney
	o.data[1] = o.stores.Pollen
	o.data[2] = 0.001 * o.stores.DecentHoney / o.energyHoney
	o.data[3] = o.stores.IdealPollen

	return o.data
}
