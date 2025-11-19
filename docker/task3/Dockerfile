FROM python:3.12-slim
WORKDIR /app
ARG DEFAULT_USER_NAME=student
ENV USER_NAME=$DEFAULT_USER_NAME
COPY show_env.py .
CMD ["python", "show_env.py"]