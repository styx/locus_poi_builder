package main

var foldersRoot = map[string]int{
	"accommodation":           1,
	"financial_post_services": 2,
	"culture_tourism":         3,
	"public_services":         4,
	"hiking_cycling":          5,
	"nature":                  6,
	"emergency_health":        7,
	"place_of_worship":        8,
	"food_drink":              9,
	"shopping":                10,
	"sport_leisure":           11,
	"car_services":            12,
}

// 1	alpine_hut
// tourism	alpine_hut

// 2	camp_caravan
// tourism	camp_site
// tourism	caravan_site

// 3	motel
// tourism	motel

// 4	hostel
// tourism	hostel

// 5	hotel
// tourism	hotel

// 6	guest_house_chalet
// tourism	guest_house
// tourism	chalet

// 7	bank
// amenity	bank

// 8	exchange
// amenity	bureau_de_change

// 9	atm
// amenity	atm

// 10	post_office
// amenity	post_office

// 11	post_box
// amenity	post_box

// 12	info
// tourism	information
// information	office

// 13	museum
// tourism	museum

// 14	cinema
// amenity	cinema

// 15	theatre
// amenity	theatre

// 16	castle_ruin_monument
// historic	ruins
// historic	memorial
// historic	monument
// historic	castle
// man_made	obelisk
// tourism	memorial
// man_made	tower

// 17	attraction
// tourism	attraction
// historic	castle
// historic	memorial
// amenity	place_of_worship
// historic	ruins
// historic	building
// tourism	zoo
// natural	stone
// natural	peak
// historic	manor
// historic	monument
// historic	fort
// man_made	water_well
// man_made	flagpole
// man_made	tower
// man_made	lighthouse
// historic	stone
// historic	cannon
// historic	tank
// amenity	planetarium
// leisure	garden
// amenity	cinema
// amenity	monastery
// historic	monastery
// leisure	nature_reserve
// leisure	park
// landuse	park
// natural	spring
// historic	industrial
// historic	battlefield
// historic	boundary_stone
// historic	mine
// shop	mall
// man_made	water_tower
// landuse	churchyard
// historic	church
// historic	tomb
// historic	city_gate

// 18	toilets
// amenity	toilets

// 19	townhall
// amenity	townhall

// 20	library
// amenity	library

// 21	education
// amenity	university
// amenity	college
// amenity	school
// amenity	music_school
// amenity	driving_school
// amenity	language_school

// 22	embassy
// amenity	embassy

// 23	telephone
// amenity	telephone
// emergency	phone

// 24	grave_yard
// amenity	grave_yard
// landuse	cemetery

// 25	bicycle_parking
// amenity	bicycle_parking

// 26	guidepost
// information	guidepost

// 27	map
// information	map

// 28	picnic_site
// tourism	picnic_site
// leisure	picnic_table
// amenity	bbq

// 29	shelter
// amenity	shelter

// 30	viewpoint
// tourism	viewpoint
// historic	memorial
// historic	monument

// 31	protected_area
// boundary	protected_area

// 32	peak
// natural	peak

// 33	spring
// natural	spring

// 34	mine_cave
// man_made	mineshaft
// natural	cave_entrance

// 35	glacier
// natural	glacier

// 36	fire_station
// amenity	fire_station

// 37	police
// amenity	police

// 38	doctor_dentist
// amenity	dentist
// amenity	doctors

// 39	hospital_clinic
// amenity	hospital
// amenity	clinic

// 40	veterinary
// amenity	veterinary
// shop	veterinary
// shop	pet

// 41	buddhist
// religion	buddhist

// 42	christian
// religion	christian

// 43	hindu
// religion	hindu

// 44	jewish
// religion	jewish

// 45	muslim
// religion	muslim

// 46	shinto
// religion	shinto

// 47	taoist
// religion	taoist

// 48	bar_pub
// amenity	bar
// amenity	pub

// 49	cafe
// amenity	cafe

// 50	restaurant
// amenity	restaurant

// 51	fast_food
// amenity	fast_food

// 52	confectionery
// shop	confectionery

// 53	drinking_water
// amenity	drinking_water

// 54	veg_food
// diet:vegan	yes
// diet:vegetarian	yes

// 55	department_store
// shop	department_store
// shop	mall

// 56	pharmacy
// amenity	pharmacy

// 57	bakery
// shop	bakery

// 58	other
// shop	general

