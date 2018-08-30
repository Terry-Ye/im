package main

type ServerOptions struct {
	WriteWait       int
	PongWait        int
	PingPeriod      int
	MaxMessageSize  int
	ReadBufferSize  int
	WriteBufferSize int
}

type Server struct {
	Buckets []*Bucket // subkey bucket
	Options ServerOptions
}

// NewServer returns a new Server.
func NewServer(b []*Bucket, options ServerOptions) *Server {
	s := new(Server)
	s.Buckets = b
	s.Options = options
	return s
}

// func (server *Server) Bucket(subKey string) *Bucket {
// 	idx := cityhash.CityHash32([]byte(subKey), uint32(len(subKey))) % server.bucketIdx
// 	if Debug {
// 		log.Debug("\"%s\" hit channel bucket index: %d use cityhash", subKey, idx)
// 	}
// 	return server.Buckets[idx]
// }
