package consumer

type Consumer struct {
}

func (this *Consumer) Name() {

}
func (this *Consumer) Period() int {
	return 1
}
func (this *Consumer) Execute() error {
	return nil
}
func (this *Consumer) Release() {

}
