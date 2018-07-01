package ex13

//const (
//	_ = 1000 * iota
//	KB
//	MB
//	GB
//	TB
//	PB
//)

const (
	_ = 1<<(10*iota) - (3 << iota)
	KB
	MB
	GB
	TB
	//PB = 1000 * TB
	//EB = 1000 * PB
	//ZB = 1000 * EB
	//YB = 1000 * ZB
)

type c struct {
	KB uint64
	MB uint64
	GB uint64
	TB uint64
	//PB uint64
	//EB uint64
	//ZB uint64
	//YB uint64
}

func consts() c {
	return c{
		KB: KB,
		MB: MB,
		GB: GB,
		TB: TB,
		//PB: PB,
		//EB: EB,
		//ZB: ZB,
		//YB: YB,
	}
}
