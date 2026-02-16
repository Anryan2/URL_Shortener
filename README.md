# URL_Shortner
**URL_Shortner** - это сервис, который должен предоставлять API по созданию
сокращенных ссылок.

### Функциональность проекта:

* Генерация коротких ссылок для длинных URL.

* Перенаправление по короткому URL на оригинальную ссылку.

* Поддержка хранения ссылок в памяти приложения или базе данных PostgreSQL.

* Запуск через Docker.

* Реализованный функционал покрыт Unit-тестами.

### Запуск проекта
Требования:

Установлен *Docker* и (при использовании базы данных) *PostgreSQL*.

```
docker build -t url .
```

#### Запуск сервера с хранилищем в памяти

```
docker run -d -p 8080:8080 -e STORAGE_TYPE=memory url
```

#### Запуск сервера с PostgreSQL (замените строку подключения к базе данных)

```
docker run -d -p 8080:8080 -p 5432:5432 -e STORAGE_TYPE=postgres -e DB_CONN_STR=postgres://postgres:12345@host.docker.internal/postgres?sslmode=disable url
```

### API эндпоинты

#### Создание короткой ссылки (POST /shorten)
**Запуск:**
```
curl.exe -X POST "localhost:8080/shorten" -d  '{\"url\":\"youtu.be\"}'
```

**Ответ:**
```
{"short_url":"http://localhost:8080/ZFsTPc0LDK"}
```

#### Редирект по короткой ссылке (GET /)
**Запуск:**
```
curl.exe -X GET "localhost:8080/ZFsTPc0LDK"
```

**Ответ:**
```
<a href="/youtu.be">Found</a>.
```

**Либо можно перейти по ссылке localhost:8080/ZFsTPc0LDK**

**Ответ:**
```
youtu.be.
```


