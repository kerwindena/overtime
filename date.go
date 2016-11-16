package overtime

type Date interface {
	String() string
	Day() int
}

type Time interface {
	String() string
}
