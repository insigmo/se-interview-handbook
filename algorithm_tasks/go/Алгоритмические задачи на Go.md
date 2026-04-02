В этой папке собраны задачи по алгоритмам и структурам данных на языке Go. Задачи которые решаются, были взяты из [файла](../../docs/leetcode_tasks), где перечислен необходимый минимум для прохождения АА секции на собеседованиях. Каждый файл содержит:
1. **Шаблон функции** - вам нужно реализовать решение задачи
2. **Тесты** - готовые тестовые случаи для локального дебага и проверки решения

## Как работать с задачами

1. Откройте любой файл с расширением `_test.go`
2. Найдите функцию, которую нужно реализовать (например, `isPalindrome(x int) bool`)
3. Напишите свою реализацию этой функции(удалив предварительно тело функции)
4. Запустите тесты командой `go test` в соответствующей папке
5. Проходите все тестовые случаи, пока решение не будет корректным

## Таблица задач

| Задача                                                                                                                            | Описание                                                          | Категория     |
|-----------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------|---------------|
| [Concatenation of array](../go/array/concatenation-of-array_test.go)                                                              | Конкатенация массива с самим собой                                | array         |
| [Set mismatch](../go/array/set-mismatch_test.go)                                                                                  | Найти ошибку в наборе чисел (повторяющееся и отсутствующее число) | array         |
| [Shuffle the array](../go/array/shuffle-the-array_test.go)                                                                        | Перемешивание массива по заданному алгоритму                      | array         |
| [Max consecutive ones](../go/array/max-consecutive-ones_test.go)                                                                  | Найти максимальное количество последовательных единиц в массиве   | array         |
| [Longest common prefix](../go/array/longest-common-prefix_test.go)                                                                | Найти самый длинный общий префикс среди строк                     | array         |
| [Find all numbers disappeared in an array](../go/array/find-all-numbers-disappeared-in-an-array_test.go)                          | Найти все числа, которые отсутствуют в массиве                    | array         |
| [How many numbers are smaller than the current number](../go/array/how-many-numbers-are-smaller-than-the-current-number_test.go)  | Для каждого элемента массива посчитать, сколько чисел меньше него | array         |
| [Two sums](../go/array/two-sums_test.go)                                                                                          | Найти два числа в массиве, сумма которых равна заданному значению | array         |
| [Two sum ii input array is sorted](../go/array/two-sum-ii-input-array-is-sorted_test.go)                                          | Найти два числа в отсортированном массиве с заданной суммой       | array         |
| [3sum](../go/array/3sum_test.go)                                                                                                  | Найти все тройки чисел в массиве с нулевой суммой                 | array         |
| [4sum](../go/array/4sum_test.go)                                                                                                  | Найти все четвёрки чисел в массиве с заданной суммой              | array         |
| [Valid anagram](../go/array/valid-anagram_test.go)                                                                                | Проверить, является ли строка анаграммой другой строки            | array         |
| [Group anagrams](../go/array/group-anagrams_test.go)                                                                              | Сгруппировать строки-анаграммы                                    | array         |
| [Find all anagrams in a string](../go/array/find-all-anagrams-in-a-string_test.go)                                                | Найти все анаграммы подстроки в строке                            | array         |
| [Longest repeating character replacement](../go/array/longest-repeating-character-replacement_test.go)                            | Найти длину самой длинной подстроки после замены символов         | array         |
| [Container with most water](../go/array/container-with-most-water_test.go)                                                        | Найти контейнер с наибольшим объёмом воды                         | array         |
| [Merge intervals](../go/array/merge-intervals_test.go)                                                                            | Объединить пересекающиеся интервалы                               | array         |
| [Partition labels](../go/array/partition-labels_test.go)                                                                          | Разбить строку на максимальное количество частей                  | array         |
| [Top k frequent elements](../go/array/top-k-frequent-elements_test.go)                                                            | Найти k наиболее часто встречающихся элементов                    | array         |
| [Top k frequent words](../go/array/top-k-frequent-words_test.go)                                                                  | Найти k наиболее часто встречающихся слов                         | array         |
| [Binary search](../go/binary_search/binary-search_test.go)                                                                        | Классический алгоритм бинарного поиска                            | binary_search |
| [Search a 2d matrix](../go/binary_search/search-a-2d-matrix_test.go)                                                              | Поиск элемента в отсортированной 2D матрице                       | binary_search |
| [Guess number higher or lower](../go/binary_search/guess-number-higher-or-lower_test.go)                                          | Угадать число с использованием подсказок "выше" или "ниже"        | binary_search |
| [Search in rotated sorted array](../go/binary_search/search-in-rotated-sorted-array_test.go)                                      | Поиск в отсортированном массиве, который был повернут             | binary_search |
| [Search in rotated sorted array ii](../go/binary_search/search-in-rotated-sorted-array-ii_test.go)                                | Поиск в отсортированном повернутом массиве с дубликатами          | binary_search |
| [Find minimum in rotated sorted array](../go/binary_search/find-minimum-in-rotated-sorted-array_test.go)                          | Найти минимальный элемент в отсортированном повернутом массиве    | binary_search |
| [Best time to buy and sell stock](../go/greedy/best-time-to-buy-and-sell-stock_test.go)                                           | Найти максимальную прибыль от одной сделки купли-продажи          | greedy        |
| [Best time to buy and sell stock ii](../go/greedy/best-time-to-buy-and-sell-stock-ii_test.go)                                     | Найти максимальную прибыль от нескольких сделок купли-продажи     | greedy        |
| [Best time to buy and sell stock with cooldown](../go/greedy/best-time-to-buy-and-sell-stock-with-cooldown_test.go)               | Максимальная прибыль от сделок с периодом ожидания                | greedy        |
| [Best time to buy and sell stock with transaction fee](../go/greedy/best-time-to-buy-and-sell-stock-with-transaction-fee_test.go) | Максимальная прибыль от сделок с комиссией за транзакцию          | greedy        |
| [Single number](../go/hash_tables/single-number_test.go)                                                                          | Найти единственное число, которое встречается один раз в массиве  | hash_tables   |
| [Sliding window maximum](../go/heapq/sliding-window-maximum_test.go)                                                              | Найти максимум в скользящем окне                                  | heapq         |
| [Sliding window median](../go/heapq/sliding-window-median_test.go)                                                                | Найти медиану в скользящем окне                                   | heapq         |
| [Add two numbers](../go/linked_list/add-two-numbers_test.go)                                                                      | Сложить два числа, представленных в виде связанных списков        | linked_list   |
| [Linked list cycle ii](../go/linked_list/linked-list-cycle-ii_test.go)                                                            | Найти начало цикла в связанном списке                             | linked_list   |
| [Linked list cycle](../go/linked_list/linked-list-cycle_test.go)                                                                  | Проверить наличие цикла в связанном списке                        | linked_list   |
| [Merge k sorted lists](../go/linked_list/merge_k_sorted_lists_test.go)                                                            | Объединить k отсортированных связанных списков                    | linked_list   |
| [Is palindrome](../go/math/is-palindrome_test.go)                                                                                 | Проверить, является ли число палиндромом                          | math          |
| [Roman to integer](../go/math/roman-to-integer_test.go)                                                                           | Преобразовать римское число в целое                               | math          |
| [Valid parentheses](../go/stack/valid-parentheses_test.go)                                                                        | Проверить корректность скобочной последовательности               | stack         |
| [Balanced binary tree](../go/tree/balanced-binary-tree_test.go)                                                                   | Проверить, является ли дерево сбалансированным                    | tree          |
| [Same tree](../go/tree/same-tree_test.go)                                                                                         | Проверить, одинаковы ли два бинарных дерева                       | tree          |
| [Symmetric tree](../go/tree/symmetric-tree_test.go)                                                                               | Проверить, является ли дерево симметричным                        | tree          |
| [Path sum](../go/tree/path-sum_test.go)                                                                                           | Проверить наличие пути с заданной суммой в дереве                 | tree          |
| [Path sum ii](../go/tree/path-sum-ii_test.go)                                                                                     | Найти все пути с заданной суммой в дереве                         | tree          |
| [Number of islands](../go/tree/number-of-islands_test.go)                                                                         | Подсчитать количество островов в матрице                          | tree          |
| [Remove invalid parentheses](../go/tree/remove-invalid-parentheses_test.go)                                                       | Удалить минимум скобок для получения корректного выражения        | tree          |
