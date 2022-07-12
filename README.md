## API Rest Go com Mux/Gorilla

## Configurando o Postgres com Docker

Atenção: as configurações deste tópico devem ser feitas antes de
tentar iniciar a aplicação na máquina local.

===================================

Para executar o projeto em ambiente local, primeiramente é
necessário inicializar o Docker.

Ao receber a mensagem "Docker Desktop is running", abrir
a pasta deste projeto no terminal e executar o seguinte comando:

`docker-compose up` -> Inicializa o container do Docker

No arquivo *docker-compose.yml*, está configurado o serviço *pgadmin-compose*.
Este serviço permite fazer o login no Postgres com o e-mail e senha informados nos
atributos de environment. Para acessar a página de login, acessar:

`localhost:54321` -> porta definida dentro do serviço *pgadmin-compose* do
docker-compose

Após fazer o login, criar novo servidor no Postgres. O host deve ser o
endereço IP da máquina do Docker que foi inicializada acima.

Para identificar o IP do host, abrir nova janela do terminal (a que subiu
o container do Docker deve continuar em execução) e digitar o seguinte
comando:

`docker-compose exec postgres sh` -> indica ao docker que deve abrir
a máquina rodando o serviço que foi definido no *docker-compose* com o nome
"postgres"

Dentro do container, executar o comando `hostname -i` e copiar o endereço
IP retornado. Este endereço deve ser utilizado no host do novo servidor
que está sendo criado no Postgres.

O host do novo servidor deverá ser configurado como:

`{ip do docker container}:5432` -> a porta foi definida dentro do *docker-compose*,
no serviço "postgres"