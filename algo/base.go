package algo

import (
	"os"
	"compress/gzip"
	"io/ioutil"
	"encoding/json"
)

//************ 算法接口

type IModelAlgo interface {
	Init(string)
	PredictSingle(*Features) float32
	TransformSingle(*Features ) *Features
}

//*************** 算法基类
type ModelAlgoBase struct {
	FeaturesMap FeaturesMapEncoder `json:"feature_map"`
	Features    []string           `json:"features"`
	Description string             `json:"description"`
}

func (self *ModelAlgoBase) PredictSingle(features *Features) float32{
	return 0
}

func (self *ModelAlgoBase) TransformSingle(features *Features) *Features{
	features =self.FeaturesMap.Transform(features)
	return features
}

//模型加载 json -> gzip
func loadModel(path string,model interface{}) bool{
	fr,oerr :=os.Open(path)
	defer fr.Close()
	if oerr !=nil{
		return false
	}
	gzf,gerr :=gzip.NewReader(fr)
	if gerr!=nil{
		return false
	}
	data,rerr :=ioutil.ReadAll(gzf)
	if rerr!=nil{
		return false
	}
	jerr :=json.Unmarshal(data,model)
	if jerr !=nil{
		return false
	}
	return true

}