package pokepaste

var pokemonData = map[string]map[string]interface{}{
	"Missingno": {
		"id":   uint(0),
		"form": uint(0),
		"type": "???",
	},
	"Bulbasaur": {
		"id":   uint(1),
		"form": uint(0),
		"type": "grass",
	},
	"Ivysaur": {
		"id":   uint(2),
		"form": uint(0),
		"type": "grass",
	},
	"Venusaur": {
		"id":   uint(3),
		"form": uint(0),
		"type": "grass",
	},
	"Venusaur-Mega": {
		"id":   uint(3),
		"form": uint(1),
		"type": "grass",
	},
	"Venusaur-Gmax": {
		"id":   uint(3),
		"form": uint(2),
		"type": "grass",
	},
	"Charmander": {
		"id":   uint(4),
		"form": uint(0),
		"type": "fire",
	},
	"Charmeleon": {
		"id":   uint(5),
		"form": uint(0),
		"type": "fire",
	},
	"Charizard": {
		"id":   uint(6),
		"form": uint(0),
		"type": "fire",
	},
	"Charizard-Mega-X": {
		"id":   uint(6),
		"form": uint(1),
		"type": "fire",
	},
	"Charizard-Mega-Y": {
		"id":   uint(6),
		"form": uint(2),
		"type": "fire",
	},
	"Charizard-Gmax": {
		"id":   uint(6),
		"form": uint(3),
		"type": "fire",
	},
	"Squirtle": {
		"id":   uint(7),
		"form": uint(0),
		"type": "water",
	},
	"Wartortle": {
		"id":   uint(8),
		"form": uint(0),
		"type": "water",
	},
	"Blastoise": {
		"id":   uint(9),
		"form": uint(0),
		"type": "water",
	},
	"Blastoise-Mega": {
		"id":   uint(9),
		"form": uint(1),
		"type": "water",
	},
	"Blastoise-Gmax": {
		"id":   uint(9),
		"form": uint(2),
		"type": "water",
	},
	"Caterpie": {
		"id":   uint(10),
		"form": uint(0),
		"type": "bug",
	},
	"Metapod": {
		"id":   uint(11),
		"form": uint(0),
		"type": "bug",
	},
	"Butterfree": {
		"id":   uint(12),
		"form": uint(0),
		"type": "bug",
	},
	"Butterfree-Gmax": {
		"id":   uint(12),
		"form": uint(1),
		"type": "bug",
	},
	"Weedle": {
		"id":   uint(13),
		"form": uint(0),
		"type": "bug",
	},
	"Kakuna": {
		"id":   uint(14),
		"form": uint(0),
		"type": "bug",
	},
	"Beedrill": {
		"id":   uint(15),
		"form": uint(0),
		"type": "bug",
	},
	"Beedrill-Mega": {
		"id":   uint(15),
		"form": uint(1),
		"type": "bug",
	},
	"Pidgey": {
		"id":   uint(16),
		"form": uint(0),
		"type": "normal",
	},
	"Pidgeotto": {
		"id":   uint(17),
		"form": uint(0),
		"type": "normal",
	},
	"Pidgeot": {
		"id":   uint(18),
		"form": uint(0),
		"type": "normal",
	},
	"Pidgeot-Mega": {
		"id":   uint(18),
		"form": uint(1),
		"type": "normal",
	},
	"Rattata": {
		"id":   uint(19),
		"form": uint(0),
		"type": "normal",
	},
	"Rattata-Alola": {
		"id":   uint(19),
		"form": uint(1),
		"type": "dark",
	},
	"Raticate": {
		"id":   uint(20),
		"form": uint(0),
		"type": "normal",
	},
	"Raticate-Alola": {
		"id":   uint(20),
		"form": uint(1),
		"type": "dark",
	},
	"Raticate-Alola-Totem": {
		"id":   uint(20),
		"form": uint(1),
		"type": "dark",
	},
	"Spearow": {
		"id":   uint(21),
		"form": uint(0),
		"type": "normal",
	},
	"Fearow": {
		"id":   uint(22),
		"form": uint(0),
		"type": "normal",
	},
	"Ekans": {
		"id":   uint(23),
		"form": uint(0),
		"type": "poison",
	},
	"Arbok": {
		"id":   uint(24),
		"form": uint(0),
		"type": "poison",
	},
	"Pikachu": {
		"id":   uint(25),
		"form": uint(0),
		"type": "electric",
	},
	"Pikachu-Gmax": {
		"id":   uint(25),
		"form": uint(1),
		"type": "electric",
	},
	"Raichu": {
		"id":   uint(26),
		"form": uint(0),
		"type": "electric",
	},
	"Raichu-Alola": {
		"id":   uint(26),
		"form": uint(1),
		"type": "electric",
	},
	"Sandshrew": {
		"id":   uint(27),
		"form": uint(0),
		"type": "ground",
	},
	"Sandshrew-Alola": {
		"id":   uint(27),
		"form": uint(1),
		"type": "ice",
	},
	"Sandslash": {
		"id":   uint(28),
		"form": uint(0),
		"type": "ground",
	},
	"Sandslash-Alola": {
		"id":   uint(28),
		"form": uint(1),
		"type": "ice",
	},
	"Nidoran-F": {
		"id":   uint(29),
		"form": uint(0),
		"type": "poison",
	},
	"Nidorina": {
		"id":   uint(30),
		"form": uint(0),
		"type": "poison",
	},
	"Nidoqueen": {
		"id":   uint(31),
		"form": uint(0),
		"type": "poison",
	},
	"Nidoran-M": {
		"id":   uint(32),
		"form": uint(0),
		"type": "poison",
	},
	"Nidorino": {
		"id":   uint(33),
		"form": uint(0),
		"type": "poison",
	},
	"Nidoking": {
		"id":   uint(34),
		"form": uint(0),
		"type": "poison",
	},
	"Clefairy": {
		"id":   uint(35),
		"form": uint(0),
		"type": "fairy",
	},
	"Clefable": {
		"id":   uint(36),
		"form": uint(0),
		"type": "fairy",
	},
	"Vulpix": {
		"id":   uint(37),
		"form": uint(0),
		"type": "fire",
	},
	"Vulpix-Alola": {
		"id":   uint(37),
		"form": uint(1),
		"type": "ice",
	},
	"Ninetales": {
		"id":   uint(38),
		"form": uint(0),
		"type": "fire",
	},
	"Ninetales-Alola": {
		"id":   uint(38),
		"form": uint(1),
		"type": "ice",
	},
	"Jigglypuff": {
		"id":   uint(39),
		"form": uint(0),
		"type": "normal",
	},
	"Wigglytuff": {
		"id":   uint(40),
		"form": uint(0),
		"type": "normal",
	},
	"Zubat": {
		"id":   uint(41),
		"form": uint(0),
		"type": "poison",
	},
	"Golbat": {
		"id":   uint(42),
		"form": uint(0),
		"type": "poison",
	},
	"Oddish": {
		"id":   uint(43),
		"form": uint(0),
		"type": "grass",
	},
	"Gloom": {
		"id":   uint(44),
		"form": uint(0),
		"type": "grass",
	},
	"Vileplume": {
		"id":   uint(45),
		"form": uint(0),
		"type": "grass",
	},
	"Paras": {
		"id":   uint(46),
		"form": uint(0),
		"type": "bug",
	},
	"Parasect": {
		"id":   uint(47),
		"form": uint(0),
		"type": "bug",
	},
	"Venonat": {
		"id":   uint(48),
		"form": uint(0),
		"type": "bug",
	},
	"Venomoth": {
		"id":   uint(49),
		"form": uint(0),
		"type": "bug",
	},
	"Diglett": {
		"id":   uint(50),
		"form": uint(0),
		"type": "ground",
	},
	"Diglett-Alola": {
		"id":   uint(50),
		"form": uint(1),
		"type": "ground",
	},
	"Dugtrio": {
		"id":   uint(51),
		"form": uint(0),
		"type": "ground",
	},
	"Dugtrio-Alola": {
		"id":   uint(51),
		"form": uint(1),
		"type": "ground",
	},
	"Meowth": {
		"id":   uint(52),
		"form": uint(0),
		"type": "normal",
	},
	"Meowth-Alola": {
		"id":   uint(52),
		"form": uint(1),
		"type": "dark",
	},
	"Meowth-Galar": {
		"id":   uint(52),
		"form": uint(2),
		"type": "steel",
	},
	"Meowth-Gmax": {
		"id":   uint(52),
		"form": uint(3),
		"type": "steel",
	},
	"Persian": {
		"id":   uint(53),
		"form": uint(0),
		"type": "normal",
	},
	"Persian-Alola": {
		"id":   uint(53),
		"form": uint(1),
		"type": "dark",
	},
	"Psyduck": {
		"id":   uint(54),
		"form": uint(0),
		"type": "water",
	},
	"Golduck": {
		"id":   uint(55),
		"form": uint(0),
		"type": "water",
	},
	"Mankey": {
		"id":   uint(56),
		"form": uint(0),
		"type": "fighting",
	},
	"Primeape": {
		"id":   uint(57),
		"form": uint(0),
		"type": "fighting",
	},
	"Growlithe": {
		"id":   uint(58),
		"form": uint(0),
		"type": "fire",
	},
	"Arcanine": {
		"id":   uint(59),
		"form": uint(0),
		"type": "fire",
	},
	"Poliwag": {
		"id":   uint(60),
		"form": uint(0),
		"type": "water",
	},
	"Poliwhirl": {
		"id":   uint(61),
		"form": uint(0),
		"type": "water",
	},
	"Poliwrath": {
		"id":   uint(62),
		"form": uint(0),
		"type": "water",
	},
	"Abra": {
		"id":   uint(63),
		"form": uint(0),
		"type": "psychic",
	},
	"Kadabra": {
		"id":   uint(64),
		"form": uint(0),
		"type": "psychic",
	},
	"Alakazam": {
		"id":   uint(65),
		"form": uint(0),
		"type": "psychic",
	},
	"Alakazam-Mega": {
		"id":   uint(65),
		"form": uint(1),
		"type": "psychic",
	},
	"Machop": {
		"id":   uint(66),
		"form": uint(0),
		"type": "fighting",
	},
	"Machoke": {
		"id":   uint(67),
		"form": uint(0),
		"type": "fighting",
	},
	"Machamp": {
		"id":   uint(68),
		"form": uint(0),
		"type": "fighting",
	},
	"Machamp-Gmax": {
		"id":   uint(68),
		"form": uint(1),
		"type": "fighting",
	},
	"Bellsprout": {
		"id":   uint(69),
		"form": uint(0),
		"type": "grass",
	},
	"Weepinbell": {
		"id":   uint(70),
		"form": uint(0),
		"type": "grass",
	},
	"Victreebel": {
		"id":   uint(71),
		"form": uint(0),
		"type": "grass",
	},
	"Tentacool": {
		"id":   uint(72),
		"form": uint(0),
		"type": "water",
	},
	"Tentacruel": {
		"id":   uint(73),
		"form": uint(0),
		"type": "water",
	},
	"Geodude": {
		"id":   uint(74),
		"form": uint(0),
		"type": "rock",
	},
	"Geodude-Alola": {
		"id":   uint(74),
		"form": uint(1),
		"type": "rock",
	},
	"Graveler": {
		"id":   uint(75),
		"form": uint(0),
		"type": "rock",
	},
	"Graveler-Alola": {
		"id":   uint(75),
		"form": uint(1),
		"type": "rock",
	},
	"Golem": {
		"id":   uint(76),
		"form": uint(0),
		"type": "rock",
	},
	"Golem-Alola": {
		"id":   uint(76),
		"form": uint(1),
		"type": "rock",
	},
	"Ponyta": {
		"id":   uint(77),
		"form": uint(0),
		"type": "fire",
	},
	"Ponyta-Galar": {
		"id":   uint(77),
		"form": uint(1),
		"type": "psychic",
	},
	"Rapidash": {
		"id":   uint(78),
		"form": uint(0),
		"type": "fire",
	},
	"Rapidash-Galar": {
		"id":   uint(78),
		"form": uint(1),
		"type": "psychic",
	},
	"Slowpoke": {
		"id":   uint(79),
		"form": uint(0),
		"type": "water",
	},
	"Slowpoke-Galar": {
		"id":   uint(79),
		"form": uint(1),
		"type": "psychic",
	},
	"Slowbro": {
		"id":   uint(80),
		"form": uint(0),
		"type": "water",
	},
	"Slowbro-Mega": {
		"id":   uint(80),
		"form": uint(1),
		"type": "water",
	},
	"Slowbro-Galar": {
		"id":   uint(80),
		"form": uint(2),
		"type": "poison",
	},
	"Magnemite": {
		"id":   uint(81),
		"form": uint(0),
		"type": "electric",
	},
	"Magneton": {
		"id":   uint(82),
		"form": uint(0),
		"type": "electric",
	},
	"Farfetch'd": {
		"id":   uint(83),
		"form": uint(0),
		"type": "normal",
	},
	"Farfetch&#x27;d": {
		"id":   uint(83),
		"form": uint(0),
		"type": "normal",
	},
	"Farfetch’d": {
		"id":   uint(83),
		"form": uint(0),
		"type": "normal",
	},
	"Farfetch&#39;d": {
		"id":   uint(83),
		"form": uint(0),
		"type": "normal",
	},
	"Farfetch'd-Galar": {
		"id":   uint(83),
		"form": uint(1),
		"type": "fighting",
	},
	"Farfetch&#x27;d-Galar": {
		"id":   uint(83),
		"form": uint(1),
		"type": "fighting",
	},
	"Farfetch’d-Galar": {
		"id":   uint(83),
		"form": uint(1),
		"type": "fighting",
	},
	"Farfetch&#39;d-Galar": {
		"id":   uint(83),
		"form": uint(1),
		"type": "fighting",
	},
	"Doduo": {
		"id":   uint(84),
		"form": uint(0),
		"type": "normal",
	},
	"Dodrio": {
		"id":   uint(85),
		"form": uint(0),
		"type": "normal",
	},
	"Seel": {
		"id":   uint(86),
		"form": uint(0),
		"type": "water",
	},
	"Dewgong": {
		"id":   uint(87),
		"form": uint(0),
		"type": "water",
	},
	"Grimer": {
		"id":   uint(88),
		"form": uint(0),
		"type": "poison",
	},
	"Grimer-Alola": {
		"id":   uint(88),
		"form": uint(1),
		"type": "poison",
	},
	"Muk": {
		"id":   uint(89),
		"form": uint(0),
		"type": "poison",
	},
	"Muk-Alola": {
		"id":   uint(89),
		"form": uint(1),
		"type": "poison",
	},
	"Shellder": {
		"id":   uint(90),
		"form": uint(0),
		"type": "water",
	},
	"Cloyster": {
		"id":   uint(91),
		"form": uint(0),
		"type": "water",
	},
	"Gastly": {
		"id":   uint(92),
		"form": uint(0),
		"type": "ghost",
	},
	"Haunter": {
		"id":   uint(93),
		"form": uint(0),
		"type": "ghost",
	},
	"Gengar": {
		"id":   uint(94),
		"form": uint(0),
		"type": "ghost",
	},
	"Gengar-Mega": {
		"id":   uint(94),
		"form": uint(1),
		"type": "ghost",
	},
	"Gengar-Gmax": {
		"id":   uint(94),
		"form": uint(2),
		"type": "ghost",
	},
	"Onix": {
		"id":   uint(95),
		"form": uint(0),
		"type": "rock",
	},
	"Drowzee": {
		"id":   uint(96),
		"form": uint(0),
		"type": "psychic",
	},
	"Hypno": {
		"id":   uint(97),
		"form": uint(0),
		"type": "psychic",
	},
	"Krabby": {
		"id":   uint(98),
		"form": uint(0),
		"type": "water",
	},
	"Kingler": {
		"id":   uint(99),
		"form": uint(0),
		"type": "water",
	},
	"Kingler-Gmax": {
		"id":   uint(99),
		"form": uint(1),
		"type": "water",
	},
	"Voltorb": {
		"id":   uint(100),
		"form": uint(0),
		"type": "electric",
	},
	"Electrode": {
		"id":   uint(101),
		"form": uint(0),
		"type": "electric",
	},
	"Exeggcute": {
		"id":   uint(102),
		"form": uint(0),
		"type": "grass",
	},
	"Exeggutor": {
		"id":   uint(103),
		"form": uint(0),
		"type": "grass",
	},
	"Exeggutor-Alola": {
		"id":   uint(103),
		"form": uint(1),
		"type": "grass",
	},
	"Cubone": {
		"id":   uint(104),
		"form": uint(0),
		"type": "ground",
	},
	"Marowak": {
		"id":   uint(105),
		"form": uint(0),
		"type": "ground",
	},
	"Marowak-Alola": {
		"id":   uint(105),
		"form": uint(1),
		"type": "fire",
	},
	"Marowak-Alola-Totem": {
		"id":   uint(105),
		"form": uint(1),
		"type": "fire",
	},
	"Hitmonlee": {
		"id":   uint(106),
		"form": uint(0),
		"type": "fighting",
	},
	"Hitmonchan": {
		"id":   uint(107),
		"form": uint(0),
		"type": "fighting",
	},
	"Lickitung": {
		"id":   uint(108),
		"form": uint(0),
		"type": "normal",
	},
	"Koffing": {
		"id":   uint(109),
		"form": uint(0),
		"type": "poison",
	},
	"Weezing": {
		"id":   uint(110),
		"form": uint(0),
		"type": "poison",
	},
	"Weezing-Galar": {
		"id":   uint(110),
		"form": uint(1),
		"type": "poison",
	},
	"Rhyhorn": {
		"id":   uint(111),
		"form": uint(0),
		"type": "ground",
	},
	"Rhydon": {
		"id":   uint(112),
		"form": uint(0),
		"type": "ground",
	},
	"Chansey": {
		"id":   uint(113),
		"form": uint(0),
		"type": "normal",
	},
	"Tangela": {
		"id":   uint(114),
		"form": uint(0),
		"type": "grass",
	},
	"Kangaskhan": {
		"id":   uint(115),
		"form": uint(0),
		"type": "normal",
	},
	"Kangaskhan-Mega": {
		"id":   uint(115),
		"form": uint(1),
		"type": "normal",
	},
	"Horsea": {
		"id":   uint(116),
		"form": uint(0),
		"type": "water",
	},
	"Seadra": {
		"id":   uint(117),
		"form": uint(0),
		"type": "water",
	},
	"Goldeen": {
		"id":   uint(118),
		"form": uint(0),
		"type": "water",
	},
	"Seaking": {
		"id":   uint(119),
		"form": uint(0),
		"type": "water",
	},
	"Staryu": {
		"id":   uint(120),
		"form": uint(0),
		"type": "water",
	},
	"Starmie": {
		"id":   uint(121),
		"form": uint(0),
		"type": "water",
	},
	"Mr. Mime": {
		"id":   uint(122),
		"form": uint(0),
		"type": "psychic",
	},
	"Mr. Mime-Galar": {
		"id":   uint(122),
		"form": uint(1),
		"type": "ice",
	},
	"Scyther": {
		"id":   uint(123),
		"form": uint(0),
		"type": "bug",
	},
	"Jynx": {
		"id":   uint(124),
		"form": uint(0),
		"type": "ice",
	},
	"Electabuzz": {
		"id":   uint(125),
		"form": uint(0),
		"type": "electric",
	},
	"Magmar": {
		"id":   uint(126),
		"form": uint(0),
		"type": "fire",
	},
	"Pinsir": {
		"id":   uint(127),
		"form": uint(0),
		"type": "bug",
	},
	"Pinsir-Mega": {
		"id":   uint(127),
		"form": uint(1),
		"type": "bug",
	},
	"Tauros": {
		"id":   uint(128),
		"form": uint(0),
		"type": "normal",
	},
	"Magikarp": {
		"id":   uint(129),
		"form": uint(0),
		"type": "water",
	},
	"Gyarados": {
		"id":   uint(130),
		"form": uint(0),
		"type": "water",
	},
	"Gyarados-Mega": {
		"id":   uint(130),
		"form": uint(1),
		"type": "water",
	},
	"Lapras": {
		"id":   uint(131),
		"form": uint(0),
		"type": "water",
	},
	"Lapras-Gmax": {
		"id":   uint(131),
		"form": uint(1),
		"type": "water",
	},
	"Ditto": {
		"id":   uint(132),
		"form": uint(0),
		"type": "normal",
	},
	"Eevee": {
		"id":   uint(133),
		"form": uint(0),
		"type": "normal",
	},
	"Eevee-Gmax": {
		"id":   uint(133),
		"form": uint(1),
		"type": "normal",
	},
	"Vaporeon": {
		"id":   uint(134),
		"form": uint(0),
		"type": "water",
	},
	"Jolteon": {
		"id":   uint(135),
		"form": uint(0),
		"type": "electric",
	},
	"Flareon": {
		"id":   uint(136),
		"form": uint(0),
		"type": "fire",
	},
	"Porygon": {
		"id":   uint(137),
		"form": uint(0),
		"type": "normal",
	},
	"Omanyte": {
		"id":   uint(138),
		"form": uint(0),
		"type": "rock",
	},
	"Omastar": {
		"id":   uint(139),
		"form": uint(0),
		"type": "rock",
	},
	"Kabuto": {
		"id":   uint(140),
		"form": uint(0),
		"type": "rock",
	},
	"Kabutops": {
		"id":   uint(141),
		"form": uint(0),
		"type": "rock",
	},
	"Aerodactyl": {
		"id":   uint(142),
		"form": uint(0),
		"type": "rock",
	},
	"Aerodactyl-Mega": {
		"id":   uint(142),
		"form": uint(1),
		"type": "rock",
	},
	"Snorlax": {
		"id":   uint(143),
		"form": uint(0),
		"type": "normal",
	},
	"Snorlax-Gmax": {
		"id":   uint(143),
		"form": uint(1),
		"type": "normal",
	},
	"Articuno": {
		"id":   uint(144),
		"form": uint(0),
		"type": "ice",
	},
	"Articuno-Galar": {
		"id":   uint(144),
		"form": uint(1),
		"type": "psychic",
	},
	"Zapdos": {
		"id":   uint(145),
		"form": uint(0),
		"type": "electric",
	},
	"Zapdos-Galar": {
		"id":   uint(145),
		"form": uint(1),
		"type": "fighting",
	},
	"Moltres": {
		"id":   uint(146),
		"form": uint(0),
		"type": "fire",
	},
	"Moltres-Galar": {
		"id":   uint(146),
		"form": uint(1),
		"type": "dark",
	},
	"Dratini": {
		"id":   uint(147),
		"form": uint(0),
		"type": "dragon",
	},
	"Dragonair": {
		"id":   uint(148),
		"form": uint(0),
		"type": "dragon",
	},
	"Dragonite": {
		"id":   uint(149),
		"form": uint(0),
		"type": "dragon",
	},
	"Mewtwo": {
		"id":   uint(150),
		"form": uint(0),
		"type": "psychic",
	},
	"Mewtwo-Mega-X": {
		"id":   uint(150),
		"form": uint(1),
		"type": "psychic",
	},
	"Mewtwo-Mega-Y": {
		"id":   uint(150),
		"form": uint(2),
		"type": "psychic",
	},
	"Mew": {
		"id":   uint(151),
		"form": uint(0),
		"type": "psychic",
	},
	"Chikorita": {
		"id":   uint(152),
		"form": uint(0),
		"type": "grass",
	},
	"Bayleef": {
		"id":   uint(153),
		"form": uint(0),
		"type": "grass",
	},
	"Meganium": {
		"id":   uint(154),
		"form": uint(0),
		"type": "grass",
	},
	"Cyndaquil": {
		"id":   uint(155),
		"form": uint(0),
		"type": "fire",
	},
	"Quilava": {
		"id":   uint(156),
		"form": uint(0),
		"type": "fire",
	},
	"Typhlosion": {
		"id":   uint(157),
		"form": uint(0),
		"type": "fire",
	},
	"Totodile": {
		"id":   uint(158),
		"form": uint(0),
		"type": "water",
	},
	"Croconaw": {
		"id":   uint(159),
		"form": uint(0),
		"type": "water",
	},
	"Feraligatr": {
		"id":   uint(160),
		"form": uint(0),
		"type": "water",
	},
	"Sentret": {
		"id":   uint(161),
		"form": uint(0),
		"type": "normal",
	},
	"Furret": {
		"id":   uint(162),
		"form": uint(0),
		"type": "normal",
	},
	"Hoothoot": {
		"id":   uint(163),
		"form": uint(0),
		"type": "normal",
	},
	"Noctowl": {
		"id":   uint(164),
		"form": uint(0),
		"type": "normal",
	},
	"Ledyba": {
		"id":   uint(165),
		"form": uint(0),
		"type": "bug",
	},
	"Ledian": {
		"id":   uint(166),
		"form": uint(0),
		"type": "bug",
	},
	"Spinarak": {
		"id":   uint(167),
		"form": uint(0),
		"type": "bug",
	},
	"Ariados": {
		"id":   uint(168),
		"form": uint(0),
		"type": "bug",
	},
	"Crobat": {
		"id":   uint(169),
		"form": uint(0),
		"type": "poison",
	},
	"Chinchou": {
		"id":   uint(170),
		"form": uint(0),
		"type": "water",
	},
	"Lanturn": {
		"id":   uint(171),
		"form": uint(0),
		"type": "water",
	},
	"Pichu": {
		"id":   uint(172),
		"form": uint(0),
		"type": "electric",
	},
	"Cleffa": {
		"id":   uint(173),
		"form": uint(0),
		"type": "fairy",
	},
	"Igglybuff": {
		"id":   uint(174),
		"form": uint(0),
		"type": "normal",
	},
	"Togepi": {
		"id":   uint(175),
		"form": uint(0),
		"type": "fairy",
	},
	"Togetic": {
		"id":   uint(176),
		"form": uint(0),
		"type": "fairy",
	},
	"Natu": {
		"id":   uint(177),
		"form": uint(0),
		"type": "psychic",
	},
	"Xatu": {
		"id":   uint(178),
		"form": uint(0),
		"type": "psychic",
	},
	"Mareep": {
		"id":   uint(179),
		"form": uint(0),
		"type": "electric",
	},
	"Flaaffy": {
		"id":   uint(180),
		"form": uint(0),
		"type": "electric",
	},
	"Ampharos": {
		"id":   uint(181),
		"form": uint(0),
		"type": "electric",
	},
	"Ampharos-Mega": {
		"id":   uint(181),
		"form": uint(1),
		"type": "electric",
	},
	"Bellossom": {
		"id":   uint(182),
		"form": uint(0),
		"type": "grass",
	},
	"Marill": {
		"id":   uint(183),
		"form": uint(0),
		"type": "water",
	},
	"Azumarill": {
		"id":   uint(184),
		"form": uint(0),
		"type": "water",
	},
	"Sudowoodo": {
		"id":   uint(185),
		"form": uint(0),
		"type": "rock",
	},
	"Politoed": {
		"id":   uint(186),
		"form": uint(0),
		"type": "water",
	},
	"Hoppip": {
		"id":   uint(187),
		"form": uint(0),
		"type": "grass",
	},
	"Skiploom": {
		"id":   uint(188),
		"form": uint(0),
		"type": "grass",
	},
	"Jumpluff": {
		"id":   uint(189),
		"form": uint(0),
		"type": "grass",
	},
	"Aipom": {
		"id":   uint(190),
		"form": uint(0),
		"type": "normal",
	},
	"Sunkern": {
		"id":   uint(191),
		"form": uint(0),
		"type": "grass",
	},
	"Sunflora": {
		"id":   uint(192),
		"form": uint(0),
		"type": "grass",
	},
	"Yanma": {
		"id":   uint(193),
		"form": uint(0),
		"type": "bug",
	},
	"Wooper": {
		"id":   uint(194),
		"form": uint(0),
		"type": "water",
	},
	"Quagsire": {
		"id":   uint(195),
		"form": uint(0),
		"type": "water",
	},
	"Espeon": {
		"id":   uint(196),
		"form": uint(0),
		"type": "psychic",
	},
	"Umbreon": {
		"id":   uint(197),
		"form": uint(0),
		"type": "dark",
	},
	"Murkrow": {
		"id":   uint(198),
		"form": uint(0),
		"type": "dark",
	},
	"Slowking": {
		"id":   uint(199),
		"form": uint(0),
		"type": "water",
	},
	"Slowking-Galar": {
		"id":   uint(199),
		"form": uint(1),
		"type": "poison",
	},
	"Misdreavus": {
		"id":   uint(200),
		"form": uint(0),
		"type": "ghost",
	},
	"Unown": {
		"id":   uint(201),
		"form": uint(0),
		"type": "psychic",
	},
	"Wobbuffet": {
		"id":   uint(202),
		"form": uint(0),
		"type": "psychic",
	},
	"Girafarig": {
		"id":   uint(203),
		"form": uint(0),
		"type": "normal",
	},
	"Pineco": {
		"id":   uint(204),
		"form": uint(0),
		"type": "bug",
	},
	"Forretress": {
		"id":   uint(205),
		"form": uint(0),
		"type": "bug",
	},
	"Dunsparce": {
		"id":   uint(206),
		"form": uint(0),
		"type": "normal",
	},
	"Gligar": {
		"id":   uint(207),
		"form": uint(0),
		"type": "ground",
	},
	"Steelix": {
		"id":   uint(208),
		"form": uint(0),
		"type": "steel",
	},
	"Steelix-Mega": {
		"id":   uint(208),
		"form": uint(1),
		"type": "steel",
	},
	"Snubbull": {
		"id":   uint(209),
		"form": uint(0),
		"type": "fairy",
	},
	"Granbull": {
		"id":   uint(210),
		"form": uint(0),
		"type": "fairy",
	},
	"Qwilfish": {
		"id":   uint(211),
		"form": uint(0),
		"type": "water",
	},
	"Scizor": {
		"id":   uint(212),
		"form": uint(0),
		"type": "bug",
	},
	"Scizor-Mega": {
		"id":   uint(212),
		"form": uint(1),
		"type": "bug",
	},
	"Shuckle": {
		"id":   uint(213),
		"form": uint(0),
		"type": "bug",
	},
	"Heracross": {
		"id":   uint(214),
		"form": uint(0),
		"type": "bug",
	},
	"Heracross-Mega": {
		"id":   uint(214),
		"form": uint(1),
		"type": "bug",
	},
	"Sneasel": {
		"id":   uint(215),
		"form": uint(0),
		"type": "dark",
	},
	"Teddiursa": {
		"id":   uint(216),
		"form": uint(0),
		"type": "normal",
	},
	"Ursaring": {
		"id":   uint(217),
		"form": uint(0),
		"type": "normal",
	},
	"Slugma": {
		"id":   uint(218),
		"form": uint(0),
		"type": "fire",
	},
	"Magcargo": {
		"id":   uint(219),
		"form": uint(0),
		"type": "fire",
	},
	"Swinub": {
		"id":   uint(220),
		"form": uint(0),
		"type": "ice",
	},
	"Piloswine": {
		"id":   uint(221),
		"form": uint(0),
		"type": "ice",
	},
	"Corsola": {
		"id":   uint(222),
		"form": uint(0),
		"type": "water",
	},
	"Corsola-Galar": {
		"id":   uint(222),
		"form": uint(1),
		"type": "ghost",
	},
	"Remoraid": {
		"id":   uint(223),
		"form": uint(0),
		"type": "water",
	},
	"Octillery": {
		"id":   uint(224),
		"form": uint(0),
		"type": "water",
	},
	"Delibird": {
		"id":   uint(225),
		"form": uint(0),
		"type": "ice",
	},
	"Mantine": {
		"id":   uint(226),
		"form": uint(0),
		"type": "water",
	},
	"Skarmory": {
		"id":   uint(227),
		"form": uint(0),
		"type": "steel",
	},
	"Houndour": {
		"id":   uint(228),
		"form": uint(0),
		"type": "dark",
	},
	"Houndoom": {
		"id":   uint(229),
		"form": uint(0),
		"type": "dark",
	},
	"Houndoom-Mega": {
		"id":   uint(229),
		"form": uint(1),
		"type": "dark",
	},
	"Kingdra": {
		"id":   uint(230),
		"form": uint(0),
		"type": "water",
	},
	"Phanpy": {
		"id":   uint(231),
		"form": uint(0),
		"type": "ground",
	},
	"Donphan": {
		"id":   uint(232),
		"form": uint(0),
		"type": "ground",
	},
	"Porygon2": {
		"id":   uint(233),
		"form": uint(0),
		"type": "normal",
	},
	"Stantler": {
		"id":   uint(234),
		"form": uint(0),
		"type": "normal",
	},
	"Smeargle": {
		"id":   uint(235),
		"form": uint(0),
		"type": "normal",
	},
	"Tyrogue": {
		"id":   uint(236),
		"form": uint(0),
		"type": "fighting",
	},
	"Hitmontop": {
		"id":   uint(237),
		"form": uint(0),
		"type": "fighting",
	},
	"Smoochum": {
		"id":   uint(238),
		"form": uint(0),
		"type": "ice",
	},
	"Elekid": {
		"id":   uint(239),
		"form": uint(0),
		"type": "electric",
	},
	"Magby": {
		"id":   uint(240),
		"form": uint(0),
		"type": "fire",
	},
	"Miltank": {
		"id":   uint(241),
		"form": uint(0),
		"type": "normal",
	},
	"Blissey": {
		"id":   uint(242),
		"form": uint(0),
		"type": "normal",
	},
	"Raikou": {
		"id":   uint(243),
		"form": uint(0),
		"type": "electric",
	},
	"Entei": {
		"id":   uint(244),
		"form": uint(0),
		"type": "fire",
	},
	"Suicune": {
		"id":   uint(245),
		"form": uint(0),
		"type": "water",
	},
	"Larvitar": {
		"id":   uint(246),
		"form": uint(0),
		"type": "rock",
	},
	"Pupitar": {
		"id":   uint(247),
		"form": uint(0),
		"type": "rock",
	},
	"Tyranitar": {
		"id":   uint(248),
		"form": uint(0),
		"type": "rock",
	},
	"Tyranitar-Mega": {
		"id":   uint(248),
		"form": uint(1),
		"type": "rock",
	},
	"Lugia": {
		"id":   uint(249),
		"form": uint(0),
		"type": "psychic",
	},
	"Ho-Oh": {
		"id":   uint(250),
		"form": uint(0),
		"type": "fire",
	},
	"Celebi": {
		"id":   uint(251),
		"form": uint(0),
		"type": "psychic",
	},
	"Treecko": {
		"id":   uint(252),
		"form": uint(0),
		"type": "grass",
	},
	"Grovyle": {
		"id":   uint(253),
		"form": uint(0),
		"type": "grass",
	},
	"Sceptile": {
		"id":   uint(254),
		"form": uint(0),
		"type": "grass",
	},
	"Sceptile-Mega": {
		"id":   uint(254),
		"form": uint(1),
		"type": "grass",
	},
	"Torchic": {
		"id":   uint(255),
		"form": uint(0),
		"type": "fire",
	},
	"Combusken": {
		"id":   uint(256),
		"form": uint(0),
		"type": "fire",
	},
	"Blaziken": {
		"id":   uint(257),
		"form": uint(0),
		"type": "fire",
	},
	"Blaziken-Mega": {
		"id":   uint(257),
		"form": uint(1),
		"type": "fire",
	},
	"Mudkip": {
		"id":   uint(258),
		"form": uint(0),
		"type": "water",
	},
	"Marshtomp": {
		"id":   uint(259),
		"form": uint(0),
		"type": "water",
	},
	"Swampert": {
		"id":   uint(260),
		"form": uint(0),
		"type": "water",
	},
	"Swampert-Mega": {
		"id":   uint(260),
		"form": uint(1),
		"type": "water",
	},
	"Poochyena": {
		"id":   uint(261),
		"form": uint(0),
		"type": "dark",
	},
	"Mightyena": {
		"id":   uint(262),
		"form": uint(0),
		"type": "dark",
	},
	"Zigzagoon": {
		"id":   uint(263),
		"form": uint(0),
		"type": "normal",
	},
	"Zigzagoon-Galar": {
		"id":   uint(263),
		"form": uint(1),
		"type": "dark",
	},
	"Linoone": {
		"id":   uint(264),
		"form": uint(0),
		"type": "normal",
	},
	"Linoone-Galar": {
		"id":   uint(264),
		"form": uint(1),
		"type": "dark",
	},
	"Wurmple": {
		"id":   uint(265),
		"form": uint(0),
		"type": "bug",
	},
	"Silcoon": {
		"id":   uint(266),
		"form": uint(0),
		"type": "bug",
	},
	"Beautifly": {
		"id":   uint(267),
		"form": uint(0),
		"type": "bug",
	},
	"Cascoon": {
		"id":   uint(268),
		"form": uint(0),
		"type": "bug",
	},
	"Dustox": {
		"id":   uint(269),
		"form": uint(0),
		"type": "bug",
	},
	"Lotad": {
		"id":   uint(270),
		"form": uint(0),
		"type": "water",
	},
	"Lombre": {
		"id":   uint(271),
		"form": uint(0),
		"type": "water",
	},
	"Ludicolo": {
		"id":   uint(272),
		"form": uint(0),
		"type": "water",
	},
	"Seedot": {
		"id":   uint(273),
		"form": uint(0),
		"type": "grass",
	},
	"Nuzleaf": {
		"id":   uint(274),
		"form": uint(0),
		"type": "grass",
	},
	"Shiftry": {
		"id":   uint(275),
		"form": uint(0),
		"type": "grass",
	},
	"Taillow": {
		"id":   uint(276),
		"form": uint(0),
		"type": "normal",
	},
	"Swellow": {
		"id":   uint(277),
		"form": uint(0),
		"type": "normal",
	},
	"Wingull": {
		"id":   uint(278),
		"form": uint(0),
		"type": "water",
	},
	"Pelipper": {
		"id":   uint(279),
		"form": uint(0),
		"type": "water",
	},
	"Ralts": {
		"id":   uint(280),
		"form": uint(0),
		"type": "psychic",
	},
	"Kirlia": {
		"id":   uint(281),
		"form": uint(0),
		"type": "psychic",
	},
	"Gardevoir": {
		"id":   uint(282),
		"form": uint(0),
		"type": "psychic",
	},
	"Gardevoir-Mega": {
		"id":   uint(282),
		"form": uint(1),
		"type": "psychic",
	},
	"Surskit": {
		"id":   uint(283),
		"form": uint(0),
		"type": "bug",
	},
	"Masquerain": {
		"id":   uint(284),
		"form": uint(0),
		"type": "bug",
	},
	"Shroomish": {
		"id":   uint(285),
		"form": uint(0),
		"type": "grass",
	},
	"Breloom": {
		"id":   uint(286),
		"form": uint(0),
		"type": "grass",
	},
	"Slakoth": {
		"id":   uint(287),
		"form": uint(0),
		"type": "normal",
	},
	"Vigoroth": {
		"id":   uint(288),
		"form": uint(0),
		"type": "normal",
	},
	"Slaking": {
		"id":   uint(289),
		"form": uint(0),
		"type": "normal",
	},
	"Nincada": {
		"id":   uint(290),
		"form": uint(0),
		"type": "bug",
	},
	"Ninjask": {
		"id":   uint(291),
		"form": uint(0),
		"type": "bug",
	},
	"Shedinja": {
		"id":   uint(292),
		"form": uint(0),
		"type": "bug",
	},
	"Whismur": {
		"id":   uint(293),
		"form": uint(0),
		"type": "normal",
	},
	"Loudred": {
		"id":   uint(294),
		"form": uint(0),
		"type": "normal",
	},
	"Exploud": {
		"id":   uint(295),
		"form": uint(0),
		"type": "normal",
	},
	"Makuhita": {
		"id":   uint(296),
		"form": uint(0),
		"type": "fighting",
	},
	"Hariyama": {
		"id":   uint(297),
		"form": uint(0),
		"type": "fighting",
	},
	"Azurill": {
		"id":   uint(298),
		"form": uint(0),
		"type": "normal",
	},
	"Nosepass": {
		"id":   uint(299),
		"form": uint(0),
		"type": "rock",
	},
	"Skitty": {
		"id":   uint(300),
		"form": uint(0),
		"type": "normal",
	},
	"Delcatty": {
		"id":   uint(301),
		"form": uint(0),
		"type": "normal",
	},
	"Sableye": {
		"id":   uint(302),
		"form": uint(0),
		"type": "dark",
	},
	"Sableye-Mega": {
		"id":   uint(302),
		"form": uint(1),
		"type": "dark",
	},
	"Mawile": {
		"id":   uint(303),
		"form": uint(0),
		"type": "steel",
	},
	"Mawile-Mega": {
		"id":   uint(303),
		"form": uint(1),
		"type": "steel",
	},
	"Aron": {
		"id":   uint(304),
		"form": uint(0),
		"type": "steel",
	},
	"Lairon": {
		"id":   uint(305),
		"form": uint(0),
		"type": "steel",
	},
	"Aggron": {
		"id":   uint(306),
		"form": uint(0),
		"type": "steel",
	},
	"Aggron-Mega": {
		"id":   uint(306),
		"form": uint(1),
		"type": "steel",
	},
	"Meditite": {
		"id":   uint(307),
		"form": uint(0),
		"type": "fighting",
	},
	"Medicham": {
		"id":   uint(308),
		"form": uint(0),
		"type": "fighting",
	},
	"Medicham-Mega": {
		"id":   uint(308),
		"form": uint(1),
		"type": "fighting",
	},
	"Electrike": {
		"id":   uint(309),
		"form": uint(0),
		"type": "electric",
	},
	"Manectric": {
		"id":   uint(310),
		"form": uint(0),
		"type": "electric",
	},
	"Manectric-Mega": {
		"id":   uint(310),
		"form": uint(1),
		"type": "electric",
	},
	"Plusle": {
		"id":   uint(311),
		"form": uint(0),
		"type": "electric",
	},
	"Minun": {
		"id":   uint(312),
		"form": uint(0),
		"type": "electric",
	},
	"Volbeat": {
		"id":   uint(313),
		"form": uint(0),
		"type": "bug",
	},
	"Illumise": {
		"id":   uint(314),
		"form": uint(0),
		"type": "bug",
	},
	"Roselia": {
		"id":   uint(315),
		"form": uint(0),
		"type": "grass",
	},
	"Gulpin": {
		"id":   uint(316),
		"form": uint(0),
		"type": "poison",
	},
	"Swalot": {
		"id":   uint(317),
		"form": uint(0),
		"type": "poison",
	},
	"Carvanha": {
		"id":   uint(318),
		"form": uint(0),
		"type": "water",
	},
	"Sharpedo": {
		"id":   uint(319),
		"form": uint(0),
		"type": "water",
	},
	"Sharpedo-Mega": {
		"id":   uint(319),
		"form": uint(1),
		"type": "water",
	},
	"Wailmer": {
		"id":   uint(320),
		"form": uint(0),
		"type": "water",
	},
	"Wailord": {
		"id":   uint(321),
		"form": uint(0),
		"type": "water",
	},
	"Numel": {
		"id":   uint(322),
		"form": uint(0),
		"type": "fire",
	},
	"Camerupt": {
		"id":   uint(323),
		"form": uint(0),
		"type": "fire",
	},
	"Camerupt-Mega": {
		"id":   uint(323),
		"form": uint(1),
		"type": "fire",
	},
	"Torkoal": {
		"id":   uint(324),
		"form": uint(0),
		"type": "fire",
	},
	"Spoink": {
		"id":   uint(325),
		"form": uint(0),
		"type": "psychic",
	},
	"Grumpig": {
		"id":   uint(326),
		"form": uint(0),
		"type": "psychic",
	},
	"Spinda": {
		"id":   uint(327),
		"form": uint(0),
		"type": "normal",
	},
	"Trapinch": {
		"id":   uint(328),
		"form": uint(0),
		"type": "ground",
	},
	"Vibrava": {
		"id":   uint(329),
		"form": uint(0),
		"type": "ground",
	},
	"Flygon": {
		"id":   uint(330),
		"form": uint(0),
		"type": "ground",
	},
	"Cacnea": {
		"id":   uint(331),
		"form": uint(0),
		"type": "grass",
	},
	"Cacturne": {
		"id":   uint(332),
		"form": uint(0),
		"type": "grass",
	},
	"Swablu": {
		"id":   uint(333),
		"form": uint(0),
		"type": "normal",
	},
	"Altaria": {
		"id":   uint(334),
		"form": uint(0),
		"type": "dragon",
	},
	"Altaria-Mega": {
		"id":   uint(334),
		"form": uint(1),
		"type": "dragon",
	},
	"Zangoose": {
		"id":   uint(335),
		"form": uint(0),
		"type": "normal",
	},
	"Seviper": {
		"id":   uint(336),
		"form": uint(0),
		"type": "poison",
	},
	"Lunatone": {
		"id":   uint(337),
		"form": uint(0),
		"type": "rock",
	},
	"Solrock": {
		"id":   uint(338),
		"form": uint(0),
		"type": "rock",
	},
	"Barboach": {
		"id":   uint(339),
		"form": uint(0),
		"type": "water",
	},
	"Whiscash": {
		"id":   uint(340),
		"form": uint(0),
		"type": "water",
	},
	"Corphish": {
		"id":   uint(341),
		"form": uint(0),
		"type": "water",
	},
	"Crawdaunt": {
		"id":   uint(342),
		"form": uint(0),
		"type": "water",
	},
	"Baltoy": {
		"id":   uint(343),
		"form": uint(0),
		"type": "ground",
	},
	"Claydol": {
		"id":   uint(344),
		"form": uint(0),
		"type": "ground",
	},
	"Lileep": {
		"id":   uint(345),
		"form": uint(0),
		"type": "rock",
	},
	"Cradily": {
		"id":   uint(346),
		"form": uint(0),
		"type": "rock",
	},
	"Anorith": {
		"id":   uint(347),
		"form": uint(0),
		"type": "rock",
	},
	"Armaldo": {
		"id":   uint(348),
		"form": uint(0),
		"type": "rock",
	},
	"Feebas": {
		"id":   uint(349),
		"form": uint(0),
		"type": "water",
	},
	"Milotic": {
		"id":   uint(350),
		"form": uint(0),
		"type": "water",
	},
	"Castform": {
		"id":   uint(351),
		"form": uint(0),
		"type": "normal",
	},
	"Castform-Sunny": {
		"id":   uint(351),
		"form": uint(1),
		"type": "fire",
	},
	"Castform-Rainy": {
		"id":   uint(351),
		"form": uint(2),
		"type": "water",
	},
	"Castform-Snowy": {
		"id":   uint(351),
		"form": uint(3),
		"type": "ice",
	},
	"Kecleon": {
		"id":   uint(352),
		"form": uint(0),
		"type": "normal",
	},
	"Shuppet": {
		"id":   uint(353),
		"form": uint(0),
		"type": "ghost",
	},
	"Banette": {
		"id":   uint(354),
		"form": uint(0),
		"type": "ghost",
	},
	"Banette-Mega": {
		"id":   uint(354),
		"form": uint(1),
		"type": "ghost",
	},
	"Duskull": {
		"id":   uint(355),
		"form": uint(0),
		"type": "ghost",
	},
	"Dusclops": {
		"id":   uint(356),
		"form": uint(0),
		"type": "ghost",
	},
	"Tropius": {
		"id":   uint(357),
		"form": uint(0),
		"type": "grass",
	},
	"Chimecho": {
		"id":   uint(358),
		"form": uint(0),
		"type": "psychic",
	},
	"Absol": {
		"id":   uint(359),
		"form": uint(0),
		"type": "dark",
	},
	"Absol-Mega": {
		"id":   uint(359),
		"form": uint(1),
		"type": "dark",
	},
	"Wynaut": {
		"id":   uint(360),
		"form": uint(0),
		"type": "psychic",
	},
	"Snorunt": {
		"id":   uint(361),
		"form": uint(0),
		"type": "ice",
	},
	"Glalie": {
		"id":   uint(362),
		"form": uint(0),
		"type": "ice",
	},
	"Glalie-Mega": {
		"id":   uint(362),
		"form": uint(1),
		"type": "ice",
	},
	"Spheal": {
		"id":   uint(363),
		"form": uint(0),
		"type": "ice",
	},
	"Sealeo": {
		"id":   uint(364),
		"form": uint(0),
		"type": "ice",
	},
	"Walrein": {
		"id":   uint(365),
		"form": uint(0),
		"type": "ice",
	},
	"Clamperl": {
		"id":   uint(366),
		"form": uint(0),
		"type": "water",
	},
	"Huntail": {
		"id":   uint(367),
		"form": uint(0),
		"type": "water",
	},
	"Gorebyss": {
		"id":   uint(368),
		"form": uint(0),
		"type": "water",
	},
	"Relicanth": {
		"id":   uint(369),
		"form": uint(0),
		"type": "water",
	},
	"Luvdisc": {
		"id":   uint(370),
		"form": uint(0),
		"type": "water",
	},
	"Bagon": {
		"id":   uint(371),
		"form": uint(0),
		"type": "dragon",
	},
	"Shelgon": {
		"id":   uint(372),
		"form": uint(0),
		"type": "dragon",
	},
	"Salamence": {
		"id":   uint(373),
		"form": uint(0),
		"type": "dragon",
	},
	"Salamence-Mega": {
		"id":   uint(373),
		"form": uint(1),
		"type": "dragon",
	},
	"Beldum": {
		"id":   uint(374),
		"form": uint(0),
		"type": "steel",
	},
	"Metang": {
		"id":   uint(375),
		"form": uint(0),
		"type": "steel",
	},
	"Metagross": {
		"id":   uint(376),
		"form": uint(0),
		"type": "steel",
	},
	"Metagross-Mega": {
		"id":   uint(376),
		"form": uint(1),
		"type": "steel",
	},
	"Regirock": {
		"id":   uint(377),
		"form": uint(0),
		"type": "rock",
	},
	"Regice": {
		"id":   uint(378),
		"form": uint(0),
		"type": "ice",
	},
	"Registeel": {
		"id":   uint(379),
		"form": uint(0),
		"type": "steel",
	},
	"Latias": {
		"id":   uint(380),
		"form": uint(0),
		"type": "dragon",
	},
	"Latias-Mega": {
		"id":   uint(380),
		"form": uint(1),
		"type": "dragon",
	},
	"Latios": {
		"id":   uint(381),
		"form": uint(0),
		"type": "dragon",
	},
	"Latios-Mega": {
		"id":   uint(381),
		"form": uint(1),
		"type": "dragon",
	},
	"Kyogre": {
		"id":   uint(382),
		"form": uint(0),
		"type": "water",
	},
	"Kyogre-Primal": {
		"id":   uint(382),
		"form": uint(1),
		"type": "water",
	},
	"Groudon": {
		"id":   uint(383),
		"form": uint(0),
		"type": "ground",
	},
	"Groudon-Primal": {
		"id":   uint(383),
		"form": uint(1),
		"type": "ground",
	},
	"Rayquaza": {
		"id":   uint(384),
		"form": uint(0),
		"type": "dragon",
	},
	"Rayquaza-Mega": {
		"id":   uint(384),
		"form": uint(1),
		"type": "dragon",
	},
	"Jirachi": {
		"id":   uint(385),
		"form": uint(0),
		"type": "steel",
	},
	"Deoxys": {
		"id":   uint(386),
		"form": uint(0),
		"type": "psychic",
	},
	"Deoxys-Attack": {
		"id":   uint(386),
		"form": uint(1),
		"type": "psychic",
	},
	"Deoxys-Defense": {
		"id":   uint(386),
		"form": uint(2),
		"type": "psychic",
	},
	"Deoxys-Speed": {
		"id":   uint(386),
		"form": uint(3),
		"type": "psychic",
	},
	"Turtwig": {
		"id":   uint(387),
		"form": uint(0),
		"type": "grass",
	},
	"Grotle": {
		"id":   uint(388),
		"form": uint(0),
		"type": "grass",
	},
	"Torterra": {
		"id":   uint(389),
		"form": uint(0),
		"type": "grass",
	},
	"Chimchar": {
		"id":   uint(390),
		"form": uint(0),
		"type": "fire",
	},
	"Monferno": {
		"id":   uint(391),
		"form": uint(0),
		"type": "fire",
	},
	"Infernape": {
		"id":   uint(392),
		"form": uint(0),
		"type": "fire",
	},
	"Piplup": {
		"id":   uint(393),
		"form": uint(0),
		"type": "water",
	},
	"Prinplup": {
		"id":   uint(394),
		"form": uint(0),
		"type": "water",
	},
	"Empoleon": {
		"id":   uint(395),
		"form": uint(0),
		"type": "water",
	},
	"Starly": {
		"id":   uint(396),
		"form": uint(0),
		"type": "normal",
	},
	"Staravia": {
		"id":   uint(397),
		"form": uint(0),
		"type": "normal",
	},
	"Staraptor": {
		"id":   uint(398),
		"form": uint(0),
		"type": "normal",
	},
	"Bidoof": {
		"id":   uint(399),
		"form": uint(0),
		"type": "normal",
	},
	"Bibarel": {
		"id":   uint(400),
		"form": uint(0),
		"type": "normal",
	},
	"Kricketot": {
		"id":   uint(401),
		"form": uint(0),
		"type": "bug",
	},
	"Kricketune": {
		"id":   uint(402),
		"form": uint(0),
		"type": "bug",
	},
	"Shinx": {
		"id":   uint(403),
		"form": uint(0),
		"type": "electric",
	},
	"Luxio": {
		"id":   uint(404),
		"form": uint(0),
		"type": "electric",
	},
	"Luxray": {
		"id":   uint(405),
		"form": uint(0),
		"type": "electric",
	},
	"Budew": {
		"id":   uint(406),
		"form": uint(0),
		"type": "grass",
	},
	"Roserade": {
		"id":   uint(407),
		"form": uint(0),
		"type": "grass",
	},
	"Cranidos": {
		"id":   uint(408),
		"form": uint(0),
		"type": "rock",
	},
	"Rampardos": {
		"id":   uint(409),
		"form": uint(0),
		"type": "rock",
	},
	"Shieldon": {
		"id":   uint(410),
		"form": uint(0),
		"type": "rock",
	},
	"Bastiodon": {
		"id":   uint(411),
		"form": uint(0),
		"type": "rock",
	},
	"Burmy": {
		"id":   uint(412),
		"form": uint(0),
		"type": "bug",
	},
	"Burmy-Sandy": {
		"id":   uint(412),
		"form": uint(1),
		"type": "bug",
	},
	"Burmy-Trash": {
		"id":   uint(412),
		"form": uint(2),
		"type": "bug",
	},
	"Wormadam": {
		"id":   uint(413),
		"form": uint(0),
		"type": "bug",
	},
	"Wormadam-Sandy": {
		"id":   uint(413),
		"form": uint(1),
		"type": "bug",
	},
	"Wormadam-Trash": {
		"id":   uint(413),
		"form": uint(2),
		"type": "bug",
	},
	"Mothim": {
		"id":   uint(414),
		"form": uint(0),
		"type": "bug",
	},
	"Combee": {
		"id":   uint(415),
		"form": uint(0),
		"type": "bug",
	},
	"Vespiquen": {
		"id":   uint(416),
		"form": uint(0),
		"type": "bug",
	},
	"Pachirisu": {
		"id":   uint(417),
		"form": uint(0),
		"type": "electric",
	},
	"Buizel": {
		"id":   uint(418),
		"form": uint(0),
		"type": "water",
	},
	"Floatzel": {
		"id":   uint(419),
		"form": uint(0),
		"type": "water",
	},
	"Cherubi": {
		"id":   uint(420),
		"form": uint(0),
		"type": "grass",
	},
	"Cherrim": {
		"id":   uint(421),
		"form": uint(0),
		"type": "grass",
	},
	"Cherrim-Sunshine": {
		"id":   uint(421),
		"form": uint(1),
		"type": "grass",
	},
	"Shellos": {
		"id":   uint(422),
		"form": uint(0),
		"type": "water",
	},
	"Shellos-East": {
		"id":   uint(422),
		"form": uint(1),
		"type": "water",
	},
	"Gastrodon": {
		"id":   uint(423),
		"form": uint(0),
		"type": "water",
	},
	"Gastrodon-East": {
		"id":   uint(423),
		"form": uint(1),
		"type": "water",
	},
	"Ambipom": {
		"id":   uint(424),
		"form": uint(0),
		"type": "normal",
	},
	"Drifloon": {
		"id":   uint(425),
		"form": uint(0),
		"type": "ghost",
	},
	"Drifblim": {
		"id":   uint(426),
		"form": uint(0),
		"type": "ghost",
	},
	"Buneary": {
		"id":   uint(427),
		"form": uint(0),
		"type": "normal",
	},
	"Lopunny": {
		"id":   uint(428),
		"form": uint(0),
		"type": "normal",
	},
	"Lopunny-Mega": {
		"id":   uint(428),
		"form": uint(1),
		"type": "normal",
	},
	"Mismagius": {
		"id":   uint(429),
		"form": uint(0),
		"type": "ghost",
	},
	"Honchkrow": {
		"id":   uint(430),
		"form": uint(0),
		"type": "dark",
	},
	"Glameow": {
		"id":   uint(431),
		"form": uint(0),
		"type": "normal",
	},
	"Purugly": {
		"id":   uint(432),
		"form": uint(0),
		"type": "normal",
	},
	"Chingling": {
		"id":   uint(433),
		"form": uint(0),
		"type": "psychic",
	},
	"Stunky": {
		"id":   uint(434),
		"form": uint(0),
		"type": "poison",
	},
	"Skuntank": {
		"id":   uint(435),
		"form": uint(0),
		"type": "poison",
	},
	"Bronzor": {
		"id":   uint(436),
		"form": uint(0),
		"type": "steel",
	},
	"Bronzong": {
		"id":   uint(437),
		"form": uint(0),
		"type": "steel",
	},
	"Bonsly": {
		"id":   uint(438),
		"form": uint(0),
		"type": "rock",
	},
	"Mime Jr.": {
		"id":   uint(439),
		"form": uint(0),
		"type": "psychic",
	},
	"Happiny": {
		"id":   uint(440),
		"form": uint(0),
		"type": "normal",
	},
	"Chatot": {
		"id":   uint(441),
		"form": uint(0),
		"type": "normal",
	},
	"Spiritomb": {
		"id":   uint(442),
		"form": uint(0),
		"type": "ghost",
	},
	"Gible": {
		"id":   uint(443),
		"form": uint(0),
		"type": "dragon",
	},
	"Gabite": {
		"id":   uint(444),
		"form": uint(0),
		"type": "dragon",
	},
	"Garchomp": {
		"id":   uint(445),
		"form": uint(0),
		"type": "dragon",
	},
	"Garchomp-Mega": {
		"id":   uint(445),
		"form": uint(1),
		"type": "dragon",
	},
	"Munchlax": {
		"id":   uint(446),
		"form": uint(0),
		"type": "normal",
	},
	"Riolu": {
		"id":   uint(447),
		"form": uint(0),
		"type": "fighting",
	},
	"Lucario": {
		"id":   uint(448),
		"form": uint(0),
		"type": "fighting",
	},
	"Lucario-Mega": {
		"id":   uint(448),
		"form": uint(1),
		"type": "fighting",
	},
	"Hippopotas": {
		"id":   uint(449),
		"form": uint(0),
		"type": "ground",
	},
	"Hippowdon": {
		"id":   uint(450),
		"form": uint(0),
		"type": "ground",
	},
	"Skorupi": {
		"id":   uint(451),
		"form": uint(0),
		"type": "poison",
	},
	"Drapion": {
		"id":   uint(452),
		"form": uint(0),
		"type": "poison",
	},
	"Croagunk": {
		"id":   uint(453),
		"form": uint(0),
		"type": "poison",
	},
	"Toxicroak": {
		"id":   uint(454),
		"form": uint(0),
		"type": "poison",
	},
	"Carnivine": {
		"id":   uint(455),
		"form": uint(0),
		"type": "grass",
	},
	"Finneon": {
		"id":   uint(456),
		"form": uint(0),
		"type": "water",
	},
	"Lumineon": {
		"id":   uint(457),
		"form": uint(0),
		"type": "water",
	},
	"Mantyke": {
		"id":   uint(458),
		"form": uint(0),
		"type": "water",
	},
	"Snover": {
		"id":   uint(459),
		"form": uint(0),
		"type": "grass",
	},
	"Abomasnow": {
		"id":   uint(460),
		"form": uint(0),
		"type": "grass",
	},
	"Abomasnow-Mega": {
		"id":   uint(460),
		"form": uint(1),
		"type": "grass",
	},
	"Weavile": {
		"id":   uint(461),
		"form": uint(0),
		"type": "dark",
	},
	"Magnezone": {
		"id":   uint(462),
		"form": uint(0),
		"type": "electric",
	},
	"Lickilicky": {
		"id":   uint(463),
		"form": uint(0),
		"type": "normal",
	},
	"Rhyperior": {
		"id":   uint(464),
		"form": uint(0),
		"type": "ground",
	},
	"Tangrowth": {
		"id":   uint(465),
		"form": uint(0),
		"type": "grass",
	},
	"Electivire": {
		"id":   uint(466),
		"form": uint(0),
		"type": "electric",
	},
	"Magmortar": {
		"id":   uint(467),
		"form": uint(0),
		"type": "fire",
	},
	"Togekiss": {
		"id":   uint(468),
		"form": uint(0),
		"type": "fairy",
	},
	"Yanmega": {
		"id":   uint(469),
		"form": uint(0),
		"type": "bug",
	},
	"Leafeon": {
		"id":   uint(470),
		"form": uint(0),
		"type": "grass",
	},
	"Glaceon": {
		"id":   uint(471),
		"form": uint(0),
		"type": "ice",
	},
	"Gliscor": {
		"id":   uint(472),
		"form": uint(0),
		"type": "ground",
	},
	"Mamoswine": {
		"id":   uint(473),
		"form": uint(0),
		"type": "ice",
	},
	"Porygon-Z": {
		"id":   uint(474),
		"form": uint(0),
		"type": "normal",
	},
	"Gallade": {
		"id":   uint(475),
		"form": uint(0),
		"type": "psychic",
	},
	"Gallade-Mega": {
		"id":   uint(475),
		"form": uint(1),
		"type": "psychic",
	},
	"Probopass": {
		"id":   uint(476),
		"form": uint(0),
		"type": "rock",
	},
	"Dusknoir": {
		"id":   uint(477),
		"form": uint(0),
		"type": "ghost",
	},
	"Froslass": {
		"id":   uint(478),
		"form": uint(0),
		"type": "ice",
	},
	"Rotom": {
		"id":   uint(479),
		"form": uint(0),
		"type": "electric",
	},
	"Rotom-Heat": {
		"id":   uint(479),
		"form": uint(1),
		"type": "electric",
	},
	"Rotom-Wash": {
		"id":   uint(479),
		"form": uint(2),
		"type": "electric",
	},
	"Rotom-Frost": {
		"id":   uint(479),
		"form": uint(3),
		"type": "electric",
	},
	"Rotom-Fan": {
		"id":   uint(479),
		"form": uint(4),
		"type": "electric",
	},
	"Rotom-Mow": {
		"id":   uint(479),
		"form": uint(5),
		"type": "electric",
	},
	"Uxie": {
		"id":   uint(480),
		"form": uint(0),
		"type": "psychic",
	},
	"Mesprit": {
		"id":   uint(481),
		"form": uint(0),
		"type": "psychic",
	},
	"Azelf": {
		"id":   uint(482),
		"form": uint(0),
		"type": "psychic",
	},
	"Dialga": {
		"id":   uint(483),
		"form": uint(0),
		"type": "steel",
	},
	"Palkia": {
		"id":   uint(484),
		"form": uint(0),
		"type": "water",
	},
	"Heatran": {
		"id":   uint(485),
		"form": uint(0),
		"type": "fire",
	},
	"Regigigas": {
		"id":   uint(486),
		"form": uint(0),
		"type": "normal",
	},
	"Giratina": {
		"id":   uint(487),
		"form": uint(0),
		"type": "ghost",
	},
	"Giratina-Origin": {
		"id":   uint(487),
		"form": uint(1),
		"type": "ghost",
	},
	"Cresselia": {
		"id":   uint(488),
		"form": uint(0),
		"type": "psychic",
	},
	"Phione": {
		"id":   uint(489),
		"form": uint(0),
		"type": "water",
	},
	"Manaphy": {
		"id":   uint(490),
		"form": uint(0),
		"type": "water",
	},
	"Darkrai": {
		"id":   uint(491),
		"form": uint(0),
		"type": "dark",
	},
	"Shaymin": {
		"id":   uint(492),
		"form": uint(0),
		"type": "grass",
	},
	"Shaymin-Sky": {
		"id":   uint(492),
		"form": uint(1),
		"type": "grass",
	},
	"Arceus-Fairy": {
		"id":   uint(493),
		"form": uint(17),
		"type": "fairy",
	},
	"Arceus-Dark": {
		"id":   uint(493),
		"form": uint(16),
		"type": "dark",
	},
	"Arceus-Dragon": {
		"id":   uint(493),
		"form": uint(15),
		"type": "dragon",
	},
	"Arceus-Ice": {
		"id":   uint(493),
		"form": uint(14),
		"type": "ice",
	},
	"Arceus-Psychic": {
		"id":   uint(493),
		"form": uint(13),
		"type": "psychic",
	},
	"Arceus-Electric": {
		"id":   uint(493),
		"form": uint(12),
		"type": "electric",
	},
	"Arceus-Grass": {
		"id":   uint(493),
		"form": uint(11),
		"type": "grass",
	},
	"Arceus-Water": {
		"id":   uint(493),
		"form": uint(10),
		"type": "water",
	},
	"Arceus-Fire": {
		"id":   uint(493),
		"form": uint(9),
		"type": "fire",
	},
	"Arceus-Steel": {
		"id":   uint(493),
		"form": uint(8),
		"type": "steel",
	},
	"Arceus-Ghost": {
		"id":   uint(493),
		"form": uint(7),
		"type": "ghost",
	},
	"Arceus-Bug": {
		"id":   uint(493),
		"form": uint(6),
		"type": "bug",
	},
	"Arceus-Rock": {
		"id":   uint(493),
		"form": uint(5),
		"type": "rock",
	},
	"Arceus-Ground": {
		"id":   uint(493),
		"form": uint(4),
		"type": "ground",
	},
	"Arceus-Poison": {
		"id":   uint(493),
		"form": uint(3),
		"type": "poison",
	},
	"Arceus-Flying": {
		"id":   uint(493),
		"form": uint(2),
		"type": "flying",
	},
	"Arceus-Fighting": {
		"id":   uint(493),
		"form": uint(1),
		"type": "fighting",
	},
	"Arceus": {
		"id":   uint(493),
		"form": uint(0),
		"type": "normal",
	},
	"Victini": {
		"id":   uint(494),
		"form": uint(0),
		"type": "psychic",
	},
	"Snivy": {
		"id":   uint(495),
		"form": uint(0),
		"type": "grass",
	},
	"Servine": {
		"id":   uint(496),
		"form": uint(0),
		"type": "grass",
	},
	"Serperior": {
		"id":   uint(497),
		"form": uint(0),
		"type": "grass",
	},
	"Tepig": {
		"id":   uint(498),
		"form": uint(0),
		"type": "fire",
	},
	"Pignite": {
		"id":   uint(499),
		"form": uint(0),
		"type": "fire",
	},
	"Emboar": {
		"id":   uint(500),
		"form": uint(0),
		"type": "fire",
	},
	"Oshawott": {
		"id":   uint(501),
		"form": uint(0),
		"type": "water",
	},
	"Dewott": {
		"id":   uint(502),
		"form": uint(0),
		"type": "water",
	},
	"Samurott": {
		"id":   uint(503),
		"form": uint(0),
		"type": "water",
	},
	"Patrat": {
		"id":   uint(504),
		"form": uint(0),
		"type": "normal",
	},
	"Watchog": {
		"id":   uint(505),
		"form": uint(0),
		"type": "normal",
	},
	"Lillipup": {
		"id":   uint(506),
		"form": uint(0),
		"type": "normal",
	},
	"Herdier": {
		"id":   uint(507),
		"form": uint(0),
		"type": "normal",
	},
	"Stoutland": {
		"id":   uint(508),
		"form": uint(0),
		"type": "normal",
	},
	"Purrloin": {
		"id":   uint(509),
		"form": uint(0),
		"type": "dark",
	},
	"Liepard": {
		"id":   uint(510),
		"form": uint(0),
		"type": "dark",
	},
	"Pansage": {
		"id":   uint(511),
		"form": uint(0),
		"type": "grass",
	},
	"Simisage": {
		"id":   uint(512),
		"form": uint(0),
		"type": "grass",
	},
	"Pansear": {
		"id":   uint(513),
		"form": uint(0),
		"type": "fire",
	},
	"Simisear": {
		"id":   uint(514),
		"form": uint(0),
		"type": "fire",
	},
	"Panpour": {
		"id":   uint(515),
		"form": uint(0),
		"type": "water",
	},
	"Simipour": {
		"id":   uint(516),
		"form": uint(0),
		"type": "water",
	},
	"Munna": {
		"id":   uint(517),
		"form": uint(0),
		"type": "psychic",
	},
	"Musharna": {
		"id":   uint(518),
		"form": uint(0),
		"type": "psychic",
	},
	"Pidove": {
		"id":   uint(519),
		"form": uint(0),
		"type": "normal",
	},
	"Tranquill": {
		"id":   uint(520),
		"form": uint(0),
		"type": "normal",
	},
	"Unfezant": {
		"id":   uint(521),
		"form": uint(0),
		"type": "normal",
	},
	"Blitzle": {
		"id":   uint(522),
		"form": uint(0),
		"type": "electric",
	},
	"Zebstrika": {
		"id":   uint(523),
		"form": uint(0),
		"type": "electric",
	},
	"Roggenrola": {
		"id":   uint(524),
		"form": uint(0),
		"type": "rock",
	},
	"Boldore": {
		"id":   uint(525),
		"form": uint(0),
		"type": "rock",
	},
	"Gigalith": {
		"id":   uint(526),
		"form": uint(0),
		"type": "rock",
	},
	"Woobat": {
		"id":   uint(527),
		"form": uint(0),
		"type": "psychic",
	},
	"Swoobat": {
		"id":   uint(528),
		"form": uint(0),
		"type": "psychic",
	},
	"Drilbur": {
		"id":   uint(529),
		"form": uint(0),
		"type": "ground",
	},
	"Excadrill": {
		"id":   uint(530),
		"form": uint(0),
		"type": "ground",
	},
	"Audino": {
		"id":   uint(531),
		"form": uint(0),
		"type": "normal",
	},
	"Audino-Mega": {
		"id":   uint(531),
		"form": uint(1),
		"type": "normal",
	},
	"Timburr": {
		"id":   uint(532),
		"form": uint(0),
		"type": "fighting",
	},
	"Gurdurr": {
		"id":   uint(533),
		"form": uint(0),
		"type": "fighting",
	},
	"Conkeldurr": {
		"id":   uint(534),
		"form": uint(0),
		"type": "fighting",
	},
	"Tympole": {
		"id":   uint(535),
		"form": uint(0),
		"type": "water",
	},
	"Palpitoad": {
		"id":   uint(536),
		"form": uint(0),
		"type": "water",
	},
	"Seismitoad": {
		"id":   uint(537),
		"form": uint(0),
		"type": "water",
	},
	"Throh": {
		"id":   uint(538),
		"form": uint(0),
		"type": "fighting",
	},
	"Sawk": {
		"id":   uint(539),
		"form": uint(0),
		"type": "fighting",
	},
	"Sewaddle": {
		"id":   uint(540),
		"form": uint(0),
		"type": "bug",
	},
	"Swadloon": {
		"id":   uint(541),
		"form": uint(0),
		"type": "bug",
	},
	"Leavanny": {
		"id":   uint(542),
		"form": uint(0),
		"type": "bug",
	},
	"Venipede": {
		"id":   uint(543),
		"form": uint(0),
		"type": "bug",
	},
	"Whirlipede": {
		"id":   uint(544),
		"form": uint(0),
		"type": "bug",
	},
	"Scolipede": {
		"id":   uint(545),
		"form": uint(0),
		"type": "bug",
	},
	"Cottonee": {
		"id":   uint(546),
		"form": uint(0),
		"type": "grass",
	},
	"Whimsicott": {
		"id":   uint(547),
		"form": uint(0),
		"type": "grass",
	},
	"Petilil": {
		"id":   uint(548),
		"form": uint(0),
		"type": "grass",
	},
	"Lilligant": {
		"id":   uint(549),
		"form": uint(0),
		"type": "grass",
	},
	"Basculin": {
		"id":   uint(550),
		"form": uint(0),
		"type": "water",
	},
	"Basculin-Blue-Striped": {
		"id":   uint(550),
		"form": uint(1),
		"type": "water",
	},
	"Sandile": {
		"id":   uint(551),
		"form": uint(0),
		"type": "ground",
	},
	"Krokorok": {
		"id":   uint(552),
		"form": uint(0),
		"type": "ground",
	},
	"Krookodile": {
		"id":   uint(553),
		"form": uint(0),
		"type": "ground",
	},
	"Darumaka": {
		"id":   uint(554),
		"form": uint(0),
		"type": "fire",
	},
	"Darumaka-Galar": {
		"id":   uint(554),
		"form": uint(1),
		"type": "ice",
	},
	"Darmanitan": {
		"id":   uint(555),
		"form": uint(0),
		"type": "fire",
	},
	"Darmanitan-Zen": {
		"id":   uint(555),
		"form": uint(1),
		"type": "fire",
	},
	"Darmanitan-Galar": {
		"id":   uint(555),
		"form": uint(2),
		"type": "ice",
	},
	"Darmanitan-Galar-Zen": {
		"id":   uint(555),
		"form": uint(3),
		"type": "ice",
	},
	"Maractus": {
		"id":   uint(556),
		"form": uint(0),
		"type": "grass",
	},
	"Dwebble": {
		"id":   uint(557),
		"form": uint(0),
		"type": "bug",
	},
	"Crustle": {
		"id":   uint(558),
		"form": uint(0),
		"type": "bug",
	},
	"Scraggy": {
		"id":   uint(559),
		"form": uint(0),
		"type": "dark",
	},
	"Scrafty": {
		"id":   uint(560),
		"form": uint(0),
		"type": "dark",
	},
	"Sigilyph": {
		"id":   uint(561),
		"form": uint(0),
		"type": "psychic",
	},
	"Yamask": {
		"id":   uint(562),
		"form": uint(0),
		"type": "ghost",
	},
	"Yamask-Galar": {
		"id":   uint(562),
		"form": uint(1),
		"type": "ground",
	},
	"Cofagrigus": {
		"id":   uint(563),
		"form": uint(0),
		"type": "ghost",
	},
	"Tirtouga": {
		"id":   uint(564),
		"form": uint(0),
		"type": "water",
	},
	"Carracosta": {
		"id":   uint(565),
		"form": uint(0),
		"type": "water",
	},
	"Archen": {
		"id":   uint(566),
		"form": uint(0),
		"type": "rock",
	},
	"Archeops": {
		"id":   uint(567),
		"form": uint(0),
		"type": "rock",
	},
	"Trubbish": {
		"id":   uint(568),
		"form": uint(0),
		"type": "poison",
	},
	"Garbodor": {
		"id":   uint(569),
		"form": uint(0),
		"type": "poison",
	},
	"Garbodor-Gmax": {
		"id":   uint(569),
		"form": uint(1),
		"type": "poison",
	},
	"Zorua": {
		"id":   uint(570),
		"form": uint(0),
		"type": "dark",
	},
	"Zoroark": {
		"id":   uint(571),
		"form": uint(0),
		"type": "dark",
	},
	"Minccino": {
		"id":   uint(572),
		"form": uint(0),
		"type": "normal",
	},
	"Cinccino": {
		"id":   uint(573),
		"form": uint(0),
		"type": "normal",
	},
	"Gothita": {
		"id":   uint(574),
		"form": uint(0),
		"type": "psychic",
	},
	"Gothorita": {
		"id":   uint(575),
		"form": uint(0),
		"type": "psychic",
	},
	"Gothitelle": {
		"id":   uint(576),
		"form": uint(0),
		"type": "psychic",
	},
	"Solosis": {
		"id":   uint(577),
		"form": uint(0),
		"type": "psychic",
	},
	"Duosion": {
		"id":   uint(578),
		"form": uint(0),
		"type": "psychic",
	},
	"Reuniclus": {
		"id":   uint(579),
		"form": uint(0),
		"type": "psychic",
	},
	"Ducklett": {
		"id":   uint(580),
		"form": uint(0),
		"type": "water",
	},
	"Swanna": {
		"id":   uint(581),
		"form": uint(0),
		"type": "water",
	},
	"Vanillite": {
		"id":   uint(582),
		"form": uint(0),
		"type": "ice",
	},
	"Vanillish": {
		"id":   uint(583),
		"form": uint(0),
		"type": "ice",
	},
	"Vanilluxe": {
		"id":   uint(584),
		"form": uint(0),
		"type": "ice",
	},
	"Deerling": {
		"id":   uint(585),
		"form": uint(0),
		"type": "normal",
	},
	"Deerling-Summer": {
		"id":   uint(585),
		"form": uint(1),
		"type": "normal",
	},
	"Deerling-Autumn": {
		"id":   uint(585),
		"form": uint(2),
		"type": "normal",
	},
	"Deerling-Winter": {
		"id":   uint(585),
		"form": uint(3),
		"type": "normal",
	},
	"Sawsbuck": {
		"id":   uint(586),
		"form": uint(0),
		"type": "normal",
	},
	"Sawsbuck-Summer": {
		"id":   uint(586),
		"form": uint(1),
		"type": "normal",
	},
	"Sawsbuck-Autumn": {
		"id":   uint(586),
		"form": uint(2),
		"type": "normal",
	},
	"Sawsbuck-Winter": {
		"id":   uint(586),
		"form": uint(3),
		"type": "normal",
	},
	"Emolga": {
		"id":   uint(587),
		"form": uint(0),
		"type": "electric",
	},
	"Karrablast": {
		"id":   uint(588),
		"form": uint(0),
		"type": "bug",
	},
	"Escavalier": {
		"id":   uint(589),
		"form": uint(0),
		"type": "bug",
	},
	"Foongus": {
		"id":   uint(590),
		"form": uint(0),
		"type": "grass",
	},
	"Amoonguss": {
		"id":   uint(591),
		"form": uint(0),
		"type": "grass",
	},
	"Frillish": {
		"id":   uint(592),
		"form": uint(0),
		"type": "water",
	},
	"Jellicent": {
		"id":   uint(593),
		"form": uint(0),
		"type": "water",
	},
	"Alomomola": {
		"id":   uint(594),
		"form": uint(0),
		"type": "water",
	},
	"Joltik": {
		"id":   uint(595),
		"form": uint(0),
		"type": "bug",
	},
	"Galvantula": {
		"id":   uint(596),
		"form": uint(0),
		"type": "bug",
	},
	"Ferroseed": {
		"id":   uint(597),
		"form": uint(0),
		"type": "grass",
	},
	"Ferrothorn": {
		"id":   uint(598),
		"form": uint(0),
		"type": "grass",
	},
	"Klink": {
		"id":   uint(599),
		"form": uint(0),
		"type": "steel",
	},
	"Klang": {
		"id":   uint(600),
		"form": uint(0),
		"type": "steel",
	},
	"Klinklang": {
		"id":   uint(601),
		"form": uint(0),
		"type": "steel",
	},
	"Tynamo": {
		"id":   uint(602),
		"form": uint(0),
		"type": "electric",
	},
	"Eelektrik": {
		"id":   uint(603),
		"form": uint(0),
		"type": "electric",
	},
	"Eelektross": {
		"id":   uint(604),
		"form": uint(0),
		"type": "electric",
	},
	"Elgyem": {
		"id":   uint(605),
		"form": uint(0),
		"type": "psychic",
	},
	"Beheeyem": {
		"id":   uint(606),
		"form": uint(0),
		"type": "psychic",
	},
	"Litwick": {
		"id":   uint(607),
		"form": uint(0),
		"type": "ghost",
	},
	"Lampent": {
		"id":   uint(608),
		"form": uint(0),
		"type": "ghost",
	},
	"Chandelure": {
		"id":   uint(609),
		"form": uint(0),
		"type": "ghost",
	},
	"Axew": {
		"id":   uint(610),
		"form": uint(0),
		"type": "dragon",
	},
	"Fraxure": {
		"id":   uint(611),
		"form": uint(0),
		"type": "dragon",
	},
	"Haxorus": {
		"id":   uint(612),
		"form": uint(0),
		"type": "dragon",
	},
	"Cubchoo": {
		"id":   uint(613),
		"form": uint(0),
		"type": "ice",
	},
	"Beartic": {
		"id":   uint(614),
		"form": uint(0),
		"type": "ice",
	},
	"Cryogonal": {
		"id":   uint(615),
		"form": uint(0),
		"type": "ice",
	},
	"Shelmet": {
		"id":   uint(616),
		"form": uint(0),
		"type": "bug",
	},
	"Accelgor": {
		"id":   uint(617),
		"form": uint(0),
		"type": "bug",
	},
	"Stunfisk": {
		"id":   uint(618),
		"form": uint(0),
		"type": "ground",
	},
	"Stunfisk-Galar": {
		"id":   uint(618),
		"form": uint(1),
		"type": "ground",
	},
	"Mienfoo": {
		"id":   uint(619),
		"form": uint(0),
		"type": "fighting",
	},
	"Mienshao": {
		"id":   uint(620),
		"form": uint(0),
		"type": "fighting",
	},
	"Druddigon": {
		"id":   uint(621),
		"form": uint(0),
		"type": "dragon",
	},
	"Golett": {
		"id":   uint(622),
		"form": uint(0),
		"type": "ground",
	},
	"Golurk": {
		"id":   uint(623),
		"form": uint(0),
		"type": "ground",
	},
	"Pawniard": {
		"id":   uint(624),
		"form": uint(0),
		"type": "dark",
	},
	"Bisharp": {
		"id":   uint(625),
		"form": uint(0),
		"type": "dark",
	},
	"Bouffalant": {
		"id":   uint(626),
		"form": uint(0),
		"type": "normal",
	},
	"Rufflet": {
		"id":   uint(627),
		"form": uint(0),
		"type": "normal",
	},
	"Braviary": {
		"id":   uint(628),
		"form": uint(0),
		"type": "normal",
	},
	"Vullaby": {
		"id":   uint(629),
		"form": uint(0),
		"type": "dark",
	},
	"Mandibuzz": {
		"id":   uint(630),
		"form": uint(0),
		"type": "dark",
	},
	"Heatmor": {
		"id":   uint(631),
		"form": uint(0),
		"type": "fire",
	},
	"Durant": {
		"id":   uint(632),
		"form": uint(0),
		"type": "bug",
	},
	"Deino": {
		"id":   uint(633),
		"form": uint(0),
		"type": "dark",
	},
	"Zweilous": {
		"id":   uint(634),
		"form": uint(0),
		"type": "dark",
	},
	"Hydreigon": {
		"id":   uint(635),
		"form": uint(0),
		"type": "dark",
	},
	"Larvesta": {
		"id":   uint(636),
		"form": uint(0),
		"type": "bug",
	},
	"Volcarona": {
		"id":   uint(637),
		"form": uint(0),
		"type": "bug",
	},
	"Cobalion": {
		"id":   uint(638),
		"form": uint(0),
		"type": "steel",
	},
	"Terrakion": {
		"id":   uint(639),
		"form": uint(0),
		"type": "rock",
	},
	"Virizion": {
		"id":   uint(640),
		"form": uint(0),
		"type": "grass",
	},
	"Tornadus": {
		"id":   uint(641),
		"form": uint(0),
		"type": "flying",
	},
	"Tornadus-Therian": {
		"id":   uint(641),
		"form": uint(1),
		"type": "flying",
	},
	"Thundurus": {
		"id":   uint(642),
		"form": uint(0),
		"type": "electric",
	},
	"Thundurus-Therian": {
		"id":   uint(642),
		"form": uint(1),
		"type": "electric",
	},
	"Reshiram": {
		"id":   uint(643),
		"form": uint(0),
		"type": "dragon",
	},
	"Zekrom": {
		"id":   uint(644),
		"form": uint(0),
		"type": "dragon",
	},
	"Landorus": {
		"id":   uint(645),
		"form": uint(0),
		"type": "ground",
	},
	"Landorus-Therian": {
		"id":   uint(645),
		"form": uint(1),
		"type": "ground",
	},
	"Kyurem": {
		"id":   uint(646),
		"form": uint(0),
		"type": "dragon",
	},
	"Kyurem-White": {
		"id":   uint(646),
		"form": uint(1),
		"type": "dragon",
	},
	"Kyurem-Black": {
		"id":   uint(646),
		"form": uint(2),
		"type": "dragon",
	},
	"Keldeo": {
		"id":   uint(647),
		"form": uint(0),
		"type": "water",
	},
	"Keldeo-Resolute": {
		"id":   uint(647),
		"form": uint(1),
		"type": "water",
	},
	"Meloetta": {
		"id":   uint(648),
		"form": uint(0),
		"type": "normal",
	},
	"Meloetta-Pirouette": {
		"id":   uint(648),
		"form": uint(1),
		"type": "normal",
	},
	"Genesect": {
		"id":   uint(649),
		"form": uint(0),
		"type": "bug",
	},
	"Chespin": {
		"id":   uint(650),
		"form": uint(0),
		"type": "grass",
	},
	"Quilladin": {
		"id":   uint(651),
		"form": uint(0),
		"type": "grass",
	},
	"Chesnaught": {
		"id":   uint(652),
		"form": uint(0),
		"type": "grass",
	},
	"Fennekin": {
		"id":   uint(653),
		"form": uint(0),
		"type": "fire",
	},
	"Braixen": {
		"id":   uint(654),
		"form": uint(0),
		"type": "fire",
	},
	"Delphox": {
		"id":   uint(655),
		"form": uint(0),
		"type": "fire",
	},
	"Froakie": {
		"id":   uint(656),
		"form": uint(0),
		"type": "water",
	},
	"Frogadier": {
		"id":   uint(657),
		"form": uint(0),
		"type": "water",
	},
	"Greninja-Ash": {
		"id":   uint(658),
		"form": uint(2),
		"type": "water",
	},
	"Greninja": {
		"id":   uint(658),
		"form": uint(0),
		"type": "water",
	},
	"Bunnelby": {
		"id":   uint(659),
		"form": uint(0),
		"type": "normal",
	},
	"Diggersby": {
		"id":   uint(660),
		"form": uint(0),
		"type": "normal",
	},
	"Fletchling": {
		"id":   uint(661),
		"form": uint(0),
		"type": "normal",
	},
	"Fletchinder": {
		"id":   uint(662),
		"form": uint(0),
		"type": "fire",
	},
	"Talonflame": {
		"id":   uint(663),
		"form": uint(0),
		"type": "fire",
	},
	"Scatterbug": {
		"id":   uint(664),
		"form": uint(0),
		"type": "bug",
	},
	"Spewpa": {
		"id":   uint(665),
		"form": uint(0),
		"type": "bug",
	},
	"Vivillon-Icysnow": {
		"id":   uint(666),
		"form": uint(0),
		"type": "bug",
	},
	"Vivillon-Polar": {
		"id":   uint(666),
		"form": uint(1),
		"type": "bug",
	},
	"Vivillon-Tundra": {
		"id":   uint(666),
		"form": uint(2),
		"type": "bug",
	},
	"Vivillon-Continental": {
		"id":   uint(666),
		"form": uint(3),
		"type": "bug",
	},
	"Vivillon-Garden": {
		"id":   uint(666),
		"form": uint(4),
		"type": "bug",
	},
	"Vivillon-Elegant": {
		"id":   uint(666),
		"form": uint(5),
		"type": "bug",
	},
	"Vivillon": {
		"id":   uint(666),
		"form": uint(6),
		"type": "bug",
	},
	"Vivillon-Modern": {
		"id":   uint(666),
		"form": uint(7),
		"type": "bug",
	},
	"Vivillon-Marine": {
		"id":   uint(666),
		"form": uint(8),
		"type": "bug",
	},
	"Vivillon-Archipelago": {
		"id":   uint(666),
		"form": uint(9),
		"type": "bug",
	},
	"Vivillon-Highplains": {
		"id":   uint(666),
		"form": uint(10),
		"type": "bug",
	},
	"Vivillon-Sandstorm": {
		"id":   uint(666),
		"form": uint(11),
		"type": "bug",
	},
	"Vivillon-River": {
		"id":   uint(666),
		"form": uint(12),
		"type": "bug",
	},
	"Vivillon-Monsoon": {
		"id":   uint(666),
		"form": uint(13),
		"type": "bug",
	},
	"Vivillon-Savanna": {
		"id":   uint(666),
		"form": uint(14),
		"type": "bug",
	},
	"Vivillon-Sun": {
		"id":   uint(666),
		"form": uint(15),
		"type": "bug",
	},
	"Vivillon-Ocean": {
		"id":   uint(666),
		"form": uint(16),
		"type": "bug",
	},
	"Vivillon-Jungle": {
		"id":   uint(666),
		"form": uint(17),
		"type": "bug",
	},
	"Vivillon-Fancy": {
		"id":   uint(666),
		"form": uint(18),
		"type": "bug",
	},
	"Vivillon-Pokeball": {
		"id":   uint(666),
		"form": uint(19),
		"type": "bug",
	},
	"Litleo": {
		"id":   uint(667),
		"form": uint(0),
		"type": "fire",
	},
	"Pyroar": {
		"id":   uint(668),
		"form": uint(0),
		"type": "fire",
	},
	"Flabebe": {
		"id":   uint(669),
		"form": uint(0),
		"type": "fairy",
	},
	"Flabebe-Yellow": {
		"id":   uint(669),
		"form": uint(1),
		"type": "fairy",
	},
	"Flabebe-Orange": {
		"id":   uint(669),
		"form": uint(2),
		"type": "fairy",
	},
	"Flabebe-Blue": {
		"id":   uint(669),
		"form": uint(3),
		"type": "fairy",
	},
	"Flabebe-White": {
		"id":   uint(669),
		"form": uint(4),
		"type": "fairy",
	},
	"Floette": {
		"id":   uint(670),
		"form": uint(0),
		"type": "fairy",
	},
	"Floette-Yellow": {
		"id":   uint(670),
		"form": uint(1),
		"type": "fairy",
	},
	"Floette-Orange": {
		"id":   uint(670),
		"form": uint(2),
		"type": "fairy",
	},
	"Floette-Blue": {
		"id":   uint(670),
		"form": uint(3),
		"type": "fairy",
	},
	"Floette-White": {
		"id":   uint(670),
		"form": uint(4),
		"type": "fairy",
	},
	"Floette-Eternal": {
		"id":   uint(670),
		"form": uint(5),
		"type": "fairy",
	},
	"Florges": {
		"id":   uint(671),
		"form": uint(0),
		"type": "fairy",
	},
	"Florges-Yellow": {
		"id":   uint(671),
		"form": uint(1),
		"type": "fairy",
	},
	"Florges-Orange": {
		"id":   uint(671),
		"form": uint(2),
		"type": "fairy",
	},
	"Florges-Blue": {
		"id":   uint(671),
		"form": uint(3),
		"type": "fairy",
	},
	"Florges-White": {
		"id":   uint(671),
		"form": uint(4),
		"type": "fairy",
	},
	"Skiddo": {
		"id":   uint(672),
		"form": uint(0),
		"type": "grass",
	},
	"Gogoat": {
		"id":   uint(673),
		"form": uint(0),
		"type": "grass",
	},
	"Pancham": {
		"id":   uint(674),
		"form": uint(0),
		"type": "fighting",
	},
	"Pangoro": {
		"id":   uint(675),
		"form": uint(0),
		"type": "fighting",
	},
	"Furfrou": {
		"id":   uint(676),
		"form": uint(0),
		"type": "normal",
	},
	"Espurr": {
		"id":   uint(677),
		"form": uint(0),
		"type": "psychic",
	},
	"Meowstic": {
		"id":   uint(678),
		"form": uint(0),
		"type": "psychic",
	},
	"Meowstic-F": {
		"id":   uint(678),
		"form": uint(1),
		"type": "psychic",
	},
	"Honedge": {
		"id":   uint(679),
		"form": uint(0),
		"type": "steel",
	},
	"Doublade": {
		"id":   uint(680),
		"form": uint(0),
		"type": "steel",
	},
	"Aegislash": {
		"id":   uint(681),
		"form": uint(0),
		"type": "steel",
	},
	"Aegislash-Blade": {
		"id":   uint(681),
		"form": uint(1),
		"type": "steel",
	},
	"Spritzee": {
		"id":   uint(682),
		"form": uint(0),
		"type": "fairy",
	},
	"Aromatisse": {
		"id":   uint(683),
		"form": uint(0),
		"type": "fairy",
	},
	"Swirlix": {
		"id":   uint(684),
		"form": uint(0),
		"type": "fairy",
	},
	"Slurpuff": {
		"id":   uint(685),
		"form": uint(0),
		"type": "fairy",
	},
	"Inkay": {
		"id":   uint(686),
		"form": uint(0),
		"type": "dark",
	},
	"Malamar": {
		"id":   uint(687),
		"form": uint(0),
		"type": "dark",
	},
	"Binacle": {
		"id":   uint(688),
		"form": uint(0),
		"type": "rock",
	},
	"Barbaracle": {
		"id":   uint(689),
		"form": uint(0),
		"type": "rock",
	},
	"Skrelp": {
		"id":   uint(690),
		"form": uint(0),
		"type": "poison",
	},
	"Dragalge": {
		"id":   uint(691),
		"form": uint(0),
		"type": "poison",
	},
	"Clauncher": {
		"id":   uint(692),
		"form": uint(0),
		"type": "water",
	},
	"Clawitzer": {
		"id":   uint(693),
		"form": uint(0),
		"type": "water",
	},
	"Helioptile": {
		"id":   uint(694),
		"form": uint(0),
		"type": "electric",
	},
	"Heliolisk": {
		"id":   uint(695),
		"form": uint(0),
		"type": "electric",
	},
	"Tyrunt": {
		"id":   uint(696),
		"form": uint(0),
		"type": "rock",
	},
	"Tyrantrum": {
		"id":   uint(697),
		"form": uint(0),
		"type": "rock",
	},
	"Amaura": {
		"id":   uint(698),
		"form": uint(0),
		"type": "rock",
	},
	"Aurorus": {
		"id":   uint(699),
		"form": uint(0),
		"type": "rock",
	},
	"Sylveon": {
		"id":   uint(700),
		"form": uint(0),
		"type": "fairy",
	},
	"Hawlucha": {
		"id":   uint(701),
		"form": uint(0),
		"type": "fighting",
	},
	"Dedenne": {
		"id":   uint(702),
		"form": uint(0),
		"type": "electric",
	},
	"Carbink": {
		"id":   uint(703),
		"form": uint(0),
		"type": "rock",
	},
	"Goomy": {
		"id":   uint(704),
		"form": uint(0),
		"type": "dragon",
	},
	"Sliggoo": {
		"id":   uint(705),
		"form": uint(0),
		"type": "dragon",
	},
	"Goodra": {
		"id":   uint(706),
		"form": uint(0),
		"type": "dragon",
	},
	"Klefki": {
		"id":   uint(707),
		"form": uint(0),
		"type": "steel",
	},
	"Phantump": {
		"id":   uint(708),
		"form": uint(0),
		"type": "ghost",
	},
	"Trevenant": {
		"id":   uint(709),
		"form": uint(0),
		"type": "ghost",
	},
	"Pumpkaboo": {
		"id":   uint(710),
		"form": uint(0),
		"type": "ghost",
	},
	"Pumpkaboo-Small": {
		"id":   uint(710),
		"form": uint(1),
		"type": "ghost",
	},
	"Pumpkaboo-Large": {
		"id":   uint(710),
		"form": uint(2),
		"type": "ghost",
	},
	"Pumpkaboo-Super": {
		"id":   uint(710),
		"form": uint(3),
		"type": "ghost",
	},
	"Gourgeist": {
		"id":   uint(711),
		"form": uint(0),
		"type": "ghost",
	},
	"Gourgeist-Small": {
		"id":   uint(711),
		"form": uint(1),
		"type": "ghost",
	},
	"Gourgeist-Large": {
		"id":   uint(711),
		"form": uint(2),
		"type": "ghost",
	},
	"Gourgeist-Super": {
		"id":   uint(711),
		"form": uint(3),
		"type": "ghost",
	},
	"Bergmite": {
		"id":   uint(712),
		"form": uint(0),
		"type": "ice",
	},
	"Avalugg": {
		"id":   uint(713),
		"form": uint(0),
		"type": "ice",
	},
	"Noibat": {
		"id":   uint(714),
		"form": uint(0),
		"type": "flying",
	},
	"Noivern": {
		"id":   uint(715),
		"form": uint(0),
		"type": "flying",
	},
	"Xerneas": {
		"id":   uint(716),
		"form": uint(0),
		"type": "fairy",
	},
	"Yveltal": {
		"id":   uint(717),
		"form": uint(0),
		"type": "dark",
	},
	"Zygarde": {
		"id":   uint(718),
		"form": uint(0),
		"type": "dragon",
	},
	"Zygarde-10%": {
		"id":   uint(718),
		"form": uint(1),
		"type": "dragon",
	},
	"Zygarde-Complete": {
		"id":   uint(718),
		"form": uint(4),
		"type": "dragon",
	},
	"Diancie": {
		"id":   uint(719),
		"form": uint(0),
		"type": "rock",
	},
	"Diancie-Mega": {
		"id":   uint(719),
		"form": uint(1),
		"type": "rock",
	},
	"Hoopa": {
		"id":   uint(720),
		"form": uint(0),
		"type": "psychic",
	},
	"Hoopa-Unbound": {
		"id":   uint(720),
		"form": uint(1),
		"type": "psychic",
	},
	"Volcanion": {
		"id":   uint(721),
		"form": uint(0),
		"type": "fire",
	},
	"Rowlet": {
		"id":   uint(722),
		"form": uint(0),
		"type": "grass",
	},
	"Dartrix": {
		"id":   uint(723),
		"form": uint(0),
		"type": "grass",
	},
	"Decidueye": {
		"id":   uint(724),
		"form": uint(0),
		"type": "grass",
	},
	"Litten": {
		"id":   uint(725),
		"form": uint(0),
		"type": "fire",
	},
	"Torracat": {
		"id":   uint(726),
		"form": uint(0),
		"type": "fire",
	},
	"Incineroar": {
		"id":   uint(727),
		"form": uint(0),
		"type": "fire",
	},
	"Popplio": {
		"id":   uint(728),
		"form": uint(0),
		"type": "water",
	},
	"Brionne": {
		"id":   uint(729),
		"form": uint(0),
		"type": "water",
	},
	"Primarina": {
		"id":   uint(730),
		"form": uint(0),
		"type": "water",
	},
	"Pikipek": {
		"id":   uint(731),
		"form": uint(0),
		"type": "normal",
	},
	"Trumbeak": {
		"id":   uint(732),
		"form": uint(0),
		"type": "normal",
	},
	"Toucannon": {
		"id":   uint(733),
		"form": uint(0),
		"type": "normal",
	},
	"Yungoos": {
		"id":   uint(734),
		"form": uint(0),
		"type": "normal",
	},
	"Gumshoos": {
		"id":   uint(735),
		"form": uint(0),
		"type": "normal",
	},
	"Gumshoos-Totem": {
		"id":   uint(735),
		"form": uint(0),
		"type": "normal",
	},
	"Grubbin": {
		"id":   uint(736),
		"form": uint(0),
		"type": "bug",
	},
	"Charjabug": {
		"id":   uint(737),
		"form": uint(0),
		"type": "bug",
	},
	"Vikavolt": {
		"id":   uint(738),
		"form": uint(0),
		"type": "bug",
	},
	"Vikavolt-Totem": {
		"id":   uint(738),
		"form": uint(0),
		"type": "bug",
	},
	"Crabrawler": {
		"id":   uint(739),
		"form": uint(0),
		"type": "fighting",
	},
	"Crabominable": {
		"id":   uint(740),
		"form": uint(0),
		"type": "fighting",
	},
	"Oricorio": {
		"id":   uint(741),
		"form": uint(0),
		"type": "fire",
	},
	"Oricorio-Pom-Pom": {
		"id":   uint(741),
		"form": uint(1),
		"type": "electric",
	},
	"Oricorio-Pa'u": {
		"id":   uint(741),
		"form": uint(2),
		"type": "psychic",
	},
	"Oricorio-Sensu": {
		"id":   uint(741),
		"form": uint(3),
		"type": "ghost",
	},
	"Cutiefly": {
		"id":   uint(742),
		"form": uint(0),
		"type": "bug",
	},
	"Ribombee": {
		"id":   uint(743),
		"form": uint(0),
		"type": "bug",
	},
	"Ribombee-Totem": {
		"id":   uint(743),
		"form": uint(0),
		"type": "bug",
	},
	"Rockruff": {
		"id":   uint(744),
		"form": uint(0),
		"type": "rock",
	},
	"Lycanroc": {
		"id":   uint(745),
		"form": uint(0),
		"type": "rock",
	},
	"Lycanroc-Midnight": {
		"id":   uint(745),
		"form": uint(1),
		"type": "rock",
	},
	"Lycanroc-Dusk": {
		"id":   uint(745),
		"form": uint(2),
		"type": "rock",
	},
	"Wishiwashi": {
		"id":   uint(746),
		"form": uint(0),
		"type": "water",
	},
	"Wishiwashi-School": {
		"id":   uint(746),
		"form": uint(1),
		"type": "water",
	},
	"Mareanie": {
		"id":   uint(747),
		"form": uint(0),
		"type": "poison",
	},
	"Toxapex": {
		"id":   uint(748),
		"form": uint(0),
		"type": "poison",
	},
	"Mudbray": {
		"id":   uint(749),
		"form": uint(0),
		"type": "ground",
	},
	"Mudsdale": {
		"id":   uint(750),
		"form": uint(0),
		"type": "ground",
	},
	"Dewpider": {
		"id":   uint(751),
		"form": uint(0),
		"type": "water",
	},
	"Araquanid": {
		"id":   uint(752),
		"form": uint(0),
		"type": "water",
	},
	"Araquanid-Totem": {
		"id":   uint(752),
		"form": uint(0),
		"type": "water",
	},
	"Fomantis": {
		"id":   uint(753),
		"form": uint(0),
		"type": "grass",
	},
	"Lurantis": {
		"id":   uint(754),
		"form": uint(0),
		"type": "grass",
	},
	"Lurantis-Totem": {
		"id":   uint(754),
		"form": uint(0),
		"type": "grass",
	},
	"Morelull": {
		"id":   uint(755),
		"form": uint(0),
		"type": "grass",
	},
	"Shiinotic": {
		"id":   uint(756),
		"form": uint(0),
		"type": "grass",
	},
	"Salandit": {
		"id":   uint(757),
		"form": uint(0),
		"type": "poison",
	},
	"Salazzle": {
		"id":   uint(758),
		"form": uint(0),
		"type": "poison",
	},
	"Salazzle-Totem": {
		"id":   uint(758),
		"form": uint(0),
		"type": "poison",
	},
	"Stufful": {
		"id":   uint(759),
		"form": uint(0),
		"type": "normal",
	},
	"Bewear": {
		"id":   uint(760),
		"form": uint(0),
		"type": "normal",
	},
	"Bounsweet": {
		"id":   uint(761),
		"form": uint(0),
		"type": "grass",
	},
	"Steenee": {
		"id":   uint(762),
		"form": uint(0),
		"type": "grass",
	},
	"Tsareena": {
		"id":   uint(763),
		"form": uint(0),
		"type": "grass",
	},
	"Comfey": {
		"id":   uint(764),
		"form": uint(0),
		"type": "fairy",
	},
	"Oranguru": {
		"id":   uint(765),
		"form": uint(0),
		"type": "normal",
	},
	"Passimian": {
		"id":   uint(766),
		"form": uint(0),
		"type": "fighting",
	},
	"Wimpod": {
		"id":   uint(767),
		"form": uint(0),
		"type": "bug",
	},
	"Golisopod": {
		"id":   uint(768),
		"form": uint(0),
		"type": "bug",
	},
	"Sandygast": {
		"id":   uint(769),
		"form": uint(0),
		"type": "ghost",
	},
	"Palossand": {
		"id":   uint(770),
		"form": uint(0),
		"type": "ghost",
	},
	"Pyukumuku": {
		"id":   uint(771),
		"form": uint(0),
		"type": "water",
	},
	"Type: Null": {
		"id":   uint(772),
		"form": uint(0),
		"type": "normal",
	},
	"Silvally-Fairy": {
		"id":   uint(773),
		"form": uint(17),
		"type": "fairy",
	},
	"Silvally-Dark": {
		"id":   uint(773),
		"form": uint(16),
		"type": "dark",
	},
	"Silvally-Dragon": {
		"id":   uint(773),
		"form": uint(15),
		"type": "dragon",
	},
	"Silvally-Ice": {
		"id":   uint(773),
		"form": uint(14),
		"type": "ice",
	},
	"Silvally-Psychic": {
		"id":   uint(773),
		"form": uint(13),
		"type": "psychic",
	},
	"Silvally-Electric": {
		"id":   uint(773),
		"form": uint(12),
		"type": "electric",
	},
	"Silvally-Grass": {
		"id":   uint(773),
		"form": uint(11),
		"type": "grass",
	},
	"Silvally-Water": {
		"id":   uint(773),
		"form": uint(10),
		"type": "water",
	},
	"Silvally-Fire": {
		"id":   uint(773),
		"form": uint(9),
		"type": "fire",
	},
	"Silvally-Steel": {
		"id":   uint(773),
		"form": uint(8),
		"type": "steel",
	},
	"Silvally-Ghost": {
		"id":   uint(773),
		"form": uint(7),
		"type": "ghost",
	},
	"Silvally-Bug": {
		"id":   uint(773),
		"form": uint(6),
		"type": "bug",
	},
	"Silvally-Rock": {
		"id":   uint(773),
		"form": uint(5),
		"type": "rock",
	},
	"Silvally-Ground": {
		"id":   uint(773),
		"form": uint(4),
		"type": "ground",
	},
	"Silvally-Poison": {
		"id":   uint(773),
		"form": uint(3),
		"type": "poison",
	},
	"Silvally-Flying": {
		"id":   uint(773),
		"form": uint(2),
		"type": "flying",
	},
	"Silvally-Fighting": {
		"id":   uint(773),
		"form": uint(1),
		"type": "fighting",
	},
	"Silvally": {
		"id":   uint(773),
		"form": uint(0),
		"type": "normal",
	},
	"Minior-Meteor": {
		"id":   uint(774),
		"form": uint(0),
		"type": "rock",
	},
	"Minior": {
		"id":   uint(774),
		"form": uint(7),
		"type": "rock",
	},
	"Minior-Orange": {
		"id":   uint(774),
		"form": uint(8),
		"type": "rock",
	},
	"Minior-Green": {
		"id":   uint(774),
		"form": uint(10),
		"type": "rock",
	},
	"Minior-Blue": {
		"id":   uint(774),
		"form": uint(11),
		"type": "rock",
	},
	"Minior-Indigo": {
		"id":   uint(774),
		"form": uint(12),
		"type": "rock",
	},
	"Minior-Violet": {
		"id":   uint(774),
		"form": uint(13),
		"type": "rock",
	},
	"Komala": {
		"id":   uint(775),
		"form": uint(0),
		"type": "normal",
	},
	"Turtonator": {
		"id":   uint(776),
		"form": uint(0),
		"type": "fire",
	},
	"Togedemaru": {
		"id":   uint(777),
		"form": uint(0),
		"type": "electric",
	},
	"Togedemaru-Totem": {
		"id":   uint(777),
		"form": uint(0),
		"type": "electric",
	},
	"Mimikyu": {
		"id":   uint(778),
		"form": uint(0),
		"type": "ghost",
	},
	"Mimikyu-Totem": {
		"id":   uint(778),
		"form": uint(0),
		"type": "ghost",
	},
	"Bruxish": {
		"id":   uint(779),
		"form": uint(0),
		"type": "water",
	},
	"Drampa": {
		"id":   uint(780),
		"form": uint(0),
		"type": "normal",
	},
	"Dhelmise": {
		"id":   uint(781),
		"form": uint(0),
		"type": "ghost",
	},
	"Jangmo-o": {
		"id":   uint(782),
		"form": uint(0),
		"type": "dragon",
	},
	"Hakamo-o": {
		"id":   uint(783),
		"form": uint(0),
		"type": "dragon",
	},
	"Kommo-o": {
		"id":   uint(784),
		"form": uint(0),
		"type": "dragon",
	},
	"Kommo-o-Totem": {
		"id":   uint(784),
		"form": uint(0),
		"type": "dragon",
	},
	"Tapu Koko": {
		"id":   uint(785),
		"form": uint(0),
		"type": "electric",
	},
	"Tapu Lele": {
		"id":   uint(786),
		"form": uint(0),
		"type": "psychic",
	},
	"Tapu Bulu": {
		"id":   uint(787),
		"form": uint(0),
		"type": "grass",
	},
	"Tapu Fini": {
		"id":   uint(788),
		"form": uint(0),
		"type": "water",
	},
	"Cosmog": {
		"id":   uint(789),
		"form": uint(0),
		"type": "psychic",
	},
	"Cosmoem": {
		"id":   uint(790),
		"form": uint(0),
		"type": "psychic",
	},
	"Solgaleo": {
		"id":   uint(791),
		"form": uint(0),
		"type": "psychic",
	},
	"Lunala": {
		"id":   uint(792),
		"form": uint(0),
		"type": "psychic",
	},
	"Nihilego": {
		"id":   uint(793),
		"form": uint(0),
		"type": "rock",
	},
	"Buzzwole": {
		"id":   uint(794),
		"form": uint(0),
		"type": "bug",
	},
	"Pheromosa": {
		"id":   uint(795),
		"form": uint(0),
		"type": "bug",
	},
	"Xurkitree": {
		"id":   uint(796),
		"form": uint(0),
		"type": "electric",
	},
	"Celesteela": {
		"id":   uint(797),
		"form": uint(0),
		"type": "steel",
	},
	"Kartana": {
		"id":   uint(798),
		"form": uint(0),
		"type": "grass",
	},
	"Guzzlord": {
		"id":   uint(799),
		"form": uint(0),
		"type": "dark",
	},
	"Necrozma": {
		"id":   uint(800),
		"form": uint(0),
		"type": "psychic",
	},
	"Necrozma-Dusk-Mane": {
		"id":   uint(800),
		"form": uint(1),
		"type": "psychic",
	},
	"Necrozma-Dawn-Wings": {
		"id":   uint(800),
		"form": uint(2),
		"type": "psychic",
	},
	"Necrozma-Ultra": {
		"id":   uint(800),
		"form": uint(3),
		"type": "psychic",
	},
	"Magearna": {
		"id":   uint(801),
		"form": uint(0),
		"type": "steel",
	},
	"Marshadow": {
		"id":   uint(802),
		"form": uint(0),
		"type": "ghost",
	},
	"Poipole": {
		"id":   uint(803),
		"form": uint(0),
		"type": "poison",
	},
	"Naganadel": {
		"id":   uint(804),
		"form": uint(0),
		"type": "poison",
	},
	"Stakataka": {
		"id":   uint(805),
		"form": uint(0),
		"type": "rock",
	},
	"Blacephalon": {
		"id":   uint(806),
		"form": uint(0),
		"type": "fire",
	},
	"Zeraora": {
		"id":   uint(807),
		"form": uint(0),
		"type": "electric",
	},
	"Meltan": {
		"id":   uint(808),
		"form": uint(0),
		"type": "steel",
	},
	"Melmetal": {
		"id":   uint(809),
		"form": uint(0),
		"type": "steel",
	},
	"Melmetal-Gmax": {
		"id":   uint(809),
		"form": uint(1),
		"type": "steel",
	},
	"Grookey": {
		"id":   uint(810),
		"form": uint(0),
		"type": "grass",
	},
	"Thwackey": {
		"id":   uint(811),
		"form": uint(0),
		"type": "grass",
	},
	"Rillaboom": {
		"id":   uint(812),
		"form": uint(0),
		"type": "grass",
	},
	"Rillaboom-Gmax": {
		"id":   uint(812),
		"form": uint(1),
		"type": "grass",
	},
	"Scorbunny": {
		"id":   uint(813),
		"form": uint(0),
		"type": "fire",
	},
	"Raboot": {
		"id":   uint(814),
		"form": uint(0),
		"type": "fire",
	},
	"Cinderace": {
		"id":   uint(815),
		"form": uint(0),
		"type": "fire",
	},
	"Cinderace-Gmax": {
		"id":   uint(815),
		"form": uint(1),
		"type": "fire",
	},
	"Sobble": {
		"id":   uint(816),
		"form": uint(0),
		"type": "water",
	},
	"Drizzile": {
		"id":   uint(817),
		"form": uint(0),
		"type": "water",
	},
	"Inteleon": {
		"id":   uint(818),
		"form": uint(0),
		"type": "water",
	},
	"Inteleon-Gmax": {
		"id":   uint(818),
		"form": uint(1),
		"type": "water",
	},
	"Skwovet": {
		"id":   uint(819),
		"form": uint(0),
		"type": "normal",
	},
	"Greedent": {
		"id":   uint(820),
		"form": uint(0),
		"type": "normal",
	},
	"Rookidee": {
		"id":   uint(821),
		"form": uint(0),
		"type": "flying",
	},
	"Corvisquire": {
		"id":   uint(822),
		"form": uint(0),
		"type": "flying",
	},
	"Corviknight": {
		"id":   uint(823),
		"form": uint(0),
		"type": "flying",
	},
	"Corviknight-Gmax": {
		"id":   uint(823),
		"form": uint(1),
		"type": "flying",
	},
	"Blipbug": {
		"id":   uint(824),
		"form": uint(0),
		"type": "bug",
	},
	"Dottler": {
		"id":   uint(825),
		"form": uint(0),
		"type": "bug",
	},
	"Orbeetle": {
		"id":   uint(826),
		"form": uint(0),
		"type": "bug",
	},
	"Orbeetle-Gmax": {
		"id":   uint(826),
		"form": uint(1),
		"type": "bug",
	},
	"Nickit": {
		"id":   uint(827),
		"form": uint(0),
		"type": "dark",
	},
	"Thievul": {
		"id":   uint(828),
		"form": uint(0),
		"type": "dark",
	},
	"Gossifleur": {
		"id":   uint(829),
		"form": uint(0),
		"type": "grass",
	},
	"Eldegoss": {
		"id":   uint(830),
		"form": uint(0),
		"type": "grass",
	},
	"Wooloo": {
		"id":   uint(831),
		"form": uint(0),
		"type": "normal",
	},
	"Dubwool": {
		"id":   uint(832),
		"form": uint(0),
		"type": "normal",
	},
	"Chewtle": {
		"id":   uint(833),
		"form": uint(0),
		"type": "water",
	},
	"Drednaw": {
		"id":   uint(834),
		"form": uint(0),
		"type": "water",
	},
	"Drednaw-Gmax": {
		"id":   uint(834),
		"form": uint(1),
		"type": "water",
	},
	"Yamper": {
		"id":   uint(835),
		"form": uint(0),
		"type": "electric",
	},
	"Boltund": {
		"id":   uint(836),
		"form": uint(0),
		"type": "electric",
	},
	"Rolycoly": {
		"id":   uint(837),
		"form": uint(0),
		"type": "rock",
	},
	"Carkol": {
		"id":   uint(838),
		"form": uint(0),
		"type": "rock",
	},
	"Coalossal": {
		"id":   uint(839),
		"form": uint(0),
		"type": "rock",
	},
	"Coalossal-Gmax": {
		"id":   uint(839),
		"form": uint(1),
		"type": "rock",
	},
	"Applin": {
		"id":   uint(840),
		"form": uint(0),
		"type": "grass",
	},
	"Flapple": {
		"id":   uint(841),
		"form": uint(0),
		"type": "grass",
	},
	"Flapple-Gmax": {
		"id":   uint(841),
		"form": uint(1),
		"type": "grass",
	},
	"Appletun": {
		"id":   uint(842),
		"form": uint(0),
		"type": "grass",
	},
	"Appletun-Gmax": {
		"id":   uint(842),
		"form": uint(1),
		"type": "grass",
	},
	"Silicobra": {
		"id":   uint(843),
		"form": uint(0),
		"type": "ground",
	},
	"Sandaconda": {
		"id":   uint(844),
		"form": uint(0),
		"type": "ground",
	},
	"Sandaconda-Gmax": {
		"id":   uint(844),
		"form": uint(1),
		"type": "ground",
	},
	"Cramorant": {
		"id":   uint(845),
		"form": uint(0),
		"type": "flying",
	},
	"Cramorant-Gulping": {
		"id":   uint(845),
		"form": uint(1),
		"type": "flying",
	},
	"Cramorant-Gorging": {
		"id":   uint(845),
		"form": uint(2),
		"type": "flying",
	},
	"Arrokuda": {
		"id":   uint(846),
		"form": uint(0),
		"type": "water",
	},
	"Barraskewda": {
		"id":   uint(847),
		"form": uint(0),
		"type": "water",
	},
	"Toxel": {
		"id":   uint(848),
		"form": uint(0),
		"type": "electric",
	},
	"Toxtricity": {
		"id":   uint(849),
		"form": uint(0),
		"type": "electric",
	},
	"Toxtricity-Low-Key": {
		"id":   uint(849),
		"form": uint(1),
		"type": "electric",
	},
	"Toxtricity-Gmax": {
		"id":   uint(844),
		"form": uint(2),
		"type": "electric",
	},
	"Sizzlipede": {
		"id":   uint(850),
		"form": uint(0),
		"type": "fire",
	},
	"Centiskorch": {
		"id":   uint(851),
		"form": uint(0),
		"type": "fire",
	},
	"Centiskorch-Gmax": {
		"id":   uint(851),
		"form": uint(1),
		"type": "fire",
	},
	"Clobbopus": {
		"id":   uint(852),
		"form": uint(0),
		"type": "fighting",
	},
	"Grapploct": {
		"id":   uint(853),
		"form": uint(0),
		"type": "fighting",
	},
	"Sinistea": {
		"id":   uint(854),
		"form": uint(0),
		"type": "ghost",
	},
	"Polteageist": {
		"id":   uint(855),
		"form": uint(0),
		"type": "ghost",
	},
	"Hatenna": {
		"id":   uint(856),
		"form": uint(0),
		"type": "psychic",
	},
	"Hattrem": {
		"id":   uint(857),
		"form": uint(0),
		"type": "psychic",
	},
	"Hatterene": {
		"id":   uint(858),
		"form": uint(0),
		"type": "psychic",
	},
	"Hatterene-Gmax": {
		"id":   uint(858),
		"form": uint(1),
		"type": "psychic",
	},
	"Impidimp": {
		"id":   uint(859),
		"form": uint(0),
		"type": "dark",
	},
	"Morgrem": {
		"id":   uint(860),
		"form": uint(0),
		"type": "dark",
	},
	"Grimmsnarl": {
		"id":   uint(861),
		"form": uint(0),
		"type": "dark",
	},
	"Grimmsnarl-Gmax": {
		"id":   uint(861),
		"form": uint(1),
		"type": "dark",
	},
	"Obstagoon": {
		"id":   uint(862),
		"form": uint(0),
		"type": "dark",
	},
	"Perrserker": {
		"id":   uint(863),
		"form": uint(0),
		"type": "steel",
	},
	"Cursola": {
		"id":   uint(864),
		"form": uint(0),
		"type": "ghost",
	},
	"Sirfetch'd": {
		"id":   uint(865),
		"form": uint(0),
		"type": "fighting",
	},
	"Sirfetch&#x27;d": {
		"id":   uint(83),
		"form": uint(0),
		"type": "normal",
	},
	"Sirfetch’d": {
		"id":   uint(83),
		"form": uint(0),
		"type": "normal",
	},
	"Sirfetch&#39;d": {
		"id":   uint(865),
		"form": uint(0),
		"type": "fighting",
	},
	"Mr. Rime": {
		"id":   uint(866),
		"form": uint(0),
		"type": "ice",
	},
	"Runerigus": {
		"id":   uint(867),
		"form": uint(0),
		"type": "ground",
	},
	"Milcery": {
		"id":   uint(868),
		"form": uint(0),
		"type": "fairy",
	},
	"Alcremie": {
		"id":   uint(869),
		"form": uint(0),
		"type": "fairy",
	},
	"Alcremie-Gmax": {
		"id":   uint(869),
		"form": uint(1),
		"type": "fairy",
	},
	"Falinks": {
		"id":   uint(870),
		"form": uint(0),
		"type": "fighting",
	},
	"Pincurchin": {
		"id":   uint(871),
		"form": uint(0),
		"type": "electric",
	},
	"Snom": {
		"id":   uint(872),
		"form": uint(0),
		"type": "ice",
	},
	"Frosmoth": {
		"id":   uint(873),
		"form": uint(0),
		"type": "ice",
	},
	"Stonjourner": {
		"id":   uint(874),
		"form": uint(0),
		"type": "rock",
	},
	"Eiscue": {
		"id":   uint(875),
		"form": uint(0),
		"type": "ice",
	},
	"Eiscue-Noice": {
		"id":   uint(875),
		"form": uint(1),
		"type": "ice",
	},
	"Indeedee": {
		"id":   uint(876),
		"form": uint(0),
		"type": "psychic",
	},
	"Indeedee-F": {
		"id":   uint(876),
		"form": uint(1),
		"type": "psychic",
	},
	"Morpeko": {
		"id":   uint(877),
		"form": uint(0),
		"type": "electric",
	},
	"Morpeko-Hangry": {
		"id":   uint(877),
		"form": uint(1),
		"type": "electric",
	},
	"Cufant": {
		"id":   uint(878),
		"form": uint(0),
		"type": "steel",
	},
	"Copperajah": {
		"id":   uint(879),
		"form": uint(0),
		"type": "steel",
	},
	"Copperajah-Gmax": {
		"id":   uint(879),
		"form": uint(1),
		"type": "steel",
	},
	"Dracozolt": {
		"id":   uint(880),
		"form": uint(0),
		"type": "electric",
	},
	"Arctozolt": {
		"id":   uint(881),
		"form": uint(0),
		"type": "electric",
	},
	"Dracovish": {
		"id":   uint(882),
		"form": uint(0),
		"type": "water",
	},
	"Arctovish": {
		"id":   uint(883),
		"form": uint(0),
		"type": "water",
	},
	"Duraludon": {
		"id":   uint(884),
		"form": uint(0),
		"type": "steel",
	},
	"Duraludon-Gmax": {
		"id":   uint(884),
		"form": uint(1),
		"type": "steel",
	},
	"Dreepy": {
		"id":   uint(885),
		"form": uint(0),
		"type": "dragon",
	},
	"Drakloak": {
		"id":   uint(886),
		"form": uint(0),
		"type": "dragon",
	},
	"Dragapult": {
		"id":   uint(887),
		"form": uint(0),
		"type": "dragon",
	},
	"Zacian": {
		"id":   uint(888),
		"form": uint(0),
		"type": "fairy",
	},
	"Zacian-Crowned": {
		"id":   uint(888),
		"form": uint(1),
		"type": "fairy",
	},
	"Zamazenta": {
		"id":   uint(889),
		"form": uint(0),
		"type": "fighting",
	},
	"Zamazenta-Crowned": {
		"id":   uint(889),
		"form": uint(1),
		"type": "fighting",
	},
	"Eternatus": {
		"id":   uint(890),
		"form": uint(0),
		"type": "poison",
	},
	"Eternatus-Eternamax": {
		"id":   uint(890),
		"form": uint(1),
		"type": "poison",
	},
	"Kubfu": {
		"id":   uint(891),
		"form": uint(0),
		"type": "fighting",
	},
	"Urshifu": {
		"id":   uint(892),
		"form": uint(0),
		"type": "fighting",
	},
	"Urshifu-Rapid-Strike": {
		"id":   uint(892),
		"form": uint(1),
		"type": "fighting",
	},
	"Urshifu-Gmax": {
		"id":   uint(892),
		"form": uint(2),
		"type": "fighting",
	},
	"Urshifu-Rapid-Strike-Gmax": {
		"id":   uint(892),
		"form": uint(3),
		"type": "fighting",
	},
	"Zarude": {
		"id":   uint(893),
		"form": uint(0),
		"type": "dark",
	},
	"Regieleki": {
		"id":   uint(894),
		"form": uint(0),
		"type": "electric",
	},
	"Regidrago": {
		"id":   uint(895),
		"form": uint(0),
		"type": "dragon",
	},
	"Glastrier": {
		"id":   uint(896),
		"form": uint(0),
		"type": "ice",
	},
	"Spectrier": {
		"id":   uint(897),
		"form": uint(0),
		"type": "ghost",
	},
	"Calyrex": {
		"id":   uint(898),
		"form": uint(0),
		"type": "psychic",
	},
	"Calyrex-Ice": {
		"id":   uint(898),
		"form": uint(1),
		"type": "psychic",
	},
	"Calyrex-Shadow": {
		"id":   uint(898),
		"form": uint(2),
		"type": "psychic",
	},
	"Syclant": {
		"id":   uint(10001),
		"form": uint(0),
		"type": "ice",
	},
	"Revenankh": {
		"id":   uint(10002),
		"form": uint(0),
		"type": "ghost",
	},
	"Pyroak": {
		"id":   uint(10003),
		"form": uint(0),
		"type": "fire",
	},
	"Fidgit": {
		"id":   uint(10004),
		"form": uint(0),
		"type": "poison",
	},
	"Stratagem": {
		"id":   uint(10005),
		"form": uint(0),
		"type": "rock",
	},
	"Arghonaut": {
		"id":   uint(10006),
		"form": uint(0),
		"type": "water",
	},
	"Kitsunoh": {
		"id":   uint(10007),
		"form": uint(0),
		"type": "steel",
	},
	"Cyclohm": {
		"id":   uint(10008),
		"form": uint(0),
		"type": "electric",
	},
	"Colossoil": {
		"id":   uint(10009),
		"form": uint(0),
		"type": "dark",
	},
	"Krilowatt": {
		"id":   uint(10010),
		"form": uint(0),
		"type": "electric",
	},
	"Voodoom": {
		"id":   uint(10011),
		"form": uint(0),
		"type": "fighting",
	},
	"Tomohawk": {
		"id":   uint(10012),
		"form": uint(0),
		"type": "flying",
	},
	"Necturna": {
		"id":   uint(10013),
		"form": uint(0),
		"type": "grass",
	},
	"Mollux": {
		"id":   uint(10014),
		"form": uint(0),
		"type": "fire",
	},
	"Aurumoth": {
		"id":   uint(10015),
		"form": uint(0),
		"type": "bug",
	},
	"Malaconda": {
		"id":   uint(10016),
		"form": uint(0),
		"type": "dark",
	},
	"Cawmodore": {
		"id":   uint(10017),
		"form": uint(0),
		"type": "steel",
	},
	"Volkraken": {
		"id":   uint(10018),
		"form": uint(0),
		"type": "water",
	},
	"Plasmanta": {
		"id":   uint(10019),
		"form": uint(0),
		"type": "electric",
	},
	"Naviathan": {
		"id":   uint(10020),
		"form": uint(0),
		"type": "water",
	},
	"Crucibelle-Mega": {
		"id":   uint(10021),
		"form": uint(1),
		"type": "rock",
	},
	"Crucibelle": {
		"id":   uint(10021),
		"form": uint(0),
		"type": "rock",
	},
	"Kerfluffle": {
		"id":   uint(10022),
		"form": uint(0),
		"type": "fairy",
	},
	"Pajantom": {
		"id":   uint(10023),
		"form": uint(0),
		"type": "dragon",
	},
	"Jumabo": {
		"id":   uint(10024),
		"form": uint(0),
		"type": "grass",
	},
	"Caribolt": {
		"id":   uint(10025),
		"form": uint(0),
		"type": "grass",
	},
	"Smokomodo": {
		"id":   uint(10026),
		"form": uint(0),
		"type": "fire",
	},
	"Snaelstrom": {
		"id":   uint(10027),
		"form": uint(0),
		"type": "water",
	},
	"Equilibra": {
		"id":   uint(10028),
		"form": uint(0),
		"type": "steel",
	},
	"Astrolotl": {
		"id":   uint(10029),
		"form": uint(0),
		"type": "fire",
	},
	"Miasmaw": {
		"id":   uint(10030),
		"form": uint(0),
		"type": "bug",
	},
}

