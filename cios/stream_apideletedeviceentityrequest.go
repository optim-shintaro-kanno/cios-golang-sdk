/*
 * Collection utility of ApiDeleteDeviceEntityRequest Struct
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

type ApiDeleteDeviceEntityRequestStream []ApiDeleteDeviceEntityRequest

func ApiDeleteDeviceEntityRequestStreamOf(arg ...ApiDeleteDeviceEntityRequest) ApiDeleteDeviceEntityRequestStream {
	return arg
}
func ApiDeleteDeviceEntityRequestStreamFrom(arg []ApiDeleteDeviceEntityRequest) ApiDeleteDeviceEntityRequestStream {
	return arg
}
func CreateApiDeleteDeviceEntityRequestStream(arg ...ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	tmp := ApiDeleteDeviceEntityRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiDeleteDeviceEntityRequestStream(arg []ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	tmp := ApiDeleteDeviceEntityRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiDeleteDeviceEntityRequestStream) Add(arg ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	return self.AddAll(arg)
}
func (self *ApiDeleteDeviceEntityRequestStream) AddAll(arg ...ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) AddSafe(arg *ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) Aggregate(fn func(ApiDeleteDeviceEntityRequest, ApiDeleteDeviceEntityRequest) ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	result := ApiDeleteDeviceEntityRequestStreamOf()
	self.ForEach(func(v ApiDeleteDeviceEntityRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiDeleteDeviceEntityRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) AllMatch(fn func(ApiDeleteDeviceEntityRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiDeleteDeviceEntityRequestStream) AnyMatch(fn func(ApiDeleteDeviceEntityRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiDeleteDeviceEntityRequestStream) Clone() *ApiDeleteDeviceEntityRequestStream {
	temp := make([]ApiDeleteDeviceEntityRequest, self.Len())
	copy(temp, *self)
	return (*ApiDeleteDeviceEntityRequestStream)(&temp)
}
func (self *ApiDeleteDeviceEntityRequestStream) Copy() *ApiDeleteDeviceEntityRequestStream {
	return self.Clone()
}
func (self *ApiDeleteDeviceEntityRequestStream) Concat(arg []ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiDeleteDeviceEntityRequestStream) Contains(arg ApiDeleteDeviceEntityRequest) bool {
	return self.FindIndex(func(_arg ApiDeleteDeviceEntityRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiDeleteDeviceEntityRequestStream) Clean() *ApiDeleteDeviceEntityRequestStream {
	*self = ApiDeleteDeviceEntityRequestStreamOf()
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) Delete(index int) *ApiDeleteDeviceEntityRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiDeleteDeviceEntityRequestStream) DeleteRange(startIndex, endIndex int) *ApiDeleteDeviceEntityRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) Distinct() *ApiDeleteDeviceEntityRequestStream {
	caches := map[string]bool{}
	result := ApiDeleteDeviceEntityRequestStreamOf()
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
func (self *ApiDeleteDeviceEntityRequestStream) Each(fn func(ApiDeleteDeviceEntityRequest)) *ApiDeleteDeviceEntityRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) EachRight(fn func(ApiDeleteDeviceEntityRequest)) *ApiDeleteDeviceEntityRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) Equals(arr []ApiDeleteDeviceEntityRequest) bool {
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
func (self *ApiDeleteDeviceEntityRequestStream) Filter(fn func(ApiDeleteDeviceEntityRequest, int) bool) *ApiDeleteDeviceEntityRequestStream {
	result := ApiDeleteDeviceEntityRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) FilterSlim(fn func(ApiDeleteDeviceEntityRequest, int) bool) *ApiDeleteDeviceEntityRequestStream {
	result := ApiDeleteDeviceEntityRequestStreamOf()
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
func (self *ApiDeleteDeviceEntityRequestStream) Find(fn func(ApiDeleteDeviceEntityRequest, int) bool) *ApiDeleteDeviceEntityRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiDeleteDeviceEntityRequestStream) FindOr(fn func(ApiDeleteDeviceEntityRequest, int) bool, or ApiDeleteDeviceEntityRequest) ApiDeleteDeviceEntityRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiDeleteDeviceEntityRequestStream) FindIndex(fn func(ApiDeleteDeviceEntityRequest, int) bool) int {
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
func (self *ApiDeleteDeviceEntityRequestStream) First() *ApiDeleteDeviceEntityRequest {
	return self.Get(0)
}
func (self *ApiDeleteDeviceEntityRequestStream) FirstOr(arg ApiDeleteDeviceEntityRequest) ApiDeleteDeviceEntityRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteDeviceEntityRequestStream) ForEach(fn func(ApiDeleteDeviceEntityRequest, int)) *ApiDeleteDeviceEntityRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) ForEachRight(fn func(ApiDeleteDeviceEntityRequest, int)) *ApiDeleteDeviceEntityRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) GroupBy(fn func(ApiDeleteDeviceEntityRequest, int) string) map[string][]ApiDeleteDeviceEntityRequest {
	m := map[string][]ApiDeleteDeviceEntityRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiDeleteDeviceEntityRequestStream) GroupByValues(fn func(ApiDeleteDeviceEntityRequest, int) string) [][]ApiDeleteDeviceEntityRequest {
	var tmp [][]ApiDeleteDeviceEntityRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiDeleteDeviceEntityRequestStream) IndexOf(arg ApiDeleteDeviceEntityRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiDeleteDeviceEntityRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiDeleteDeviceEntityRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiDeleteDeviceEntityRequestStream) Last() *ApiDeleteDeviceEntityRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiDeleteDeviceEntityRequestStream) LastOr(arg ApiDeleteDeviceEntityRequest) ApiDeleteDeviceEntityRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteDeviceEntityRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiDeleteDeviceEntityRequestStream) Limit(limit int) *ApiDeleteDeviceEntityRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiDeleteDeviceEntityRequestStream) Map(fn func(ApiDeleteDeviceEntityRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntityRequestStream) Map2Int(fn func(ApiDeleteDeviceEntityRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntityRequestStream) Map2Int32(fn func(ApiDeleteDeviceEntityRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntityRequestStream) Map2Int64(fn func(ApiDeleteDeviceEntityRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntityRequestStream) Map2Float32(fn func(ApiDeleteDeviceEntityRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntityRequestStream) Map2Float64(fn func(ApiDeleteDeviceEntityRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntityRequestStream) Map2Bool(fn func(ApiDeleteDeviceEntityRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntityRequestStream) Map2Bytes(fn func(ApiDeleteDeviceEntityRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntityRequestStream) Map2String(fn func(ApiDeleteDeviceEntityRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntityRequestStream) Max(fn func(ApiDeleteDeviceEntityRequest, int) float64) *ApiDeleteDeviceEntityRequest {
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
func (self *ApiDeleteDeviceEntityRequestStream) Min(fn func(ApiDeleteDeviceEntityRequest, int) float64) *ApiDeleteDeviceEntityRequest {
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
func (self *ApiDeleteDeviceEntityRequestStream) NoneMatch(fn func(ApiDeleteDeviceEntityRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiDeleteDeviceEntityRequestStream) Get(index int) *ApiDeleteDeviceEntityRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiDeleteDeviceEntityRequestStream) GetOr(index int, arg ApiDeleteDeviceEntityRequest) ApiDeleteDeviceEntityRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteDeviceEntityRequestStream) Peek(fn func(*ApiDeleteDeviceEntityRequest, int)) *ApiDeleteDeviceEntityRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *ApiDeleteDeviceEntityRequestStream) Reduce(fn func(ApiDeleteDeviceEntityRequest, ApiDeleteDeviceEntityRequest, int) ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	return self.ReduceInit(fn, ApiDeleteDeviceEntityRequest{})
}
func (self *ApiDeleteDeviceEntityRequestStream) ReduceInit(fn func(ApiDeleteDeviceEntityRequest, ApiDeleteDeviceEntityRequest, int) ApiDeleteDeviceEntityRequest, initialValue ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	result := ApiDeleteDeviceEntityRequestStreamOf()
	self.ForEach(func(v ApiDeleteDeviceEntityRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) ReduceInterface(fn func(interface{}, ApiDeleteDeviceEntityRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiDeleteDeviceEntityRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiDeleteDeviceEntityRequestStream) ReduceString(fn func(string, ApiDeleteDeviceEntityRequest, int) string) []string {
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
func (self *ApiDeleteDeviceEntityRequestStream) ReduceInt(fn func(int, ApiDeleteDeviceEntityRequest, int) int) []int {
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
func (self *ApiDeleteDeviceEntityRequestStream) ReduceInt32(fn func(int32, ApiDeleteDeviceEntityRequest, int) int32) []int32 {
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
func (self *ApiDeleteDeviceEntityRequestStream) ReduceInt64(fn func(int64, ApiDeleteDeviceEntityRequest, int) int64) []int64 {
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
func (self *ApiDeleteDeviceEntityRequestStream) ReduceFloat32(fn func(float32, ApiDeleteDeviceEntityRequest, int) float32) []float32 {
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
func (self *ApiDeleteDeviceEntityRequestStream) ReduceFloat64(fn func(float64, ApiDeleteDeviceEntityRequest, int) float64) []float64 {
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
func (self *ApiDeleteDeviceEntityRequestStream) ReduceBool(fn func(bool, ApiDeleteDeviceEntityRequest, int) bool) []bool {
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
func (self *ApiDeleteDeviceEntityRequestStream) Reverse() *ApiDeleteDeviceEntityRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) Replace(fn func(ApiDeleteDeviceEntityRequest, int) ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	return self.ForEach(func(v ApiDeleteDeviceEntityRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiDeleteDeviceEntityRequestStream) Select(fn func(ApiDeleteDeviceEntityRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiDeleteDeviceEntityRequestStream) Set(index int, val ApiDeleteDeviceEntityRequest) *ApiDeleteDeviceEntityRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) Skip(skip int) *ApiDeleteDeviceEntityRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiDeleteDeviceEntityRequestStream) SkippingEach(fn func(ApiDeleteDeviceEntityRequest, int) int) *ApiDeleteDeviceEntityRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) Slice(startIndex, n int) *ApiDeleteDeviceEntityRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiDeleteDeviceEntityRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) Sort(fn func(i, j int) bool) *ApiDeleteDeviceEntityRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiDeleteDeviceEntityRequestStream) Tail() *ApiDeleteDeviceEntityRequest {
	return self.Last()
}
func (self *ApiDeleteDeviceEntityRequestStream) TailOr(arg ApiDeleteDeviceEntityRequest) ApiDeleteDeviceEntityRequest {
	return self.LastOr(arg)
}
func (self *ApiDeleteDeviceEntityRequestStream) ToList() []ApiDeleteDeviceEntityRequest {
	return self.Val()
}
func (self *ApiDeleteDeviceEntityRequestStream) Unique() *ApiDeleteDeviceEntityRequestStream {
	return self.Distinct()
}
func (self *ApiDeleteDeviceEntityRequestStream) Val() []ApiDeleteDeviceEntityRequest {
	if self == nil {
		return []ApiDeleteDeviceEntityRequest{}
	}
	return *self.Copy()
}
func (self *ApiDeleteDeviceEntityRequestStream) While(fn func(ApiDeleteDeviceEntityRequest, int) bool) *ApiDeleteDeviceEntityRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) Where(fn func(ApiDeleteDeviceEntityRequest) bool) *ApiDeleteDeviceEntityRequestStream {
	result := ApiDeleteDeviceEntityRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeleteDeviceEntityRequestStream) WhereSlim(fn func(ApiDeleteDeviceEntityRequest) bool) *ApiDeleteDeviceEntityRequestStream {
	result := ApiDeleteDeviceEntityRequestStreamOf()
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
