package producer

type SimpleProducer struct {
}

func NewSimpleProducer() *SimpleProducer {
	return &SimpleProducer{}
}

func (this *SimpleProducer) Handle() {
	//ζιηδΊ§
	//global.QUEUE.Publish("simple producer")
}
