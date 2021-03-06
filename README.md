# <div align="center">π Curry CΓ’mbios</div>

## <div align="center">Valor do cΓ’mbio(USDxBRL) em lojas no Paraguay</div>

<br>

## π Sobre

Curry Γ© um WebCrawler escrito em Golang com finalidade de verificar o valor do cΓ’mbio de **DΓ³lar** para **Real** (**USD**x**BRL**) em algumas lojas no Paraguay.

* * *

## π Uso

Para utilizar, basta seguir os passos abaixo:

### π Rodando

```sh
git clone https://github.com/bitebait/curry.git
cd curry/
go run app.go
```

Exemplo de saida do terminal:

```sh
 ββββββββββ   ββββββββββ βββββββ βββ   βββ
βββββββββββ   βββββββββββββββββββββββ ββββ
βββ     βββ   βββββββββββββββββββ βββββββ 
βββ     βββ   βββββββββββββββββββ  βββββ  
ββββββββββββββββββββ  ββββββ  βββ   βββ   
 βββββββ βββββββ βββ  ββββββ  βββ   βββ   
2021/11/27 15:20:58 Running and Listening on :8000
2021/11/27 15:20:58 API Endpoint: /api
2021/11/27 15:20:59 Running crawler...
```

Agora basta acessar <http://127.0.0.1:8000/api>.
Caso tenha alterado o endpoint favor verifique o log de saida do terminal ou o arquivo config.yml.

β οΈοΈ **NΓ£o esqueΓ§a de configurar o arquivo *config.yml* conforme as suas necessidades!** β οΈ

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
  host: # Server host IP
  port: '8000' # Listen Port

api:
  endpoint: /api # api endpoint

db:
  file_name: database.db # sqlite file name

cache:
  max_age: 12  # Refresh cache every 12 hours

currency:
  currency: BRL # Don't change
```

**π·οΈ Para mais informaΓ§Γ΅es sobre spiders disponΓ­veis de uma olhada em
[/spiders](https://github.com/bitebait/curry/tree/master/crawler/spiders) e [AllSpiders()](https://github.com/bitebait/curry/blob/master/crawler/spiders/spiders.go)**.

* * *
<br>

## β **Lista de lojas monitoradas**

<br>

1. **[alboradainfo](https://www.alboradainfo.com/)**
1. **[atacadocollections](https://www.atacadocollections.com/)**
1. **[atacadogames](https://www.atacadogames.com/)**
1. **[atlanticoshop](https://www.atlanticoshop.com.br/)**
1. **[audiumelectronics](https://www.audiumelectronics.com/home)**
1. **[cellshop](https://www.cellshop.com/br/)**
1. **[eleganciacompany](https://www.eleganciacompany.com/)**
1. **[madridcenter](https://www.madridcenter.com/)**
1. **[megaeletronicos](https://www.megaeletronicos.com/br)**
1. **[mobilezone](https://www.mobilezone.com.br/)**
1. **[oneclick](https://oneclick.com.py/)**
1. **[pioneerinter](https://www.pioneerinter.com/)**
1. **[shoppingcentropioneer](https://shoppingcentropioneer.com/)**
1. **[tecombras](https://www.tecombras.net/)**
1. **[topdek](https://www.topdek.com.br/br)**
1. **[victoriastore](https://www.victoriastore.com.br/)**
1. **[visaovip](http://www.visaovip.com/)**
1. **[pontocom](https://www.pontocom.com/)**
1. **[cambioschaco](https://www.cambioschaco.com.py/pt-br/)**
1. **[bonanzacambios](https://bonanzacambios.com.py/)**
1. **[mercosurcambios](https://site.mercosurcambios.com/)**
1. **[dolarpy](https://www.dolarpy.com.br/)**

<br>

### π ExoneraΓ§Γ£o de responsabilidade e problemas conhecidos

* Qualquer uso do script Γ© de responsabilidade apenas do usuΓ‘rio. Os usuΓ‘rios do script devem agir de acordo com os
  termos dos sites acessados.
* Como acontece com todos os sites, a estrutura do site pode mudar no futuro e, portanto, como costuma acontecer com
  scripts de scraping, descontinue-o. NΓ£o Γ© realmente uma questΓ£o de saber se o cΓ³digo-fonte do site irΓ‘ mudar, mas sim
  quando (entΓ£o aproveite enquanto ainda estΓ‘ funcionando)

<br>

### π LicenΓ§a

Todo o conteΓΊdo apresentado nos sites pertence aos criadores originais.

A licenΓ§a abaixo se refere apenas ao script e nΓ£o ao conteΓΊdo scrapado.

[LicenΓ§a - MIT](https://github.com/bitebait/curry/blob/master/LICENSE)

<br>

### π₯ Sinta-se Γ  vontade para contribuir com o cΓ³digo (; π₯
