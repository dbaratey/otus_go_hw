## Домашнее задание №12 «Заготовка сервиса Календарь»
Необходимо реализовать скелет сервиса «Календарь», который будет дорабатываться в дальнейшем.

Описание то, к чему мы должны придти, представлено [в техническом задании](./CALENDAR.MD).

На данный момент сервис будет состоять из следующих логически выделенных частей:

### 0) «Точка входа», запускающая сервис
Обычно располагается в `cmd/`. В `main()` происходит инициализация компонентов сервиса
(клиент к хранилищу, логгер, конфигурация и пр.), конструирование главного объекта-сервиса на их
основе и его запуск. Допускается использование https://github.com/spf13/cobra.

### 1) Конфигурирование сервиса
При необходимости создать отдельный пакет, отвечающий за работу с конфигом.

Сервис должен иметь конфиг, считываемый из файла. Для этого потребуется:
* обработка аргументов командной строки;
* чтение файла конфигурации формата json, yaml и пр. на выбор разработчика;
* заполнение программной структуры конфига из файла (Unmarshal и пр. способы);
* допустимо, если все действия выше за вас делает сторонний модуль.

Конфигурация понадобится для инициализации различных компонентов системы,
её возможные поля будут рассмотрены далее.

Соответственно сервис будет запускаться командой вида
```bash
./calendar --config=/path/to/config.yaml
```
где `--config`  - путь к файлу конфигурации.

В репозитории должен присутствовать образец конфига.

### 2) Логирование в сервисе
При необходимости создать отдельный пакет, отвечающий за работу с логером.

Параметры, которыми необходимо инициализировать логгер:
* log_file - путь к файлу логов;
* log_level - уровень логирования (error / warn / info / debug);
* пр. на усмотрение разработчика.

Логгер может быть как глобальной переменной, так и компонентом сервиса.

### 3) Работа с хранилищем
Создать отдельный пакет, отвечающий за работу с хранилищем.

Создать интерфейс хранилища событий, состоящий из методов для работы с ним:
* добавление события в хранилище;
* изменение события в хранилище;
* удаление события из хранилища;
* листинг событий;
* пр. на усмотрение разработчика.

Описание сущности и методов представлено в [ТЗ](./CALENDAR.MD).

Создать объекты ошибок, соответствующие бизнес ошибкам, которые необходимо выделить.
Например, `ErrDateBusy` - данное время уже занято другим событием.

Создать две реализации интерфейса выше:
* **in-memory**: храним события в памяти (т.е. просто складываем объекты в словари/слайсы, не забывая про критические секции);
* **sql**: храним события в полноценной СУБД путем использования SQL-запросов в соответствующих методах.

Вынести в конфиг параметр, отвечающий за то, какую из реализаций использовать при старте.

Для работоспособности второй реализации необходимо:
* установить СУБД (например PostgreSQL) локально (или сразу через Docker, если знаете как);
* создать базу данных и пользователя для проекта календарь;
* реализовать схему данных (таблицы, индексы) в виде отдельного SQL или go-файла (файл миграции)
и сохранить его в репозиторий;
* применять миграции руками или на старте сервиса;
* вынести настройки подключения к БД в конфиг проекта.

Полезные библиотеки:
* https://github.com/jmoiron/sqlx
* https://github.com/pressly/goose#go-migrations

**Использовать ORM (например, gorm) не допускается**.

Календарь должен использовать хранилище через интерфейс.

### 4) Запуск простого HTTP-сервера
Запуск календаря должен стартовать HTTP-сервер. `host` и `port` сервера вынести в конфиг.

На данном этапе сервер не должен быть связан с бизнес логикой приложения и должен иметь
только один "hello-world" endpoint ("/", "/hello", etc.).

Информация об обработанном запросе должна выводиться в log-файл:
* IP клиента;
* дата и время запроса;
* метод, path и версия HTTP;
* код ответа;
* latency (время обработки запроса, посчитанное, например, с помощью middleware);
* user agent, если есть.

Пример лога:
```text
66.249.65.3 [25/Feb/2020:19:11:24 +0600] GET /hello?q=1 HTTP/1.1 200 30 "Mozilla/5.0"
```

### 5) Юнит-тесты
Минимальный обязательный набор - тесты на **in-memory** реализацию хранилища (на основную логику, бизнес-ошибки и конкуррентно-безопасность).

Остальные тесты на усмотрение разработчика.

### 6) Makefile
Проект должен иметь в корне файлы go.mod и Makefile, последний должен описывать команды:
* `make build` - скомпилировать бинарный файл сервиса;
* [`make run`] - опционально, собрать и запустить сервис с конфигом по умолчанию;
* `make test` - запустить юнит-тесты (с флагом -race);
* `make lint` - запустить golangci-lint (при необходимости добавить свой `.golangci.yml`);
* [`make migrate`] - опционально, если миграции применяются руками;
* пр. на усмотрение разработчика

### Об архитектуре
Проект должен следовать:
* https://github.com/golang-standards/project-layout
* https://golang.org/doc/effective_go.html#package-names
* https://rakyll.org/style-packages/
* https://en.wikipedia.org/wiki/Dependency_injection

Важно понять, что в Go нет серебряной пули по архитектуре.
Ссылки ниже могут дать полезные концепции, но не стоит слепо следовать им:
* https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
* https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html
* https://github.com/marcusolsson/goddd
* чистая архитектура (clean architecture):
    - https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
    - https://medium.com/@zhashkevych/чистая-архитектура-на-golang-cccbfdc95eba

Используйте стандартный layout, соблюдайте направление зависимостей, выделяйте пакеты по концепции,
не забывайте про внедрение зависимостей через интерфейсы и у вас всё получится!

Не забываем про стайл гайд, например:
* https://github.com/uber-go/guide/blob/master/style.md
* https://github.com/cristaloleg/go-advice

### В данном ДЗ не нужно
* Реализовывать HTTP, GRPC и пр. API к микросервису.
* Писать .proto-файлы.

Это всё будет позже.

### Критерии оценки
- Makefile заполнен и пайплайн зеленый - 1 балл
- Понятность и чистота кода (включая факт, что проект разбит
на пакеты по определенной логике) - до 2 баллов
- Реализовано конфигурирование сервиса - 1 балл
- Используется логгер и он настраивается из конфига - 1 балл
- Реализовано хранилище:
    - in-memory - 1 балл
    - sql + миграции - 2 балла
- Запускается простой HTTP-сервер - 1 балл
- Присутствуют юнит-тесты - 1 балл

#### Зачёт от 7 баллов