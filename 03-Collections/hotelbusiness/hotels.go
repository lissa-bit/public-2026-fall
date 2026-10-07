//go:build !solution

package hotelbusiness

import "sort"

// ComputeLoad builds a sorted hotel occupancy timeline.
func ComputeLoad(guests []Guest) []Load {
	n := len(guests)

	if n == 0 {
		return nil
	}

	delta := make(map[int]int, n*2)

	for i := 0; i < n; i++ {
		g := guests[i]
		delta[g.CheckInDate]++
		delta[g.CheckOutDate]--
	}

	dates := make([]int, 0, len(delta))

	for d := range delta {
		dates = append(dates, d)
	}

	sort.Ints(dates)

	res := make([]Load, 0, len(dates))
	count := 0

	for i := 0; i < len(dates); i++ {
		date := dates[i]

		count += delta[date]
		if delta[date] != 0 {
			res = append(res, Load{StartDate: date, GuestCount: count})
		}
	}

	return res
}
