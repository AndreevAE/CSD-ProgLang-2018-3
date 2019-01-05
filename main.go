/* Реализовать на языке Go tcp клиент/сервер (общее приложение, режим определяется параметрами), способное установить
сокет соединение и последовательно обмениваться сообщениями. На каждой стороне реализуется алгоритм переменных ключей
(смотри реализацию на Python - protector.py). На каждом шаге обмена вычисляется следующий ключ и сравнивается с
полученным от второй стороны.

Шаг 1. Установление соединения. Клиент подключается к серверу и передает стартовую строку и первый ключ

Шаг 2. Сервер на основе строки и ключа генерирует новый ключ и отдает его клиенту

Шаг 3. Клиент сравнивает полученный ключ со следующим ключом, и, если все успешно, создает новый ключ и отправляет
следующее сообщение на сервер.

Шаг 4..10 - аналогично

На каждом шагу приложение должно выводить в консоли текущий статус, текущий ключ и отправленное/полученное сообщение.

По желанию можно дополнить код функцией чата и вводить попутное сообщение/ответ из консоли.

 При запуске программа должна принимать два параметра командной строки:

 1) порт - режим сервера или ip:port - режим клиента
 2) -n 100 - кол-во одновременных подключений

пример клиента и сервера: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
*/

package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}