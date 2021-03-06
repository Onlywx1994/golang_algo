package algo


type GradientBoostingClassifier struct {
	ModelAlgoBase
	MaxDepth int `json:"max_depth"`
	FeatureCount int `json:"n_features"`
	InitValue  float32 `json:"init_value"`
	ClassCount  int `json:"n_classes"`
	EstimatorCount  int `json:"n_estimators"`
	LearningRate  float32  `json:"learning_rate"`
	Estimators []DecisionTreeRegressor `json:"estimators"`
}

func (self *GradientBoostingClassifier) Init(path string){
	loadModel(path,self)
}

func (self *GradientBoostingClassifier) PredictSingle(features *Features) float32{
	socre :=self.InitValue
	for i :=0;i<self.EstimatorCount;i++{
		socre+=self.LearningRate*self.Estimators[i].PredictSingle(features)
	}
	return Expit(socre)
}

func (self *GradientBoostingClassifier) PredictSingleLeafs(features *Features) []int{
	leafs :=make([]int ,self.EstimatorCount)
	for i :=0;i<self.EstimatorCount;i++{
		leafs[i]=self.Estimators[i].PredictSingleLeaf(features)
	}
	return leafs
}