package inst

import (
	//"github.com/PaulioRandall/scarlet-go/token2/inst"
	"github.com/PaulioRandall/scarlet-go/token2/value"
	//"github.com/PaulioRandall/scarlet-go/token2/tree"
)

// InstData represents all data required by instructions.
type InstData map[DataRef]value.Value

// DataSet is a builder for creating InstData structs.
type DataSet struct {
	data InstData
	ref  DataRef
}

// NewDataSet returns a new initialised DataSet.
func NewDataSet() *DataSet {
	return &DataSet{
		data: InstData{},
	}
}

// Insert adds a new data value to the DataSet with a unique data reference.
func (ds *DataSet) Insert(v value.Value) DataRef {
	r := ds.newRef()
	ds.data[r] = v
	return r
}

// Compile returns the constructed InstData.
func (ds *DataSet) Compile() InstData {
	return ds.data
}

func (ds *DataSet) newRef() DataRef {
	ds.ref++
	return ds.ref
}
