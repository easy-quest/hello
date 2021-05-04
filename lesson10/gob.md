-   Установка Golang
    -   Установка Go на Windows
    -   Установка Go на Ubuntu
    -   Установка Go на Debian
    -   Установка Go на MacOS
    -   Установка Go на CentOS
-   Курс Go
-   Создание сайта
-   Строки
-   Время
-   Числа

Golang

-   Установка Golang
    -   Установка Go на Windows
    -   Установка Go на Ubuntu
    -   Установка Go на Debian
    -   Установка Go на MacOS
    -   Установка Go на CentOS
-   Курс Go
-   Создание сайта
-   Строки
-   Время
-   Числа

Поиск



СЕРИАЛИЗАЦИЯ ДАННЫХ В БИНАРНЫЙ ФОРМАТ GOB


Помимо хорошо известных JSON и XML, Go также предлагает бинарный формат
— gob. В данной инструкции описывается базовый концепт использования
пакета gob из encoding.


Как сериализовать структуру в бинарный формат gob?

1. Создайте файл gob.go со следующим содержимым:

[]

ПРЕМИУМ 👑 канал по Golang

Рекомендуем вам супер TELEGRAM канал по Golang где собраны все материалы
для качественного изучения языка. Удивите всех своими знаниями на
собеседовании! 😎

Подписаться на канал

Уроки, статьи и Видео

Мы публикуем в паблике ВК и Telegram качественные обучающие материалы
для быстрого изучения Go. Подпишитесь на нас в ВК и в Telegram.
Поддержите сообщество Go программистов.

Go в ВК ЧАТ в Telegram

gob.go

Go

package main import ( "bytes" "encoding/gob" "fmt" ) type User struct {
FirstName string LastName string Age int Active bool } func (u User)
String() string { return
fmt.Sprintf(`{"FirstName":%s,"LastName":%s,"Age":%d,"Active":%v }`,
u.FirstName, u.LastName, u.Age, u.Active) } type SimpleUser struct {
FirstName string LastName string } func (u SimpleUser) String() string {
return fmt.Sprintf(`{"FirstName":%s,"LastName":%s}`, u.FirstName,
u.LastName) } func main() { var buff bytes.Buffer // Кодирование
значения enc := gob.NewEncoder(&buff) user := User{ "Radomir",
"Sohlich", 30, true, } enc.Encode(user) fmt.Printf("%X\n", buff.Bytes())
// Декодирование значения out := User{} dec := gob.NewDecoder(&buff)
dec.Decode(&out) fmt.Println(out.String()) enc.Encode(user) out2 :=
SimpleUser{} dec.Decode(&out2) fmt.Println(out2.String()) }

+-----------------------------------+-----------------------------------+
| 1                                 | package main                      |
|                                   |                                   |
| 2                                 |                                   |
|                                   |                                   |
| 3                                 | import (                          |
|                                   |                                   |
| 4                                 | "bytes"                           |
|                                   |                                   |
| 5                                 | "encoding/gob"                    |
|                                   |                                   |
| 6                                 | "fmt"                             |
|                                   |                                   |
| 7                                 | )                                 |
|                                   |                                   |
| 8                                 |                                   |
|                                   |                                   |
| 9                                 | type User struct {                |
|                                   |                                   |
| 10                                | FirstName string                  |
|                                   |                                   |
| 11                                | LastName  string                  |
|                                   |                                   |
| 12                                | Age       int                     |
|                                   |                                   |
| 13                                | Active    bool                    |
|                                   |                                   |
| 14                                | }                                 |
|                                   |                                   |
| 15                                |                                   |
|                                   |                                   |
| 16                                | func (u User) String() string {   |
|                                   |                                   |
| 17                                | return                            |
|                                   | fmt.Sprintf(`{"FirstName":%s,"Las |
| 18                                | tName":%s,"Age":%d,"Active":%v    |
|                                   | }`,                               |
| 19                                |                                   |
|                                   | u.FirstName, u.LastName, u.Age,   |
| 20                                | u.Active)                         |
|                                   |                                   |
| 21                                | }                                 |
|                                   |                                   |
| 22                                |                                   |
|                                   |                                   |
| 23                                | type SimpleUser struct {          |
|                                   |                                   |
| 24                                | FirstName string                  |
|                                   |                                   |
| 25                                | LastName  string                  |
|                                   |                                   |
| 26                                | }                                 |
|                                   |                                   |
| 27                                |                                   |
|                                   |                                   |
| 28                                | func (u SimpleUser) String()      |
|                                   | string {                          |
| 29                                |                                   |
|                                   | return                            |
| 30                                | fmt.Sprintf(`{"FirstName":%s,"Las |
|                                   | tName":%s}`,                      |
| 31                                |                                   |
|                                   | u.FirstName, u.LastName)          |
| 32                                |                                   |
|                                   | }                                 |
| 33                                |                                   |
|                                   |                                   |
| 34                                |                                   |
|                                   | func main() {                     |
| 35                                |                                   |
|                                   |                                   |
| 36                                |                                   |
|                                   | var buff bytes.Buffer             |
| 37                                |                                   |
|                                   |                                   |
| 38                                |                                   |
|                                   | // Кодирование значения           |
| 39                                |                                   |
|                                   | enc := gob.NewEncoder(&buff)      |
| 40                                |                                   |
|                                   | user := User{                     |
| 41                                |                                   |
|                                   | "Radomir",                        |
| 42                                |                                   |
|                                   | "Sohlich",                        |
| 43                                |                                   |
|                                   | 30,                               |
| 44                                |                                   |
|                                   | true,                             |
| 45                                |                                   |
|                                   | }                                 |
| 46                                |                                   |
|                                   | enc.Encode(user)                  |
| 47                                |                                   |
|                                   | fmt.Printf("%X\n", buff.Bytes())  |
| 48                                |                                   |
|                                   |                                   |
| 49                                |                                   |
|                                   | // Декодирование значения         |
| 50                                |                                   |
|                                   | out := User{}                     |
| 51                                |                                   |
|                                   | dec := gob.NewDecoder(&buff)      |
| 52                                |                                   |
|                                   | dec.Decode(&out)                  |
| 53                                |                                   |
|                                   | fmt.Println(out.String())         |
| 54                                |                                   |
|                                   |                                   |
| 55                                |                                   |
|                                   | enc.Encode(user)                  |
| 56                                |                                   |
|                                   | out2 := SimpleUser{}              |
| 57                                |                                   |
|                                   | dec.Decode(&out2)                 |
|                                   |                                   |
|                                   | fmt.Println(out2.String())        |
|                                   |                                   |
|                                   |                                   |
|                                   |                                   |
|                                   | }                                 |
+-----------------------------------+-----------------------------------+

