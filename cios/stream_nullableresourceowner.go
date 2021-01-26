/*
 * Collection utility of NullableResourceOwner Struct
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

type NullableResourceOwnerStream []NullableResourceOwner

func NullableResourceOwnerStreamOf(arg ...NullableResourceOwner) NullableResourceOwnerStream {
	return arg
}
func NullableResourceOwnerStreamFrom(arg []NullableResourceOwner) NullableResourceOwnerStream {
	return arg
}
func CreateNullableResourceOwnerStream(arg ...NullableResourceOwner) *NullableResourceOwnerStream {
	tmp := NullableResourceOwnerStreamOf(arg...)
	return &tmp
}
func GenerateNullableResourceOwnerStream(arg []NullableResourceOwner) *NullableResourceOwnerStream {
	tmp := NullableResourceOwnerStreamFrom(arg)
	return &tmp
}

func (self *NullableResourceOwnerStream) Add(arg NullableResourceOwner) *NullableResourceOwnerStream {
	return self.AddAll(arg)
}
func (self *NullableResourceOwnerStream) AddAll(arg ...NullableResourceOwner) *NullableResourceOwnerStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableResourceOwnerStream) AddSafe(arg *NullableResourceOwner) *NullableResourceOwnerStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableResourceOwnerStream) Aggregate(fn func(NullableResourceOwner, NullableResourceOwner) NullableResourceOwner) *NullableResourceOwnerStream {
	result := NullableResourceOwnerStreamOf()
	self.ForEach(func(v NullableResourceOwner, i int) {
		if i == 0 {
			result.Add(fn(NullableResourceOwner{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableResourceOwnerStream) AllMatch(fn func(NullableResourceOwner, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableResourceOwnerStream) AnyMatch(fn func(NullableResourceOwner, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableResourceOwnerStream) Clone() *NullableResourceOwnerStream {
	temp := make([]NullableResourceOwner, self.Len())
	copy(temp, *self)
	return (*NullableResourceOwnerStream)(&temp)
}
func (self *NullableResourceOwnerStream) Copy() *NullableResourceOwnerStream {
	return self.Clone()
}
func (self *NullableResourceOwnerStream) Concat(arg []NullableResourceOwner) *NullableResourceOwnerStream {
	return self.AddAll(arg...)
}
func (self *NullableResourceOwnerStream) Contains(arg NullableResourceOwner) bool {
	return self.FindIndex(func(_arg NullableResourceOwner, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableResourceOwnerStream) Clean() *NullableResourceOwnerStream {
	*self = NullableResourceOwnerStreamOf()
	return self
}
func (self *NullableResourceOwnerStream) Delete(index int) *NullableResourceOwnerStream {
	return self.DeleteRange(index, index)
}
func (self *NullableResourceOwnerStream) DeleteRange(startIndex, endIndex int) *NullableResourceOwnerStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableResourceOwnerStream) Distinct() *NullableResourceOwnerStream {
	caches := map[string]bool{}
	result := NullableResourceOwnerStreamOf()
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
func (self *NullableResourceOwnerStream) Each(fn func(NullableResourceOwner)) *NullableResourceOwnerStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableResourceOwnerStream) EachRight(fn func(NullableResourceOwner)) *NullableResourceOwnerStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableResourceOwnerStream) Equals(arr []NullableResourceOwner) bool {
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
func (self *NullableResourceOwnerStream) Filter(fn func(NullableResourceOwner, int) bool) *NullableResourceOwnerStream {
	result := NullableResourceOwnerStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableResourceOwnerStream) FilterSlim(fn func(NullableResourceOwner, int) bool) *NullableResourceOwnerStream {
	result := NullableResourceOwnerStreamOf()
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
func (self *NullableResourceOwnerStream) Find(fn func(NullableResourceOwner, int) bool) *NullableResourceOwner {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableResourceOwnerStream) FindOr(fn func(NullableResourceOwner, int) bool, or NullableResourceOwner) NullableResourceOwner {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableResourceOwnerStream) FindIndex(fn func(NullableResourceOwner, int) bool) int {
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
func (self *NullableResourceOwnerStream) First() *NullableResourceOwner {
	return self.Get(0)
}
func (self *NullableResourceOwnerStream) FirstOr(arg NullableResourceOwner) NullableResourceOwner {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableResourceOwnerStream) ForEach(fn func(NullableResourceOwner, int)) *NullableResourceOwnerStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableResourceOwnerStream) ForEachRight(fn func(NullableResourceOwner, int)) *NullableResourceOwnerStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableResourceOwnerStream) GroupBy(fn func(NullableResourceOwner, int) string) map[string][]NullableResourceOwner {
	m := map[string][]NullableResourceOwner{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableResourceOwnerStream) GroupByValues(fn func(NullableResourceOwner, int) string) [][]NullableResourceOwner {
	var tmp [][]NullableResourceOwner
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableResourceOwnerStream) IndexOf(arg NullableResourceOwner) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableResourceOwnerStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableResourceOwnerStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableResourceOwnerStream) Last() *NullableResourceOwner {
	return self.Get(self.Len() - 1)
}
func (self *NullableResourceOwnerStream) LastOr(arg NullableResourceOwner) NullableResourceOwner {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableResourceOwnerStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableResourceOwnerStream) Limit(limit int) *NullableResourceOwnerStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableResourceOwnerStream) Map(fn func(NullableResourceOwner, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerStream) Map2Int(fn func(NullableResourceOwner, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerStream) Map2Int32(fn func(NullableResourceOwner, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerStream) Map2Int64(fn func(NullableResourceOwner, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerStream) Map2Float32(fn func(NullableResourceOwner, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerStream) Map2Float64(fn func(NullableResourceOwner, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerStream) Map2Bool(fn func(NullableResourceOwner, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerStream) Map2Bytes(fn func(NullableResourceOwner, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerStream) Map2String(fn func(NullableResourceOwner, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerStream) Max(fn func(NullableResourceOwner, int) float64) *NullableResourceOwner {
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
func (self *NullableResourceOwnerStream) Min(fn func(NullableResourceOwner, int) float64) *NullableResourceOwner {
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
func (self *NullableResourceOwnerStream) NoneMatch(fn func(NullableResourceOwner, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableResourceOwnerStream) Get(index int) *NullableResourceOwner {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableResourceOwnerStream) GetOr(index int, arg NullableResourceOwner) NullableResourceOwner {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableResourceOwnerStream) Peek(fn func(*NullableResourceOwner, int)) *NullableResourceOwnerStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableResourceOwnerStream) Reduce(fn func(NullableResourceOwner, NullableResourceOwner, int) NullableResourceOwner) *NullableResourceOwnerStream {
	return self.ReduceInit(fn, NullableResourceOwner{})
}
func (self *NullableResourceOwnerStream) ReduceInit(fn func(NullableResourceOwner, NullableResourceOwner, int) NullableResourceOwner, initialValue NullableResourceOwner) *NullableResourceOwnerStream {
	result := NullableResourceOwnerStreamOf()
	self.ForEach(func(v NullableResourceOwner, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableResourceOwnerStream) ReduceInterface(fn func(interface{}, NullableResourceOwner, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableResourceOwner{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableResourceOwnerStream) ReduceString(fn func(string, NullableResourceOwner, int) string) []string {
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
func (self *NullableResourceOwnerStream) ReduceInt(fn func(int, NullableResourceOwner, int) int) []int {
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
func (self *NullableResourceOwnerStream) ReduceInt32(fn func(int32, NullableResourceOwner, int) int32) []int32 {
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
func (self *NullableResourceOwnerStream) ReduceInt64(fn func(int64, NullableResourceOwner, int) int64) []int64 {
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
func (self *NullableResourceOwnerStream) ReduceFloat32(fn func(float32, NullableResourceOwner, int) float32) []float32 {
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
func (self *NullableResourceOwnerStream) ReduceFloat64(fn func(float64, NullableResourceOwner, int) float64) []float64 {
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
func (self *NullableResourceOwnerStream) ReduceBool(fn func(bool, NullableResourceOwner, int) bool) []bool {
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
func (self *NullableResourceOwnerStream) Reverse() *NullableResourceOwnerStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableResourceOwnerStream) Replace(fn func(NullableResourceOwner, int) NullableResourceOwner) *NullableResourceOwnerStream {
	return self.ForEach(func(v NullableResourceOwner, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableResourceOwnerStream) Select(fn func(NullableResourceOwner) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableResourceOwnerStream) Set(index int, val NullableResourceOwner) *NullableResourceOwnerStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableResourceOwnerStream) Skip(skip int) *NullableResourceOwnerStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableResourceOwnerStream) SkippingEach(fn func(NullableResourceOwner, int) int) *NullableResourceOwnerStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableResourceOwnerStream) Slice(startIndex, n int) *NullableResourceOwnerStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableResourceOwner{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableResourceOwnerStream) Sort(fn func(i, j int) bool) *NullableResourceOwnerStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableResourceOwnerStream) Tail() *NullableResourceOwner {
	return self.Last()
}
func (self *NullableResourceOwnerStream) TailOr(arg NullableResourceOwner) NullableResourceOwner {
	return self.LastOr(arg)
}
func (self *NullableResourceOwnerStream) ToList() []NullableResourceOwner {
	return self.Val()
}
func (self *NullableResourceOwnerStream) Unique() *NullableResourceOwnerStream {
	return self.Distinct()
}
func (self *NullableResourceOwnerStream) Val() []NullableResourceOwner {
	if self == nil {
		return []NullableResourceOwner{}
	}
	return *self.Copy()
}
func (self *NullableResourceOwnerStream) While(fn func(NullableResourceOwner, int) bool) *NullableResourceOwnerStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableResourceOwnerStream) Where(fn func(NullableResourceOwner) bool) *NullableResourceOwnerStream {
	result := NullableResourceOwnerStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableResourceOwnerStream) WhereSlim(fn func(NullableResourceOwner) bool) *NullableResourceOwnerStream {
	result := NullableResourceOwnerStreamOf()
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
