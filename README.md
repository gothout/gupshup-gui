# 🧱 Arquitetura do Projeto Gupshup-GUI

Este projeto segue uma arquitetura **limpa, escalável e modular** baseada em princípios de *Domain-Driven Design (DDD)* e práticas modernas de desenvolvimento em Go.

---

## 📐 Camadas da Arquitetura

### 1. `Handler` (Entrada HTTP)

Responsável por:

- Interpretar a requisição vinda do `Gin`.
- Realizar o bind e validação dos dados de entrada (query, path, body).
- Invocar o Controller correto.
- Retornar a resposta apropriada (status + body).

**Exemplo:**  
```go
group.GET("/apps/:app_id/templates", templateHandler.GetTemplatesHandler)
````

---

### 2. `Controller` (Orquestração)

Responsável por:

* Receber os dados processados do `Handler`.
* Realizar validações de alto nível e regras de orquestração.
* Encaminhar chamadas ao(s) `Service(s)` necessário(s).
* Ser desacoplado do contexto HTTP (não usa `gin.Context`).

**Exemplo:**

```go
func (c *templateControllerImpl) GetTemplates(appID string) ([]template.PartnerTemplate, error)
```

---

### 3. `Service` (Regras de Negócio)

Responsável por:

* Conter toda a lógica de negócio.
* Tratar cache, autenticação, chamadas externas, parse de dados, etc.
* Ser desacoplado de qualquer framework (puro Go).

**Exemplo:**

```go
func (s *templateServiceImpl) GetTemplates(appID string) ([]template.PartnerTemplate, error)
```

---

### 4. `Model` (Entidades e DTOs)

Responsável por:

* Representar as estruturas de dados.
* Definir `structs` que espelham JSONs da Gupshup ou entidades internas.
* Contar com funções auxiliares como `MarshalJSON()` quando necessário.

**Exemplo:**

```go
type PartnerTemplate struct {
	ID   string `json:"id"`
	Meta json.RawMessage `json:"meta"` // via Marshal customizado
}
```

---

### 5. `Binding` (Validação de Entrada)

Responsável por:

* Isolar os parâmetros vindos de URL ou JSON.
* Validar os campos obrigatórios.
* Reduzir acoplamento entre HTTP e regras de negócio.

**Exemplo:**

```go
type AppIDInput struct {
	AppID string `json:"appid" binding:"required"`
}
```

---

## 📦 Organização de Pastas

```
internal/
  app/
    controller/
    handler/
    service/
    model/
    binding/
package/
  configuration/
    rest_err/        // Padrão de erro uniforme
    config/          // URL e variáveis globais
```

---

## ✅ Boas Práticas Adotadas

* ✅ **Organização por camadas**: `Handler → Controller → Service → Model`.
* ✅ **Separação de responsabilidades**: cada camada possui responsabilidade única.
* ✅ **Autenticação cacheada** com `patrickmn/go-cache` e renovação automática.
* ✅ **Retry com tratamento de erro 429** (`Too Many Requests`) e reautenticação silenciosa.
* ✅ **Token por App com controle de acesso** e limitação de requisições (5 por appId).
* ✅ **Uso de `sync.RWMutex`** para evitar race conditions em acesso ao cache.
* ✅ **Custom `MarshalJSON`** para preservar a resposta original da API.
* ✅ **Leitura segura e decodificação defensiva de JSONs aninhados**.
* ✅ **Código desacoplado de Gin no Controller e Service**.
* ✅ **Mensagens de erro centralizadas e padronizadas via `rest_err`**.

---

## 📌 Exemplo de Fluxo Completo

```bash
GET /app/apps/abc123/templates
```

1. Handler recebe `abc123` da URL.
2. Faz `binding` e valida se é um appId válido.
3. Controller orquestra a requisição.
4. Service consulta o cache:

   * Se token do app for novo ou passou de 5 requisições, renova.
   * Se erro 429, força novo login no token principal.
5. Templates são decodificados e retornados.
6. Resposta final é idêntica à da Gupshup.

---

## 📌 Conclusão

A arquitetura foi pensada para:

* Crescer com múltiplos domínios (ex: campaign, bot, broadcast).
* Evitar duplicação de lógica e responsabilidade.
* Permitir testes de unidade em todas as camadas.
* Facilitar substituição futura de bibliotecas como `Gin`.

> Este padrão deve ser seguido por todos os módulos da aplicação.

---

```