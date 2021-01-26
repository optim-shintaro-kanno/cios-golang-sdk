/*
 * Collection utility of ConnectTokenResponse Struct
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

type ConnectTokenResponseStream []ConnectTokenResponse

func ConnectTokenResponseStreamOf(arg ...ConnectTokenResponse) ConnectTokenResponseStream {
	return arg
}
func ConnectTokenResponseStreamFrom(arg []ConnectTokenResponse) ConnectTokenResponseStream {
	return arg
}
func CreateConnectTokenResponseStream(arg ...ConnectTokenResponse) *ConnectTokenResponseStream {
	tmp := ConnectTokenResponseStreamOf(arg...)
	return &tmp
}
func GenerateConnectTokenResponseStream(arg []ConnectTokenResponse) *ConnectTokenResponseStream {
	tmp := ConnectTokenResponseStreamFrom(arg)
	return &tmp
}

func (self *ConnectTokenResponseStream) Add(arg ConnectTokenResponse) *ConnectTokenResponseStream {
	return self.AddAll(arg)
}
func (self *ConnectTokenResponseStream) AddAll(arg ...ConnectTokenResponse) *ConnectTokenResponseStream {
	*self = append(*self, arg...)
	return self
}
func (self *ConnectTokenResponseStream) AddSafe(arg *ConnectTokenResponse) *ConnectTokenResponseStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ConnectTokenResponseStream) Aggregate(fn func(ConnectTokenResponse, ConnectTokenResponse) ConnectTokenResponse) *ConnectTokenResponseStream {
	result := ConnectTokenResponseStreamOf()
	self.ForEach(func(v ConnectTokenResponse, i int) {
		if i == 0 {
			result.Add(fn(ConnectTokenResponse{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ConnectTokenResponseStream) AllMatch(fn func(ConnectTokenResponse, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ConnectTokenResponseStream) AnyMatch(fn func(ConnectTokenResponse, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ConnectTokenResponseStream) Clone() *ConnectTokenResponseStream {
	temp := make([]ConnectTokenResponse, self.Len())
	copy(temp, *self)
	return (*ConnectTokenResponseStream)(&temp)
}
func (self *ConnectTokenResponseStream) Copy() *ConnectTokenResponseStream {
	return self.Clone()
}
func (self *ConnectTokenResponseStream) Concat(arg []ConnectTokenResponse) *ConnectTokenResponseStream {
	return self.AddAll(arg...)
}
func (self *ConnectTokenResponseStream) Contains(arg ConnectTokenResponse) bool {
	return self.FindIndex(func(_arg ConnectTokenResponse, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ConnectTokenResponseStream) Clean() *ConnectTokenResponseStream {
	*self = ConnectTokenResponseStreamOf()
	return self
}
func (self *ConnectTokenResponseStream) Delete(index int) *ConnectTokenResponseStream {
	return self.DeleteRange(index, index)
}
func (self *ConnectTokenResponseStream) DeleteRange(startIndex, endIndex int) *ConnectTokenResponseStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ConnectTokenResponseStream) Distinct() *ConnectTokenResponseStream {
	caches := map[string]bool{}
	result := ConnectTokenResponseStreamOf()
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
func (self *ConnectTokenResponseStream) Each(fn func(ConnectTokenResponse)) *ConnectTokenResponseStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ConnectTokenResponseStream) EachRight(fn func(ConnectTokenResponse)) *ConnectTokenResponseStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ConnectTokenResponseStream) Equals(arr []ConnectTokenResponse) bool {
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
func (self *ConnectTokenResponseStream) Filter(fn func(ConnectTokenResponse, int) bool) *ConnectTokenResponseStream {
	result := ConnectTokenResponseStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ConnectTokenResponseStream) FilterSlim(fn func(ConnectTokenResponse, int) bool) *ConnectTokenResponseStream {
	result := ConnectTokenResponseStreamOf()
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
func (self *ConnectTokenResponseStream) Find(fn func(ConnectTokenResponse, int) bool) *ConnectTokenResponse {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ConnectTokenResponseStream) FindOr(fn func(ConnectTokenResponse, int) bool, or ConnectTokenResponse) ConnectTokenResponse {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ConnectTokenResponseStream) FindIndex(fn func(ConnectTokenResponse, int) bool) int {
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
func (self *ConnectTokenResponseStream) First() *ConnectTokenResponse {
	return self.Get(0)
}
func (self *ConnectTokenResponseStream) FirstOr(arg ConnectTokenResponse) ConnectTokenResponse {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ConnectTokenResponseStream) ForEach(fn func(ConnectTokenResponse, int)) *ConnectTokenResponseStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ConnectTokenResponseStream) ForEachRight(fn func(ConnectTokenResponse, int)) *ConnectTokenResponseStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ConnectTokenResponseStream) GroupBy(fn func(ConnectTokenResponse, int) string) map[string][]ConnectTokenResponse {
	m := map[string][]ConnectTokenResponse{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ConnectTokenResponseStream) GroupByValues(fn func(ConnectTokenResponse, int) string) [][]ConnectTokenResponse {
	var tmp [][]ConnectTokenResponse
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ConnectTokenResponseStream) IndexOf(arg ConnectTokenResponse) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ConnectTokenResponseStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ConnectTokenResponseStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ConnectTokenResponseStream) Last() *ConnectTokenResponse {
	return self.Get(self.Len() - 1)
}
func (self *ConnectTokenResponseStream) LastOr(arg ConnectTokenResponse) ConnectTokenResponse {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ConnectTokenResponseStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ConnectTokenResponseStream) Limit(limit int) *ConnectTokenResponseStream {
	self.Slice(0, limit)
	return self
}

func (self *ConnectTokenResponseStream) Map(fn func(ConnectTokenResponse, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ConnectTokenResponseStream) Map2Int(fn func(ConnectTokenResponse, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ConnectTokenResponseStream) Map2Int32(fn func(ConnectTokenResponse, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ConnectTokenResponseStream) Map2Int64(fn func(ConnectTokenResponse, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ConnectTokenResponseStream) Map2Float32(fn func(ConnectTokenResponse, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ConnectTokenResponseStream) Map2Float64(fn func(ConnectTokenResponse, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ConnectTokenResponseStream) Map2Bool(fn func(ConnectTokenResponse, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ConnectTokenResponseStream) Map2Bytes(fn func(ConnectTokenResponse, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ConnectTokenResponseStream) Map2String(fn func(ConnectTokenResponse, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ConnectTokenResponseStream) Max(fn func(ConnectTokenResponse, int) float64) *ConnectTokenResponse {
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
func (self *ConnectTokenResponseStream) Min(fn func(ConnectTokenResponse, int) float64) *ConnectTokenResponse {
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
func (self *ConnectTokenResponseStream) NoneMatch(fn func(ConnectTokenResponse, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ConnectTokenResponseStream) Get(index int) *ConnectTokenResponse {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ConnectTokenResponseStream) GetOr(index int, arg ConnectTokenResponse) ConnectTokenResponse {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ConnectTokenResponseStream) Peek(fn func(*ConnectTokenResponse, int)) *ConnectTokenResponseStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *ConnectTokenResponseStream) Reduce(fn func(ConnectTokenResponse, ConnectTokenResponse, int) ConnectTokenResponse) *ConnectTokenResponseStream {
	return self.ReduceInit(fn, ConnectTokenResponse{})
}
func (self *ConnectTokenResponseStream) ReduceInit(fn func(ConnectTokenResponse, ConnectTokenResponse, int) ConnectTokenResponse, initialValue ConnectTokenResponse) *ConnectTokenResponseStream {
	result := ConnectTokenResponseStreamOf()
	self.ForEach(func(v ConnectTokenResponse, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ConnectTokenResponseStream) ReduceInterface(fn func(interface{}, ConnectTokenResponse, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ConnectTokenResponse{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ConnectTokenResponseStream) ReduceString(fn func(string, ConnectTokenResponse, int) string) []string {
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
func (self *ConnectTokenResponseStream) ReduceInt(fn func(int, ConnectTokenResponse, int) int) []int {
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
func (self *ConnectTokenResponseStream) ReduceInt32(fn func(int32, ConnectTokenResponse, int) int32) []int32 {
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
func (self *ConnectTokenResponseStream) ReduceInt64(fn func(int64, ConnectTokenResponse, int) int64) []int64 {
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
func (self *ConnectTokenResponseStream) ReduceFloat32(fn func(float32, ConnectTokenResponse, int) float32) []float32 {
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
func (self *ConnectTokenResponseStream) ReduceFloat64(fn func(float64, ConnectTokenResponse, int) float64) []float64 {
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
func (self *ConnectTokenResponseStream) ReduceBool(fn func(bool, ConnectTokenResponse, int) bool) []bool {
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
func (self *ConnectTokenResponseStream) Reverse() *ConnectTokenResponseStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ConnectTokenResponseStream) Replace(fn func(ConnectTokenResponse, int) ConnectTokenResponse) *ConnectTokenResponseStream {
	return self.ForEach(func(v ConnectTokenResponse, i int) { self.Set(i, fn(v, i)) })
}
func (self *ConnectTokenResponseStream) Select(fn func(ConnectTokenResponse) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ConnectTokenResponseStream) Set(index int, val ConnectTokenResponse) *ConnectTokenResponseStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ConnectTokenResponseStream) Skip(skip int) *ConnectTokenResponseStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ConnectTokenResponseStream) SkippingEach(fn func(ConnectTokenResponse, int) int) *ConnectTokenResponseStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ConnectTokenResponseStream) Slice(startIndex, n int) *ConnectTokenResponseStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ConnectTokenResponse{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ConnectTokenResponseStream) Sort(fn func(i, j int) bool) *ConnectTokenResponseStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ConnectTokenResponseStream) Tail() *ConnectTokenResponse {
	return self.Last()
}
func (self *ConnectTokenResponseStream) TailOr(arg ConnectTokenResponse) ConnectTokenResponse {
	return self.LastOr(arg)
}
func (self *ConnectTokenResponseStream) ToList() []ConnectTokenResponse {
	return self.Val()
}
func (self *ConnectTokenResponseStream) Unique() *ConnectTokenResponseStream {
	return self.Distinct()
}
func (self *ConnectTokenResponseStream) Val() []ConnectTokenResponse {
	if self == nil {
		return []ConnectTokenResponse{}
	}
	return *self.Copy()
}
func (self *ConnectTokenResponseStream) While(fn func(ConnectTokenResponse, int) bool) *ConnectTokenResponseStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ConnectTokenResponseStream) Where(fn func(ConnectTokenResponse) bool) *ConnectTokenResponseStream {
	result := ConnectTokenResponseStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ConnectTokenResponseStream) WhereSlim(fn func(ConnectTokenResponse) bool) *ConnectTokenResponseStream {
	result := ConnectTokenResponseStreamOf()
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
