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
