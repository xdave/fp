package fp

import (
	"reflect"
	"sync"
)

// Performs operation on ALL items of collection (slice or map) concurrently,
// returning a new collection with the results (maintains order, too)
func Map(collection interface{}, operation interface{}) interface{} {
	fn := reflect.ValueOf(operation)
	cl := reflect.ValueOf(collection)
	typ := reflect.TypeOf(collection)
	switch cl.Kind() {
	case reflect.Slice:
		results := reflect.MakeSlice(typ, cl.Len(), cl.Cap())
		var wg sync.WaitGroup
		for i := range make([]struct{}, cl.Len()) {
			wg.Add(1)
			go func(i int, results reflect.Value) {
				item := cl.Index(i)
				index := reflect.ValueOf(i)
				returned := fn.Call([]reflect.Value{item, index})
				results.Index(i).Set(returned[0])
				wg.Done()
			}(i, results)
		}
		wg.Wait()
		return results.Interface()
	case reflect.Map:
		results := reflect.MakeMap(typ)
		keys := cl.MapKeys()
		var wg sync.WaitGroup
		for i := range make([]struct{}, len(keys)) {
			wg.Add(1)
			go func(i int, results reflect.Value) {
				index := keys[i]
				item := cl.MapIndex(index)
				returned := fn.Call([]reflect.Value{item, index})
				results.SetMapIndex(index, returned[0])
				wg.Done()
			}(i, results)
		}
		wg.Wait()
		return results.Interface()
	default:
		panic("Kind of collection is not slice or map!")
	}
}

// Performs an operation on ALL items of collection (slice or map)
// accumulating the results
// TODO: Concurrency
func Reduce(collection interface{}, operation interface{}) interface{} {
	fn := reflect.ValueOf(operation)
	cl := reflect.ValueOf(collection)
	var r reflect.Value
	switch cl.Kind() {
	case reflect.Slice:
		for i := range make([]struct{}, cl.Len()) {
			item := cl.Index(i)
			if i == 0 {
				r = item
			} else {
				returned := fn.Call([]reflect.Value{r, item})
				r = returned[0]
			}
		}
	case reflect.Map:
		keys := cl.MapKeys()
		for i := range make([]struct{}, len(keys)) {
			index := keys[i]
			item := cl.MapIndex(index)
			if i == 0 {
				r = item
			} else {
				returned := fn.Call([]reflect.Value{r, item})
				r = returned[0]
			}
		}
	default:
		panic("Kind of collection is not slice or map!")
	}
	return r.Interface()
}
