# account

Serviço responsável por cadastro de contas e transções.

## Executando o projeto

### Via docker-compose

```sh
docker-compose up --build
```

### Via IDE

#### Adicione as seguintes variáves de ambiente

|Nome         |Valor      |
--------------|-------------
|DB_HOST      |localhost   |
|DB_NAME      |postgres    |
|DB_USERNAME  |postgres    |
|DB_PASSWORD  |account     |
|DB_PORT      |5432        |
|DB_SSL       |false       |

#### Inicie o banco de dados

```sh
docker-compose up --build postgres
```

Execute o arquivo `main.go`.  
O serviço será iniciado na porta `8080`.

## Endpoints

### Criar uma conta

#### URL

```
[POST] /api/accounts
```

#### Payload

```json
{
    "document_number": "123"
}
```

#### Resposta

```json
{
    "account_id": 1,
    "document_number": "123"
}
```

#### cURL
```sh
curl --location --request POST 'http://localhost:8080/api/accounts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "document_number": "123"
}'
```

### Recuperar uma conta

#### URL

```
[POST] /api/accounts/:account_id
```

#### Resposta

```json
{
    "account_id": 1,
    "document_number": "123"
}
```

#### cURL

```sh
curl --location --request GET 'http://localhost:8080/api/accounts/1'
```

### Criar uma transação

#### URL
```
[POST] /api/transactions
```

#### Payload

```json
{
    "account_id": 1,
    "operation_type": "FOR_CACHE",
    "amount": 123.45
}
```

#### Operações Disponíveis

|Nome       |Descrição        |
|-----------|-----------------|
|FOR_CACHE  |Compra à vista   |
|INSTALLMENT|Compra parcelada |
|WITHDRAWN  |Saque            |
|PAYMENT    |Pagamento        |

#### Resposta

```json
{
    "account_id": 1,
    "amount": 123.45,
    "operation_type": "FOR_CACHE",
    "transaction_id": 1
}
```

## Executando os testes unitários

```sh
go test ./...
```

## Postman

No diretório `postman` existe uma collection que pode ser importada no Postman para a realização de testes funcionais.
