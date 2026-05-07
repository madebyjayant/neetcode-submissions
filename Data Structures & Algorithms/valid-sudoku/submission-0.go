func isValidSudoku(board [][]byte) bool {
	// for row
	for i:=0;i<9;i++{
		freqMap := make(map[string]int, 8)
		for j:=0;j<9;j++{
			if string(board[i][j])!="."{
				key := string(board[i][j])
				if _, ok := freqMap[key];!ok{
					freqMap[key]++
					continue;
				}
				return false
			}
		}
	}

	// for col
	for j:=0;j<9;j++{
		freqMap := make(map[string]int,8)
		for i:=0;i<9;i++{
			if string(board[i][j])!="."{
				key := string(board[i][j])
				if _, ok := freqMap[key];!ok{
					freqMap[key]++
					continue;
				}
				return false
			}
		}
	}

	// for squares
	offset:=3;

	for k:=0; k<9;k+=offset{
		for l:=0;l<9;l+=offset{
			// iterate through sq.
			freqMap := make(map[string]int,8)
			for i:=k;i<k+offset;i++{
				for j:=l;j<l+offset;j++{
					if string(board[i][j])!="."{
						key := string(board[i][j])
						if _, ok := freqMap[key];!ok{
							freqMap[key]++
							continue;
						}
						return false
					}
				}
			}
		}
	}


	return true
}
