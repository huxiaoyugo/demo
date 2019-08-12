package batchUpdateEng


type BatchOption struct {
	onceMaxCount  int
	isShowLog  bool
}

type Option func(option *BatchOption)

func apply(opts *[]Option) BatchOption {
	op := BatchOption{}
	for _, opFunc := range *opts {
		opFunc(&op)
	}
	return op
}

func WithOnceMaxCount(maxCount int) Option{
	return func(option *BatchOption) {
		option.onceMaxCount = maxCount
	}
}

func WithIsShowLog(isShowLog bool) Option{
	return func(option *BatchOption) {
		option.isShowLog = isShowLog
	}
}