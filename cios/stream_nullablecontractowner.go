/*
 * Collection utility of NullableContractOwner Struct
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

type NullableContractOwnerStream []NullableContractOwner

func NullableContractOwnerStreamOf(arg ...NullableContractOwner) NullableContractOwnerStream {
	return arg
}
func NullableContractOwnerStreamFrom(arg []NullableContractOwner) NullableContractOwnerStream {
	return arg
}
func CreateNullableContractOwnerStream(arg ...NullableContractOwner) *NullableContractOwnerStream {
	tmp := NullableContractOwnerStreamOf(arg...)
	return &tmp
}
func GenerateNullableContractOwnerStream(arg []NullableContractOwner) *NullableContractOwnerStream {
	tmp := NullableContractOwnerStreamFrom(arg)
	return &tmp
}

func (self *NullableContractOwnerStream) Add(arg NullableContractOwner) *NullableContractOwnerStream {
	return self.AddAll(arg)
}
func (self *NullableContractOwnerStream) AddAll(arg ...NullableContractOwner) *NullableContractOwnerStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableContractOwnerStream) AddSafe(arg *NullableContractOwner) *NullableContractOwnerStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableContractOwnerStream) Aggregate(fn func(NullableContractOwner, NullableContractOwner) NullableContractOwner) *NullableContractOwnerStream {
	result := NullableContractOwnerStreamOf()
	self.ForEach(func(v NullableContractOwner, i int) {
		if i == 0 {
			result.Add(fn(NullableContractOwner{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableContractOwnerStream) AllMatch(fn func(NullableContractOwner, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableContractOwnerStream) AnyMatch(fn func(NullableContractOwner, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableContractOwnerStream) Clone() *NullableContractOwnerStream {
	temp := make([]NullableContractOwner, self.Len())
	copy(temp, *self)
	return (*NullableContractOwnerStream)(&temp)
}
func (self *NullableContractOwnerStream) Copy() *NullableContractOwnerStream {
	return self.Clone()
}
func (self *NullableContractOwnerStream) Concat(arg []NullableContractOwner) *NullableContractOwnerStream {
	return self.AddAll(arg...)
}
func (self *NullableContractOwnerStream) Contains(arg NullableContractOwner) bool {
	return self.FindIndex(func(_arg NullableContractOwner, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableContractOwnerStream) Clean() *NullableContractOwnerStream {
	*self = NullableContractOwnerStreamOf()
	return self
}
func (self *NullableContractOwnerStream) Delete(index int) *NullableContractOwnerStream {
	return self.DeleteRange(index, index)
}
func (self *NullableContractOwnerStream) DeleteRange(startIndex, endIndex int) *NullableContractOwnerStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableContractOwnerStream) Distinct() *NullableContractOwnerStream {
	caches := map[string]bool{}
	result := NullableContractOwnerStreamOf()
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
func (self *NullableContractOwnerStream) Each(fn func(NullableContractOwner)) *NullableContractOwnerStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableContractOwnerStream) EachRight(fn func(NullableContractOwner)) *NullableContractOwnerStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableContractOwnerStream) Equals(arr []NullableContractOwner) bool {
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
func (self *NullableContractOwnerStream) Filter(fn func(NullableContractOwner, int) bool) *NullableContractOwnerStream {
	result := NullableContractOwnerStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableContractOwnerStream) FilterSlim(fn func(NullableContractOwner, int) bool) *NullableContractOwnerStream {
	result := NullableContractOwnerStreamOf()
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
func (self *NullableContractOwnerStream) Find(fn func(NullableContractOwner, int) bool) *NullableContractOwner {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableContractOwnerStream) FindOr(fn func(NullableContractOwner, int) bool, or NullableContractOwner) NullableContractOwner {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableContractOwnerStream) FindIndex(fn func(NullableContractOwner, int) bool) int {
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
func (self *NullableContractOwnerStream) First() *NullableContractOwner {
	return self.Get(0)
}
func (self *NullableContractOwnerStream) FirstOr(arg NullableContractOwner) NullableContractOwner {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableContractOwnerStream) ForEach(fn func(NullableContractOwner, int)) *NullableContractOwnerStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableContractOwnerStream) ForEachRight(fn func(NullableContractOwner, int)) *NullableContractOwnerStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableContractOwnerStream) GroupBy(fn func(NullableContractOwner, int) string) map[string][]NullableContractOwner {
	m := map[string][]NullableContractOwner{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableContractOwnerStream) GroupByValues(fn func(NullableContractOwner, int) string) [][]NullableContractOwner {
	var tmp [][]NullableContractOwner
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableContractOwnerStream) IndexOf(arg NullableContractOwner) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableContractOwnerStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableContractOwnerStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableContractOwnerStream) Last() *NullableContractOwner {
	return self.Get(self.Len() - 1)
}
func (self *NullableContractOwnerStream) LastOr(arg NullableContractOwner) NullableContractOwner {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableContractOwnerStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableContractOwnerStream) Limit(limit int) *NullableContractOwnerStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableContractOwnerStream) Map(fn func(NullableContractOwner, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractOwnerStream) Map2Int(fn func(NullableContractOwner, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractOwnerStream) Map2Int32(fn func(NullableContractOwner, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractOwnerStream) Map2Int64(fn func(NullableContractOwner, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractOwnerStream) Map2Float32(fn func(NullableContractOwner, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractOwnerStream) Map2Float64(fn func(NullableContractOwner, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractOwnerStream) Map2Bool(fn func(NullableContractOwner, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractOwnerStream) Map2Bytes(fn func(NullableContractOwner, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractOwnerStream) Map2String(fn func(NullableContractOwner, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractOwnerStream) Max(fn func(NullableContractOwner, int) float64) *NullableContractOwner {
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
func (self *NullableContractOwnerStream) Min(fn func(NullableContractOwner, int) float64) *NullableContractOwner {
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
func (self *NullableContractOwnerStream) NoneMatch(fn func(NullableContractOwner, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableContractOwnerStream) Get(index int) *NullableContractOwner {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableContractOwnerStream) GetOr(index int, arg NullableContractOwner) NullableContractOwner {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableContractOwnerStream) Peek(fn func(*NullableContractOwner, int)) *NullableContractOwnerStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableContractOwnerStream) Reduce(fn func(NullableContractOwner, NullableContractOwner, int) NullableContractOwner) *NullableContractOwnerStream {
	return self.ReduceInit(fn, NullableContractOwner{})
}
func (self *NullableContractOwnerStream) ReduceInit(fn func(NullableContractOwner, NullableContractOwner, int) NullableContractOwner, initialValue NullableContractOwner) *NullableContractOwnerStream {
	result := NullableContractOwnerStreamOf()
	self.ForEach(func(v NullableContractOwner, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableContractOwnerStream) ReduceInterface(fn func(interface{}, NullableContractOwner, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableContractOwner{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableContractOwnerStream) ReduceString(fn func(string, NullableContractOwner, int) string) []string {
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
func (self *NullableContractOwnerStream) ReduceInt(fn func(int, NullableContractOwner, int) int) []int {
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
func (self *NullableContractOwnerStream) ReduceInt32(fn func(int32, NullableContractOwner, int) int32) []int32 {
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
func (self *NullableContractOwnerStream) ReduceInt64(fn func(int64, NullableContractOwner, int) int64) []int64 {
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
func (self *NullableContractOwnerStream) ReduceFloat32(fn func(float32, NullableContractOwner, int) float32) []float32 {
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
func (self *NullableContractOwnerStream) ReduceFloat64(fn func(float64, NullableContractOwner, int) float64) []float64 {
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
func (self *NullableContractOwnerStream) ReduceBool(fn func(bool, NullableContractOwner, int) bool) []bool {
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
func (self *NullableContractOwnerStream) Reverse() *NullableContractOwnerStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableContractOwnerStream) Replace(fn func(NullableContractOwner, int) NullableContractOwner) *NullableContractOwnerStream {
	return self.ForEach(func(v NullableContractOwner, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableContractOwnerStream) Select(fn func(NullableContractOwner) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableContractOwnerStream) Set(index int, val NullableContractOwner) *NullableContractOwnerStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableContractOwnerStream) Skip(skip int) *NullableContractOwnerStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableContractOwnerStream) SkippingEach(fn func(NullableContractOwner, int) int) *NullableContractOwnerStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableContractOwnerStream) Slice(startIndex, n int) *NullableContractOwnerStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableContractOwner{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableContractOwnerStream) Sort(fn func(i, j int) bool) *NullableContractOwnerStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableContractOwnerStream) Tail() *NullableContractOwner {
	return self.Last()
}
func (self *NullableContractOwnerStream) TailOr(arg NullableContractOwner) NullableContractOwner {
	return self.LastOr(arg)
}
func (self *NullableContractOwnerStream) ToList() []NullableContractOwner {
	return self.Val()
}
func (self *NullableContractOwnerStream) Unique() *NullableContractOwnerStream {
	return self.Distinct()
}
func (self *NullableContractOwnerStream) Val() []NullableContractOwner {
	if self == nil {
		return []NullableContractOwner{}
	}
	return *self.Copy()
}
func (self *NullableContractOwnerStream) While(fn func(NullableContractOwner, int) bool) *NullableContractOwnerStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableContractOwnerStream) Where(fn func(NullableContractOwner) bool) *NullableContractOwnerStream {
	result := NullableContractOwnerStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableContractOwnerStream) WhereSlim(fn func(NullableContractOwner) bool) *NullableContractOwnerStream {
	result := NullableContractOwnerStreamOf()
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
