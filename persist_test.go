package persist_test

import (
	"os"
	"testing"
	"time"

	"github.com/cheekybits/is"
	"github.com/keks/persist"
)

type obj struct {
	Name   string
	Number int
	When   time.Time
}

func TestPersist(t *testing.T) {
	is := is.New(t)

	f, err := os.Create("./file.tmp")
	is.NoErr(err)
	defer os.Remove("./file.tmp")

	o := &obj{
		Name:   "Mat",
		Number: 47,
		When:   time.Now(),
	}

	// save it
	err = persist.Save(f, o)
	is.NoErr(err)

	// load it
	var o2 obj
	err = persist.Load(f, &o2)
	is.NoErr(err)

	is.Equal(o.Name, o2.Name)
	is.Equal(o.Number, o2.Number)
	is.True(o.When.Equal(o2.When))

	// load it, twice
	o2.Name = ""
	o2.Number = 0
	err = persist.Load(f, &o2)
	is.NoErr(err)

	is.Equal(o.Name, o2.Name)
	is.Equal(o.Number, o2.Number)
	is.True(o.When.Equal(o2.When))
}
