package cmd

import "fmt"

type Operation int

const (
	Read Operation = iota
	Write
)

var Operations = [...]string{
	"Read",
	"Write",
}

func (op Operation) String() string {
	return Operations[op]
}

type Request struct {
	Op    Operation
	Key   string
	Value string
}

type Response struct {
	OK    bool
	Value string
}

func (req Request) String() string {
	switch req.Op {
	case Read:
		return fmt.Sprintf("Read request for key %q", req.Key)
	case Write:
		return fmt.Sprintf("Write request for key %q with value %q",
			req.Key, req.Value)
	default:
		return "Unkown request type"
	}
}

func (resp Response) String() string {
	return fmt.Sprintf("Reponse with OK: %t and value: %q",
		resp.OK, resp.Value)
}
