package p2p

// Peer is an interface that represents the remote node.
type Peer interface{

}

// Transport is anything that handles the communication
// between the nodes(aka peers) in the network. This can
// be of the form (TCP, UDP, wesockets, ...)
type Transport interface{}
