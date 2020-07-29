package _16puzzle

import "../base"

func GetExample() base.Board {
	// true: 14
	//b := [base.size][base.size]int{
	//	{2, 3, 6, 4},
	//	{1, 5, 7, 8},
	//	{9, 14, 10, 11},
	//	{13, 15, 12, 0},
	//}
	// true: 20, slightly better with H3, but makes no sense since the path is so short
	b := [size][size]int{
		{2, 3, 7, 6},
		{1, 0, 5, 4},
		{9, 14, 10, 8},
		{13, 15, 12, 11},
	}
	// true: 30, BFS and IDDFS cannot be used from now on
	//b := [base.size][base.size]int{
	//	{2, 7, 5, 6},
	//	{1, 3, 10, 4},
	//	{13, 14, 0, 8},
	//	{15, 9, 12, 11},
	//}
	// true: 46
	//b := [base.size][base.size]int{
	//	{7, 8, 5, 6},
	//	{1, 3, 10, 4},
	//	{14, 13, 0, 11},
	//	{15, 9, 12, 2},
	//}
	// true: 78, A* search uses so much memory, can only be executed with IDA* search
	//b := [base.size][base.size]int{
	//	{0, 15, 14, 13},
	//	{12, 11, 10, 9},
	//	{8, 7, 6, 5},
	//	{4, 3, 2, 1},
	//}
	// true: 72, corner case for WD heuristics
	//b := [base.size][base.size]int{
	//	{1, 5, 9, 13},
	//	{2, 6, 10, 14},
	//	{3, 7, 11, 15},
	//	{4, 8, 12, 0},
	//}
	// true: 58
	//b := [base.size][base.size]int{
	//	{5, 6, 10, 15},
	//	{7, 9, 14, 1},
	//	{2, 12, 13, 11},
	//	{8, 3, 4, 0},
	//}
	return newBoard(b)
}
