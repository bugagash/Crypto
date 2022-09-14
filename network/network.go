package network

type Package struct {
	Option int
	Data string
}

const (
	ENDBYTES = "\000\005\007\001\007\005\000"
)

func Send(address string, pack *Package) *Package {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil
	}
	conn.Write([]byte(SerializePackage(pack) + ENDBYTES))
	var (
		res = new(Package)
		ch = make(chan bool)
	)
	go func() {
		res = readPackage(conn)
		ch <- true
	}
}