# Агрегатор сайтов по заведению рекламы

0. Запусти hello world
1. Пройти https://academy.threedots.tech/trainings/one-evening-go/exercise/159ce245-e0ea-4aae-884a-932de793293c?utm_content=cta&utm_source=one-evening-go-landing
2. Добавить ручки (делай в одном файле main.go, писать данные в глобальную переменную)
    1. `POST` `/profile` request: body{"name", "age", ...}, response: 200, id; 400 BadRequest, 500 если что-то пошло по пизде
    2. `GET` `/profile/:id` response: 200, body{"name", "age", ...}; 400 BadRequest, 500 если что-то пошло по пизде
    3. `PUT` `/profile/:id` request: body{"name", "age", ...}, response: 200, id; 400 BadRequest, 500 если что-то пошло по пизде
    4. `DELETE` `/profile/:id` response: 200, id; 400 BadRequest, 500 если что-то пошло по пизде

3. слайс vs массив, структуру слайса, что проиходит при append
4. map, бакеты, коллизии, миграции