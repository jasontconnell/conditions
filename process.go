package conditions

func Load(c string) Condition {
	return Condition{}
}

func (c Condition) Test(t interface{}) bool {
	return false
}
