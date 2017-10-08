package conf

type (
	Db struct {
		Host string
		Port int
	}
)

var (
	DBCONFIG *Db
)

func init(){
	DBCONFIG = &Db{
		Host:"127.0.0.1",
		Port:8888,
	}
}