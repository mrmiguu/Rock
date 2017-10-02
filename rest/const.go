package rest

const (
	DefaultClientAddr = "/"
	DefaultServerAddr = ":80"
	POST              = "9b466094ec991a03cb95c489c19c4d75635f0ae5"
	GET               = "783923e57ba5e8f1044632c31fd806ee24814bb5"
	V                 = "▼"

	Tbool byte = iota
	Tstring
	Tint
	Tint8
	Tint16
	Tint32
	Tint64
	Tuint
	Tuint8
	Tuint16
	Tuint32
	Tuint64
	Tuintptr
	Tbyte
	Tbytes
	Trune
	Tfloat32
	Tfloat64
	Tcomplex64
	Tcomplex128
)
