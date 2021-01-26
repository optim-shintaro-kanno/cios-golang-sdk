/*
 * Collection utility of MultipleDataStoreChannel Struct
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

type MultipleDataStoreChannelStream []MultipleDataStoreChannel

func MultipleDataStoreChannelStreamOf(arg ...MultipleDataStoreChannel) MultipleDataStoreChannelStream {
	return arg
}
func MultipleDataStoreChannelStreamFrom(arg []MultipleDataStoreChannel) MultipleDataStoreChannelStream {
	return arg
}
func CreateMultipleDataStoreChannelStream(arg ...MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	tmp := MultipleDataStoreChannelStreamOf(arg...)
	return &tmp
}
func GenerateMultipleDataStoreChannelStream(arg []MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	tmp := MultipleDataStoreChannelStreamFrom(arg)
	return &tmp
}

func (self *MultipleDataStoreChannelStream) Add(arg MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	return self.AddAll(arg)
}
func (self *MultipleDataStoreChannelStream) AddAll(arg ...MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultipleDataStoreChannelStream) AddSafe(arg *MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultipleDataStoreChannelStream) Aggregate(fn func(MultipleDataStoreChannel, MultipleDataStoreChannel) MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	result := MultipleDataStoreChannelStreamOf()
	self.ForEach(func(v MultipleDataStoreChannel, i int) {
		if i == 0 {
			result.Add(fn(MultipleDataStoreChannel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultipleDataStoreChannelStream) AllMatch(fn func(MultipleDataStoreChannel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultipleDataStoreChannelStream) AnyMatch(fn func(MultipleDataStoreChannel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultipleDataStoreChannelStream) Clone() *MultipleDataStoreChannelStream {
	temp := make([]MultipleDataStoreChannel, self.Len())
	copy(temp, *self)
	return (*MultipleDataStoreChannelStream)(&temp)
}
func (self *MultipleDataStoreChannelStream) Copy() *MultipleDataStoreChannelStream {
	return self.Clone()
}
func (self *MultipleDataStoreChannelStream) Concat(arg []MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	return self.AddAll(arg...)
}
func (self *MultipleDataStoreChannelStream) Contains(arg MultipleDataStoreChannel) bool {
	return self.FindIndex(func(_arg MultipleDataStoreChannel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultipleDataStoreChannelStream) Clean() *MultipleDataStoreChannelStream {
	*self = MultipleDataStoreChannelStreamOf()
	return self
}
func (self *MultipleDataStoreChannelStream) Delete(index int) *MultipleDataStoreChannelStream {
	return self.DeleteRange(index, index)
}
func (self *MultipleDataStoreChannelStream) DeleteRange(startIndex, endIndex int) *MultipleDataStoreChannelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultipleDataStoreChannelStream) Distinct() *MultipleDataStoreChannelStream {
	caches := map[string]bool{}
	result := MultipleDataStoreChannelStreamOf()
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
func (self *MultipleDataStoreChannelStream) Each(fn func(MultipleDataStoreChannel)) *MultipleDataStoreChannelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultipleDataStoreChannelStream) EachRight(fn func(MultipleDataStoreChannel)) *MultipleDataStoreChannelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultipleDataStoreChannelStream) Equals(arr []MultipleDataStoreChannel) bool {
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
func (self *MultipleDataStoreChannelStream) Filter(fn func(MultipleDataStoreChannel, int) bool) *MultipleDataStoreChannelStream {
	result := MultipleDataStoreChannelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDataStoreChannelStream) FilterSlim(fn func(MultipleDataStoreChannel, int) bool) *MultipleDataStoreChannelStream {
	result := MultipleDataStoreChannelStreamOf()
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
func (self *MultipleDataStoreChannelStream) Find(fn func(MultipleDataStoreChannel, int) bool) *MultipleDataStoreChannel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultipleDataStoreChannelStream) FindOr(fn func(MultipleDataStoreChannel, int) bool, or MultipleDataStoreChannel) MultipleDataStoreChannel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultipleDataStoreChannelStream) FindIndex(fn func(MultipleDataStoreChannel, int) bool) int {
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
func (self *MultipleDataStoreChannelStream) First() *MultipleDataStoreChannel {
	return self.Get(0)
}
func (self *MultipleDataStoreChannelStream) FirstOr(arg MultipleDataStoreChannel) MultipleDataStoreChannel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDataStoreChannelStream) ForEach(fn func(MultipleDataStoreChannel, int)) *MultipleDataStoreChannelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultipleDataStoreChannelStream) ForEachRight(fn func(MultipleDataStoreChannel, int)) *MultipleDataStoreChannelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultipleDataStoreChannelStream) GroupBy(fn func(MultipleDataStoreChannel, int) string) map[string][]MultipleDataStoreChannel {
	m := map[string][]MultipleDataStoreChannel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultipleDataStoreChannelStream) GroupByValues(fn func(MultipleDataStoreChannel, int) string) [][]MultipleDataStoreChannel {
	var tmp [][]MultipleDataStoreChannel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultipleDataStoreChannelStream) IndexOf(arg MultipleDataStoreChannel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultipleDataStoreChannelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultipleDataStoreChannelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultipleDataStoreChannelStream) Last() *MultipleDataStoreChannel {
	return self.Get(self.Len() - 1)
}
func (self *MultipleDataStoreChannelStream) LastOr(arg MultipleDataStoreChannel) MultipleDataStoreChannel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDataStoreChannelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultipleDataStoreChannelStream) Limit(limit int) *MultipleDataStoreChannelStream {
	self.Slice(0, limit)
	return self
}

func (self *MultipleDataStoreChannelStream) Map(fn func(MultipleDataStoreChannel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreChannelStream) Map2Int(fn func(MultipleDataStoreChannel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreChannelStream) Map2Int32(fn func(MultipleDataStoreChannel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreChannelStream) Map2Int64(fn func(MultipleDataStoreChannel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreChannelStream) Map2Float32(fn func(MultipleDataStoreChannel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreChannelStream) Map2Float64(fn func(MultipleDataStoreChannel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreChannelStream) Map2Bool(fn func(MultipleDataStoreChannel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreChannelStream) Map2Bytes(fn func(MultipleDataStoreChannel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreChannelStream) Map2String(fn func(MultipleDataStoreChannel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreChannelStream) Max(fn func(MultipleDataStoreChannel, int) float64) *MultipleDataStoreChannel {
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
func (self *MultipleDataStoreChannelStream) Min(fn func(MultipleDataStoreChannel, int) float64) *MultipleDataStoreChannel {
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
func (self *MultipleDataStoreChannelStream) NoneMatch(fn func(MultipleDataStoreChannel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultipleDataStoreChannelStream) Get(index int) *MultipleDataStoreChannel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultipleDataStoreChannelStream) GetOr(index int, arg MultipleDataStoreChannel) MultipleDataStoreChannel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDataStoreChannelStream) Peek(fn func(*MultipleDataStoreChannel, int)) *MultipleDataStoreChannelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *MultipleDataStoreChannelStream) Reduce(fn func(MultipleDataStoreChannel, MultipleDataStoreChannel, int) MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	return self.ReduceInit(fn, MultipleDataStoreChannel{})
}
func (self *MultipleDataStoreChannelStream) ReduceInit(fn func(MultipleDataStoreChannel, MultipleDataStoreChannel, int) MultipleDataStoreChannel, initialValue MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	result := MultipleDataStoreChannelStreamOf()
	self.ForEach(func(v MultipleDataStoreChannel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultipleDataStoreChannelStream) ReduceInterface(fn func(interface{}, MultipleDataStoreChannel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultipleDataStoreChannel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultipleDataStoreChannelStream) ReduceString(fn func(string, MultipleDataStoreChannel, int) string) []string {
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
func (self *MultipleDataStoreChannelStream) ReduceInt(fn func(int, MultipleDataStoreChannel, int) int) []int {
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
func (self *MultipleDataStoreChannelStream) ReduceInt32(fn func(int32, MultipleDataStoreChannel, int) int32) []int32 {
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
func (self *MultipleDataStoreChannelStream) ReduceInt64(fn func(int64, MultipleDataStoreChannel, int) int64) []int64 {
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
func (self *MultipleDataStoreChannelStream) ReduceFloat32(fn func(float32, MultipleDataStoreChannel, int) float32) []float32 {
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
func (self *MultipleDataStoreChannelStream) ReduceFloat64(fn func(float64, MultipleDataStoreChannel, int) float64) []float64 {
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
func (self *MultipleDataStoreChannelStream) ReduceBool(fn func(bool, MultipleDataStoreChannel, int) bool) []bool {
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
func (self *MultipleDataStoreChannelStream) Reverse() *MultipleDataStoreChannelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultipleDataStoreChannelStream) Replace(fn func(MultipleDataStoreChannel, int) MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	return self.ForEach(func(v MultipleDataStoreChannel, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultipleDataStoreChannelStream) Select(fn func(MultipleDataStoreChannel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultipleDataStoreChannelStream) Set(index int, val MultipleDataStoreChannel) *MultipleDataStoreChannelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultipleDataStoreChannelStream) Skip(skip int) *MultipleDataStoreChannelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultipleDataStoreChannelStream) SkippingEach(fn func(MultipleDataStoreChannel, int) int) *MultipleDataStoreChannelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultipleDataStoreChannelStream) Slice(startIndex, n int) *MultipleDataStoreChannelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultipleDataStoreChannel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultipleDataStoreChannelStream) Sort(fn func(i, j int) bool) *MultipleDataStoreChannelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultipleDataStoreChannelStream) Tail() *MultipleDataStoreChannel {
	return self.Last()
}
func (self *MultipleDataStoreChannelStream) TailOr(arg MultipleDataStoreChannel) MultipleDataStoreChannel {
	return self.LastOr(arg)
}
func (self *MultipleDataStoreChannelStream) ToList() []MultipleDataStoreChannel {
	return self.Val()
}
func (self *MultipleDataStoreChannelStream) Unique() *MultipleDataStoreChannelStream {
	return self.Distinct()
}
func (self *MultipleDataStoreChannelStream) Val() []MultipleDataStoreChannel {
	if self == nil {
		return []MultipleDataStoreChannel{}
	}
	return *self.Copy()
}
func (self *MultipleDataStoreChannelStream) While(fn func(MultipleDataStoreChannel, int) bool) *MultipleDataStoreChannelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultipleDataStoreChannelStream) Where(fn func(MultipleDataStoreChannel) bool) *MultipleDataStoreChannelStream {
	result := MultipleDataStoreChannelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDataStoreChannelStream) WhereSlim(fn func(MultipleDataStoreChannel) bool) *MultipleDataStoreChannelStream {
	result := MultipleDataStoreChannelStreamOf()
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
