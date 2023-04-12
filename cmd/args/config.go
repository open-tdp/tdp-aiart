package args

var Debug bool

var Dataset struct {
	Dir    string
	Secret string
}

var Database struct {
	Type   string
	Host   string
	User   string
	Passwd string
	Name   string
	Option string
}

var Logger struct {
	Dir    string
	Level  string
	Target string
}

var Server struct {
	Listen string
	JwtKey string
}