var moveData = map[string]map[string]interface{}{
	"Acupressure": {
		"id":             uint(367),
		"type":           "normal",
		"classification": uint(0),
	},
	"After You": {
		"id":             uint(495),
		"type":           "normal",
		"classification": uint(0),
	},
	"Assist": {
		"id":             uint(274),
		"type":           "normal",
		"classification": uint(0),
	},
	"Attract": {
		"id":             uint(213),
		"type":           "normal",
		"classification": uint(0),
	},
	"Barrage": {
		"id":             uint(140),
		"type":           "normal",
		"classification": uint(1),
	},
	"Baton Pass": {
		"id":             uint(226),
		"type":           "normal",
		"classification": uint(0),
	},
	"Belly Drum": {
		"id":             uint(187),
		"type":           "normal",
		"classification": uint(0),
	},
	"Bestow": {
		"id":             uint(516),
		"type":           "normal",
		"classification": uint(0),
	},
	"Bide": {
		"id":             uint(117),
		"type":           "normal",
		"classification": uint(1),
	},
	"Bind": {
		"id":             uint(20),
		"type":           "normal",
		"classification": uint(1),
	},
	"Block": {
		"id":             uint(335),
		"type":           "normal",
		"classification": uint(0),
	},
	"Body Slam": {
		"id":             uint(34),
		"type":           "normal",
		"classification": uint(1),
	},
	"Boomburst": {
		"id":             uint(586),
		"type":           "normal",
		"classification": uint(2),
	},
	"Camouflage": {
		"id":             uint(293),
		"type":           "normal",
		"classification": uint(0),
	},
	"Captivate": {
		"id":             uint(445),
		"type":           "normal",
		"classification": uint(0),
	},
	"Celebrate": {
		"id":             uint(606),
		"type":           "normal",
		"classification": uint(0),
	},
	"Chip Away": {
		"id":             uint(498),
		"type":           "normal",
		"classification": uint(1),
	},
	"Comet Punch": {
		"id":             uint(4),
		"type":           "normal",
		"classification": uint(1),
	},
	"Confide": {
		"id":             uint(590),
		"type":           "normal",
		"classification": uint(0),
	},
	"Constrict": {
		"id":             uint(132),
		"type":           "normal",
		"classification": uint(1),
	},
	"Conversion": {
		"id":             uint(160),
		"type":           "normal",
		"classification": uint(0),
	},
	"Conversion uint(2": {
		"id":             uint(176),
		"type":           "normal",
		"classification": uint(0),
	},
	"Copycat": {
		"id":             uint(383),
		"type":           "normal",
		"classification": uint(0),
	},
	"Covet": {
		"id":             uint(343),
		"type":           "normal",
		"classification": uint(1),
	},
	"Crush Claw": {
		"id":             uint(306),
		"type":           "normal",
		"classification": uint(1),
	},
	"Crush Grip": {
		"id":             uint(462),
		"type":           "normal",
		"classification": uint(1),
	},
	"Cut": {
		"id":             uint(15),
		"type":           "normal",
		"classification": uint(1),
	},
	"Defense Curl": {
		"id":             uint(111),
		"type":           "normal",
		"classification": uint(0),
	},
	"Disable": {
		"id":             uint(50),
		"type":           "normal",
		"classification": uint(0),
	},
	"Dizzy Punch": {
		"id":             uint(146),
		"type":           "normal",
		"classification": uint(1),
	},
	"Double Hit": {
		"id":             uint(458),
		"type":           "normal",
		"classification": uint(1),
	},
	"Double Slap": {
		"id":             uint(3),
		"type":           "normal",
		"classification": uint(1),
	},
	"Double Team": {
		"id":             uint(104),
		"type":           "normal",
		"classification": uint(0),
	},
	"Double-Edge": {
		"id":             uint(38),
		"type":           "normal",
		"classification": uint(1),
	},
	"Echoed Voice": {
		"id":             uint(497),
		"type":           "normal",
		"classification": uint(2),
	},
	"Egg Bomb": {
		"id":             uint(121),
		"type":           "normal",
		"classification": uint(1),
	},
	"Encore": {
		"id":             uint(227),
		"type":           "normal",
		"classification": uint(0),
	},
	"Endeavor": {
		"id":             uint(283),
		"type":           "normal",
		"classification": uint(1),
	},
	"Endure": {
		"id":             uint(203),
		"type":           "normal",
		"classification": uint(0),
	},
	"Entrainment": {
		"id":             uint(494),
		"type":           "normal",
		"classification": uint(0),
	},
	"Explosion": {
		"id":             uint(153),
		"type":           "normal",
		"classification": uint(1),
	},
	"Extreme Speed": {
		"id":             uint(245),
		"type":           "normal",
		"classification": uint(1),
	},
	"Facade": {
		"id":             uint(263),
		"type":           "normal",
		"classification": uint(1),
	},
	"Fake Out": {
		"id":             uint(252),
		"type":           "normal",
		"classification": uint(1),
	},
	"False Swipe": {
		"id":             uint(206),
		"type":           "normal",
		"classification": uint(1),
	},
	"Feint": {
		"id":             uint(364),
		"type":           "normal",
		"classification": uint(1),
	},
	"Flail": {
		"id":             uint(175),
		"type":           "normal",
		"classification": uint(1),
	},
	"Flash": {
		"id":             uint(148),
		"type":           "normal",
		"classification": uint(0),
	},
	"Focus Energy": {
		"id":             uint(116),
		"type":           "normal",
		"classification": uint(0),
	},
	"Follow Me": {
		"id":             uint(266),
		"type":           "normal",
		"classification": uint(0),
	},
	"Foresight": {
		"id":             uint(193),
		"type":           "normal",
		"classification": uint(0),
	},
	"Frustration": {
		"id":             uint(218),
		"type":           "normal",
		"classification": uint(1),
	},
	"Fury Attack": {
		"id":             uint(31),
		"type":           "normal",
		"classification": uint(1),
	},
	"Fury Swipes": {
		"id":             uint(154),
		"type":           "normal",
		"classification": uint(1),
	},
	"Giga Impact": {
		"id":             uint(416),
		"type":           "normal",
		"classification": uint(1),
	},
	"Glare": {
		"id":             uint(137),
		"type":           "normal",
		"classification": uint(0),
	},
	"Growl": {
		"id":             uint(45),
		"type":           "normal",
		"classification": uint(0),
	},
	"Growth": {
		"id":             uint(74),
		"type":           "normal",
		"classification": uint(0),
	},
	"Guillotine": {
		"id":             uint(12),
		"type":           "normal",
		"classification": uint(1),
	},
	"Happy Hour": {
		"id":             uint(603),
		"type":           "normal",
		"classification": uint(0),
	},
	"Harden": {
		"id":             uint(106),
		"type":           "normal",
		"classification": uint(0),
	},
	"Head Charge": {
		"id":             uint(543),
		"type":           "normal",
		"classification": uint(1),
	},
	"Headbutt": {
		"id":             uint(29),
		"type":           "normal",
		"classification": uint(1),
	},
	"Heal Bell": {
		"id":             uint(215),
		"type":           "normal",
		"classification": uint(0),
	},
	"Helping Hand": {
		"id":             uint(270),
		"type":           "normal",
		"classification": uint(0),
	},
	"Hidden Power": {
		"id":             uint(237),
		"type":           "normal",
		"classification": uint(2),
	},
	"Hidden Power Fighting": {
		"id":             uint(237),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Hidden Power Flying": {
		"id":             uint(237),
		"type":           "flying",
		"classification": uint(2),
	},
	"Hidden Power Poison": {
		"id":             uint(237),
		"type":           "poison",
		"classification": uint(2),
	},
	"Hidden Power Ground": {
		"id":             uint(237),
		"type":           "ground",
		"classification": uint(2),
	},
	"Hidden Power Rock": {
		"id":             uint(237),
		"type":           "rock",
		"classification": uint(2),
	},
	"Hidden Power Bug": {
		"id":             uint(237),
		"type":           "bug",
		"classification": uint(2),
	},
	"Hidden Power Ghost": {
		"id":             uint(237),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Hidden Power Steel": {
		"id":             uint(237),
		"type":           "steel",
		"classification": uint(2),
	},
	"Hidden Power Fire": {
		"id":             uint(237),
		"type":           "fire",
		"classification": uint(2),
	},
	"Hidden Power Water": {
		"id":             uint(237),
		"type":           "water",
		"classification": uint(2),
	},
	"Hidden Power Grass": {
		"id":             uint(237),
		"type":           "grass",
		"classification": uint(2),
	},
	"Hidden Power Electric": {
		"id":             uint(237),
		"type":           "electric",
		"classification": uint(2),
	},
	"Hidden Power Psychic": {
		"id":             uint(237),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Hidden Power Ice": {
		"id":             uint(237),
		"type":           "ice",
		"classification": uint(2),
	},
	"Hidden Power Dragon": {
		"id":             uint(237),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Hidden Power Dark": {
		"id":             uint(237),
		"type":           "dark",
		"classification": uint(2),
	},
	"Hidden Power Fairy": {
		"id":             uint(237),
		"type":           "normal",
		"classification": uint(2),
	},
	"Hold Back": {
		"id":             uint(610),
		"type":           "normal",
		"classification": uint(1),
	},
	"Hold Hands": {
		"id":             uint(607),
		"type":           "normal",
		"classification": uint(0),
	},
	"Horn Attack": {
		"id":             uint(30),
		"type":           "normal",
		"classification": uint(1),
	},
	"Horn Drill": {
		"id":             uint(32),
		"type":           "normal",
		"classification": uint(1),
	},
	"Howl": {
		"id":             uint(336),
		"type":           "normal",
		"classification": uint(0),
	},
	"Hyper Beam": {
		"id":             uint(63),
		"type":           "normal",
		"classification": uint(2),
	},
	"Hyper Fang": {
		"id":             uint(158),
		"type":           "normal",
		"classification": uint(1),
	},
	"Hyper Voice": {
		"id":             uint(304),
		"type":           "normal",
		"classification": uint(2),
	},
	"Judgment": {
		"id":             uint(449),
		"type":           "normal",
		"classification": uint(2),
	},
	"Laser Focus": {
		"id":             uint(673),
		"type":           "normal",
		"classification": uint(0),
	},
	"Last Resort": {
		"id":             uint(387),
		"type":           "normal",
		"classification": uint(1),
	},
	"Leer": {
		"id":             uint(43),
		"type":           "normal",
		"classification": uint(0),
	},
	"Lock-On": {
		"id":             uint(199),
		"type":           "normal",
		"classification": uint(0),
	},
	"Lovely Kiss": {
		"id":             uint(142),
		"type":           "normal",
		"classification": uint(0),
	},
	"Lucky Chant": {
		"id":             uint(381),
		"type":           "normal",
		"classification": uint(0),
	},
	"Me First": {
		"id":             uint(382),
		"type":           "normal",
		"classification": uint(0),
	},
	"Mean Look": {
		"id":             uint(212),
		"type":           "normal",
		"classification": uint(0),
	},
	"Mega Kick": {
		"id":             uint(25),
		"type":           "normal",
		"classification": uint(1),
	},
	"Mega Punch": {
		"id":             uint(5),
		"type":           "normal",
		"classification": uint(1),
	},
	"Metronome": {
		"id":             uint(118),
		"type":           "normal",
		"classification": uint(0),
	},
	"Milk Drink": {
		"id":             uint(208),
		"type":           "normal",
		"classification": uint(0),
	},
	"Mimic": {
		"id":             uint(102),
		"type":           "normal",
		"classification": uint(0),
	},
	"Mind Blown": {
		"id":             uint(720),
		"type":           "fire",
		"classification": uint(2),
	},
	"Mind Reader": {
		"id":             uint(170),
		"type":           "normal",
		"classification": uint(0),
	},
	"Minimize": {
		"id":             uint(107),
		"type":           "normal",
		"classification": uint(0),
	},
	"Morning Sun": {
		"id":             uint(234),
		"type":           "normal",
		"classification": uint(0),
	},
	"Multi-Attack": {
		"id":             uint(718),
		"type":           "normal",
		"classification": uint(1),
	},
	"Natural Gift": {
		"id":             uint(363),
		"type":           "normal",
		"classification": uint(1),
	},
	"Nature Power": {
		"id":             uint(267),
		"type":           "normal",
		"classification": uint(0),
	},
	"Noble Roar": {
		"id":             uint(568),
		"type":           "normal",
		"classification": uint(0),
	},
	"Odor Sleuth": {
		"id":             uint(316),
		"type":           "normal",
		"classification": uint(0),
	},
	"Pain Split": {
		"id":             uint(220),
		"type":           "normal",
		"classification": uint(0),
	},
	"Pay Day": {
		"id":             uint(6),
		"type":           "normal",
		"classification": uint(1),
	},
	"Perish Song": {
		"id":             uint(195),
		"type":           "normal",
		"classification": uint(0),
	},
	"Photon Geyser": {
		"id":             uint(722),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Plasma Fists": {
		"id":             uint(721),
		"type":           "electric",
		"classification": uint(1),
	},
	"Play Nice": {
		"id":             uint(589),
		"type":           "normal",
		"classification": uint(0),
	},
	"Pound": {
		"id":             uint(1),
		"type":           "normal",
		"classification": uint(1),
	},
	"Present": {
		"id":             uint(217),
		"type":           "normal",
		"classification": uint(1),
	},
	"Protect": {
		"id":             uint(182),
		"type":           "normal",
		"classification": uint(0),
	},
	"Psych Up": {
		"id":             uint(244),
		"type":           "normal",
		"classification": uint(0),
	},
	"Quick Attack": {
		"id":             uint(98),
		"type":           "normal",
		"classification": uint(1),
	},
	"Rage": {
		"id":             uint(99),
		"type":           "normal",
		"classification": uint(1),
	},
	"Rapid Spin": {
		"id":             uint(229),
		"type":           "normal",
		"classification": uint(1),
	},
	"Razor Wind": {
		"id":             uint(13),
		"type":           "normal",
		"classification": uint(2),
	},
	"Recover": {
		"id":             uint(105),
		"type":           "normal",
		"classification": uint(0),
	},
	"Recycle": {
		"id":             uint(278),
		"type":           "normal",
		"classification": uint(0),
	},
	"Reflect Type": {
		"id":             uint(513),
		"type":           "normal",
		"classification": uint(0),
	},
	"Refresh": {
		"id":             uint(287),
		"type":           "normal",
		"classification": uint(0),
	},
	"Relic Song": {
		"id":             uint(547),
		"type":           "normal",
		"classification": uint(2),
	},
	"Retaliate": {
		"id":             uint(514),
		"type":           "normal",
		"classification": uint(1),
	},
	"Return": {
		"id":             uint(216),
		"type":           "normal",
		"classification": uint(1),
	},
	"Revelation Dance": {
		"id":             uint(686),
		"type":           "normal",
		"classification": uint(2),
	},
	"Roar": {
		"id":             uint(46),
		"type":           "normal",
		"classification": uint(0),
	},
	"Rock Climb": {
		"id":             uint(431),
		"type":           "normal",
		"classification": uint(1),
	},
	"Round": {
		"id":             uint(496),
		"type":           "normal",
		"classification": uint(2),
	},
	"Safeguard": {
		"id":             uint(219),
		"type":           "normal",
		"classification": uint(0),
	},
	"Scary Face": {
		"id":             uint(184),
		"type":           "normal",
		"classification": uint(0),
	},
	"Scratch": {
		"id":             uint(10),
		"type":           "normal",
		"classification": uint(1),
	},
	"Screech": {
		"id":             uint(103),
		"type":           "normal",
		"classification": uint(0),
	},
	"Secret Power": {
		"id":             uint(290),
		"type":           "normal",
		"classification": uint(1),
	},
	"Self-Destruct": {
		"id":             uint(120),
		"type":           "normal",
		"classification": uint(1),
	},
	"Sharpen": {
		"id":             uint(159),
		"type":           "normal",
		"classification": uint(0),
	},
	"Shell Smash": {
		"id":             uint(504),
		"type":           "normal",
		"classification": uint(0),
	},
	"Simple Beam": {
		"id":             uint(493),
		"type":           "normal",
		"classification": uint(0),
	},
	"Sing": {
		"id":             uint(47),
		"type":           "normal",
		"classification": uint(0),
	},
	"Sketch": {
		"id":             uint(166),
		"type":           "normal",
		"classification": uint(0),
	},
	"Skull Bash": {
		"id":             uint(130),
		"type":           "normal",
		"classification": uint(1),
	},
	"Slack Off": {
		"id":             uint(303),
		"type":           "normal",
		"classification": uint(0),
	},
	"Slam": {
		"id":             uint(21),
		"type":           "normal",
		"classification": uint(1),
	},
	"Slash": {
		"id":             uint(163),
		"type":           "normal",
		"classification": uint(1),
	},
	"Sleep Talk": {
		"id":             uint(214),
		"type":           "normal",
		"classification": uint(0),
	},
	"Smelling Salts": {
		"id":             uint(265),
		"type":           "normal",
		"classification": uint(1),
	},
	"Smokescreen": {
		"id":             uint(108),
		"type":           "normal",
		"classification": uint(0),
	},
	"Snore": {
		"id":             uint(173),
		"type":           "normal",
		"classification": uint(2),
	},
	"Soft-Boiled": {
		"id":             uint(135),
		"type":           "normal",
		"classification": uint(0),
	},
	"Sonic Boom": {
		"id":             uint(49),
		"type":           "normal",
		"classification": uint(2),
	},
	"Spectral Thief": {
		"id":             uint(712),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Spike Cannon": {
		"id":             uint(131),
		"type":           "normal",
		"classification": uint(1),
	},
	"Spit Up": {
		"id":             uint(255),
		"type":           "normal",
		"classification": uint(2),
	},
	"Splash": {
		"id":             uint(150),
		"type":           "normal",
		"classification": uint(0),
	},
	"Spotlight": {
		"id":             uint(671),
		"type":           "normal",
		"classification": uint(0),
	},
	"Stockpile": {
		"id":             uint(254),
		"type":           "normal",
		"classification": uint(0),
	},
	"Stomp": {
		"id":             uint(23),
		"type":           "normal",
		"classification": uint(1),
	},
	"Strength": {
		"id":             uint(70),
		"type":           "normal",
		"classification": uint(1),
	},
	"Substitute": {
		"id":             uint(164),
		"type":           "normal",
		"classification": uint(0),
	},
	"Super Fang": {
		"id":             uint(162),
		"type":           "normal",
		"classification": uint(1),
	},
	"Supersonic": {
		"id":             uint(48),
		"type":           "normal",
		"classification": uint(0),
	},
	"Swagger": {
		"id":             uint(207),
		"type":           "normal",
		"classification": uint(0),
	},
	"Swallow": {
		"id":             uint(256),
		"type":           "normal",
		"classification": uint(0),
	},
	"Sweet Scent": {
		"id":             uint(230),
		"type":           "normal",
		"classification": uint(0),
	},
	"Swift": {
		"id":             uint(129),
		"type":           "normal",
		"classification": uint(2),
	},
	"Swords Dance": {
		"id":             uint(14),
		"type":           "normal",
		"classification": uint(0),
	},
	"Tackle": {
		"id":             uint(33),
		"type":           "normal",
		"classification": uint(1),
	},
	"Tail Slap": {
		"id":             uint(541),
		"type":           "normal",
		"classification": uint(1),
	},
	"Tail Whip": {
		"id":             uint(39),
		"type":           "normal",
		"classification": uint(0),
	},
	"Take Down": {
		"id":             uint(36),
		"type":           "normal",
		"classification": uint(1),
	},
	"Tearful Look": {
		"id":             uint(715),
		"type":           "normal",
		"classification": uint(0),
	},
	"Techno Blast": {
		"id":             uint(546),
		"type":           "normal",
		"classification": uint(2),
	},
	"Teeter Dance": {
		"id":             uint(298),
		"type":           "normal",
		"classification": uint(0),
	},
	"Thrash": {
		"id":             uint(37),
		"type":           "normal",
		"classification": uint(1),
	},
	"Tickle": {
		"id":             uint(321),
		"type":           "normal",
		"classification": uint(0),
	},
	"Transform": {
		"id":             uint(144),
		"type":           "normal",
		"classification": uint(0),
	},
	"Tri Attack": {
		"id":             uint(161),
		"type":           "normal",
		"classification": uint(2),
	},
	"Trump Card": {
		"id":             uint(376),
		"type":           "normal",
		"classification": uint(2),
	},
	"Uproar": {
		"id":             uint(253),
		"type":           "normal",
		"classification": uint(2),
	},
	"Vice Grip": {
		"id":             uint(11),
		"type":           "normal",
		"classification": uint(1),
	},
	"Weather Ball": {
		"id":             uint(311),
		"type":           "normal",
		"classification": uint(2),
	},
	"Whirlwind": {
		"id":             uint(18),
		"type":           "normal",
		"classification": uint(0),
	},
	"Wish": {
		"id":             uint(273),
		"type":           "normal",
		"classification": uint(0),
	},
	"Work Up": {
		"id":             uint(526),
		"type":           "normal",
		"classification": uint(0),
	},
	"Wrap": {
		"id":             uint(35),
		"type":           "normal",
		"classification": uint(1),
	},
	"Wring Out": {
		"id":             uint(378),
		"type":           "normal",
		"classification": uint(2),
	},
	"Yawn": {
		"id":             uint(281),
		"type":           "normal",
		"classification": uint(0),
	},
	"Arm Thrust": {
		"id":             uint(292),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Aura Sphere": {
		"id":             uint(396),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Brick Break": {
		"id":             uint(280),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Bulk Up": {
		"id":             uint(339),
		"type":           "fighting",
		"classification": uint(0),
	},
	"Circle Throw": {
		"id":             uint(509),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Close Combat": {
		"id":             uint(370),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Counter": {
		"id":             uint(68),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Cross Chop": {
		"id":             uint(238),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Detect": {
		"id":             uint(197),
		"type":           "fighting",
		"classification": uint(0),
	},
	"Double Kick": {
		"id":             uint(24),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Drain Punch": {
		"id":             uint(409),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Dynamic Punch": {
		"id":             uint(223),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Final Gambit": {
		"id":             uint(515),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Flying Press": {
		"id":             uint(560),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Focus Blast": {
		"id":             uint(411),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Focus Punch": {
		"id":             uint(264),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Force Palm": {
		"id":             uint(395),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Hammer Arm": {
		"id":             uint(359),
		"type":           "fighting",
		"classification": uint(1),
	},
	"High Jump Kick": {
		"id":             uint(136),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Jump Kick": {
		"id":             uint(26),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Karate Chop": {
		"id":             uint(2),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Low Kick": {
		"id":             uint(67),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Low Sweep": {
		"id":             uint(490),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Mach Punch": {
		"id":             uint(183),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Mat Block": {
		"id":             uint(561),
		"type":           "fighting",
		"classification": uint(0),
	},
	"Power-Up Punch": {
		"id":             uint(612),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Quick Guard": {
		"id":             uint(501),
		"type":           "fighting",
		"classification": uint(0),
	},
	"Revenge": {
		"id":             uint(279),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Reversal": {
		"id":             uint(179),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Rock Smash": {
		"id":             uint(249),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Rolling Kick": {
		"id":             uint(27),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Sacred Sword": {
		"id":             uint(533),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Secret Sword": {
		"id":             uint(548),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Seismic Toss": {
		"id":             uint(69),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Sky Uppercut": {
		"id":             uint(327),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Storm Throw": {
		"id":             uint(480),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Submission": {
		"id":             uint(66),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Superpower": {
		"id":             uint(276),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Triple Kick": {
		"id":             uint(167),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Vacuum Wave": {
		"id":             uint(410),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Vital Throw": {
		"id":             uint(233),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Wake-Up Slap": {
		"id":             uint(358),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Blast Burn": {
		"id":             uint(307),
		"type":           "fire",
		"classification": uint(2),
	},
	"Blaze Kick": {
		"id":             uint(299),
		"type":           "fire",
		"classification": uint(1),
	},
	"Blue Flare": {
		"id":             uint(551),
		"type":           "fire",
		"classification": uint(2),
	},
	"Burn Up": {
		"id":             uint(682),
		"type":           "fire",
		"classification": uint(2),
	},
	"Ember": {
		"id":             uint(52),
		"type":           "fire",
		"classification": uint(2),
	},
	"Eruption": {
		"id":             uint(284),
		"type":           "fire",
		"classification": uint(2),
	},
	"Fiery Dance": {
		"id":             uint(552),
		"type":           "fire",
		"classification": uint(2),
	},
	"Fire Blast": {
		"id":             uint(126),
		"type":           "fire",
		"classification": uint(2),
	},
	"Fire Fang": {
		"id":             uint(424),
		"type":           "fire",
		"classification": uint(1),
	},
	"Fire Lash": {
		"id":             uint(680),
		"type":           "fire",
		"classification": uint(1),
	},
	"Fire Pledge": {
		"id":             uint(519),
		"type":           "fire",
		"classification": uint(2),
	},
	"Fire Punch": {
		"id":             uint(7),
		"type":           "fire",
		"classification": uint(1),
	},
	"Fire Spin": {
		"id":             uint(83),
		"type":           "fire",
		"classification": uint(2),
	},
	"Flame Burst": {
		"id":             uint(481),
		"type":           "fire",
		"classification": uint(2),
	},
	"Flame Charge": {
		"id":             uint(488),
		"type":           "fire",
		"classification": uint(1),
	},
	"Flame Wheel": {
		"id":             uint(172),
		"type":           "fire",
		"classification": uint(1),
	},
	"Flamethrower": {
		"id":             uint(53),
		"type":           "fire",
		"classification": uint(2),
	},
	"Flare Blitz": {
		"id":             uint(394),
		"type":           "fire",
		"classification": uint(1),
	},
	"Fusion Flare": {
		"id":             uint(558),
		"type":           "fire",
		"classification": uint(2),
	},
	"Heat Crash": {
		"id":             uint(535),
		"type":           "fire",
		"classification": uint(1),
	},
	"Heat Wave": {
		"id":             uint(257),
		"type":           "fire",
		"classification": uint(2),
	},
	"Incinerate": {
		"id":             uint(510),
		"type":           "fire",
		"classification": uint(2),
	},
	"Inferno": {
		"id":             uint(517),
		"type":           "fire",
		"classification": uint(2),
	},
	"Lava Plume": {
		"id":             uint(436),
		"type":           "fire",
		"classification": uint(2),
	},
	"Magma Storm": {
		"id":             uint(463),
		"type":           "fire",
		"classification": uint(2),
	},
	"Mystical Fire": {
		"id":             uint(595),
		"type":           "fire",
		"classification": uint(2),
	},
	"Overheat": {
		"id":             uint(315),
		"type":           "fire",
		"classification": uint(2),
	},
	"Sacred Fire": {
		"id":             uint(221),
		"type":           "fire",
		"classification": uint(1),
	},
	"Searing Shot": {
		"id":             uint(545),
		"type":           "fire",
		"classification": uint(2),
	},
	"Shell Trap": {
		"id":             uint(704),
		"type":           "fire",
		"classification": uint(2),
	},
	"Sunny Day": {
		"id":             uint(241),
		"type":           "fire",
		"classification": uint(0),
	},
	"V-create": {
		"id":             uint(557),
		"type":           "fire",
		"classification": uint(1),
	},
	"Will-O-Wisp": {
		"id":             uint(261),
		"type":           "fire",
		"classification": uint(0),
	},
	"Aurora Beam": {
		"id":             uint(62),
		"type":           "ice",
		"classification": uint(2),
	},
	"Aurora Veil": {
		"id":             uint(694),
		"type":           "ice",
		"classification": uint(0),
	},
	"Avalanche": {
		"id":             uint(419),
		"type":           "ice",
		"classification": uint(1),
	},
	"Blizzard": {
		"id":             uint(59),
		"type":           "ice",
		"classification": uint(2),
	},
	"Freeze Shock": {
		"id":             uint(553),
		"type":           "ice",
		"classification": uint(1),
	},
	"Freeze-Dry": {
		"id":             uint(573),
		"type":           "ice",
		"classification": uint(2),
	},
	"Frost Breath": {
		"id":             uint(524),
		"type":           "ice",
		"classification": uint(2),
	},
	"Glaciate": {
		"id":             uint(549),
		"type":           "ice",
		"classification": uint(2),
	},
	"Hail": {
		"id":             uint(258),
		"type":           "ice",
		"classification": uint(0),
	},
	"Haze": {
		"id":             uint(114),
		"type":           "ice",
		"classification": uint(0),
	},
	"Ice Ball": {
		"id":             uint(301),
		"type":           "ice",
		"classification": uint(1),
	},
	"Ice Beam": {
		"id":             uint(58),
		"type":           "ice",
		"classification": uint(2),
	},
	"Ice Burn": {
		"id":             uint(554),
		"type":           "ice",
		"classification": uint(2),
	},
	"Ice Fang": {
		"id":             uint(423),
		"type":           "ice",
		"classification": uint(1),
	},
	"Ice Hammer": {
		"id":             uint(665),
		"type":           "ice",
		"classification": uint(1),
	},
	"Ice Punch": {
		"id":             uint(8),
		"type":           "ice",
		"classification": uint(1),
	},
	"Ice Shard": {
		"id":             uint(420),
		"type":           "ice",
		"classification": uint(1),
	},
	"Icicle Crash": {
		"id":             uint(556),
		"type":           "ice",
		"classification": uint(1),
	},
	"Icicle Spear": {
		"id":             uint(333),
		"type":           "ice",
		"classification": uint(1),
	},
	"Icy Wind": {
		"id":             uint(196),
		"type":           "ice",
		"classification": uint(2),
	},
	"Mist": {
		"id":             uint(54),
		"type":           "ice",
		"classification": uint(0),
	},
	"Powder Snow": {
		"id":             uint(181),
		"type":           "ice",
		"classification": uint(2),
	},
	"Sheer Cold": {
		"id":             uint(329),
		"type":           "ice",
		"classification": uint(2),
	},
	"Bolt Strike": {
		"id":             uint(550),
		"type":           "electric",
		"classification": uint(1),
	},
	"Charge": {
		"id":             uint(268),
		"type":           "electric",
		"classification": uint(0),
	},
	"Charge Beam": {
		"id":             uint(451),
		"type":           "electric",
		"classification": uint(2),
	},
	"Discharge": {
		"id":             uint(435),
		"type":           "electric",
		"classification": uint(2),
	},
	"Eerie Impulse": {
		"id":             uint(598),
		"type":           "electric",
		"classification": uint(0),
	},
	"Electric Terrain": {
		"id":             uint(604),
		"type":           "electric",
		"classification": uint(0),
	},
	"Electrify": {
		"id":             uint(582),
		"type":           "electric",
		"classification": uint(0),
	},
	"Electro Ball": {
		"id":             uint(486),
		"type":           "electric",
		"classification": uint(2),
	},
	"Electroweb": {
		"id":             uint(527),
		"type":           "electric",
		"classification": uint(2),
	},
	"Fusion Bolt": {
		"id":             uint(559),
		"type":           "electric",
		"classification": uint(1),
	},
	"Ion Deluge": {
		"id":             uint(569),
		"type":           "electric",
		"classification": uint(0),
	},
	"Magnet Rise": {
		"id":             uint(393),
		"type":           "electric",
		"classification": uint(0),
	},
	"Magnetic Flux": {
		"id":             uint(602),
		"type":           "electric",
		"classification": uint(0),
	},
	"Nuzzle": {
		"id":             uint(609),
		"type":           "electric",
		"classification": uint(1),
	},
	"Parabolic Charge": {
		"id":             uint(570),
		"type":           "electric",
		"classification": uint(2),
	},
	"Shock Wave": {
		"id":             uint(351),
		"type":           "electric",
		"classification": uint(2),
	},
	"Spark": {
		"id":             uint(209),
		"type":           "electric",
		"classification": uint(1),
	},
	"Thunder": {
		"id":             uint(87),
		"type":           "electric",
		"classification": uint(2),
	},
	"Thunder Fang": {
		"id":             uint(422),
		"type":           "electric",
		"classification": uint(1),
	},
	"Thunder Punch": {
		"id":             uint(9),
		"type":           "electric",
		"classification": uint(1),
	},
	"Thunder Shock": {
		"id":             uint(84),
		"type":           "electric",
		"classification": uint(2),
	},
	"Thunder Wave": {
		"id":             uint(86),
		"type":           "electric",
		"classification": uint(0),
	},
	"Thunderbolt": {
		"id":             uint(85),
		"type":           "electric",
		"classification": uint(2),
	},
	"Volt Switch": {
		"id":             uint(521),
		"type":           "electric",
		"classification": uint(2),
	},
	"Volt Tackle": {
		"id":             uint(344),
		"type":           "electric",
		"classification": uint(1),
	},
	"Wild Charge": {
		"id":             uint(528),
		"type":           "electric",
		"classification": uint(1),
	},
	"Zap Cannon": {
		"id":             uint(192),
		"type":           "electric",
		"classification": uint(2),
	},
	"Zing Zap": {
		"id":             uint(716),
		"type":           "electric",
		"classification": uint(1),
	},
	"Acrobatics": {
		"id":             uint(512),
		"type":           "flying",
		"classification": uint(1),
	},
	"Aerial Ace": {
		"id":             uint(332),
		"type":           "flying",
		"classification": uint(1),
	},
	"Aeroblast": {
		"id":             uint(177),
		"type":           "flying",
		"classification": uint(2),
	},
	"Air Cutter": {
		"id":             uint(314),
		"type":           "flying",
		"classification": uint(2),
	},
	"Air Slash": {
		"id":             uint(403),
		"type":           "flying",
		"classification": uint(2),
	},
	"Beak Blast": {
		"id":             uint(690),
		"type":           "flying",
		"classification": uint(1),
	},
	"Bounce": {
		"id":             uint(340),
		"type":           "flying",
		"classification": uint(1),
	},
	"Brave Bird": {
		"id":             uint(413),
		"type":           "flying",
		"classification": uint(1),
	},
	"Chatter": {
		"id":             uint(448),
		"type":           "flying",
		"classification": uint(2),
	},
	"Defog": {
		"id":             uint(432),
		"type":           "flying",
		"classification": uint(3),
	},
	"Dragon Ascent": {
		"id":             uint(620),
		"type":           "flying",
		"classification": uint(1),
	},
	"Drill Peck": {
		"id":             uint(65),
		"type":           "flying",
		"classification": uint(1),
	},
	"Feather Dance": {
		"id":             uint(297),
		"type":           "flying",
		"classification": uint(0),
	},
	"Fly": {
		"id":             uint(19),
		"type":           "flying",
		"classification": uint(1),
	},
	"Gust": {
		"id":             uint(16),
		"type":           "flying",
		"classification": uint(2),
	},
	"Hurricane": {
		"id":             uint(542),
		"type":           "flying",
		"classification": uint(2),
	},
	"Mirror Move": {
		"id":             uint(119),
		"type":           "flying",
		"classification": uint(0),
	},
	"Oblivion Wing": {
		"id":             uint(613),
		"type":           "flying",
		"classification": uint(2),
	},
	"Peck": {
		"id":             uint(64),
		"type":           "flying",
		"classification": uint(1),
	},
	"Pluck": {
		"id":             uint(365),
		"type":           "flying",
		"classification": uint(1),
	},
	"Roost": {
		"id":             uint(355),
		"type":           "flying",
		"classification": uint(0),
	},
	"Sky Attack": {
		"id":             uint(143),
		"type":           "flying",
		"classification": uint(1),
	},
	"Sky Drop": {
		"id":             uint(507),
		"type":           "flying",
		"classification": uint(1),
	},
	"Tailwind": {
		"id":             uint(366),
		"type":           "flying",
		"classification": uint(0),
	},
	"Wing Attack": {
		"id":             uint(17),
		"type":           "flying",
		"classification": uint(1),
	},
	"Absorb": {
		"id":             uint(71),
		"type":           "grass",
		"classification": uint(2),
	},
	"Aromatherapy": {
		"id":             uint(312),
		"type":           "grass",
		"classification": uint(0),
	},
	"Bullet Seed": {
		"id":             uint(331),
		"type":           "grass",
		"classification": uint(1),
	},
	"Cotton Guard": {
		"id":             uint(538),
		"type":           "grass",
		"classification": uint(0),
	},
	"Cotton Spore": {
		"id":             uint(178),
		"type":           "grass",
		"classification": uint(0),
	},
	"Energy Ball": {
		"id":             uint(412),
		"type":           "grass",
		"classification": uint(2),
	},
	"Forest's Curse": {
		"id":             uint(571),
		"type":           "grass",
		"classification": uint(0),
	},
	"Frenzy Plant": {
		"id":             uint(338),
		"type":           "grass",
		"classification": uint(2),
	},
	"Giga Drain": {
		"id":             uint(202),
		"type":           "grass",
		"classification": uint(2),
	},
	"Grass Knot": {
		"id":             uint(447),
		"type":           "grass",
		"classification": uint(2),
	},
	"Grass Pledge": {
		"id":             uint(520),
		"type":           "grass",
		"classification": uint(2),
	},
	"Grass Whistle": {
		"id":             uint(320),
		"type":           "grass",
		"classification": uint(0),
	},
	"Grassy Terrain": {
		"id":             uint(580),
		"type":           "grass",
		"classification": uint(0),
	},
	"Horn Leech": {
		"id":             uint(532),
		"type":           "grass",
		"classification": uint(1),
	},
	"Ingrain": {
		"id":             uint(275),
		"type":           "grass",
		"classification": uint(0),
	},
	"Leaf Blade": {
		"id":             uint(348),
		"type":           "grass",
		"classification": uint(1),
	},
	"Leaf Storm": {
		"id":             uint(437),
		"type":           "grass",
		"classification": uint(2),
	},
	"Leaf Tornado": {
		"id":             uint(536),
		"type":           "grass",
		"classification": uint(2),
	},
	"Leafage": {
		"id":             uint(670),
		"type":           "grass",
		"classification": uint(1),
	},
	"Leech Seed": {
		"id":             uint(73),
		"type":           "grass",
		"classification": uint(0),
	},
	"Magical Leaf": {
		"id":             uint(345),
		"type":           "grass",
		"classification": uint(2),
	},
	"Mega Drain": {
		"id":             uint(72),
		"type":           "grass",
		"classification": uint(2),
	},
	"Needle Arm": {
		"id":             uint(302),
		"type":           "grass",
		"classification": uint(1),
	},
	"Petal Blizzard": {
		"id":             uint(572),
		"type":           "grass",
		"classification": uint(1),
	},
	"Petal Dance": {
		"id":             uint(80),
		"type":           "grass",
		"classification": uint(2),
	},
	"Power Whip": {
		"id":             uint(438),
		"type":           "grass",
		"classification": uint(1),
	},
	"Razor Leaf": {
		"id":             uint(75),
		"type":           "grass",
		"classification": uint(1),
	},
	"Seed Bomb": {
		"id":             uint(402),
		"type":           "grass",
		"classification": uint(1),
	},
	"Seed Flare": {
		"id":             uint(465),
		"type":           "grass",
		"classification": uint(2),
	},
	"Sleep Powder": {
		"id":             uint(79),
		"type":           "grass",
		"classification": uint(0),
	},
	"Solar Beam": {
		"id":             uint(76),
		"type":           "grass",
		"classification": uint(2),
	},
	"Solar Blade": {
		"id":             uint(669),
		"type":           "grass",
		"classification": uint(1),
	},
	"Spiky Shield": {
		"id":             uint(596),
		"type":           "grass",
		"classification": uint(0),
	},
	"Spore": {
		"id":             uint(147),
		"type":           "grass",
		"classification": uint(0),
	},
	"Strength Sap": {
		"id":             uint(668),
		"type":           "grass",
		"classification": uint(0),
	},
	"Stun Spore": {
		"id":             uint(78),
		"type":           "grass",
		"classification": uint(0),
	},
	"Synthesis": {
		"id":             uint(235),
		"type":           "grass",
		"classification": uint(0),
	},
	"Trop Kick": {
		"id":             uint(688),
		"type":           "grass",
		"classification": uint(1),
	},
	"Vine Whip": {
		"id":             uint(22),
		"type":           "grass",
		"classification": uint(1),
	},
	"Wood Hammer": {
		"id":             uint(452),
		"type":           "grass",
		"classification": uint(1),
	},
	"Worry Seed": {
		"id":             uint(388),
		"type":           "grass",
		"classification": uint(0),
	},
	"Bone Club": {
		"id":             uint(125),
		"type":           "ground",
		"classification": uint(1),
	},
	"Bone Rush": {
		"id":             uint(198),
		"type":           "ground",
		"classification": uint(1),
	},
	"Bonemerang": {
		"id":             uint(155),
		"type":           "ground",
		"classification": uint(1),
	},
	"Bulldoze": {
		"id":             uint(523),
		"type":           "ground",
		"classification": uint(1),
	},
	"Dig": {
		"id":             uint(91),
		"type":           "ground",
		"classification": uint(1),
	},
	"Drill Run": {
		"id":             uint(529),
		"type":           "ground",
		"classification": uint(1),
	},
	"Earth Power": {
		"id":             uint(414),
		"type":           "ground",
		"classification": uint(2),
	},
	"Earthquake": {
		"id":             uint(89),
		"type":           "ground",
		"classification": uint(1),
	},
	"Fissure": {
		"id":             uint(90),
		"type":           "ground",
		"classification": uint(1),
	},
	"High Horsepower": {
		"id":             uint(667),
		"type":           "ground",
		"classification": uint(1),
	},
	"Land's Wrath": {
		"id":             uint(616),
		"type":           "ground",
		"classification": uint(1),
	},
	"Magnitude": {
		"id":             uint(222),
		"type":           "ground",
		"classification": uint(1),
	},
	"Mud Bomb": {
		"id":             uint(426),
		"type":           "ground",
		"classification": uint(2),
	},
	"Mud Shot": {
		"id":             uint(341),
		"type":           "ground",
		"classification": uint(2),
	},
	"Mud Sport": {
		"id":             uint(300),
		"type":           "ground",
		"classification": uint(0),
	},
	"Mud-Slap": {
		"id":             uint(189),
		"type":           "ground",
		"classification": uint(2),
	},
	"Precipice Blades": {
		"id":             uint(619),
		"type":           "ground",
		"classification": uint(1),
	},
	"Rototiller": {
		"id":             uint(563),
		"type":           "ground",
		"classification": uint(0),
	},
	"Sand Attack": {
		"id":             uint(28),
		"type":           "ground",
		"classification": uint(0),
	},
	"Sand Tomb": {
		"id":             uint(328),
		"type":           "ground",
		"classification": uint(1),
	},
	"Shore Up": {
		"id":             uint(659),
		"type":           "ground",
		"classification": uint(0),
	},
	"Spikes": {
		"id":             uint(191),
		"type":           "ground",
		"classification": uint(0),
	},
	"Stomping Tantrum": {
		"id":             uint(707),
		"type":           "ground",
		"classification": uint(1),
	},
	"Thousand Arrows": {
		"id":             uint(614),
		"type":           "ground",
		"classification": uint(1),
	},
	"Thousand Waves": {
		"id":             uint(615),
		"type":           "ground",
		"classification": uint(1),
	},
	"Acid": {
		"id":             uint(51),
		"type":           "poison",
		"classification": uint(2),
	},
	"Acid Armor": {
		"id":             uint(151),
		"type":           "poison",
		"classification": uint(0),
	},
	"Acid Spray": {
		"id":             uint(491),
		"type":           "poison",
		"classification": uint(2),
	},
	"Baneful Bunker": {
		"id":             uint(661),
		"type":           "poison",
		"classification": uint(0),
	},
	"Belch": {
		"id":             uint(562),
		"type":           "poison",
		"classification": uint(2),
	},
	"Clear Smog": {
		"id":             uint(499),
		"type":           "poison",
		"classification": uint(2),
	},
	"Coil": {
		"id":             uint(489),
		"type":           "poison",
		"classification": uint(0),
	},
	"Cross Poison": {
		"id":             uint(440),
		"type":           "poison",
		"classification": uint(1),
	},
	"Gastro Acid": {
		"id":             uint(380),
		"type":           "poison",
		"classification": uint(0),
	},
	"Gunk Shot": {
		"id":             uint(441),
		"type":           "poison",
		"classification": uint(1),
	},
	"Poison Fang": {
		"id":             uint(305),
		"type":           "poison",
		"classification": uint(1),
	},
	"Poison Gas": {
		"id":             uint(139),
		"type":           "poison",
		"classification": uint(0),
	},
	"Poison Jab": {
		"id":             uint(398),
		"type":           "poison",
		"classification": uint(1),
	},
	"Poison Powder": {
		"id":             uint(77),
		"type":           "poison",
		"classification": uint(0),
	},
	"Poison Sting": {
		"id":             uint(40),
		"type":           "poison",
		"classification": uint(1),
	},
	"Poison Tail": {
		"id":             uint(342),
		"type":           "poison",
		"classification": uint(1),
	},
	"Purify": {
		"id":             uint(685),
		"type":           "poison",
		"classification": uint(0),
	},
	"Sludge": {
		"id":             uint(124),
		"type":           "poison",
		"classification": uint(2),
	},
	"Sludge Bomb": {
		"id":             uint(188),
		"type":           "poison",
		"classification": uint(2),
	},
	"Sludge Wave": {
		"id":             uint(482),
		"type":           "poison",
		"classification": uint(2),
	},
	"Smog": {
		"id":             uint(123),
		"type":           "poison",
		"classification": uint(2),
	},
	"Toxic": {
		"id":             uint(92),
		"type":           "poison",
		"classification": uint(0),
	},
	"Toxic Spikes": {
		"id":             uint(390),
		"type":           "poison",
		"classification": uint(0),
	},
	"Toxic Thread": {
		"id":             uint(672),
		"type":           "poison",
		"classification": uint(0),
	},
	"Venom Drench": {
		"id":             uint(599),
		"type":           "poison",
		"classification": uint(0),
	},
	"Venoshock": {
		"id":             uint(474),
		"type":           "poison",
		"classification": uint(2),
	},
	"Attack Order": {
		"id":             uint(454),
		"type":           "bug",
		"classification": uint(1),
	},
	"Bug Bite": {
		"id":             uint(450),
		"type":           "bug",
		"classification": uint(1),
	},
	"Bug Buzz": {
		"id":             uint(405),
		"type":           "bug",
		"classification": uint(2),
	},
	"Defend Order": {
		"id":             uint(455),
		"type":           "bug",
		"classification": uint(0),
	},
	"Fell Stinger": {
		"id":             uint(565),
		"type":           "bug",
		"classification": uint(1),
	},
	"First Impression": {
		"id":             uint(660),
		"type":           "bug",
		"classification": uint(1),
	},
	"Fury Cutter": {
		"id":             uint(210),
		"type":           "bug",
		"classification": uint(1),
	},
	"Heal Order": {
		"id":             uint(456),
		"type":           "bug",
		"classification": uint(0),
	},
	"Infestation": {
		"id":             uint(611),
		"type":           "bug",
		"classification": uint(2),
	},
	"Leech Life": {
		"id":             uint(141),
		"type":           "bug",
		"classification": uint(1),
	},
	"Lunge": {
		"id":             uint(679),
		"type":           "bug",
		"classification": uint(1),
	},
	"Megahorn": {
		"id":             uint(224),
		"type":           "bug",
		"classification": uint(1),
	},
	"Pin Missile": {
		"id":             uint(42),
		"type":           "bug",
		"classification": uint(1),
	},
	"Pollen Puff": {
		"id":             uint(676),
		"type":           "bug",
		"classification": uint(2),
	},
	"Powder": {
		"id":             uint(600),
		"type":           "bug",
		"classification": uint(0),
	},
	"Quiver Dance": {
		"id":             uint(483),
		"type":           "bug",
		"classification": uint(0),
	},
	"Rage Powder": {
		"id":             uint(476),
		"type":           "bug",
		"classification": uint(0),
	},
	"Signal Beam": {
		"id":             uint(324),
		"type":           "bug",
		"classification": uint(2),
	},
	"Silver Wind": {
		"id":             uint(318),
		"type":           "bug",
		"classification": uint(2),
	},
	"Spider Web": {
		"id":             uint(169),
		"type":           "bug",
		"classification": uint(0),
	},
	"Steamroller": {
		"id":             uint(537),
		"type":           "bug",
		"classification": uint(1),
	},
	"Sticky Web": {
		"id":             uint(564),
		"type":           "bug",
		"classification": uint(0),
	},
	"String Shot": {
		"id":             uint(81),
		"type":           "bug",
		"classification": uint(0),
	},
	"Struggle Bug": {
		"id":             uint(522),
		"type":           "bug",
		"classification": uint(2),
	},
	"Tail Glow": {
		"id":             uint(294),
		"type":           "bug",
		"classification": uint(0),
	},
	"Twineedle": {
		"id":             uint(41),
		"type":           "bug",
		"classification": uint(1),
	},
	"U-turn": {
		"id":             uint(369),
		"type":           "bug",
		"classification": uint(1),
	},
	"X-Scissor": {
		"id":             uint(404),
		"type":           "bug",
		"classification": uint(1),
	},
	"Assurance": {
		"id":             uint(372),
		"type":           "dark",
		"classification": uint(1),
	},
	"Beat Up": {
		"id":             uint(251),
		"type":           "dark",
		"classification": uint(1),
	},
	"Bite": {
		"id":             uint(44),
		"type":           "dark",
		"classification": uint(1),
	},
	"Brutal Swing": {
		"id":             uint(693),
		"type":           "dark",
		"classification": uint(1),
	},
	"Crunch": {
		"id":             uint(242),
		"type":           "dark",
		"classification": uint(1),
	},
	"Dark Pulse": {
		"id":             uint(399),
		"type":           "dark",
		"classification": uint(2),
	},
	"Dark Void": {
		"id":             uint(464),
		"type":           "dark",
		"classification": uint(0),
	},
	"Darkest Lariat": {
		"id":             uint(663),
		"type":           "dark",
		"classification": uint(1),
	},
	"Embargo": {
		"id":             uint(373),
		"type":           "dark",
		"classification": uint(0),
	},
	"Fake Tears": {
		"id":             uint(313),
		"type":           "dark",
		"classification": uint(0),
	},
	"Feint Attack": {
		"id":             uint(185),
		"type":           "dark",
		"classification": uint(1),
	},
	"Flatter": {
		"id":             uint(260),
		"type":           "dark",
		"classification": uint(0),
	},
	"Fling": {
		"id":             uint(374),
		"type":           "dark",
		"classification": uint(1),
	},
	"Foul Play": {
		"id":             uint(492),
		"type":           "dark",
		"classification": uint(1),
	},
	"Hone Claws": {
		"id":             uint(468),
		"type":           "dark",
		"classification": uint(0),
	},
	"Hyperspace Fury": {
		"id":             uint(621),
		"type":           "dark",
		"classification": uint(1),
	},
	"Knock Off": {
		"id":             uint(282),
		"type":           "dark",
		"classification": uint(1),
	},
	"Memento": {
		"id":             uint(262),
		"type":           "dark",
		"classification": uint(0),
	},
	"Nasty Plot": {
		"id":             uint(417),
		"type":           "dark",
		"classification": uint(0),
	},
	"Night Daze": {
		"id":             uint(539),
		"type":           "dark",
		"classification": uint(2),
	},
	"Night Slash": {
		"id":             uint(400),
		"type":           "dark",
		"classification": uint(1),
	},
	"Parting Shot": {
		"id":             uint(575),
		"type":           "dark",
		"classification": uint(0),
	},
	"Payback": {
		"id":             uint(371),
		"type":           "dark",
		"classification": uint(1),
	},
	"Power Trip": {
		"id":             uint(681),
		"type":           "dark",
		"classification": uint(1),
	},
	"Punishment": {
		"id":             uint(386),
		"type":           "dark",
		"classification": uint(1),
	},
	"Pursuit": {
		"id":             uint(228),
		"type":           "dark",
		"classification": uint(1),
	},
	"Quash": {
		"id":             uint(511),
		"type":           "dark",
		"classification": uint(0),
	},
	"Snarl": {
		"id":             uint(555),
		"type":           "dark",
		"classification": uint(2),
	},
	"Snatch": {
		"id":             uint(289),
		"type":           "dark",
		"classification": uint(0),
	},
	"Sucker Punch": {
		"id":             uint(389),
		"type":           "dark",
		"classification": uint(1),
	},
	"Switcheroo": {
		"id":             uint(415),
		"type":           "dark",
		"classification": uint(3),
	},
	"Taunt": {
		"id":             uint(269),
		"type":           "dark",
		"classification": uint(0),
	},
	"Thief": {
		"id":             uint(168),
		"type":           "dark",
		"classification": uint(1),
	},
	"Throat Chop": {
		"id":             uint(675),
		"type":           "dark",
		"classification": uint(1),
	},
	"Topsy-Turvy": {
		"id":             uint(576),
		"type":           "dark",
		"classification": uint(0),
	},
	"Torment": {
		"id":             uint(259),
		"type":           "dark",
		"classification": uint(0),
	},
	"Aqua Jet": {
		"id":             uint(453),
		"type":           "water",
		"classification": uint(1),
	},
	"Aqua Ring": {
		"id":             uint(392),
		"type":           "water",
		"classification": uint(0),
	},
	"Aqua Tail": {
		"id":             uint(401),
		"type":           "water",
		"classification": uint(1),
	},
	"Brine": {
		"id":             uint(362),
		"type":           "water",
		"classification": uint(2),
	},
	"Bubble": {
		"id":             uint(145),
		"type":           "water",
		"classification": uint(2),
	},
	"Bubble Beam": {
		"id":             uint(61),
		"type":           "water",
		"classification": uint(2),
	},
	"Clamp": {
		"id":             uint(128),
		"type":           "water",
		"classification": uint(1),
	},
	"Crabhammer": {
		"id":             uint(152),
		"type":           "water",
		"classification": uint(1),
	},
	"Dive": {
		"id":             uint(291),
		"type":           "water",
		"classification": uint(1),
	},
	"Hydro Cannon": {
		"id":             uint(308),
		"type":           "water",
		"classification": uint(2),
	},
	"Hydro Pump": {
		"id":             uint(56),
		"type":           "water",
		"classification": uint(2),
	},
	"Liquidation": {
		"id":             uint(710),
		"type":           "water",
		"classification": uint(1),
	},
	"Muddy Water": {
		"id":             uint(330),
		"type":           "water",
		"classification": uint(2),
	},
	"Octazooka": {
		"id":             uint(190),
		"type":           "water",
		"classification": uint(2),
	},
	"Origin Pulse": {
		"id":             uint(618),
		"type":           "water",
		"classification": uint(2),
	},
	"Rain Dance": {
		"id":             uint(240),
		"type":           "water",
		"classification": uint(0),
	},
	"Razor Shell": {
		"id":             uint(534),
		"type":           "water",
		"classification": uint(1),
	},
	"Scald": {
		"id":             uint(503),
		"type":           "water",
		"classification": uint(2),
	},
	"Soak": {
		"id":             uint(487),
		"type":           "water",
		"classification": uint(0),
	},
	"Sparkling Aria": {
		"id":             uint(664),
		"type":           "water",
		"classification": uint(2),
	},
	"Steam Eruption": {
		"id":             uint(592),
		"type":           "water",
		"classification": uint(2),
	},
	"Surf": {
		"id":             uint(57),
		"type":           "water",
		"classification": uint(2),
	},
	"Water Gun": {
		"id":             uint(55),
		"type":           "water",
		"classification": uint(2),
	},
	"Water Pledge": {
		"id":             uint(518),
		"type":           "water",
		"classification": uint(2),
	},
	"Water Pulse": {
		"id":             uint(352),
		"type":           "water",
		"classification": uint(2),
	},
	"Water Shuriken": {
		"id":             uint(594),
		"type":           "water",
		"classification": uint(2),
	},
	"Water Sport": {
		"id":             uint(346),
		"type":           "water",
		"classification": uint(0),
	},
	"Water Spout": {
		"id":             uint(323),
		"type":           "water",
		"classification": uint(2),
	},
	"Waterfall": {
		"id":             uint(127),
		"type":           "water",
		"classification": uint(1),
	},
	"Whirlpool": {
		"id":             uint(250),
		"type":           "water",
		"classification": uint(2),
	},
	"Withdraw": {
		"id":             uint(110),
		"type":           "water",
		"classification": uint(0),
	},
	"Agility": {
		"id":             uint(97),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Ally Switch": {
		"id":             uint(502),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Amnesia": {
		"id":             uint(133),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Barrier": {
		"id":             uint(112),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Calm Mind": {
		"id":             uint(347),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Confusion": {
		"id":             uint(93),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Cosmic Power": {
		"id":             uint(322),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Dream Eater": {
		"id":             uint(138),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Extrasensory": {
		"id":             uint(326),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Future Sight": {
		"id":             uint(248),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Gravity": {
		"id":             uint(356),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Guard Split": {
		"id":             uint(470),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Guard Swap": {
		"id":             uint(385),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Heal Block": {
		"id":             uint(377),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Heal Pulse": {
		"id":             uint(505),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Healing Wish": {
		"id":             uint(361),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Heart Stamp": {
		"id":             uint(531),
		"type":           "psychic",
		"classification": uint(1),
	},
	"Heart Swap": {
		"id":             uint(391),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Hyperspace Hole": {
		"id":             uint(593),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Hypnosis": {
		"id":             uint(95),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Imprison": {
		"id":             uint(286),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Instruct": {
		"id":             uint(689),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Kinesis": {
		"id":             uint(134),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Light Screen": {
		"id":             uint(113),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Lunar Dance": {
		"id":             uint(461),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Luster Purge": {
		"id":             uint(295),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Magic Coat": {
		"id":             uint(277),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Magic Room": {
		"id":             uint(478),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Meditate": {
		"id":             uint(96),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Miracle Eye": {
		"id":             uint(357),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Mirror Coat": {
		"id":             uint(243),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Mist Ball": {
		"id":             uint(296),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Power Split": {
		"id":             uint(471),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Power Swap": {
		"id":             uint(384),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Power Trick": {
		"id":             uint(379),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Prismatic Laser": {
		"id":             uint(711),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psybeam": {
		"id":             uint(60),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psychic": {
		"id":             uint(94),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psychic Fangs": {
		"id":             uint(706),
		"type":           "psychic",
		"classification": uint(1),
	},
	"Psychic Terrain": {
		"id":             uint(678),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Psycho Boost": {
		"id":             uint(354),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psycho Cut": {
		"id":             uint(427),
		"type":           "psychic",
		"classification": uint(1),
	},
	"Psycho Shift": {
		"id":             uint(375),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Psyshock": {
		"id":             uint(473),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psystrike": {
		"id":             uint(540),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psywave": {
		"id":             uint(149),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Reflect": {
		"id":             uint(115),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Rest": {
		"id":             uint(156),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Role Play": {
		"id":             uint(272),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Skill Swap": {
		"id":             uint(285),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Speed Swap": {
		"id":             uint(683),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Stored Power": {
		"id":             uint(500),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Synchronoise": {
		"id":             uint(485),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Telekinesis": {
		"id":             uint(477),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Teleport": {
		"id":             uint(100),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Trick": {
		"id":             uint(271),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Trick Room": {
		"id":             uint(433),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Wonder Room": {
		"id":             uint(472),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Zen Headbutt": {
		"id":             uint(428),
		"type":           "psychic",
		"classification": uint(1),
	},
	"Clanging Scales": {
		"id":             uint(691),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Core Enforcer": {
		"id":             uint(687),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Draco Meteor": {
		"id":             uint(434),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Dragon Breath": {
		"id":             uint(225),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Dragon Claw": {
		"id":             uint(337),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Dragon Dance": {
		"id":             uint(349),
		"type":           "dragon",
		"classification": uint(0),
	},
	"Dragon Hammer": {
		"id":             uint(692),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Dragon Pulse": {
		"id":             uint(406),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Dragon Rage": {
		"id":             uint(82),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Dragon Rush": {
		"id":             uint(407),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Dragon Tail": {
		"id":             uint(525),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Dual Chop": {
		"id":             uint(530),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Outrage": {
		"id":             uint(200),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Roar of Time": {
		"id":             uint(459),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Spacial Rend": {
		"id":             uint(460),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Twister": {
		"id":             uint(239),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Accelerock": {
		"id":             uint(709),
		"type":           "rock",
		"classification": uint(1),
	},
	"Ancient Power": {
		"id":             uint(246),
		"type":           "rock",
		"classification": uint(2),
	},
	"Diamond Storm": {
		"id":             uint(591),
		"type":           "rock",
		"classification": uint(1),
	},
	"Head Smash": {
		"id":             uint(457),
		"type":           "rock",
		"classification": uint(1),
	},
	"Power Gem": {
		"id":             uint(408),
		"type":           "rock",
		"classification": uint(2),
	},
	"Rock Blast": {
		"id":             uint(350),
		"type":           "rock",
		"classification": uint(1),
	},
	"Rock Polish": {
		"id":             uint(397),
		"type":           "rock",
		"classification": uint(0),
	},
	"Rock Slide": {
		"id":             uint(157),
		"type":           "rock",
		"classification": uint(1),
	},
	"Rock Throw": {
		"id":             uint(88),
		"type":           "rock",
		"classification": uint(1),
	},
	"Rock Tomb": {
		"id":             uint(317),
		"type":           "rock",
		"classification": uint(1),
	},
	"Rock Wrecker": {
		"id":             uint(439),
		"type":           "rock",
		"classification": uint(1),
	},
	"Rollout": {
		"id":             uint(205),
		"type":           "rock",
		"classification": uint(1),
	},
	"Sandstorm": {
		"id":             uint(201),
		"type":           "rock",
		"classification": uint(0),
	},
	"Smack Down": {
		"id":             uint(479),
		"type":           "rock",
		"classification": uint(1),
	},
	"Stealth Rock": {
		"id":             uint(446),
		"type":           "rock",
		"classification": uint(0),
	},
	"Stone Edge": {
		"id":             uint(444),
		"type":           "rock",
		"classification": uint(1),
	},
	"Wide Guard": {
		"id":             uint(469),
		"type":           "rock",
		"classification": uint(0),
	},
	"Astonish": {
		"id":             uint(310),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Confuse Ray": {
		"id":             uint(109),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Curse": {
		"id":             uint(174),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Destiny Bond": {
		"id":             uint(194),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Grudge": {
		"id":             uint(288),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Hex": {
		"id":             uint(506),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Lick": {
		"id":             uint(122),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Moongeist Beam": {
		"id":             uint(714),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Night Shade": {
		"id":             uint(101),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Nightmare": {
		"id":             uint(171),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Ominous Wind": {
		"id":             uint(466),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Phantom Force": {
		"id":             uint(566),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Shadow Ball": {
		"id":             uint(247),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Shadow Bone": {
		"id":             uint(708),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Shadow Claw": {
		"id":             uint(421),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Shadow Force": {
		"id":             uint(467),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Shadow Punch": {
		"id":             uint(325),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Shadow Sneak": {
		"id":             uint(425),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Spirit Shackle": {
		"id":             uint(662),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Spite": {
		"id":             uint(180),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Trick-or-Treat": {
		"id":             uint(567),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Anchor Shot": {
		"id":             uint(677),
		"type":           "steel",
		"classification": uint(1),
	},
	"Autotomize": {
		"id":             uint(475),
		"type":           "steel",
		"classification": uint(0),
	},
	"Bullet Punch": {
		"id":             uint(418),
		"type":           "steel",
		"classification": uint(1),
	},
	"Doom Desire": {
		"id":             uint(353),
		"type":           "steel",
		"classification": uint(2),
	},
	"Flash Cannon": {
		"id":             uint(430),
		"type":           "steel",
		"classification": uint(2),
	},
	"Gear Grind": {
		"id":             uint(544),
		"type":           "steel",
		"classification": uint(1),
	},
	"Gear Up": {
		"id":             uint(674),
		"type":           "steel",
		"classification": uint(0),
	},
	"Gyro Ball": {
		"id":             uint(360),
		"type":           "steel",
		"classification": uint(1),
	},
	"Heavy Slam": {
		"id":             uint(484),
		"type":           "steel",
		"classification": uint(1),
	},
	"Iron Defense": {
		"id":             uint(334),
		"type":           "steel",
		"classification": uint(0),
	},
	"Iron Head": {
		"id":             uint(442),
		"type":           "steel",
		"classification": uint(1),
	},
	"Iron Tail": {
		"id":             uint(231),
		"type":           "steel",
		"classification": uint(1),
	},
	"King's Shield": {
		"id":             uint(588),
		"type":           "steel",
		"classification": uint(0),
	},
	"Magnet Bomb": {
		"id":             uint(443),
		"type":           "steel",
		"classification": uint(1),
	},
	"Metal Burst": {
		"id":             uint(368),
		"type":           "steel",
		"classification": uint(1),
	},
	"Metal Claw": {
		"id":             uint(232),
		"type":           "steel",
		"classification": uint(1),
	},
	"Metal Sound": {
		"id":             uint(319),
		"type":           "steel",
		"classification": uint(0),
	},
	"Meteor Mash": {
		"id":             uint(309),
		"type":           "steel",
		"classification": uint(1),
	},
	"Mirror Shot": {
		"id":             uint(429),
		"type":           "steel",
		"classification": uint(2),
	},
	"Shift Gear": {
		"id":             uint(508),
		"type":           "steel",
		"classification": uint(0),
	},
	"Smart Strike": {
		"id":             uint(684),
		"type":           "steel",
		"classification": uint(1),
	},
	"Steel Wing": {
		"id":             uint(211),
		"type":           "steel",
		"classification": uint(1),
	},
	"Sunsteel Strike": {
		"id":             uint(713),
		"type":           "steel",
		"classification": uint(1),
	},
	"Aromatic Mist": {
		"id":             uint(597),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Baby-Doll Eyes": {
		"id":             uint(608),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Charm": {
		"id":             uint(204),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Crafty Shield": {
		"id":             uint(578),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Dazzling Gleam": {
		"id":             uint(605),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Disarming Voice": {
		"id":             uint(574),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Draining Kiss": {
		"id":             uint(577),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Fairy Lock": {
		"id":             uint(587),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Fairy Wind": {
		"id":             uint(584),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Fleur Cannon": {
		"id":             uint(705),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Floral Healing": {
		"id":             uint(666),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Flower Shield": {
		"id":             uint(579),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Geomancy": {
		"id":             uint(601),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Misty Terrain": {
		"id":             uint(581),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Moonblast": {
		"id":             uint(585),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Moonlight": {
		"id":             uint(236),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Nature's Madness": {
		"id":             uint(717),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Play Rough": {
		"id":             uint(583),
		"type":           "fairy",
		"classification": uint(1),
	},
	"Sweet Kiss": {
		"id":             uint(186),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Dynamax Cannon": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(2),
	},
	"Snipe Shot": {
		"id":             0,
		"type":           "water",
		"classification": uint(2),
	},
	"Jaw Lock": {
		"id":             0,
		"type":           "dark",
		"classification": uint(1),
	},
	"Stuff Cheeks": {
		"id":             0,
		"type":           "normal",
		"classification": uint(0),
	},
	"No Retreat": {
		"id":             0,
		"type":           "fighting",
		"classification": uint(0),
	},
	"Tar Shot": {
		"id":             0,
		"type":           "rock",
		"classification": uint(0),
	},
	"Magic Powder": {
		"id":             0,
		"type":           "psychic",
		"classification": uint(0),
	},
	"Dragon Darts": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(1),
	},
	"Teatime": {
		"id":             0,
		"type":           "normal",
		"classification": uint(0),
	},
	"Octolock": {
		"id":             0,
		"type":           "fighting",
		"classification": uint(0),
	},
	"Bolt Beak": {
		"id":             0,
		"type":           "electric",
		"classification": uint(1),
	},
	"Fishious Rend": {
		"id":             0,
		"type":           "water",
		"classification": uint(1),
	},
	"Court Change": {
		"id":             0,
		"type":           "normal",
		"classification": uint(0),
	},
	"Clangorous Soul": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(0),
	},
	"Body Press": {
		"id":             0,
		"type":           "fighting",
		"classification": uint(1),
	},
	"Decorate": {
		"id":             0,
		"type":           "fairy",
		"classification": uint(0),
	},
	"Drum Beating": {
		"id":             0,
		"type":           "grass",
		"classification": uint(1),
	},
	"Snap Trap": {
		"id":             0,
		"type":           "grass",
		"classification": uint(1),
	},
	"Pyro Ball": {
		"id":             0,
		"type":           "fire",
		"classification": uint(1),
	},
	"Behemoth Blade": {
		"id":             0,
		"type":           "steel",
		"classification": uint(1),
	},
	"Behemoth Bash": {
		"id":             0,
		"type":           "steel",
		"classification": uint(1),
	},
	"Aura Wheel": {
		"id":             0,
		"type":           "electric",
		"classification": uint(1),
	},
	"Breaking Swipe": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(1),
	},
	"Branch Poke": {
		"id":             0,
		"type":           "grass",
		"classification": uint(1),
	},
	"Overdrive": {
		"id":             0,
		"type":           "electric",
		"classification": uint(2),
	},
	"Apple Acid": {
		"id":             0,
		"type":           "grass",
		"classification": uint(2),
	},
	"Grav Apple": {
		"id":             0,
		"type":           "grass",
		"classification": uint(1),
	},
	"Spirit Break": {
		"id":             0,
		"type":           "fairy",
		"classification": uint(1),
	},
	"Strange Steam": {
		"id":             0,
		"type":           "fairy",
		"classification": uint(2),
	},
	"Life Dew": {
		"id":             0,
		"type":           "water",
		"classification": uint(0),
	},
	"Obstruct": {
		"id":             0,
		"type":           "dark",
		"classification": uint(0),
	},
	"False Surrender": {
		"id":             0,
		"type":           "dark",
		"classification": uint(1),
	},
	"Meteor Assault": {
		"id":             0,
		"type":           "fighting",
		"classification": uint(1),
	},
	"Eternabeam": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(2),
	},
	"Steel Beam": {
		"id":             0,
		"type":           "steel",
		"classification": uint(2),
	},

	"Expanding Force": {
		"id":             0,
		"type":           "psychic",
		"classification": uint(2),
	},
	"Steel Roller": {
		"id":             0,
		"type":           "steel",
		"classification": uint(1),
	},
	"Scale Shot": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(1),
	},
	"Meteor Beam": {
		"id":             0,
		"type":           "rock",
		"classification": uint(2),
	},
	"Shell Side Arm": {
		"id":             0,
		"type":           "poison",
		"classification": uint(2),
	},
	"Misty Explosion": {
		"id":             0,
		"type":           "fairy",
		"classification": uint(2),
	},
	"Grassy Glide": {
		"id":             0,
		"type":           "grass",
		"classification": uint(1),
	},
	"Rising Voltage": {
		"id":             0,
		"type":           "electric",
		"classification": uint(2),
	},
	"Terrain Pulse": {
		"id":             0,
		"type":           "normal",
		"classification": uint(2),
	},
	"Skitter Smack": {
		"id":             0,
		"type":           "bug",
		"classification": uint(1),
	},
	"Burning Jealousy": {
		"id":             0,
		"type":           "fire",
		"classification": uint(2),
	},
	"Lash Out": {
		"id":             0,
		"type":           "dark",
		"classification": uint(1),
	},
	"Poltergeist": {
		"id":             0,
		"type":           "ghost",
		"classification": uint(1),
	},
	"Corrosive Gas": {
		"id":             0,
		"type":           "poison",
		"classification": uint(0),
	},
	"Coaching": {
		"id":             0,
		"type":           "fighting",
		"classification": uint(0),
	},
	"Flip Turn": {
		"id":             0,
		"type":           "water",
		"classification": uint(1),
	},
	"Triple Axel": {
		"id":             0,
		"type":           "ice",
		"classification": uint(1),
	},
	"Dual Wingbeat": {
		"id":             0,
		"type":           "flying",
		"classification": uint(1),
	},
	"Scorching Sands": {
		"id":             0,
		"type":           "ground",
		"classification": uint(2),
	},
	"Jungle Healing": {
		"id":             0,
		"type":           "grass",
		"classification": uint(0),
	},
	"Wicked Blow": {
		"id":             0,
		"type":           "dark",
		"classification": uint(1),
	},
	"Surging Strikes": {
		"id":             0,
		"type":           "water",
		"classification": uint(1),
	},
	"Thunder Cage": {
		"id":             0,
		"type":           "electric",
		"classification": uint(2),
	},
	"Dragon Energy": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(2),
	},
	"Freezing Glare": {
		"id":             0,
		"type":           "psychic",
		"classification": uint(2),
	},
	"Fiery Wrath": {
		"id":             0,
		"type":           "dark",
		"classification": uint(2),
	},
	"Thunderous Kick": {
		"id":             0,
		"type":           "fighting",
		"classification": uint(1),
	},
	"Glacial Lance": {
		"id":             0,
		"type":           "ice",
		"classification": uint(1),
	},
	"Astral Barrage": {
		"id":             0,
		"type":           "ghost",
		"classification": uint(2),
	},
	"Eerie Spell": {
		"id":             0,
		"type":           "psychic",
		"classification": uint(2),
	},
}

