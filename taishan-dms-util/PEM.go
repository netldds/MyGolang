package main

import (
	"encoding/pem"
	"fmt"
	"os"
)

var RsaPriKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDCEjvR9tkr02d+ok7sfiASE6NBod3Nb7px3+WmCCdFck9Htmm4
2stMUiJS9AIYszPYZPCGc5vNvmNswD8Afq4s6qtlZoG9AySu4WV5r+peKM8xsJsP
G7sBylNxNkznvlaSMyyw6RYClwuGqE3LYqZVcGJ2F8ky6R5lIkwd02u3BwIDAQAB
AoGBAIUiexU0EZHGKkauVBRbBedAv4uD3qRTVIVQQrve9gnnPEoG8ooU1siZw+BC
crm9HgECZhrmbmj7hjuRuua9AQCvLewdVu4ZDNF3k4O/eqm89KNMt0ifDeKWRx3C
xhyCdvhwfQCBQ3bzDG56sdq58CTllBc6fCFOSqqo2DvaGhX5AkEA6TYSFWFcfr1V
Mhlix+vxNONPUD9kZBA04acP0GUotuwiXb00D8KKQ/OFtKajetZj7lVrnHhC9eLw
dBslGjr7NQJBANUJDQiOIjCgkqbY6d25n3HCJS7YQc1aYRuRlvCqOc0mt73ak4P6
Szry3qCNO+pqSy4OxHD3zazRx5mxd0FY9MsCQG1PGL7Iuc/18n7fAzvtzUsa2Ewm
ymlUZ1T1NyZYo/LJT3pcepCAgMpE1IDOMoYbAw/tHdljTQ9vZYEmUAexaZkCQCLt
hfaGHzLr0L+MRuO0gGDNXP1ONZOuosc7Wo0Ay9NH6s403QTBb74tfbTDEzS+0q6t
eyWua0lPZ7NaNlw/cnsCQCbv5ITU8IXS+RcuWyLFBVcWKhPO4GHHKbpZjiurSvuU
DKcI4ny/GzM/zgZc4NwaKlanyYyLu0jn4iJgRUb1nSk=
-----END RSA PRIVATE KEY-----`

func main() {
	block := pem.Block{
		Type:    "mail",
		Headers: map[string]string{"header": "title"},
		Bytes:   []byte("hello"),
	}
	err := pem.Encode(os.Stdout, &block)
	if err != nil {
		fmt.Println(err)
	}
}
