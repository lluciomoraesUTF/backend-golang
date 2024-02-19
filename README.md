# backend-golang
Repositório feito para o processo seletivo da empresa DigitalSys
 Para a boa execução do código deve-se ter o PostgreSQL instalado na sua máquina. 
 A versão usada do PostgreSQL foi a 16.2-1 que pode ser encontrada no seguinte link:
 https://www.enterprisedb.com/downloads/postgres-postgresql-downloads. É fundamental que o Script SQL seja rodado no Postgre o servidor dele esteja na porta 5432, sua porta default.
 A linguagem GO também deve estar instalada na máquina a versão usada foi a go 1.22.0 o download pode ser feito através do seguinte link: https://go.dev/dl/
 As dependências do Go também devem ser instaladas para isso entre no diretório em que o repositório foi baixado via Prompt de Comando (CMD) ou via Gerenciador de Arquivos e onde se localiza o caminho do arquivo digite cmd.
 Após isso vamos baixar a primeira das dependências usadas no projeto o Framework GIN para isso digite no prompt de comando go get -u github.com/gin-gonic/gin para baixar o framework. 
 Também deve ser instalada a biblioteca ORM para isso digite no prompt de comando go get -u gorm.io/gorm.
 Também é fundamental que seja feito o download do driver específico para o PostgreSQL para isso digite no cmd o comando go get -u gorm.io/driver/postgres.
 Para ver os testes unitários e de integração deve se baixar a biblioteca Testfy que pode ser baixada escrevendo o comando go get -u github.com/stretchr/testify/assert
 no cmd.

