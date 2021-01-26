/*
 * Collection utility of NullableCollectionLocation Struct
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

type NullableCollectionLocationStream []NullableCollectionLocation

func NullableCollectionLocationStreamOf(arg ...NullableCollectionLocation) NullableCollectionLocationStream {
	return arg
}
func NullableCollectionLocationStreamFrom(arg []NullableCollectionLocation) NullableCollectionLocationStream {
	return arg
}
func CreateNullableCollectionLocationStream(arg ...NullableCollectionLocation) *NullableCollectionLocationStream {
	tmp := NullableCollectionLocationStreamOf(arg...)
	return &tmp
}
func GenerateNullableCollectionLocationStream(arg []NullableCollectionLocation) *NullableCollectionLocationStream {
	tmp := NullableCollectionLocationStreamFrom(arg)
	return &tmp
}

func (self *NullableCollectionLocationStream) Add(arg NullableCollectionLocation) *NullableCollectionLocationStream {
	return self.AddAll(arg)
}
func (self *NullableCollectionLocationStream) AddAll(arg ...NullableCollectionLocation) *NullableCollectionLocationStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableCollectionLocationStream) AddSafe(arg *NullableCollectionLocation) *NullableCollectionLocationStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableCollectionLocationStream) Aggregate(fn func(NullableCollectionLocation, NullableCollectionLocation) NullableCollectionLocation) *NullableCollectionLocationStream {
	result := NullableCollectionLocationStreamOf()
	self.ForEach(func(v NullableCollectionLocation, i int) {
		if i == 0 {
			result.Add(fn(NullableCollectionLocation{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableCollectionLocationStream) AllMatch(fn func(NullableCollectionLocation, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableCollectionLocationStream) AnyMatch(fn func(NullableCollectionLocation, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableCollectionLocationStream) Clone() *NullableCollectionLocationStream {
	temp := make([]NullableCollectionLocation, self.Len())
	copy(temp, *self)
	return (*NullableCollectionLocationStream)(&temp)
}
func (self *NullableCollectionLocationStream) Copy() *NullableCollectionLocationStream {
	return self.Clone()
}
func (self *NullableCollectionLocationStream) Concat(arg []NullableCollectionLocation) *NullableCollectionLocationStream {
	return self.AddAll(arg...)
}
func (self *NullableCollectionLocationStream) Contains(arg NullableCollectionLocation) bool {
	return self.FindIndex(func(_arg NullableCollectionLocation, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableCollectionLocationStream) Clean() *NullableCollectionLocationStream {
	*self = NullableCollectionLocationStreamOf()
	return self
}
func (self *NullableCollectionLocationStream) Delete(index int) *NullableCollectionLocationStream {
	return self.DeleteRange(index, index)
}
func (self *NullableCollectionLocationStream) DeleteRange(startIndex, endIndex int) *NullableCollectionLocationStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableCollectionLocationStream) Distinct() *NullableCollectionLocationStream {
	caches := map[string]bool{}
	result := NullableCollectionLocationStreamOf()
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
func (self *NullableCollectionLocationStream) Each(fn func(NullableCollectionLocation)) *NullableCollectionLocationStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableCollectionLocationStream) EachRight(fn func(NullableCollectionLocation)) *NullableCollectionLocationStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableCollectionLocationStream) Equals(arr []NullableCollectionLocation) bool {
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
func (self *NullableCollectionLocationStream) Filter(fn func(NullableCollectionLocation, int) bool) *NullableCollectionLocationStream {
	result := NullableCollectionLocationStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableCollectionLocationStream) FilterSlim(fn func(NullableCollectionLocation, int) bool) *NullableCollectionLocationStream {
	result := NullableCollectionLocationStreamOf()
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
func (self *NullableCollectionLocationStream) Find(fn func(NullableCollectionLocation, int) bool) *NullableCollectionLocation {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableCollectionLocationStream) FindOr(fn func(NullableCollectionLocation, int) bool, or NullableCollectionLocation) NullableCollectionLocation {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableCollectionLocationStream) FindIndex(fn func(NullableCollectionLocation, int) bool) int {
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
func (self *NullableCollectionLocationStream) First() *NullableCollectionLocation {
	return self.Get(0)
}
func (self *NullableCollectionLocationStream) FirstOr(arg NullableCollectionLocation) NullableCollectionLocation {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableCollectionLocationStream) ForEach(fn func(NullableCollectionLocation, int)) *NullableCollectionLocationStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableCollectionLocationStream) ForEachRight(fn func(NullableCollectionLocation, int)) *NullableCollectionLocationStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableCollectionLocationStream) GroupBy(fn func(NullableCollectionLocation, int) string) map[string][]NullableCollectionLocation {
	m := map[string][]NullableCollectionLocation{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableCollectionLocationStream) GroupByValues(fn func(NullableCollectionLocation, int) string) [][]NullableCollectionLocation {
	var tmp [][]NullableCollectionLocation
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableCollectionLocationStream) IndexOf(arg NullableCollectionLocation) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableCollectionLocationStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableCollectionLocationStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableCollectionLocationStream) Last() *NullableCollectionLocation {
	return self.Get(self.Len() - 1)
}
func (self *NullableCollectionLocationStream) LastOr(arg NullableCollectionLocation) NullableCollectionLocation {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableCollectionLocationStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableCollectionLocationStream) Limit(limit int) *NullableCollectionLocationStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableCollectionLocationStream) Map(fn func(NullableCollectionLocation, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLocationStream) Map2Int(fn func(NullableCollectionLocation, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLocationStream) Map2Int32(fn func(NullableCollectionLocation, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLocationStream) Map2Int64(fn func(NullableCollectionLocation, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLocationStream) Map2Float32(fn func(NullableCollectionLocation, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLocationStream) Map2Float64(fn func(NullableCollectionLocation, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLocationStream) Map2Bool(fn func(NullableCollectionLocation, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLocationStream) Map2Bytes(fn func(NullableCollectionLocation, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLocationStream) Map2String(fn func(NullableCollectionLocation, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableCollectionLocationStream) Max(fn func(NullableCollectionLocation, int) float64) *NullableCollectionLocation {
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
func (self *NullableCollectionLocationStream) Min(fn func(NullableCollectionLocation, int) float64) *NullableCollectionLocation {
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
func (self *NullableCollectionLocationStream) NoneMatch(fn func(NullableCollectionLocation, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableCollectionLocationStream) Get(index int) *NullableCollectionLocation {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableCollectionLocationStream) GetOr(index int, arg NullableCollectionLocation) NullableCollectionLocation {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableCollectionLocationStream) Peek(fn func(*NullableCollectionLocation, int)) *NullableCollectionLocationStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableCollectionLocationStream) Reduce(fn func(NullableCollectionLocation, NullableCollectionLocation, int) NullableCollectionLocation) *NullableCollectionLocationStream {
	return self.ReduceInit(fn, NullableCollectionLocation{})
}
func (self *NullableCollectionLocationStream) ReduceInit(fn func(NullableCollectionLocation, NullableCollectionLocation, int) NullableCollectionLocation, initialValue NullableCollectionLocation) *NullableCollectionLocationStream {
	result := NullableCollectionLocationStreamOf()
	self.ForEach(func(v NullableCollectionLocation, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableCollectionLocationStream) ReduceInterface(fn func(interface{}, NullableCollectionLocation, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableCollectionLocation{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableCollectionLocationStream) ReduceString(fn func(string, NullableCollectionLocation, int) string) []string {
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
func (self *NullableCollectionLocationStream) ReduceInt(fn func(int, NullableCollectionLocation, int) int) []int {
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
func (self *NullableCollectionLocationStream) ReduceInt32(fn func(int32, NullableCollectionLocation, int) int32) []int32 {
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
func (self *NullableCollectionLocationStream) ReduceInt64(fn func(int64, NullableCollectionLocation, int) int64) []int64 {
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
func (self *NullableCollectionLocationStream) ReduceFloat32(fn func(float32, NullableCollectionLocation, int) float32) []float32 {
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
func (self *NullableCollectionLocationStream) ReduceFloat64(fn func(float64, NullableCollectionLocation, int) float64) []float64 {
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
func (self *NullableCollectionLocationStream) ReduceBool(fn func(bool, NullableCollectionLocation, int) bool) []bool {
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
func (self *NullableCollectionLocationStream) Reverse() *NullableCollectionLocationStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableCollectionLocationStream) Replace(fn func(NullableCollectionLocation, int) NullableCollectionLocation) *NullableCollectionLocationStream {
	return self.ForEach(func(v NullableCollectionLocation, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableCollectionLocationStream) Select(fn func(NullableCollectionLocation) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableCollectionLocationStream) Set(index int, val NullableCollectionLocation) *NullableCollectionLocationStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableCollectionLocationStream) Skip(skip int) *NullableCollectionLocationStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableCollectionLocationStream) SkippingEach(fn func(NullableCollectionLocation, int) int) *NullableCollectionLocationStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableCollectionLocationStream) Slice(startIndex, n int) *NullableCollectionLocationStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableCollectionLocation{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableCollectionLocationStream) Sort(fn func(i, j int) bool) *NullableCollectionLocationStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableCollectionLocationStream) Tail() *NullableCollectionLocation {
	return self.Last()
}
func (self *NullableCollectionLocationStream) TailOr(arg NullableCollectionLocation) NullableCollectionLocation {
	return self.LastOr(arg)
}
func (self *NullableCollectionLocationStream) ToList() []NullableCollectionLocation {
	return self.Val()
}
func (self *NullableCollectionLocationStream) Unique() *NullableCollectionLocationStream {
	return self.Distinct()
}
func (self *NullableCollectionLocationStream) Val() []NullableCollectionLocation {
	if self == nil {
		return []NullableCollectionLocation{}
	}
	return *self.Copy()
}
func (self *NullableCollectionLocationStream) While(fn func(NullableCollectionLocation, int) bool) *NullableCollectionLocationStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableCollectionLocationStream) Where(fn func(NullableCollectionLocation) bool) *NullableCollectionLocationStream {
	result := NullableCollectionLocationStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableCollectionLocationStream) WhereSlim(fn func(NullableCollectionLocation) bool) *NullableCollectionLocationStream {
	result := NullableCollectionLocationStreamOf()
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
