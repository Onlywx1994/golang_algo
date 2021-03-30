package algo

type XgboostClassifier struct {
	ModelAlgoBase
	MaxDepth int `json:"max_depth"`
	FeatureCount int `json:"n_features"`
	InitValue float32 `json:"init_value"`
	ClassCount int `json:"n_classes"`
	EstimatorCount int `json:"n_estimators"`
	LearningRate float32 `json:"learning_rate"`
	Estimators []DecisionTreeRegressor `json:"estimators"`
}

func (self *XgboostClassifier) Init(path string){
	loadModel(path,self)
}

func (self *XgboostClassifier) PredictSingle(features *Features)float32{
	score :=self.InitValue
	for i:=0;i<self.EstimatorCount;i++{
		score +=self.Estimators[i].PredictSingle(features)
	}
	return Expit(score)
}


func (self *XgboostClassifier)PredictSingleLeafs(features *Features)[]int{
	leafs :=make([]int ,self.EstimatorCount)
	for i:=0;i<self.EstimatorCount;i++{
		leafs[i]=self.Estimators[i].PredictSingleLeaf(features)
	}
	return leafs
}