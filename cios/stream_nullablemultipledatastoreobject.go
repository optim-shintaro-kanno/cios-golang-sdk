/*
 * Collection utility of NullableMultipleDataStoreObject Struct
 *
 * Generated by: Go Streamer
 * Author
 */

package cios

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type NullableMultipleDataStoreObjectStream []NullableMultipleDataStoreObject

func NullableMultipleDataStoreObjectStreamOf(arg ...NullableMultipleDataStoreObject) NullableMultipleDataStoreObjectStream {
	return arg
}
func NullableMultipleDataStoreObjectStreamFrom(arg []NullableMultipleDataStoreObject) NullableMultipleDataStoreObjectStream {
	return arg
}
func CreateNullableMultipleDataStoreObjectStream(arg ...NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	tmp := NullableMultipleDataStoreObjectStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleDataStoreObjectStream(arg []NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	tmp := NullableMultipleDataStoreObjectStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleDataStoreObjectStream) Add(arg NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleDataStoreObjectStream) AddAll(arg ...NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleDataStoreObjectStream) AddSafe(arg *NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Aggregate(fn func(NullableMultipleDataStoreObject, NullableMultipleDataStoreObject) NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	result := NullableMultipleDataStoreObjectStreamOf()
	self.ForEach(func(v NullableMultipleDataStoreObject, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleDataStoreObject{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleDataStoreObjectStream) AllMatch(fn func(NullableMultipleDataStoreObject, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleDataStoreObjectStream) AnyMatch(fn func(NullableMultipleDataStoreObject, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleDataStoreObjectStream) Clone() *NullableMultipleDataStoreObjectStream {
	temp := make([]NullableMultipleDataStoreObject, self.Len())
	copy(temp, *self)
	return (*NullableMultipleDataStoreObjectStream)(&temp)
}
func (self *NullableMultipleDataStoreObjectStream) Copy() *NullableMultipleDataStoreObjectStream {
	return self.Clone()
}
func (self *NullableMultipleDataStoreObjectStream) Concat(arg []NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleDataStoreObjectStream) Contains(arg NullableMultipleDataStoreObject) bool {
	return self.FindIndex(func(_arg NullableMultipleDataStoreObject, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleDataStoreObjectStream) Clean() *NullableMultipleDataStoreObjectStream {
	*self = NullableMultipleDataStoreObjectStreamOf()
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Delete(index int) *NullableMultipleDataStoreObjectStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleDataStoreObjectStream) DeleteRange(startIndex, endIndex int) *NullableMultipleDataStoreObjectStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Distinct() *NullableMultipleDataStoreObjectStream {
	caches := map[string]bool{}
	result := NullableMultipleDataStoreObjectStreamOf()
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if !f {
				result = append(result, v)
			}
		} else if caches[key] = true; !f {
			result = append(result, v)
		}

	}
	*self = result
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Each(fn func(NullableMultipleDataStoreObject)) *NullableMultipleDataStoreObjectStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleDataStoreObjectStream) EachRight(fn func(NullableMultipleDataStoreObject)) *NullableMultipleDataStoreObjectStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Equals(arr []NullableMultipleDataStoreObject) bool {
	if (*self == nil) != (arr == nil) || len(*self) != len(arr) {
		return false
	}
	for i := range *self {
		if !reflect.DeepEqual((*self)[i], arr[i]) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleDataStoreObjectStream) Filter(fn func(NullableMultipleDataStoreObject, int) bool) *NullableMultipleDataStoreObjectStream {
	result := NullableMultipleDataStoreObjectStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDataStoreObjectStream) FilterSlim(fn func(NullableMultipleDataStoreObject, int) bool) *NullableMultipleDataStoreObjectStream {
	result := NullableMultipleDataStoreObjectStreamOf()
	caches := map[string]bool{}
	for i, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v, i); caches[key] {
			result.Add(v)

		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Find(fn func(NullableMultipleDataStoreObject, int) bool) *NullableMultipleDataStoreObject {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleDataStoreObjectStream) FindOr(fn func(NullableMultipleDataStoreObject, int) bool, or NullableMultipleDataStoreObject) NullableMultipleDataStoreObject {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleDataStoreObjectStream) FindIndex(fn func(NullableMultipleDataStoreObject, int) bool) int {
	if self == nil {
		return -1
	}
	for i, v := range *self {
		if fn(v, i) {
			return i
		}
	}
	return -1
}
func (self *NullableMultipleDataStoreObjectStream) First() *NullableMultipleDataStoreObject {
	return self.Get(0)
}
func (self *NullableMultipleDataStoreObjectStream) FirstOr(arg NullableMultipleDataStoreObject) NullableMultipleDataStoreObject {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDataStoreObjectStream) ForEach(fn func(NullableMultipleDataStoreObject, int)) *NullableMultipleDataStoreObjectStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleDataStoreObjectStream) ForEachRight(fn func(NullableMultipleDataStoreObject, int)) *NullableMultipleDataStoreObjectStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleDataStoreObjectStream) GroupBy(fn func(NullableMultipleDataStoreObject, int) string) map[string][]NullableMultipleDataStoreObject {
	m := map[string][]NullableMultipleDataStoreObject{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleDataStoreObjectStream) GroupByValues(fn func(NullableMultipleDataStoreObject, int) string) [][]NullableMultipleDataStoreObject {
	var tmp [][]NullableMultipleDataStoreObject
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleDataStoreObjectStream) IndexOf(arg NullableMultipleDataStoreObject) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleDataStoreObjectStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleDataStoreObjectStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleDataStoreObjectStream) Last() *NullableMultipleDataStoreObject {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleDataStoreObjectStream) LastOr(arg NullableMultipleDataStoreObject) NullableMultipleDataStoreObject {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDataStoreObjectStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleDataStoreObjectStream) Limit(limit int) *NullableMultipleDataStoreObjectStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleDataStoreObjectStream) Map(fn func(NullableMultipleDataStoreObject, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreObjectStream) Map2Int(fn func(NullableMultipleDataStoreObject, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreObjectStream) Map2Int32(fn func(NullableMultipleDataStoreObject, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreObjectStream) Map2Int64(fn func(NullableMultipleDataStoreObject, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreObjectStream) Map2Float32(fn func(NullableMultipleDataStoreObject, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreObjectStream) Map2Float64(fn func(NullableMultipleDataStoreObject, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreObjectStream) Map2Bool(fn func(NullableMultipleDataStoreObject, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreObjectStream) Map2Bytes(fn func(NullableMultipleDataStoreObject, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreObjectStream) Map2String(fn func(NullableMultipleDataStoreObject, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDataStoreObjectStream) Max(fn func(NullableMultipleDataStoreObject, int) float64) *NullableMultipleDataStoreObject {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Max(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *NullableMultipleDataStoreObjectStream) Min(fn func(NullableMultipleDataStoreObject, int) float64) *NullableMultipleDataStoreObject {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Min(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *NullableMultipleDataStoreObjectStream) NoneMatch(fn func(NullableMultipleDataStoreObject, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleDataStoreObjectStream) Get(index int) *NullableMultipleDataStoreObject {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleDataStoreObjectStream) GetOr(index int, arg NullableMultipleDataStoreObject) NullableMultipleDataStoreObject {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDataStoreObjectStream) Peek(fn func(*NullableMultipleDataStoreObject, int)) *NullableMultipleDataStoreObjectStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableMultipleDataStoreObjectStream) Reduce(fn func(NullableMultipleDataStoreObject, NullableMultipleDataStoreObject, int) NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	return self.ReduceInit(fn, NullableMultipleDataStoreObject{})
}
func (self *NullableMultipleDataStoreObjectStream) ReduceInit(fn func(NullableMultipleDataStoreObject, NullableMultipleDataStoreObject, int) NullableMultipleDataStoreObject, initialValue NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	result := NullableMultipleDataStoreObjectStreamOf()
	self.ForEach(func(v NullableMultipleDataStoreObject, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleDataStoreObjectStream) ReduceInterface(fn func(interface{}, NullableMultipleDataStoreObject, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleDataStoreObject{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDataStoreObjectStream) ReduceString(fn func(string, NullableMultipleDataStoreObject, int) string) []string {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDataStoreObjectStream) ReduceInt(fn func(int, NullableMultipleDataStoreObject, int) int) []int {
	result := []int{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDataStoreObjectStream) ReduceInt32(fn func(int32, NullableMultipleDataStoreObject, int) int32) []int32 {
	result := []int32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDataStoreObjectStream) ReduceInt64(fn func(int64, NullableMultipleDataStoreObject, int) int64) []int64 {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDataStoreObjectStream) ReduceFloat32(fn func(float32, NullableMultipleDataStoreObject, int) float32) []float32 {
	result := []float32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDataStoreObjectStream) ReduceFloat64(fn func(float64, NullableMultipleDataStoreObject, int) float64) []float64 {
	result := []float64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDataStoreObjectStream) ReduceBool(fn func(bool, NullableMultipleDataStoreObject, int) bool) []bool {
	result := []bool{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDataStoreObjectStream) Reverse() *NullableMultipleDataStoreObjectStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Replace(fn func(NullableMultipleDataStoreObject, int) NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	return self.ForEach(func(v NullableMultipleDataStoreObject, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleDataStoreObjectStream) Select(fn func(NullableMultipleDataStoreObject) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleDataStoreObjectStream) Set(index int, val NullableMultipleDataStoreObject) *NullableMultipleDataStoreObjectStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Skip(skip int) *NullableMultipleDataStoreObjectStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleDataStoreObjectStream) SkippingEach(fn func(NullableMultipleDataStoreObject, int) int) *NullableMultipleDataStoreObjectStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Slice(startIndex, n int) *NullableMultipleDataStoreObjectStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleDataStoreObject{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Sort(fn func(i, j int) bool) *NullableMultipleDataStoreObjectStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleDataStoreObjectStream) Tail() *NullableMultipleDataStoreObject {
	return self.Last()
}
func (self *NullableMultipleDataStoreObjectStream) TailOr(arg NullableMultipleDataStoreObject) NullableMultipleDataStoreObject {
	return self.LastOr(arg)
}
func (self *NullableMultipleDataStoreObjectStream) ToList() []NullableMultipleDataStoreObject {
	return self.Val()
}
func (self *NullableMultipleDataStoreObjectStream) Unique() *NullableMultipleDataStoreObjectStream {
	return self.Distinct()
}
func (self *NullableMultipleDataStoreObjectStream) Val() []NullableMultipleDataStoreObject {
	if self == nil {
		return []NullableMultipleDataStoreObject{}
	}
	return *self.Copy()
}
func (self *NullableMultipleDataStoreObjectStream) While(fn func(NullableMultipleDataStoreObject, int) bool) *NullableMultipleDataStoreObjectStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleDataStoreObjectStream) Where(fn func(NullableMultipleDataStoreObject) bool) *NullableMultipleDataStoreObjectStream {
	result := NullableMultipleDataStoreObjectStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDataStoreObjectStream) WhereSlim(fn func(NullableMultipleDataStoreObject) bool) *NullableMultipleDataStoreObjectStream {
	result := NullableMultipleDataStoreObjectStreamOf()
	caches := map[string]bool{}
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v); caches[key] {
			result.Add(v)
		}
	}
	*self = result
	return self
}
