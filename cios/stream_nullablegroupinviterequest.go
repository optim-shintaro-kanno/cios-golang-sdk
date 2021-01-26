/*
 * Collection utility of NullableGroupInviteRequest Struct
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

type NullableGroupInviteRequestStream []NullableGroupInviteRequest

func NullableGroupInviteRequestStreamOf(arg ...NullableGroupInviteRequest) NullableGroupInviteRequestStream {
	return arg
}
func NullableGroupInviteRequestStreamFrom(arg []NullableGroupInviteRequest) NullableGroupInviteRequestStream {
	return arg
}
func CreateNullableGroupInviteRequestStream(arg ...NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	tmp := NullableGroupInviteRequestStreamOf(arg...)
	return &tmp
}
func GenerateNullableGroupInviteRequestStream(arg []NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	tmp := NullableGroupInviteRequestStreamFrom(arg)
	return &tmp
}

func (self *NullableGroupInviteRequestStream) Add(arg NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	return self.AddAll(arg)
}
func (self *NullableGroupInviteRequestStream) AddAll(arg ...NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableGroupInviteRequestStream) AddSafe(arg *NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableGroupInviteRequestStream) Aggregate(fn func(NullableGroupInviteRequest, NullableGroupInviteRequest) NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	result := NullableGroupInviteRequestStreamOf()
	self.ForEach(func(v NullableGroupInviteRequest, i int) {
		if i == 0 {
			result.Add(fn(NullableGroupInviteRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableGroupInviteRequestStream) AllMatch(fn func(NullableGroupInviteRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableGroupInviteRequestStream) AnyMatch(fn func(NullableGroupInviteRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableGroupInviteRequestStream) Clone() *NullableGroupInviteRequestStream {
	temp := make([]NullableGroupInviteRequest, self.Len())
	copy(temp, *self)
	return (*NullableGroupInviteRequestStream)(&temp)
}
func (self *NullableGroupInviteRequestStream) Copy() *NullableGroupInviteRequestStream {
	return self.Clone()
}
func (self *NullableGroupInviteRequestStream) Concat(arg []NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	return self.AddAll(arg...)
}
func (self *NullableGroupInviteRequestStream) Contains(arg NullableGroupInviteRequest) bool {
	return self.FindIndex(func(_arg NullableGroupInviteRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableGroupInviteRequestStream) Clean() *NullableGroupInviteRequestStream {
	*self = NullableGroupInviteRequestStreamOf()
	return self
}
func (self *NullableGroupInviteRequestStream) Delete(index int) *NullableGroupInviteRequestStream {
	return self.DeleteRange(index, index)
}
func (self *NullableGroupInviteRequestStream) DeleteRange(startIndex, endIndex int) *NullableGroupInviteRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableGroupInviteRequestStream) Distinct() *NullableGroupInviteRequestStream {
	caches := map[string]bool{}
	result := NullableGroupInviteRequestStreamOf()
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
func (self *NullableGroupInviteRequestStream) Each(fn func(NullableGroupInviteRequest)) *NullableGroupInviteRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableGroupInviteRequestStream) EachRight(fn func(NullableGroupInviteRequest)) *NullableGroupInviteRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableGroupInviteRequestStream) Equals(arr []NullableGroupInviteRequest) bool {
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
func (self *NullableGroupInviteRequestStream) Filter(fn func(NullableGroupInviteRequest, int) bool) *NullableGroupInviteRequestStream {
	result := NullableGroupInviteRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableGroupInviteRequestStream) FilterSlim(fn func(NullableGroupInviteRequest, int) bool) *NullableGroupInviteRequestStream {
	result := NullableGroupInviteRequestStreamOf()
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
func (self *NullableGroupInviteRequestStream) Find(fn func(NullableGroupInviteRequest, int) bool) *NullableGroupInviteRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableGroupInviteRequestStream) FindOr(fn func(NullableGroupInviteRequest, int) bool, or NullableGroupInviteRequest) NullableGroupInviteRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableGroupInviteRequestStream) FindIndex(fn func(NullableGroupInviteRequest, int) bool) int {
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
func (self *NullableGroupInviteRequestStream) First() *NullableGroupInviteRequest {
	return self.Get(0)
}
func (self *NullableGroupInviteRequestStream) FirstOr(arg NullableGroupInviteRequest) NullableGroupInviteRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableGroupInviteRequestStream) ForEach(fn func(NullableGroupInviteRequest, int)) *NullableGroupInviteRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableGroupInviteRequestStream) ForEachRight(fn func(NullableGroupInviteRequest, int)) *NullableGroupInviteRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableGroupInviteRequestStream) GroupBy(fn func(NullableGroupInviteRequest, int) string) map[string][]NullableGroupInviteRequest {
	m := map[string][]NullableGroupInviteRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableGroupInviteRequestStream) GroupByValues(fn func(NullableGroupInviteRequest, int) string) [][]NullableGroupInviteRequest {
	var tmp [][]NullableGroupInviteRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableGroupInviteRequestStream) IndexOf(arg NullableGroupInviteRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableGroupInviteRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableGroupInviteRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableGroupInviteRequestStream) Last() *NullableGroupInviteRequest {
	return self.Get(self.Len() - 1)
}
func (self *NullableGroupInviteRequestStream) LastOr(arg NullableGroupInviteRequest) NullableGroupInviteRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableGroupInviteRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableGroupInviteRequestStream) Limit(limit int) *NullableGroupInviteRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableGroupInviteRequestStream) Map(fn func(NullableGroupInviteRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupInviteRequestStream) Map2Int(fn func(NullableGroupInviteRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupInviteRequestStream) Map2Int32(fn func(NullableGroupInviteRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupInviteRequestStream) Map2Int64(fn func(NullableGroupInviteRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupInviteRequestStream) Map2Float32(fn func(NullableGroupInviteRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupInviteRequestStream) Map2Float64(fn func(NullableGroupInviteRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupInviteRequestStream) Map2Bool(fn func(NullableGroupInviteRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupInviteRequestStream) Map2Bytes(fn func(NullableGroupInviteRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupInviteRequestStream) Map2String(fn func(NullableGroupInviteRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupInviteRequestStream) Max(fn func(NullableGroupInviteRequest, int) float64) *NullableGroupInviteRequest {
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
func (self *NullableGroupInviteRequestStream) Min(fn func(NullableGroupInviteRequest, int) float64) *NullableGroupInviteRequest {
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
func (self *NullableGroupInviteRequestStream) NoneMatch(fn func(NullableGroupInviteRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableGroupInviteRequestStream) Get(index int) *NullableGroupInviteRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableGroupInviteRequestStream) GetOr(index int, arg NullableGroupInviteRequest) NullableGroupInviteRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableGroupInviteRequestStream) Peek(fn func(*NullableGroupInviteRequest, int)) *NullableGroupInviteRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableGroupInviteRequestStream) Reduce(fn func(NullableGroupInviteRequest, NullableGroupInviteRequest, int) NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	return self.ReduceInit(fn, NullableGroupInviteRequest{})
}
func (self *NullableGroupInviteRequestStream) ReduceInit(fn func(NullableGroupInviteRequest, NullableGroupInviteRequest, int) NullableGroupInviteRequest, initialValue NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	result := NullableGroupInviteRequestStreamOf()
	self.ForEach(func(v NullableGroupInviteRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableGroupInviteRequestStream) ReduceInterface(fn func(interface{}, NullableGroupInviteRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableGroupInviteRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableGroupInviteRequestStream) ReduceString(fn func(string, NullableGroupInviteRequest, int) string) []string {
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
func (self *NullableGroupInviteRequestStream) ReduceInt(fn func(int, NullableGroupInviteRequest, int) int) []int {
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
func (self *NullableGroupInviteRequestStream) ReduceInt32(fn func(int32, NullableGroupInviteRequest, int) int32) []int32 {
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
func (self *NullableGroupInviteRequestStream) ReduceInt64(fn func(int64, NullableGroupInviteRequest, int) int64) []int64 {
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
func (self *NullableGroupInviteRequestStream) ReduceFloat32(fn func(float32, NullableGroupInviteRequest, int) float32) []float32 {
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
func (self *NullableGroupInviteRequestStream) ReduceFloat64(fn func(float64, NullableGroupInviteRequest, int) float64) []float64 {
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
func (self *NullableGroupInviteRequestStream) ReduceBool(fn func(bool, NullableGroupInviteRequest, int) bool) []bool {
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
func (self *NullableGroupInviteRequestStream) Reverse() *NullableGroupInviteRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableGroupInviteRequestStream) Replace(fn func(NullableGroupInviteRequest, int) NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	return self.ForEach(func(v NullableGroupInviteRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableGroupInviteRequestStream) Select(fn func(NullableGroupInviteRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableGroupInviteRequestStream) Set(index int, val NullableGroupInviteRequest) *NullableGroupInviteRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableGroupInviteRequestStream) Skip(skip int) *NullableGroupInviteRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableGroupInviteRequestStream) SkippingEach(fn func(NullableGroupInviteRequest, int) int) *NullableGroupInviteRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableGroupInviteRequestStream) Slice(startIndex, n int) *NullableGroupInviteRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableGroupInviteRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableGroupInviteRequestStream) Sort(fn func(i, j int) bool) *NullableGroupInviteRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableGroupInviteRequestStream) Tail() *NullableGroupInviteRequest {
	return self.Last()
}
func (self *NullableGroupInviteRequestStream) TailOr(arg NullableGroupInviteRequest) NullableGroupInviteRequest {
	return self.LastOr(arg)
}
func (self *NullableGroupInviteRequestStream) ToList() []NullableGroupInviteRequest {
	return self.Val()
}
func (self *NullableGroupInviteRequestStream) Unique() *NullableGroupInviteRequestStream {
	return self.Distinct()
}
func (self *NullableGroupInviteRequestStream) Val() []NullableGroupInviteRequest {
	if self == nil {
		return []NullableGroupInviteRequest{}
	}
	return *self.Copy()
}
func (self *NullableGroupInviteRequestStream) While(fn func(NullableGroupInviteRequest, int) bool) *NullableGroupInviteRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableGroupInviteRequestStream) Where(fn func(NullableGroupInviteRequest) bool) *NullableGroupInviteRequestStream {
	result := NullableGroupInviteRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableGroupInviteRequestStream) WhereSlim(fn func(NullableGroupInviteRequest) bool) *NullableGroupInviteRequestStream {
	result := NullableGroupInviteRequestStreamOf()
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
