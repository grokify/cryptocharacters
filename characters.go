package cryptocharacters

import "math"

const (
	CharAlice   = "Alice"
	CharBob     = "Bob"
	CharCarlos  = "Carlos"
	CharCarol   = "Carol"
	CharChad    = "Chad"
	CharCharlie = "Charlie"
	CharChuck   = "Chuck"
	CharCraig   = "Craig"
	CharDan     = "Dan"
	CharDarth   = "Darth"
	CharDave    = "Dave"
	CharDavid   = "David"
	CharErin    = "Erin"
	CharEve     = "Eve"
	CharFaythe  = "Faythe"
	CharFrank   = "Frank"
	CharGrace   = "Grace"
	CharHeidi   = "Heidi"
	CharIvan    = "Ivan"
	CharJudy    = "Judy"
	CharMallory = "Mallory"
	CharMallet  = "Mallet"
	CharMichael = "Michael"
	CharNiaj    = "Niaj"
	CharOlivia  = "Olivia"
	CharOscar   = "Oscar"
	CharPat     = "Pat"
	CharPeggy   = "Peggy"
	CharRupert  = "Rupert"
	CharSybil   = "Sybil"
	CharTed     = "Ted"
	CharTrent   = "Trent"
	CharTrudy   = "Trudy"
	CharVanna   = "Vanna"
	CharVictor  = "Victor"
	CharWalter  = "Walter"
	CharWendy   = "Wendy"
	CharYves    = "Yves"
)

var nameSlice = []string{
	CharAlice,
	CharBob,
	CharCarlos,
	CharCarol,
	CharChad,
	CharCharlie,
	CharChuck,
	CharCraig,
	CharDan,
	CharDarth,
	CharDave,
	CharDavid,
	CharErin,
	CharEve,
	CharFaythe,
	CharFrank,
	CharGrace,
	CharHeidi,
	CharIvan,
	CharJudy,
	CharMallory,
	CharMallet,
	CharMichael,
	CharNiaj,
	CharOlivia,
	CharOscar,
	CharPat,
	CharPeggy,
	CharRupert,
	CharSybil,
	CharTed,
	CharTrent,
	CharTrudy,
	CharVanna,
	CharVictor,
	CharWalter,
	CharWendy,
	CharYves,
}

func SecurityUsers() []string {
	return nameSlice
}

func NameX(x uint) string {
	return nameSlice[int(math.Mod(float64(x), float64(len(nameSlice))))]
}
