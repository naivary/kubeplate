package inputer

import "github.com/hashicorp/go-plugin"

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "INPUTER_PLUGIN",
	MagicCookieValue: "9d39d112-2c25-4e68-9f8b-55b10b02090b",
}
