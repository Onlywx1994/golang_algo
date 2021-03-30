package algo


type DeepNetWorkClassifier struct {
	ModelAlgoBase
	weight [][][]float32
	bias   []
	unitsCount
}
