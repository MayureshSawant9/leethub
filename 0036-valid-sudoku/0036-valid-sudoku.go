func isValidSudoku(board [][]byte) bool {
    
    type Pair struct{
        row int
        column int
    }

    rowDict := make(map[int]map[byte]struct{})
    columnDict := make(map[int]map[byte]struct{})
    boxDict := make(map[Pair]map[byte]struct{})

    for rowIndex, row := range board {
        for columnIndex, num := range row {
            if num == '.' {
                continue
            }

            if _, ok := rowDict[rowIndex]; !ok {
                rowDict[rowIndex] = make(map[byte]struct{})
            }
            if _, ok := columnDict[columnIndex]; !ok {
                columnDict[columnIndex] = make(map[byte]struct{})
            }
            boxKey := Pair{rowIndex/3, columnIndex/3}
            if _, ok := boxDict[boxKey]; !ok {
                boxDict[boxKey] = make(map[byte]struct{})
            }


            if _, ok := rowDict[rowIndex][num]; ok {
                return false
            } else if _, ok := columnDict[columnIndex][num]; ok {
                return false
            } else if _, ok := boxDict[boxKey][num]; ok {
                return false
            }
            
            rowDict[rowIndex][num] = struct{}{}
            columnDict[columnIndex][num] = struct{}{}
            boxDict[boxKey][num] = struct{}{}
        }
    }
    return true
}