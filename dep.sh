#!/bin/bash

# Baixa as dependências do projeto Go
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/stretchr/testify/assert
