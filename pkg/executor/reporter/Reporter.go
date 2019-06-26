package reporter

type Reporter struct {
}

func (this *Reporter) Name() {

}

func (this *Reporter) Period() int {
	return 1
}

func (this *Reporter) Execute() error {
	return nil
}

func (this *Reporter) Release() {

}
