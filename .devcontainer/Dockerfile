FROM ubuntu:jammy
RUN apt-get update && apt-get install -y vim
RUN apt-get install -y curl

# 日本語化
RUN apt-get install -y language-pack-ja-base language-pack-ja locales
RUN locale-gen ja_JP.UTF-8
ENV LANG ja_JP.UTF-8