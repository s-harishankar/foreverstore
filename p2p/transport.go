package p2p

// Peer is an interface that represents the remote node
type Peer interface {}

// Transport is anything that handles communication 
// between peers in the network. This can be of the form TCP, UDP,
// websockers, etc.
type Transport interface {}
