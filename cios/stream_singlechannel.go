/*
 * Collection utility of SingleChannel Struct
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

type SingleChannelStream []SingleChannel

func SingleChannelStreamOf(arg ...SingleChannel) SingleChannelStream {
	return arg
}
func SingleChannelStreamFrom(arg []SingleChannel) SingleChannelStream {
	return arg
}
func CreateSingleChannelStream(arg ...SingleChannel) *SingleChannelStream {
	tmp := SingleChannelStreamOf(arg...)
	return &tmp
}
func GenerateSingleChannelStream(arg []SingleChannel) *SingleChannelStream {
	tmp := SingleChannelStreamFrom(arg)
	return &tmp
}

func (self *SingleChannelStream) Add(arg SingleChannel) *SingleChannelStream {
	return self.AddAll(arg)
}
func (self *SingleChannelStream) AddAll(arg ...SingleChannel) *SingleChannelStream {
	*self = append(*self, arg...)
	return self
}
func (self *SingleChannelStream) AddSafe(arg *SingleChannel) *SingleChannelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *SingleChannelStream) Aggregate(fn func(SingleChannel, SingleChannel) SingleChannel) *SingleChannelStream {
	result := SingleChannelStreamOf()
	self.ForEach(func(v SingleChannel, i int) {
		if i == 0 {
			result.Add(fn(SingleChannel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *SingleChannelStream) AllMatch(fn func(SingleChannel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *SingleChannelStream) AnyMatch(fn func(SingleChannel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SingleChannelStream) Clone() *SingleChannelStream {
	temp := make([]SingleChannel, self.Len())
	copy(temp, *self)
	return (*SingleChannelStream)(&temp)
}
func (self *SingleChannelStream) Copy() *SingleChannelStream {
	return self.Clone()
}
func (self *SingleChannelStream) Concat(arg []SingleChannel) *SingleChannelStream {
	return self.AddAll(arg...)
}
func (self *SingleChannelStream) Contains(arg SingleChannel) bool {
	return self.FindIndex(func(_arg SingleChannel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *SingleChannelStream) Clean() *SingleChannelStream {
	*self = SingleChannelStreamOf()
	return self
}
func (self *SingleChannelStream) Delete(index int) *SingleChannelStream {
	return self.DeleteRange(index, index)
}
func (self *SingleChannelStream) DeleteRange(startIndex, endIndex int) *SingleChannelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *SingleChannelStream) Distinct() *SingleChannelStream {
	caches := map[string]bool{}
	result := SingleChannelStreamOf()
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
func (self *SingleChannelStream) Each(fn func(SingleChannel)) *SingleChannelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *SingleChannelStream) EachRight(fn func(SingleChannel)) *SingleChannelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *SingleChannelStream) Equals(arr []SingleChannel) bool {
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
func (self *SingleChannelStream) Filter(fn func(SingleChannel, int) bool) *SingleChannelStream {
	result := SingleChannelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleChannelStream) FilterSlim(fn func(SingleChannel, int) bool) *SingleChannelStream {
	result := SingleChannelStreamOf()
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
func (self *SingleChannelStream) Find(fn func(SingleChannel, int) bool) *SingleChannel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *SingleChannelStream) FindOr(fn func(SingleChannel, int) bool, or SingleChannel) SingleChannel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *SingleChannelStream) FindIndex(fn func(SingleChannel, int) bool) int {
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
func (self *SingleChannelStream) First() *SingleChannel {
	return self.Get(0)
}
func (self *SingleChannelStream) FirstOr(arg SingleChannel) SingleChannel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *SingleChannelStream) ForEach(fn func(SingleChannel, int)) *SingleChannelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *SingleChannelStream) ForEachRight(fn func(SingleChannel, int)) *SingleChannelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SingleChannelStream) GroupBy(fn func(SingleChannel, int) string) map[string][]SingleChannel {
	m := map[string][]SingleChannel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *SingleChannelStream) GroupByValues(fn func(SingleChannel, int) string) [][]SingleChannel {
	var tmp [][]SingleChannel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SingleChannelStream) IndexOf(arg SingleChannel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SingleChannelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SingleChannelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SingleChannelStream) Last() *SingleChannel {
	return self.Get(self.Len() - 1)
}
func (self *SingleChannelStream) LastOr(arg SingleChannel) SingleChannel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *SingleChannelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SingleChannelStream) Limit(limit int) *SingleChannelStream {
	self.Slice(0, limit)
	return self
}

func (self *SingleChannelStream) Map(fn func(SingleChannel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleChannelStream) Map2Int(fn func(SingleChannel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleChannelStream) Map2Int32(fn func(SingleChannel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleChannelStream) Map2Int64(fn func(SingleChannel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleChannelStream) Map2Float32(fn func(SingleChannel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleChannelStream) Map2Float64(fn func(SingleChannel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleChannelStream) Map2Bool(fn func(SingleChannel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleChannelStream) Map2Bytes(fn func(SingleChannel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleChannelStream) Map2String(fn func(SingleChannel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleChannelStream) Max(fn func(SingleChannel, int) float64) *SingleChannel {
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
func (self *SingleChannelStream) Min(fn func(SingleChannel, int) float64) *SingleChannel {
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
func (self *SingleChannelStream) NoneMatch(fn func(SingleChannel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *SingleChannelStream) Get(index int) *SingleChannel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *SingleChannelStream) GetOr(index int, arg SingleChannel) SingleChannel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *SingleChannelStream) Peek(fn func(*SingleChannel, int)) *SingleChannelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *SingleChannelStream) Reduce(fn func(SingleChannel, SingleChannel, int) SingleChannel) *SingleChannelStream {
	return self.ReduceInit(fn, SingleChannel{})
}
func (self *SingleChannelStream) ReduceInit(fn func(SingleChannel, SingleChannel, int) SingleChannel, initialValue SingleChannel) *SingleChannelStream {
	result := SingleChannelStreamOf()
	self.ForEach(func(v SingleChannel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *SingleChannelStream) ReduceInterface(fn func(interface{}, SingleChannel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(SingleChannel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SingleChannelStream) ReduceString(fn func(string, SingleChannel, int) string) []string {
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
func (self *SingleChannelStream) ReduceInt(fn func(int, SingleChannel, int) int) []int {
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
func (self *SingleChannelStream) ReduceInt32(fn func(int32, SingleChannel, int) int32) []int32 {
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
func (self *SingleChannelStream) ReduceInt64(fn func(int64, SingleChannel, int) int64) []int64 {
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
func (self *SingleChannelStream) ReduceFloat32(fn func(float32, SingleChannel, int) float32) []float32 {
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
func (self *SingleChannelStream) ReduceFloat64(fn func(float64, SingleChannel, int) float64) []float64 {
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
func (self *SingleChannelStream) ReduceBool(fn func(bool, SingleChannel, int) bool) []bool {
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
func (self *SingleChannelStream) Reverse() *SingleChannelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *SingleChannelStream) Replace(fn func(SingleChannel, int) SingleChannel) *SingleChannelStream {
	return self.ForEach(func(v SingleChannel, i int) { self.Set(i, fn(v, i)) })
}
func (self *SingleChannelStream) Select(fn func(SingleChannel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *SingleChannelStream) Set(index int, val SingleChannel) *SingleChannelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *SingleChannelStream) Skip(skip int) *SingleChannelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *SingleChannelStream) SkippingEach(fn func(SingleChannel, int) int) *SingleChannelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *SingleChannelStream) Slice(startIndex, n int) *SingleChannelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []SingleChannel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *SingleChannelStream) Sort(fn func(i, j int) bool) *SingleChannelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SingleChannelStream) Tail() *SingleChannel {
	return self.Last()
}
func (self *SingleChannelStream) TailOr(arg SingleChannel) SingleChannel {
	return self.LastOr(arg)
}
func (self *SingleChannelStream) ToList() []SingleChannel {
	return self.Val()
}
func (self *SingleChannelStream) Unique() *SingleChannelStream {
	return self.Distinct()
}
func (self *SingleChannelStream) Val() []SingleChannel {
	if self == nil {
		return []SingleChannel{}
	}
	return *self.Copy()
}
func (self *SingleChannelStream) While(fn func(SingleChannel, int) bool) *SingleChannelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *SingleChannelStream) Where(fn func(SingleChannel) bool) *SingleChannelStream {
	result := SingleChannelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleChannelStream) WhereSlim(fn func(SingleChannel) bool) *SingleChannelStream {
	result := SingleChannelStreamOf()
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
