# api students

API to manage 'Golang do zero' course students

Routes

- GET /students - lista de estudantes
- POST /students - cria um estudante
- GET /students/:id - retorna um estudante
- PUT /students/:id - atualiza um estudante
- DELETE /students/:id - deleta um estudante
- GET /students?active=<true/false> - lista de estudantes ativos ou não

Informação de estuante:
- name (string)
- CPF (int)
- Email (string)
- Age (int)
- Active (bool)
