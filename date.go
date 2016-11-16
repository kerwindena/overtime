package overtime

type Date interface {
	String() string
	GetDay() int
}

type Time interface {
	String() string
}
