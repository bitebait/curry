# <div align="center">🍛 Curry Câmbios</div>

## <div align="center">Valor do câmbio(USDxBRL) em lojas no Paraguay</div>

<br>

## 🎓 Sobre

Curry é um WebCrawler escrito em Golang com finalidade de verificar o valor do câmbio de **Dólar** para **Real** (**USD**x**BRL**) em algumas lojas no Paraguay.

* * *

## :tada: Features

1. :white_check_mark: Web Crawler.
1. :white_check_mark: API JSON Endpoint.
1. :white_check_mark: In-Memory Cache.
1. :white_check_mark: Scheduler.
1. :white_check_mark: SQLite Database to keep history.
1. :white_check_mark: Custom settings.
1. :white_check_mark: Easy to add new spiders.

* * *

## 📌 Uso

Para utilizar, basta seguir os passos abaixo:

### 📜 Rodando

```sh
git clone https://github.com/bitebait/curry.git
cd curry/
go run .
```

Exemplo de saida do terminal:

```
 ██████╗██╗   ██╗██████╗ ██████╗ ██╗   ██╗
██╔════╝██║   ██║██╔══██╗██╔══██╗╚██╗ ██╔╝
██║     ██║   ██║██████╔╝██████╔╝ ╚████╔╝ 
██║     ██║   ██║██╔══██╗██╔══██╗  ╚██╔╝  
╚██████╗╚██████╔╝██║  ██║██║  ██║   ██║   
 ╚═════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   
2021/11/27 19:06:00 Running and Listening on :8000
2021/11/27 19:06:00 API Endpoint: /api
2021/11/27 19:06:01 Running crawler...
...
2021/11/27 19:06:09 FINISHED: 30 of 30 urls visited.
2021/11/27 19:06:09 CRAWLER function took 8.403559105s.
```

Agora basta acessar <http://127.0.0.1:8000/api>.
Caso tenha alterado o endpoint favor verifique o log de saida do terminal ou o arquivo config.yml.

⚠️️ **Não esqueça de configurar o arquivo *config.yml* conforme as suas necessidades!** ⚠️

<br>

*API Json Result:*

```json
{
  "createdAt": "2021-11-27T15:21:06.343499301-03:00",
  "items": [
    {
      "name": "example_store1",
      "currency": "BRL",
      "value": "5.7",
      "url": "https://www.example_store1.com/"
    },
    {
      "name": "example_store2",
      "currency": "BRL",
      "value": "5.72",
      "url": "https://www.example_store2.com/"
    }
  ]
}
```

*config.yml:*

```yaml
api:
  port: 8000 # api listen port
  endpoint: /api # api endpoint

db:
  file_name: database.db # sqlite file name

cache:
  max_age: 12  # Refresh cache every 12 hours

currency:
  currency: BRL # Don't change
```

**🕷️ Para mais informações sobre spiders disponíveis e como adicionar uma nova spider, de uma olhada em
[/spiders](https://github.com/bitebait/curry/tree/master/crawler/spiders) e [func NewSpider()](https://github.com/bitebait/curry/blob/master/crawler/spiders/spiders.go)**.

* * *
<br>

## ✅ **Lista de lojas monitoradas**

<br>

1. **[alboradainfo](https://www.alboradainfo.com/)**
1. **[aromastore](https://www.aromastore.com.br/)**
1. **[atacadocollections](https://www.atacadocollections.com/)**
1. **[atacadoconnect](https://atacadoconnect.com/)**
1. **[audiumelectronics](https://www.audiumelectronics.com/home)**
1. **[bonanzacambios](https://bonanzacambios.com.py/)**
1. **[cambioschaco](https://www.cambioschaco.com.py/pt-br/)**
1. **[digitalcenterpy](https://digitalcenterpy.com/)**
1. **[cellshop](https://www.cellshop.com/br/)**
1. **[comprasparaguai](https://www.comprasparaguai.com.br/)**
1. **[dolarpy](https://www.dolarpy.com.br/)**
1. **[gabahobby](https://www.gabahobby.com/)**
1. **[hbgames](http://www.hbgamespy.com/)**
1. **[icompy](http://icompy.com/)**
1. **[infinitysport](https://www.infinitysport.com.py/)**
1. **[lgimportados](https://www.lgimportados.com/)**
1. **[madridcenter](https://www.madridcenter.com/)**
1. **[megaeletro](https://www.megaeletro.com.py/br)**
1. **[megaeletronicos](https://www.megaeletronicos.com/br)**
1. **[mercosurcambios](https://site.mercosurcambios.com/)**
1. **[mundodocelular](https://www.mundodocelular.com/)**
1. **[oneclick](https://oneclick.com.py/)**
1. **[pioneerinter](https://www.pioneerinter.com/)**
1. **[probook](https://www.probook.com.py/)**
1. **[shoppingcentropioneer](https://shoppingcentropioneer.com/)**
1. **[tcheloco](https://www.tcheloco.com.py/br/)**
1. **[tecombras](https://www.tecombras.net/)**
1. **[topdek](https://www.topdek.com.br/br)**
1. **[victoriastore](https://www.victoriastore.com.br/)**
1. **[visaovip](http://www.visaovip.com/)**

<br>

### 📄 Exoneração de responsabilidade e problemas conhecidos

* Qualquer uso do script é de responsabilidade apenas do usuário. Os usuários do script devem agir de acordo com os
  termos dos sites acessados.
* Como acontece com todos os sites, a estrutura do site pode mudar no futuro e, portanto, como costuma acontecer com
  scripts de scraping, descontinue-o. Não é realmente uma questão de saber se o código-fonte do site irá mudar, mas sim
  quando (então aproveite enquanto ainda está funcionando)

<br>

### 🔒 Licença

Todo o conteúdo apresentado nos sites pertence aos criadores originais.

A licença abaixo se refere apenas ao script e não ao conteúdo scrapado.

[Licença - MIT](https://github.com/bitebait/curry/blob/master/LICENSE)

<br>

### 🔥 Sinta-se à vontade para contribuir com o código (; 🔥
