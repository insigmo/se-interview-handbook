# Хакаем собесы на бэк и Go

## Общая информация

Цифры в скобка после вопроса \- сколько раз вопрос встретился на разных собесах.  
Пиши в чат, если хочешь спросить или порекомендовать источник по вопросам.  
Документ активно пополняется источниками.  
Просьба не распространять документ за пределами сообщества. 

## Список собесов

#### [Скрининг в авито](https://www.youtube.com/watch?v=yKyg_je6vTA) (прошел дальше)

#### [Собес в авито](https://youtu.be/alOrzdaWYyA) (оффер на 350к, оценили в middle+/senior-)

#### [Собес в касперский](https://youtu.be/qPu1yk-1tzk) (оффер на 300к)

#### [Собес в yabbi](https://youtu.be/982mC6NjSoc) (сразу позвали знакомиться с командой, скипнул из\-за макс. вилки 250к)

#### [Собес в сбер маркет](https://youtu.be/7efsnTlvMqQ) (позвали знакомиться с командой, грейд middle-/middle \- скипнул)

#### [Скрининг в uzum](https://youtu.be/F3Q06b3FKi4) (прошел дальше, скипнул)

#### [Собес в kode](https://youtu.be/JE0a5p7Jj04) (проигнорили)

#### [System design в авито](https://www.youtube.com/watch?v=ac2_7_uFYT8) (оффер 350к)

#### [какой-то аутсорс](https://www.youtube.com/watch?v=Sr6R-XmO4c4), оффер 250к

#### Короткий стандартный собес в ВБ, оффер на 270к \- запись потерялась видимо. По опыту \- у них стандартного опросника видимо нет, собес сильно зависит от команды.

## Лайвкодинг

#### [Параллельные http запросы (горутины, waitGroup, каналы)](https://www.youtube.com/watch?v=alOrzdaWYyA&t=220s)

#### [На каналы и select](https://www.youtube.com/watch?v=alOrzdaWYyA&t=1890s)

#### [Написать многопоточный логгер (мьютексы, каналы, горутины)](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=1113s)

#### [На WaitGroup и каналы](https://www.youtube.com/watch?v=982mC6NjSoc&t=2330s)

#### [Пайплайн каналов](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=4055s)

#### [Merge каналов](https://www.youtube.com/watch?v=F3Q06b3FKi4&t=560s)

#### [Тротлер (задача на каналы)](https://www.youtube.com/watch?v=F3Q06b3FKi4&t=1200s)

#### [Семафор](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=1555s)

#### [SQL запрос на group by и having](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=2005s)

#### Источники

