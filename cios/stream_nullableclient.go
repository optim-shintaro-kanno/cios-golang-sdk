/*
 * Collection utility of NullableClient Struct
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

type NullableClientStream []NullableClient

func NullableClientStreamOf(arg ...NullableClient) NullableClientStream {
	return arg
}
func NullableClientStreamFrom(arg []NullableClient) NullableClientStream {
	return arg
}
func CreateNullableClientStream(arg ...NullableClient) *NullableClientStream {
	tmp := NullableClientStreamOf(arg...)
	return &tmp
}
func GenerateNullableClientStream(arg []NullableClient) *NullableClientStream {
	tmp := NullableClientStreamFrom(arg)
	return &tmp
}

func (self *NullableClientStream) Add(arg NullableClient) *NullableClientStream {
	return self.AddAll(arg)
}
func (self *NullableClientStream) AddAll(arg ...NullableClient) *NullableClientStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableClientStream) AddSafe(arg *NullableClient) *NullableClientStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableClientStream) Aggregate(fn func(NullableClient, NullableClient) NullableClient) *NullableClientStream {
	result := NullableClientStreamOf()
	self.ForEach(func(v NullableClient, i int) {
		if i == 0 {
			result.Add(fn(NullableClient{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableClientStream) AllMatch(fn func(NullableClient, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableClientStream) AnyMatch(fn func(NullableClient, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableClientStream) Clone() *NullableClientStream {
	temp := make([]NullableClient, self.Len())
	copy(temp, *self)
	return (*NullableClientStream)(&temp)
}
func (self *NullableClientStream) Copy() *NullableClientStream {
	return self.Clone()
}
func (self *NullableClientStream) Concat(arg []NullableClient) *NullableClientStream {
	return self.AddAll(arg...)
}
func (self *NullableClientStream) Contains(arg NullableClient) bool {
	return self.FindIndex(func(_arg NullableClient, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableClientStream) Clean() *NullableClientStream {
	*self = NullableClientStreamOf()
	return self
}
func (self *NullableClientStream) Delete(index int) *NullableClientStream {
	return self.DeleteRange(index, index)
}
func (self *NullableClientStream) DeleteRange(startIndex, endIndex int) *NullableClientStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableClientStream) Distinct() *NullableClientStream {
	caches := map[string]bool{}
	result := NullableClientStreamOf()
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
func (self *NullableClientStream) Each(fn func(NullableClient)) *NullableClientStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableClientStream) EachRight(fn func(NullableClient)) *NullableClientStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableClientStream) Equals(arr []NullableClient) bool {
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
func (self *NullableClientStream) Filter(fn func(NullableClient, int) bool) *NullableClientStream {
	result := NullableClientStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableClientStream) FilterSlim(fn func(NullableClient, int) bool) *NullableClientStream {
	result := NullableClientStreamOf()
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
func (self *NullableClientStream) Find(fn func(NullableClient, int) bool) *NullableClient {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableClientStream) FindOr(fn func(NullableClient, int) bool, or NullableClient) NullableClient {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableClientStream) FindIndex(fn func(NullableClient, int) bool) int {
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
func (self *NullableClientStream) First() *NullableClient {
	return self.Get(0)
}
func (self *NullableClientStream) FirstOr(arg NullableClient) NullableClient {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableClientStream) ForEach(fn func(NullableClient, int)) *NullableClientStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableClientStream) ForEachRight(fn func(NullableClient, int)) *NullableClientStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableClientStream) GroupBy(fn func(NullableClient, int) string) map[string][]NullableClient {
	m := map[string][]NullableClient{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableClientStream) GroupByValues(fn func(NullableClient, int) string) [][]NullableClient {
	var tmp [][]NullableClient
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableClientStream) IndexOf(arg NullableClient) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableClientStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableClientStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableClientStream) Last() *NullableClient {
	return self.Get(self.Len() - 1)
}
func (self *NullableClientStream) LastOr(arg NullableClient) NullableClient {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableClientStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableClientStream) Limit(limit int) *NullableClientStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableClientStream) Map(fn func(NullableClient, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableClientStream) Map2Int(fn func(NullableClient, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableClientStream) Map2Int32(fn func(NullableClient, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableClientStream) Map2Int64(fn func(NullableClient, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableClientStream) Map2Float32(fn func(NullableClient, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableClientStream) Map2Float64(fn func(NullableClient, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableClientStream) Map2Bool(fn func(NullableClient, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableClientStream) Map2Bytes(fn func(NullableClient, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableClientStream) Map2String(fn func(NullableClient, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableClientStream) Max(fn func(NullableClient, int) float64) *NullableClient {
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
func (self *NullableClientStream) Min(fn func(NullableClient, int) float64) *NullableClient {
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
func (self *NullableClientStream) NoneMatch(fn func(NullableClient, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableClientStream) Get(index int) *NullableClient {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableClientStream) GetOr(index int, arg NullableClient) NullableClient {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableClientStream) Peek(fn func(*NullableClient, int)) *NullableClientStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableClientStream) Reduce(fn func(NullableClient, NullableClient, int) NullableClient) *NullableClientStream {
	return self.ReduceInit(fn, NullableClient{})
}
func (self *NullableClientStream) ReduceInit(fn func(NullableClient, NullableClient, int) NullableClient, initialValue NullableClient) *NullableClientStream {
	result := NullableClientStreamOf()
	self.ForEach(func(v NullableClient, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableClientStream) ReduceInterface(fn func(interface{}, NullableClient, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableClient{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableClientStream) ReduceString(fn func(string, NullableClient, int) string) []string {
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
func (self *NullableClientStream) ReduceInt(fn func(int, NullableClient, int) int) []int {
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
func (self *NullableClientStream) ReduceInt32(fn func(int32, NullableClient, int) int32) []int32 {
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
func (self *NullableClientStream) ReduceInt64(fn func(int64, NullableClient, int) int64) []int64 {
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
func (self *NullableClientStream) ReduceFloat32(fn func(float32, NullableClient, int) float32) []float32 {
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
func (self *NullableClientStream) ReduceFloat64(fn func(float64, NullableClient, int) float64) []float64 {
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
func (self *NullableClientStream) ReduceBool(fn func(bool, NullableClient, int) bool) []bool {
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
func (self *NullableClientStream) Reverse() *NullableClientStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableClientStream) Replace(fn func(NullableClient, int) NullableClient) *NullableClientStream {
	return self.ForEach(func(v NullableClient, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableClientStream) Select(fn func(NullableClient) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableClientStream) Set(index int, val NullableClient) *NullableClientStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableClientStream) Skip(skip int) *NullableClientStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableClientStream) SkippingEach(fn func(NullableClient, int) int) *NullableClientStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableClientStream) Slice(startIndex, n int) *NullableClientStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableClient{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableClientStream) Sort(fn func(i, j int) bool) *NullableClientStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableClientStream) Tail() *NullableClient {
	return self.Last()
}
func (self *NullableClientStream) TailOr(arg NullableClient) NullableClient {
	return self.LastOr(arg)
}
func (self *NullableClientStream) ToList() []NullableClient {
	return self.Val()
}
func (self *NullableClientStream) Unique() *NullableClientStream {
	return self.Distinct()
}
func (self *NullableClientStream) Val() []NullableClient {
	if self == nil {
		return []NullableClient{}
	}
	return *self.Copy()
}
func (self *NullableClientStream) While(fn func(NullableClient, int) bool) *NullableClientStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableClientStream) Where(fn func(NullableClient) bool) *NullableClientStream {
	result := NullableClientStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableClientStream) WhereSlim(fn func(NullableClient) bool) *NullableClientStream {
	result := NullableClientStreamOf()
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
