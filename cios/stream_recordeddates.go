/*
 * Collection utility of RecordedDates Struct
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

type RecordedDatesStream []RecordedDates

func RecordedDatesStreamOf(arg ...RecordedDates) RecordedDatesStream {
	return arg
}
func RecordedDatesStreamFrom(arg []RecordedDates) RecordedDatesStream {
	return arg
}
func CreateRecordedDatesStream(arg ...RecordedDates) *RecordedDatesStream {
	tmp := RecordedDatesStreamOf(arg...)
	return &tmp
}
func GenerateRecordedDatesStream(arg []RecordedDates) *RecordedDatesStream {
	tmp := RecordedDatesStreamFrom(arg)
	return &tmp
}

func (self *RecordedDatesStream) Add(arg RecordedDates) *RecordedDatesStream {
	return self.AddAll(arg)
}
func (self *RecordedDatesStream) AddAll(arg ...RecordedDates) *RecordedDatesStream {
	*self = append(*self, arg...)
	return self
}
func (self *RecordedDatesStream) AddSafe(arg *RecordedDates) *RecordedDatesStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *RecordedDatesStream) Aggregate(fn func(RecordedDates, RecordedDates) RecordedDates) *RecordedDatesStream {
	result := RecordedDatesStreamOf()
	self.ForEach(func(v RecordedDates, i int) {
		if i == 0 {
			result.Add(fn(RecordedDates{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *RecordedDatesStream) AllMatch(fn func(RecordedDates, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *RecordedDatesStream) AnyMatch(fn func(RecordedDates, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *RecordedDatesStream) Clone() *RecordedDatesStream {
	temp := make([]RecordedDates, self.Len())
	copy(temp, *self)
	return (*RecordedDatesStream)(&temp)
}
func (self *RecordedDatesStream) Copy() *RecordedDatesStream {
	return self.Clone()
}
func (self *RecordedDatesStream) Concat(arg []RecordedDates) *RecordedDatesStream {
	return self.AddAll(arg...)
}
func (self *RecordedDatesStream) Contains(arg RecordedDates) bool {
	return self.FindIndex(func(_arg RecordedDates, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *RecordedDatesStream) Clean() *RecordedDatesStream {
	*self = RecordedDatesStreamOf()
	return self
}
func (self *RecordedDatesStream) Delete(index int) *RecordedDatesStream {
	return self.DeleteRange(index, index)
}
func (self *RecordedDatesStream) DeleteRange(startIndex, endIndex int) *RecordedDatesStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *RecordedDatesStream) Distinct() *RecordedDatesStream {
	caches := map[string]bool{}
	result := RecordedDatesStreamOf()
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
func (self *RecordedDatesStream) Each(fn func(RecordedDates)) *RecordedDatesStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *RecordedDatesStream) EachRight(fn func(RecordedDates)) *RecordedDatesStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *RecordedDatesStream) Equals(arr []RecordedDates) bool {
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
func (self *RecordedDatesStream) Filter(fn func(RecordedDates, int) bool) *RecordedDatesStream {
	result := RecordedDatesStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *RecordedDatesStream) FilterSlim(fn func(RecordedDates, int) bool) *RecordedDatesStream {
	result := RecordedDatesStreamOf()
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
func (self *RecordedDatesStream) Find(fn func(RecordedDates, int) bool) *RecordedDates {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *RecordedDatesStream) FindOr(fn func(RecordedDates, int) bool, or RecordedDates) RecordedDates {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *RecordedDatesStream) FindIndex(fn func(RecordedDates, int) bool) int {
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
func (self *RecordedDatesStream) First() *RecordedDates {
	return self.Get(0)
}
func (self *RecordedDatesStream) FirstOr(arg RecordedDates) RecordedDates {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *RecordedDatesStream) ForEach(fn func(RecordedDates, int)) *RecordedDatesStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *RecordedDatesStream) ForEachRight(fn func(RecordedDates, int)) *RecordedDatesStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *RecordedDatesStream) GroupBy(fn func(RecordedDates, int) string) map[string][]RecordedDates {
	m := map[string][]RecordedDates{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *RecordedDatesStream) GroupByValues(fn func(RecordedDates, int) string) [][]RecordedDates {
	var tmp [][]RecordedDates
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *RecordedDatesStream) IndexOf(arg RecordedDates) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *RecordedDatesStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *RecordedDatesStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *RecordedDatesStream) Last() *RecordedDates {
	return self.Get(self.Len() - 1)
}
func (self *RecordedDatesStream) LastOr(arg RecordedDates) RecordedDates {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *RecordedDatesStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *RecordedDatesStream) Limit(limit int) *RecordedDatesStream {
	self.Slice(0, limit)
	return self
}

func (self *RecordedDatesStream) Map(fn func(RecordedDates, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RecordedDatesStream) Map2Int(fn func(RecordedDates, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RecordedDatesStream) Map2Int32(fn func(RecordedDates, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RecordedDatesStream) Map2Int64(fn func(RecordedDates, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RecordedDatesStream) Map2Float32(fn func(RecordedDates, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RecordedDatesStream) Map2Float64(fn func(RecordedDates, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RecordedDatesStream) Map2Bool(fn func(RecordedDates, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RecordedDatesStream) Map2Bytes(fn func(RecordedDates, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RecordedDatesStream) Map2String(fn func(RecordedDates, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RecordedDatesStream) Max(fn func(RecordedDates, int) float64) *RecordedDates {
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
func (self *RecordedDatesStream) Min(fn func(RecordedDates, int) float64) *RecordedDates {
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
func (self *RecordedDatesStream) NoneMatch(fn func(RecordedDates, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *RecordedDatesStream) Get(index int) *RecordedDates {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *RecordedDatesStream) GetOr(index int, arg RecordedDates) RecordedDates {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *RecordedDatesStream) Peek(fn func(*RecordedDates, int)) *RecordedDatesStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *RecordedDatesStream) Reduce(fn func(RecordedDates, RecordedDates, int) RecordedDates) *RecordedDatesStream {
	return self.ReduceInit(fn, RecordedDates{})
}
func (self *RecordedDatesStream) ReduceInit(fn func(RecordedDates, RecordedDates, int) RecordedDates, initialValue RecordedDates) *RecordedDatesStream {
	result := RecordedDatesStreamOf()
	self.ForEach(func(v RecordedDates, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *RecordedDatesStream) ReduceInterface(fn func(interface{}, RecordedDates, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(RecordedDates{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *RecordedDatesStream) ReduceString(fn func(string, RecordedDates, int) string) []string {
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
func (self *RecordedDatesStream) ReduceInt(fn func(int, RecordedDates, int) int) []int {
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
func (self *RecordedDatesStream) ReduceInt32(fn func(int32, RecordedDates, int) int32) []int32 {
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
func (self *RecordedDatesStream) ReduceInt64(fn func(int64, RecordedDates, int) int64) []int64 {
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
func (self *RecordedDatesStream) ReduceFloat32(fn func(float32, RecordedDates, int) float32) []float32 {
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
func (self *RecordedDatesStream) ReduceFloat64(fn func(float64, RecordedDates, int) float64) []float64 {
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
func (self *RecordedDatesStream) ReduceBool(fn func(bool, RecordedDates, int) bool) []bool {
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
func (self *RecordedDatesStream) Reverse() *RecordedDatesStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *RecordedDatesStream) Replace(fn func(RecordedDates, int) RecordedDates) *RecordedDatesStream {
	return self.ForEach(func(v RecordedDates, i int) { self.Set(i, fn(v, i)) })
}
func (self *RecordedDatesStream) Select(fn func(RecordedDates) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *RecordedDatesStream) Set(index int, val RecordedDates) *RecordedDatesStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *RecordedDatesStream) Skip(skip int) *RecordedDatesStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *RecordedDatesStream) SkippingEach(fn func(RecordedDates, int) int) *RecordedDatesStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *RecordedDatesStream) Slice(startIndex, n int) *RecordedDatesStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []RecordedDates{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *RecordedDatesStream) Sort(fn func(i, j int) bool) *RecordedDatesStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *RecordedDatesStream) Tail() *RecordedDates {
	return self.Last()
}
func (self *RecordedDatesStream) TailOr(arg RecordedDates) RecordedDates {
	return self.LastOr(arg)
}
func (self *RecordedDatesStream) ToList() []RecordedDates {
	return self.Val()
}
func (self *RecordedDatesStream) Unique() *RecordedDatesStream {
	return self.Distinct()
}
func (self *RecordedDatesStream) Val() []RecordedDates {
	if self == nil {
		return []RecordedDates{}
	}
	return *self.Copy()
}
func (self *RecordedDatesStream) While(fn func(RecordedDates, int) bool) *RecordedDatesStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *RecordedDatesStream) Where(fn func(RecordedDates) bool) *RecordedDatesStream {
	result := RecordedDatesStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *RecordedDatesStream) WhereSlim(fn func(RecordedDates) bool) *RecordedDatesStream {
	result := RecordedDatesStreamOf()
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
