package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пусть даны три точки p_1, p_2, p_3. Найдём значение знаковой площади S треугольника p_1 p_2 p_3,
// т.е. площади этого треугольника, взятой со знаком плюс или минус в зависимости от типа поворота,
// образуемого точками p_1, p_2, p_3: против часовой стрелки или по ней соответственно.

// Ориентированной площадью треугольника ABC называется величина (ABC), равная его площади, взятой со знаком плюс,
// если обход треугольника в порядке A–B–C–A совершается против
// часовой стрелки и со знаком минус, если по часовой стрелке

// Если мы научимся вычислять такую знаковую ("ориентированную") площадь, то сможем и находить обычную площадь любого
// треугольника, а также сможем проверять, по часовой стрелке или против направлена какая-либо тройка точек.

// Вычисление
// Воспользуемся понятием косого (псевдоскалярного) произведения векторов.
// Оно как раз равно удвоенной знаковой площади треугольника:
// S = 1/2 * vec(AB) ^ vec(AC)
// a = vec(AB)
// b = vec(AC)
// a ^ b = |a||b|sin(∠(a, b)) = 2S => (опр вики: https://ru.wikipedia.org/wiki/%D0%9F%D1%81%D0%B5%D0%B2%D0%B4%D0%BE%D1%81%D0%BA%D0%B0%D0%BB%D1%8F%D1%80%D0%BD%D0%BE%D0%B5_%D0%BF%D1%80%D0%BE%D0%B8%D0%B7%D0%B2%D0%B5%D0%B4%D0%B5%D0%BD%D0%B8%D0%B5),
//
// где угол ∠(a, b) берётся ориентированным, т.е. это угол вращения между этими векторами против часовой стрелки.
//
// (Модуль косого произведения двух векторов равен модулю векторного произведения их.)
//
// Косое произведение вычисляется как величина определителя, составленного из координат точек:
//      | x_1 y_1 1 |
// 2S = | x_2 y_2 1 |
//      | x_3 y_3 1 |
// Раскрывая определитель, можно получить такую формулу:
//
// 2S = x_1 (y_2 - y_3) + x_2 (y_3 - y_1) + x_3 (y_1 - y_2)
//
// Можно сгруппировать третье слагаемое с первыми двумя, избавившись от одного умножения:
//
// 2S = (x_2 - x_1) (y_3 - y_1) - (y_2 - y_1) (x_3 - x_1)

// tArea2 - функция вычисляющая удвоенную знаковую площадь треугольника
func tArea2(x1, y1, x2, y2, x3, y3 int) int {
	return (x2-x1)*(y3-y1) - (y2-y1)*(x3-x1)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// tArea - функция вычисляющая площадь треугольника
func tArea(x1, y1, x2, y2, x3, y3 int) float64 {
	return float64(abs(tArea2(x1, y1, x2, y2, x3, y3))) / 2.0
}

// clockwise - функция проверяющая образует ли указанная тройка точек поворот по часовой стрелке
func clockwise(x1, y1, x2, y2, x3, y3 int) bool {
	return tArea2(x1, y1, x2, y2, x3, y3) < 0
}

// counterClockwise - функция проверяющая образует ли указанная тройка точек поворот против часовой стрелки
func counterClockwise(x1, y1, x2, y2, x3, y3 int) bool {
	return tArea2(x1, y1, x2, y2, x3, y3) > 0
}

func TestTArea(t *testing.T) {
	x1 := 0
	y1 := 0

	x2 := 1
	y2 := 1

	x3 := 2
	y3 := 0

	// Тут идем по часовой
	//  0,0 -> 1,1 -> 2,0
	s := tArea(x1, y1, x2, y2, x3, y3)
	c := clockwise(x1, y1, x2, y2, x3, y3)
	cc := counterClockwise(x1, y1, x2, y2, x3, y3)
	assert.Equal(t, float64(1), s)
	assert.Equal(t, true, c)
	assert.Equal(t, false, cc)

	x1 = 0
	y1 = 0

	x2 = 2
	y2 = 0

	x3 = 1
	y3 = 1

	// Тут идем против часовой
	//  0,0 -> 2,0 -> 1,1
	s = tArea(x1, y1, x2, y2, x3, y3)
	c = clockwise(x1, y1, x2, y2, x3, y3)
	cc = counterClockwise(x1, y1, x2, y2, x3, y3)
	assert.Equal(t, float64(1), s)
	assert.Equal(t, false, c)
	assert.Equal(t, true, cc)
}