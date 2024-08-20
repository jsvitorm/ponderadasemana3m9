# TDD em GoLang com GIN

## Testes em Go

Em Go, os testes são uma parte crucial do desenvolvimento de software de alta qualidade. A linguagem suporta testes automatizados diretamente com sua ferramenta de teste embutida, que permite aos desenvolvedores escrever, executar e verificar testes com facilidade. Para organizar os testes, geralmente criamos arquivos de teste com o sufixo `_test.go`.

## Desenvolvimento Orientado a Testes (TDD)

O TDD é uma abordagem de desenvolvimento de software onde os testes são escritos antes do código funcional. Este processo envolve três etapas principais:

1. **Red (Vermelho)**: Escrever um teste que falhe inicialmente, pois a funcionalidade ainda não foi implementada.
2. **Green (Verde)**: Escrever o código mínimo necessário para fazer o teste passar.
3. **Refactor (Refatorar)**: Melhorar o código garantindo que os testes continuem passando.

## Aplicação do TDD em Go

Aplicar TDD em Go envolve criar funções de teste que verificam se o código atende aos requisitos especificados. Vamos passar pelo ciclo TDD (Red, Green, Refactor) para criar e testar uma função de conexão ao banco de dados usando Gorm e um handler HTTP simples usando Gin.

### Red: Escrever testes que falhem

Primeiro, escrevemos os testes que esperamos que falhem, pois as funções ainda não foram completamente implementadas.

**Teste para o handler HTTP `GetGreeting`:**

```go
package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/greeting", GetGreeting)

	// Execute request
	req, _ := http.NewRequest("GET", "/greeting", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Hello, World!"}`, w.Body.String())

}
```


Quando rodamos esse teste, ele falhará por não ter sua implementação ou por estar incompleta

IMGs


### Green: Código necessário para o teste funcionar

Agora escrevemos o código que passará nos testes

```


func GetGreeting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func ConnectDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

```

imgs

Com esse código conseguiremos passar nos testes


### Refactor: Melhorar o código mantendo os testes passando
Com os testes passando, podemos refatorar o código se necessário. Por exemplo, poderíamos melhorar a configuração do banco de dados ou a estrutura dos handlers para maior clareza e manutenção. No entanto, no exemplo atual, o código já está relativamente simples e claro.

## Conclusão
Seguindo o ciclo TDD (Red, Green, Refactor), garantimos que cada nova funcionalidade é bem testada e que o código é continuamente melhorado para manter a qualidade e legibilidade. Essa prática ajuda a prevenir bugs e facilita a manutenção e evolução do software.


