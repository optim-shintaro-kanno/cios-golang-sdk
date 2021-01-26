/*
 * Collection utility of MultiplePoint Struct
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

type MultiplePointStream []MultiplePoint

func MultiplePointStreamOf(arg ...MultiplePoint) MultiplePointStream {
	return arg
}
func MultiplePointStreamFrom(arg []MultiplePoint) MultiplePointStream {
	return arg
}
func CreateMultiplePointStream(arg ...MultiplePoint) *MultiplePointStream {
	tmp := MultiplePointStreamOf(arg...)
	return &tmp
}
func GenerateMultiplePointStream(arg []MultiplePoint) *MultiplePointStream {
	tmp := MultiplePointStreamFrom(arg)
	return &tmp
}

func (self *MultiplePointStream) Add(arg MultiplePoint) *MultiplePointStream {
	return self.AddAll(arg)
}
func (self *MultiplePointStream) AddAll(arg ...MultiplePoint) *MultiplePointStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultiplePointStream) AddSafe(arg *MultiplePoint) *MultiplePointStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultiplePointStream) Aggregate(fn func(MultiplePoint, MultiplePoint) MultiplePoint) *MultiplePointStream {
	result := MultiplePointStreamOf()
	self.ForEach(func(v MultiplePoint, i int) {
		if i == 0 {
			result.Add(fn(MultiplePoint{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultiplePointStream) AllMatch(fn func(MultiplePoint, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultiplePointStream) AnyMatch(fn func(MultiplePoint, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultiplePointStream) Clone() *MultiplePointStream {
	temp := make([]MultiplePoint, self.Len())
	copy(temp, *self)
	return (*MultiplePointStream)(&temp)
}
func (self *MultiplePointStream) Copy() *MultiplePointStream {
	return self.Clone()
}
func (self *MultiplePointStream) Concat(arg []MultiplePoint) *MultiplePointStream {
	return self.AddAll(arg...)
}
func (self *MultiplePointStream) Contains(arg MultiplePoint) bool {
	return self.FindIndex(func(_arg MultiplePoint, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultiplePointStream) Clean() *MultiplePointStream {
	*self = MultiplePointStreamOf()
	return self
}
func (self *MultiplePointStream) Delete(index int) *MultiplePointStream {
	return self.DeleteRange(index, index)
}
func (self *MultiplePointStream) DeleteRange(startIndex, endIndex int) *MultiplePointStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultiplePointStream) Distinct() *MultiplePointStream {
	caches := map[string]bool{}
	result := MultiplePointStreamOf()
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
func (self *MultiplePointStream) Each(fn func(MultiplePoint)) *MultiplePointStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultiplePointStream) EachRight(fn func(MultiplePoint)) *MultiplePointStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultiplePointStream) Equals(arr []MultiplePoint) bool {
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
func (self *MultiplePointStream) Filter(fn func(MultiplePoint, int) bool) *MultiplePointStream {
	result := MultiplePointStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultiplePointStream) FilterSlim(fn func(MultiplePoint, int) bool) *MultiplePointStream {
	result := MultiplePointStreamOf()
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
func (self *MultiplePointStream) Find(fn func(MultiplePoint, int) bool) *MultiplePoint {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultiplePointStream) FindOr(fn func(MultiplePoint, int) bool, or MultiplePoint) MultiplePoint {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultiplePointStream) FindIndex(fn func(MultiplePoint, int) bool) int {
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
func (self *MultiplePointStream) First() *MultiplePoint {
	return self.Get(0)
}
func (self *MultiplePointStream) FirstOr(arg MultiplePoint) MultiplePoint {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultiplePointStream) ForEach(fn func(MultiplePoint, int)) *MultiplePointStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultiplePointStream) ForEachRight(fn func(MultiplePoint, int)) *MultiplePointStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultiplePointStream) GroupBy(fn func(MultiplePoint, int) string) map[string][]MultiplePoint {
	m := map[string][]MultiplePoint{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultiplePointStream) GroupByValues(fn func(MultiplePoint, int) string) [][]MultiplePoint {
	var tmp [][]MultiplePoint
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultiplePointStream) IndexOf(arg MultiplePoint) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultiplePointStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultiplePointStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultiplePointStream) Last() *MultiplePoint {
	return self.Get(self.Len() - 1)
}
func (self *MultiplePointStream) LastOr(arg MultiplePoint) MultiplePoint {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultiplePointStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultiplePointStream) Limit(limit int) *MultiplePointStream {
	self.Slice(0, limit)
	return self
}

func (self *MultiplePointStream) Map(fn func(MultiplePoint, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultiplePointStream) Map2Int(fn func(MultiplePoint, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultiplePointStream) Map2Int32(fn func(MultiplePoint, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultiplePointStream) Map2Int64(fn func(MultiplePoint, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultiplePointStream) Map2Float32(fn func(MultiplePoint, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultiplePointStream) Map2Float64(fn func(MultiplePoint, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultiplePointStream) Map2Bool(fn func(MultiplePoint, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultiplePointStream) Map2Bytes(fn func(MultiplePoint, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultiplePointStream) Map2String(fn func(MultiplePoint, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultiplePointStream) Max(fn func(MultiplePoint, int) float64) *MultiplePoint {
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
func (self *MultiplePointStream) Min(fn func(MultiplePoint, int) float64) *MultiplePoint {
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
func (self *MultiplePointStream) NoneMatch(fn func(MultiplePoint, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultiplePointStream) Get(index int) *MultiplePoint {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultiplePointStream) GetOr(index int, arg MultiplePoint) MultiplePoint {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultiplePointStream) Peek(fn func(*MultiplePoint, int)) *MultiplePointStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *MultiplePointStream) Reduce(fn func(MultiplePoint, MultiplePoint, int) MultiplePoint) *MultiplePointStream {
	return self.ReduceInit(fn, MultiplePoint{})
}
func (self *MultiplePointStream) ReduceInit(fn func(MultiplePoint, MultiplePoint, int) MultiplePoint, initialValue MultiplePoint) *MultiplePointStream {
	result := MultiplePointStreamOf()
	self.ForEach(func(v MultiplePoint, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultiplePointStream) ReduceInterface(fn func(interface{}, MultiplePoint, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultiplePoint{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultiplePointStream) ReduceString(fn func(string, MultiplePoint, int) string) []string {
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
func (self *MultiplePointStream) ReduceInt(fn func(int, MultiplePoint, int) int) []int {
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
func (self *MultiplePointStream) ReduceInt32(fn func(int32, MultiplePoint, int) int32) []int32 {
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
func (self *MultiplePointStream) ReduceInt64(fn func(int64, MultiplePoint, int) int64) []int64 {
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
func (self *MultiplePointStream) ReduceFloat32(fn func(float32, MultiplePoint, int) float32) []float32 {
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
func (self *MultiplePointStream) ReduceFloat64(fn func(float64, MultiplePoint, int) float64) []float64 {
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
func (self *MultiplePointStream) ReduceBool(fn func(bool, MultiplePoint, int) bool) []bool {
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
func (self *MultiplePointStream) Reverse() *MultiplePointStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultiplePointStream) Replace(fn func(MultiplePoint, int) MultiplePoint) *MultiplePointStream {
	return self.ForEach(func(v MultiplePoint, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultiplePointStream) Select(fn func(MultiplePoint) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultiplePointStream) Set(index int, val MultiplePoint) *MultiplePointStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultiplePointStream) Skip(skip int) *MultiplePointStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultiplePointStream) SkippingEach(fn func(MultiplePoint, int) int) *MultiplePointStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultiplePointStream) Slice(startIndex, n int) *MultiplePointStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultiplePoint{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultiplePointStream) Sort(fn func(i, j int) bool) *MultiplePointStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultiplePointStream) Tail() *MultiplePoint {
	return self.Last()
}
func (self *MultiplePointStream) TailOr(arg MultiplePoint) MultiplePoint {
	return self.LastOr(arg)
}
func (self *MultiplePointStream) ToList() []MultiplePoint {
	return self.Val()
}
func (self *MultiplePointStream) Unique() *MultiplePointStream {
	return self.Distinct()
}
func (self *MultiplePointStream) Val() []MultiplePoint {
	if self == nil {
		return []MultiplePoint{}
	}
	return *self.Copy()
}
func (self *MultiplePointStream) While(fn func(MultiplePoint, int) bool) *MultiplePointStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultiplePointStream) Where(fn func(MultiplePoint) bool) *MultiplePointStream {
	result := MultiplePointStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultiplePointStream) WhereSlim(fn func(MultiplePoint) bool) *MultiplePointStream {
	result := MultiplePointStreamOf()
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
