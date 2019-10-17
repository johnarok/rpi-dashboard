FROM golang:1.13 as builder
RUN apt-get update && apt-get install libwebkit2gtk-4.0 -y
