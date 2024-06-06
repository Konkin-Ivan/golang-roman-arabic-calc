FROM golang:latest

# Устанавливаем bash (если он не установлен)
RUN apt update && apt install -y bash
# Создаем пользователя
RUN useradd -m -s /bin/bash -G sudo myuser
WORKDIR /app/app
# Переключение на пользователя
USER myuser