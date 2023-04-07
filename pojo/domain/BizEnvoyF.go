package domain

type BizEnvoyF struct {
	Id          float64
	ServiceName string
	Times       string
}

func (BizEnvoyF) TableName() string {
	return "service_update"
}
