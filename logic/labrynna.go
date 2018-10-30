package logic

// this all assumes that you start in the forest of time and that the time
// portals on the screens next to the maku tree are always active.

var labrynnaNodes = map[string]*Node{
	// forest of time
	"start":          And(),
	"starting chest": AndSlot("start"),
	"nayru's house":  AndSlot("start"),

	// lynna / south shore / palace
	"lynna city":         Or("break bush", "flute", "echoes"),
	"lynna village":      Or("lynna city", "echoes"),
	"black tower worker": AndSlot("lynna village"),
	"maku tree": OrSlot("rescue nayru",
		And("lynna village", "shovel", "kill normal")),
	"south lynna tree": AndSlot("lynna city", "sword", "seed item"),
	"lynna city chest": OrSlot("ember seeds", "currents"),
	"shore present": Or("flute", "ricky's gloves",
		And("break bush", "feather"), And("lynna city", "bracelet"),
		And("currents", Or("feather", "flippers", "raft"))),
	"south shore dirt": AndSlot("shore present", "shovel"),
	"tingle": And("feather", Or("sword", "boomerang"),
		Or("currents", And(Or("break bush", "shore present"),
			Or("any seed shooter", "ricky's gloves", "ricky's flute")))),
	"tingle's gift": AndSlot("tingle"),
	"tingle's upgrade": AndSlot("tingle", Or( // 3 types of seeds
		And("ember seeds", Or(
			And("scent seeds",
				Or("pegasus seeds", "gale seeds", "mystery seeds")),
			And("pegasus seeds", Or("gale seeds", "mystery seeds")),
			And("gale seeds", "mystery seeds"))),
		And("scent seeds", Or(
			And("pegasus seeds", Or("gale seeds", "mystery seeds")),
			And("gale seeds", "mystery seeds"))),
		And("pegasus seeds", "gale seeds", "mystery seeds"))),
	"raft":               And("lynna village", "cheval rope", "island chart"),
	"shop, 150 rupees":   AndSlot("lynna city"),
	"ambi's palace tree": AndSlot("lynna village", "sword", "seed item"),
	"ambi's palace chest": AndSlot("lynna village", Or(
		HardAnd("satchel", "scent seeds", "pegasus seeds"),
		And("break bush", "mermaid suit"))),
	"rescue nayru": AndSlot("ambi's palace chest", "mystery seeds",
		"switch hook"),
	"mayor plen's house": AndSlot("long hook"),
	"maku seed": And("d1 essence", "d2 essence", "d3 essence", "d4 essence",
		"d5 essence", "d6 essence", "d7 essence", "d8 essence"),

	// yoll graveyard
	"yoll graveyard": And("ember seeds"),
	"yoll moosh":     Or("moosh's flute", And("yoll graveyard", "kill ghini")),
	"cheval's grave": Or("yoll moosh", "bomb jump 3"),
	"cheval's test": AndSlot("cheval's grave", "bracelet",
		Or("feather", "flippers")),
	"cheval's invention": AndSlot("cheval's grave", "flippers"),
	"grave under tree":   AndSlot("yoll graveyard"),
	"syrup":              And("yoll graveyard", Or("flippers", "bomb jump 2")),
	"enter d1":           And("yoll graveyard", "graveyard key"),

	// western woods
	"fairies' woods chest": AndSlot("lynna city", Or(
		And(Or("bracelet", "flippers", "dimitri's flute"),
			Or("feather", "ricky's flute", "moosh's flute", "ages")),
		And("bracelet", "currents"))),
	"deku forest": And("lynna city", Or("bracelet",
		And(Or("flippers", "dimitri's flute"), "ages"))),
	"deku forest cave east": AndSlot("deku forest",
		Or("feather", "bracelet", "ages")),
	"deku forest cave west": AndSlot("deku forest", "bracelet",
		Or("feather", "ember seeds", "ages")),
	"deku forest tree": AndSlot("deku forest", Or("ember seeds", "ages"),
		"sword", "seed item"),
	"deku forest soldier": AndSlot("deku forest", "mystery seeds"),
	"enter d2":            And("deku forest", Or("bombs", "currents")),

	// crescent island
	"crescent past": Or("raft", And("lynna city", "mermaid suit"),
		And("crescent present west", "echoes", "ages")),
	"tokay crystal cave": AndSlot("crescent past",
		Or("shovel", "break crystal")),
	"tokay bomb cave":       AndSlot("crescent past", "bracelet", "bombs"),
	"wild tokay game":       AndSlot("crescent past", "bombs", "bracelet"),
	"crescent present east": And("crescent past", "harp"),
	"crescent island tree": AndSlot("crescent present east", "scent seedling",
		"sword", "seed item"),
	"crescent present west": Or("dimitri's flute",
		And("lynna city", "mermaid suit"), And("crescent past", "harp")),
	"enter d3":               And("crescent present west"),
	"hidden tokay cave":      AndSlot("lynna city", "mermaid suit"),
	"crescent seafloor cave": AndSlot("lynna city", "mermaid suit"),
	"tokay pot cave":         AndSlot("crescent past", "long hook"),

	// nuun / symmetry city / talus peaks
	"ricky nuun":   Root(),
	"dimitri nuun": Root(),
	"moosh nuun":   Root("start"),
	"nuun": And("lynna city", Or("currents",
		And(Or("bracelet", "flippers", "dimitri's flute"), "ember shooter"))),
	"nuun highlands cave": AndSlot("nuun", Or("dimitri's flute",
		And(Or("ricky nuun", "moosh nuun"), Or("flute", "currents")))),
	"symmetry present": And("nuun", Or("ages", "flute",
		And("currents", Or("ricky nuun", "moosh nuun")))),
	"symmetry city tree":    AndSlot("sword", "seed item", "symmetry present"),
	"symmetry past":         And("symmetry present", "echoes"),
	"symmetry city brother": AndSlot("symmetry past"),
	"tokkey's composition":  AndSlot("symmetry past", "flippers"),
	"restoration wall": Or(
		And("lynna city", "ages",
			Or("bracelet", "flippers", "dimitri's flute")),
		And("symmetry past", "currents", "bracelet", "flippers")),
	// TODO do gale seeds work for the ceremony?
	"patch": And(Or("sword", "shield", "boomerang", "switch hook",
		HardOr("scent seeds", "cane")), "restoration wall"),
	"talus peaks chest": OrSlot("restoration wall",
		And(Or("bracelet", "flippers", "dimitri's flute"), "ages")),
	"enter d4": And("symmetry present", "tuni nut", "patch"),

	// rolling ridge. what a nightmare
	"goron elder": AndSlot("bomb flower", "switch hook",
		Or("feather", "ages")),
	"ridge west past": Or("goron elder",
		And("ridge west present", "bracelet", "echoes")),
	"ridge west present": Or("ridge upper present",
		And("switch hook", "currents", Or("feather", "ages")),
		And("currents", Or("ridge west past", "ridge base past"))),
	"ridge west cave":         AndSlot("ridge west present"),
	"rolling ridge west tree": AndSlot("sword", "seed item", "ridge west past"),
	"under moblin keep": AndSlot("ridge west present", "feather",
		"flippers"),
	"defeat great moblin": AndSlot("ridge west present", "pegasus satchel",
		"bracelet"),
	"ridge upper present": Or("ridge base present",
		And("defeat great moblin", "feather")),
	"enter d5": And("crown key", "ridge upper present"),
	"ridge base present": Or("ridge upper present",
		And("currents", Or("ridge base past east", "ridge base past west"))),
	"enter d6 present":    And("mermaid key", "ridge base present"),
	"pool in d6 entrance": And("ridge base present", "mermaid suit"),
	"goron dance present": AndSlot("ridge base present"),
	"goron dance past":    AndSlot("ridge base past", "goron letter"),
	"ridge mid past": Or("ridge base past west",
		And("ridge upper present", "ages"),
		And("ridge base past east", "brother emblem", "feather")),
	"ridge mid present": Or(
		And("ridge mid past", "currents"),
		And("ridge base present", "brother emblem",
			Or("switch hook", "jump 3"))),
	"target carts":     And("ridge mid past", "switch hook", "currents"),
	"target carts 1":   AndSlot("target carts"),
	"target carts 2":   AndSlot("target carts"),
	"shooting gallery": AndSlot("target carts", "sword"),
	"rolling ridge east tree": AndSlot("sword", "seed item",
		Or("target carts", And("ridge mid present", "ages"))),
	"ridge base past east": Or("target carts",
		And(Or("echoes", "lynna city"), Or("feather", "ages"), "mermaid suit"),
		And("ridge mid past", "feather", "brother emblem"),
		And("ridge mid present", "ages"),
		And("ridge base present", "echoes")),
	"ridge base past west": And("ridge base past east",
		Or("flippers", Hard("jump 3"))),
	"ridge base past":     AndSlot("ridge base past west", "bombs"),
	"enter d6 past":       And("old mermaid key", "ridge base past west"),
	"ridge diamonds past": AndSlot("ridge base past west", "switch hook"),
	"bomb goron head": AndSlot("bombs", Or(
		And("ridge base past west", "switch hook"),
		And("ridge upper present", "ages"))),
	"big bang game":         AndSlot("goronade", "ridge mid present"),
	"ridge NE cave present": AndSlot("ridge mid present"),
	"trade rock brisket":    AndSlot("rock brisket", "ridge base present"),
	"trade goron vase":      AndSlot("goron vase", "ridge base past east"),
	"trade lava juice":      AndSlot("lava juice", "ridge mid past"),
	"goron's hiding place":  AndSlot("ridge west present", "bombs"),
	"ridge base chest":      AndSlot("ridge west present"),
	"goron diamond cave": AndSlot("ridge mid present",
		Or("switch hook", "jump 3")),
	"ridge bush cave": AndSlot("ridge mid past", "switch hook"),

	// zora village / zora seas. only accessible with tune of ages, so no
	// distinctions between past and present are necessary.
	"zora village":         And("mermaid suit", "ages", "switch hook"),
	"zora village tree":    AndSlot("zora village", "sword", "seed item"),
	"zora village present": AndSlot("zora village"),
	"zora palace chest":    AndSlot("zora village"),
	"zora NW cave":         AndSlot("zora village", "bombs", "power glove"),
	"fairies' coast chest": AndSlot("zora village"),
	"king zora":            AndSlot("zora village", "syrup"),
	"library present":      AndSlot("zora village", "library key"),
	"library past": AndSlot("zora village", "library key",
		"book of seals"),
	"clean seas":      And("zora village", "fairy powder"),
	"zora seas chest": AndSlot("clean seas"),
	"enter d7":        And("king zora", "clean seas"),
	"zora cave past":  AndSlot("mermaid suit", "ages", "long hook"),
	"zora's reward":   AndSlot("d7 essence"),

	// sea of storms / sea of no return
	"piratian captain":   AndSlot("lynna city", "mermaid suit", "zora scale"),
	"sea of storms past": AndSlot("lynna city", "mermaid suit", "zora scale"),
	"sea of storms present": AndSlot("lynna city", "mermaid suit",
		"zora scale", "currents"),
	"enter d8": And("crescent past", "tokay eyeball", "kill normal", "break pot",
		"bombs", Or("cane", Hard()), "mermaid suit", "feather"),
	"sea of no return": AndSlot("enter d8", "power glove"),

	// trading sequence
	"old zora": AndSlot("yoll graveyard", "graveyard key", "bracelet",
		"crescent present east", "symmetry past",
		Or("switch hook", "mermaid suit")),
}
