package geometry

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Алгоритм Грэхема — алгоритм построения выпуклой оболочки в двумерном пространстве.
// В этом алгоритме задача о выпуклой оболочке решается с помощью стека, сформированного из точек-кандидатов.
// Все точки входного множества заносятся в стек, а потом точки, не являющиеся вершинами выпуклой оболочки,
// со временем удаляются из него.
// По завершении работы алгоритма в стеке остаются только вершины оболочки в порядке их обхода против часовой стрелки.

// Точка с координатами
type point struct{ x, y int }

// Вообще point тут - радиус-вектор - вектор из начала координат, ведущий в эту точку.
func vectorMul(a, b *point) int {
	// Простое векторное произведение
	// Если b «слева» от a, то векторное произведение положительное.
	//Если b «справа» от a, то векторное произведение отрицательное.
	// a*b = ∣a∣*∣b∣*sinθ = a.x*b.y - b.x*a.y
	return a.x*b.y - b.x*a.y
}

func vectorMinus(a, b *point) *point {
	// Вычитание двух векторов
	return &point{a.x - b.x, a.y - b.y}
}

// graham - за время O(NlogN) с использованием только операций сравнения, сложения и умножения.
// Алгоритм является асимптотически оптимальным (доказано, что не существует алгоритма с лучшей асимптотикой)
func graham(points []*point) []*point {
	//  p0 - самая левая точка по x координате, если таких несколько то и самая низкая по координате y
	var p0 = points[0]

	// Ищем такую точку
	for i := range points {
		if points[i].x < p0.x || (points[i].x == p0.x && points[i].y < p0.y) {
			p0 = points[i]
		}
	}

	// Сортируем точки по полярному углу относительно точки p0
	sort.Slice(points, func(i, j int) bool {
		return vectorMul(vectorMinus(points[i], p0), vectorMinus(points[j], p0)) > 0
	})

	res := make([]*point, 0)
	for i := range points {
		// Удаляем последнюю точку со стека пока она образует невыпуклость
		for len(res) >= 2 {
			// v1 - от точки points[i] до точки res[len(res)-1]
			var v1 = vectorMinus(points[i], res[len(res)-1])

			// v2 - от точки res[len(res)-1] до точки res[len(res)-2]
			var v2 = vectorMinus(res[len(res)-1], res[len(res)-2])
			// если два последних вектора заворачивают влево, удаляем последнюю точку
			if vectorMul(v1, v2) > 0 {
				res = res[:len(res)-1]
			} else {
				break
			}
		}
		res = append(res, points[i])
	}

	return res
}

func TestGraham(t *testing.T) {
	pts := []*point{{3, 0}, {7, 3}, {12, 2}, {10, 5}, {5, 15}}
	res := graham(pts)
	assert.Equal(t, []*point{{3, 0}, {12, 2}, {5, 15}}, res)
}
