/*
 * Collection utility of ApiUpdatePointRequest Struct
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

type ApiUpdatePointRequestStream []ApiUpdatePointRequest

func ApiUpdatePointRequestStreamOf(arg ...ApiUpdatePointRequest) ApiUpdatePointRequestStream {
	return arg
}
func ApiUpdatePointRequestStreamFrom(arg []ApiUpdatePointRequest) ApiUpdatePointRequestStream {
	return arg
}
func CreateApiUpdatePointRequestStream(arg ...ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	tmp := ApiUpdatePointRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiUpdatePointRequestStream(arg []ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	tmp := ApiUpdatePointRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiUpdatePointRequestStream) Add(arg ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	return self.AddAll(arg)
}
func (self *ApiUpdatePointRequestStream) AddAll(arg ...ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiUpdatePointRequestStream) AddSafe(arg *ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiUpdatePointRequestStream) Aggregate(fn func(ApiUpdatePointRequest, ApiUpdatePointRequest) ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	result := ApiUpdatePointRequestStreamOf()
	self.ForEach(func(v ApiUpdatePointRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiUpdatePointRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiUpdatePointRequestStream) AllMatch(fn func(ApiUpdatePointRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiUpdatePointRequestStream) AnyMatch(fn func(ApiUpdatePointRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiUpdatePointRequestStream) Clone() *ApiUpdatePointRequestStream {
	temp := make([]ApiUpdatePointRequest, self.Len())
	copy(temp, *self)
	return (*ApiUpdatePointRequestStream)(&temp)
}
func (self *ApiUpdatePointRequestStream) Copy() *ApiUpdatePointRequestStream {
	return self.Clone()
}
func (self *ApiUpdatePointRequestStream) Concat(arg []ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiUpdatePointRequestStream) Contains(arg ApiUpdatePointRequest) bool {
	return self.FindIndex(func(_arg ApiUpdatePointRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiUpdatePointRequestStream) Clean() *ApiUpdatePointRequestStream {
	*self = ApiUpdatePointRequestStreamOf()
	return self
}
func (self *ApiUpdatePointRequestStream) Delete(index int) *ApiUpdatePointRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiUpdatePointRequestStream) DeleteRange(startIndex, endIndex int) *ApiUpdatePointRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiUpdatePointRequestStream) Distinct() *ApiUpdatePointRequestStream {
	caches := map[string]bool{}
	result := ApiUpdatePointRequestStreamOf()
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
func (self *ApiUpdatePointRequestStream) Each(fn func(ApiUpdatePointRequest)) *ApiUpdatePointRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiUpdatePointRequestStream) EachRight(fn func(ApiUpdatePointRequest)) *ApiUpdatePointRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiUpdatePointRequestStream) Equals(arr []ApiUpdatePointRequest) bool {
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
func (self *ApiUpdatePointRequestStream) Filter(fn func(ApiUpdatePointRequest, int) bool) *ApiUpdatePointRequestStream {
	result := ApiUpdatePointRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiUpdatePointRequestStream) FilterSlim(fn func(ApiUpdatePointRequest, int) bool) *ApiUpdatePointRequestStream {
	result := ApiUpdatePointRequestStreamOf()
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
func (self *ApiUpdatePointRequestStream) Find(fn func(ApiUpdatePointRequest, int) bool) *ApiUpdatePointRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiUpdatePointRequestStream) FindOr(fn func(ApiUpdatePointRequest, int) bool, or ApiUpdatePointRequest) ApiUpdatePointRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiUpdatePointRequestStream) FindIndex(fn func(ApiUpdatePointRequest, int) bool) int {
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
func (self *ApiUpdatePointRequestStream) First() *ApiUpdatePointRequest {
	return self.Get(0)
}
func (self *ApiUpdatePointRequestStream) FirstOr(arg ApiUpdatePointRequest) ApiUpdatePointRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdatePointRequestStream) ForEach(fn func(ApiUpdatePointRequest, int)) *ApiUpdatePointRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiUpdatePointRequestStream) ForEachRight(fn func(ApiUpdatePointRequest, int)) *ApiUpdatePointRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiUpdatePointRequestStream) GroupBy(fn func(ApiUpdatePointRequest, int) string) map[string][]ApiUpdatePointRequest {
	m := map[string][]ApiUpdatePointRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiUpdatePointRequestStream) GroupByValues(fn func(ApiUpdatePointRequest, int) string) [][]ApiUpdatePointRequest {
	var tmp [][]ApiUpdatePointRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiUpdatePointRequestStream) IndexOf(arg ApiUpdatePointRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiUpdatePointRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiUpdatePointRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiUpdatePointRequestStream) Last() *ApiUpdatePointRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiUpdatePointRequestStream) LastOr(arg ApiUpdatePointRequest) ApiUpdatePointRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdatePointRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiUpdatePointRequestStream) Limit(limit int) *ApiUpdatePointRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiUpdatePointRequestStream) Map(fn func(ApiUpdatePointRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdatePointRequestStream) Map2Int(fn func(ApiUpdatePointRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdatePointRequestStream) Map2Int32(fn func(ApiUpdatePointRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdatePointRequestStream) Map2Int64(fn func(ApiUpdatePointRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdatePointRequestStream) Map2Float32(fn func(ApiUpdatePointRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdatePointRequestStream) Map2Float64(fn func(ApiUpdatePointRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdatePointRequestStream) Map2Bool(fn func(ApiUpdatePointRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdatePointRequestStream) Map2Bytes(fn func(ApiUpdatePointRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdatePointRequestStream) Map2String(fn func(ApiUpdatePointRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdatePointRequestStream) Max(fn func(ApiUpdatePointRequest, int) float64) *ApiUpdatePointRequest {
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
func (self *ApiUpdatePointRequestStream) Min(fn func(ApiUpdatePointRequest, int) float64) *ApiUpdatePointRequest {
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
func (self *ApiUpdatePointRequestStream) NoneMatch(fn func(ApiUpdatePointRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiUpdatePointRequestStream) Get(index int) *ApiUpdatePointRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiUpdatePointRequestStream) GetOr(index int, arg ApiUpdatePointRequest) ApiUpdatePointRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdatePointRequestStream) Peek(fn func(*ApiUpdatePointRequest, int)) *ApiUpdatePointRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *ApiUpdatePointRequestStream) Reduce(fn func(ApiUpdatePointRequest, ApiUpdatePointRequest, int) ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	return self.ReduceInit(fn, ApiUpdatePointRequest{})
}
func (self *ApiUpdatePointRequestStream) ReduceInit(fn func(ApiUpdatePointRequest, ApiUpdatePointRequest, int) ApiUpdatePointRequest, initialValue ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	result := ApiUpdatePointRequestStreamOf()
	self.ForEach(func(v ApiUpdatePointRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiUpdatePointRequestStream) ReduceInterface(fn func(interface{}, ApiUpdatePointRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiUpdatePointRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiUpdatePointRequestStream) ReduceString(fn func(string, ApiUpdatePointRequest, int) string) []string {
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
func (self *ApiUpdatePointRequestStream) ReduceInt(fn func(int, ApiUpdatePointRequest, int) int) []int {
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
func (self *ApiUpdatePointRequestStream) ReduceInt32(fn func(int32, ApiUpdatePointRequest, int) int32) []int32 {
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
func (self *ApiUpdatePointRequestStream) ReduceInt64(fn func(int64, ApiUpdatePointRequest, int) int64) []int64 {
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
func (self *ApiUpdatePointRequestStream) ReduceFloat32(fn func(float32, ApiUpdatePointRequest, int) float32) []float32 {
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
func (self *ApiUpdatePointRequestStream) ReduceFloat64(fn func(float64, ApiUpdatePointRequest, int) float64) []float64 {
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
func (self *ApiUpdatePointRequestStream) ReduceBool(fn func(bool, ApiUpdatePointRequest, int) bool) []bool {
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
func (self *ApiUpdatePointRequestStream) Reverse() *ApiUpdatePointRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiUpdatePointRequestStream) Replace(fn func(ApiUpdatePointRequest, int) ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	return self.ForEach(func(v ApiUpdatePointRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiUpdatePointRequestStream) Select(fn func(ApiUpdatePointRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiUpdatePointRequestStream) Set(index int, val ApiUpdatePointRequest) *ApiUpdatePointRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiUpdatePointRequestStream) Skip(skip int) *ApiUpdatePointRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiUpdatePointRequestStream) SkippingEach(fn func(ApiUpdatePointRequest, int) int) *ApiUpdatePointRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiUpdatePointRequestStream) Slice(startIndex, n int) *ApiUpdatePointRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiUpdatePointRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiUpdatePointRequestStream) Sort(fn func(i, j int) bool) *ApiUpdatePointRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiUpdatePointRequestStream) Tail() *ApiUpdatePointRequest {
	return self.Last()
}
func (self *ApiUpdatePointRequestStream) TailOr(arg ApiUpdatePointRequest) ApiUpdatePointRequest {
	return self.LastOr(arg)
}
func (self *ApiUpdatePointRequestStream) ToList() []ApiUpdatePointRequest {
	return self.Val()
}
func (self *ApiUpdatePointRequestStream) Unique() *ApiUpdatePointRequestStream {
	return self.Distinct()
}
func (self *ApiUpdatePointRequestStream) Val() []ApiUpdatePointRequest {
	if self == nil {
		return []ApiUpdatePointRequest{}
	}
	return *self.Copy()
}
func (self *ApiUpdatePointRequestStream) While(fn func(ApiUpdatePointRequest, int) bool) *ApiUpdatePointRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiUpdatePointRequestStream) Where(fn func(ApiUpdatePointRequest) bool) *ApiUpdatePointRequestStream {
	result := ApiUpdatePointRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiUpdatePointRequestStream) WhereSlim(fn func(ApiUpdatePointRequest) bool) *ApiUpdatePointRequestStream {
	result := ApiUpdatePointRequestStreamOf()
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
