- должен быть установлен го 1.16
- клонируем проект
- с корня проекта делаем **go run main.go**
- на последней строке видим такую строку

``
[GIN-debug] Listening and serving HTTP on :8080
``
Запрос на добавление работника 

``curl --location --request POST 'http://localhost:8080/employee' \
--header 'Content-Type: application/json' \
--data-raw '{
"name": "Test",
"email": "berezniy.d@gmail.com"
}'``

Запрос на получение списка работников

``curl --location --request GET 'http://localhost:8080/employee'``

C нереализованных требований 
- логер запросов\ответов (Делал бы через мидлварре)
- страница на хтмл (решил сделать REST API, так как страницы хтмл отходят в прошлое, клиент могу дописать на Vue)

