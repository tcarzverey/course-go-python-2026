# HW 1
## Дедлайн: 14.10.2025

Задания нацелены на проверку знаний самого Go, основных структур данных языка и concurrency.

Для всех заданий есть публичный набор тестов, лежащий в этом репозитории, а так же приватные тесты, запускаемые
отдельно.

Также решения дополнительно проверяются на предмет плагиата, и, среди сдавших на хороший результат, могут быть выбраны N
случайных человек, которые вызываются на защиту ДЗ: буквально 5-10 минут онлайн созвона с уточняюшими вопросами по
решениям.


## Task 0: Подготовка окружения

1. Установите последнюю версию Go: https://go.dev/doc/install
2. (Опционально) Попробуйте создать свой модуль и поработать с командой go: https://go.dev/doc/tutorial/getting-started
3. (Неактуально для GitHub Classroom) Склонируйте этот репозиторий 
   ```bash
   git clone git@github.com:tcarzverey/course-go-python.git # если клонировать осноной репозиторий
   # или свой форк
   git clone git@github.com:<youraccount>/course-go-python.git
   cd course-go-python
   ```
4. Настройте свою IDE
    1. Goland:
       [инструкция](https://www.jetbrains.com/help/go/installing-and-configuring-goland.html#setting-up-your-work-environment)
    2. VSCode: [инструкция](https://github.com/golang/vscode-go/tree/master?tab=readme-ov-file#quick-start)
5. Подтяните все необходимые зависимости
   ```bash
   go mod tidy
   ```
6. Запустите какой-нибудь из примеров в этом репозитории, чтобы убедиться что все работает
   ```bash
   go run ./examples/lecture2/cli/args arg1 arg2
   ```

Основные команды которыми вы будете пользоваться:

1. Запуск пакета/файла main
   ```bash
   go run path/to/your/package
   go run path/to/your/package/main_file.go
   go run ./ # запустить main файл в текущей директории
   ```
   Чтобы команда сработала нужно чтобы в пакете/файле было `package main` и `func main()`
2. Сборка и запуск бинарного файла
   ```bash
   go build path/to/your/package # бинарник будет назван так же как директория, т.е. package
   go build -o my_command path/to/your/package # бинарник будет назван output_name
   ./my_command arg1 arg2 # запуск бинарник с аргументами
   ```
3. Запуск тестов для пакета
   ```bash
   go test ./path/to/your/package
   go test ./path/to/your/package -v # -v добавляет подробный вывод теста
   go test ./... # запустить все тесты во всех вложенных пакетах
   ```
4. Полезные инструменты и утилиты
   ```bash
   go fmt ./... # отформатировать весь код в проекте
   go vet ./... # проверить код на возможные ошибки
   go install golang.org/x/tools/cmd/goimports@latest # установить утилиту goimports (очень полезная для автоматического импорта)
   # Добавьте в PATH ваш GOBIN, чтобы команды, установленные через install, были доступны как в примере дальше
   goimports -w . # автоматически импортировать все что нужно и удалить неиспользуемые импорты, а еще сделать gofmt
   ```

## Task 1: Dice Roller (3 балла)

[Описание](./dice/README.md)

## Task 2: Error Handler (3 балла)

[Описание](./handler/README.md)

## Task 3: URL Aggregator (4 балла)

[Описание](./urls/README.md)

## Дополнительно после выполнения

1. Проверьте свои решения на предмет самых частых ошибок: https://go.dev/wiki/CodeReviewComments
2. Убедитесь что тесты, которые добавлены к задачам проходят: `go test ./homeworks/hw1/...` 
3. Перечитайте README.md файлы в задачах, убедитесь что не забыли каких-то сценариев и корнеркейсов, т.к. 
   при проверке будет запускаться дополнительные тесты.
4. Прогоните форматтеры `gofmt` и `goimports`