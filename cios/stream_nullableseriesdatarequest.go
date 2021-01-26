/*
 * Collection utility of NullableSeriesDataRequest Struct
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

type NullableSeriesDataRequestStream []NullableSeriesDataRequest

func NullableSeriesDataRequestStreamOf(arg ...NullableSeriesDataRequest) NullableSeriesDataRequestStream {
	return arg
}
func NullableSeriesDataRequestStreamFrom(arg []NullableSeriesDataRequest) NullableSeriesDataRequestStream {
	return arg
}
func CreateNullableSeriesDataRequestStream(arg ...NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	tmp := NullableSeriesDataRequestStreamOf(arg...)
	return &tmp
}
func GenerateNullableSeriesDataRequestStream(arg []NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	tmp := NullableSeriesDataRequestStreamFrom(arg)
	return &tmp
}

func (self *NullableSeriesDataRequestStream) Add(arg NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	return self.AddAll(arg)
}
func (self *NullableSeriesDataRequestStream) AddAll(arg ...NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSeriesDataRequestStream) AddSafe(arg *NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSeriesDataRequestStream) Aggregate(fn func(NullableSeriesDataRequest, NullableSeriesDataRequest) NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	result := NullableSeriesDataRequestStreamOf()
	self.ForEach(func(v NullableSeriesDataRequest, i int) {
		if i == 0 {
			result.Add(fn(NullableSeriesDataRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesDataRequestStream) AllMatch(fn func(NullableSeriesDataRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSeriesDataRequestStream) AnyMatch(fn func(NullableSeriesDataRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSeriesDataRequestStream) Clone() *NullableSeriesDataRequestStream {
	temp := make([]NullableSeriesDataRequest, self.Len())
	copy(temp, *self)
	return (*NullableSeriesDataRequestStream)(&temp)
}
func (self *NullableSeriesDataRequestStream) Copy() *NullableSeriesDataRequestStream {
	return self.Clone()
}
func (self *NullableSeriesDataRequestStream) Concat(arg []NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	return self.AddAll(arg...)
}
func (self *NullableSeriesDataRequestStream) Contains(arg NullableSeriesDataRequest) bool {
	return self.FindIndex(func(_arg NullableSeriesDataRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSeriesDataRequestStream) Clean() *NullableSeriesDataRequestStream {
	*self = NullableSeriesDataRequestStreamOf()
	return self
}
func (self *NullableSeriesDataRequestStream) Delete(index int) *NullableSeriesDataRequestStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSeriesDataRequestStream) DeleteRange(startIndex, endIndex int) *NullableSeriesDataRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSeriesDataRequestStream) Distinct() *NullableSeriesDataRequestStream {
	caches := map[string]bool{}
	result := NullableSeriesDataRequestStreamOf()
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
func (self *NullableSeriesDataRequestStream) Each(fn func(NullableSeriesDataRequest)) *NullableSeriesDataRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSeriesDataRequestStream) EachRight(fn func(NullableSeriesDataRequest)) *NullableSeriesDataRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSeriesDataRequestStream) Equals(arr []NullableSeriesDataRequest) bool {
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
func (self *NullableSeriesDataRequestStream) Filter(fn func(NullableSeriesDataRequest, int) bool) *NullableSeriesDataRequestStream {
	result := NullableSeriesDataRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesDataRequestStream) FilterSlim(fn func(NullableSeriesDataRequest, int) bool) *NullableSeriesDataRequestStream {
	result := NullableSeriesDataRequestStreamOf()
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
func (self *NullableSeriesDataRequestStream) Find(fn func(NullableSeriesDataRequest, int) bool) *NullableSeriesDataRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesDataRequestStream) FindOr(fn func(NullableSeriesDataRequest, int) bool, or NullableSeriesDataRequest) NullableSeriesDataRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSeriesDataRequestStream) FindIndex(fn func(NullableSeriesDataRequest, int) bool) int {
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
func (self *NullableSeriesDataRequestStream) First() *NullableSeriesDataRequest {
	return self.Get(0)
}
func (self *NullableSeriesDataRequestStream) FirstOr(arg NullableSeriesDataRequest) NullableSeriesDataRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesDataRequestStream) ForEach(fn func(NullableSeriesDataRequest, int)) *NullableSeriesDataRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSeriesDataRequestStream) ForEachRight(fn func(NullableSeriesDataRequest, int)) *NullableSeriesDataRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSeriesDataRequestStream) GroupBy(fn func(NullableSeriesDataRequest, int) string) map[string][]NullableSeriesDataRequest {
	m := map[string][]NullableSeriesDataRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSeriesDataRequestStream) GroupByValues(fn func(NullableSeriesDataRequest, int) string) [][]NullableSeriesDataRequest {
	var tmp [][]NullableSeriesDataRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSeriesDataRequestStream) IndexOf(arg NullableSeriesDataRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSeriesDataRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSeriesDataRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSeriesDataRequestStream) Last() *NullableSeriesDataRequest {
	return self.Get(self.Len() - 1)
}
func (self *NullableSeriesDataRequestStream) LastOr(arg NullableSeriesDataRequest) NullableSeriesDataRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesDataRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSeriesDataRequestStream) Limit(limit int) *NullableSeriesDataRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSeriesDataRequestStream) Map(fn func(NullableSeriesDataRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataRequestStream) Map2Int(fn func(NullableSeriesDataRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataRequestStream) Map2Int32(fn func(NullableSeriesDataRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataRequestStream) Map2Int64(fn func(NullableSeriesDataRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataRequestStream) Map2Float32(fn func(NullableSeriesDataRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataRequestStream) Map2Float64(fn func(NullableSeriesDataRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataRequestStream) Map2Bool(fn func(NullableSeriesDataRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataRequestStream) Map2Bytes(fn func(NullableSeriesDataRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataRequestStream) Map2String(fn func(NullableSeriesDataRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataRequestStream) Max(fn func(NullableSeriesDataRequest, int) float64) *NullableSeriesDataRequest {
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
func (self *NullableSeriesDataRequestStream) Min(fn func(NullableSeriesDataRequest, int) float64) *NullableSeriesDataRequest {
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
func (self *NullableSeriesDataRequestStream) NoneMatch(fn func(NullableSeriesDataRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSeriesDataRequestStream) Get(index int) *NullableSeriesDataRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesDataRequestStream) GetOr(index int, arg NullableSeriesDataRequest) NullableSeriesDataRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesDataRequestStream) Peek(fn func(*NullableSeriesDataRequest, int)) *NullableSeriesDataRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableSeriesDataRequestStream) Reduce(fn func(NullableSeriesDataRequest, NullableSeriesDataRequest, int) NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	return self.ReduceInit(fn, NullableSeriesDataRequest{})
}
func (self *NullableSeriesDataRequestStream) ReduceInit(fn func(NullableSeriesDataRequest, NullableSeriesDataRequest, int) NullableSeriesDataRequest, initialValue NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	result := NullableSeriesDataRequestStreamOf()
	self.ForEach(func(v NullableSeriesDataRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesDataRequestStream) ReduceInterface(fn func(interface{}, NullableSeriesDataRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSeriesDataRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSeriesDataRequestStream) ReduceString(fn func(string, NullableSeriesDataRequest, int) string) []string {
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
func (self *NullableSeriesDataRequestStream) ReduceInt(fn func(int, NullableSeriesDataRequest, int) int) []int {
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
func (self *NullableSeriesDataRequestStream) ReduceInt32(fn func(int32, NullableSeriesDataRequest, int) int32) []int32 {
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
func (self *NullableSeriesDataRequestStream) ReduceInt64(fn func(int64, NullableSeriesDataRequest, int) int64) []int64 {
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
func (self *NullableSeriesDataRequestStream) ReduceFloat32(fn func(float32, NullableSeriesDataRequest, int) float32) []float32 {
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
func (self *NullableSeriesDataRequestStream) ReduceFloat64(fn func(float64, NullableSeriesDataRequest, int) float64) []float64 {
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
func (self *NullableSeriesDataRequestStream) ReduceBool(fn func(bool, NullableSeriesDataRequest, int) bool) []bool {
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
func (self *NullableSeriesDataRequestStream) Reverse() *NullableSeriesDataRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSeriesDataRequestStream) Replace(fn func(NullableSeriesDataRequest, int) NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	return self.ForEach(func(v NullableSeriesDataRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSeriesDataRequestStream) Select(fn func(NullableSeriesDataRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSeriesDataRequestStream) Set(index int, val NullableSeriesDataRequest) *NullableSeriesDataRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSeriesDataRequestStream) Skip(skip int) *NullableSeriesDataRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSeriesDataRequestStream) SkippingEach(fn func(NullableSeriesDataRequest, int) int) *NullableSeriesDataRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSeriesDataRequestStream) Slice(startIndex, n int) *NullableSeriesDataRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSeriesDataRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSeriesDataRequestStream) Sort(fn func(i, j int) bool) *NullableSeriesDataRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSeriesDataRequestStream) Tail() *NullableSeriesDataRequest {
	return self.Last()
}
func (self *NullableSeriesDataRequestStream) TailOr(arg NullableSeriesDataRequest) NullableSeriesDataRequest {
	return self.LastOr(arg)
}
func (self *NullableSeriesDataRequestStream) ToList() []NullableSeriesDataRequest {
	return self.Val()
}
func (self *NullableSeriesDataRequestStream) Unique() *NullableSeriesDataRequestStream {
	return self.Distinct()
}
func (self *NullableSeriesDataRequestStream) Val() []NullableSeriesDataRequest {
	if self == nil {
		return []NullableSeriesDataRequest{}
	}
	return *self.Copy()
}
func (self *NullableSeriesDataRequestStream) While(fn func(NullableSeriesDataRequest, int) bool) *NullableSeriesDataRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSeriesDataRequestStream) Where(fn func(NullableSeriesDataRequest) bool) *NullableSeriesDataRequestStream {
	result := NullableSeriesDataRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesDataRequestStream) WhereSlim(fn func(NullableSeriesDataRequest) bool) *NullableSeriesDataRequestStream {
	result := NullableSeriesDataRequestStreamOf()
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
