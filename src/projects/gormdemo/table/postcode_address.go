package table

type PostcodeAddress struct {
	Id       int64  `gorm:"column:id;primaryKey" json:"id"`
	Postcode string `json:"postcode"`
	State    string `json:"state"`
	City     string `json:"city"`
	Street   string `json:"street"`
	Ctime    int64  `json:"ctime"`
	Utime    int64  `json:"utime"`
}

func (e *PostcodeAddress) TableName() string {
	return "postcode_address"
}
