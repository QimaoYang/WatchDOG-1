FROM python:3.6-alpine as builder

COPY ./requirements.txt ./requirements.txt

RUN pip install -r ./requirements.txt

FROM alpine:latest as prod

COPY --from=builder $VIRTUAL_ENV $VIRTUAL_ENV

USER root

WORKDIR /home/app/

COPY . .

CMD ["python3", "app.py"]