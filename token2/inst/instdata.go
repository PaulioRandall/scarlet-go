package inst

import (
	//"github.com/PaulioRandall/scarlet-go/token2/inst"
	"github.com/PaulioRandall/scarlet-go/token2/value"
	//"github.com/PaulioRandall/scarlet-go/token2/tree"
)

type InstData map[DataRef]value.Value

type DataSet struct {
	data InstData
	ref  DataRef
}

func (ds *DataSet) Insert(v value.Value) DataRef {
	r := ds.newRef()
	ds.data[r] = v
	return r
}

func (ds *DataSet) Compile() InstData {
	return ds.data
}

func (ds *DataSet) newRef() DataRef {
	r := ds.ref
	ds.ref++
	return r
}
