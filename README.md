# <div align="center">🍛 Curry - Câmbios 🍛</div>

## <div align="center">Valor do câmbio(USDxBRL) em lojas no Paraguay</div>

<br>

## 🎓 Sobre

Curry é um WebCrawler escrito em Golang com finalidade de verificar os valores de câmbio(USDxBRL)
em algumas lojas no Paraguay.

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
2021/11/27 01:59:09 Running and Listening on http://127.0.0.1:8000
2021/11/27 01:59:09 API Endpoint: /api
2021/11/27 01:59:10 Running crawler...

```

Agora basta acessar http://127.0.0.1:8000/api. Caso tenha alterado o endpoint favor verifique o log de saida do terminal
ou o arquivo config.yml.

⚠️️ **Não esqueça de configurar o arquivo *config.yml* conforme as suas necessidades!** ⚠️

<br>

*API Json Result:*

```json
{
  "id": 1,
  "createdAt": "2021-11-27T02:09:58.310052709-03:00",
  "updateAt": "2021-11-27T02:09:58.310052709-03:00",
  "deletedAt": null,
  "items": [
    {
      "id": 1,
      "name": "store1",
      "value": "R$ 5,78 BRL",
      "url": "https://www.store1.com.br/br"
    },
    {
      "id": 2,
      "name": "store2",
      "value": "R$ 5,76 BRL",
      "url": "https://www.store2.com/home"
    }
  ]
}
```

*config.yml:*

```yaml
app:
  host: 127.0.0.1
  port: '8000'

api:
  endpoint: /api

db:
  file_name: database.db

cache:
  maxage: 12  # Run crawler again after 12 hours

currency:
  symbol: BRL # Check https://github.com/joiggama/money/blob/v2.0.0/currencies.go
  symbol_space: true # true: R$ 5,78 | false:  R$5,78
  show_currency: true # true: R$ 5,78 BRL | false: R$ 5,78
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