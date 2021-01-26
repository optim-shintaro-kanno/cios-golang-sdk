/*
 * Collection utility of NullableFloat64 Struct
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

type NullableFloat64Stream []NullableFloat64

func NullableFloat64StreamOf(arg ...NullableFloat64) NullableFloat64Stream {
	return arg
}
func NullableFloat64StreamFrom(arg []NullableFloat64) NullableFloat64Stream {
	return arg
}
func CreateNullableFloat64Stream(arg ...NullableFloat64) *NullableFloat64Stream {
	tmp := NullableFloat64StreamOf(arg...)
	return &tmp
}
func GenerateNullableFloat64Stream(arg []NullableFloat64) *NullableFloat64Stream {
	tmp := NullableFloat64StreamFrom(arg)
	return &tmp
}

func (self *NullableFloat64Stream) Add(arg NullableFloat64) *NullableFloat64Stream {
	return self.AddAll(arg)
}
func (self *NullableFloat64Stream) AddAll(arg ...NullableFloat64) *NullableFloat64Stream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableFloat64Stream) AddSafe(arg *NullableFloat64) *NullableFloat64Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableFloat64Stream) Aggregate(fn func(NullableFloat64, NullableFloat64) NullableFloat64) *NullableFloat64Stream {
	result := NullableFloat64StreamOf()
	self.ForEach(func(v NullableFloat64, i int) {
		if i == 0 {
			result.Add(fn(NullableFloat64{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableFloat64Stream) AllMatch(fn func(NullableFloat64, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableFloat64Stream) AnyMatch(fn func(NullableFloat64, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableFloat64Stream) Clone() *NullableFloat64Stream {
	temp := make([]NullableFloat64, self.Len())
	copy(temp, *self)
	return (*NullableFloat64Stream)(&temp)
}
func (self *NullableFloat64Stream) Copy() *NullableFloat64Stream {
	return self.Clone()
}
func (self *NullableFloat64Stream) Concat(arg []NullableFloat64) *NullableFloat64Stream {
	return self.AddAll(arg...)
}
func (self *NullableFloat64Stream) Contains(arg NullableFloat64) bool {
	return self.FindIndex(func(_arg NullableFloat64, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableFloat64Stream) Clean() *NullableFloat64Stream {
	*self = NullableFloat64StreamOf()
	return self
}
func (self *NullableFloat64Stream) Delete(index int) *NullableFloat64Stream {
	return self.DeleteRange(index, index)
}
func (self *NullableFloat64Stream) DeleteRange(startIndex, endIndex int) *NullableFloat64Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableFloat64Stream) Distinct() *NullableFloat64Stream {
	caches := map[string]bool{}
	result := NullableFloat64StreamOf()
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
func (self *NullableFloat64Stream) Each(fn func(NullableFloat64)) *NullableFloat64Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableFloat64Stream) EachRight(fn func(NullableFloat64)) *NullableFloat64Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableFloat64Stream) Equals(arr []NullableFloat64) bool {
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
func (self *NullableFloat64Stream) Filter(fn func(NullableFloat64, int) bool) *NullableFloat64Stream {
	result := NullableFloat64StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableFloat64Stream) FilterSlim(fn func(NullableFloat64, int) bool) *NullableFloat64Stream {
	result := NullableFloat64StreamOf()
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
func (self *NullableFloat64Stream) Find(fn func(NullableFloat64, int) bool) *NullableFloat64 {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableFloat64Stream) FindOr(fn func(NullableFloat64, int) bool, or NullableFloat64) NullableFloat64 {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableFloat64Stream) FindIndex(fn func(NullableFloat64, int) bool) int {
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
func (self *NullableFloat64Stream) First() *NullableFloat64 {
	return self.Get(0)
}
func (self *NullableFloat64Stream) FirstOr(arg NullableFloat64) NullableFloat64 {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableFloat64Stream) ForEach(fn func(NullableFloat64, int)) *NullableFloat64Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableFloat64Stream) ForEachRight(fn func(NullableFloat64, int)) *NullableFloat64Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableFloat64Stream) GroupBy(fn func(NullableFloat64, int) string) map[string][]NullableFloat64 {
	m := map[string][]NullableFloat64{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableFloat64Stream) GroupByValues(fn func(NullableFloat64, int) string) [][]NullableFloat64 {
	var tmp [][]NullableFloat64
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableFloat64Stream) IndexOf(arg NullableFloat64) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableFloat64Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableFloat64Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableFloat64Stream) Last() *NullableFloat64 {
	return self.Get(self.Len() - 1)
}
func (self *NullableFloat64Stream) LastOr(arg NullableFloat64) NullableFloat64 {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableFloat64Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableFloat64Stream) Limit(limit int) *NullableFloat64Stream {
	self.Slice(0, limit)
	return self
}

func (self *NullableFloat64Stream) Map(fn func(NullableFloat64, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableFloat64Stream) Map2Int(fn func(NullableFloat64, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableFloat64Stream) Map2Int32(fn func(NullableFloat64, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableFloat64Stream) Map2Int64(fn func(NullableFloat64, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableFloat64Stream) Map2Float32(fn func(NullableFloat64, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableFloat64Stream) Map2Float64(fn func(NullableFloat64, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableFloat64Stream) Map2Bool(fn func(NullableFloat64, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableFloat64Stream) Map2Bytes(fn func(NullableFloat64, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableFloat64Stream) Map2String(fn func(NullableFloat64, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableFloat64Stream) Max(fn func(NullableFloat64, int) float64) *NullableFloat64 {
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
func (self *NullableFloat64Stream) Min(fn func(NullableFloat64, int) float64) *NullableFloat64 {
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
func (self *NullableFloat64Stream) NoneMatch(fn func(NullableFloat64, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableFloat64Stream) Get(index int) *NullableFloat64 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableFloat64Stream) GetOr(index int, arg NullableFloat64) NullableFloat64 {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableFloat64Stream) Peek(fn func(*NullableFloat64, int)) *NullableFloat64Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableFloat64Stream) Reduce(fn func(NullableFloat64, NullableFloat64, int) NullableFloat64) *NullableFloat64Stream {
	return self.ReduceInit(fn, NullableFloat64{})
}
func (self *NullableFloat64Stream) ReduceInit(fn func(NullableFloat64, NullableFloat64, int) NullableFloat64, initialValue NullableFloat64) *NullableFloat64Stream {
	result := NullableFloat64StreamOf()
	self.ForEach(func(v NullableFloat64, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableFloat64Stream) ReduceInterface(fn func(interface{}, NullableFloat64, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableFloat64{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableFloat64Stream) ReduceString(fn func(string, NullableFloat64, int) string) []string {
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
func (self *NullableFloat64Stream) ReduceInt(fn func(int, NullableFloat64, int) int) []int {
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
func (self *NullableFloat64Stream) ReduceInt32(fn func(int32, NullableFloat64, int) int32) []int32 {
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
func (self *NullableFloat64Stream) ReduceInt64(fn func(int64, NullableFloat64, int) int64) []int64 {
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
func (self *NullableFloat64Stream) ReduceFloat32(fn func(float32, NullableFloat64, int) float32) []float32 {
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
func (self *NullableFloat64Stream) ReduceFloat64(fn func(float64, NullableFloat64, int) float64) []float64 {
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
func (self *NullableFloat64Stream) ReduceBool(fn func(bool, NullableFloat64, int) bool) []bool {
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
func (self *NullableFloat64Stream) Reverse() *NullableFloat64Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableFloat64Stream) Replace(fn func(NullableFloat64, int) NullableFloat64) *NullableFloat64Stream {
	return self.ForEach(func(v NullableFloat64, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableFloat64Stream) Select(fn func(NullableFloat64) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableFloat64Stream) Set(index int, val NullableFloat64) *NullableFloat64Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableFloat64Stream) Skip(skip int) *NullableFloat64Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableFloat64Stream) SkippingEach(fn func(NullableFloat64, int) int) *NullableFloat64Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableFloat64Stream) Slice(startIndex, n int) *NullableFloat64Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableFloat64{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableFloat64Stream) Sort(fn func(i, j int) bool) *NullableFloat64Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableFloat64Stream) Tail() *NullableFloat64 {
	return self.Last()
}
func (self *NullableFloat64Stream) TailOr(arg NullableFloat64) NullableFloat64 {
	return self.LastOr(arg)
}
func (self *NullableFloat64Stream) ToList() []NullableFloat64 {
	return self.Val()
}
func (self *NullableFloat64Stream) Unique() *NullableFloat64Stream {
	return self.Distinct()
}
func (self *NullableFloat64Stream) Val() []NullableFloat64 {
	if self == nil {
		return []NullableFloat64{}
	}
	return *self.Copy()
}
func (self *NullableFloat64Stream) While(fn func(NullableFloat64, int) bool) *NullableFloat64Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableFloat64Stream) Where(fn func(NullableFloat64) bool) *NullableFloat64Stream {
	result := NullableFloat64StreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableFloat64Stream) WhereSlim(fn func(NullableFloat64) bool) *NullableFloat64Stream {
	result := NullableFloat64StreamOf()
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
