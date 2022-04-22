package models

// CODE IS FROM EITHER THE BLACK HAT GO BOOK OR BLACK HAT GO REPO WITH MINOR MODIFICATION


type CmdLineOpts struct {
	Input string
	Output string
	Meta bool
	Supress bool
	Offset string
	Inject bool
	Payload string
	Type string
	Encode bool
	Decode bool
	Key string
}