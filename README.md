# <div align="center">🍛 Curry - Câmbios 🍛</div>

## <div align="center">Valor do câmbio(USDxBRL) em lojas no Paraguay</div>

<br>

## 🎓 Sobre

Curry é um WebCrawler escrito em Golang com finalidade de verificar os valores de câmbio(USDxBRL)
em algumas lojas no Paraguay.

Demo: https://gocurry.herokuapp.com/api

* * *

## 📌 Uso

Para utilizar, basta seguir os passos abaixo:

### 📜 Rodando

```sh
git clone https://github.com/bitebait/curry.git
cd curry/
go run app.go
```

Exemplo de saida do terminal:

```sh
 ██████╗██╗   ██╗██████╗ ██████╗ ██╗   ██╗
██╔════╝██║   ██║██╔══██╗██╔══██╗╚██╗ ██╔╝
██║     ██║   ██║██████╔╝██████╔╝ ╚████╔╝ 
██║     ██║   ██║██╔══██╗██╔══██╗  ╚██╔╝  
╚██████╗╚██████╔╝██║  ██║██║  ██║   ██║   
 ╚═════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   
2021/11/27 15:20:58 Running and Listening on :8000
2021/11/27 15:20:58 API Endpoint: /api
2021/11/27 15:20:59 Running crawler...
```

Agora basta acessar http://127.0.0.1:8000/api. Caso tenha alterado o endpoint favor verifique o log de saida do terminal
ou o arquivo config.yml.

⚠️️ **Não esqueça de configurar o arquivo *config.yml* conforme as suas necessidades!** ⚠️

<br>

*API Json Result:*

```json
{
  "id": 1,
  "createdAt": "2021-11-27T15:21:06.343499301-03:00",
  "updateAt": "2021-11-27T15:21:06.343499301-03:00",
  "deletedAt": null,
  "items": [
    {
      "id": 1,
      "name": "example_store1",
      "currency": "BRL",
      "value": 5.7,
      "url": "https://www.example_store1.com/"
    },
    {
      "id": 2,
      "name": "example_store2",
      "currency": "BRL",
      "value": 5.72,
      "url": "https://www.example_store2.com/"
    }
  ]
}
```

*config.yml:*

```yaml
app:
  host:
  port: '8000'

api:
  endpoint: /api

db:
  file_name: database.db

cache:
  maxage: 12  # MaxAge = 12 Hours

currency:
  currency: BRL # Don't change
```

**🕷️ Para mais informações sobre spiders disponíveis de uma olhada em
[/spiders](https://github.com/bitebait/curry/tree/master/crawler/spiders) e
[AllSpiders()](https://github.com/bitebait/curry/blob/master/crawler/spiders/spiders.go)**

<br>

### 📄 Exoneração de responsabilidade e problemas conhecidos

- Qualquer uso do script é de responsabilidade apenas do usuário. Os usuários do script devem agir de acordo com os
  termos dos sites acessados.
- Como acontece com todos os sites, a estrutura do site pode mudar no futuro e, portanto, como costuma acontecer com
  scripts de scraping, descontinue-o. Não é realmente uma questão de saber se o código-fonte do site irá mudar, mas sim
  quando (então aproveite enquanto ainda está funcionando)

<br>

### 🔒 Licença

Todo o conteúdo apresentado nos sites pertence aos criadores originais.

A licença abaixo se refere apenas ao script e não ao conteúdo scrapado.

[Licença - MIT](https://github.com/bitebait/curry/blob/master/LICENSE)

<br>

### 🔥 Sinta-se à vontade para contribuir com o código (; 🔥