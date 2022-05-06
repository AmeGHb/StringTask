package main

import "fmt"

func main() {
	fmt.Println(brutforceCase(""))
}

func brutforceCase(str string) int {

	/*
		Пометки и концепт работы функции - русский язык.
		Пометки внутри кода - английский язык.

		Время на выполение:
			- Python (30 min): Первое выполнение задачи.
			- Тесты & Code review (45 min).
			- Golang (15 min): Перевод кода на язык Go.
			- Оформление (30 min): Оформление функции и комментарии.

		+:
			- Стабильный алгоритм O(n).
			- Легко поддерживать (простой алгоритм).
		-:
			- Перебор ведётся по каждому элементу последовательности.

		Идеи для улучшения внутри этого алгоритма:
			-(1) Если <maxCount> больше чем сумма оставшихся элементов и <currentCount>,
			то остальная часть последовательности не перебирается {Это оптимизирует
			код, однако влияние на больших масштабах окажет незначительное}.

		Дальнейшие действия:
			- Обращаюсь к TeamLead;
			- Если O(n) - допустимая скорость, то проводятся финальные проверки и
				тесты и push.
			- Если нет, то необходимо либо улучшить существующий алгоритм, либо
				найти новый.

		Идеи для более оптимизированного алгоритма:
			- На базе алгоритма нахождения подстроки в строке Боэра-Мура можно взять
				концепцию перевода каретки (итерация не по всей строке, а по <maxCount>).
				Идея в следующем: при получении занчения <maxCount> > 0 мы можем
				итерироваться на <maxCount>. Если внутри подстроки больше 3 элементов,
				то продолжаем итерацию. Если нет, то возвращаемся и ищем новый <maxCount>.

				В лучшем случае алгоритм даст O(n / <maxCount>), в худшем: O(2n).
	*/

	var stringLen int = len(str)

	switch stringLen {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	}

	var letter1 string = string(str[0])
	var letter2 string = string(str[1])
	var maxCount int = 0
	var currentCount int = 2

	// <lastLetterRowCount> will show how much the same letter was in the
	// line before a new letter appeared. When the program will recount
	// <currentCount> it will add this number + 1. For example: "abbbbcc" ->
	// The pointer at the end of "b" line and <lastLetterRowCount> = 4. The next
	// step will make a new pair of letters: from ("a", "b") to ("b", "c") and
	// make a new <currentCount> = <lastLetterRowCount> + 1 = 4 + 1 = 5,
	// because the new greatest line is -> "bbbbc".
	var lastLetterRowCount int

	if letter1 == letter2 {
		lastLetterRowCount = 2
	} else {
		lastLetterRowCount = 1
	}

	for index_ := 2; index_ < stringLen; index_++ {

		// *reference: Идеи для улучшения внутри этого алгоритма (1).
		if stringLen-index_+currentCount < maxCount {
			return maxCount
		}

		if bfCase__manager(str, &letter1, &letter2, index_, &currentCount,
			&maxCount, &lastLetterRowCount) {
			currentCount += 1
		}
	}

	if currentCount > maxCount {
		return currentCount
	}
	return maxCount
}

func bfCase__manager(str string,
	letter1, letter2 *string,
	index_ int,
	currentCount, maxCount, lastLetterRowCount *int) bool {

	var currentString string = string(str[index_])

	if currentString == *letter1 || currentString == *letter2 {

		// <lastLetterRowCount> parsing.
		if currentString == string(str[index_-1]) {
			*lastLetterRowCount += 1
		} else {
			*lastLetterRowCount = 1
		}

	} else {

		if *letter1 == *letter2 {

			// "aaab". <letter1> = "a", <letter2> = "a", <currentString> = "b".
			// Changing <letter2> to "b" and keep moving forward.
			*letter2 = currentString

		} else {

			if *currentCount > *maxCount {
				*maxCount = *currentCount
			}

			*letter1 = string(str[index_-1])
			*letter2 = currentString
			*currentCount = *lastLetterRowCount + 1
			*lastLetterRowCount = 1
			return false
		}
	}
	return true
}
