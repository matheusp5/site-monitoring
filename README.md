# Monitoramento de Sites

Um problema que infelizmente ocorre frequentemente na minha vida é os meus sites por algum motivo cairem.<br>
Ele é totalmente feito em Go e toda vez que aberto, verifica seus sites.

<h1> Site caiu? </h1>
Se o sistema identificar algo errado com seu site, ele vai lhe informar e mostrar o erro.<br>
Ele faz isso de uma forma simples, mas que ajuda muito.

<h1> Como executar </h1>
O primeiro passo é ter o Go instalado em seu PC.
Caso não tenha, pode instalar usando o seguinte site <a href="https://go.dev">Golang</a>.
Após isso, abra seu terminal e vá a pasta que instalou o arquivo e execute o comando.
<br>
<br>

```go
go build main.go
```
Ele vai gerar um arquivo .exe

<h1> Meus sites </h1>
O usuário deverá modificar um array preenchendo com seus devidos sites.<br>
Abra o arquivo principal: main.go e a partir da linha 45, o array contendo todos os sites vai ser criado.<br>
Ai é simples, adicionar ou tirar sites do array.<br>
<br>

```go
sites := []string{
        "https://alura.com.br",               // Site teste
	"https://stackoverflow.com",          // Site teste
	"https://google.com",                 // Site teste
	"https://google.com/VaiDarErro",      // Site teste
	"https://google.com/VaiDaErroTambem", // Site teste
}
```
<br>
<br>

Caso tenha alguma duvida, estou aberto pra perguntas!<br>
Repositorio criado por Matheus

Bom código!
