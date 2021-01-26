/*
 * Collection utility of SingleDataStoreChannel Struct
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

type SingleDataStoreChannelStream []SingleDataStoreChannel

func SingleDataStoreChannelStreamOf(arg ...SingleDataStoreChannel) SingleDataStoreChannelStream {
	return arg
}
func SingleDataStoreChannelStreamFrom(arg []SingleDataStoreChannel) SingleDataStoreChannelStream {
	return arg
}
func CreateSingleDataStoreChannelStream(arg ...SingleDataStoreChannel) *SingleDataStoreChannelStream {
	tmp := SingleDataStoreChannelStreamOf(arg...)
	return &tmp
}
func GenerateSingleDataStoreChannelStream(arg []SingleDataStoreChannel) *SingleDataStoreChannelStream {
	tmp := SingleDataStoreChannelStreamFrom(arg)
	return &tmp
}

func (self *SingleDataStoreChannelStream) Add(arg SingleDataStoreChannel) *SingleDataStoreChannelStream {
	return self.AddAll(arg)
}
func (self *SingleDataStoreChannelStream) AddAll(arg ...SingleDataStoreChannel) *SingleDataStoreChannelStream {
	*self = append(*self, arg...)
	return self
}
func (self *SingleDataStoreChannelStream) AddSafe(arg *SingleDataStoreChannel) *SingleDataStoreChannelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *SingleDataStoreChannelStream) Aggregate(fn func(SingleDataStoreChannel, SingleDataStoreChannel) SingleDataStoreChannel) *SingleDataStoreChannelStream {
	result := SingleDataStoreChannelStreamOf()
	self.ForEach(func(v SingleDataStoreChannel, i int) {
		if i == 0 {
			result.Add(fn(SingleDataStoreChannel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *SingleDataStoreChannelStream) AllMatch(fn func(SingleDataStoreChannel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *SingleDataStoreChannelStream) AnyMatch(fn func(SingleDataStoreChannel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SingleDataStoreChannelStream) Clone() *SingleDataStoreChannelStream {
	temp := make([]SingleDataStoreChannel, self.Len())
	copy(temp, *self)
	return (*SingleDataStoreChannelStream)(&temp)
}
func (self *SingleDataStoreChannelStream) Copy() *SingleDataStoreChannelStream {
	return self.Clone()
}
func (self *SingleDataStoreChannelStream) Concat(arg []SingleDataStoreChannel) *SingleDataStoreChannelStream {
	return self.AddAll(arg...)
}
func (self *SingleDataStoreChannelStream) Contains(arg SingleDataStoreChannel) bool {
	return self.FindIndex(func(_arg SingleDataStoreChannel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *SingleDataStoreChannelStream) Clean() *SingleDataStoreChannelStream {
	*self = SingleDataStoreChannelStreamOf()
	return self
}
func (self *SingleDataStoreChannelStream) Delete(index int) *SingleDataStoreChannelStream {
	return self.DeleteRange(index, index)
}
func (self *SingleDataStoreChannelStream) DeleteRange(startIndex, endIndex int) *SingleDataStoreChannelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *SingleDataStoreChannelStream) Distinct() *SingleDataStoreChannelStream {
	caches := map[string]bool{}
	result := SingleDataStoreChannelStreamOf()
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
func (self *SingleDataStoreChannelStream) Each(fn func(SingleDataStoreChannel)) *SingleDataStoreChannelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *SingleDataStoreChannelStream) EachRight(fn func(SingleDataStoreChannel)) *SingleDataStoreChannelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *SingleDataStoreChannelStream) Equals(arr []SingleDataStoreChannel) bool {
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
func (self *SingleDataStoreChannelStream) Filter(fn func(SingleDataStoreChannel, int) bool) *SingleDataStoreChannelStream {
	result := SingleDataStoreChannelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleDataStoreChannelStream) FilterSlim(fn func(SingleDataStoreChannel, int) bool) *SingleDataStoreChannelStream {
	result := SingleDataStoreChannelStreamOf()
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
func (self *SingleDataStoreChannelStream) Find(fn func(SingleDataStoreChannel, int) bool) *SingleDataStoreChannel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *SingleDataStoreChannelStream) FindOr(fn func(SingleDataStoreChannel, int) bool, or SingleDataStoreChannel) SingleDataStoreChannel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *SingleDataStoreChannelStream) FindIndex(fn func(SingleDataStoreChannel, int) bool) int {
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
func (self *SingleDataStoreChannelStream) First() *SingleDataStoreChannel {
	return self.Get(0)
}
func (self *SingleDataStoreChannelStream) FirstOr(arg SingleDataStoreChannel) SingleDataStoreChannel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *SingleDataStoreChannelStream) ForEach(fn func(SingleDataStoreChannel, int)) *SingleDataStoreChannelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *SingleDataStoreChannelStream) ForEachRight(fn func(SingleDataStoreChannel, int)) *SingleDataStoreChannelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SingleDataStoreChannelStream) GroupBy(fn func(SingleDataStoreChannel, int) string) map[string][]SingleDataStoreChannel {
	m := map[string][]SingleDataStoreChannel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *SingleDataStoreChannelStream) GroupByValues(fn func(SingleDataStoreChannel, int) string) [][]SingleDataStoreChannel {
	var tmp [][]SingleDataStoreChannel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SingleDataStoreChannelStream) IndexOf(arg SingleDataStoreChannel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SingleDataStoreChannelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SingleDataStoreChannelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SingleDataStoreChannelStream) Last() *SingleDataStoreChannel {
	return self.Get(self.Len() - 1)
}
func (self *SingleDataStoreChannelStream) LastOr(arg SingleDataStoreChannel) SingleDataStoreChannel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *SingleDataStoreChannelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SingleDataStoreChannelStream) Limit(limit int) *SingleDataStoreChannelStream {
	self.Slice(0, limit)
	return self
}

func (self *SingleDataStoreChannelStream) Map(fn func(SingleDataStoreChannel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDataStoreChannelStream) Map2Int(fn func(SingleDataStoreChannel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDataStoreChannelStream) Map2Int32(fn func(SingleDataStoreChannel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDataStoreChannelStream) Map2Int64(fn func(SingleDataStoreChannel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDataStoreChannelStream) Map2Float32(fn func(SingleDataStoreChannel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDataStoreChannelStream) Map2Float64(fn func(SingleDataStoreChannel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDataStoreChannelStream) Map2Bool(fn func(SingleDataStoreChannel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDataStoreChannelStream) Map2Bytes(fn func(SingleDataStoreChannel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDataStoreChannelStream) Map2String(fn func(SingleDataStoreChannel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDataStoreChannelStream) Max(fn func(SingleDataStoreChannel, int) float64) *SingleDataStoreChannel {
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
func (self *SingleDataStoreChannelStream) Min(fn func(SingleDataStoreChannel, int) float64) *SingleDataStoreChannel {
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
func (self *SingleDataStoreChannelStream) NoneMatch(fn func(SingleDataStoreChannel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *SingleDataStoreChannelStream) Get(index int) *SingleDataStoreChannel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *SingleDataStoreChannelStream) GetOr(index int, arg SingleDataStoreChannel) SingleDataStoreChannel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *SingleDataStoreChannelStream) Peek(fn func(*SingleDataStoreChannel, int)) *SingleDataStoreChannelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *SingleDataStoreChannelStream) Reduce(fn func(SingleDataStoreChannel, SingleDataStoreChannel, int) SingleDataStoreChannel) *SingleDataStoreChannelStream {
	return self.ReduceInit(fn, SingleDataStoreChannel{})
}
func (self *SingleDataStoreChannelStream) ReduceInit(fn func(SingleDataStoreChannel, SingleDataStoreChannel, int) SingleDataStoreChannel, initialValue SingleDataStoreChannel) *SingleDataStoreChannelStream {
	result := SingleDataStoreChannelStreamOf()
	self.ForEach(func(v SingleDataStoreChannel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *SingleDataStoreChannelStream) ReduceInterface(fn func(interface{}, SingleDataStoreChannel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(SingleDataStoreChannel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SingleDataStoreChannelStream) ReduceString(fn func(string, SingleDataStoreChannel, int) string) []string {
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
func (self *SingleDataStoreChannelStream) ReduceInt(fn func(int, SingleDataStoreChannel, int) int) []int {
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
func (self *SingleDataStoreChannelStream) ReduceInt32(fn func(int32, SingleDataStoreChannel, int) int32) []int32 {
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
func (self *SingleDataStoreChannelStream) ReduceInt64(fn func(int64, SingleDataStoreChannel, int) int64) []int64 {
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
func (self *SingleDataStoreChannelStream) ReduceFloat32(fn func(float32, SingleDataStoreChannel, int) float32) []float32 {
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
func (self *SingleDataStoreChannelStream) ReduceFloat64(fn func(float64, SingleDataStoreChannel, int) float64) []float64 {
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
func (self *SingleDataStoreChannelStream) ReduceBool(fn func(bool, SingleDataStoreChannel, int) bool) []bool {
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
func (self *SingleDataStoreChannelStream) Reverse() *SingleDataStoreChannelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *SingleDataStoreChannelStream) Replace(fn func(SingleDataStoreChannel, int) SingleDataStoreChannel) *SingleDataStoreChannelStream {
	return self.ForEach(func(v SingleDataStoreChannel, i int) { self.Set(i, fn(v, i)) })
}
func (self *SingleDataStoreChannelStream) Select(fn func(SingleDataStoreChannel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *SingleDataStoreChannelStream) Set(index int, val SingleDataStoreChannel) *SingleDataStoreChannelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *SingleDataStoreChannelStream) Skip(skip int) *SingleDataStoreChannelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *SingleDataStoreChannelStream) SkippingEach(fn func(SingleDataStoreChannel, int) int) *SingleDataStoreChannelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *SingleDataStoreChannelStream) Slice(startIndex, n int) *SingleDataStoreChannelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []SingleDataStoreChannel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *SingleDataStoreChannelStream) Sort(fn func(i, j int) bool) *SingleDataStoreChannelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SingleDataStoreChannelStream) Tail() *SingleDataStoreChannel {
	return self.Last()
}
func (self *SingleDataStoreChannelStream) TailOr(arg SingleDataStoreChannel) SingleDataStoreChannel {
	return self.LastOr(arg)
}
func (self *SingleDataStoreChannelStream) ToList() []SingleDataStoreChannel {
	return self.Val()
}
func (self *SingleDataStoreChannelStream) Unique() *SingleDataStoreChannelStream {
	return self.Distinct()
}
func (self *SingleDataStoreChannelStream) Val() []SingleDataStoreChannel {
	if self == nil {
		return []SingleDataStoreChannel{}
	}
	return *self.Copy()
}
func (self *SingleDataStoreChannelStream) While(fn func(SingleDataStoreChannel, int) bool) *SingleDataStoreChannelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *SingleDataStoreChannelStream) Where(fn func(SingleDataStoreChannel) bool) *SingleDataStoreChannelStream {
	result := SingleDataStoreChannelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleDataStoreChannelStream) WhereSlim(fn func(SingleDataStoreChannel) bool) *SingleDataStoreChannelStream {
	result := SingleDataStoreChannelStreamOf()
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
