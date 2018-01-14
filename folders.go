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

func tagsToSubs(tags map[string]string) *[]int {
	subs := make([]int, 0, 4)
	for key, val := range tags {
		tagVals, tagValsPresent := tagSubFolder[key]
		if tagValsPresent {
			tagSubs, tagSubsPresent := tagVals[val]
			if tagSubsPresent {
				subs = append(subs, tagSubs...)
			}
		}
	}
	return &subs
}

func tagsToFolders(tags map[string]string) *[][]int {
	result := make([][]int, 0, 2)
	subs := tagsToSubs(tags)
	for _, sub := range *subs {
		result = append(result, []int{subToRoot(sub), sub})
	}

	return &result
}
