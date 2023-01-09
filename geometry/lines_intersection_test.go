package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Точка пересечения прямых
// Пусть нам даны две прямые, заданные своими коэффициентами A_1, B_1, C_1 и A_2, B_2, C_2.
// Требуется найти их точку пересечения, или выяснить, что прямые параллельны.

// Если две прямые не параллельны, то они пересекаются.
// Чтобы найти точку пересечения, достаточно составить из двух уравнений прямых систему и решить её:
//
// {
// {	A_1 x + B_1 y + C_1 = 0
// {    A_2 x + B_2 y + C_2 = 0
// {
//
// Пользуясь формулой Крамера, сразу находим решение системы, которое и будет искомой точкой пересечения:
//
// x = - ((C_1 * B_2) - (C_2 * B_1)) / ((A_1 * B_2) - (A_2 * B_1))
// y = - ((C_2 * A_1) - (C_1 * A_2)) / ((A_1 * B_2) - (A_2 * B_1))
// Если знаменатель нулевой => ((A_1 * B_2) - (A_2 * B_1)) == 0 => то система решений не имеет
// (прямые параллельны и не совпадают) или имеет бесконечно много (прямые совпадают).
// Если необходимо различить эти два случая, надо проверить, что коэффициенты C прямых пропорциональны с
// тем же коэффициентом пропорциональности, что и коэффициенты A и B,
// для чего достаточно посчитать два определителя, если они оба равны нулю, то прямые совпадают:
//  | A_1, C_1 | | B_1 C_1 |
//  | A_2, C_2 | | B_2 C_2 |

const eps = 1e-6

func det(a, b, c, d float64) float64 {
	return a*d - b*c
}

func abs_float64(x float64) float64 {
	if x > 0 {
		return x
	}
	return -x
}

// intersect - определяет точку пересечения двух прямых и возвращает их координаты + флаг
func intersect(a1, b1, c1, a2, b2, c2 float64) (float64, float64, bool) {
	// Проверка на паралельность
	zn := det(a1, b1, a2, b2)
	if abs_float64(zn) < eps {
		return 0, 0, false
	}
	//

	ptx := -det(c1, b1, c2, b2) / zn
	pty := -det(a1, c1, a2, c2) / zn
	return ptx, pty, true
}

// parallel определяет паралельны ли две прямые или нет
func parallel(a1, b1, c1, a2, b2, c2 float64) bool {
	return abs_float64(det(a1, b1, a2, b2)) < eps
}

// equivalent определяет совпадают ли две прямые или нет
func equivalent(a1, b1, c1, a2, b2, c2 float64) bool {
	return abs_float64(det(a1, b1, a2, b2)) < eps &&
		abs_float64(det(a1, c1, a2, c2)) < eps &&
		abs_float64(det(b1, c1, b2, c2)) < eps
}

func TestIntersection(t *testing.T) {
	a1, b1, c1 := 0.0, 0.0, 0.0
	a2, b2, c2 := 0.0, 0.0, 0.0

	_, _, ok := intersect(a1, b1, c1, a2, b2, c2)
	assert.Equal(t, false, ok)
	assert.Equal(t, true, equivalent(a1, b1, c1, a2, b2, c2))

	a1, b1, c1 = 0.0, 0.0, 555.0
	_, _, ok = intersect(a1, b1, c1, a2, b2, c2)
	assert.Equal(t, false, ok)
	assert.Equal(t, true, parallel(a1, b1, c1, a2, b2, c2))

	a1, b1, c1 = 53.0, 43.0, 555.0
	a2, b2, c2 = 34.0, -54.0, 65.0
	x, y, ok := intersect(a1, b1, c1, a2, b2, c2)

	assert.Equal(t, true, ok)
	assert.Equal(t, -7.577474560592044, x)
	assert.Equal(t, -3.567298797409806, y)
}
