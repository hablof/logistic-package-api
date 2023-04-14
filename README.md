Проект выполнен в рамках курса [Route256 2021 OZON](https://rutracker.org/forum/viewtopic.php?t=6201055) и является форком репозиториев [раз](https://github.com/ozonmp/omp-demo-api) и [два](https://github.com/ozonmp/omp-template-api).

### Репозиторий с заданиями
https://github.com/ozonmp/omp-docs

# О проекте
Проект включает в себя помимо настоящего репозиториия, ещё два:
* https://github.com/hablof/omp-bot — телеграм бот
* https://github.com/hablof/logistic-package-api-facade — фасад

## gRPC-server
Первый из двух бинарников, собирающихся из данного репозитория — gRPC-server. 

### API
Protopuf контракт API описан в `api\hablof\logistic_package_api\v1\logistic_package_api.proto`, предоставляет CRUD-методы. Код grpc объектов, методов, валидации вынесен в отдельный модуль `pkg\logistic-package-api`.

В обработчиках запросов есть возможность поднять уровень логгирования с помощью метаданных. Для этого необходимо передать  по ключу `"log_level"` значение `"debug"`.

### Service
Слой сервиса не содержит логики домена и создан исключетельно в целях разделения пакетов `api` и `repo` (отсутсвуют импорты).

### Repository
Данные хранятся в Postgres. Методы пакета `repo` повторяют методы `api`, однако, при этом, реализуют паттерн [transactional outbox](https://microservices.io/patterns/data/transactional-outbox.html). Помимо изменения основной таблицы, хранящей записи о сущностях домена `package`, при выполнении CUD-методов добавляется запись в таблицу `package_event`, описывающая произошедшие изменения.

Текст SQL-запросов собирается с помощью [squirell](https://github.com/Masterminds/squirrel)

#### Миграции
Миграции оприсаны в `db\migrations`.

Индексы созданы:
* на столбце `package_id` в таблице `package`, поскольку по нему идёт условие `WHERE` в запросе `Describe`
* на столбце `package_event_id` в таблице `package_event`, поскольку по нему идёт условие `WHERE` в запросах Cleaner'а (см. ниже)
* на столбце `event_status` в таблице `package_event`, поскольку по нему идёт условие `WHERE` в запросах Consumer'а (см. ниже)

## Retranslator
Второй бинарник — retranslator. Работает в асинхронном режиме. Состоит из:
* `consumer`
* `producer`
* `sender`
* `cleaner`

Термины `consumer` и `producer` употреблены по отношению к внешнему миру.

### Consumer
По таймеру горутины `consumer`'ов "забирают" из таблицы `package_event` данные о событии репозитория на обработку: читают записи и изменяют значение в столбце `event_status` на `Locked`. Полученные данные отправляются в буферезированный канал `eventsChannel`.

### Producer
Горутины `producer`'ов читают данные из `eventsChannel` и пытаются отправить их в кафка-топик вызывая методы `sender`'а. Пишут в буферезированный канал `cleanerChannel` сообщение с `id` события и результатом отправи в кафку.

### Cleaner
Если отправка в кафку была успешна, `cleaner` удаляет из репозитория запись о событии, если неуспешна — `cleaner` возвращает значение в столбце `event_status` на значение по умолчанию `Unlocked`.  Запись события готова к повторной обработке.

Cleaner пытается набрать батч событий и обработать их в базе одним запросом. Если батч не накапливается, обработка вызывается прниудительно по таймеру.

### Sender 
Sender сериализует данные о событиях в protobuf и отправляет в топик `omp-package-events`. 
Предусмотренны повторные попытки отправки с некоторым интервалом.

Код модели вынесен в отдельный модуль `pkg\kafka-proto`.

### At-least-once
Порядок действий при завершении работы ретранслятора позволяет избежать потерь сообщений.

## Метрики
Сервисы собирают метрики с помощью [прометей-клиента](https://github.com/prometheus/client_golang)

### gRPC-server
grpc-сервер собирает две метрики:
* `logistic_package_api_not_found_total` _Counter_ — общее количество NotFound событий
* `logistic_package_api_cud_event_total` _Counter_ — общее количество CUD событий

Метрики grpc-сервера доступны на `:9100/metrics`

### Retranslator

Ретранслятор собирает одну метрику
* `logistic_package_api_retranslator_events_processing` _Gauge_ — количество событий, обрабатывемых в текущий момент в ретрансляторе

Метрики ретранслятора доступны на `:9101/metrics`

### Grafana и Prometheus

С заданым интервалои Prometheus считывает метрики сервисов. Метрики запрашиваются Grafan'ой и отображаются на графике: `logistic_package_api_retranslator_events_processing` — напрямую, а `logistic_package_api_not_found_total` и `logistic_package_api_cud_event_total` — расчитывая количество происходящих событий в секунду в окне в 1 минуту.

Grafana доступна на `:3000`

## Jaeger

Обработчики запросов grpc-серверa записывает трейсы и отправляет их в Jaeger. Трейсы имеют вложенные спаны методов репозиторя и SQL-запросов.

Jaeger доступен на `:16686`

## Docker

Описаны докерфайлы для образов grpc-серверa и ретранслятора.

В docker-compose прокидываются конфиги для gRPC-сервера, ретранслятора, бота и фасада.

## Makefile

В мейкфайле описаны команды для локального запуска приложения, для генерации protobuf модулей с помощью утилиты buf, сборки докер-образов, и их запуска.