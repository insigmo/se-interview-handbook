// https://leetcode.com/problems/remove-invalid-parentheses/description/

package tree

import "testing"

// removeInvalidParentheses возвращает все возможные валидные строки,
// полученные из исходной строки s путём удаления минимального числа скобок.
// Порядок результатов не гарантирован (итерация по map).
func removeInvalidParentheses(s string) []string {
	seen := map[string]struct{}{}

	leftToRemove, rightToRemove := countMinRemovals(s)
	collectValidStrings(s, 0, leftToRemove, rightToRemove, seen)

	result := make([]string, 0, len(seen))
	for validStr := range seen {
		result = append(result, validStr)
	}
	return result
}

// isValid проверяет, является ли строка валидной последовательностью скобок.
// Проходит по каждому символу: увеличивает счётчик на '(' и уменьшает на ')'.
// Если счётчик уходит в минус — встретилась лишняя ')' без пары, строка невалидна.
// В конце счётчик должен быть ноль: все '(' закрыты.
func isValid(str string) bool {
	openCount := 0
	for _, char := range str {
		if char == '(' {
			openCount++
		} else if char == ')' {
			openCount--
			if openCount < 0 {
				return false
			}
		}
	}
	return openCount == 0
}

// countMinRemovals вычисляет минимальное количество '(' и ')', которые нужно удалить,
// чтобы строка стала валидной. Логика: идём слева направо и отслеживаем
// незакрытые '(' через счётчик openLeft. Если встречаем ')', а openLeft == 0,
// значит эту ')' некому закрыть — она лишняя, увеличиваем rightToRemove.
// Иначе ')' закрывает последнюю незакрытую '(', уменьшаем openLeft.
// В конце openLeft содержит число '(', которые так и не закрылись — их тоже надо удалить.
func countMinRemovals(str string) (leftToRemove int, rightToRemove int) {
	openLeft := 0
	for _, char := range str {
		if char == '(' {
			openLeft++
		} else if char == ')' {
			if openLeft == 0 {
				rightToRemove++
			} else {
				openLeft--
			}
		}
	}
	leftToRemove = openLeft
	return
}

// collectValidStrings рекурсивно перебирает все варианты удаления скобок и
// собирает уникальные валидные строки в множество seen.
//
// Параметры:
//
//	str           — текущая строка на данном шаге рекурсии
//	startIndex    — индекс, с которого начинаем перебор (не возвращаемся назад,
//	                чтобы не генерировать дубли)
//	leftToRemove  — сколько '(' ещё нужно удалить
//	rightToRemove — сколько ')' ещё нужно удалить
//	seen          — множество уже найденных валидных строк (дедупликация)
//
// Базовый случай: оба счётчика обнулились — больше удалять нечего.
// Проверяем валидность и, если строка валидна, добавляем в seen.
//
// Рекурсивный случай: перебираем позиции от startIndex до конца строки.
// Пропускаем позицию, если символ совпадает с предыдущим — это привело бы
// к дублирующемуся результату (например, два стоящих рядом '(' дают одно
// и то же после удаления любого из них).
// Если нашли '(' и ещё нужно удалять '(' — удаляем и уходим в рекурсию.
// Аналогично для ')'.
func collectValidStrings(str string, startIndex, leftToRemove, rightToRemove int, seen map[string]struct{}) {
	if leftToRemove == 0 && rightToRemove == 0 {
		if isValid(str) {
			seen[str] = struct{}{}
		}
		return
	}

	for currentIndex := startIndex; currentIndex < len(str); currentIndex++ {
		isDuplicatePosition := currentIndex != startIndex && str[currentIndex] == str[currentIndex-1]
		if isDuplicatePosition {
			continue
		}

		currentChar := str[currentIndex]
		strWithCharRemoved := str[:currentIndex] + str[currentIndex+1:]

		if currentChar == '(' && leftToRemove > 0 {
			collectValidStrings(strWithCharRemoved, currentIndex, leftToRemove-1, rightToRemove, seen)
		}
		if currentChar == ')' && rightToRemove > 0 {
			collectValidStrings(strWithCharRemoved, currentIndex, leftToRemove, rightToRemove-1, seen)
		}
	}
}

func TestRemoveInvalidParentheses(t *testing.T) {
	assertEqual := func(t *testing.T, got, want []string) {
		t.Helper()

		gotSet := make(map[string]struct{})
		for _, s := range got {
			if !isValid(s) {
				t.Fatalf("result contains invalid string: %q", s)
			}
			gotSet[s] = struct{}{}
		}

		if len(gotSet) != len(got) {
			t.Fatalf("result contains duplicates: %v", got)
		}

		wantSet := make(map[string]struct{})
		for _, s := range want {
			wantSet[s] = struct{}{}
		}

		if len(gotSet) != len(wantSet) {
			t.Fatalf("got %v, want %v", got, want)
		}

		for s := range wantSet {
			if _, ok := gotSet[s]; !ok {
				t.Fatalf("missing result %q in %v", s, got)
			}
		}

		for s := range gotSet {
			if _, ok := wantSet[s]; !ok {
				t.Fatalf("unexpected result %q in %v", s, got)
			}
		}
	}

	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "example one",
			s:    "()())()",
			want: []string{"(())()", "()()()"},
		},
		{
			name: "example two with letters",
			s:    "(a)())()",
			want: []string{"(a())()", "(a)()()"},
		},
		{
			name: "example three all invalid",
			s:    ")(",
			want: []string{""},
		},
		{
			name: "already valid no removals",
			s:    "(a(b)c)",
			want: []string{"(a(b)c)"},
		},
		{
			name: "letters only",
			s:    "abc",
			want: []string{"abc"},
		},
		{
			name: "all opening parentheses",
			s:    "(((",
			want: []string{""},
		},
		{
			name: "all closing parentheses",
			s:    ")))",
			want: []string{""},
		},
		{
			name: "single invalid close in middle",
			s:    "a)b(c)d",
			want: []string{"ab(c)d"},
		},
		{
			name: "single invalid open at end",
			s:    "x(",
			want: []string{"x"},
		},
		{
			name: "two unique minimal results",
			s:    "()(()",
			want: []string{"()()"},
		},
		{
			name: "deduplicate identical outcomes",
			s:    "())()",
			want: []string{"()()"},
		},
		{
			name: "duplicate removals with letters",
			s:    "((a)",
			want: []string{"(a)"},
		},
		{
			name: "remove both sides around letters",
			s:    ")a(",
			want: []string{"a"},
		},
		{
			name: "nested valid unchanged",
			s:    "((()))",
			want: []string{"((()))"},
		},
		{
			name: "complex with multiple answers",
			s:    "(r(()()(",
			want: []string{"(r())", "(r)()", "r(())", "r()()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeInvalidParentheses(tt.s)
			assertEqual(t, got, tt.want)
		})
	}
}