var itemData = map[string]map[string]interface{}{
	"Ability Capsule": {
		"id": uint(645),
	},
	"Abomasite": {
		"id": uint(674),
	},
	"Absolite": {
		"id": uint(677),
	},
	"Absorb Bulb": {
		"id": uint(545),
	},
	"Adamant Orb": {
		"id": uint(135),
	},
	"Adrenaline Orb": {
		"id": uint(846),
	},
	"Aerodactylite": {
		"id": uint(672),
	},
	"Aggronite": {
		"id": uint(667),
	},
	"Aguav Berry": {
		"id": uint(162),
	},
	"Air Balloon": {
		"id": uint(541),
	},
	"Alakazite": {
		"id": uint(679),
	},
	"Aloraichium Z": {
		"id":   uint(803),
		"type": "electric",
	},
	"Altarianite": {
		"id": uint(755),
	},
	"Ampharosite": {
		"id": uint(658),
	},
	"Amulet Coin": {
		"id": uint(223),
	},
	"Antidote": {
		"id": uint(18),
	},
	"Apicot Berry": {
		"id": uint(205),
	},
	"Armor Fossil": {
		"id": uint(104),
	},
	"Aspear Berry": {
		"id": uint(153),
	},
	"Assault Vest": {
		"id": uint(640),
	},
	"Audinite": {
		"id": uint(757),
	},
	"Awakening": {
		"id": uint(21),
	},
	"Babiri Berry": {
		"id":   uint(199),
		"type": "steel",
	},
	"Balm Mushroom": {
		"id": uint(580),
	},
	"Banettite": {
		"id": uint(668),
	},
	"Beast Ball": {
		"id": uint(851),
	},
	"Beedrillite": {
		"id": uint(770),
	},
	"Berry Juice": {
		"id": uint(43),
	},
	"Big Malasada": {
		"id": uint(852),
	},
	"Big Mushroom": {
		"id": uint(87),
	},
	"Big Nugget": {
		"id": uint(581),
	},
	"Big Pearl": {
		"id": uint(89),
	},
	"Big Root": {
		"id": uint(296),
	},
	"Binding Band": {
		"id": uint(544),
	},
	"Black Belt": {
		"id":   uint(241),
		"type": "fighting",
	},
	"Black Glasses": {
		"id":   uint(240),
		"type": "dark",
	},
	"Black Sludge": {
		"id": uint(281),
	},
	"Blastoisinite": {
		"id": uint(661),
	},
	"Blazikenite": {
		"id": uint(664),
	},
	"Blue Orb": {
		"id": uint(535),
	},
	"Blue Shard": {
		"id": uint(73),
	},
	"Bluk Berry": {
		"id": uint(165),
	},
	"Bottle Cap": {
		"id": uint(795),
	},
	"Bright Powder": {
		"id": uint(213),
	},
	"Bug Gem": {
		"id":   uint(558),
		"type": "bug",
	},
	"Bug Memory": {
		"id":   uint(909),
		"type": "bug",
	},
	"Buginium Z": {
		"id":   uint(787),
		"type": "bug",
	},
	"Burn Drive": {
		"id":   uint(118),
		"type": "fire",
	},
	"Burn Heal": {
		"id": uint(19),
	},
	"Calcium": {
		"id": uint(49),
	},
	"Cameruptite": {
		"id": uint(767),
	},
	"Carbos": {
		"id": uint(48),
	},
	"Casteliacone": {
		"id": uint(591),
	},
	"Cell Battery": {
		"id": uint(546),
	},
	"Charcoal": {
		"id":   uint(249),
		"type": "fire",
	},
	"Charizardite X": {
		"id": uint(660),
	},
	"Charizardite Y": {
		"id": uint(678),
	},
	"Charti Berry": {
		"id":   uint(195),
		"type": "rock",
	},
	"Cheri Berry": {
		"id": uint(149),
	},
	"Chesto Berry": {
		"id": uint(150),
	},
	"Chilan Berry": {
		"id":   uint(200),
		"type": "normal",
	},
	"Chill Drive": {
		"id":   uint(119),
		"type": "ice",
	},
	"Choice Band": {
		"id": uint(220),
	},
	"Choice Scarf": {
		"id": uint(287),
	},
	"Choice Specs": {
		"id": uint(297),
	},
	"Chople Berry": {
		"id":   uint(189),
		"type": "fighting",
	},
	"Cleanse Tag": {
		"id": uint(224),
	},
	"Clever Wing": {
		"id": uint(569),
	},
	"Coba Berry": {
		"id":   uint(192),
		"type": "flying",
	},
	"Colbur Berry": {
		"id":   uint(198),
		"type": "dark",
	},
	"Comet Shard": {
		"id": uint(583),
	},
	"Cover Fossil": {
		"id": uint(572),
	},
	"Damp Rock": {
		"id":   uint(285),
		"type": "water",
	},
	"Dark Gem": {
		"id":   uint(562),
		"type": "dark",
	},
	"Dark Memory": {
		"id":   uint(919),
		"type": "dark",
	},
	"Darkinium Z": {
		"id":   uint(791),
		"type": "dark",
	},
	"Dawn Stone": {
		"id": uint(109),
	},
	"Decidium Z": {
		"id":   uint(798),
		"type": "ghost",
	},
	"Deep Sea Scale": {
		"id": uint(227),
	},
	"Deep Sea Tooth": {
		"id": uint(226),
	},
	"Destiny Knot": {
		"id": uint(280),
	},
	"Diancite": {
		"id": uint(764),
	},
	"Dire Hit": {
		"id": uint(56),
	},
	"Dive Ball": {
		"id": uint(7),
	},
	"Douse Drive": {
		"id":   uint(116),
		"type": "water",
	},
	"Draco Plate": {
		"id":   uint(311),
		"type": "dragon",
	},
	"Dragon Fang": {
		"id":   uint(250),
		"type": "dragon",
	},
	"Dragon Gem": {
		"id":   uint(561),
		"type": "dragon",
	},
	"Dragon Memory": {
		"id":   uint(918),
		"type": "dragon",
	},
	"Dragon Scale": {
		"id": uint(235),
	},
	"Dragonium Z": {
		"id":   uint(790),
		"type": "dragon",
	},
	"Dread Plate": {
		"id":   uint(312),
		"type": "dark",
	},
	"Dubious Disc": {
		"id": uint(324),
	},
	"Dusk Ball": {
		"id": uint(13),
	},
	"Dusk Stone": {
		"id": uint(108),
	},
	"Earth Plate": {
		"id":   uint(305),
		"type": "ground",
	},
	"Eevium Z": {
		"id":   uint(805),
		"type": "normal",
	},
	"Eject Button": {
		"id": uint(547),
	},
	"Electirizer": {
		"id": uint(322),
	},
	"Electric Gem": {
		"id":   uint(550),
		"type": "electric",
	},
	"Electric Memory": {
		"id":   uint(915),
		"type": "electric",
	},
	"Electric Seed": {
		"id": uint(881),
	},
	"Electrium Z": {
		"id":   uint(779),
		"type": "electric",
	},
	"Elixir": {
		"id": uint(40),
	},
	"Energy Powder": {
		"id": uint(34),
	},
	"Energy Root": {
		"id": uint(35),
	},
	"Escape Rope": {
		"id": uint(78),
	},
	"Ether": {
		"id": uint(38),
	},
	"Everstone": {
		"id": uint(229),
	},
	"Eviolite": {
		"id": uint(538),
	},
	"Expert Belt": {
		"id": uint(268),
	},
	"Fairium Z": {
		"id":   uint(793),
		"type": "fairy",
	},
	"Fairy Gem": {
		"id":   uint(715),
		"type": "fairy",
	},
	"Fairy Memory": {
		"id":   uint(920),
		"type": "fairy",
	},
	"Fast Ball": {
		"id": uint(492),
	},
	"Festival Ticket": {
		"id": uint(844),
	},
	"Fighting Gem": {
		"id":   uint(553),
		"type": "fighting",
	},
	"Fighting Memory": {
		"id": uint(904),
	},
	"Fightinium Z": {
		"id":   uint(782),
		"type": "fighting",
	},
	"Figy Berry": {
		"id": uint(159),
	},
	"Fire Gem": {
		"id":   uint(548),
		"type": "fire",
	},
	"Fire Memory": {
		"id":   uint(912),
		"type": "fire",
	},
	"Fire Stone": {
		"id": uint(82),
	},
	"Firium Z": {
		"id":   uint(777),
		"type": "fire",
	},
	"Fist Plate": {
		"id":   uint(303),
		"type": "fighting",
	},
	"Flame Orb": {
		"id": uint(273),
	},
	"Flame Plate": {
		"id":   uint(298),
		"type": "fire",
	},
	"Float Stone": {
		"id": uint(539),
	},
	"Flying Gem": {
		"id":   uint(556),
		"type": "flying",
	},
	"Flying Memory": {
		"id":   uint(905),
		"type": "flying",
	},
	"Flyinium Z": {
		"id":   uint(785),
		"type": "flying",
	},
	"Focus Band": {
		"id": uint(230),
	},
	"Focus Sash": {
		"id": uint(275),
	},
	"Fresh Water": {
		"id": uint(30),
	},
	"Friend Ball": {
		"id": uint(497),
	},
	"Full Heal": {
		"id": uint(27),
	},
	"Full Incense": {
		"id": uint(316),
	},
	"Full Restore": {
		"id": uint(23),
	},
	"Galladite": {
		"id": uint(756),
	},
	"Ganlon Berry": {
		"id": uint(202),
	},
	"Garchompite": {
		"id": uint(683),
	},
	"Gardevoirite": {
		"id": uint(657),
	},
	"Gengarite": {
		"id": uint(656),
	},
	"Genius Wing": {
		"id": uint(568),
	},
	"Ghost Gem": {
		"id":   uint(560),
		"type": "ghost",
	},
	"Ghost Memory": {
		"id":   uint(910),
		"type": "ghost",
	},
	"Ghostium Z": {
		"id":   uint(789),
		"type": "ghost",
	},
	"Glalitite": {
		"id": uint(763),
	},
	"Gold Bottle Cap": {
		"id": uint(796),
	},
	"Grass Gem": {
		"id":   uint(551),
		"type": "grass",
	},
	"Grass Memory": {
		"id":   uint(914),
		"type": "grass",
	},
	"Grassium Z": {
		"id":   uint(780),
		"type": "grass",
	},
	"Grassy Seed": {
		"id": uint(884),
	},
	"Great Ball": {
		"id": uint(3),
	},
	"Green Shard": {
		"id": uint(75),
	},
	"Grepa Berry": {
		"id": uint(173),
	},
	"Grip Claw": {
		"id": uint(286),
	},
	"Griseous Orb": {
		"id": uint(112),
	},
	"Ground Gem": {
		"id":   uint(555),
		"type": "ground",
	},
	"Ground Memory": {
		"id":   uint(907),
		"type": "ground",
	},
	"Groundium Z": {
		"id":   uint(784),
		"type": "ground",
	},
	"Guard Spec.": {
		"id": uint(55),
	},
	"Gyaradosite": {
		"id": uint(676),
	},
	"HP Up": {
		"id": uint(45),
	},
	"Haban Berry": {
		"id":   uint(197),
		"type": "dragon",
	},
	"Hard Stone": {
		"id":   uint(238),
		"type": "rock",
	},
	"Heal Ball": {
		"id": uint(14),
	},
	"Heal Powder": {
		"id": uint(36),
	},
	"Health Wing": {
		"id": uint(565),
	},
	"Heart Scale": {
		"id": uint(93),
	},
	"Heat Rock": {
		"id":   uint(284),
		"type": "fire",
	},
	"Heavy Ball": {
		"id": uint(495),
	},
	"Heracronite": {
		"id": uint(680),
	},
	"Hondew Berry": {
		"id": uint(172),
	},
	"Honey": {
		"id": uint(94),
	},
	"Houndoominite": {
		"id": uint(666),
	},
	"Hyper Potion": {
		"id": uint(25),
	},
	"Iapapa Berry": {
		"id": uint(163),
	},
	"Ice Gem": {
		"id":   uint(552),
		"type": "ice",
	},
	"Ice Heal": {
		"id": uint(20),
	},
	"Ice Memory": {
		"id":   uint(917),
		"type": "ice",
	},
	"Ice Stone": {
		"id": uint(849),
	},
	"Icicle Plate": {
		"id":   uint(302),
		"type": "ice",
	},
	"Icium Z": {
		"id":   uint(781),
		"type": "ice",
	},
	"Icy Rock": {
		"id":   uint(282),
		"type": "ice",
	},
	"Incinium Z": {
		"id":   uint(799),
		"type": "dark",
	},
	"Insect Plate": {
		"id":   uint(308),
		"type": "bug",
	},
	"Iron": {
		"id": uint(47),
	},
	"Iron Ball": {
		"id": uint(278),
	},
	"Iron Plate": {
		"id":   uint(313),
		"type": "steel",
	},
	"Kangaskhanite": {
		"id": uint(675),
	},
	"Kasib Berry": {
		"id":   uint(196),
		"type": "ghost",
	},
	"Kebia Berry": {
		"id":   uint(190),
		"type": "poison",
	},
	"Kee Berry": {
		"id": uint(687),
	},
	"Kelpsy Berry": {
		"id": uint(170),
	},
	"King's Rock": {
		"id": uint(221),
	},
	"Kommonium Z": {
		"id":   uint(926),
		"type": "dragon",
	},
	"Lagging Tail": {
		"id": uint(279),
	},
	"Lansat Berry": {
		"id": uint(206),
	},
	"Latiasite": {
		"id": uint(684),
	},
	"Latiosite": {
		"id": uint(685),
	},
	"Lava Cookie": {
		"id": uint(42),
	},
	"Lax Incense": {
		"id": uint(255),
	},
	"Leaf Stone": {
		"id": uint(85),
	},
	"Leftovers": {
		"id": uint(234),
	},
	"Lemonade": {
		"id": uint(32),
	},
	"Leppa Berry": {
		"id": uint(154),
	},
	"Level Ball": {
		"id": uint(493),
	},
	"Liechi Berry": {
		"id": uint(201),
	},
	"Life Orb": {
		"id": uint(270),
	},
	"Light Ball": {
		"id": uint(236),
	},
	"Light Clay": {
		"id": uint(269),
	},
	"Lopunnite": {
		"id": uint(768),
	},
	"Love Ball": {
		"id": uint(496),
	},
	"Lucarionite": {
		"id": uint(673),
	},
	"Luck Incense": {
		"id": uint(319),
	},
	"Lucky Egg": {
		"id": uint(231),
	},
	"Lucky Punch": {
		"id": uint(256),
	},
	"Lum Berry": {
		"id": uint(157),
	},
	"Luminous Moss": {
		"id": uint(648),
	},
	"Lumiose Galette": {
		"id": uint(708),
	},
	"Lunalium Z": {
		"id":   uint(922),
		"type": "ghost",
	},
	"Lure Ball": {
		"id": uint(494),
	},
	"Lustrous Orb": {
		"id": uint(136),
	},
	"Luxury Ball": {
		"id": uint(11),
	},
	"Lycanium Z": {
		"id":   uint(925),
		"type": "rock",
	},
	"Magmarizer": {
		"id": uint(323),
	},
	"Magnet": {
		"id":   uint(242),
		"type": "electric",
	},
	"Mago Berry": {
		"id": uint(161),
	},
	"Manectite": {
		"id": uint(682),
	},
	"Maranga Berry": {
		"id": uint(688),
	},
	"Marshadium Z": {
		"id":   uint(802),
		"type": "ghost",
	},
	"Master Ball": {
		"id": uint(1),
	},
	"Mawilite": {
		"id": uint(681),
	},
	"Max Elixir": {
		"id": uint(41),
	},
	"Max Ether": {
		"id": uint(39),
	},
	"Max Potion": {
		"id": uint(24),
	},
	"Max Repel": {
		"id": uint(77),
	},
	"Max Revive": {
		"id": uint(29),
	},
	"Meadow Plate": {
		"id":   uint(301),
		"type": "grass",
	},
	"Medichamite": {
		"id": uint(665),
	},
	"Mental Herb": {
		"id": uint(219),
	},
	"Metagrossite": {
		"id": uint(758),
	},
	"Metal Coat": {
		"id":   uint(233),
		"type": "steel",
	},
	"Metal Powder": {
		"id": uint(257),
	},
	"Metronome": {
		"id": uint(277),
	},
	"Mewnium Z": {
		"id":   uint(806),
		"type": "psychic",
	},
	"Mewtwonite X": {
		"id": uint(662),
	},
	"Mewtwonite Y": {
		"id": uint(663),
	},
	"Mimikium Z": {
		"id":   uint(924),
		"type": "fairy",
	},
	"Mind Plate": {
		"id":   uint(307),
		"type": "psychic",
	},
	"Miracle Seed": {
		"id":   uint(239),
		"type": "grass",
	},
	"Misty Seed": {
		"id": uint(883),
	},
	"Moomoo Milk": {
		"id": uint(33),
	},
	"Moon Ball": {
		"id": uint(498),
	},
	"Moon Stone": {
		"id": uint(81),
	},
	"Muscle Band": {
		"id": uint(266),
	},
	"Muscle Wing": {
		"id": uint(566),
	},
	"Mystic Water": {
		"id":   uint(243),
		"type": "water",
	},
	"Nest Ball": {
		"id": uint(8),
	},
	"Net Ball": {
		"id": uint(6),
	},
	"Never-Melt Ice": {
		"id":   uint(246),
		"type": "ice",
	},
	"Normal Gem": {
		"id":   uint(564),
		"type": "normal",
	},
	"Normalium Z": {
		"id":   uint(776),
		"type": "normal",
	},
	"Nugget": {
		"id": uint(92),
	},
	"Occa Berry": {
		"id":   uint(184),
		"type": "fire",
	},
	"Odd Incense": {
		"id":   uint(314),
		"type": "psychic",
	},
	"Old Gateau": {
		"id": uint(54),
	},
	"Oran Berry": {
		"id": uint(155),
	},
	"Oval Stone": {
		"id": uint(110),
	},
	"PP Max": {
		"id": uint(53),
	},
	"PP Up": {
		"id": uint(51),
	},
	"Paralyze Heal": {
		"id": uint(22),
	},
	"Passho Berry": {
		"id":   uint(185),
		"type": "water",
	},
	"Payapa Berry": {
		"id":   uint(193),
		"type": "psychic",
	},
	"Pearl": {
		"id": uint(88),
	},
	"Pearl String": {
		"id": uint(582),
	},
	"Pecha Berry": {
		"id": uint(151),
	},
	"Persim Berry": {
		"id": uint(156),
	},
	"Petaya Berry": {
		"id": uint(204),
	},
	"Pidgeotite": {
		"id": uint(762),
	},
	"Pikanium Z": {
		"id":   uint(794),
		"type": "electric",
	},
	"Pikashunium Z": {
		"id":   uint(835),
		"type": "electric",
	},
	"Pinap Berry": {
		"id": uint(168),
	},
	"Pink Nectar": {
		"id": uint(855),
	},
	"Pinsirite": {
		"id": uint(671),
	},
	"Pixie Plate": {
		"id":   uint(644),
		"type": "fairy",
	},
	"Plume Fossil": {
		"id": uint(573),
	},
	"Poison Barb": {
		"id":   uint(245),
		"type": "poison",
	},
	"Poison Gem": {
		"id":   uint(554),
		"type": "poison",
	},
	"Poison Memory": {
		"id":   uint(906),
		"type": "poison",
	},
	"Poisonium Z": {
		"id":   uint(783),
		"type": "poison",
	},
	"Pok\u00e9 Ball": {
		"id": uint(4),
	},
	"Pok\u00e9 Doll": {
		"id": uint(63),
	},
	"Pok\u00e9 Toy": {
		"id": uint(577),
	},
	"Pomeg Berry": {
		"id": uint(169),
	},
	"Potion": {
		"id": uint(17),
	},
	"Power Anklet": {
		"id": uint(293),
	},
	"Power Band": {
		"id": uint(292),
	},
	"Power Belt": {
		"id": uint(290),
	},
	"Power Bracer": {
		"id": uint(289),
	},
	"Power Herb": {
		"id": uint(271),
	},
	"Power Lens": {
		"id": uint(291),
	},
	"Power Weight": {
		"id": uint(294),
	},
	"Premier Ball": {
		"id": uint(12),
	},
	"Pretty Wing": {
		"id": uint(571),
	},
	"Primarium Z": {
		"id":   uint(800),
		"type": "water",
	},
	"Prism Scale": {
		"id": uint(537),
	},
	"Protective Pads": {
		"id": uint(880),
	},
	"Protector": {
		"id": uint(321),
	},
	"Protein": {
		"id": uint(46),
	},
	"Psychic Gem": {
		"id":   uint(557),
		"type": "psychic",
	},
	"Psychic Memory": {
		"id":   uint(916),
		"type": "psychic",
	},
	"Psychic Seed": {
		"id": uint(882),
	},
	"Psychium Z": {
		"id":   uint(786),
		"type": "psychic",
	},
	"Pure Incense": {
		"id": uint(320),
	},
	"Purple Nectar": {
		"id": uint(856),
	},
	"Qualot Berry": {
		"id": uint(171),
	},
	"Quick Ball": {
		"id": uint(15),
	},
	"Quick Claw": {
		"id": uint(217),
	},
	"Quick Powder": {
		"id": uint(274),
	},
	"Rage Candy Bar": {
		"id": uint(504),
	},
	"Rare Bone": {
		"id": uint(106),
	},
	"Rare Candy": {
		"id": uint(50),
	},
	"Rawst Berry": {
		"id": uint(152),
	},
	"Razor Claw": {
		"id": uint(326),
	},
	"Razor Fang": {
		"id": uint(327),
	},
	"Reaper Cloth": {
		"id": uint(325),
	},
	"Red Card": {
		"id": uint(542),
	},
	"Red Nectar": {
		"id": uint(853),
	},
	"Red Orb": {
		"id": uint(534),
	},
	"Red Shard": {
		"id": uint(72),
	},
	"Repeat Ball": {
		"id": uint(9),
	},
	"Repel": {
		"id": uint(79),
	},
	"Resist Wing": {
		"id": uint(567),
	},
	"Revival Herb": {
		"id": uint(37),
	},
	"Revive": {
		"id": uint(28),
	},
	"Rindo Berry": {
		"id":   uint(187),
		"type": "grass",
	},
	"Ring Target": {
		"id": uint(543),
	},
	"Rock Gem": {
		"id":   uint(559),
		"type": "rock",
	},
	"Rock Incense": {
		"id":   uint(315),
		"type": "rock",
	},
	"Rock Memory": {
		"id":   uint(908),
		"type": "rock",
	},
	"Rockium Z": {
		"id":   uint(788),
		"type": "rock",
	},
	"Rocky Helmet": {
		"id": uint(540),
	},
	"Rose Incense": {
		"id":   uint(318),
		"type": "grass",
	},
	"Roseli Berry": {
		"id":   uint(686),
		"type": "fairy",
	},
	"Sablenite": {
		"id": uint(754),
	},
	"Sachet": {
		"id": uint(647),
	},
	"Sacred Ash": {
		"id": uint(44),
	},
	"Safety Goggles": {
		"id": uint(650),
	},
	"Salac Berry": {
		"id": uint(203),
	},
	"Salamencite": {
		"id": uint(769),
	},
	"Sceptilite": {
		"id": uint(753),
	},
	"Scizorite": {
		"id": uint(670),
	},
	"Scope Lens": {
		"id": uint(232),
	},
	"Sea Incense": {
		"id":   uint(254),
		"type": "water",
	},
	"Shalour Sable": {
		"id": uint(709),
	},
	"Sharp Beak": {
		"id":   uint(244),
		"type": "flying",
	},
	"Sharpedonite": {
		"id": uint(759),
	},
	"Shed Shell": {
		"id": uint(295),
	},
	"Shell Bell": {
		"id": uint(253),
	},
	"Shiny Stone": {
		"id": uint(107),
	},
	"Shock Drive": {
		"id":   uint(117),
		"type": "electric",
	},
	"Shuca Berry": {
		"id":   uint(191),
		"type": "ground",
	},
	"Silk Scarf": {
		"id":   uint(251),
		"type": "normal",
	},
	"SilverPowder": {
		"id":   uint(222),
		"type": "bug",
	},
	"Silver Powder": {
		"id":   uint(222),
		"type": "bug",
	},
	"Sitrus Berry": {
		"id": uint(158),
	},
	"Skull Fossil": {
		"id": uint(105),
	},
	"Sky Plate": {
		"id":   uint(306),
		"type": "flying",
	},
	"Slowbronite": {
		"id": uint(760),
	},
	"Smoke Ball": {
		"id": uint(228),
	},
	"Smooth Rock": {
		"id":   uint(283),
		"type": "rock",
	},
	"Snorlium Z": {
		"id":   uint(804),
		"type": "normal",
	},
	"Snowball": {
		"id": uint(649),
	},
	"Soda Pop": {
		"id": uint(31),
	},
	"Soft Sand": {
		"id":   uint(237),
		"type": "ground",
	},
	"Solganium Z": {
		"id":   uint(921),
		"type": "steel",
	},
	"Soothe Bell": {
		"id": uint(218),
	},
	"Soul Dew": {
		"id": uint(225),
	},
	"Spell Tag": {
		"id":   uint(247),
		"type": "ghost",
	},
	"Splash Plate": {
		"id":   uint(299),
		"type": "water",
	},
	"Spooky Plate": {
		"id":   uint(310),
		"type": "ghost",
	},
	"Star Piece": {
		"id": uint(91),
	},
	"Stardust": {
		"id": uint(90),
	},
	"Starf Berry": {
		"id": uint(207),
	},
	"Steel Gem": {
		"id":   uint(563),
		"type": "steel",
	},
	"Steel Memory": {
		"id":   uint(911),
		"type": "steel",
	},
	"Steelium Z": {
		"id":   uint(792),
		"type": "steel",
	},
	"Steelixite": {
		"id": uint(761),
	},
	"Stick": {
		"id": uint(259),
	},
	"Sticky Barb": {
		"id": uint(288),
	},
	"Stone Plate": {
		"id":   uint(309),
		"type": "rock",
	},
	"Strange Souvenir": {
		"id": uint(704),
	},
	"Sun Stone": {
		"id": uint(80),
	},
	"Super Potion": {
		"id": uint(26),
	},
	"Super Repel": {
		"id": uint(76),
	},
	"Swampertite": {
		"id": uint(752),
	},
	"Sweet Heart": {
		"id": uint(134),
	},
	"Swift Wing": {
		"id": uint(570),
	},
	"Tamato Berry": {
		"id": uint(174),
	},
	"Tanga Berry": {
		"id":   uint(194),
		"type": "bug",
	},
	"Tapunium Z": {
		"id":   uint(801),
		"type": "fairy",
	},
	"Terrain Extender": {
		"id": uint(879),
	},
	"Thick Club": {
		"id": uint(258),
	},
	"Thunder Stone": {
		"id": uint(83),
	},
	"Timer Ball": {
		"id": uint(10),
	},
	"Tiny Mushroom": {
		"id": uint(86),
	},
	"Toxic Orb": {
		"id": uint(272),
	},
	"Toxic Plate": {
		"id":   uint(304),
		"type": "poison",
	},
	"Twisted Spoon": {
		"id":   uint(248),
		"type": "psychic",
	},
	"Tyranitarite": {
		"id": uint(669),
	},
	"Ultra Ball": {
		"id": uint(2),
	},
	"Ultranecrozium Z": {
		"id":   uint(923),
		"type": "psychic",
	},
	"Up-Grade": {
		"id": uint(252),
	},
	"Venusaurite": {
		"id": uint(659),
	},
	"Wacan Berry": {
		"id":   uint(186),
		"type": "electric",
	},
	"Water Gem": {
		"id":   uint(549),
		"type": "water",
	},
	"Water Memory": {
		"id":   uint(913),
		"type": "water",
	},
	"Water Stone": {
		"id": uint(84),
	},
	"Waterium Z": {
		"id":   uint(778),
		"type": "water",
	},
	"Wave Incense": {
		"id":   uint(317),
		"type": "water",
	},
	"Weakness Policy": {
		"id": uint(639),
	},
	"Whipped Dream": {
		"id": uint(646),
	},
	"White Herb": {
		"id": uint(214),
	},
	"Wide Lens": {
		"id": uint(265),
	},
	"Wiki Berry": {
		"id": uint(160),
	},
	"Wise Glasses": {
		"id": uint(267),
	},
	"X Accuracy": {
		"id": uint(60),
	},
	"X Attack": {
		"id": uint(57),
	},
	"X Defense": {
		"id": uint(58),
	},
	"X Sp. Atk": {
		"id": uint(61),
	},
	"X Sp. Def": {
		"id": uint(62),
	},
	"X Speed": {
		"id": uint(59),
	},
	"Yache Berry": {
		"id":   uint(188),
		"type": "ice",
	},
	"Yellow Nectar": {
		"id": uint(854),
	},
	"Yellow Shard": {
		"id": uint(74),
	},
	"Zap Plate": {
		"id":   uint(300),
		"type": "electric",
	},
	"Zinc": {
		"id": uint(52),
	},
	"Zoom Lens": {
		"id": uint(276),
	},
	"S Tier": {
		"id": uint(1),
	},
	"A Tier": {
		"id": uint(2),
	},
	"B Tier": {
		"id": uint(3),
	},
	"C Tier": {
		"id": uint(4),
	},
	// adding SwSh items below sorted by id and only what's available on Showdown
	"Leek": {
		"id": uint(259),
	},
	"Rusted Sword": {
		"id": uint(1103),
	},

	"Rusted Shield": {
		"id": uint(1104),
	},

	"Fossilized Bird": {
		"id": uint(1105),
	},

	"Fossilized Fish": {
		"id": uint(1106),
	},

	"Fossilized Drake": {
		"id": uint(1107),
	},

	"Fossilized Dino": {
		"id": uint(1108),
	},

	"Strawberry Sweet": {
		"id": uint(1109),
	},

	"Love Sweet": {
		"id": uint(1110),
	},

	"Berry Sweet": {
		"id": uint(1111),
	},

	"Clover Sweet": {
		"id": uint(1112),
	},

	"Flower Sweet": {
		"id": uint(1113),
	},

	"Star Sweet": {
		"id": uint(1114),
	},

	"Ribbon Sweet": {
		"id": uint(1115),
	},

	"Sweet Apple": {
		"id": uint(1116),
	},

	"Tart Apple": {
		"id": uint(1117),
	},

	"Throat Spray": {
		"id": uint(1118),
	},

	"Eject Pack": {
		"id": uint(1119),
	},

	"Heavy-Duty Boots": {
		"id": uint(1120),
	},

	"Blunder Policy": {
		"id": uint(1121),
	},

	"Room Service": {
		"id": uint(1122),
	},

	"Utility Umbrella": {
		"id": uint(1123),
	},
	//I took the liberty of making each ID correspond to the first instance of each TR's type
	"TR00": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR01": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR02": {
		"id":   uint(1132),
		"type": "fire",
	},

	"TR03": {
		"id":   uint(1133),
		"type": "water",
	},

	"TR04": {
		"id":   uint(1133),
		"type": "water",
	},

	"TR05": {
		"id":   uint(1135),
		"type": "ice",
	},

	"TR06": {
		"id":   uint(1135),
		"type": "ice",
	},

	"TR07": {
		"id":   uint(1137),
		"type": "fighting",
	},

	"TR08": {
		"id":   uint(1138),
		"type": "electric",
	},

	"TR09": {
		"id":   uint(1138),
		"type": "electric",
	},

	"TR10": {
		"id":   uint(1140),
		"type": "ground",
	},

	"TR11": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR12": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR13": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR14": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR15": {
		"id":   uint(1132),
		"type": "fire",
	},

	"TR16": {
		"id":   uint(1133),
		"type": "water",
	},

	"TR17": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR18": {
		"id":   uint(1148),
		"type": "bug",
	},

	"TR19": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR20": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR21": {
		"id":   uint(1137),
		"type": "fighting",
	},

	"TR22": {
		"id":   uint(1152),
		"type": "poison",
	},

	"TR23": {
		"id":   uint(1140),
		"type": "ground",
	},

	"TR24": {
		"id":   uint(1154),
		"type": "dragon",
	},

	"TR25": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR26": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR27": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR28": {
		"id":   uint(1148),
		"type": "bug",
	},

	"TR29": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR30": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR31": {
		"id":   uint(1161),
		"type": "steel",
	},

	"TR32": {
		"id":   uint(1162),
		"type": "dark",
	},

	"TR33": {
		"id":   uint(1163),
		"type": "ghost",
	},

	"TR34": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR35": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR36": {
		"id":   uint(1132),
		"type": "fire",
	},

	"TR37": {
		"id":   uint(1162),
		"type": "dark",
	},

	"TR38": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR39": {
		"id":   uint(1137),
		"type": "fighting",
	},

	"TR40": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR41": {
		"id":   uint(1132),
		"type": "fire",
	},

	"TR42": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR43": {
		"id":   uint(1132),
		"type": "fire",
	},

	"TR44": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR45": {
		"id":   uint(1133),
		"type": "water",
	},

	"TR46": {
		"id":   uint(1161),
		"type": "steel",
	},

	"TR47": {
		"id":   uint(1154),
		"type": "dragon",
	},

	"TR48": {
		"id":   uint(1137),
		"type": "fighting",
	},

	"TR49": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR50": {
		"id":   uint(1180),
		"type": "grass",
	},

	"TR51": {
		"id":   uint(1154),
		"type": "dragon",
	},

	"TR52": {
		"id":   uint(1161),
		"type": "steel",
	},

	"TR53": {
		"id":   uint(1137),
		"type": "fighting",
	},

	"TR54": {
		"id":   uint(1152),
		"type": "poison",
	},

	"TR55": {
		"id":   uint(1132),
		"type": "fire",
	},

	"TR56": {
		"id":   uint(1137),
		"type": "fighting",
	},

	"TR57": {
		"id":   uint(1152),
		"type": "poison",
	},

	"TR58": {
		"id":   uint(1162),
		"type": "dark",
	},

	"TR59": {
		"id":   uint(1180),
		"type": "grass",
	},

	"TR60": {
		"id":   uint(1148),
		"type": "bug",
	},

	"TR61": {
		"id":   uint(1148),
		"type": "bug",
	},

	"TR62": {
		"id":   uint(1154),
		"type": "dragon",
	},

	"TR63": {
		"id":   uint(1193),
		"type": "rock",
	},

	"TR64": {
		"id":   uint(1137),
		"type": "fighting",
	},

	"TR65": {
		"id":   uint(1180),
		"type": "grass",
	},

	"TR66": {
		"id":   uint(1196),
		"type": "flying",
	},

	"TR67": {
		"id":   uint(1140),
		"type": "ground",
	},

	"TR68": {
		"id":   uint(1162),
		"type": "dark",
	},

	"TR69": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR70": {
		"id":   uint(1161),
		"type": "steel",
	},

	"TR71": {
		"id":   uint(1180),
		"type": "grass",
	},

	"TR72": {
		"id":   uint(1180),
		"type": "grass",
	},

	"TR73": {
		"id":   uint(1152),
		"type": "poison",
	},

	"TR74": {
		"id":   uint(1161),
		"type": "steel",
	},

	"TR75": {
		"id":   uint(1193),
		"type": "rock",
	},

	"TR76": {
		"id":   uint(1193),
		"type": "rock",
	},

	"TR77": {
		"id":   uint(1180),
		"type": "grass",
	},

	"TR78": {
		"id":   uint(1152),
		"type": "poison",
	},

	"TR79": {
		"id":   uint(1161),
		"type": "steel",
	},

	"TR80": {
		"id":   uint(1138),
		"type": "electric",
	},

	"TR81": {
		"id":   uint(1162),
		"type": "dark",
	},

	"TR82": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR83": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR84": {
		"id":   uint(1133),
		"type": "water",
	},

	"TR85": {
		"id":   uint(1130),
		"type": "normal",
	},

	"TR86": {
		"id":   uint(1138),
		"type": "electric",
	},

	"TR87": {
		"id":   uint(1140),
		"type": "ground",
	},

	"TR88": {
		"id":   uint(1132),
		"type": "fire",
	},

	"TR89": {
		"id":   uint(1196),
		"type": "flying",
	},

	"TR90": {
		"id":   uint(1220),
		"type": "fairy",
	},

	"TR91": {
		"id":   uint(1152),
		"type": "poison",
	},

	"TR92": {
		"id":   uint(1220),
		"type": "fairy",
	},

	"TR93": {
		"id":   uint(1162),
		"type": "dark",
	},

	"TR94": {
		"id":   uint(1140),
		"type": "ground",
	},

	"TR95": {
		"id":   uint(1162),
		"type": "dark",
	},

	"TR96": {
		"id":   uint(1148),
		"type": "bug",
	},

	"TR97": {
		"id":   uint(1141),
		"type": "psychic",
	},

	"TR98": {
		"id":   uint(1133),
		"type": "water",
	},

	"TR99": {
		"id":   uint(1137),
		"type": "fighting",
	},

	"Cracked Pot": {
		"id": uint(1253),
	},

	"Chipped Pot": {
		"id": uint(1254),
	},

	"Galarica Cuff": {
		"id": uint(1582),
	},

	"Galarica Wreath": {
		"id": uint(1592),
	},
}
