package model

import (
	"github.com/mlange-42/ark-tools/app"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs/globals"
	"github.com/mlange-42/beecs/params"
	"github.com/mlange-42/beecs/sys"
)

// Default sets up the default beecs model with the standard sub-models.
//
// If the argument m is nil, a new model instance is created.
// If it is non-nil, the model is reset and re-used, saving some time for initialization and memory allocation.
func Default(p params.Params, app *app.App) *app.App {

	// Add parameters and other resources

	app = initializeModel(p, app)

	// Initialization

	app.AddSystem(&sys.InitStore{})          // initializes DecentHoney and IdealPollen now to mimic BEEHAVE more closely.
	app.AddSystem(&sys.InitCohorts{})        // initializes NewCohort-global now. Needed to introduce this somewhere, seemed to make the most sense here.
	app.AddSystem(&sys.InitPopulation{})     // unchanged
	app.AddSystem(&sys.InitPatchesList{})    // unchanged
	app.AddSystem(&sys.InitForagingPeriod{}) // unchanged

	// Sub-models

	app.AddSystem(&sys.CalcAff{})            // unchanged
	app.AddSystem(&sys.CalcForagingPeriod{}) // unchanged
	app.AddSystem(&sys.ReplenishPatches{})   // unchanged

	app.AddSystem(&sys.MortalityCohorts{})   // moved to the beginning, because this happens while ageing in Netlogo. Does not "take effect" in the model until next CountPopulation
	app.AddSystem(&sys.AgeCohorts{})         // happens with mortality in Netlogo and has to happen before new cohorts (including eggs) get initiated
	app.AddSystem(&sys.EggLaying{})          // unchanged, just relocated here
	app.AddSystem(&sys.TransitionForagers{}) // now does not actually initiate and fully transition foragers anymore. Only calculates how many foragers are to be initiated and decreases IHbee count

	app.AddSystem(&sys.CountPopulation{}) // included additional countingproc here to capture mortality effects and decrease of IHbees through TransitionForagers on BroodCare. Original BEEHAVE works this exact way. Importance: NECESSARY FOR ACCURATE FUNCTION
	app.AddSystem(&sys.BroodCare{})       // Moved in between the creation of new cohorts and the population decrease from the transition of cohorts to the next lifestage/job. BEEHAVE works this way.

	app.AddSystem(&sys.NewCohorts{})      // included this new subsystem to initialize the new cohorts, that were calculated in AgeCohorts. IHbees and Drones now get initiated here (mimics Netlogo exactly)
	app.AddSystem(&sys.CountPopulation{}) // included yet an additional countingproc here to capture NewCohorts and brood changes. Does not have a large effect; only necessary here because Foraging now recalculates decent honey for foraging decisionmaking. Could potentially be omitted, but BEEHAVE has it.

	app.AddSystem(&sys.Foraging{})          // remains largely unchanged, but initializes the new foragers from TransitionForagers now. Also introduced the calculation of decent honey to mimic the original model even closer
	app.AddSystem(&sys.MortalityForagers{}) // put this behind Foraging subsystem, because Netlogo also runs this at the end of the foraging IBM subsystem

	app.AddSystem(&sys.CountPopulation{})   // last counting proc that counts "final" amounts for feeding and next timestep. Moved before ConsumptionProcs because forager/cohort amounts affect consumption and MortalityForager has changed foragercount. Importance: NECESSARY FOR ACCURATE FUNCTION
	app.AddSystem(&sys.HoneyConsumption{})  // deactivated DecentHoney calculation here, as this gets updated in Foraging now (as is the case in BEEHAVE).
	app.AddSystem(&sys.PollenConsumption{}) // unchanged

	app.AddSystem(&sys.FixedTermination{}) // unchanged

	return app
}

// WithSystems sets up a beecs model with the given systems instead of the default ones.
//
// If the argument m is nil, a new model instance is created.
// If it is non-nil, the model is reset and re-used, saving some time for initialization and memory allocation.
func WithSystems(p params.Params, sys []app.System, app *app.App) *app.App {

	app = initializeModel(p, app)

	for _, s := range sys {
		app.AddSystem(s)
	}

	return app
}

func initializeModel(p params.Params, a *app.App) *app.App {
	if a == nil {
		a = app.New()
	} else {
		a.Reset()
	}

	p.Apply(&a.World)

	factory := globals.NewForagerFactory(&a.World)
	ecs.AddResource(&a.World, &factory)

	stats := globals.PopulationStats{}
	ecs.AddResource(&a.World, &stats)

	consumptionStats := globals.ConsumptionStats{}
	ecs.AddResource(&a.World, &consumptionStats)

	foragingStats := globals.ForagingStats{}
	ecs.AddResource(&a.World, &foragingStats)

	return a
}
