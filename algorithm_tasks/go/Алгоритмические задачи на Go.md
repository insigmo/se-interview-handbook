В этой папке собраны задачи по алгоритмам и структурам данных на языке Go. Каждый файл содержит:
1. **Шаблон функции** - вам нужно реализовать решение задачи
2. **Тесты** - готовые тестовые случаи для локального дебага и проверки решения

## Как работать с задачами

1. Откройте любой файл с расширением `_test.go`
2. Найдите функцию, которую нужно реализовать (например, `isPalindrome(x int) bool`)
3. Напишите свою реализацию этой функции(удалив предварительно тело функции)
4. Запустите тесты командой `go test` в соответствующей папке
5. Проходите все тестовые случаи, пока решение не будет корректным

## Таблица задач

| Задача                                                                                                                           | Описание                                                          | Категория     |
|----------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------|---------------|
| [Concatenation of array](../go/array/concatenation-of-array_test.go)                                                             | Конкатенация массива с самим собой                                | array         |
| [Set mismatch](../go/array/set-mismatch_test.go)                                                                                 | Найти ошибку в наборе чисел (повторяющееся и отсутствующее число) | array         |
| [Shuffle the array](../go/array/shuffle-the-array_test.go)                                                                       | Перемешивание массива по заданному алгоритму                      | array         |
| [Max consecutive ones](../go/array/max-consecutive-ones_test.go)                                                                 | Найти максимальное количество последовательных единиц в массиве   | array         |
| [Longest common prefix](../go/array/longest-common-prefix_test.go)                                                               | Найти самый длинный общий префикс среди строк                     | array         |
| [Find all numbers disappeared in an array](../go/array/find-all-numbers-disappeared-in-an-array_test.go)                         | Найти все числа, которые отсутствуют в массиве                    | array         |
| [How many numbers are smaller than the current number](../go/array/how-many-numbers-are-smaller-than-the-current-number_test.go) | Для каждого элемента массива посчитать, сколько чисел меньше него | array         |
| [Binary search](../go/binary_search/binary_search_test.go)                                                                       | Классический алгоритм бинарного поиска                            | binary_search |
| [Search a 2d matrix](../go/binary_search/search-a-2d-matrix_test.go)                                                             | Поиск элемента в отсортированной 2D матрице                       | binary_search |
| [Guess number higher or lower](../go/binary_search/guess-number-higher-or-lower_test.go)                                         | Угадать число с использованием подсказок "выше" или "ниже"        | binary_search |
| [Search in rotated sorted array](../go/binary_search/search-in-rotated-sorted-array_test.go)                                     | Поиск в отсортированном массиве, который был повернут             | binary_search |
| [Search in rotated sorted array ii](../go/binary_search/search-in-rotated-sorted-array-ii_test.go)                               | Поиск в отсортированном повернутом массиве с дубликатами          | binary_search |
| [Find minimum in rotated sorted array](../go/binary_search/find-minimum-in-rotated-sorted-array_test.go)                         | Найти минимальный элемент в отсортированном повернутом массиве    | binary_search |
| [Two sums](hash_tables/two-sums_test.go)                                                                                         | Найти два числа в массиве, сумма которых равна заданному значению | hash_tables   |
| [Single number](hash_tables/single-number_test.go)                                                                               | Найти единственное число, которое встречается один раз в массиве  | hash_tables   |
| [Add two numbers](../go/linked_list/add-two-numbers_test.go)                                                                     | Сложить два числа, представленных в виде связанных списков        | linked_list   |
| [Linked list cycle ii](../go/linked_list/linked-list-cycle-ii_test.go)                                                           | Найти начало цикла в связанном списке                             | linked_list   |
| [Linked list cycle](../go/linked_list/linked-list-cycle_test.go)                                                                 | Проверить наличие цикла в связанном списке                        | linked_list   |
| [Is palindrome](../go/math/is-palindrome_test.go)                                                                                | Проверить, является ли число палиндромом                          | math          |
| [Roman to integer](../go/math/roman-to-integer_test.go)                                                                          | Преобразовать римское число в целое                               | math          |
