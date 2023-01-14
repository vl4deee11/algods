package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Нахождение площади многоугольника без самопересечений
//
// Пусть дан простой многоугольник (т.е. без самопересечений, но не обязательно выпуклый), требуется вычислить его площадь
// Алгоритм такой: перебрать все рёбра и сложить площади трапеций, ограниченных каждым ребром.
// Площадь нужно брать с тем знаком, с каким она получится (именно благодаря знаку вся "лишняя" площадь сократится).
// Т.е. формула такова: S = 0.5 * (SUM i in [0, n]: (X_i+1 - X_i) * (Y_i+1 + Y_i))
// Проще говоря мы опускаем из каждой точки X_i+1,X_i, перпендикуляр на ось OX, считаем площадь трапеции образованную
// точками [P_i+1, P_i, P_OX_i+1, P_OX_i], отнимаем кусок площади или прибавляем в зависимости
// от уменьшения или увеличения по Х (допускаем, что весь многоугольник лежит в верхней полуплоскости)
// Более подробно с объяснениями на картинках можно увидеть тут: https://wiki.algocode.ru/index.php?title=%D0%9F%D0%BB%D0%BE%D1%89%D0%B0%D0%B4%D1%8C_%D0%BC%D0%BD%D0%BE%D0%B3%D0%BE%D1%83%D0%B3%D0%BE%D0%BB%D1%8C%D0%BD%D0%B8%D0%BA%D0%B0
// polygonSquare - за O(N)
func polygonSquare(fig [][2]float64) float64 {
	res := 0.0

	for i := 0; i < len(fig); i++ {
		p1 := fig[len(fig)-1]
		if i > 0 {
			p1 = fig[i-1]
		}
		p2 := fig[i]
		res += (p2[0] - p1[0]) * (p1[1] + p2[1])
	}

	return fabs(res) / 2
}

func fabs(x float64) float64 {
	if x > 0 {
		return x
	}
	return -x
}

func TestPolygonSquare(t *testing.T) {
	// Пример взял из https://upload.wikimedia.org/wikipedia/commons/thumb/0/0b/Polygon_area_formula_%28English%29.svg/1012px-Polygon_area_formula_%28English%29.svg.png?20111229180347
	fig := [][2]float64{{3, 4}, {5, 11}, {12, 8}, {9, 5}, {5, 6}}
	res := polygonSquare(fig)
	assert.Equal(t, 30.0, res)
}
