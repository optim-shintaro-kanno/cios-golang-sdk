/*
 * Collection utility of ApiSubscribeMessageRequest Struct
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

type ApiSubscribeMessageRequestStream []ApiSubscribeMessageRequest

func ApiSubscribeMessageRequestStreamOf(arg ...ApiSubscribeMessageRequest) ApiSubscribeMessageRequestStream {
	return arg
}
func ApiSubscribeMessageRequestStreamFrom(arg []ApiSubscribeMessageRequest) ApiSubscribeMessageRequestStream {
	return arg
}
func CreateApiSubscribeMessageRequestStream(arg ...ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	tmp := ApiSubscribeMessageRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiSubscribeMessageRequestStream(arg []ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	tmp := ApiSubscribeMessageRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiSubscribeMessageRequestStream) Add(arg ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	return self.AddAll(arg)
}
func (self *ApiSubscribeMessageRequestStream) AddAll(arg ...ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiSubscribeMessageRequestStream) AddSafe(arg *ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiSubscribeMessageRequestStream) Aggregate(fn func(ApiSubscribeMessageRequest, ApiSubscribeMessageRequest) ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	result := ApiSubscribeMessageRequestStreamOf()
	self.ForEach(func(v ApiSubscribeMessageRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiSubscribeMessageRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiSubscribeMessageRequestStream) AllMatch(fn func(ApiSubscribeMessageRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiSubscribeMessageRequestStream) AnyMatch(fn func(ApiSubscribeMessageRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiSubscribeMessageRequestStream) Clone() *ApiSubscribeMessageRequestStream {
	temp := make([]ApiSubscribeMessageRequest, self.Len())
	copy(temp, *self)
	return (*ApiSubscribeMessageRequestStream)(&temp)
}
func (self *ApiSubscribeMessageRequestStream) Copy() *ApiSubscribeMessageRequestStream {
	return self.Clone()
}
func (self *ApiSubscribeMessageRequestStream) Concat(arg []ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiSubscribeMessageRequestStream) Contains(arg ApiSubscribeMessageRequest) bool {
	return self.FindIndex(func(_arg ApiSubscribeMessageRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiSubscribeMessageRequestStream) Clean() *ApiSubscribeMessageRequestStream {
	*self = ApiSubscribeMessageRequestStreamOf()
	return self
}
func (self *ApiSubscribeMessageRequestStream) Delete(index int) *ApiSubscribeMessageRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiSubscribeMessageRequestStream) DeleteRange(startIndex, endIndex int) *ApiSubscribeMessageRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiSubscribeMessageRequestStream) Distinct() *ApiSubscribeMessageRequestStream {
	caches := map[string]bool{}
	result := ApiSubscribeMessageRequestStreamOf()
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
func (self *ApiSubscribeMessageRequestStream) Each(fn func(ApiSubscribeMessageRequest)) *ApiSubscribeMessageRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiSubscribeMessageRequestStream) EachRight(fn func(ApiSubscribeMessageRequest)) *ApiSubscribeMessageRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiSubscribeMessageRequestStream) Equals(arr []ApiSubscribeMessageRequest) bool {
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
func (self *ApiSubscribeMessageRequestStream) Filter(fn func(ApiSubscribeMessageRequest, int) bool) *ApiSubscribeMessageRequestStream {
	result := ApiSubscribeMessageRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiSubscribeMessageRequestStream) FilterSlim(fn func(ApiSubscribeMessageRequest, int) bool) *ApiSubscribeMessageRequestStream {
	result := ApiSubscribeMessageRequestStreamOf()
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
func (self *ApiSubscribeMessageRequestStream) Find(fn func(ApiSubscribeMessageRequest, int) bool) *ApiSubscribeMessageRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiSubscribeMessageRequestStream) FindOr(fn func(ApiSubscribeMessageRequest, int) bool, or ApiSubscribeMessageRequest) ApiSubscribeMessageRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiSubscribeMessageRequestStream) FindIndex(fn func(ApiSubscribeMessageRequest, int) bool) int {
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
func (self *ApiSubscribeMessageRequestStream) First() *ApiSubscribeMessageRequest {
	return self.Get(0)
}
func (self *ApiSubscribeMessageRequestStream) FirstOr(arg ApiSubscribeMessageRequest) ApiSubscribeMessageRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiSubscribeMessageRequestStream) ForEach(fn func(ApiSubscribeMessageRequest, int)) *ApiSubscribeMessageRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiSubscribeMessageRequestStream) ForEachRight(fn func(ApiSubscribeMessageRequest, int)) *ApiSubscribeMessageRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiSubscribeMessageRequestStream) GroupBy(fn func(ApiSubscribeMessageRequest, int) string) map[string][]ApiSubscribeMessageRequest {
	m := map[string][]ApiSubscribeMessageRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiSubscribeMessageRequestStream) GroupByValues(fn func(ApiSubscribeMessageRequest, int) string) [][]ApiSubscribeMessageRequest {
	var tmp [][]ApiSubscribeMessageRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiSubscribeMessageRequestStream) IndexOf(arg ApiSubscribeMessageRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiSubscribeMessageRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiSubscribeMessageRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiSubscribeMessageRequestStream) Last() *ApiSubscribeMessageRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiSubscribeMessageRequestStream) LastOr(arg ApiSubscribeMessageRequest) ApiSubscribeMessageRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiSubscribeMessageRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiSubscribeMessageRequestStream) Limit(limit int) *ApiSubscribeMessageRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiSubscribeMessageRequestStream) Map(fn func(ApiSubscribeMessageRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSubscribeMessageRequestStream) Map2Int(fn func(ApiSubscribeMessageRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSubscribeMessageRequestStream) Map2Int32(fn func(ApiSubscribeMessageRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSubscribeMessageRequestStream) Map2Int64(fn func(ApiSubscribeMessageRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSubscribeMessageRequestStream) Map2Float32(fn func(ApiSubscribeMessageRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSubscribeMessageRequestStream) Map2Float64(fn func(ApiSubscribeMessageRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSubscribeMessageRequestStream) Map2Bool(fn func(ApiSubscribeMessageRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSubscribeMessageRequestStream) Map2Bytes(fn func(ApiSubscribeMessageRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSubscribeMessageRequestStream) Map2String(fn func(ApiSubscribeMessageRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiSubscribeMessageRequestStream) Max(fn func(ApiSubscribeMessageRequest, int) float64) *ApiSubscribeMessageRequest {
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
func (self *ApiSubscribeMessageRequestStream) Min(fn func(ApiSubscribeMessageRequest, int) float64) *ApiSubscribeMessageRequest {
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
func (self *ApiSubscribeMessageRequestStream) NoneMatch(fn func(ApiSubscribeMessageRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiSubscribeMessageRequestStream) Get(index int) *ApiSubscribeMessageRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiSubscribeMessageRequestStream) GetOr(index int, arg ApiSubscribeMessageRequest) ApiSubscribeMessageRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiSubscribeMessageRequestStream) Peek(fn func(*ApiSubscribeMessageRequest, int)) *ApiSubscribeMessageRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *ApiSubscribeMessageRequestStream) Reduce(fn func(ApiSubscribeMessageRequest, ApiSubscribeMessageRequest, int) ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	return self.ReduceInit(fn, ApiSubscribeMessageRequest{})
}
func (self *ApiSubscribeMessageRequestStream) ReduceInit(fn func(ApiSubscribeMessageRequest, ApiSubscribeMessageRequest, int) ApiSubscribeMessageRequest, initialValue ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	result := ApiSubscribeMessageRequestStreamOf()
	self.ForEach(func(v ApiSubscribeMessageRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiSubscribeMessageRequestStream) ReduceInterface(fn func(interface{}, ApiSubscribeMessageRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiSubscribeMessageRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiSubscribeMessageRequestStream) ReduceString(fn func(string, ApiSubscribeMessageRequest, int) string) []string {
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
func (self *ApiSubscribeMessageRequestStream) ReduceInt(fn func(int, ApiSubscribeMessageRequest, int) int) []int {
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
func (self *ApiSubscribeMessageRequestStream) ReduceInt32(fn func(int32, ApiSubscribeMessageRequest, int) int32) []int32 {
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
func (self *ApiSubscribeMessageRequestStream) ReduceInt64(fn func(int64, ApiSubscribeMessageRequest, int) int64) []int64 {
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
func (self *ApiSubscribeMessageRequestStream) ReduceFloat32(fn func(float32, ApiSubscribeMessageRequest, int) float32) []float32 {
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
func (self *ApiSubscribeMessageRequestStream) ReduceFloat64(fn func(float64, ApiSubscribeMessageRequest, int) float64) []float64 {
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
func (self *ApiSubscribeMessageRequestStream) ReduceBool(fn func(bool, ApiSubscribeMessageRequest, int) bool) []bool {
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
func (self *ApiSubscribeMessageRequestStream) Reverse() *ApiSubscribeMessageRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiSubscribeMessageRequestStream) Replace(fn func(ApiSubscribeMessageRequest, int) ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	return self.ForEach(func(v ApiSubscribeMessageRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiSubscribeMessageRequestStream) Select(fn func(ApiSubscribeMessageRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiSubscribeMessageRequestStream) Set(index int, val ApiSubscribeMessageRequest) *ApiSubscribeMessageRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiSubscribeMessageRequestStream) Skip(skip int) *ApiSubscribeMessageRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiSubscribeMessageRequestStream) SkippingEach(fn func(ApiSubscribeMessageRequest, int) int) *ApiSubscribeMessageRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiSubscribeMessageRequestStream) Slice(startIndex, n int) *ApiSubscribeMessageRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiSubscribeMessageRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiSubscribeMessageRequestStream) Sort(fn func(i, j int) bool) *ApiSubscribeMessageRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiSubscribeMessageRequestStream) Tail() *ApiSubscribeMessageRequest {
	return self.Last()
}
func (self *ApiSubscribeMessageRequestStream) TailOr(arg ApiSubscribeMessageRequest) ApiSubscribeMessageRequest {
	return self.LastOr(arg)
}
func (self *ApiSubscribeMessageRequestStream) ToList() []ApiSubscribeMessageRequest {
	return self.Val()
}
func (self *ApiSubscribeMessageRequestStream) Unique() *ApiSubscribeMessageRequestStream {
	return self.Distinct()
}
func (self *ApiSubscribeMessageRequestStream) Val() []ApiSubscribeMessageRequest {
	if self == nil {
		return []ApiSubscribeMessageRequest{}
	}
	return *self.Copy()
}
func (self *ApiSubscribeMessageRequestStream) While(fn func(ApiSubscribeMessageRequest, int) bool) *ApiSubscribeMessageRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiSubscribeMessageRequestStream) Where(fn func(ApiSubscribeMessageRequest) bool) *ApiSubscribeMessageRequestStream {
	result := ApiSubscribeMessageRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiSubscribeMessageRequestStream) WhereSlim(fn func(ApiSubscribeMessageRequest) bool) *ApiSubscribeMessageRequestStream {
	result := ApiSubscribeMessageRequestStreamOf()
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
