/*
 * Collection utility of ApiGetContractUserRequest Struct
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

type ApiGetContractUserRequestStream []ApiGetContractUserRequest

func ApiGetContractUserRequestStreamOf(arg ...ApiGetContractUserRequest) ApiGetContractUserRequestStream {
	return arg
}
func ApiGetContractUserRequestStreamFrom(arg []ApiGetContractUserRequest) ApiGetContractUserRequestStream {
	return arg
}
func CreateApiGetContractUserRequestStream(arg ...ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	tmp := ApiGetContractUserRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetContractUserRequestStream(arg []ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	tmp := ApiGetContractUserRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetContractUserRequestStream) Add(arg ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetContractUserRequestStream) AddAll(arg ...ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetContractUserRequestStream) AddSafe(arg *ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetContractUserRequestStream) Aggregate(fn func(ApiGetContractUserRequest, ApiGetContractUserRequest) ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	result := ApiGetContractUserRequestStreamOf()
	self.ForEach(func(v ApiGetContractUserRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetContractUserRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetContractUserRequestStream) AllMatch(fn func(ApiGetContractUserRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetContractUserRequestStream) AnyMatch(fn func(ApiGetContractUserRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetContractUserRequestStream) Clone() *ApiGetContractUserRequestStream {
	temp := make([]ApiGetContractUserRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetContractUserRequestStream)(&temp)
}
func (self *ApiGetContractUserRequestStream) Copy() *ApiGetContractUserRequestStream {
	return self.Clone()
}
func (self *ApiGetContractUserRequestStream) Concat(arg []ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetContractUserRequestStream) Contains(arg ApiGetContractUserRequest) bool {
	return self.FindIndex(func(_arg ApiGetContractUserRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetContractUserRequestStream) Clean() *ApiGetContractUserRequestStream {
	*self = ApiGetContractUserRequestStreamOf()
	return self
}
func (self *ApiGetContractUserRequestStream) Delete(index int) *ApiGetContractUserRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetContractUserRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetContractUserRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetContractUserRequestStream) Distinct() *ApiGetContractUserRequestStream {
	caches := map[string]bool{}
	result := ApiGetContractUserRequestStreamOf()
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
func (self *ApiGetContractUserRequestStream) Each(fn func(ApiGetContractUserRequest)) *ApiGetContractUserRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetContractUserRequestStream) EachRight(fn func(ApiGetContractUserRequest)) *ApiGetContractUserRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetContractUserRequestStream) Equals(arr []ApiGetContractUserRequest) bool {
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
func (self *ApiGetContractUserRequestStream) Filter(fn func(ApiGetContractUserRequest, int) bool) *ApiGetContractUserRequestStream {
	result := ApiGetContractUserRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetContractUserRequestStream) FilterSlim(fn func(ApiGetContractUserRequest, int) bool) *ApiGetContractUserRequestStream {
	result := ApiGetContractUserRequestStreamOf()
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
func (self *ApiGetContractUserRequestStream) Find(fn func(ApiGetContractUserRequest, int) bool) *ApiGetContractUserRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetContractUserRequestStream) FindOr(fn func(ApiGetContractUserRequest, int) bool, or ApiGetContractUserRequest) ApiGetContractUserRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetContractUserRequestStream) FindIndex(fn func(ApiGetContractUserRequest, int) bool) int {
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
func (self *ApiGetContractUserRequestStream) First() *ApiGetContractUserRequest {
	return self.Get(0)
}
func (self *ApiGetContractUserRequestStream) FirstOr(arg ApiGetContractUserRequest) ApiGetContractUserRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetContractUserRequestStream) ForEach(fn func(ApiGetContractUserRequest, int)) *ApiGetContractUserRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetContractUserRequestStream) ForEachRight(fn func(ApiGetContractUserRequest, int)) *ApiGetContractUserRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetContractUserRequestStream) GroupBy(fn func(ApiGetContractUserRequest, int) string) map[string][]ApiGetContractUserRequest {
	m := map[string][]ApiGetContractUserRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetContractUserRequestStream) GroupByValues(fn func(ApiGetContractUserRequest, int) string) [][]ApiGetContractUserRequest {
	var tmp [][]ApiGetContractUserRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetContractUserRequestStream) IndexOf(arg ApiGetContractUserRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetContractUserRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetContractUserRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetContractUserRequestStream) Last() *ApiGetContractUserRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetContractUserRequestStream) LastOr(arg ApiGetContractUserRequest) ApiGetContractUserRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetContractUserRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetContractUserRequestStream) Limit(limit int) *ApiGetContractUserRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetContractUserRequestStream) Map(fn func(ApiGetContractUserRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractUserRequestStream) Map2Int(fn func(ApiGetContractUserRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractUserRequestStream) Map2Int32(fn func(ApiGetContractUserRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractUserRequestStream) Map2Int64(fn func(ApiGetContractUserRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractUserRequestStream) Map2Float32(fn func(ApiGetContractUserRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractUserRequestStream) Map2Float64(fn func(ApiGetContractUserRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractUserRequestStream) Map2Bool(fn func(ApiGetContractUserRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractUserRequestStream) Map2Bytes(fn func(ApiGetContractUserRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractUserRequestStream) Map2String(fn func(ApiGetContractUserRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractUserRequestStream) Max(fn func(ApiGetContractUserRequest, int) float64) *ApiGetContractUserRequest {
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
func (self *ApiGetContractUserRequestStream) Min(fn func(ApiGetContractUserRequest, int) float64) *ApiGetContractUserRequest {
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
func (self *ApiGetContractUserRequestStream) NoneMatch(fn func(ApiGetContractUserRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetContractUserRequestStream) Get(index int) *ApiGetContractUserRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetContractUserRequestStream) GetOr(index int, arg ApiGetContractUserRequest) ApiGetContractUserRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetContractUserRequestStream) Peek(fn func(*ApiGetContractUserRequest, int)) *ApiGetContractUserRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *ApiGetContractUserRequestStream) Reduce(fn func(ApiGetContractUserRequest, ApiGetContractUserRequest, int) ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	return self.ReduceInit(fn, ApiGetContractUserRequest{})
}
func (self *ApiGetContractUserRequestStream) ReduceInit(fn func(ApiGetContractUserRequest, ApiGetContractUserRequest, int) ApiGetContractUserRequest, initialValue ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	result := ApiGetContractUserRequestStreamOf()
	self.ForEach(func(v ApiGetContractUserRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetContractUserRequestStream) ReduceInterface(fn func(interface{}, ApiGetContractUserRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetContractUserRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetContractUserRequestStream) ReduceString(fn func(string, ApiGetContractUserRequest, int) string) []string {
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
func (self *ApiGetContractUserRequestStream) ReduceInt(fn func(int, ApiGetContractUserRequest, int) int) []int {
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
func (self *ApiGetContractUserRequestStream) ReduceInt32(fn func(int32, ApiGetContractUserRequest, int) int32) []int32 {
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
func (self *ApiGetContractUserRequestStream) ReduceInt64(fn func(int64, ApiGetContractUserRequest, int) int64) []int64 {
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
func (self *ApiGetContractUserRequestStream) ReduceFloat32(fn func(float32, ApiGetContractUserRequest, int) float32) []float32 {
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
func (self *ApiGetContractUserRequestStream) ReduceFloat64(fn func(float64, ApiGetContractUserRequest, int) float64) []float64 {
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
func (self *ApiGetContractUserRequestStream) ReduceBool(fn func(bool, ApiGetContractUserRequest, int) bool) []bool {
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
func (self *ApiGetContractUserRequestStream) Reverse() *ApiGetContractUserRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetContractUserRequestStream) Replace(fn func(ApiGetContractUserRequest, int) ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	return self.ForEach(func(v ApiGetContractUserRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetContractUserRequestStream) Select(fn func(ApiGetContractUserRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetContractUserRequestStream) Set(index int, val ApiGetContractUserRequest) *ApiGetContractUserRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetContractUserRequestStream) Skip(skip int) *ApiGetContractUserRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetContractUserRequestStream) SkippingEach(fn func(ApiGetContractUserRequest, int) int) *ApiGetContractUserRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetContractUserRequestStream) Slice(startIndex, n int) *ApiGetContractUserRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetContractUserRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetContractUserRequestStream) Sort(fn func(i, j int) bool) *ApiGetContractUserRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetContractUserRequestStream) Tail() *ApiGetContractUserRequest {
	return self.Last()
}
func (self *ApiGetContractUserRequestStream) TailOr(arg ApiGetContractUserRequest) ApiGetContractUserRequest {
	return self.LastOr(arg)
}
func (self *ApiGetContractUserRequestStream) ToList() []ApiGetContractUserRequest {
	return self.Val()
}
func (self *ApiGetContractUserRequestStream) Unique() *ApiGetContractUserRequestStream {
	return self.Distinct()
}
func (self *ApiGetContractUserRequestStream) Val() []ApiGetContractUserRequest {
	if self == nil {
		return []ApiGetContractUserRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetContractUserRequestStream) While(fn func(ApiGetContractUserRequest, int) bool) *ApiGetContractUserRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetContractUserRequestStream) Where(fn func(ApiGetContractUserRequest) bool) *ApiGetContractUserRequestStream {
	result := ApiGetContractUserRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetContractUserRequestStream) WhereSlim(fn func(ApiGetContractUserRequest) bool) *ApiGetContractUserRequestStream {
	result := ApiGetContractUserRequestStreamOf()
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
