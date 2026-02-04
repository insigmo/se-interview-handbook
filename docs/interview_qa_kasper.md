# Типовые вопросы с собеседований на QA (Собес проходил в Касперский)

# Типовые вопросы с собеседований

---

## 🌐 Сеть

<details>
<summary><b>DHCP</b></summary>

> **DHCP** — протокол автоматической выдачи IP-адреса и сетевых параметров клиенту.
</details>

<details>
<summary><b>Что происходит, когда компьютер включают в сеть (DHCP сервер)</b></summary>

> Процесс называется **DORA**:
> 1. Discover — клиент ищет DHCP-сервер
> 2. Offer — сервер предлагает IP
> 3. Request — клиент запрашивает IP
> 4. Acknowledge — сервер подтверждает
</details>

<details>
<summary><b>DNS</b></summary>

> **DNS** — система преобразования доменных имен в IP-адреса.
</details>

<details>
<summary><b>Что происходит, когда DNS сервер не знает домен (например ya.ru)</b></summary>

> DNS идет по иерархии:
> - root DNS
> - TLD (.ru)
> - авторитетный DNS  
> Если домен не найден — возвращается **NXDOMAIN**.
> - своими словами - идет общение между днс серверами - аля рукопожатие
</details>

<details>
<summary><b>Где больше IP адресов: /23 или /24. Сколько IP в /24 маске. Почему нельзя использовать первый и последний IP</b></summary>

> В **/23** больше IP, чем в /24.  
> **/24 = 256 IP**, из них **254 usable**.
>
> Первый IP — адрес сети  
> Последний IP — broadcast
</details>

---

## 🐧 Администрирование Linux

<details>
<summary><b>Где находятся логи сервера</b></summary>

> Основные логи находятся в каталоге:
> ```
> /var/log/
> ```
</details>

<details>
<summary><b>Права файла: rwx, 4 2 1, как выдать, как поменять владельца, как выдать на каталог, почему r это 4 w это 2 а x это 1 </b></summary>

> r = 4, w = 2, x = 1  
> Пример:
> ```
> chmod 755 file
> ```
>
> Сменить владельца:
> ```
> chown user:group file
> ```
>
> Рекурсивно для каталога:
> ```
> chmod -R 755 dir
> chown -R user:group dir
> ```
> 
> r = 4, w = 2, x = 1 (что означают цифры)
> ```
> r (read) - чтение = 100 (двоичное) = 4 (восьмеричное) 
> w (write) - запись = 010 (двоичное) = 2 (восьмеричное)
> x (execute) - выполнение = 001 (двоичное) = 1 (восьмеричное)
</details>

<details>
<summary><b>Как подключиться по ssh к серверу</b></summary>

> ```
> ssh user@ip
> ```
</details>

<details>
<summary><b>Как скопировать файл с сервера</b></summary>

> ```
> scp user@ip:/path/file .
> ```
</details>

<details>
<summary><b>Как в логах вывести все строки со словом error</b></summary>

> ```
> grep error logfile
> grep -i error logfile
> ```
</details>

<details>
<summary><b>Как смотреть нагрузку сервера. Почему процесс может занимать 130% CPU</b></summary>

> Команды:
> ```
> top
> htop
> ```
>
> 130% CPU возможно, потому что процесс использует несколько ядер.
> 100% = одно ядро.
</details>

<details>
<summary><b>Как закрыть процесс. Он не закрывается — как закрыть точно.  </b></summary>

> Обычное завершение:
> ```
> kill PID
> ```
>
> Принудительное:
> ```
> kill -9 PID
> ```
</details>

<details>
<summary><b> в команде kill -9 PID что значит -9 </b></summary>

> В команде kill -9 PID параметр -9 — это номер сигнала. Это эквивалентно сигналу SIGKILL.
> 
>
> Другие часто используемые сигналы:
> ```
>  Сигнал	Номер	Описание
>  SIGHUP 	1	"Hang up" — перечитать конфигурационные файлы
>  SIGINT 	2	Прерывание (Ctrl+C)
>  SIGQUIT	3	Выход с созданием core dump
>  SIGKILL	9	Принудительное убийство (нельзя перехватить)
>  SIGTERM	15	Корректное завершение (по умолчанию)
>  SIGSTOP	19	Приостановка процесса (нельзя перехватить)
>  SIGCONT	18	Продолжение приостановленного процесса


</details>


---

## 🪟 Администрирование Windows

<details>
<summary><b>Как проходит загрузка системы</b></summary>

> BIOS/UEFI → Boot Manager → загрузка ядра Windows → службы → пользователь.
</details>

<details>
<summary><b>Чем отличаются BIOS и UEFI</b></summary>

> BIOS — устаревший, MBR  
> UEFI — современный, GPT, Secure Boot
</details>

<details>
<summary><b>Что такое базовый диск</b></summary>

> Базовый диск — стандартный тип диска в Windows с обычными разделами.
</details>

<details>
<summary><b>Как посмотреть все логи в Windows</b></summary>

> **Event Viewer (Просмотр событий)**:
> - System
> - Application
> - Security
</details>

<details>
<summary><b>Что такое домен контроллер</b></summary>

> Сервер, который управляет аутентификацией пользователей в домене.
</details>

<details>
<summary><b>Что такое Active Directory</b></summary>

> Служба Microsoft для управления пользователями, компьютерами и политиками.
</details>

---

## 🌍 Web

<details>
<summary><b>DevTools</b></summary>

> Инструменты браузера для отладки web-приложений (Network, Console, Elements).
</details>

<details>
<summary><b>HTTP коды (1xx, 2xx, 3xx, 4xx, 5xx)</b></summary>

> 1xx — информация  
> 2xx — успех  
> 3xx — редирект  
> 4xx — ошибка клиента  
> 5xx — ошибка сервера
</details>

<details>
<summary><b>Что такое WebSocket</b></summary>

> Протокол для постоянного двустороннего соединения клиент ↔ сервер.
</details>

<details>
<summary><b>Как тестировать web на разных устройствах</b></summary>

> - DevTools (responsive mode)
> - разные браузеры
> - реальные устройства
</details>
