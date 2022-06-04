FROM 889883130442.dkr.ecr.us-west-2.amazonaws.com/fair-images:python-3.8.7-3

EXPOSE 5000
EXPOSE 50051

ARG GITHUB_OAUTH_TOKEN

RUN apt-get update && apt-get install -y --no-install-recommends \
    libffi-dev \
    gcc \
    ssh

RUN apt-get install -y gnupg2

RUN mkdir -p /usr/share/man/man1
RUN apt-get install -y openjdk-11-jre pdftk

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app


ADD requirements.txt /usr/src/app/
RUN pip install --no-cache-dir -r requirements.txt

ADD . /usr/src/app

CMD ["python", "servers.py"]