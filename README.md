# API Cryptography

Esta é a solução para o desafio proposto em [Backend Brasil](https://github.com/backend-br/desafios/blob/master/cryptography/PROBLEM.md). Esta API permite o cadastro de um documento, um cartão de crédito e um valor. Os dados sensíveis, como o documento e o cartão de crédito, são criptografados antes de serem salvos no banco de dados e descriptografados ao serem retornados ao usuário.

## Funcionalidades

- **Cadastro de Dados:** Permite cadastrar um documento, um cartão de crédito e um valor.
- **Criptografia:** Os dados sensíveis são criptografados antes de serem armazenados.
- **Descriptografia:** Os dados sensíveis são descriptografados ao serem retornados.
- **CRUD Completo:** Possui rotas para criar, visualizar, editar e excluir registros.

## Tarefas

- [ ] Aprender a criptografar com Go
- [ ] Rota de Registro
  - [ ] Criar rota
  - [ ] Validar dados do corpo da requisição
  - [ ] Criptografar dados
  - [ ] Salvar no banco de dados
- [ ] Rota para Visualizar os Dados
- [ ] Rota para Editar os Dados
- [ ] Rota para Excluir Usuário
