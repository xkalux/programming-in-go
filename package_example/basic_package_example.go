package csstudent

import "math/rand"

var stdName = []string{
	"Mr.PHONTHAKRON  MENTA",
	"MissMANITA  SIRIPHURIPHAKORN",
	"MissPHIYADA  DAENGWICHAI",
	"Mr.YODSANAN  MANGAN",
	"Mr.TAKSIN  YODKEEREE",
	"MissTAKSAPORN  CHIRAVATANASOMKUL",
	"Mr.SORNTHEP  CHIMSUNTORN",
	"Mr.KISAPHON  KAEOPHA",
	"Mr.THANACHOT  YINKHUNTHOD",
	"Mr.RATTANAN  JIAMSRANOI",
	"Mr.RATTANAN  JIAMSRANOI",
	"Mr.RATTANAN  JIAMSRANOI",
	"Mr.RATTANAN  JIAMSRANOI",
	"Mr.RATTANAN  JIAMSRANOI",
	"Mr.ANUCHA  NAMLEE",
	"Mr.NATDANAI  PORSUNGNOEN",
	"Mr.NATTHARAT  THEPPHOT",
	"MissMANASSAWEE  SANGBOONTHAI",
	"MissPAWEENA THOTHATSA",
	"Mr.KONGKAT  CHONKEAW",
	"Mr.CHIYANAT  NAKLANG",
	"MissCHAYADA  HOMNAN",
	"Mr.YODSAWAT  CHANSUNGNOEN",
	"Mr.PHUTHANET  SUKSEELA",
	"Mr.JATSADA  THAINSUECHAROEN",
	"Mr.THANATHON  WONGPHAYAKYOTHIN",
	"Mr.HARIRAK  NONTABUT",
}

func Get() string {
	return stdName[rand.Intn(len(stdName))]
}