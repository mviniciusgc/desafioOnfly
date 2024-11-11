# Microsserviço de Gestão de Pedidos de Viagem Corporativa

Este é um microsserviço desenvolvido em Golang para gerenciar pedidos de viagem corporativa. O serviço oferece uma API REST para realizar operações de criação, atualização, consulta e listagem de pedidos de viagem.

## Funcionalidades

O microsserviço permite realizar as seguintes operações:

### 1. **Criar um Pedido de Viagem**
Cadastrar um novo pedido de viagem, com informações como ID do pedido, nome do solicitante, destino, data de ida, data de volta e status.

**Rota:**
POST http://localhost:8080/travel

**Exemplo de corpo da requisição:**
```json
{
  "RequesterName": "John oioioioi",
  "Destination": "Paris",
  "DepartureDate": "2024-12-01T12:00:00Z",
  "ReturnDate": "2024-12-15T12:00:00Z",
  "Status": "cancelado",
  "CreatedAt": "2024-11-01T12:00:00Z"
}
```
2. Listar Todos os Pedidos de Viagem
Retorna todos os pedidos de viagem cadastrados, com a possibilidade de filtrar por status.

Rota:
GET http://localhost:8080/travel/listAll?status=cancelado

3. Consultar Pedido de Viagem pelo ID
Retorna as informações completas de um pedido de viagem com base no ID fornecido.

Rota:
GET http://localhost:8080/travel/{id}

4. Atualizar o Status de um Pedido de Viagem
Permite atualizar o status de um pedido de viagem para o status (solicitado, aprovado, cancelado).

Rota:
PATCH http://localhost:8080/travel/updateStatus/{id}
Exemplo de corpo da requisição:
```json
{
  "Status": "aprovado"
}
```
Estrutura do Banco de Dados
O microsserviço utiliza duas tabelas no banco de dados PostgreSQL:

logs: Armazena logs do sistema.
travel: Armazena os pedidos de viagem.

Como Rodar o Projeto
Clone o repositório:
git clone https://github.com/usuario/projeto-viagem.git
cd desafioOnfly

Inicie os containers com Docker Compose: O Docker Compose é utilizado para subir o PostgreSQL e a imagem do Golang.
docker-compose up
Isso pode levar alguns minutos.

O serviço estará disponível em http://localhost:8080.

Tecnologias Utilizadas
Golang: Linguagem de programação principal.
PostgreSQL: Banco de dados utilizado.
Docker Compose: Orquestração de containers.
