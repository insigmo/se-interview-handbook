# Вопросы и задачи с собеседований

## ОЗОН (разработка)

### Задача / Вопрос
Требуется реализовать функцию `uniqRandn`, которая генерирует слайс длины *n* уникальных, рандомных чисел.

**Python:**
```python
from random import randint

def uniqRadn(n: int) -> list[int]:
    answer = set()
    max_int = 2 ** 64 - 1
    while len(answer) < n:
        random_int = randint(0, max_int)
        if random_int not in answer:
            answer.add(random_int)
    return list(answer)
````

---

### Дан отсортированный массив чисел

Нужно вернуть отсортированный массив квадратов этих чисел.

```python
def func(nums: list[int]) -> list[int]:
    left_index = 0
    right_index = len(nums) - 1
    result = []

    for i in range(right_index + 1):
        left_value = nums[left_index] ** 2
        right_value = nums[right_index] ** 2
        if left_value > right_value:
            result.append(nums[left_index] ** 2)
            left_index += 1
        else:
            result.append(nums[right_index] ** 2)
            right_index -= 1

    return result[::-1]
```

---

### Аналог `zip` из Python на Go

(Задача без приведённого решения.)

---

### Восстановление пароля по хешу

Дано:

- база с хешированными паролями (`hashPassword`)
    
- алфавит допустимых символов (`alphabet`)
    

Нужно реализовать `RecoverPassword`, чтобы тест `TestRecoverPassword` проходил успешно.

---

### Реализовать LRU Cache (задача со звёздочкой)

---

## SQL Задачи

### Таблицы

**user**

```
id | firstname | lastname | birth
1  | Ivan      | Petrov   | 1996-05-01
2  | Anna      | Petrova  | 1999-06-01
3  | Anna      | Petrova  | 1990-10-02
```

**purchase**

```
sku | price | user_id | date
1   | 5500  | 1       | 2021-02-15
1   | 5700  | 1       | 2021-01-15
2   | 4000  | 1       | 2021-02-14
3   | 8000  | 2       | 2021-03-01
4   | 400   | 2       | 2021-03-02
```

**ban_list**

```
user_id | date_from
1       | 2021-03-08
```

---

### Вывести уникальные комбинации пользователя и SKU до бана

```sql
SELECT DISTINCT
    u.firstname,
    u.lastname,
    p.sku
FROM users AS u
JOIN purchase AS p ON p.user_id = u.id
LEFT JOIN ban_list AS b ON b.user_id = u.id
WHERE b.date_from > p.date OR b.date_from IS NULL
ORDER BY
    u.firstname,
    u.lastname,
    p.sku;
```

---

### Пользователи, совершившие покупок > 5000р

```sql
SELECT u.id, u.firstname, u.lastname, SUM(p.price)
FROM users AS u
JOIN purchase AS p ON p.user_id = u.id
GROUP BY u.id
HAVING SUM(p.price) > 9000
ORDER BY SUM(p.price) DESC;
```

---

## Модель библиотеки

Три сущности: Автор, Книга, Читатель.  
Книга физически одна и может быть у одного читателя.

### Таблицы

```sql
CREATE TABLE author (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE reader (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    reader_id INTEGER REFERENCES reader(id)
);

CREATE TABLE author_book(
    book_id INTEGER NOT NULL REFERENCES book(id),
    author_id INTEGER NOT NULL REFERENCES author(id),
    CONSTRAINT author_book_pk PRIMARY KEY (book_id, author_id)
);
```

### Книги "на руках"

```sql
SELECT book.name
FROM book
WHERE book.reader_id IS NOT NULL;
```

### Авторы, у которых > 3 книг

```sql
SELECT author.name
FROM author
JOIN author_book ON author.id = author_book.author_id
GROUP BY author.id
HAVING COUNT(author_book.author_id) > 3
ORDER BY COUNT(author_book.author_id) DESC;
```

### Топ-3 читаемых автора

```sql
SELECT author.name, COUNT(book.reader_id) AS book_count
FROM book
JOIN author_book ON book.id = author_book.book_id
JOIN author ON author_book.author_id = author.id
WHERE book.reader_id IS NOT NULL
GROUP BY author.id
ORDER BY book_count DESC
LIMIT 3;
```

---

## Индекс для

```sql
SELECT *
FROM employee
WHERE sex = 'm' AND salary > 300000 AND age = 20
ORDER BY created_at;
```

### Оптимальный индекс

```sql
CREATE INDEX ON employee USING btree (salary, age);
```
