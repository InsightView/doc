package main

type region struct {
	x, y int
	width, height int
}

func searchRegion(matrix [][]int, r region, target int) bool {
	if r.width == 0 || r.height == 0{
		return false
	}

	if r.width == 1 && r.height == 1{
		return matrix[r.y][r.x] == target
	}

	midX := (r.x + r.width/2)
	midY := (r.y + r.height/2)
	if matrix[midY][midX] == target {
		return true
	}

	if r.width == 1 {
		if matrix[midY][r.x] > target {
			reg := region{
				x: r.x,
				y: r.y,
				width: 1,
				height: r.height/2
			}
			return searchRegion(matrix, reg, target)
		} else {
			reg := region{
				x: r.x,
				y: midY,
				width: 1,
				height: r.height - r.height/2
			}
		}
	}
}

func searchMatric() bool{
	
	return false
}


func main(){
	
}
