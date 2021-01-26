/*
 * Collection utility of NullableMultipleResourceOwner Struct
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

type NullableMultipleResourceOwnerStream []NullableMultipleResourceOwner

func NullableMultipleResourceOwnerStreamOf(arg ...NullableMultipleResourceOwner) NullableMultipleResourceOwnerStream {
	return arg
}
func NullableMultipleResourceOwnerStreamFrom(arg []NullableMultipleResourceOwner) NullableMultipleResourceOwnerStream {
	return arg
}
func CreateNullableMultipleResourceOwnerStream(arg ...NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	tmp := NullableMultipleResourceOwnerStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleResourceOwnerStream(arg []NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	tmp := NullableMultipleResourceOwnerStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleResourceOwnerStream) Add(arg NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleResourceOwnerStream) AddAll(arg ...NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleResourceOwnerStream) AddSafe(arg *NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleResourceOwnerStream) Aggregate(fn func(NullableMultipleResourceOwner, NullableMultipleResourceOwner) NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	result := NullableMultipleResourceOwnerStreamOf()
	self.ForEach(func(v NullableMultipleResourceOwner, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleResourceOwner{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleResourceOwnerStream) AllMatch(fn func(NullableMultipleResourceOwner, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleResourceOwnerStream) AnyMatch(fn func(NullableMultipleResourceOwner, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleResourceOwnerStream) Clone() *NullableMultipleResourceOwnerStream {
	temp := make([]NullableMultipleResourceOwner, self.Len())
	copy(temp, *self)
	return (*NullableMultipleResourceOwnerStream)(&temp)
}
func (self *NullableMultipleResourceOwnerStream) Copy() *NullableMultipleResourceOwnerStream {
	return self.Clone()
}
func (self *NullableMultipleResourceOwnerStream) Concat(arg []NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleResourceOwnerStream) Contains(arg NullableMultipleResourceOwner) bool {
	return self.FindIndex(func(_arg NullableMultipleResourceOwner, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleResourceOwnerStream) Clean() *NullableMultipleResourceOwnerStream {
	*self = NullableMultipleResourceOwnerStreamOf()
	return self
}
func (self *NullableMultipleResourceOwnerStream) Delete(index int) *NullableMultipleResourceOwnerStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleResourceOwnerStream) DeleteRange(startIndex, endIndex int) *NullableMultipleResourceOwnerStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleResourceOwnerStream) Distinct() *NullableMultipleResourceOwnerStream {
	caches := map[string]bool{}
	result := NullableMultipleResourceOwnerStreamOf()
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
func (self *NullableMultipleResourceOwnerStream) Each(fn func(NullableMultipleResourceOwner)) *NullableMultipleResourceOwnerStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleResourceOwnerStream) EachRight(fn func(NullableMultipleResourceOwner)) *NullableMultipleResourceOwnerStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleResourceOwnerStream) Equals(arr []NullableMultipleResourceOwner) bool {
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
func (self *NullableMultipleResourceOwnerStream) Filter(fn func(NullableMultipleResourceOwner, int) bool) *NullableMultipleResourceOwnerStream {
	result := NullableMultipleResourceOwnerStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleResourceOwnerStream) FilterSlim(fn func(NullableMultipleResourceOwner, int) bool) *NullableMultipleResourceOwnerStream {
	result := NullableMultipleResourceOwnerStreamOf()
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
func (self *NullableMultipleResourceOwnerStream) Find(fn func(NullableMultipleResourceOwner, int) bool) *NullableMultipleResourceOwner {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleResourceOwnerStream) FindOr(fn func(NullableMultipleResourceOwner, int) bool, or NullableMultipleResourceOwner) NullableMultipleResourceOwner {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleResourceOwnerStream) FindIndex(fn func(NullableMultipleResourceOwner, int) bool) int {
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
func (self *NullableMultipleResourceOwnerStream) First() *NullableMultipleResourceOwner {
	return self.Get(0)
}
func (self *NullableMultipleResourceOwnerStream) FirstOr(arg NullableMultipleResourceOwner) NullableMultipleResourceOwner {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleResourceOwnerStream) ForEach(fn func(NullableMultipleResourceOwner, int)) *NullableMultipleResourceOwnerStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleResourceOwnerStream) ForEachRight(fn func(NullableMultipleResourceOwner, int)) *NullableMultipleResourceOwnerStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleResourceOwnerStream) GroupBy(fn func(NullableMultipleResourceOwner, int) string) map[string][]NullableMultipleResourceOwner {
	m := map[string][]NullableMultipleResourceOwner{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleResourceOwnerStream) GroupByValues(fn func(NullableMultipleResourceOwner, int) string) [][]NullableMultipleResourceOwner {
	var tmp [][]NullableMultipleResourceOwner
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleResourceOwnerStream) IndexOf(arg NullableMultipleResourceOwner) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleResourceOwnerStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleResourceOwnerStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleResourceOwnerStream) Last() *NullableMultipleResourceOwner {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleResourceOwnerStream) LastOr(arg NullableMultipleResourceOwner) NullableMultipleResourceOwner {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleResourceOwnerStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleResourceOwnerStream) Limit(limit int) *NullableMultipleResourceOwnerStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleResourceOwnerStream) Map(fn func(NullableMultipleResourceOwner, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleResourceOwnerStream) Map2Int(fn func(NullableMultipleResourceOwner, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleResourceOwnerStream) Map2Int32(fn func(NullableMultipleResourceOwner, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleResourceOwnerStream) Map2Int64(fn func(NullableMultipleResourceOwner, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleResourceOwnerStream) Map2Float32(fn func(NullableMultipleResourceOwner, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleResourceOwnerStream) Map2Float64(fn func(NullableMultipleResourceOwner, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleResourceOwnerStream) Map2Bool(fn func(NullableMultipleResourceOwner, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleResourceOwnerStream) Map2Bytes(fn func(NullableMultipleResourceOwner, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleResourceOwnerStream) Map2String(fn func(NullableMultipleResourceOwner, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleResourceOwnerStream) Max(fn func(NullableMultipleResourceOwner, int) float64) *NullableMultipleResourceOwner {
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
func (self *NullableMultipleResourceOwnerStream) Min(fn func(NullableMultipleResourceOwner, int) float64) *NullableMultipleResourceOwner {
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
func (self *NullableMultipleResourceOwnerStream) NoneMatch(fn func(NullableMultipleResourceOwner, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleResourceOwnerStream) Get(index int) *NullableMultipleResourceOwner {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleResourceOwnerStream) GetOr(index int, arg NullableMultipleResourceOwner) NullableMultipleResourceOwner {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleResourceOwnerStream) Peek(fn func(*NullableMultipleResourceOwner, int)) *NullableMultipleResourceOwnerStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableMultipleResourceOwnerStream) Reduce(fn func(NullableMultipleResourceOwner, NullableMultipleResourceOwner, int) NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	return self.ReduceInit(fn, NullableMultipleResourceOwner{})
}
func (self *NullableMultipleResourceOwnerStream) ReduceInit(fn func(NullableMultipleResourceOwner, NullableMultipleResourceOwner, int) NullableMultipleResourceOwner, initialValue NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	result := NullableMultipleResourceOwnerStreamOf()
	self.ForEach(func(v NullableMultipleResourceOwner, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleResourceOwnerStream) ReduceInterface(fn func(interface{}, NullableMultipleResourceOwner, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleResourceOwner{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleResourceOwnerStream) ReduceString(fn func(string, NullableMultipleResourceOwner, int) string) []string {
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
func (self *NullableMultipleResourceOwnerStream) ReduceInt(fn func(int, NullableMultipleResourceOwner, int) int) []int {
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
func (self *NullableMultipleResourceOwnerStream) ReduceInt32(fn func(int32, NullableMultipleResourceOwner, int) int32) []int32 {
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
func (self *NullableMultipleResourceOwnerStream) ReduceInt64(fn func(int64, NullableMultipleResourceOwner, int) int64) []int64 {
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
func (self *NullableMultipleResourceOwnerStream) ReduceFloat32(fn func(float32, NullableMultipleResourceOwner, int) float32) []float32 {
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
func (self *NullableMultipleResourceOwnerStream) ReduceFloat64(fn func(float64, NullableMultipleResourceOwner, int) float64) []float64 {
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
func (self *NullableMultipleResourceOwnerStream) ReduceBool(fn func(bool, NullableMultipleResourceOwner, int) bool) []bool {
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
func (self *NullableMultipleResourceOwnerStream) Reverse() *NullableMultipleResourceOwnerStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleResourceOwnerStream) Replace(fn func(NullableMultipleResourceOwner, int) NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	return self.ForEach(func(v NullableMultipleResourceOwner, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleResourceOwnerStream) Select(fn func(NullableMultipleResourceOwner) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleResourceOwnerStream) Set(index int, val NullableMultipleResourceOwner) *NullableMultipleResourceOwnerStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleResourceOwnerStream) Skip(skip int) *NullableMultipleResourceOwnerStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleResourceOwnerStream) SkippingEach(fn func(NullableMultipleResourceOwner, int) int) *NullableMultipleResourceOwnerStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleResourceOwnerStream) Slice(startIndex, n int) *NullableMultipleResourceOwnerStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleResourceOwner{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleResourceOwnerStream) Sort(fn func(i, j int) bool) *NullableMultipleResourceOwnerStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleResourceOwnerStream) Tail() *NullableMultipleResourceOwner {
	return self.Last()
}
func (self *NullableMultipleResourceOwnerStream) TailOr(arg NullableMultipleResourceOwner) NullableMultipleResourceOwner {
	return self.LastOr(arg)
}
func (self *NullableMultipleResourceOwnerStream) ToList() []NullableMultipleResourceOwner {
	return self.Val()
}
func (self *NullableMultipleResourceOwnerStream) Unique() *NullableMultipleResourceOwnerStream {
	return self.Distinct()
}
func (self *NullableMultipleResourceOwnerStream) Val() []NullableMultipleResourceOwner {
	if self == nil {
		return []NullableMultipleResourceOwner{}
	}
	return *self.Copy()
}
func (self *NullableMultipleResourceOwnerStream) While(fn func(NullableMultipleResourceOwner, int) bool) *NullableMultipleResourceOwnerStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleResourceOwnerStream) Where(fn func(NullableMultipleResourceOwner) bool) *NullableMultipleResourceOwnerStream {
	result := NullableMultipleResourceOwnerStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleResourceOwnerStream) WhereSlim(fn func(NullableMultipleResourceOwner) bool) *NullableMultipleResourceOwnerStream {
	result := NullableMultipleResourceOwnerStreamOf()
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