На гошном лайвкодинге любят многопоточку, типичные паттерны описаны в  
[https://go.dev/blog/pipelines](https://go.dev/blog/pipelines)  

## GO

### Слайсы

#### Как устроен слайс (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1780s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1780s) 

#### Чем слайс отличается от массива (2)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=20s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=20s)  
[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1780s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1780s) 

#### Как конвертировать массив в слайс (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1885s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1885s) 

#### Задача на append к слайсу (2)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=1795s](https://www.youtube.com/watch?v=982mC6NjSoc&t=1795s)  
[https://www.youtube.com/watch?v=F3Q06b3FKi4\&t=325s](https://www.youtube.com/watch?v=F3Q06b3FKi4&t=325s)

#### Источники

Все что нужно знать от авторов языка:  
[https://go.dev/blog/slices-intro](https://go.dev/blog/slices-intro)   
[https://go.dev/blog/slices](https://go.dev/blog/slices) 

### Map

#### Как устроена (2)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2040s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2040s) 

#### Что будет если писать конкурентно в мапу (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2260s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2260s) 

#### Сложность доступа к мапе (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=40s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=40s)

#### Сет (множество) в го (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=52s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=52s) 

#### Как устроена sync.Map (2)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2310s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2310s)

#### Источники

[https://go.dev/blog/maps](https://go.dev/blog/maps)  
[https://www.youtube.com/watch?v=P\_SXTUiA-9Y](https://www.youtube.com/watch?v=P_SXTUiA-9Y)

[https://pkg.go.dev/sync\#Map](https://pkg.go.dev/sync#Map)   
[https://golangforall.com/ru/post/sync-map.html](https://golangforall.com/ru/post/sync-map.html)  
[https://habr.com/ru/articles/338718/](https://habr.com/ru/articles/338718/) 

### Горутины

#### Что такое горутина, как устроены в рантайме, преимущество перед тредами (4)

[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=820s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=820s)   
[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=865s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=865s)  
[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2530s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2530s)  
[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=2168s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=2168s) 

#### Чем отличается конкурентность и параллелизм (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2493s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2493s) 

#### Сколько по дефолту будет потоков (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2635s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2635s)

#### Источники 

Вся серия статей  
[https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)  
[https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)  
[https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html](https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html)

[https://www.kelche.co/blog/go/golang-scheduling/](https://www.kelche.co/blog/go/golang-scheduling/) 

### Интерфейсы

#### Что такое интерфейсы (2)

[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=2755s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=2755s)

#### Сравнение интерфейса с nil, внутреннее устройство интерфейса (1)

[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=770s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=770s)   
[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2430s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2430s) (todo: плохой ответ, посоветуйте источник)

#### Задача на интерфейсы и Error, errors.As errors.Is (1)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=1940s](https://www.youtube.com/watch?v=982mC6NjSoc&t=1940s) 

#### Источники 

[https://go.dev/tour/methods/9](https://go.dev/tour/methods/9)  
[https://go.dev/blog/go1.13-errors](https://go.dev/blog/go1.13-errors) 

### Каналы

#### Какие бывают каналы, чем отличается буферизированный от небуферизированной канала (4)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=85s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=85s)  
[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=1756s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=1756s)  
[https://www.youtube.com/watch?v=982mC6NjSoc\&t=2700s](https://www.youtube.com/watch?v=982mC6NjSoc&t=2700s)   
[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2885s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2885s) 

#### Внутреннее устройство каналов (2) [https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=545s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=545s)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2846s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2846s)   

#### Аксиомы каналов (2)

[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=1800s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=1800s) 

#### Код с каналами (много раз) [https://www.youtube.com/watch?v=F3Q06b3FKi4\&t=180s](https://www.youtube.com/watch?v=F3Q06b3FKi4&t=180s)

#### Источники

[https://go.dev/tour/concurrency/2](https://go.dev/tour/concurrency/2)  
[https://dave.cheney.net/2014/03/19/channel-axioms](https://dave.cheney.net/2014/03/19/channel-axioms)  
[https://www.youtube.com/watch?v=ZTJcaP4G4JM](https://www.youtube.com/watch?v=ZTJcaP4G4JM)   
 

### Context

#### Зачем нужен контекст (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=89s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=89s)  
[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=2850s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=2850s) 

#### Источники

[https://go.dev/blog/contex](https://go.dev/blog/context)[t](https://pkg.go.dev/context)  

### Defer, recover, panic

#### Исключения в го, их аналог (2)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=120s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=120s)  
[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=3105s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=3105s)  

#### Как работает defer (2)

[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=660s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=660s)   
[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=3060s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=3060s) 

#### Можно ли поймать панику другой горутины (2)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=3133s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=3133s) 

#### Код с defer (2)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=3350s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=3350s)  
[https://www.youtube.com/watch?v=F3Q06b3FKi4\&t=85s](https://www.youtube.com/watch?v=F3Q06b3FKi4&t=85s)  

#### Источники

[https://go.dev/blog/defer-panic-and-recover](https://go.dev/blog/defer-panic-and-recover) 

### string и \[\]byte

#### Как строки представлены в go (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1600s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1600s) 

#### Задача на string и \[\]byte (1)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=1649s](https://www.youtube.com/watch?v=982mC6NjSoc&t=1649s) 

#### Alias чего является byte и rune (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1575s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1575s) 

#### Что будет, если на строку вызвать len (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1650s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1650s) 

#### Как правильно посчитать число символов (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1695s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1695s) 

#### Почему плохо кастовать из string в \[\]byte (1)

#### Источники

[https://go.dev/blog/strings](https://go.dev/blog/strings) 

### Захват переменной в цикле (2)

[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=1390s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=1390s) 

### Sync, Мьютексы

#### Что такое мьютекс, их виды (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=140s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=140s) 

#### Задача на мьютекс и atomic (1)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=2255s](https://www.youtube.com/watch?v=982mC6NjSoc&t=2255s) 

#### Какие бывают примитивы синхронизации (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=3205s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=3205s) 

#### Как устроен мьютекс (1) todo: плохой ответ, посоветуйте источник

#### Код на мьютекс (много, см. лайвкодинг)

[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=1260s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=1260s)

#### Источники

[https://go.dev/tour/concurrency/9](https://go.dev/tour/concurrency/9)   
На мидловые собесы из sync надо знать  
Mutex, Map, RWMutex, WaitGroup  
[https://pkg.go.dev/sync](https://pkg.go.dev/sync) 

### Указатели

#### Receiver (2)

[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=2830s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=2830s)  
[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=229s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=229s)  

#### Задача на указатель (1)

[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=0s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=0s)  
[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=1260s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=1260s) 

#### Источники

[https://go.dev/tour/methods/4](https://go.dev/tour/methods/4) 

### ООП

#### ООП в go (2)

[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=2648s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=2648s)  
[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1170s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1170s)  

### WaitGroup

Очень часто на лайвкодинге, см. лайвкодинг

#### Источники

[https://pkg.go.dev/sync\#example-WaitGroup](https://pkg.go.dev/sync#example-WaitGroup) 

### Select

Очень часто на лайвкодинге, см. лайвкодинг

#### Что такое select и в каком порядке исполняется (3)

[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=370s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=370s)  
[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2923s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2923s) 

#### Источники

[https://go.dev/tour/concurrency/5](https://go.dev/tour/concurrency/5) 

### Прочее

#### Как устроен garbage collector (2)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=2670s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=2670s)  
Источники  
На мидловый собес достаточно начало статьи  
[https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html)  
и мб еще как настраивать GOGC (хотя меня ни разу не спрашивали)  
[https://tip.golang.org/doc/gc-guide](https://tip.golang.org/doc/gc-guide) \- “key takeaway is that doubling GOGC will double heap memory overheads and roughly halve GC CPU cost”

#### Работал ли с профайлером (1)

[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=2625s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=2625s) 

#### Расшифровать solid, dry, kiss, yagni (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1270s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1270s) 

#### Что такое императивный и декларитивный язык (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1435s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1435s) 

#### Какие есть встроенные типы в go (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=1519s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=1519s) 

## БАЗЫ

### Индексы

#### Зачем нужны индексы, примеры индексов, минусы (4)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=180s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=180s)  
[https://www.youtube.com/watch?v=982mC6NjSoc\&t=960s](https://www.youtube.com/watch?v=982mC6NjSoc&t=960s)   
[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=695s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=695s)   
[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=677s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=677s) 

#### Какой индекс для jsonb (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=730s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=730s) 

#### Составные индексы, примеры когда они полезны/бесполезны (1)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=960s](https://www.youtube.com/watch?v=982mC6NjSoc&t=960s) 

#### Почему поиск в b-tree быстрый, сложность поиска (2)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=208s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=208s)  
[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=757s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=757s)  

#### Источники

[https://www.postgresql.org/docs/current/indexes-types.html](https://www.postgresql.org/docs/current/indexes-types.html)  
[https://www.postgresql.org/docs/current/indexes-multicolumn.html](https://www.postgresql.org/docs/current/indexes-multicolumn.html)   
[https://www.postgresql.org/docs/current/using-explain.html](https://www.postgresql.org/docs/current/using-explain.html) 

### Транзакции

#### Зачем нужны транзакции, уровни изоляций, расшифровать ACID (3)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=250s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=250s)  
[https://www.youtube.com/watch?v=982mC6NjSoc\&t=840s](https://www.youtube.com/watch?v=982mC6NjSoc&t=840s)  
[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=560s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=560s)  

#### Источники

[https://www.youtube.com/watch?v=e9a4ESSHQ74](https://www.youtube.com/watch?v=e9a4ESSHQ74) 

### SQL

#### Чем отличается where и having (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=313s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=313s) 

#### Перечислить агрегатные функции (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=353s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=353s) 

#### Что такое триггер и примеры использования (1) todo: плохой ответ, посоветуйте источник

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=370s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=370s) 

#### Зачем нужен foreign key (1)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=1170s](https://www.youtube.com/watch?v=982mC6NjSoc&t=1170s) 

#### Оконные функции (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=510s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=510s) 

#### Рекурсивные запросы, сte (1)  todo: плохой ответ, посоветуйте источник

#### Какие бывают join (1)

#### [https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=640s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=640s) 

### Kafka

#### Основные термины кафки (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=877s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=877s) 

#### Какую библиотеку юзаете для кафки (1)

[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=1092s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=1092s)

#### Источники

[https://www.youtube.com/watch?v=-AZOi3kP9Js](https://www.youtube.com/watch?v=-AZOi3kP9Js)  
[https://www.youtube.com/watch?v=j4bqyAMMb7o\&list=PLa7VYi0yPIH0KbnJQcMv5N9iW8HkZHztH](https://www.youtube.com/watch?v=j4bqyAMMb7o&list=PLa7VYi0yPIH0KbnJQcMv5N9iW8HkZHztH) (первые 11 видосов, они маленькие)

### Базы прочее

#### С какими базами работал, на каком уровне, какие бывают (4)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=645s](https://www.youtube.com/watch?v=982mC6NjSoc&t=645s)  
[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=469s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=469s)  
[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=386s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=386s) 

#### Нормализация в БД, зачем нужна (1) (todo: плохой ответ, посоветуйте источник)

[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=445s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=445s) 

#### Почему реляционные базы называются реляционные (1)

[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=414s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=414s) 

#### Настраивал ли постгрес (1)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=700s](https://www.youtube.com/watch?v=982mC6NjSoc&t=700s) 

#### Проектировал ли структуру базы (1)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=725s](https://www.youtube.com/watch?v=982mC6NjSoc&t=725s) 

#### Mongo vs Postgres (1)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=800s](https://www.youtube.com/watch?v=982mC6NjSoc&t=800s) 

#### Был ли опыт с redis, что юзал в нем (2)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=805s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=805s) 

#### Был ли опыт с elastic (1)

## ПРОТОКОЛЫ

#### Рассказать про протоколы межсерверного взаимодействия (1)

[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=1020s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=1020s) 

#### Опыт с вебсокетами (1) (todo: посоветуйте источник)

### TCP

#### TCP vs UDP (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=429s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=429s) 

#### Как устроен TCP (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=429s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=429s) 

### HTTP

#### Из каких частей состоит HTTP запрос (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=502s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=502s) 

#### HTTP vs HTTPS (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=549s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=549s) 

#### Зачем нужны таймаут и как их подобрать (1) todo: плохой ответ, посоветуйте источник

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=600s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=600s) 

### DNS

#### Что происходить когда вбиваешь google.com в браузере (1)

[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=3780s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=3780s) 

#### По какому протоколу работает DNS (1)

### GRPC

#### grpc: чем stream отличается от unary (1)

[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=1125s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=1125s) 

### Источники

Брошюра [Introduction to Networking Charles Severance](https://do1.dr-chuck.net/net-intro/EN_us/net-intro.pdf)

## ОС

#### Отличие процессов и потоков в ОС (2)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=696s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=696s)  
[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=988s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=988s) 

#### Как убить процесс в линуксе (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=732s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=732s) 

#### Что происходит, когда в коде просим выделить один кб памяти (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=832s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=832s) 

#### Что будет, если система не может дать память (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=902s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=902s) 

#### Дескрипторы ОС (1)

[https://www.youtube.com/watch?v=alOrzdaWYyA\&t=1090s](https://www.youtube.com/watch?v=alOrzdaWYyA&t=1090s) 

#### Куда смотреть если лагает сервер (1) todo: не очень ответ, посоветуйте источник

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=435s](https://www.youtube.com/watch?v=982mC6NjSoc&t=435s) 

#### Опыт с ci/cd (1)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=565s](https://www.youtube.com/watch?v=982mC6NjSoc&t=565s) 

#### Опыт с мониторингом и prometheus (1)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=590s](https://www.youtube.com/watch?v=982mC6NjSoc&t=590s) 

## Архитектура

#### Зачем нужно кэширование, плюсы/минусы, алгоритмы обновления (1)

[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=2859s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=2859s) 

#### Что такое шардирование и реплицирование (2)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=1250s](https://www.youtube.com/watch?v=982mC6NjSoc&t=1250s) 

#### Как масштабировать приложение, вертикальное/горизонтальное масштабирование (1)

[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=3223s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=3223s) 

#### Как масштабировать базу: синхронная репликация, шардирование (2)

[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=3380s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=3380s) 

#### Как работает репликация (2)

[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=3490s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=3490s) 

#### Балансировка нагрузки (1)

[https://www.youtube.com/watch?v=qPu1yk-1tzk\&t=3588s](https://www.youtube.com/watch?v=qPu1yk-1tzk&t=3588s) 

#### Что-то хотели спросить про кубер (1)

[https://www.youtube.com/watch?v=yKyg\_je6vTA\&t=948s](https://www.youtube.com/watch?v=yKyg_je6vTA&t=948s) 

#### Зачем нужна микросервисная архитектура (1)

[https://www.youtube.com/watch?v=7efsnTlvMqQ\&t=340s](https://www.youtube.com/watch?v=7efsnTlvMqQ&t=340s) 

#### Рассказать все про авторизацию, jwt, oauth (1)

[https://www.youtube.com/watch?v=JE0a5p7Jj04\&t=880s](https://www.youtube.com/watch?v=JE0a5p7Jj04&t=880s) 

## Git

#### Как работаем с гитом в команде (1) 

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=500s](https://www.youtube.com/watch?v=982mC6NjSoc&t=500s) 

#### Чем отличается merge от rebase (2)

[https://www.youtube.com/watch?v=982mC6NjSoc\&t=545s](https://www.youtube.com/watch?v=982mC6NjSoc&t=545s) 

