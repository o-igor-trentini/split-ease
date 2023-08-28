# Split Ease - Backend  #

## Iniciando

As dependências serão instaladas ao inicializar a aplicação pela primeira vez.

A aplicação utiliza variáveis de ambiente, dessa forma é necessário realizar uma cópia do arquivo `.env.example`
(arquivo "modelo") para `.env`:

```bash
cp .env.example .env
```
> **Observação:** A maioria das variáveis possui valor "intuitivo" ou tem seus possíveis valores indicados num
> comentário acima. Entretanto, há possibilidade de alguma variável de ambiente não atender nenhuma das duas situações.
> Quando isso acontecer é necessário entrar em contato com alguém que tenha acesso a essa informação.

Após realizar a configuração das variáveis de ambiente, o projeto pode ser executado com o seguinte comando:

```bash
go run main.go
```

Por padrão o servidor irá utilizar [http://localhost:8080](http://localhost:8080), mas isto pode ser alterado nas
variáveis de ambiente.

Após iniciar o projeto, verifique se o servidor inicializou clicando [aqui](http://localhost:8080/api/health). 

## Facilitadores ##

Foram adicionados alguns facilitadores para agilizar e padronizar algumas ações. Esses facilitadores são executados ao
executar o projeto passando algum parâmetro.

> **Atenção:** Ao passar o parâmetro ```-mr``` ou ```-migrate-and-run```, pois ele inicia o servidor HTTP. Pode ser que algum
> dos outros argumentos sejam ignorados. É recomendado usá-lo de forma individual.

### Parâmetros ###
- ```-cds``` ou ```-create-database-structure```
    - Cria a estrutura do banco:
      - Tabelas;
      - Tipos de dados;
      - Constraints;
      - Executa ```migrations```;
      - Entre outros.
- ```-cm``` ou ```-create-migration```
  - Cria um arquivo padronizado para escrever uma ```migration```.
- ```-m``` ou ```-migrate```
  - Executa as ```migrations```.
- ```-mr``` ou ```-migrate-and-run```
  - Executa as ```migrations``` e inicia o servidor HTTP. 
- ```-t``` ou ```-test```
  - Executa todos os testes do projeto e gera um relatório de cobertura

### Utilização ###

Para utilizar algum facilitador, utilize o seguinte comando na raiz do projeto:

```bash
go run main.go -t
```
