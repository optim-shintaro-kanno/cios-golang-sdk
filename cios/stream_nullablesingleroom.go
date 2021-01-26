/*
 * Collection utility of NullableSingleRoom Struct
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

type NullableSingleRoomStream []NullableSingleRoom

func NullableSingleRoomStreamOf(arg ...NullableSingleRoom) NullableSingleRoomStream {
	return arg
}
func NullableSingleRoomStreamFrom(arg []NullableSingleRoom) NullableSingleRoomStream {
	return arg
}
func CreateNullableSingleRoomStream(arg ...NullableSingleRoom) *NullableSingleRoomStream {
	tmp := NullableSingleRoomStreamOf(arg...)
	return &tmp
}
func GenerateNullableSingleRoomStream(arg []NullableSingleRoom) *NullableSingleRoomStream {
	tmp := NullableSingleRoomStreamFrom(arg)
	return &tmp
}

func (self *NullableSingleRoomStream) Add(arg NullableSingleRoom) *NullableSingleRoomStream {
	return self.AddAll(arg)
}
func (self *NullableSingleRoomStream) AddAll(arg ...NullableSingleRoom) *NullableSingleRoomStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSingleRoomStream) AddSafe(arg *NullableSingleRoom) *NullableSingleRoomStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSingleRoomStream) Aggregate(fn func(NullableSingleRoom, NullableSingleRoom) NullableSingleRoom) *NullableSingleRoomStream {
	result := NullableSingleRoomStreamOf()
	self.ForEach(func(v NullableSingleRoom, i int) {
		if i == 0 {
			result.Add(fn(NullableSingleRoom{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleRoomStream) AllMatch(fn func(NullableSingleRoom, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSingleRoomStream) AnyMatch(fn func(NullableSingleRoom, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSingleRoomStream) Clone() *NullableSingleRoomStream {
	temp := make([]NullableSingleRoom, self.Len())
	copy(temp, *self)
	return (*NullableSingleRoomStream)(&temp)
}
func (self *NullableSingleRoomStream) Copy() *NullableSingleRoomStream {
	return self.Clone()
}
func (self *NullableSingleRoomStream) Concat(arg []NullableSingleRoom) *NullableSingleRoomStream {
	return self.AddAll(arg...)
}
func (self *NullableSingleRoomStream) Contains(arg NullableSingleRoom) bool {
	return self.FindIndex(func(_arg NullableSingleRoom, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSingleRoomStream) Clean() *NullableSingleRoomStream {
	*self = NullableSingleRoomStreamOf()
	return self
}
func (self *NullableSingleRoomStream) Delete(index int) *NullableSingleRoomStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSingleRoomStream) DeleteRange(startIndex, endIndex int) *NullableSingleRoomStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSingleRoomStream) Distinct() *NullableSingleRoomStream {
	caches := map[string]bool{}
	result := NullableSingleRoomStreamOf()
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
func (self *NullableSingleRoomStream) Each(fn func(NullableSingleRoom)) *NullableSingleRoomStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSingleRoomStream) EachRight(fn func(NullableSingleRoom)) *NullableSingleRoomStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSingleRoomStream) Equals(arr []NullableSingleRoom) bool {
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
func (self *NullableSingleRoomStream) Filter(fn func(NullableSingleRoom, int) bool) *NullableSingleRoomStream {
	result := NullableSingleRoomStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleRoomStream) FilterSlim(fn func(NullableSingleRoom, int) bool) *NullableSingleRoomStream {
	result := NullableSingleRoomStreamOf()
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
func (self *NullableSingleRoomStream) Find(fn func(NullableSingleRoom, int) bool) *NullableSingleRoom {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSingleRoomStream) FindOr(fn func(NullableSingleRoom, int) bool, or NullableSingleRoom) NullableSingleRoom {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSingleRoomStream) FindIndex(fn func(NullableSingleRoom, int) bool) int {
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
func (self *NullableSingleRoomStream) First() *NullableSingleRoom {
	return self.Get(0)
}
func (self *NullableSingleRoomStream) FirstOr(arg NullableSingleRoom) NullableSingleRoom {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleRoomStream) ForEach(fn func(NullableSingleRoom, int)) *NullableSingleRoomStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSingleRoomStream) ForEachRight(fn func(NullableSingleRoom, int)) *NullableSingleRoomStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSingleRoomStream) GroupBy(fn func(NullableSingleRoom, int) string) map[string][]NullableSingleRoom {
	m := map[string][]NullableSingleRoom{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSingleRoomStream) GroupByValues(fn func(NullableSingleRoom, int) string) [][]NullableSingleRoom {
	var tmp [][]NullableSingleRoom
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSingleRoomStream) IndexOf(arg NullableSingleRoom) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSingleRoomStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSingleRoomStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSingleRoomStream) Last() *NullableSingleRoom {
	return self.Get(self.Len() - 1)
}
func (self *NullableSingleRoomStream) LastOr(arg NullableSingleRoom) NullableSingleRoom {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleRoomStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSingleRoomStream) Limit(limit int) *NullableSingleRoomStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSingleRoomStream) Map(fn func(NullableSingleRoom, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRoomStream) Map2Int(fn func(NullableSingleRoom, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRoomStream) Map2Int32(fn func(NullableSingleRoom, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRoomStream) Map2Int64(fn func(NullableSingleRoom, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRoomStream) Map2Float32(fn func(NullableSingleRoom, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRoomStream) Map2Float64(fn func(NullableSingleRoom, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRoomStream) Map2Bool(fn func(NullableSingleRoom, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRoomStream) Map2Bytes(fn func(NullableSingleRoom, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRoomStream) Map2String(fn func(NullableSingleRoom, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleRoomStream) Max(fn func(NullableSingleRoom, int) float64) *NullableSingleRoom {
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
func (self *NullableSingleRoomStream) Min(fn func(NullableSingleRoom, int) float64) *NullableSingleRoom {
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
func (self *NullableSingleRoomStream) NoneMatch(fn func(NullableSingleRoom, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSingleRoomStream) Get(index int) *NullableSingleRoom {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSingleRoomStream) GetOr(index int, arg NullableSingleRoom) NullableSingleRoom {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleRoomStream) Peek(fn func(*NullableSingleRoom, int)) *NullableSingleRoomStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableSingleRoomStream) Reduce(fn func(NullableSingleRoom, NullableSingleRoom, int) NullableSingleRoom) *NullableSingleRoomStream {
	return self.ReduceInit(fn, NullableSingleRoom{})
}
func (self *NullableSingleRoomStream) ReduceInit(fn func(NullableSingleRoom, NullableSingleRoom, int) NullableSingleRoom, initialValue NullableSingleRoom) *NullableSingleRoomStream {
	result := NullableSingleRoomStreamOf()
	self.ForEach(func(v NullableSingleRoom, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleRoomStream) ReduceInterface(fn func(interface{}, NullableSingleRoom, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSingleRoom{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSingleRoomStream) ReduceString(fn func(string, NullableSingleRoom, int) string) []string {
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
func (self *NullableSingleRoomStream) ReduceInt(fn func(int, NullableSingleRoom, int) int) []int {
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
func (self *NullableSingleRoomStream) ReduceInt32(fn func(int32, NullableSingleRoom, int) int32) []int32 {
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
func (self *NullableSingleRoomStream) ReduceInt64(fn func(int64, NullableSingleRoom, int) int64) []int64 {
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
func (self *NullableSingleRoomStream) ReduceFloat32(fn func(float32, NullableSingleRoom, int) float32) []float32 {
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
func (self *NullableSingleRoomStream) ReduceFloat64(fn func(float64, NullableSingleRoom, int) float64) []float64 {
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
func (self *NullableSingleRoomStream) ReduceBool(fn func(bool, NullableSingleRoom, int) bool) []bool {
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
func (self *NullableSingleRoomStream) Reverse() *NullableSingleRoomStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSingleRoomStream) Replace(fn func(NullableSingleRoom, int) NullableSingleRoom) *NullableSingleRoomStream {
	return self.ForEach(func(v NullableSingleRoom, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSingleRoomStream) Select(fn func(NullableSingleRoom) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSingleRoomStream) Set(index int, val NullableSingleRoom) *NullableSingleRoomStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSingleRoomStream) Skip(skip int) *NullableSingleRoomStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSingleRoomStream) SkippingEach(fn func(NullableSingleRoom, int) int) *NullableSingleRoomStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSingleRoomStream) Slice(startIndex, n int) *NullableSingleRoomStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSingleRoom{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSingleRoomStream) Sort(fn func(i, j int) bool) *NullableSingleRoomStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSingleRoomStream) Tail() *NullableSingleRoom {
	return self.Last()
}
func (self *NullableSingleRoomStream) TailOr(arg NullableSingleRoom) NullableSingleRoom {
	return self.LastOr(arg)
}
func (self *NullableSingleRoomStream) ToList() []NullableSingleRoom {
	return self.Val()
}
func (self *NullableSingleRoomStream) Unique() *NullableSingleRoomStream {
	return self.Distinct()
}
func (self *NullableSingleRoomStream) Val() []NullableSingleRoom {
	if self == nil {
		return []NullableSingleRoom{}
	}
	return *self.Copy()
}
func (self *NullableSingleRoomStream) While(fn func(NullableSingleRoom, int) bool) *NullableSingleRoomStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSingleRoomStream) Where(fn func(NullableSingleRoom) bool) *NullableSingleRoomStream {
	result := NullableSingleRoomStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleRoomStream) WhereSlim(fn func(NullableSingleRoom) bool) *NullableSingleRoomStream {
	result := NullableSingleRoomStreamOf()
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
