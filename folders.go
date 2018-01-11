package main

// 1	accommodation
// 2	financial_post_services
// 3	culture_tourism
// 4	public_services
// 5	hiking_cycling
// 6	nature
// 7	emergency_health
// 8	place_of_worship
// 9	food_drink
// 10	shopping
// 11	sport_leisure
// 12	car_services
// 13	transportation

func subToRoot(n int) int {
	if n < 7 {
		return 1
	} else if n < 12 {
		return 2
	} else if n < 19 {
		return 3
	} else if n < 25 {
		return 4
	} else if n < 31 {
		return 5
	} else if n < 36 {
		return 6
	} else if n < 41 {
		return 7
	} else if n < 48 {
		return 8
	} else if n < 55 {
		return 9
	} else if n < 61 {
		return 10
	} else if n < 67 {
		return 11
	} else if n < 72 {
		return 12
	} else if n <= 77 {
		return 13
	}

	return 0
}

// 1	alpine_hut
// 2	camp_caravan
// 3	motel
// 4	hostel
// 5	hotel
// 6	guest_house_chalet
// 7	bank
// 8	exchange
// 9	atm
// 10	post_office
// 11	post_box
// 12	info
// 13	museum
// 14	cinema
// 15	theatre
// 16	castle_ruin_monument
// 17	attraction
// 18	toilets
// 19	townhall
// 20	library
// 21	education
// 22	embassy
// 23	telephone
// 24	grave_yard
// 25	bicycle_parking
// 26	guidepost
// 27	map
// 28	picnic_site
// 29	shelter
// 30	viewpoint
// 31	protected_area
// 32	peak
// 33	spring
// 34	mine_cave
// 35	glacier
// 36	fire_station
// 37	police
// 38	doctor_dentist
// 39	hospital_clinic
// 40	veterinary
// 41	buddhist
// 42	christian
// 43	hindu
// 44	jewish
// 45	muslim
// 46	shinto
// 47	taoist
// 48	bar_pub
// 49	cafe
// 50	restaurant
// 51	fast_food
// 52	confectionery
// 53	drinking_water
// 54	veg_food
// 55	department_store
// 56	pharmacy
// 57	bakery
// 58	other
// 59	supermarket_convenience
// 60	sport_outdoor
// 61	golf
// 62	swimming
// 63	sport_centre
// 64	sport_pitch
// 65	stadium
// 66	skiing
// 67	clubs_dancing
// 68	gas_station
// 69	rest_area
// 70	parking
// 71	car_shop_and_repair
// 72	bus_and_tram_stop
// 73	bus_station
// 74	railway_station
// 75	subway
// 76	airport
// 77	ferries
