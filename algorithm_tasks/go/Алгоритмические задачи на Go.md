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

| Задача                                                                                                                                        | Описание                                                          | Папка         |
| --------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- | ------------- |
| [Concatenation of array](../go/array/concatenation-of-array_test.go)                                                                          | Конкатенация массива с самим собой                                | array         |
| [Set mismatch](algorithm_tasks/go/array/set-mismatch_test.go)                                                                                 | Найти ошибку в наборе чисел (повторяющееся и отсутствующее число) | array         |
| [Shuffle the array](algorithm_tasks/go/array/shuffle-the-array_test.go)                                                                       | Перемешивание массива по заданному алгоритму                      | array         |
| [Max consecutive ones](algorithm_tasks/go/array/max-consecutive-ones_test.go)                                                                 | Найти максимальное количество последовательных единиц в массиве   | array         |
| [Longest common prefix](algorithm_tasks/go/array/longest-common-prefix_test.go)                                                               | Найти самый длинный общий префикс среди строк                     | array         |
| [Find all numbers disappeared in an array](algorithm_tasks/go/array/find-all-numbers-disappeared-in-an-array_test.go)                         | Найти все числа, которые отсутствуют в массиве                    | array         |
| [How many numbers are smaller than the current number](algorithm_tasks/go/array/how-many-numbers-are-smaller-than-the-current-number_test.go) | Для каждого элемента массива посчитать, сколько чисел меньше него | array         |
| [Binary search](algorithm_tasks/go/binary-search/binary-search_test.go)                                                                       | Классический алгоритм бинарного поиска                            | binary-search |
| [Search a 2d matrix](algorithm_tasks/go/binary-search/search-a-2d-matrix_test.go)                                                             | Поиск элемента в отсортированной 2D матрице                       | binary-search |
| [Guess number higher or lower](algorithm_tasks/go/binary-search/guess-number-higher-or-lower_test.go)                                         | Угадать число с использованием подсказок "выше" или "ниже"        | binary-search |
| [Search in rotated sorted array](algorithm_tasks/go/binary-search/search-in-rotated-sorted-array_test.go)                                     | Поиск в отсортированном массиве, который был повернут             | binary-search |
| [Search in rotated sorted array ii](algorithm_tasks/go/binary-search/search-in-rotated-sorted-array-ii_test.go)                               | Поиск в отсортированном повернутом массиве с дубликатами          | binary-search |
| [Find minimum in rotated sorted array](algorithm_tasks/go/binary-search/find-minimum-in-rotated-sorted-array_test.go)                         | Найти минимальный элемент в отсортированном повернутом массиве    | binary-search |
| [Two sums](algorithm_tasks/go/hash-tables/two-sums.go)                                                                                        | Найти два числа в массиве, сумма которых равна заданному значению | hash-tables   |
| [Single number](algorithm_tasks/go/hash-tables/single-number_test.go)                                                                         | Найти единственное число, которое встречается один раз в массиве  | hash-tables   |
| [Add two numbers](algorithm_tasks/go/linked_list/add-two-numbers_test.go)                                                                     | Сложить два числа, представленных в виде связанных списков        | linked_list   |
| [Linked list cycle ii](algorithm_tasks/go/linked_list/linked-list-cycle-ii_test.go)                                                           | Найти начало цикла в связанном списке                             | linked_list   |
| [Linked list cycle](algorithm_tasks/go/linked_list/linked-list-cycle_test.go)                                                                 | Проверить наличие цикла в связанном списке                        | linked_list   |
| [Is palindrome](algorithm_tasks/go/math/is-palindrome_test.go)                                                                                | Проверить, является ли число палиндромом                          | math          |
| [Roman to integer](algorithm_tasks/go/math/roman-to-integer_test.go)                                                                          | Преобразовать римское число в целое                               | math          |
