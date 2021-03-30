package algo

type XgboostLRClassifier struct {
	ModelAlgoBase
	Xgboost XgboostClassifier `json:"xgboost"`
	OneHot OneHotEncoder `json:"one_hot"`
	LR LogisticRegression `json:"lr"`
	FeatureCount int `json:"n_features"`
}

func (self *XgboostLRClassifier) Init (path string){
	loadModel(path,self)
}

func (self *XgboostLRClassifier) PredictSingle (features *Features) float32{
	leafs :=self.Xgboost.PredictSingleLeafs(features)
	new_features :=self.OneHot.Transform(leafs)
	return self.LR.PredictSingle(new_features)
}