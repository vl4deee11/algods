package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func min(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Дана строка s длины n. Требуется найти все такие пары (i,j), где i<j, что подстрока s[i ... j]
// является палиндромом (т.е. читается одинаково слева направо и справа налево).
//
// Однако информацию о найденных палиндромах можно возвращать более компактно: для каждой позиции i=0 ... n-1
// найдём значения d_1[i] и d_2[i],
// обозначающие количество палиндромов соответственно нечётной и чётной длины с центром в позиции i.
//
// Например, в строке s = abababc есть три палиндрома нечётной длины с центром в символе s[3]=b, т.е. значение d_1[3]=3
// А в строке s = cbaabd есть два палиндрома чётной длины с центром в символе s[3]=a, т.е. значение d_2[3]=2
// Т.е. идея — в том, что если есть подпалиндром длины l с центром в какой-то позиции i,
// то есть также подпалиндромы длины l-2, l-4, и т.д. с центрами в i.
// Поэтому двух таких массивов d_1[i] и d_2[i] достаточно для хранения информации обо всех подпалиндромах этой строки.
// palindromesCnt - за время O(n)
// на вход s - строка, d1, d2 - размеры палиндромов центр которых в индексе dx[i]
func palindromesCnt(s string) ([]int, []int) {
	d1, d2 := make([]int, len(s)), make([]int, len(s))
	n := len(s)
	//  Вычисляем для нечетной длинны
	// Для быстрого вычисления будем поддерживать границы (l, r) самого правого из обнаруженных подпалиндрома
	// (т.е. подпалиндрома с наибольшим значением r). Изначально можно положить l=0, r=-1.
	l := 0
	r := -1

	for i := 0; i < n; i++ {
		// Если i не находится в пределах текущего подпалиндрома, т.е. i > r, то просто выполним подсчет в тупую.
		k := 1
		if i <= r {
			// Попробуем извлечь часть информации из уже подсчитанных значений d_1[].
			// А именно, отразим позицию i внутри подпалиндрома (l,r),
			// т.е. получим позицию j = l + (r - i), и рассмотрим значение d_1[j].
			// Поскольку j — позиция, симметричная позиции i, то почти всегда мы можем просто присвоить
			// d_1[i] = d_1[j].
			// палиндром вокруг j фактически "копируется" в палиндром вокруг i
			// Однако здесь есть тонкость, которую надо обработать правильно: когда "внутренний палиндром"
			// достигает границы внешнего или вылазит за неё, т.е. j-d_1[j]+1 <= l (или, что то же самое, i+d_1[j]-1 >= r).
			// Поскольку за границами внешнего палиндрома никакой симметрии не гарантируется,
			// то просто присвоить d_1[i] = d_1[j] будет уже некорректно:
			// у нас недостаточно сведений, чтобы утверждать, что в позиции i подпалиндром имеет такую же длину.
			//
			// На самом деле, чтобы правильно обрабатывать такие ситуации, надо "обрезать" длину подпалиндрома,
			// т.е. присвоить d_1[i] = r - i.
			// После этого следует пустить тривиальный алгоритм, который будет пытаться увеличить значение d_1[i],
			// пока это возможно.
			k = min(d1[l+r-i], r-i+1)
		}

		// Расходимся в разные стороны палиндрома (досчитаем в тупую)
		for i+k < n && i-k >= 0 && s[i+k] == s[i-k] {
			k++
		}

		//  Запишем
		d1[i] = k
		// Обновим кэш если надо
		if i+k-1 > r {
			l = i - k + 1
			r = i + k - 1
		}
	}

	//  Вычисляем для чётной длинны
	l = 0
	r = -1
	for i := 0; i < n; i++ {
		k := 0
		if i <= r {
			k = min(d2[l+r-i+1], r-i+1)
		}

		// Расходимся в разные стороны палиндрома
		for i+k < n && i-k-1 >= 0 && s[i+k] == s[i-k-1] {
			k++
		}

		d2[i] = k
		if i+k-1 > r {
			l = i - k
			r = i + k - 1
		}
	}
	return d1, d2
}

func TestPalindromeCnt(t *testing.T) {
	str := "abcdedcbahjh"
	d1, d2 := palindromesCnt(str)
	assert.Equal(t, []int{1, 1, 1, 1, 5, 4, 3, 2, 1, 1, 2, 1}, d1)
	str = "funnuf"
	d1, d2 = palindromesCnt(str)
	assert.Equal(t, []int{0, 0, 0, 3, 2, 1}, d2)
}
