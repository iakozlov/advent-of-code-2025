package day8

type Point struct {
	x, y, z int
}

type Edge struct {
	i, j            int // indices of points
	squaredDistance int
}

// Disjoint union set
type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	d := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		d.parent[i] = i
		d.size[i] = 1
	}
	return d
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}

	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	px, py := d.Find(x), d.Find(y)
	if px == py {
		return false
	}

	if d.size[px] < d.size[py] {
		px, py = py, px
	}
	d.parent[py] = px
	d.size[px] += d.size[py]
	return true
}

func buildEdges(points []Point) []Edge {
	n := len(points)
	var edges []Edge
	for i := 0; i < n; i++ {
		pi := points[i]
		for j := i + 1; j < n; j++ {
			pj := points[j]
			dx := pi.x - pj.x
			dy := pi.y - pj.y
			dz := pi.z - pj.z
			squaredDistance := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{i, j, squaredDistance})
		}
	}

	return edges
}
