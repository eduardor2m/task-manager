# API de Tasks em Go

Esta é uma API em Go que fornece operações CRUD (Create, Read, Update, Delete) para gerenciar tarefas.

## Configuração

Certifique-se de ter o Go instalado em sua máquina.

1. Clone este repositório:

```shell
git clone https://github.com/eduardor2m/task-manager.git
```

2. Navegue até o diretório do projeto:

```shell
cd task-manager
```

3. Instale as dependências:

```shell
go mod download
```

4. Execute o servidor:

```shell
go run src/api/app/main.go
```

## Uso

A API possui as seguintes rotas disponíveis:

- `GET /tasks`: Lista todas as tarefas.
- `GET /tasks/{id}`: Retorna os detalhes de uma tarefa específica.
- `POST /tasks`: Cria uma nova tarefa.
- `PUT /tasks/{id}`: Atualiza os detalhes de uma tarefa existente.
- `DELETE /tasks/{id}`: Exclui uma tarefa existente.

### Exemplos de requisições

#### Listar todas as tarefas

```
GET /tasks
```

#### Retornar os detalhes de uma tarefa

```
GET /tasks/{id}
```

#### Criar uma nova tarefa

```
POST /tasks

Body da requisição:
{
  "title": "Minha nova tarefa",
  "description": "Detalhes da tarefa"
}
```

#### Atualizar os detalhes de uma tarefa existente

```
PUT /tasks/{id}

Body da requisição:
{
  "title": "Tarefa atualizada",
  "description": "Detalhes atualizados da tarefa"
  "completed": true or false
}
```

#### Excluir uma tarefa existente

```
DELETE /tasks/{id}
```

## Contribuição

Contribuições são bem-vindas! Se você encontrar algum problema ou tiver uma sugestão, por favor, abra uma issue ou envie um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
