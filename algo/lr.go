package algo

type LogisticRegression struct {
	ModelAlgoBase
	CoefMap []float32  `json:"coef"`
	Intercept float32 `json:"intercept"`
}

func (self *LogisticRegression) Init(path string){
	loadModel(path,self)
}

func (self *LogisticRegression) PredictSingle(features *Features)float32{
	var score float32 =self.Intercept
	for i ,feature :=range features.ToMap() {
		score +=self.CoefMap[i]*feature
	}
	return Expit(score)
}