// 59	supermarket_convenience
// shop	supermarket
// shop	convenience

// 60	sport_outdoor
// shop	sports
// shop	bicycle
// shop	outdoor

// 61	golf
// leisure	golf_course
// leisure	miniature_golf
// sport	golf

// 62	swimming
// amenity	swimming_pool
// leisure	swimming_pool
// leisure	water_park
// natural	beach

// 63	sport_centre
// leisure	sports_centre
// leisure	fitness_centre
// club	sport
// amenity	gym

// 64	sport_pitch
// leisure	pitch

// 65	stadium
// leisure	stadium

// 66	skiing
// sport	skiing

// 67	clubs_dancing
// club	music

// 68	gas_station
// amenity	fuel

// 69	rest_area
// highway	rest_area

// 70	parking
// amenity	parking

// 71	car_shop_and_repair
// shop	car_repair
// shop	car
// shop	car_parts
// amenity	car_wash
// amenity	vehicle_inspection

// 72	bus_and_tram_stop
// highway	bus_stop
// railway	tram_stop

// 73	bus_station
// amenity	bus_station

// 74	railway_station
// railway	halt

// 75	subway
// station	subway

// 76	airport
// aeroway	aerodrome

// 77	ferries
// amenity	ferry_terminal
var tagSubFolder = map[string]map[string][]int{
	"tourism": {
		// alpine_hut
		"alpine_hut": {1},

		// camp_caravan
		"camp_site":    {2},
		"caravan_site": {2},

		// motel
		"motel": {3},

		// hostel
		"hostel": {4},

		// 5 hotel
		"hotel": {5},

		// guest_house_chalet
		"guest_house": {6},
		"chalet":      {6},

		// info
		"information": {12},

		// museum
		"museum": {13},

		// 16 castle_ruin_monument
		"memorial": {16},

		// 17	attraction
		"attraction": {17},
		"zoo":        {17},

		// 28 picnic_site
		"picnic_site": {28},

		// 30 viewpoint
		"viewpoint": {30},
	},
	"amenity": {
		// bank
		"bank": {7},
		// exchange
		"bureau_de_change": {8},
		// atm
		"atm": {9},
		// post_office
		"post_office": {10},
		// post_box
		"post_box": {11},
		// 14 cinema, 17 attraction
		"cinema": {14, 17},
		// theatre
		"theatre": {15},

		// 17 attraction
		"place_of_worship": {17},
		"planetarium":      {17},
		"monastery":        {17},

		// toilets
		"toilets": {18},

		// townhall
		"townhall": {19},

		// library
		"library": {20},

		// 21 education
		"university":      {21},
		"college":         {21},
		"school":          {21},
		"music_school":    {21},
		"driving_school":  {21},
		"language_school": {21},

		// embassy
		"embassy": {22},

		// 23 telephone
		"telephone": {23},

		// 24 grave_yard
		"grave_yard": {24},

		// bicycle_parking
		"bicycle_parking": {25},

		// 28 picnic_site
		"bbq": {28},

		// shelter
		"shelter": {29},

		// 36 fire_station
		"fire_station": {36},

		// 37 police
		"police": {37},

		// 38 doctor_dentist
		"dentist": {38},
		"doctors": {38},

		// 39 hospital_clinic
		"hospital": {39},
		"clinic":   {39},

		// 40 veterinary
		"veterinary": {40},

		// 48 bar_pub
		"pub": {48},
		"bar": {48},

		// 49 cafe
		"cafe": {49},

		// 50 restaurant
		"restaurant": {50},

		// 51 fast_food
		"fast_food": {51},

		// 53 drinking_water
		"drinking_water": {53},

		// 56 pharmacy
		"pharmacy": {56},

		// 62 swimming
		"swimming_pool": {62},

		// 63 sport_centre
		"gym": {63},

		// 68 gas_station
		"fuel": {68},

		// 70 parking
		"parking": {70},

		// 73 bus_station
		"bus_station": {73},

		// 71	car_shop_and_repair
		"car_wash":           {71},
		"vehicle_inspection": {71},

		// 77 ferries
		"ferry_terminal": {77},
	},
	"information": {
		// 12 info
		"office": {12},
		// 26 guidepost
		"guidepost": {26},
		// 27 map
		"map": {27},
	},
	"emergency": {
		// 23 telephone
		"phone": {23},
	},
	"landuse": {
		// 17 attraction
		"park":       {17},
		"churchyard": {17},

		// 24 grave_yard
		"cemetery": {24},
	},
	"leisure": {
		// 17 attraction
		"garden":         {17},
		"park":           {17},
		"nature_reserve": {17},
		// 28 picnic_site
		"picnic_table": {28},
		// 61 golf
		"golf_course":    {61},
		"miniature_golf": {61},
		// 62 swimming
		"swimming_pool": {62},
		"water_park":    {62},
		// 63 sport_centre
		"sports_centre":  {63},
		"fitness_centre": {63},
		// 64 sport_pitch
		"pitch": {64},
		// 65 stadium
		"stadium": {65},
	},
	"boundary": {
		"protected_area": {31},
	},
	"natural": {
		// 17 attraction
		"stone": {17},
		// 32 peak
		"peak": {17, 32},
		// 33 spring
		"spring": {17, 33},
		// 34 mine_cave
		"cave_entrance": {34},
		// 35 glacier
		"glacier": {35},
		// 62 swimming
		"beach": {62},
	},
	"man_made": {
		// 16 castle_ruin_monument
		// 17 attraction
		"obelisk":     {16},
		"tower":       {16, 17},
		"water_well":  {17},
		"flagpole":    {17},
		"lighthouse":  {17},
		"water_tower": {17},
		// 34 mine_cave
		"mineshaft": {34},
	},
	"religion": {
		// 41 buddhist
		"buddhist": {41},
		// 42 christian
		"christian": {42},
		// 43 hindu
		"hindu": {43},
		// 44 jewish
		"jewish": {44},
		// 45 muslim
		"muslim": {45},
		// 46 shinto
		"shinto": {46},
		// 47 taoist
		"taoist": {47},
	},
	"shop": {
		// 40 veterinary
		"veterinary": {40},
		"pet":        {40},
		// 52 confectionery
		"confectionery": {52},
		// 17 attraction
		// 55 department_store
		"mall":             {17, 55},
		"department_store": {55},
		// 57 bakery
		"bakery": {57},
		// 58 other
		"general": {58},
		// 59 supermarket_convenience
		"supermarket": {59},
		"convenience": {59},
		// 60 sport_outdoor
		"sports":  {60},
		"bicycle": {60},
		"outdoor": {60},

		// 71 car_shop_and_repair
		"car":        {71},
		"car_repair": {71},
		"car_parts":  {71},
	},
	"diet:vegan": {
		// 54 veg_food
		"yes": {54},
	},
	"diet:vegetarian": {
		// 54 veg_food
		"yes": {54},
	},
	"sport": {
		// 61 golf
		"golf": {61},
		// 66 skiing
		"skiing": {66},
	},
	"club": {
		"sport": {63},
		// 67 clubs_dancing
		"music": {67},
	},
	"highway": {
		// 69 rest_area
		"rest_area": {69},
		// 72 bus_and_tram_stop
		"bus_stop": {72},
	},
	"railway": {
		// 72 bus_and_tram_stop
		"tram_stop": {72},
		// 74 railway_station
		"halt": {74},
	},
	"station": {
		// 75 subway
		"subway": {75},
	},
	"aeroway": {
		// 76 airport
		"aerodrome": {76},
	},
	"historic": {
		// 16 castle_ruin_monument
		// 17 attraction
		"ruins":          {16, 17},
		"memorial":       {16, 17},
		"monument":       {16, 17},
		"castle":         {16, 17},
		"building":       {17},
		"manor":          {17},
		"fort":           {17},
		"stone":          {17},
		"cannon":         {17},
		"tank":           {17},
		"monastery":      {17},
		"industrial":     {17},
		"battlefield":    {17},
		"boundary_stone": {17},
		"mine":           {17},
		"church":         {17},
		"tomb":           {17},
		"city_gate":      {17},
	},
	"attraction": {
		"amusement_ride": {17},
		"animal": {17},
		"big_wheel": {17},
		"bumper_car": {17},
		"bungee_jumping": {17},
		"carousel": {17},
		"dark_ride": {17},
		"drop_tower": {17},
		"formal_garden": {17},
		"kiddie_ride": {17},
		"log_flume": {17},
		"maze": {17},
		"pirate_ship": {17},
		"river_rafting": {17},
		"roller_coaster": {17},
		"summer_toboggan": {17},
		"swing_carousel": {17},
		"train": {17},
		"water_slide": {17},
	}
}
