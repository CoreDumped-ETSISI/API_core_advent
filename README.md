env
JWT_SECRET=secret
DB_PASSWORD=secret
ADMIN_PASSWORD=secret
DB_PATH=advent-calendar.db


rutas
- GET / -> view years as a list [2024, 2023, 2022, 2021]
- GET /:year -> view a list of days [1, 2, 3]
- GET /:year/:day -> view a problem:
    - User needs to be authenticated
    - Problem is locked until a certain date {"desbloqueado": false, "tiempos_para_desbloquear": "1h 30m 15s"}
    - If user has not submitted a valid answer yet, return the problem {"problema": "Enunciado", "respuesta_valida": false}
    - If user has submitted a valid answer, return the problem and the user's answer {"problema": "Enunciado", "respuesta_valida": true, "respuesta_usuario": "Respuesta"}

- POST /:year/:day -> submit an answer to a problem {"solucion_propuesta": "Respuesta"}

- GET /ranking/:year -> view the ranking of users for a year ordered by the number of problems solved and the sum of the time it took 

        {
            1: {usuario: "Usuario", problemas_resueltos: 3, tiempo_total: "1h 30m 15s"}, 
            2: {usuario: "Usuario", problemas_resueltos: 2, tiempo_total: "1h 30m 15s"}, 
            3: {usuario: "Usuario", problemas_resueltos: 1, tiempo_total: "1h 30m 15s"}
        }

- GET /:year/resueltas -> view the problems solved by a user for a year 


- POST /admin -> login as admin {password}
- POST /admin/problemas -> create a problem 
        {
        "year": 2024,
        "dia": 1,
        "titulo": "New Problem",
        "enunciado": "Problem description",
        "solucion": "Solution description",
        "fecha_desbloqueo": "2024-11-19T12:00:00Z",
        "fecha_bloqueo": "2024-11-20T12:00:00Z"
        }
- PUT /admin/problemas/:year/:day -> update a problem
- DELETE /admin/problemas/:year/:day -> delete a problem
- GET /admin/info_problemas -> view all problems


- POST: /register -> register a user {Correo, Usuario, Constraseña}
- POST: /login -> login a user {Valor, Contraseña}

