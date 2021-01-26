/*
 * Collection utility of NullableMemberInfo Struct
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

type NullableMemberInfoStream []NullableMemberInfo

func NullableMemberInfoStreamOf(arg ...NullableMemberInfo) NullableMemberInfoStream {
	return arg
}
func NullableMemberInfoStreamFrom(arg []NullableMemberInfo) NullableMemberInfoStream {
	return arg
}
func CreateNullableMemberInfoStream(arg ...NullableMemberInfo) *NullableMemberInfoStream {
	tmp := NullableMemberInfoStreamOf(arg...)
	return &tmp
}
func GenerateNullableMemberInfoStream(arg []NullableMemberInfo) *NullableMemberInfoStream {
	tmp := NullableMemberInfoStreamFrom(arg)
	return &tmp
}

func (self *NullableMemberInfoStream) Add(arg NullableMemberInfo) *NullableMemberInfoStream {
	return self.AddAll(arg)
}
func (self *NullableMemberInfoStream) AddAll(arg ...NullableMemberInfo) *NullableMemberInfoStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMemberInfoStream) AddSafe(arg *NullableMemberInfo) *NullableMemberInfoStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMemberInfoStream) Aggregate(fn func(NullableMemberInfo, NullableMemberInfo) NullableMemberInfo) *NullableMemberInfoStream {
	result := NullableMemberInfoStreamOf()
	self.ForEach(func(v NullableMemberInfo, i int) {
		if i == 0 {
			result.Add(fn(NullableMemberInfo{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMemberInfoStream) AllMatch(fn func(NullableMemberInfo, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMemberInfoStream) AnyMatch(fn func(NullableMemberInfo, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMemberInfoStream) Clone() *NullableMemberInfoStream {
	temp := make([]NullableMemberInfo, self.Len())
	copy(temp, *self)
	return (*NullableMemberInfoStream)(&temp)
}
func (self *NullableMemberInfoStream) Copy() *NullableMemberInfoStream {
	return self.Clone()
}
func (self *NullableMemberInfoStream) Concat(arg []NullableMemberInfo) *NullableMemberInfoStream {
	return self.AddAll(arg...)
}
func (self *NullableMemberInfoStream) Contains(arg NullableMemberInfo) bool {
	return self.FindIndex(func(_arg NullableMemberInfo, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMemberInfoStream) Clean() *NullableMemberInfoStream {
	*self = NullableMemberInfoStreamOf()
	return self
}
func (self *NullableMemberInfoStream) Delete(index int) *NullableMemberInfoStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMemberInfoStream) DeleteRange(startIndex, endIndex int) *NullableMemberInfoStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMemberInfoStream) Distinct() *NullableMemberInfoStream {
	caches := map[string]bool{}
	result := NullableMemberInfoStreamOf()
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
func (self *NullableMemberInfoStream) Each(fn func(NullableMemberInfo)) *NullableMemberInfoStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMemberInfoStream) EachRight(fn func(NullableMemberInfo)) *NullableMemberInfoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMemberInfoStream) Equals(arr []NullableMemberInfo) bool {
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
func (self *NullableMemberInfoStream) Filter(fn func(NullableMemberInfo, int) bool) *NullableMemberInfoStream {
	result := NullableMemberInfoStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMemberInfoStream) FilterSlim(fn func(NullableMemberInfo, int) bool) *NullableMemberInfoStream {
	result := NullableMemberInfoStreamOf()
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
func (self *NullableMemberInfoStream) Find(fn func(NullableMemberInfo, int) bool) *NullableMemberInfo {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMemberInfoStream) FindOr(fn func(NullableMemberInfo, int) bool, or NullableMemberInfo) NullableMemberInfo {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMemberInfoStream) FindIndex(fn func(NullableMemberInfo, int) bool) int {
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
func (self *NullableMemberInfoStream) First() *NullableMemberInfo {
	return self.Get(0)
}
func (self *NullableMemberInfoStream) FirstOr(arg NullableMemberInfo) NullableMemberInfo {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMemberInfoStream) ForEach(fn func(NullableMemberInfo, int)) *NullableMemberInfoStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMemberInfoStream) ForEachRight(fn func(NullableMemberInfo, int)) *NullableMemberInfoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMemberInfoStream) GroupBy(fn func(NullableMemberInfo, int) string) map[string][]NullableMemberInfo {
	m := map[string][]NullableMemberInfo{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMemberInfoStream) GroupByValues(fn func(NullableMemberInfo, int) string) [][]NullableMemberInfo {
	var tmp [][]NullableMemberInfo
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMemberInfoStream) IndexOf(arg NullableMemberInfo) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMemberInfoStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMemberInfoStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMemberInfoStream) Last() *NullableMemberInfo {
	return self.Get(self.Len() - 1)
}
func (self *NullableMemberInfoStream) LastOr(arg NullableMemberInfo) NullableMemberInfo {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMemberInfoStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMemberInfoStream) Limit(limit int) *NullableMemberInfoStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMemberInfoStream) Map(fn func(NullableMemberInfo, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberInfoStream) Map2Int(fn func(NullableMemberInfo, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberInfoStream) Map2Int32(fn func(NullableMemberInfo, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberInfoStream) Map2Int64(fn func(NullableMemberInfo, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberInfoStream) Map2Float32(fn func(NullableMemberInfo, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberInfoStream) Map2Float64(fn func(NullableMemberInfo, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberInfoStream) Map2Bool(fn func(NullableMemberInfo, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberInfoStream) Map2Bytes(fn func(NullableMemberInfo, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberInfoStream) Map2String(fn func(NullableMemberInfo, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberInfoStream) Max(fn func(NullableMemberInfo, int) float64) *NullableMemberInfo {
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
func (self *NullableMemberInfoStream) Min(fn func(NullableMemberInfo, int) float64) *NullableMemberInfo {
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
func (self *NullableMemberInfoStream) NoneMatch(fn func(NullableMemberInfo, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMemberInfoStream) Get(index int) *NullableMemberInfo {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMemberInfoStream) GetOr(index int, arg NullableMemberInfo) NullableMemberInfo {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMemberInfoStream) Peek(fn func(*NullableMemberInfo, int)) *NullableMemberInfoStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *NullableMemberInfoStream) Reduce(fn func(NullableMemberInfo, NullableMemberInfo, int) NullableMemberInfo) *NullableMemberInfoStream {
	return self.ReduceInit(fn, NullableMemberInfo{})
}
func (self *NullableMemberInfoStream) ReduceInit(fn func(NullableMemberInfo, NullableMemberInfo, int) NullableMemberInfo, initialValue NullableMemberInfo) *NullableMemberInfoStream {
	result := NullableMemberInfoStreamOf()
	self.ForEach(func(v NullableMemberInfo, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMemberInfoStream) ReduceInterface(fn func(interface{}, NullableMemberInfo, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMemberInfo{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMemberInfoStream) ReduceString(fn func(string, NullableMemberInfo, int) string) []string {
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
func (self *NullableMemberInfoStream) ReduceInt(fn func(int, NullableMemberInfo, int) int) []int {
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
func (self *NullableMemberInfoStream) ReduceInt32(fn func(int32, NullableMemberInfo, int) int32) []int32 {
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
func (self *NullableMemberInfoStream) ReduceInt64(fn func(int64, NullableMemberInfo, int) int64) []int64 {
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
func (self *NullableMemberInfoStream) ReduceFloat32(fn func(float32, NullableMemberInfo, int) float32) []float32 {
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
func (self *NullableMemberInfoStream) ReduceFloat64(fn func(float64, NullableMemberInfo, int) float64) []float64 {
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
func (self *NullableMemberInfoStream) ReduceBool(fn func(bool, NullableMemberInfo, int) bool) []bool {
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
func (self *NullableMemberInfoStream) Reverse() *NullableMemberInfoStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMemberInfoStream) Replace(fn func(NullableMemberInfo, int) NullableMemberInfo) *NullableMemberInfoStream {
	return self.ForEach(func(v NullableMemberInfo, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMemberInfoStream) Select(fn func(NullableMemberInfo) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMemberInfoStream) Set(index int, val NullableMemberInfo) *NullableMemberInfoStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMemberInfoStream) Skip(skip int) *NullableMemberInfoStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMemberInfoStream) SkippingEach(fn func(NullableMemberInfo, int) int) *NullableMemberInfoStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMemberInfoStream) Slice(startIndex, n int) *NullableMemberInfoStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMemberInfo{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMemberInfoStream) Sort(fn func(i, j int) bool) *NullableMemberInfoStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMemberInfoStream) Tail() *NullableMemberInfo {
	return self.Last()
}
func (self *NullableMemberInfoStream) TailOr(arg NullableMemberInfo) NullableMemberInfo {
	return self.LastOr(arg)
}
func (self *NullableMemberInfoStream) ToList() []NullableMemberInfo {
	return self.Val()
}
func (self *NullableMemberInfoStream) Unique() *NullableMemberInfoStream {
	return self.Distinct()
}
func (self *NullableMemberInfoStream) Val() []NullableMemberInfo {
	if self == nil {
		return []NullableMemberInfo{}
	}
	return *self.Copy()
}
func (self *NullableMemberInfoStream) While(fn func(NullableMemberInfo, int) bool) *NullableMemberInfoStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMemberInfoStream) Where(fn func(NullableMemberInfo) bool) *NullableMemberInfoStream {
	result := NullableMemberInfoStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMemberInfoStream) WhereSlim(fn func(NullableMemberInfo) bool) *NullableMemberInfoStream {
	result := NullableMemberInfoStreamOf()
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
