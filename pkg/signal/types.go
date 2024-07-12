package signal

type Message struct {
	Opcode string `json:"opcode"`
}

// different message types

type Application struct {
	Message
	community string
	mac       string
}

type Acceptance struct {
	Message
} // Acceptance is a message that indicates that a connection has been accepted

type Rejection struct {
	Message
} // Rejection is a message that indicates that a connection has been rejected

type Offer struct {
	Message
	Payload []byte `json:"payload"`
	Mac     string `json:"mac"`
} // Offer is a message that contains an offer

type Answer struct {
	Message
	Payload []byte `json:"payload"`
	Mac     string `json:"mac"`
} // Answer is a message that contains an answer

type Candidate struct {
	Message
	Payload []byte `json:"payload"`
	Mac     string `json:"mac"`
} // Candidate is a message that contains a candidate

type Exited struct {
	Message
	Mac string `json:"mac"`
} // Exited is a message that indicates that a device has exited the community

type Resignation struct {
	Message
	Mac string `json:"mac"`
} // Resignation is a message that indicates that a device has resigned from the community

type Ready struct {
	Message
	Mac string `json:"mac"`
}

type Introduction struct {
	Message
	Mac string `json:"mac"`
}
