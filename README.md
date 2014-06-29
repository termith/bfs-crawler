### Реализация обхода Веб-графа в ширину на golang

###### Задача из курса "Параллельные и распределенные вычисления" ШАД Яндекс

##### Условия задачи

1. Напишите многопоточный поисковый робот, реализующий обход Web-графа в ширину  и сохраняющий на диск все посещенные страницы
2. При запуске роботу передаются URL начальной страницы и глубина обхода
3. Робот не должен посещать одну и ту же страницу более одного раза
4. Попытайтесь добиться максимальной скорости работы робота, обоснуйте используемый для этого подход

##### v 0.1.0 

**Без многопоточности**

`time go run main.go -url=http://ya.ru/ -depth=1`

`real	0m2.667s
user	0m0.585s
sys	0m0.071s`

##### v 1.0.0 

**Поиск и сохранение работают в разных рутинах**

Конфигурация: `Intel Atom N455 1.66 Hgz`

`time go run main.go -url=http://ya.ru/ -depth=2`

`real	1m12.813s
user	0m13.693s
sys	    0m1.396s`


