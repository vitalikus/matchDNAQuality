package main

import "fmt"

func main() {
	type Relation struct {
		totalOverlapMin     float32
		totalOverlapMax     float32
		longestSegmentMin   float32
		longestSegmentMax   float32
		usingLongestSegment bool
		relationName        string
	}

	var totalOverlap float32
	var longestSegment float32
	// var generation float32
	var relationName string

	var rel [20]Relation
	rel[0] = Relation{3400, 3748, 280, 290, true, "Parent / Child"}
	rel[1] = Relation{2643, 2802, 50, 141, true, "Sibling"}
	rel[2] = Relation{1087, 2297, -1, -1, false, "Grandparent / Grandchild"}
	rel[3] = Relation{1320, 2100, -1, -1, false, "Half-sibling"}
	rel[4] = Relation{1526, 2082, 95, 163, true, "Uncle / Aunt / Newphew / Niece"}
	rel[5] = Relation{572, 1231, -1, -1, false, "Great-Grandparent / Grandchild"}
	rel[6] = Relation{684, 1122, 54, 112, true, "Half- Uncle / Aunt / Nephew / Niece"}
	rel[7] = Relation{548, 1034, 50, 182, true, "1st Cousin"}
	rel[8] = Relation{550, 860, 60, 135, true, "Great- Uncle / Aunt / Nephew / Niece"}
	rel[9] = Relation{220, 638, 34, 106, true, "1st Cousin Once Removed"}
	rel[10] = Relation{107, 426, 21, 64, true, "1st Cousin Twice Removed"}
	rel[11] = Relation{101, 378, 31, 61, true, "2nd Cousin"}
	rel[12] = Relation{320, 330, 30, 35, true, "Half 1st Cousin"}
	rel[13] = Relation{195, 205, 23, 33, true, "Half 1st Cousin Once Removed"}
	rel[14] = Relation{19, 197, 19, 81, true, "2nd Cousin Once Removed"}
	rel[15] = Relation{43, 150, -1, -1, false, "3rd Cousin"}
	rel[16] = Relation{11, 99, -1, -1, false, "3nd Cousin Once Removed"}
	rel[17] = Relation{12, 73, -1, -1, false, "2nd Cousin Twice Removed"}
	rel[18] = Relation{5, 50, 0, 24, true, "4th Cousin"}
	rel[19] = Relation{0, 33, 0, 10, true, "5th Cousin"}

	fmt.Print("Enter a Longest Segment and Total overlap: ")
	fmt.Scanf("%f%f", &longestSegment, &totalOverlap)

	fmt.Println("Longest Segment:", longestSegment)
	fmt.Println("Total Overlap:", totalOverlap)

	for i := 0; i < 20; i++ {
		fmt.Println("Iteration: ", i)
		fmt.Println("Total Overlap Min:", rel[i].totalOverlapMin)
		fmt.Println("Total Overlap Max:", rel[i].totalOverlapMax)

		if totalOverlap >= rel[i].totalOverlapMin && totalOverlap <= rel[i].totalOverlapMax {
			fmt.Println("Total Overlap in the range")
			if rel[i].usingLongestSegment {
				fmt.Println("Using longest Segment for matching...")
				if longestSegment >= rel[i].longestSegmentMin && longestSegment <= rel[i].longestSegmentMax {
					fmt.Println("Longest Segment is in the range")
					fmt.Println("Longest Segment Min:", rel[i].longestSegmentMin)
					fmt.Println("Longest Segment Max:", rel[i].longestSegmentMax)
					relationName = rel[i].relationName
					i = 20
				}
			} else {
				relationName = rel[i].relationName
				i = 20
			}
		}
	}

	fmt.Println("Relation is", relationName)
}
