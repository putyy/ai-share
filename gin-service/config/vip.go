package config

type VipConfig struct {
	Level  int
	Name   string
	Profit int
	Price  int
	Icon   string
	Length int
	VideoParse int
}

var Vip = [4]VipConfig{}

func init() {
	Vip[0] = VipConfig{
		Level:  0,
		Name:   "免费",
		Price:  0,
		Profit: 0,
		Length: 0,
		VideoParse: 3,
	}
	Vip[1] = VipConfig{
		Level:  1,
		Name:   "VIP",
		Price:  99,
		Profit: 0,
		Length: 12,
		VideoParse: 6,
	}
	Vip[2] = VipConfig{
		Level:  2,
		Name:   "SVIP",
		Price:  1999,
		Profit: 25,
		Length: -1,
		VideoParse:  -1,
	}
	Vip[3] = VipConfig{
		Level:  3,
		Name:   "SVIP+",
		Profit: 35,
		Price:  2999,
		Length: -1,
		VideoParse:  -1,
	}
}
