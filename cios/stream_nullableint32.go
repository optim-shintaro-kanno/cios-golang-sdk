/*
 * Collection utility of NullableInt32 Struct
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

type NullableInt32Stream []NullableInt32

func NullableInt32StreamOf(arg ...NullableInt32) NullableInt32Stream {
	return arg
}
func NullableInt32StreamFrom(arg []NullableInt32) NullableInt32Stream {
	return arg
}
func CreateNullableInt32Stream(arg ...NullableInt32) *NullableInt32Stream {
	tmp := NullableInt32StreamOf(arg...)
	return &tmp
}
func GenerateNullableInt32Stream(arg []NullableInt32) *NullableInt32Stream {
	tmp := NullableInt32StreamFrom(arg)
	return &tmp
}

func (self *NullableInt32Stream) Add(arg NullableInt32) *NullableInt32Stream {
	return self.AddAll(arg)
}
func (self *NullableInt32Stream) AddAll(arg ...NullableInt32) *NullableInt32Stream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableInt32Stream) AddSafe(arg *NullableInt32) *NullableInt32Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableInt32Stream) Aggregate(fn func(NullableInt32, NullableInt32) NullableInt32) *NullableInt32Stream {
	result := NullableInt32StreamOf()
	self.ForEach(func(v NullableInt32, i int) {
		if i == 0 {
			result.Add(fn(NullableInt32{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableInt32Stream) AllMatch(fn func(NullableInt32, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableInt32Stream) AnyMatch(fn func(NullableInt32, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableInt32Stream) Clone() *NullableInt32Stream {
	temp := make([]NullableInt32, self.Len())
	copy(temp, *self)
	return (*NullableInt32Stream)(&temp)
}
func (self *NullableInt32Stream) Copy() *NullableInt32Stream {
	return self.Clone()
}
func (self *NullableInt32Stream) Concat(arg []NullableInt32) *NullableInt32Stream {
	return self.AddAll(arg...)
}
func (self *NullableInt32Stream) Contains(arg NullableInt32) bool {
	return self.FindIndex(func(_arg NullableInt32, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableInt32Stream) Clean() *NullableInt32Stream {
	*self = NullableInt32StreamOf()
	return self
}
func (self *NullableInt32Stream) Delete(index int) *NullableInt32Stream {
	return self.DeleteRange(index, index)
}
func (self *NullableInt32Stream) DeleteRange(startIndex, endIndex int) *NullableInt32Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableInt32Stream) Distinct() *NullableInt32Stream {
	caches := map[string]bool{}
	result := NullableInt32StreamOf()
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
func (self *NullableInt32Stream) Each(fn func(NullableInt32)) *NullableInt32Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableInt32Stream) EachRight(fn func(NullableInt32)) *NullableInt32Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableInt32Stream) Equals(arr []NullableInt32) bool {
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
func (self *NullableInt32Stream) Filter(fn func(NullableInt32, int) bool) *NullableInt32Stream {
	result := NullableInt32StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableInt32Stream) FilterSlim(fn func(NullableInt32, int) bool) *NullableInt32Stream {
	result := NullableInt32StreamOf()
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
func (self *NullableInt32Stream) Find(fn func(NullableInt32, int) bool) *NullableInt32 {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableInt32Stream) FindOr(fn func(NullableInt32, int) bool, or NullableInt32) NullableInt32 {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableInt32Stream) FindIndex(fn func(NullableInt32, int) bool) int {
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
func (self *NullableInt32Stream) First() *NullableInt32 {
	return self.Get(0)
}
func (self *NullableInt32Stream) FirstOr(arg NullableInt32) NullableInt32 {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableInt32Stream) ForEach(fn func(NullableInt32, int)) *NullableInt32Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableInt32Stream) ForEachRight(fn func(NullableInt32, int)) *NullableInt32Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableInt32Stream) GroupBy(fn func(NullableInt32, int) string) map[string][]NullableInt32 {
	m := map[string][]NullableInt32{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableInt32Stream) GroupByValues(fn func(NullableInt32, int) string) [][]NullableInt32 {
	var tmp [][]NullableInt32
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableInt32Stream) IndexOf(arg NullableInt32) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableInt32Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableInt32Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableInt32Stream) Last() *NullableInt32 {
	return self.Get(self.Len() - 1)
}
func (self *NullableInt32Stream) LastOr(arg NullableInt32) NullableInt32 {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableInt32Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableInt32Stream) Limit(limit int) *NullableInt32Stream {
	self.Slice(0, limit)
	return self
}

func (self *NullableInt32Stream) Map(fn func(NullableInt32, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInt32Stream) Map2Int(fn func(NullableInt32, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInt32Stream) Map2Int32(fn func(NullableInt32, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInt32Stream) Map2Int64(fn func(NullableInt32, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInt32Stream) Map2Float32(fn func(NullableInt32, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInt32Stream) Map2Float64(fn func(NullableInt32, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInt32Stream) Map2Bool(fn func(NullableInt32, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInt32Stream) Map2Bytes(fn func(NullableInt32, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInt32Stream) Map2String(fn func(NullableInt32, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableInt32Stream) Max(fn func(NullableInt32, int) float64) *NullableInt32 {
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
func (self *NullableInt32Stream) Min(fn func(NullableInt32, int) float64) *NullableInt32 {
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
func (self *NullableInt32Stream) NoneMatch(fn func(NullableInt32, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableInt32Stream) Get(index int) *NullableInt32 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableInt32Stream) GetOr(index int, arg NullableInt32) NullableInt32 {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableInt32Stream) Peek(fn func(*NullableInt32, int)) *NullableInt32Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableInt32Stream) Reduce(fn func(NullableInt32, NullableInt32, int) NullableInt32) *NullableInt32Stream {
	return self.ReduceInit(fn, NullableInt32{})
}
func (self *NullableInt32Stream) ReduceInit(fn func(NullableInt32, NullableInt32, int) NullableInt32, initialValue NullableInt32) *NullableInt32Stream {
	result := NullableInt32StreamOf()
	self.ForEach(func(v NullableInt32, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableInt32Stream) ReduceInterface(fn func(interface{}, NullableInt32, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableInt32{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableInt32Stream) ReduceString(fn func(string, NullableInt32, int) string) []string {
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
func (self *NullableInt32Stream) ReduceInt(fn func(int, NullableInt32, int) int) []int {
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
func (self *NullableInt32Stream) ReduceInt32(fn func(int32, NullableInt32, int) int32) []int32 {
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
func (self *NullableInt32Stream) ReduceInt64(fn func(int64, NullableInt32, int) int64) []int64 {
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
func (self *NullableInt32Stream) ReduceFloat32(fn func(float32, NullableInt32, int) float32) []float32 {
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
func (self *NullableInt32Stream) ReduceFloat64(fn func(float64, NullableInt32, int) float64) []float64 {
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
func (self *NullableInt32Stream) ReduceBool(fn func(bool, NullableInt32, int) bool) []bool {
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
func (self *NullableInt32Stream) Reverse() *NullableInt32Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableInt32Stream) Replace(fn func(NullableInt32, int) NullableInt32) *NullableInt32Stream {
	return self.ForEach(func(v NullableInt32, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableInt32Stream) Select(fn func(NullableInt32) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableInt32Stream) Set(index int, val NullableInt32) *NullableInt32Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableInt32Stream) Skip(skip int) *NullableInt32Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableInt32Stream) SkippingEach(fn func(NullableInt32, int) int) *NullableInt32Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableInt32Stream) Slice(startIndex, n int) *NullableInt32Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableInt32{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableInt32Stream) Sort(fn func(i, j int) bool) *NullableInt32Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableInt32Stream) Tail() *NullableInt32 {
	return self.Last()
}
func (self *NullableInt32Stream) TailOr(arg NullableInt32) NullableInt32 {
	return self.LastOr(arg)
}
func (self *NullableInt32Stream) ToList() []NullableInt32 {
	return self.Val()
}
func (self *NullableInt32Stream) Unique() *NullableInt32Stream {
	return self.Distinct()
}
func (self *NullableInt32Stream) Val() []NullableInt32 {
	if self == nil {
		return []NullableInt32{}
	}
	return *self.Copy()
}
func (self *NullableInt32Stream) While(fn func(NullableInt32, int) bool) *NullableInt32Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableInt32Stream) Where(fn func(NullableInt32) bool) *NullableInt32Stream {
	result := NullableInt32StreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableInt32Stream) WhereSlim(fn func(NullableInt32) bool) *NullableInt32Stream {
	result := NullableInt32StreamOf()
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
