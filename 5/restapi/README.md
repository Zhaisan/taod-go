# rest-api

# user-service

# REST API

GET /users -- list of users -- 200(всё окей), 404(сущность не найдена), 500(мы не смогли обработать запрос, сервер лег)
GET /users/:id -- user by id -- 200, 404, 500
POST /users/:id -- create user -- 204(no content), 4хх, Header Location: url
PUT /users/:id -- fully update user -- 204/200(либо ничего не отдаем либо полностью обновленный пользователь), 404, 400, 500
PATCH /users/:id -- partially update user -- 204/200, 404, 400, 500
DELETE /users/:id -- delete user by id -- 204, 404, 400 