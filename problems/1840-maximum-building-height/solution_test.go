package maximumbuildingheight

import "testing"

func TestMaxBuilding(t *testing.T) {
	tests := []struct {
		name         string
		n            int
		restrictions [][]int
		want         int
	}{
		{
			name:         "example 1",
			n:            5,
			restrictions: [][]int{{2, 1}, {4, 1}},
			want:         2,
		},
		{
			name:         "example 2",
			n:            6,
			restrictions: nil,
			want:         5,
		},
		{
			name:         "example 3",
			n:            10,
			restrictions: [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}},
			want:         5,
		},
		{
			name:         "minimum n without restrictions",
			n:            2,
			restrictions: nil,
			want:         1,
		},
		{
			name:         "restriction at last building",
			n:            2,
			restrictions: [][]int{{2, 0}},
			want:         0,
		},
		{
			name:         "duplicate boundary at n",
			n:            10,
			restrictions: [][]int{{10, 3}},
			want:         6,
		},
		{
			name:         "tight restriction near start",
			n:            5,
			restrictions: [][]int{{2, 0}},
			want:         3,
		},
		{
			name:         "large n without restrictions",
			n:            1_000_000_000,
			restrictions: nil,
			want:         999_999_999,
		},
		{
			name:         "large restriction height capped by slope",
			n:            100,
			restrictions: [][]int{{50, 1_000_000_000}},
			want:         99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restrictions := cloneRestrictions(tt.restrictions)
			got := maxBuilding(tt.n, restrictions)
			if got != tt.want {
				t.Fatalf("maxBuilding(%d, %v) = %d, want %d", tt.n, tt.restrictions, got, tt.want)
			}

			inPlaceRestrictions := cloneRestrictions(tt.restrictions)
			gotInPlace := maxBuildingInPlace(tt.n, inPlaceRestrictions)
			if gotInPlace != tt.want {
				t.Fatalf("maxBuildingInPlace(%d, %v) = %d, want %d", tt.n, tt.restrictions, gotInPlace, tt.want)
			}
		})
	}
}

func cloneRestrictions(restrictions [][]int) [][]int {
	if restrictions == nil {
		return nil
	}

	cloned := make([][]int, len(restrictions))
	for i, restriction := range restrictions {
		cloned[i] = append([]int(nil), restriction...)
	}
	return cloned
}
