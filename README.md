# Сборка
Бекенд написан на Go 1.17 и для сборки потребуется [Go SDK](https://go.dev/dl/). \
Также для сборки фронтенда потребуется [Node.JS](https://nodejs.org/en/download/). \
Затем, удостоверившись, что всё встало, следует из корня проекта выполнить следующие команды:
```
#Установить/обновить зависимости для бека
go mod download

#Собрать бекенд
go build ./cmd/adapteris -o ./build/adapteris
```
```
#Установить/обновить зависимости для фронта
yarn --cwd ./front install

#Собрать фронтенд
yarn --cwd ./front build
```
Или, если вы не используете прекрасный и изящный [yarn](https://yarnpkg.com/getting-started/install), то для фронта:
```
npm install --prefix ./front 
npm --prefix ./front build
```
Затем
```
#Переместить собранные файлы фронта поближе к беку
mv ./front/dist ./build/static
```
И после этого можно запустить получившийся бинарник
```
./build/adapteris
```
# Разработка
Для запуска дев-версии надо так же установить Go 1.17 и Node.js, установить зависимости фронтенда и бекенда:
```
go mod download
yarn --cwd ./front install
```
А затем 
```
#Запустить бекенд сервер
go run ./cmd/adapteris

#Запустить фронтенд сервер
yarn --cwd ./front start
```
Как можно заметить, в дев режиме фронтенд пересобирается при каждом обновлении файла. Мне также нравится, когда бекенд делает так же, для этого я использую [air](https://github.com/cosmtrek/air#installation). 
```
#Вместо go run
air 
```