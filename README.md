<p style="
  display: block;
  width: 100px;
  height: 82px;
  margin: auto;
  filter: hue-rotate(180deg);
  background-image: url(https://assets.website-files.com/60147207eb0b6f4ddaeeaf73/601b7c07082f025136d6936c_logo-eventials.svg);
  background-size: 100px 82px;
  filter: invert(100%);" />


# Desafio backend

Linguagem: **Golang**

Informações detalhadas sobre o desafio podem ser encontradas no arquivo `README.ORIGINAL`.

Informações sobre estruturação do código e algumas anotações sobre ele podem ser encontradas no arquivo `README.DEV`

**Sumário**
- [Desafio backend](#desafio-backend)
  - [Considerações feitas durante o desenvolvimento](#considerações-feitas-durante-o-desenvolvimento)
  - [Como executar?](#como-executar)
  - [Migrations](#migrations)
    - [Dependências do Soda](#dependências-do-soda)
  - [DEV](#dev)

## Considerações feitas durante o desenvolvimento

1. Criar uma tabela "companies" (id,name:varchar(200),zip:varchar(5))

2. Para o enunciado abaixo:
    > The loaded data should have the following treatment:
    Existem várias possíveis abordagens, dentre elas:
    1. Abrir cada arquivo, tratar os dados e criar novos arquivos, CSV|sql
    2. Abrir cada arquivo, tratar os dados e já inserir no banco
    3. Importar diretamente no banco, e tratar os dados por lá (solução trivial)

    > Uma vez que não é especificado a abordagem, **irei optar pela (3)**.
    > Parte da (1ª abordagem já foi feita) e pode ser encontrada dentro do arquivo `tools/fileStream.7z`
    > Este arquivo tem uma função para abrir pedaços de arquivos absurdamente grandes e então tratá-los.

3. Pressupõe-se que tem um fluxo a ser respeitado, assim, terei que editar a tabela após esta ser criada
   e populada.

4. Tem alguns erros nos dados:
    - tim dieball aparece 2x
    - epicboardshop branch não tem addressZip

    > Dessa forma, é necessário fazer o matching usando os dois campos para poder inserir o terceiro

5. Como foi pedido para fazer os updates via *HTTP*, optei por criar a API primeiro, com boa parte
   das funções rest e então, usar um script para popular com as chamadas na **API**

6. Como **não** foi pedido documentação de endpoints e, dado que ficou apertado o prazo, optei por
   deixar isso como uma "melhoria futura" da implementação.

## Como executar?

```bash
# - Cria o banco docker
# - Baixa os pacotes necessários para a aplicação
# - Roda a migration que altera a tabela inicial
make setup

# Roda a suíte de testes
make check

# Inicia a API
make start

# Executa o script que abre o arquivo q2 e envia os dados para uma rota específica da api
# este script precisa que um outro terminal esteja rodando o "make start", i.e, a api precisa estar online
make run
```

## Migrations

As migrations são executadas usando o cli [Soda]. Para executar o Soda,
você precisa resolver suas [dependências](#dependências-do-soda)

> As migrations serão responsáveis por realizar a atualização do banco,
> preferencialmente, sem perder dados e mantendo uma compatibilidade com
> a estrutura já existente.


```powershell
# Dentro da pasta src:

# Gerando uma migration
soda.exe generate fizz CreateCompaniesTable

# Executando as migrations
soda.exe migrate

# Revertendo as migrations
soda.exe migrate down

# Vendo os stauts das migrations
soda.exe migrate status
```

### Dependências do Soda

Para executar as *migrations* você precisa ter um client do postgresql no path do seu sistema.

1. Windows:
   1. Caso tenha o [dbeaver](https://dbeaver.io/) instalado, uma solução é usar o client
      fornecido por eles. Bastando assim, apenas colocá-lo no path.
      O caminho para encontrar o local do client no dbeaver é:
      `Menu -> Database -> Driver Manager -> PostgreSQL -> Native Client`
      Com isto, você conseguirá copiar o path do client que eles baixam para o programa.
      Adicione o path no ambiente do sistema. **Também é necessário criar uma cópia dos executáveis no local,
      porém, sem a extensão `.exe`**

   2. Ainda no Windows, você pode simplesmente instalar apenas as "Command Line Tools" durante a instalação do [Postgresql]

   3. Por fim, uma outra abordagem utiliza o [`WSL2`](https://docs.microsoft.com/en-us/windows/wsl/install-win10).
       Ao ter o WSL2 instalado, é possível utilizar a abordagem linux. Um outro benefício é o comando make, requisitado
       na configuração deste projeto

2. Linux:

```bash
# Caso esteja usando linux, basta instalar pelo apt
sudo apt install postgresql-client-common
```


## DEV

```bash
# remove os módulos em cache
go clean -modcache

# path da pasta de apps
go env GOMODCACHE
```



<!-- Links -->

[Soda]: https://gobuffalo.io/en/docs/db/toolbox
[Postgresql]: https://www.postgresql.org/download/windows/