/*
 * Collection utility of NullableChannelProposal Struct
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

type NullableChannelProposalStream []NullableChannelProposal

func NullableChannelProposalStreamOf(arg ...NullableChannelProposal) NullableChannelProposalStream {
	return arg
}
func NullableChannelProposalStreamFrom(arg []NullableChannelProposal) NullableChannelProposalStream {
	return arg
}
func CreateNullableChannelProposalStream(arg ...NullableChannelProposal) *NullableChannelProposalStream {
	tmp := NullableChannelProposalStreamOf(arg...)
	return &tmp
}
func GenerateNullableChannelProposalStream(arg []NullableChannelProposal) *NullableChannelProposalStream {
	tmp := NullableChannelProposalStreamFrom(arg)
	return &tmp
}

func (self *NullableChannelProposalStream) Add(arg NullableChannelProposal) *NullableChannelProposalStream {
	return self.AddAll(arg)
}
func (self *NullableChannelProposalStream) AddAll(arg ...NullableChannelProposal) *NullableChannelProposalStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableChannelProposalStream) AddSafe(arg *NullableChannelProposal) *NullableChannelProposalStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableChannelProposalStream) Aggregate(fn func(NullableChannelProposal, NullableChannelProposal) NullableChannelProposal) *NullableChannelProposalStream {
	result := NullableChannelProposalStreamOf()
	self.ForEach(func(v NullableChannelProposal, i int) {
		if i == 0 {
			result.Add(fn(NullableChannelProposal{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableChannelProposalStream) AllMatch(fn func(NullableChannelProposal, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableChannelProposalStream) AnyMatch(fn func(NullableChannelProposal, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableChannelProposalStream) Clone() *NullableChannelProposalStream {
	temp := make([]NullableChannelProposal, self.Len())
	copy(temp, *self)
	return (*NullableChannelProposalStream)(&temp)
}
func (self *NullableChannelProposalStream) Copy() *NullableChannelProposalStream {
	return self.Clone()
}
func (self *NullableChannelProposalStream) Concat(arg []NullableChannelProposal) *NullableChannelProposalStream {
	return self.AddAll(arg...)
}
func (self *NullableChannelProposalStream) Contains(arg NullableChannelProposal) bool {
	return self.FindIndex(func(_arg NullableChannelProposal, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableChannelProposalStream) Clean() *NullableChannelProposalStream {
	*self = NullableChannelProposalStreamOf()
	return self
}
func (self *NullableChannelProposalStream) Delete(index int) *NullableChannelProposalStream {
	return self.DeleteRange(index, index)
}
func (self *NullableChannelProposalStream) DeleteRange(startIndex, endIndex int) *NullableChannelProposalStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableChannelProposalStream) Distinct() *NullableChannelProposalStream {
	caches := map[string]bool{}
	result := NullableChannelProposalStreamOf()
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
func (self *NullableChannelProposalStream) Each(fn func(NullableChannelProposal)) *NullableChannelProposalStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableChannelProposalStream) EachRight(fn func(NullableChannelProposal)) *NullableChannelProposalStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableChannelProposalStream) Equals(arr []NullableChannelProposal) bool {
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
func (self *NullableChannelProposalStream) Filter(fn func(NullableChannelProposal, int) bool) *NullableChannelProposalStream {
	result := NullableChannelProposalStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableChannelProposalStream) FilterSlim(fn func(NullableChannelProposal, int) bool) *NullableChannelProposalStream {
	result := NullableChannelProposalStreamOf()
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
func (self *NullableChannelProposalStream) Find(fn func(NullableChannelProposal, int) bool) *NullableChannelProposal {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableChannelProposalStream) FindOr(fn func(NullableChannelProposal, int) bool, or NullableChannelProposal) NullableChannelProposal {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableChannelProposalStream) FindIndex(fn func(NullableChannelProposal, int) bool) int {
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
func (self *NullableChannelProposalStream) First() *NullableChannelProposal {
	return self.Get(0)
}
func (self *NullableChannelProposalStream) FirstOr(arg NullableChannelProposal) NullableChannelProposal {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableChannelProposalStream) ForEach(fn func(NullableChannelProposal, int)) *NullableChannelProposalStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableChannelProposalStream) ForEachRight(fn func(NullableChannelProposal, int)) *NullableChannelProposalStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableChannelProposalStream) GroupBy(fn func(NullableChannelProposal, int) string) map[string][]NullableChannelProposal {
	m := map[string][]NullableChannelProposal{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableChannelProposalStream) GroupByValues(fn func(NullableChannelProposal, int) string) [][]NullableChannelProposal {
	var tmp [][]NullableChannelProposal
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableChannelProposalStream) IndexOf(arg NullableChannelProposal) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableChannelProposalStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableChannelProposalStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableChannelProposalStream) Last() *NullableChannelProposal {
	return self.Get(self.Len() - 1)
}
func (self *NullableChannelProposalStream) LastOr(arg NullableChannelProposal) NullableChannelProposal {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableChannelProposalStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableChannelProposalStream) Limit(limit int) *NullableChannelProposalStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableChannelProposalStream) Map(fn func(NullableChannelProposal, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProposalStream) Map2Int(fn func(NullableChannelProposal, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProposalStream) Map2Int32(fn func(NullableChannelProposal, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProposalStream) Map2Int64(fn func(NullableChannelProposal, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProposalStream) Map2Float32(fn func(NullableChannelProposal, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProposalStream) Map2Float64(fn func(NullableChannelProposal, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProposalStream) Map2Bool(fn func(NullableChannelProposal, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProposalStream) Map2Bytes(fn func(NullableChannelProposal, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProposalStream) Map2String(fn func(NullableChannelProposal, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableChannelProposalStream) Max(fn func(NullableChannelProposal, int) float64) *NullableChannelProposal {
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
func (self *NullableChannelProposalStream) Min(fn func(NullableChannelProposal, int) float64) *NullableChannelProposal {
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
func (self *NullableChannelProposalStream) NoneMatch(fn func(NullableChannelProposal, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableChannelProposalStream) Get(index int) *NullableChannelProposal {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableChannelProposalStream) GetOr(index int, arg NullableChannelProposal) NullableChannelProposal {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableChannelProposalStream) Peek(fn func(*NullableChannelProposal, int)) *NullableChannelProposalStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableChannelProposalStream) Reduce(fn func(NullableChannelProposal, NullableChannelProposal, int) NullableChannelProposal) *NullableChannelProposalStream {
	return self.ReduceInit(fn, NullableChannelProposal{})
}
func (self *NullableChannelProposalStream) ReduceInit(fn func(NullableChannelProposal, NullableChannelProposal, int) NullableChannelProposal, initialValue NullableChannelProposal) *NullableChannelProposalStream {
	result := NullableChannelProposalStreamOf()
	self.ForEach(func(v NullableChannelProposal, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableChannelProposalStream) ReduceInterface(fn func(interface{}, NullableChannelProposal, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableChannelProposal{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableChannelProposalStream) ReduceString(fn func(string, NullableChannelProposal, int) string) []string {
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
func (self *NullableChannelProposalStream) ReduceInt(fn func(int, NullableChannelProposal, int) int) []int {
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
func (self *NullableChannelProposalStream) ReduceInt32(fn func(int32, NullableChannelProposal, int) int32) []int32 {
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
func (self *NullableChannelProposalStream) ReduceInt64(fn func(int64, NullableChannelProposal, int) int64) []int64 {
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
func (self *NullableChannelProposalStream) ReduceFloat32(fn func(float32, NullableChannelProposal, int) float32) []float32 {
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
func (self *NullableChannelProposalStream) ReduceFloat64(fn func(float64, NullableChannelProposal, int) float64) []float64 {
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
func (self *NullableChannelProposalStream) ReduceBool(fn func(bool, NullableChannelProposal, int) bool) []bool {
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
func (self *NullableChannelProposalStream) Reverse() *NullableChannelProposalStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableChannelProposalStream) Replace(fn func(NullableChannelProposal, int) NullableChannelProposal) *NullableChannelProposalStream {
	return self.ForEach(func(v NullableChannelProposal, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableChannelProposalStream) Select(fn func(NullableChannelProposal) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableChannelProposalStream) Set(index int, val NullableChannelProposal) *NullableChannelProposalStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableChannelProposalStream) Skip(skip int) *NullableChannelProposalStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableChannelProposalStream) SkippingEach(fn func(NullableChannelProposal, int) int) *NullableChannelProposalStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableChannelProposalStream) Slice(startIndex, n int) *NullableChannelProposalStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableChannelProposal{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableChannelProposalStream) Sort(fn func(i, j int) bool) *NullableChannelProposalStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableChannelProposalStream) Tail() *NullableChannelProposal {
	return self.Last()
}
func (self *NullableChannelProposalStream) TailOr(arg NullableChannelProposal) NullableChannelProposal {
	return self.LastOr(arg)
}
func (self *NullableChannelProposalStream) ToList() []NullableChannelProposal {
	return self.Val()
}
func (self *NullableChannelProposalStream) Unique() *NullableChannelProposalStream {
	return self.Distinct()
}
func (self *NullableChannelProposalStream) Val() []NullableChannelProposal {
	if self == nil {
		return []NullableChannelProposal{}
	}
	return *self.Copy()
}
func (self *NullableChannelProposalStream) While(fn func(NullableChannelProposal, int) bool) *NullableChannelProposalStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableChannelProposalStream) Where(fn func(NullableChannelProposal) bool) *NullableChannelProposalStream {
	result := NullableChannelProposalStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableChannelProposalStream) WhereSlim(fn func(NullableChannelProposal) bool) *NullableChannelProposalStream {
	result := NullableChannelProposalStreamOf()
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
