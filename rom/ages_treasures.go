package rom

func agesTreasure(id, subID byte, offset uint16,
	mode, param, text, sprite byte) *Treasure {
	return &Treasure{id, subID, Addr{0x16, offset}, mode, param, text, sprite}
}

var agesTreasures = map[string]*Treasure{
	// equip items
	"bombs, 10":       agesTreasure(0x03, 0x00, 0x54ce, 0x38, 0x10, 0x4d, 0x05),
	"iron shield":     agesTreasure(0x01, 0x01, 0x54c2, 0x0a, 0x02, 0x20, 0x14),
	"cane":            agesTreasure(0x04, 0x00, 0x5342, 0x38, 0x00, 0x73, 0x17),
	"sword 1":         agesTreasure(0x05, 0x00, 0x54e6, 0x38, 0x01, 0x1c, 0x10),
	"sword 2":         agesTreasure(0x05, 0x00, 0x54e6, 0x38, 0x01, 0x1c, 0x10),
	"boomerang":       agesTreasure(0x06, 0x00, 0x5502, 0x0a, 0x01, 0x22, 0x1c),
	"switch hook 1":   agesTreasure(0x0a, 0x00, 0x550e, 0x38, 0x01, 0x30, 0x1f),
	"switch hook 2":   agesTreasure(0x0a, 0x00, 0x550e, 0x38, 0x01, 0x30, 0x1f),
	"strange flute":   agesTreasure(0x0e, 0x00, 0x552e, 0x0a, 0x0c, 0x3b, 0x23),
	"ricky's flute":   agesTreasure(0x0e, 0x00, 0x552a, 0x0a, 0x0b, 0x38, 0x23),
	"dimitri's flute": agesTreasure(0x0e, 0x01, 0x552e, 0x0a, 0x0c, 0x39, 0x23),
	"moosh's flute":   agesTreasure(0x0e, 0x02, 0x5532, 0x0a, 0x0d, 0x3a, 0x23),
	"seed shooter":    agesTreasure(0x0f, 0x00, 0x536e, 0x38, 0x01, 0x2e, 0x21),
	"shovel":          agesTreasure(0x15, 0x00, 0x553e, 0x0a, 0x00, 0x25, 0x1b),
	"bracelet 1":      agesTreasure(0x16, 0x00, 0x554a, 0x0a, 0x01, 0x26, 0x19),
	"bracelet 2":      agesTreasure(0x16, 0x00, 0x554a, 0x0a, 0x01, 0x26, 0x19),
	"feather":         agesTreasure(0x17, 0x00, 0x555a, 0x0a, 0x01, 0x27, 0x16),
	"satchel 1":       agesTreasure(0x19, 0x02, 0x556e, 0x29, 0x00, 0x2d, 0x20),
	"satchel 2":       agesTreasure(0x19, 0x02, 0x556e, 0x29, 0x00, 0x2d, 0x20),

	// not used; tune of echoes gives it automatically
	// "harp": agesTreasure(0x11, 0x00, 0x5536, 0x0a, 0x00, 0x71, 0x68),

	// collection items
	"harp 1":          agesTreasure(0x25, 0x00, 0x53c6, 0x68, 0x00, 0x72, 0x68),
	"harp 2":          agesTreasure(0x25, 0x00, 0x53c6, 0x68, 0x00, 0x72, 0x69),
	"harp 3":          agesTreasure(0x25, 0x00, 0x53c6, 0x68, 0x00, 0x72, 0x69),
	"flippers 1":      agesTreasure(0x2e, 0x00, 0x56ce, 0x0a, 0x00, 0x31, 0x31),
	"flippers 2":      agesTreasure(0x2e, 0x00, 0x56ce, 0x0a, 0x00, 0x31, 0x31),
	"broken sword":    agesTreasure(0x41, 0x0b, 0x573a, 0x0a, 0x0b, 0x65, 0x7b),
	"graveyard key":   agesTreasure(0x42, 0x00, 0x543a, 0x29, 0x00, 0x23, 0x44),
	"crown key":       agesTreasure(0x43, 0x00, 0x543e, 0x09, 0x00, 0x3d, 0x45),
	"mermaid key":     agesTreasure(0x44, 0x00, 0x5442, 0x09, 0x00, 0x42, 0x46),
	"old mermaid key": agesTreasure(0x45, 0x00, 0x573e, 0x09, 0x00, 0x43, 0x47),
	"library key":     agesTreasure(0x46, 0x00, 0x544a, 0x02, 0x00, 0x44, 0x48),
	"ricky's gloves":  agesTreasure(0x48, 0x00, 0x5452, 0x51, 0x01, 0x67, 0x55),
	"bomb flower":     agesTreasure(0x49, 0x00, 0x5746, 0x0a, 0x00, 0x3c, 0x56),
	"tuni nut":        agesTreasure(0x4c, 0x00, 0x574e, 0x0a, 0x00, 0x37, 0x5b),
	"scent seedling":  agesTreasure(0x4d, 0x00, 0x5466, 0x0a, 0x00, 0x0d, 0x3e),
	"zora scale":      agesTreasure(0x4e, 0x00, 0x546a, 0x0a, 0x00, 0x47, 0x51),
	"tokay eyeball":   agesTreasure(0x4f, 0x00, 0x546e, 0x0a, 0x00, 0x56, 0x53),
	"fairy powder":    agesTreasure(0x51, 0x00, 0x5476, 0x0a, 0x00, 0x55, 0x58),
	"cheval rope":     agesTreasure(0x52, 0x00, 0x547a, 0x0a, 0x00, 0x7d, 0x3c),
	"island chart":    agesTreasure(0x54, 0x00, 0x5482, 0x0a, 0x00, 0x7c, 0x26),
	"book of seals":   agesTreasure(0x55, 0x00, 0x5486, 0x0a, 0x00, 0x4e, 0x52),
	"goron letter":    agesTreasure(0x59, 0x00, 0x5496, 0x02, 0x00, 0x4a, 0x49),
	"lava juice":      agesTreasure(0x5a, 0x00, 0x549a, 0x0a, 0x00, 0x41, 0x4a),
	"brother emblem":  agesTreasure(0x5b, 0x00, 0x549e, 0x0a, 0x00, 0x0c, 0x4b),
	"goron vase":      agesTreasure(0x5c, 0x00, 0x54a2, 0x0a, 0x00, 0x3f, 0x4c),
	"goronade":        agesTreasure(0x5d, 0x00, 0x5756, 0x0a, 0x00, 0x40, 0x4d),
	"rock brisket":    agesTreasure(0x5e, 0x00, 0x575e, 0x0a, 0x00, 0x3e, 0x4e),

	"rupees, 10":  agesTreasure(0x28, 0x02, 0x55aa, 0x38, 0x04, 0x03, 0x2a),
	"rupees, 20":  agesTreasure(0x28, 0x03, 0x55ae, 0x38, 0x05, 0x04, 0x2b),
	"rupees, 30":  agesTreasure(0x28, 0x04, 0x55b2, 0x38, 0x07, 0x05, 0x2b),
	"rupees, 50":  agesTreasure(0x28, 0x05, 0x55b6, 0x38, 0x0b, 0x06, 0x2c),
	"rupees, 100": agesTreasure(0x28, 0x06, 0x55ba, 0x38, 0x0c, 0x07, 0x2d),
	"rupees, 200": agesTreasure(0x28, 0x08, 0x55c2, 0x38, 0x0d, 0x09, 0x2e),

	"piece of heart": agesTreasure(0x2b, 0x01, 0x5602, 0x38, 0x01, 0x17, 0x3a),

	// rings
	"discovery ring":  agesTreasure(0x2d, 0x04, 0x563a, 0x38, 0x28, 0x54, 0x0e),
	"power ring L-1":  agesTreasure(0x2d, 0x0e, 0x5662, 0x38, 0x01, 0x54, 0x0e),
	"gold joy ring":   agesTreasure(0x2d, 0x16, 0x5682, 0x38, 0x26, 0x54, 0x0e),
	"armor ring L-1":  agesTreasure(0x2d, 0x17, 0x5686, 0x38, 0x04, 0x54, 0x0e),
	"light ring L-1":  agesTreasure(0x2d, 0x19, 0x568e, 0x38, 0x17, 0x54, 0x0e),
	"blue luck ring":  agesTreasure(0x2d, 0x1a, 0x5692, 0x38, 0x1b, 0x54, 0x0e),
	"power ring L-2":  agesTreasure(0x2d, 0x1b, 0x5696, 0x38, 0x02, 0x54, 0x0e),
	"gold luck ring":  agesTreasure(0x2d, 0x1c, 0x569a, 0x38, 0x1c, 0x54, 0x0e),
	"pegasus ring":    agesTreasure(0x2d, 0x1e, 0x56a2, 0x38, 0x11, 0x54, 0x0e),
	"green luck ring": agesTreasure(0x2d, 0x20, 0x56aa, 0x38, 0x1a, 0x54, 0x0e),
	"green holy ring": agesTreasure(0x2d, 0x21, 0x56ae, 0x38, 0x1e, 0x54, 0x0e),
	"red holy ring":   agesTreasure(0x2d, 0x22, 0x56b2, 0x38, 0x20, 0x54, 0x0e),
	"whisp ring":      agesTreasure(0x2d, 0x23, 0x56b6, 0x38, 0x39, 0x54, 0x0e),
	"whimsical ring":  agesTreasure(0x2d, 0x25, 0x56be, 0x38, 0x3e, 0x54, 0x0e),
	"toss ring":       agesTreasure(0x2d, 0x26, 0x56c2, 0x38, 0x12, 0x54, 0x0e),
	"blue ring":       agesTreasure(0x2d, 0x27, 0x56c6, 0x38, 0x08, 0x54, 0x0e),
	"like-like ring":  agesTreasure(0x2d, 0x28, 0x56ca, 0x38, 0x2c, 0x54, 0x0e),

	"d1 boss key": agesTreasure(0x31, 0x03, 0x56f2, 0x38, 0x00, 0x1b, 0x43),
	"d2 boss key": agesTreasure(0x31, 0x03, 0x56f2, 0x38, 0x00, 0x1b, 0x43),
	"d3 boss key": agesTreasure(0x31, 0x03, 0x56f2, 0x38, 0x00, 0x1b, 0x43),
	"d4 boss key": agesTreasure(0x31, 0x03, 0x56f2, 0x38, 0x00, 0x1b, 0x43),
	"d5 boss key": agesTreasure(0x31, 0x03, 0x56f2, 0x38, 0x00, 0x1b, 0x43),
	"d6 boss key": agesTreasure(0x31, 0x03, 0x56f2, 0x38, 0x00, 0x1b, 0x43),
	"d7 boss key": agesTreasure(0x31, 0x03, 0x56f2, 0x38, 0x00, 0x1b, 0x43),
	"d8 boss key": agesTreasure(0x31, 0x03, 0x56f2, 0x38, 0x00, 0x1b, 0x43),
	"compass":     agesTreasure(0x32, 0x02, 0x56fe, 0x68, 0x00, 0x19, 0x41),
	"dungeon map": agesTreasure(0x33, 0x02, 0x570a, 0x68, 0x00, 0x18, 0x40),

	"gasha seed": agesTreasure(0x34, 0x01, 0x5582, 0x38, 0x01, 0x4b, 0x0d),

	// not real treasures, just placeholders for seeds in trees
	"ember tree seeds":   &Treasure{id: 0x00},
	"scent tree seeds":   &Treasure{id: 0x01},
	"pegasus tree seeds": &Treasure{id: 0x02},
	"gale tree seeds":    &Treasure{id: 0x03},
	"mystery tree seeds": &Treasure{id: 0x04},
}
