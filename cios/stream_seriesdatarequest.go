/*
 * Collection utility of SeriesDataRequest Struct
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

type SeriesDataRequestStream []SeriesDataRequest

func SeriesDataRequestStreamOf(arg ...SeriesDataRequest) SeriesDataRequestStream {
	return arg
}
func SeriesDataRequestStreamFrom(arg []SeriesDataRequest) SeriesDataRequestStream {
	return arg
}
func CreateSeriesDataRequestStream(arg ...SeriesDataRequest) *SeriesDataRequestStream {
	tmp := SeriesDataRequestStreamOf(arg...)
	return &tmp
}
func GenerateSeriesDataRequestStream(arg []SeriesDataRequest) *SeriesDataRequestStream {
	tmp := SeriesDataRequestStreamFrom(arg)
	return &tmp
}

func (self *SeriesDataRequestStream) Add(arg SeriesDataRequest) *SeriesDataRequestStream {
	return self.AddAll(arg)
}
func (self *SeriesDataRequestStream) AddAll(arg ...SeriesDataRequest) *SeriesDataRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *SeriesDataRequestStream) AddSafe(arg *SeriesDataRequest) *SeriesDataRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *SeriesDataRequestStream) Aggregate(fn func(SeriesDataRequest, SeriesDataRequest) SeriesDataRequest) *SeriesDataRequestStream {
	result := SeriesDataRequestStreamOf()
	self.ForEach(func(v SeriesDataRequest, i int) {
		if i == 0 {
			result.Add(fn(SeriesDataRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *SeriesDataRequestStream) AllMatch(fn func(SeriesDataRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *SeriesDataRequestStream) AnyMatch(fn func(SeriesDataRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SeriesDataRequestStream) Clone() *SeriesDataRequestStream {
	temp := make([]SeriesDataRequest, self.Len())
	copy(temp, *self)
	return (*SeriesDataRequestStream)(&temp)
}
func (self *SeriesDataRequestStream) Copy() *SeriesDataRequestStream {
	return self.Clone()
}
func (self *SeriesDataRequestStream) Concat(arg []SeriesDataRequest) *SeriesDataRequestStream {
	return self.AddAll(arg...)
}
func (self *SeriesDataRequestStream) Contains(arg SeriesDataRequest) bool {
	return self.FindIndex(func(_arg SeriesDataRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *SeriesDataRequestStream) Clean() *SeriesDataRequestStream {
	*self = SeriesDataRequestStreamOf()
	return self
}
func (self *SeriesDataRequestStream) Delete(index int) *SeriesDataRequestStream {
	return self.DeleteRange(index, index)
}
func (self *SeriesDataRequestStream) DeleteRange(startIndex, endIndex int) *SeriesDataRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *SeriesDataRequestStream) Distinct() *SeriesDataRequestStream {
	caches := map[string]bool{}
	result := SeriesDataRequestStreamOf()
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
func (self *SeriesDataRequestStream) Each(fn func(SeriesDataRequest)) *SeriesDataRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *SeriesDataRequestStream) EachRight(fn func(SeriesDataRequest)) *SeriesDataRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *SeriesDataRequestStream) Equals(arr []SeriesDataRequest) bool {
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
func (self *SeriesDataRequestStream) Filter(fn func(SeriesDataRequest, int) bool) *SeriesDataRequestStream {
	result := SeriesDataRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SeriesDataRequestStream) FilterSlim(fn func(SeriesDataRequest, int) bool) *SeriesDataRequestStream {
	result := SeriesDataRequestStreamOf()
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
func (self *SeriesDataRequestStream) Find(fn func(SeriesDataRequest, int) bool) *SeriesDataRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *SeriesDataRequestStream) FindOr(fn func(SeriesDataRequest, int) bool, or SeriesDataRequest) SeriesDataRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *SeriesDataRequestStream) FindIndex(fn func(SeriesDataRequest, int) bool) int {
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
func (self *SeriesDataRequestStream) First() *SeriesDataRequest {
	return self.Get(0)
}
func (self *SeriesDataRequestStream) FirstOr(arg SeriesDataRequest) SeriesDataRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *SeriesDataRequestStream) ForEach(fn func(SeriesDataRequest, int)) *SeriesDataRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *SeriesDataRequestStream) ForEachRight(fn func(SeriesDataRequest, int)) *SeriesDataRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SeriesDataRequestStream) GroupBy(fn func(SeriesDataRequest, int) string) map[string][]SeriesDataRequest {
	m := map[string][]SeriesDataRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *SeriesDataRequestStream) GroupByValues(fn func(SeriesDataRequest, int) string) [][]SeriesDataRequest {
	var tmp [][]SeriesDataRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SeriesDataRequestStream) IndexOf(arg SeriesDataRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SeriesDataRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SeriesDataRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SeriesDataRequestStream) Last() *SeriesDataRequest {
	return self.Get(self.Len() - 1)
}
func (self *SeriesDataRequestStream) LastOr(arg SeriesDataRequest) SeriesDataRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *SeriesDataRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SeriesDataRequestStream) Limit(limit int) *SeriesDataRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *SeriesDataRequestStream) Map(fn func(SeriesDataRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SeriesDataRequestStream) Map2Int(fn func(SeriesDataRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SeriesDataRequestStream) Map2Int32(fn func(SeriesDataRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SeriesDataRequestStream) Map2Int64(fn func(SeriesDataRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SeriesDataRequestStream) Map2Float32(fn func(SeriesDataRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SeriesDataRequestStream) Map2Float64(fn func(SeriesDataRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SeriesDataRequestStream) Map2Bool(fn func(SeriesDataRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SeriesDataRequestStream) Map2Bytes(fn func(SeriesDataRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SeriesDataRequestStream) Map2String(fn func(SeriesDataRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SeriesDataRequestStream) Max(fn func(SeriesDataRequest, int) float64) *SeriesDataRequest {
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
func (self *SeriesDataRequestStream) Min(fn func(SeriesDataRequest, int) float64) *SeriesDataRequest {
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
func (self *SeriesDataRequestStream) NoneMatch(fn func(SeriesDataRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *SeriesDataRequestStream) Get(index int) *SeriesDataRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *SeriesDataRequestStream) GetOr(index int, arg SeriesDataRequest) SeriesDataRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *SeriesDataRequestStream) Peek(fn func(*SeriesDataRequest, int)) *SeriesDataRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *SeriesDataRequestStream) Reduce(fn func(SeriesDataRequest, SeriesDataRequest, int) SeriesDataRequest) *SeriesDataRequestStream {
	return self.ReduceInit(fn, SeriesDataRequest{})
}
func (self *SeriesDataRequestStream) ReduceInit(fn func(SeriesDataRequest, SeriesDataRequest, int) SeriesDataRequest, initialValue SeriesDataRequest) *SeriesDataRequestStream {
	result := SeriesDataRequestStreamOf()
	self.ForEach(func(v SeriesDataRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *SeriesDataRequestStream) ReduceInterface(fn func(interface{}, SeriesDataRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(SeriesDataRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SeriesDataRequestStream) ReduceString(fn func(string, SeriesDataRequest, int) string) []string {
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
func (self *SeriesDataRequestStream) ReduceInt(fn func(int, SeriesDataRequest, int) int) []int {
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
func (self *SeriesDataRequestStream) ReduceInt32(fn func(int32, SeriesDataRequest, int) int32) []int32 {
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
func (self *SeriesDataRequestStream) ReduceInt64(fn func(int64, SeriesDataRequest, int) int64) []int64 {
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
func (self *SeriesDataRequestStream) ReduceFloat32(fn func(float32, SeriesDataRequest, int) float32) []float32 {
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
func (self *SeriesDataRequestStream) ReduceFloat64(fn func(float64, SeriesDataRequest, int) float64) []float64 {
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
func (self *SeriesDataRequestStream) ReduceBool(fn func(bool, SeriesDataRequest, int) bool) []bool {
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
func (self *SeriesDataRequestStream) Reverse() *SeriesDataRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *SeriesDataRequestStream) Replace(fn func(SeriesDataRequest, int) SeriesDataRequest) *SeriesDataRequestStream {
	return self.ForEach(func(v SeriesDataRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *SeriesDataRequestStream) Select(fn func(SeriesDataRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *SeriesDataRequestStream) Set(index int, val SeriesDataRequest) *SeriesDataRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *SeriesDataRequestStream) Skip(skip int) *SeriesDataRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *SeriesDataRequestStream) SkippingEach(fn func(SeriesDataRequest, int) int) *SeriesDataRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *SeriesDataRequestStream) Slice(startIndex, n int) *SeriesDataRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []SeriesDataRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *SeriesDataRequestStream) Sort(fn func(i, j int) bool) *SeriesDataRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SeriesDataRequestStream) Tail() *SeriesDataRequest {
	return self.Last()
}
func (self *SeriesDataRequestStream) TailOr(arg SeriesDataRequest) SeriesDataRequest {
	return self.LastOr(arg)
}
func (self *SeriesDataRequestStream) ToList() []SeriesDataRequest {
	return self.Val()
}
func (self *SeriesDataRequestStream) Unique() *SeriesDataRequestStream {
	return self.Distinct()
}
func (self *SeriesDataRequestStream) Val() []SeriesDataRequest {
	if self == nil {
		return []SeriesDataRequest{}
	}
	return *self.Copy()
}
func (self *SeriesDataRequestStream) While(fn func(SeriesDataRequest, int) bool) *SeriesDataRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *SeriesDataRequestStream) Where(fn func(SeriesDataRequest) bool) *SeriesDataRequestStream {
	result := SeriesDataRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SeriesDataRequestStream) WhereSlim(fn func(SeriesDataRequest) bool) *SeriesDataRequestStream {
	result := SeriesDataRequestStreamOf()
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
