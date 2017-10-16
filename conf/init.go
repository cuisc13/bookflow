package conf

type (
	Db struct {
		Host string
		Port int
	}
	MgoDb struct {
		Host string
		Port int
	}
)

var (
	DBCONFIG *Db
	MGOCONF *MgoDb
)

func init(){
	DBCONFIG = &Db{
		Host:"127.0.0.1",
		Port:8888,
	}
	MGOCONF = &MgoDb{
		Host:"127.0.0.1",
		Port:27017,
	}
}