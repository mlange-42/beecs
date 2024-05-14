package model_test

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs/model"
	"github.com/mlange-42/beecs/params"
	"github.com/stretchr/testify/assert"
)

func TestSetParameter(t *testing.T) {
	p := params.Default()
	m := model.Default(&p, nil)

	err := model.SetParameter(&m.World, "params.Foragers.SquadronSize", 10)
	assert.Nil(t, err)
	assert.Equal(t, 10, ecs.GetResource[params.Foragers](&m.World).SquadronSize)

	err = model.SetParameter(&m.World, "params.Foragers.SquadronSize", "25")
	assert.Nil(t, err)
	assert.Equal(t, 25, ecs.GetResource[params.Foragers](&m.World).SquadronSize)

	err = model.SetParameter(&m.World, "params.Foragers.SquadronSize", 10.0)
	assert.NotNil(t, err)

	err = model.SetParameter(&m.World, "params.ForagerS.SquadronSize", 10)
	assert.NotNil(t, err)

	err = model.SetParameter(&m.World, "params.Foragers.SquadronSizE", 10)
	assert.NotNil(t, err)
}

func TestGetParameter(t *testing.T) {
	p := params.Default()
	m := model.Default(&p, nil)

	v, err := model.GetParameter(&m.World, "params.Foragers.SquadronSize")
	assert.Nil(t, err)
	assert.Equal(t, int64(100), v)

	_, err = model.GetParameter(&m.World, "params.ForagerS.SquadronSize")
	assert.NotNil(t, err)

	_, err = model.GetParameter(&m.World, "params.Foragers.SquadronSizE")
	assert.NotNil(t, err)
}
