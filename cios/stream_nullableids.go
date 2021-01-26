/*
 * Collection utility of NullableIds Struct
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

type NullableIdsStream []NullableIds

func NullableIdsStreamOf(arg ...NullableIds) NullableIdsStream {
	return arg
}
func NullableIdsStreamFrom(arg []NullableIds) NullableIdsStream {
	return arg
}
func CreateNullableIdsStream(arg ...NullableIds) *NullableIdsStream {
	tmp := NullableIdsStreamOf(arg...)
	return &tmp
}
func GenerateNullableIdsStream(arg []NullableIds) *NullableIdsStream {
	tmp := NullableIdsStreamFrom(arg)
	return &tmp
}

func (self *NullableIdsStream) Add(arg NullableIds) *NullableIdsStream {
	return self.AddAll(arg)
}
func (self *NullableIdsStream) AddAll(arg ...NullableIds) *NullableIdsStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableIdsStream) AddSafe(arg *NullableIds) *NullableIdsStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableIdsStream) Aggregate(fn func(NullableIds, NullableIds) NullableIds) *NullableIdsStream {
	result := NullableIdsStreamOf()
	self.ForEach(func(v NullableIds, i int) {
		if i == 0 {
			result.Add(fn(NullableIds{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableIdsStream) AllMatch(fn func(NullableIds, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableIdsStream) AnyMatch(fn func(NullableIds, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableIdsStream) Clone() *NullableIdsStream {
	temp := make([]NullableIds, self.Len())
	copy(temp, *self)
	return (*NullableIdsStream)(&temp)
}
func (self *NullableIdsStream) Copy() *NullableIdsStream {
	return self.Clone()
}
func (self *NullableIdsStream) Concat(arg []NullableIds) *NullableIdsStream {
	return self.AddAll(arg...)
}
func (self *NullableIdsStream) Contains(arg NullableIds) bool {
	return self.FindIndex(func(_arg NullableIds, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableIdsStream) Clean() *NullableIdsStream {
	*self = NullableIdsStreamOf()
	return self
}
func (self *NullableIdsStream) Delete(index int) *NullableIdsStream {
	return self.DeleteRange(index, index)
}
func (self *NullableIdsStream) DeleteRange(startIndex, endIndex int) *NullableIdsStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableIdsStream) Distinct() *NullableIdsStream {
	caches := map[string]bool{}
	result := NullableIdsStreamOf()
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
func (self *NullableIdsStream) Each(fn func(NullableIds)) *NullableIdsStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableIdsStream) EachRight(fn func(NullableIds)) *NullableIdsStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableIdsStream) Equals(arr []NullableIds) bool {
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
func (self *NullableIdsStream) Filter(fn func(NullableIds, int) bool) *NullableIdsStream {
	result := NullableIdsStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableIdsStream) FilterSlim(fn func(NullableIds, int) bool) *NullableIdsStream {
	result := NullableIdsStreamOf()
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
func (self *NullableIdsStream) Find(fn func(NullableIds, int) bool) *NullableIds {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableIdsStream) FindOr(fn func(NullableIds, int) bool, or NullableIds) NullableIds {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableIdsStream) FindIndex(fn func(NullableIds, int) bool) int {
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
func (self *NullableIdsStream) First() *NullableIds {
	return self.Get(0)
}
func (self *NullableIdsStream) FirstOr(arg NullableIds) NullableIds {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableIdsStream) ForEach(fn func(NullableIds, int)) *NullableIdsStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableIdsStream) ForEachRight(fn func(NullableIds, int)) *NullableIdsStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableIdsStream) GroupBy(fn func(NullableIds, int) string) map[string][]NullableIds {
	m := map[string][]NullableIds{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableIdsStream) GroupByValues(fn func(NullableIds, int) string) [][]NullableIds {
	var tmp [][]NullableIds
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableIdsStream) IndexOf(arg NullableIds) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableIdsStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableIdsStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableIdsStream) Last() *NullableIds {
	return self.Get(self.Len() - 1)
}
func (self *NullableIdsStream) LastOr(arg NullableIds) NullableIds {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableIdsStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableIdsStream) Limit(limit int) *NullableIdsStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableIdsStream) Map(fn func(NullableIds, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableIdsStream) Map2Int(fn func(NullableIds, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableIdsStream) Map2Int32(fn func(NullableIds, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableIdsStream) Map2Int64(fn func(NullableIds, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableIdsStream) Map2Float32(fn func(NullableIds, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableIdsStream) Map2Float64(fn func(NullableIds, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableIdsStream) Map2Bool(fn func(NullableIds, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableIdsStream) Map2Bytes(fn func(NullableIds, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableIdsStream) Map2String(fn func(NullableIds, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableIdsStream) Max(fn func(NullableIds, int) float64) *NullableIds {
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
func (self *NullableIdsStream) Min(fn func(NullableIds, int) float64) *NullableIds {
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
func (self *NullableIdsStream) NoneMatch(fn func(NullableIds, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableIdsStream) Get(index int) *NullableIds {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableIdsStream) GetOr(index int, arg NullableIds) NullableIds {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableIdsStream) Peek(fn func(*NullableIds, int)) *NullableIdsStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableIdsStream) Reduce(fn func(NullableIds, NullableIds, int) NullableIds) *NullableIdsStream {
	return self.ReduceInit(fn, NullableIds{})
}
func (self *NullableIdsStream) ReduceInit(fn func(NullableIds, NullableIds, int) NullableIds, initialValue NullableIds) *NullableIdsStream {
	result := NullableIdsStreamOf()
	self.ForEach(func(v NullableIds, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableIdsStream) ReduceInterface(fn func(interface{}, NullableIds, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableIds{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableIdsStream) ReduceString(fn func(string, NullableIds, int) string) []string {
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
func (self *NullableIdsStream) ReduceInt(fn func(int, NullableIds, int) int) []int {
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
func (self *NullableIdsStream) ReduceInt32(fn func(int32, NullableIds, int) int32) []int32 {
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
func (self *NullableIdsStream) ReduceInt64(fn func(int64, NullableIds, int) int64) []int64 {
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
func (self *NullableIdsStream) ReduceFloat32(fn func(float32, NullableIds, int) float32) []float32 {
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
func (self *NullableIdsStream) ReduceFloat64(fn func(float64, NullableIds, int) float64) []float64 {
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
func (self *NullableIdsStream) ReduceBool(fn func(bool, NullableIds, int) bool) []bool {
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
func (self *NullableIdsStream) Reverse() *NullableIdsStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableIdsStream) Replace(fn func(NullableIds, int) NullableIds) *NullableIdsStream {
	return self.ForEach(func(v NullableIds, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableIdsStream) Select(fn func(NullableIds) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableIdsStream) Set(index int, val NullableIds) *NullableIdsStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableIdsStream) Skip(skip int) *NullableIdsStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableIdsStream) SkippingEach(fn func(NullableIds, int) int) *NullableIdsStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableIdsStream) Slice(startIndex, n int) *NullableIdsStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableIds{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableIdsStream) Sort(fn func(i, j int) bool) *NullableIdsStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableIdsStream) Tail() *NullableIds {
	return self.Last()
}
func (self *NullableIdsStream) TailOr(arg NullableIds) NullableIds {
	return self.LastOr(arg)
}
func (self *NullableIdsStream) ToList() []NullableIds {
	return self.Val()
}
func (self *NullableIdsStream) Unique() *NullableIdsStream {
	return self.Distinct()
}
func (self *NullableIdsStream) Val() []NullableIds {
	if self == nil {
		return []NullableIds{}
	}
	return *self.Copy()
}
func (self *NullableIdsStream) While(fn func(NullableIds, int) bool) *NullableIdsStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableIdsStream) Where(fn func(NullableIds) bool) *NullableIdsStream {
	result := NullableIdsStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableIdsStream) WhereSlim(fn func(NullableIds) bool) *NullableIdsStream {
	result := NullableIdsStreamOf()
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
