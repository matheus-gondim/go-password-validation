# Introdução

Passwors verifier é uma API para verificar se uma senha é forte usando 6 regras:

* Quantidade mínima de caracteres (minSize)
* Quantidade mínima de caracteres maiúsculas (minUppercase)
* Quantidade mínima de caracteres minúsculas (minLowercase)
* Quantidade mínima de dígitos (minDigit)
* Quantidade mínima de caracteres especiais (minSpecialChars)
* Não ter nenhum caractere repetido em sequência (noRepeted)

Para isso deve-se fazer uma requisição POST para `/verify` seguindo o seguinte boby:

```
password: "TesteSenhaForte!1234&",
rules: [
    {rule: "minSize", value: 8},
    {rule: "minSpecialChars", value: 2},
    {rule: "noRepeted", value: 0},
    {rule: "minDigit", value: 4},
    {rule: "minUppercase", value: 3},
    {rule: "minLowercase", value: 5}
]
```
As rules devem ser removendo ou adicionando na lista conforme a necessidade do usuário.

O Retorno da API vai ser um json com dois campos `verify` e `noMatch`, conforme o exmeplo abaixo:

```
{
    "verify": false,
    "noMatch": ["minDigit"]
}
```

OBS: Há também a opção de usar a API com graphql mandando uma requisição para `/graphql`.

## Iniciando APP 

Parada rodar a API basta usar os comandos de `make run-http` para a versão da API usando http ou o comando `make run-graphql` para a API graphql