2. Запустите код через go run gob.go;
3. Посмотрите на результат в терминале:

Shell

40FF81030101045573657201FF82000104010946697273744E616D65010C0001084C6173744E616D65010C0001034167650104000106416374697665010200000019FF8201075261646F6D69720107536F686C696368013C010100
{"FirstName":Radomir,"LastName":Sohlich,"Age":30,"Active":true }
{"FirstName":Radomir,"LastName":Sohlich}

+-----------------------------------+-----------------------------------+
| 1                                 | 40FF81030101045573657201FF8200010 |
|                                   | 4010946697273744E616D65010C000108 |
| 2                                 | 4C6173744E616D65010C0001034167650 |
|                                   | 104000106416374697665010200000019 |
| 3                                 | FF8201075261646F6D69720107536F686 |
|                                   | C696368013C010100                 |
|                                   |                                   |
|                                   | {"FirstName":Radomir,"LastName":S |
|                                   | ohlich,"Age":30,"Active":true     |
|                                   | }                                 |
|                                   |                                   |
|                                   | {"FirstName":Radomir,"LastName":S |
|                                   | ohlich}                           |
+-----------------------------------+-----------------------------------+


Пакет encoding/gob для сериализации данных в бинарный формат

Для сериализации и десериализации бинарных данных gob нужен энкодер и
декодер. Функция gob.NewEncoder создает Encoder с базовым Writer. Каждый
вызов метода Encode сериализует объект в бинарный формат gob. Формат gob
является самоописывающим бинарным форматом. Это значит, что перед каждой
сериализованной структурой есть ее описание.

Для декодирования данных из сериализированной формы, ДЕКОДЕР должен
создаваться через вызов gob.NewDecoder с базовым Reader. Затем метод
Decode принимает указатель на структуру, куда нужно десериализировать
данные.

  ОБРАТИТЕ ВНИМАНИЕ, что бинарный формат gob не требует, чтобы тип
  источника и тип назначения совпадали.

Добавить комментарий Отменить ответ

Ваш адрес email не будет опубликован. Обязательные поля помечены *

Комментарий

Имя *

Email *

Сайт

Сохранить моё имя, email и адрес сайта в этом браузере для последующих
моих комментариев.

Уроки из раздела

-   Читаем ввод с клавиатуры в командной строке
-   Вывод данных и ошибок в командную строку
-   Открываем файл используя полный путь к нему
-   Читаем содержимое файла и сохраняем данные в переменную
-   Чтение и запись файлов в разной кодировке в Golang
-   Поиск позиции в файле через Seek()
-   Чтение и запись бинарных данных в Golang
-   Запись данных в нескольких файлах одновременно через MultiWriter
-   Туннель между процессами записи и чтения через io.Pipe()
-   СЕРИАЛИЗАЦИЯ ДАННЫХ В БИНАРНЫЙ ФОРМАТ GOB
-   Создание ZIP архивов в Golang
-   Парсинг большого XML файла в Go
-   Извлечение данных из поврежденного JSON массива в Golang

Изучение языка программирования Golang

[]
