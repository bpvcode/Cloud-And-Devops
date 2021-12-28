# syntax=docker/dockerfile:1

# NOTE:
# Should be the first line of every Dockerfile – it tells the Docker builder what syntax to use while parsing the Dockerfile and the location of the Docker syntax file.

FROM python:3.8-slim-buster

WORKDIR /app

# We copy just the requirements.txt first to leverage Docker cache
COPY ./requirements.txt /app/requirements.txt

RUN pip3 install -r requirements.txt

COPY . .

CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]
