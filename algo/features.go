package algo

import (
	"bytes"
	"fmt"
	"math"
	"crypto/md5"
	"encoding/binary"
)

type Features struct {
	featuresMap map[int]float32
}

func (self *Features)checkInit(){
	if self.featuresMap == nil {
		self.featuresMap =make(map[int]float32)
	}
}

func (self *Features) toString() string{
	self.checkInit()
	var buffer bytes.Buffer
	var i int=0
	for key,val :=range self.featuresMap{
		if i!=0{
			buffer.WriteString(",")
		}
		str := fmt.Sprintf("%d:%g" ,key,val)
		buffer.WriteString(str)
		i++
	}
	return buffer.String()
}

func (self *Features) ToMap() map[int]float32{
	self.checkInit()
	return self.featuresMap
}

func (self *Features) FromMap (featuresMap map[int]float32){
	for key,val :=range featuresMap{
		self.Add(key,val)
	}
}

func (self *Features) FromArray(features []float32){
	for key,val :=range features{
		self.Add(key,val)
	}
}

func (self *Features) Add (key int, val float32) bool{
	self.checkInit()
	if key>=0 &&math.Abs(float64(val))>0.000001{
		self.featuresMap[key]=val
		return true
	}
	return false
}

func (self *Features) AddArray(start,length int,vals []float32)bool{
	if vals !=nil{
		for i,val :=range vals{
			if i<length{
				self.Add(start+i,val)
			}
		}
		return true
	}
	return false
}

func (self *Features) AddCategory(start,length,minVal int ,val,def int) bool{
	new_val :=val-minVal
	if new_val>=0&&new_val<length{
		return self.Add(start+new_val,1.0)
	}else{
		return self.Add(start+(def-minVal),1.0)
	}
}

func (self *Features) AddCategories(start,length,minVal int ,vals []int,def int) bool{
	for _,val :=range vals{
		self.AddCategory(start,length,minVal,val,def)
	}
	return true
}

func (self *Features) AddHash(start,length int ,val interface{}) bool{
	bytesVal :=GetBytes(val)
	hashVal :=Md5Sum32(bytesVal)
	restVal :=int(hashVal%int32(length))
	return self.AddCategory(start,length,0,restVal,0)
}

func (self *Features) AddHashStrings(start,length int ,vals []string)bool{
	for _,val :=range vals{
		self.AddHash(start,length,val)
	}
	return true
}

func (self *Features) Get(key int )(float32,bool){
	self.checkInit()
	val,ok :=self.featuresMap[key]
	return val,ok
}
// GetBytes convert interface to []byte.
func GetBytes(v interface{}) []byte {
	switch result := v.(type) {
	case string:
		return []byte(result)
	case []byte:
		return result
	default:
		if v != nil {
			return []byte(fmt.Sprintf("%v", result))
		}
	}
	return nil
}

func Bytes2Int32(data []byte) int32 {
	var x int32
	b_buf := bytes.NewBuffer(data) // 取最后4字节
	err := binary.Read(b_buf, binary.BigEndian, &x)
	if err != nil {
		fmt.Printf("err %s\n", err)
	}
	return x
}

func Md5Sum32(data []byte) int32 {
	hash := md5.New()
	hash.Write(data)
	resByte := hash.Sum(nil)

	return Bytes2Int32(resByte[12:])
}