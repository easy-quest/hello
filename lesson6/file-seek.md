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



ПОИСК ПОЗИЦИИ В ФАЙЛЕ ЧЕРЕЗ SEEK()


В некоторых случаях может потребоваться прочитать или записать что-то в
определенное место в файле. К примеру, это может быть очень большой
файл, но мы знаем где именно нужно выполнить изменение в нем, для этого
не нужно читать его всего. Данная инструкция показывает, как
использовать поиск позиции в контексте операций над обычным текстовом
файле файла.

1. Создайте файл flatfile.txt со следующим содержимым:

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

Shell

123.Jun.......Wong...... 12..Novak.....Jurgen....
10..Thomas....Sohlich...

+-----------------------------------+-----------------------------------+
| 1                                 | 123.Jun.......Wong......          |
|                                   |                                   |
| 2                                 | 12..Novak.....Jurgen....          |
|                                   |                                   |
| 3                                 | 10..Thomas....Sohlich...          |
+-----------------------------------+-----------------------------------+

2. Создайте файл fileseek.go со следующим содержимым:

fileseek.go

Go

package main import ( "errors" "fmt" "os" ) const lineLegth = 25 func
main() { f, e := os.OpenFile("flatfile.txt", os.O_RDWR|os.O_CREATE,
os.ModePerm) if e != nil { panic(e) } defer f.Close()
fmt.Println(readRecords(2, "last", f)) if err := writeRecord(2, "first",
"Radomir", f); err != nil { panic(err) } fmt.Println(readRecords(2,
"first", f)) if err := writeRecord(10, "first", "Andrew", f); err != nil
{ panic(err) } fmt.Println(readRecords(10, "first", f))
fmt.Println(readLine(2, f)) } func readLine(line int, f *os.File)
(string, error) { lineBuffer := make([]byte, 24)
f.Seek(int64(line*lineLegth), 0) _, err := f.Read(lineBuffer) return
string(lineBuffer), err } func writeRecord(line int, column, dataStr
string, f *os.File) error { definedLen := 10 position := int64(line *
lineLegth) switch column { case "id": definedLen = 4 case "first":
position += 4 case "last": position += 14 default: return
errors.New("Столбец не определен") } if len([]byte(dataStr)) >
definedLen { return fmt.Errorf("Максимальная длина для '%s' это %d",
column, definedLen) } data := make([]byte, definedLen) for i := range
data { data[i] = '.' } copy(data, []byte(dataStr)) _, err :=
f.WriteAt(data, position) return err } func readRecords(line int, column
string, f *os.File) (string, error) { lineBuffer := make([]byte, 24)
f.ReadAt(lineBuffer, int64(line*lineLegth)) var retVal string switch
column { case "id": return string(lineBuffer[:3]), nil case "first":
return string(lineBuffer[4:13]), nil case "last": return
string(lineBuffer[14:23]), nil } return retVal, errors.New("Столбец не
определен") }

+-----------------------------------+-----------------------------------+
| 1                                 | package main                      |
|                                   |                                   |
| 2                                 |                                   |
|                                   |                                   |
| 3                                 | import (                          |
|                                   |                                   |
| 4                                 | "errors"                          |
|                                   |                                   |
| 5                                 | "fmt"                             |
|                                   |                                   |
| 6                                 | "os"                              |
|                                   |                                   |
| 7                                 | )                                 |
|                                   |                                   |
| 8                                 |                                   |
|                                   |                                   |
| 9                                 | const lineLegth = 25              |
|                                   |                                   |
| 10                                |                                   |
|                                   |                                   |
| 11                                | func main() {                     |
|                                   |                                   |
| 12                                |                                   |
|                                   |                                   |
| 13                                | f, e :=                           |
|                                   | os.OpenFile("flatfile.txt",       |
| 14                                | os.O_RDWR|os.O_CREATE,            |
|                                   | os.ModePerm)                      |
| 15                                |                                   |
|                                   | if e != nil {                     |
| 16                                |                                   |
|                                   | panic(e)                          |
| 17                                |                                   |
|                                   | }                                 |
| 18                                |                                   |
|                                   | defer f.Close()                   |
| 19                                |                                   |
|                                   |                                   |
| 20                                |                                   |
|                                   | fmt.Println(readRecords(2,        |
| 21                                | "last", f))                       |
|                                   |                                   |
| 22                                | if err := writeRecord(2, "first", |
|                                   | "Radomir", f); err != nil {       |
| 23                                |                                   |
|                                   | panic(err)                        |
| 24                                |                                   |
|                                   | }                                 |
| 25                                |                                   |
|                                   | fmt.Println(readRecords(2,        |
| 26                                | "first", f))                      |
|                                   |                                   |
| 27                                | if err := writeRecord(10,         |
|                                   | "first", "Andrew", f); err != nil |
| 28                                | {                                 |
|                                   |                                   |
| 29                                | panic(err)                        |
|                                   |                                   |
| 30                                | }                                 |
|                                   |                                   |
| 31                                | fmt.Println(readRecords(10,       |
|                                   | "first", f))                      |
| 32                                |                                   |
|                                   | fmt.Println(readLine(2, f))       |
| 33                                |                                   |
|                                   | }                                 |
| 34                                |                                   |
|                                   |                                   |
| 35                                |                                   |
|                                   | func readLine(line int, f         |
| 36                                | *os.File) (string, error) {       |
|                                   |                                   |
| 37                                | lineBuffer := make([]byte, 24)    |
|                                   |                                   |
| 38                                | f.Seek(int64(line*lineLegth), 0)  |
|                                   |                                   |
| 39                                | _, err := f.Read(lineBuffer)      |
|                                   |                                   |
| 40                                | return string(lineBuffer), err    |
|                                   |                                   |
| 41                                | }                                 |
|                                   |                                   |
| 42                                |                                   |
|                                   |                                   |
| 43                                | func writeRecord(line int,        |
|                                   | column, dataStr string, f         |
| 44                                | *os.File) error {                 |
|                                   |                                   |
| 45                                | definedLen := 10                  |
|                                   |                                   |
| 46                                | position := int64(line *          |
|                                   | lineLegth)                        |
| 47                                |                                   |
|                                   | switch column {                   |
| 48                                |                                   |
|                                   | case "id":                        |
| 49                                |                                   |
|                                   | definedLen = 4                    |
| 50                                |                                   |
|                                   | case "first":                     |
| 51                                |                                   |
|                                   | position += 4                     |
| 52                                |                                   |
|                                   | case "last":                      |
| 53                                |                                   |
|                                   | position += 14                    |
| 54                                |                                   |
|                                   | default:                          |
| 55                                |                                   |
|                                   | return errors.New("Столбец не     |
| 56                                | определен")                       |
|                                   |                                   |
| 57                                | }                                 |
|                                   |                                   |
| 58                                |                                   |
|                                   |                                   |
| 59                                | if len([]byte(dataStr)) >         |
|                                   | definedLen {                      |
| 60                                |                                   |
|                                   | return fmt.Errorf("Максимальная   |
| 61                                | длина для '%s' это %d", column,   |
|                                   | definedLen)                       |
| 62                                |                                   |
|                                   | }                                 |
| 63                                |                                   |
|                                   |                                   |
| 64                                |                                   |
|                                   | data := make([]byte, definedLen)  |
| 65                                |                                   |
|                                   | for i := range data {             |
| 66                                |                                   |
|                                   | data[i] = '.'                     |
| 67                                |                                   |
|                                   | }                                 |
| 68                                |                                   |
|                                   | copy(data, []byte(dataStr))       |
| 69                                |                                   |
|                                   | _, err := f.WriteAt(data,         |
| 70                                | position)                         |
|                                   |                                   |
| 71                                | return err                        |
|                                   |                                   |
| 72                                | }                                 |
|                                   |                                   |
| 73                                |                                   |
|                                   |                                   |
| 74                                | func readRecords(line int, column |
|                                   | string, f *os.File) (string,      |
| 75                                | error) {                          |
|                                   |                                   |
| 76                                | lineBuffer := make([]byte, 24)    |
|                                   |                                   |
| 77                                | f.ReadAt(lineBuffer,              |
|                                   | int64(line*lineLegth))            |
| 78                                |                                   |
|                                   | var retVal string                 |
| 79                                |                                   |
|                                   | switch column {                   |
|                                   |                                   |
|                                   | case "id":                        |
|                                   |                                   |
|                                   | return string(lineBuffer[:3]),    |
|                                   | nil                               |
|                                   |                                   |
|                                   | case "first":                     |
|                                   |                                   |
|                                   | return string(lineBuffer[4:13]),  |
|                                   | nil                               |
|                                   |                                   |
|                                   | case "last":                      |
|                                   |                                   |
|                                   | return string(lineBuffer[14:23]), |
|                                   | nil                               |
|                                   |                                   |
|                                   | }                                 |
|                                   |                                   |
|                                   |                                   |
|                                   |                                   |
|                                   | return retVal,                    |
|                                   | errors.New("Столбец не            |
|                                   | определен")                       |
|                                   |                                   |
|                                   | }                                 |
+-----------------------------------+-----------------------------------+

3. Запустите код через go run fileseek.go;
4. Посмотрите на вывод:

[файлы golang]

5. Отображение файла в xxd flatfile.txt:

[файлы go]


Как написать код для программы, что находит позицию в файле?

В предыдущем примере используется файл flatfile.txt, над которым
выполняется поиск в нем, прочтение и запись в определенную позицию в
файле. В общем и целом для перемещения позиции текущего указателя в File
может использоваться метод Seek. Он принимает два аргумента, это позиция
и то, как считать позицию, _0 — относительно оригиналу файла, 1 —
относительно текущей позиции, 2 — относительно конца строки_. В таком
случае вы можете переместить курсор внутри файла. Метод Seek
используется для имплементации функции readLine в коде выше.

os.File также содержит методы ReadAt и WriteAt. При использовании
методов, байты должны быть записаны/прочитаны, также должна даваться
отправная точка. Это упрощает запись и чтение для определенной позиции в
файле.

  НА ЗАМЕТКУ: Обратите внимание, что в примере предполагается, что
  каждая руна составляет один байт, что не обязательно должно быть так
  для специальных символов и так далее.

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
-   ПОИСК ПОЗИЦИИ В ФАЙЛЕ ЧЕРЕЗ SEEK()
-   Чтение и запись бинарных данных в Golang
-   Запись данных в нескольких файлах одновременно через MultiWriter
-   Туннель между процессами записи и чтения через io.Pipe()
-   Сериализация данных в бинарный формат gob
-   Создание ZIP архивов в Golang
-   Парсинг большого XML файла в Go
-   Извлечение данных из поврежденного JSON массива в Golang

Изучение языка программирования Golang

[]
