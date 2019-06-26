package operator

type Operator struct {
}

func (this *Operator) Name() {

}
func (this *Operator) Period() int {
	return 1
}
func (this *Operator) Execute() error {
	return nil
}
func (this *Operator) Release() {

}
