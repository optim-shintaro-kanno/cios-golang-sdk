/*
 * Collection utility of BucketName Struct
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

type BucketNameStream []BucketName

func BucketNameStreamOf(arg ...BucketName) BucketNameStream {
	return arg
}
func BucketNameStreamFrom(arg []BucketName) BucketNameStream {
	return arg
}
func CreateBucketNameStream(arg ...BucketName) *BucketNameStream {
	tmp := BucketNameStreamOf(arg...)
	return &tmp
}
func GenerateBucketNameStream(arg []BucketName) *BucketNameStream {
	tmp := BucketNameStreamFrom(arg)
	return &tmp
}

func (self *BucketNameStream) Add(arg BucketName) *BucketNameStream {
	return self.AddAll(arg)
}
func (self *BucketNameStream) AddAll(arg ...BucketName) *BucketNameStream {
	*self = append(*self, arg...)
	return self
}
func (self *BucketNameStream) AddSafe(arg *BucketName) *BucketNameStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *BucketNameStream) Aggregate(fn func(BucketName, BucketName) BucketName) *BucketNameStream {
	result := BucketNameStreamOf()
	self.ForEach(func(v BucketName, i int) {
		if i == 0 {
			result.Add(fn(BucketName{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *BucketNameStream) AllMatch(fn func(BucketName, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *BucketNameStream) AnyMatch(fn func(BucketName, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *BucketNameStream) Clone() *BucketNameStream {
	temp := make([]BucketName, self.Len())
	copy(temp, *self)
	return (*BucketNameStream)(&temp)
}
func (self *BucketNameStream) Copy() *BucketNameStream {
	return self.Clone()
}
func (self *BucketNameStream) Concat(arg []BucketName) *BucketNameStream {
	return self.AddAll(arg...)
}
func (self *BucketNameStream) Contains(arg BucketName) bool {
	return self.FindIndex(func(_arg BucketName, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *BucketNameStream) Clean() *BucketNameStream {
	*self = BucketNameStreamOf()
	return self
}
func (self *BucketNameStream) Delete(index int) *BucketNameStream {
	return self.DeleteRange(index, index)
}
func (self *BucketNameStream) DeleteRange(startIndex, endIndex int) *BucketNameStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *BucketNameStream) Distinct() *BucketNameStream {
	caches := map[string]bool{}
	result := BucketNameStreamOf()
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
func (self *BucketNameStream) Each(fn func(BucketName)) *BucketNameStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *BucketNameStream) EachRight(fn func(BucketName)) *BucketNameStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *BucketNameStream) Equals(arr []BucketName) bool {
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
func (self *BucketNameStream) Filter(fn func(BucketName, int) bool) *BucketNameStream {
	result := BucketNameStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *BucketNameStream) FilterSlim(fn func(BucketName, int) bool) *BucketNameStream {
	result := BucketNameStreamOf()
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
func (self *BucketNameStream) Find(fn func(BucketName, int) bool) *BucketName {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *BucketNameStream) FindOr(fn func(BucketName, int) bool, or BucketName) BucketName {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *BucketNameStream) FindIndex(fn func(BucketName, int) bool) int {
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
func (self *BucketNameStream) First() *BucketName {
	return self.Get(0)
}
func (self *BucketNameStream) FirstOr(arg BucketName) BucketName {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *BucketNameStream) ForEach(fn func(BucketName, int)) *BucketNameStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *BucketNameStream) ForEachRight(fn func(BucketName, int)) *BucketNameStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *BucketNameStream) GroupBy(fn func(BucketName, int) string) map[string][]BucketName {
	m := map[string][]BucketName{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *BucketNameStream) GroupByValues(fn func(BucketName, int) string) [][]BucketName {
	var tmp [][]BucketName
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *BucketNameStream) IndexOf(arg BucketName) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *BucketNameStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *BucketNameStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *BucketNameStream) Last() *BucketName {
	return self.Get(self.Len() - 1)
}
func (self *BucketNameStream) LastOr(arg BucketName) BucketName {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *BucketNameStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *BucketNameStream) Limit(limit int) *BucketNameStream {
	self.Slice(0, limit)
	return self
}

func (self *BucketNameStream) Map(fn func(BucketName, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketNameStream) Map2Int(fn func(BucketName, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketNameStream) Map2Int32(fn func(BucketName, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketNameStream) Map2Int64(fn func(BucketName, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketNameStream) Map2Float32(fn func(BucketName, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketNameStream) Map2Float64(fn func(BucketName, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketNameStream) Map2Bool(fn func(BucketName, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketNameStream) Map2Bytes(fn func(BucketName, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketNameStream) Map2String(fn func(BucketName, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketNameStream) Max(fn func(BucketName, int) float64) *BucketName {
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
func (self *BucketNameStream) Min(fn func(BucketName, int) float64) *BucketName {
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
func (self *BucketNameStream) NoneMatch(fn func(BucketName, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *BucketNameStream) Get(index int) *BucketName {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *BucketNameStream) GetOr(index int, arg BucketName) BucketName {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *BucketNameStream) Peek(fn func(*BucketName, int)) *BucketNameStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *BucketNameStream) Reduce(fn func(BucketName, BucketName, int) BucketName) *BucketNameStream {
	return self.ReduceInit(fn, BucketName{})
}
func (self *BucketNameStream) ReduceInit(fn func(BucketName, BucketName, int) BucketName, initialValue BucketName) *BucketNameStream {
	result := BucketNameStreamOf()
	self.ForEach(func(v BucketName, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *BucketNameStream) ReduceInterface(fn func(interface{}, BucketName, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(BucketName{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *BucketNameStream) ReduceString(fn func(string, BucketName, int) string) []string {
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
func (self *BucketNameStream) ReduceInt(fn func(int, BucketName, int) int) []int {
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
func (self *BucketNameStream) ReduceInt32(fn func(int32, BucketName, int) int32) []int32 {
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
func (self *BucketNameStream) ReduceInt64(fn func(int64, BucketName, int) int64) []int64 {
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
func (self *BucketNameStream) ReduceFloat32(fn func(float32, BucketName, int) float32) []float32 {
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
func (self *BucketNameStream) ReduceFloat64(fn func(float64, BucketName, int) float64) []float64 {
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
func (self *BucketNameStream) ReduceBool(fn func(bool, BucketName, int) bool) []bool {
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
func (self *BucketNameStream) Reverse() *BucketNameStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *BucketNameStream) Replace(fn func(BucketName, int) BucketName) *BucketNameStream {
	return self.ForEach(func(v BucketName, i int) { self.Set(i, fn(v, i)) })
}
func (self *BucketNameStream) Select(fn func(BucketName) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *BucketNameStream) Set(index int, val BucketName) *BucketNameStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *BucketNameStream) Skip(skip int) *BucketNameStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *BucketNameStream) SkippingEach(fn func(BucketName, int) int) *BucketNameStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *BucketNameStream) Slice(startIndex, n int) *BucketNameStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []BucketName{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *BucketNameStream) Sort(fn func(i, j int) bool) *BucketNameStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *BucketNameStream) Tail() *BucketName {
	return self.Last()
}
func (self *BucketNameStream) TailOr(arg BucketName) BucketName {
	return self.LastOr(arg)
}
func (self *BucketNameStream) ToList() []BucketName {
	return self.Val()
}
func (self *BucketNameStream) Unique() *BucketNameStream {
	return self.Distinct()
}
func (self *BucketNameStream) Val() []BucketName {
	if self == nil {
		return []BucketName{}
	}
	return *self.Copy()
}
func (self *BucketNameStream) While(fn func(BucketName, int) bool) *BucketNameStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *BucketNameStream) Where(fn func(BucketName) bool) *BucketNameStream {
	result := BucketNameStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *BucketNameStream) WhereSlim(fn func(BucketName) bool) *BucketNameStream {
	result := BucketNameStreamOf()
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
