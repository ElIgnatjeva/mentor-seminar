Запуск в Docker
``` bash 
docker volume create todo_data
docker run -d -p 8000:80 -v todo_data:/app/data elenaignateva/todo_app:latest 
```

Ссылка на образ в Docker Hub
https://hub.docker.com/repository/docker/elenaignateva/todo_app/general
