package settings

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	view := New()
	err := view.Load(filepath.Join("testdata", "starwars.json.conf"), FormatJSON)
	assert.NoError(t, err)
	assert.EqualValues(t, view.Get("jedi.yoda.occupation", nil), "master")
	assert.EqualValues(t, view.Get("ho.ho.ho", "ho"), "ho")
}

func TestCreate(t *testing.T) {
	view := New()
	view.Set("name", "yoda")
	view.Set("planets", []string{"naboo", "corsuant", "dagobah", "tatooin"})
	view.Set("jedi.yoda.occupation", "master")
	view.Set("jedi.yoda.hair", "grey")
	view.Set("jedi.yoda.hair.count", 5)
	view.Set("jedi.yoda.allmovies", false)
	err := view.Save(filepath.Join("testdata", "starwars.json.conf"), FormatJSON)
	assert.NoError(t, err)
}
