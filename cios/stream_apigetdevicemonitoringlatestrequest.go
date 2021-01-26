/*
 * Collection utility of ApiGetDeviceMonitoringLatestRequest Struct
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

type ApiGetDeviceMonitoringLatestRequestStream []ApiGetDeviceMonitoringLatestRequest

func ApiGetDeviceMonitoringLatestRequestStreamOf(arg ...ApiGetDeviceMonitoringLatestRequest) ApiGetDeviceMonitoringLatestRequestStream {
	return arg
}
func ApiGetDeviceMonitoringLatestRequestStreamFrom(arg []ApiGetDeviceMonitoringLatestRequest) ApiGetDeviceMonitoringLatestRequestStream {
	return arg
}
func CreateApiGetDeviceMonitoringLatestRequestStream(arg ...ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	tmp := ApiGetDeviceMonitoringLatestRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetDeviceMonitoringLatestRequestStream(arg []ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	tmp := ApiGetDeviceMonitoringLatestRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetDeviceMonitoringLatestRequestStream) Add(arg ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) AddAll(arg ...ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) AddSafe(arg *ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Aggregate(fn func(ApiGetDeviceMonitoringLatestRequest, ApiGetDeviceMonitoringLatestRequest) ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	result := ApiGetDeviceMonitoringLatestRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceMonitoringLatestRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetDeviceMonitoringLatestRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) AllMatch(fn func(ApiGetDeviceMonitoringLatestRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) AnyMatch(fn func(ApiGetDeviceMonitoringLatestRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Clone() *ApiGetDeviceMonitoringLatestRequestStream {
	temp := make([]ApiGetDeviceMonitoringLatestRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetDeviceMonitoringLatestRequestStream)(&temp)
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Copy() *ApiGetDeviceMonitoringLatestRequestStream {
	return self.Clone()
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Concat(arg []ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Contains(arg ApiGetDeviceMonitoringLatestRequest) bool {
	return self.FindIndex(func(_arg ApiGetDeviceMonitoringLatestRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Clean() *ApiGetDeviceMonitoringLatestRequestStream {
	*self = ApiGetDeviceMonitoringLatestRequestStreamOf()
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Delete(index int) *ApiGetDeviceMonitoringLatestRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetDeviceMonitoringLatestRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Distinct() *ApiGetDeviceMonitoringLatestRequestStream {
	caches := map[string]bool{}
	result := ApiGetDeviceMonitoringLatestRequestStreamOf()
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) Each(fn func(ApiGetDeviceMonitoringLatestRequest)) *ApiGetDeviceMonitoringLatestRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) EachRight(fn func(ApiGetDeviceMonitoringLatestRequest)) *ApiGetDeviceMonitoringLatestRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Equals(arr []ApiGetDeviceMonitoringLatestRequest) bool {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) Filter(fn func(ApiGetDeviceMonitoringLatestRequest, int) bool) *ApiGetDeviceMonitoringLatestRequestStream {
	result := ApiGetDeviceMonitoringLatestRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) FilterSlim(fn func(ApiGetDeviceMonitoringLatestRequest, int) bool) *ApiGetDeviceMonitoringLatestRequestStream {
	result := ApiGetDeviceMonitoringLatestRequestStreamOf()
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) Find(fn func(ApiGetDeviceMonitoringLatestRequest, int) bool) *ApiGetDeviceMonitoringLatestRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) FindOr(fn func(ApiGetDeviceMonitoringLatestRequest, int) bool, or ApiGetDeviceMonitoringLatestRequest) ApiGetDeviceMonitoringLatestRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) FindIndex(fn func(ApiGetDeviceMonitoringLatestRequest, int) bool) int {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) First() *ApiGetDeviceMonitoringLatestRequest {
	return self.Get(0)
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) FirstOr(arg ApiGetDeviceMonitoringLatestRequest) ApiGetDeviceMonitoringLatestRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) ForEach(fn func(ApiGetDeviceMonitoringLatestRequest, int)) *ApiGetDeviceMonitoringLatestRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) ForEachRight(fn func(ApiGetDeviceMonitoringLatestRequest, int)) *ApiGetDeviceMonitoringLatestRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) GroupBy(fn func(ApiGetDeviceMonitoringLatestRequest, int) string) map[string][]ApiGetDeviceMonitoringLatestRequest {
	m := map[string][]ApiGetDeviceMonitoringLatestRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) GroupByValues(fn func(ApiGetDeviceMonitoringLatestRequest, int) string) [][]ApiGetDeviceMonitoringLatestRequest {
	var tmp [][]ApiGetDeviceMonitoringLatestRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) IndexOf(arg ApiGetDeviceMonitoringLatestRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Last() *ApiGetDeviceMonitoringLatestRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) LastOr(arg ApiGetDeviceMonitoringLatestRequest) ApiGetDeviceMonitoringLatestRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Limit(limit int) *ApiGetDeviceMonitoringLatestRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetDeviceMonitoringLatestRequestStream) Map(fn func(ApiGetDeviceMonitoringLatestRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Map2Int(fn func(ApiGetDeviceMonitoringLatestRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Map2Int32(fn func(ApiGetDeviceMonitoringLatestRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Map2Int64(fn func(ApiGetDeviceMonitoringLatestRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Map2Float32(fn func(ApiGetDeviceMonitoringLatestRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Map2Float64(fn func(ApiGetDeviceMonitoringLatestRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Map2Bool(fn func(ApiGetDeviceMonitoringLatestRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Map2Bytes(fn func(ApiGetDeviceMonitoringLatestRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Map2String(fn func(ApiGetDeviceMonitoringLatestRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Max(fn func(ApiGetDeviceMonitoringLatestRequest, int) float64) *ApiGetDeviceMonitoringLatestRequest {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) Min(fn func(ApiGetDeviceMonitoringLatestRequest, int) float64) *ApiGetDeviceMonitoringLatestRequest {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) NoneMatch(fn func(ApiGetDeviceMonitoringLatestRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Get(index int) *ApiGetDeviceMonitoringLatestRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) GetOr(index int, arg ApiGetDeviceMonitoringLatestRequest) ApiGetDeviceMonitoringLatestRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Peek(fn func(*ApiGetDeviceMonitoringLatestRequest, int)) *ApiGetDeviceMonitoringLatestRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *ApiGetDeviceMonitoringLatestRequestStream) Reduce(fn func(ApiGetDeviceMonitoringLatestRequest, ApiGetDeviceMonitoringLatestRequest, int) ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	return self.ReduceInit(fn, ApiGetDeviceMonitoringLatestRequest{})
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) ReduceInit(fn func(ApiGetDeviceMonitoringLatestRequest, ApiGetDeviceMonitoringLatestRequest, int) ApiGetDeviceMonitoringLatestRequest, initialValue ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	result := ApiGetDeviceMonitoringLatestRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceMonitoringLatestRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) ReduceInterface(fn func(interface{}, ApiGetDeviceMonitoringLatestRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetDeviceMonitoringLatestRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) ReduceString(fn func(string, ApiGetDeviceMonitoringLatestRequest, int) string) []string {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) ReduceInt(fn func(int, ApiGetDeviceMonitoringLatestRequest, int) int) []int {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) ReduceInt32(fn func(int32, ApiGetDeviceMonitoringLatestRequest, int) int32) []int32 {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) ReduceInt64(fn func(int64, ApiGetDeviceMonitoringLatestRequest, int) int64) []int64 {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) ReduceFloat32(fn func(float32, ApiGetDeviceMonitoringLatestRequest, int) float32) []float32 {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) ReduceFloat64(fn func(float64, ApiGetDeviceMonitoringLatestRequest, int) float64) []float64 {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) ReduceBool(fn func(bool, ApiGetDeviceMonitoringLatestRequest, int) bool) []bool {
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
func (self *ApiGetDeviceMonitoringLatestRequestStream) Reverse() *ApiGetDeviceMonitoringLatestRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Replace(fn func(ApiGetDeviceMonitoringLatestRequest, int) ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	return self.ForEach(func(v ApiGetDeviceMonitoringLatestRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Select(fn func(ApiGetDeviceMonitoringLatestRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Set(index int, val ApiGetDeviceMonitoringLatestRequest) *ApiGetDeviceMonitoringLatestRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Skip(skip int) *ApiGetDeviceMonitoringLatestRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) SkippingEach(fn func(ApiGetDeviceMonitoringLatestRequest, int) int) *ApiGetDeviceMonitoringLatestRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Slice(startIndex, n int) *ApiGetDeviceMonitoringLatestRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetDeviceMonitoringLatestRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Sort(fn func(i, j int) bool) *ApiGetDeviceMonitoringLatestRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetDeviceMonitoringLatestRequestStream) Tail() *ApiGetDeviceMonitoringLatestRequest {
	return self.Last()
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) TailOr(arg ApiGetDeviceMonitoringLatestRequest) ApiGetDeviceMonitoringLatestRequest {
	return self.LastOr(arg)
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) ToList() []ApiGetDeviceMonitoringLatestRequest {
	return self.Val()
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Unique() *ApiGetDeviceMonitoringLatestRequestStream {
	return self.Distinct()
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Val() []ApiGetDeviceMonitoringLatestRequest {
	if self == nil {
		return []ApiGetDeviceMonitoringLatestRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) While(fn func(ApiGetDeviceMonitoringLatestRequest, int) bool) *ApiGetDeviceMonitoringLatestRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) Where(fn func(ApiGetDeviceMonitoringLatestRequest) bool) *ApiGetDeviceMonitoringLatestRequestStream {
	result := ApiGetDeviceMonitoringLatestRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceMonitoringLatestRequestStream) WhereSlim(fn func(ApiGetDeviceMonitoringLatestRequest) bool) *ApiGetDeviceMonitoringLatestRequestStream {
	result := ApiGetDeviceMonitoringLatestRequestStreamOf()
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
