
package tests

import (
	"fmt"
	"os"
	"testing"
	"github.com/voidwyrm-2/gust/internal/lexer"
)

func TestLexer(t *testing.T) {
	data, err := os.ReadFile("../examples/hello_world.gt")
	if err != nil {
		t.Fatalf("Could not read example file: %v", err)
	}

	input := string(data)
	l := lexer.New(input)

	expectedTokens := []lexer.Token{
		{Type: lexer.IDENT, Literal: "println"},
		{Type: lexer.LEFT_PAREN, Literal: "("},
		{Type: lexer.STRING, Literal: "Hello World"},
		{Type: lexer.RIGHT_PAREN, Literal: ")"},
		{Type: lexer.EOF, Literal: ""},
	}


	for i, expected := range expectedTokens {
		tok := l.NextToken()
		if tok.Type != expected.Type {
			t.Fatalf("Token type mismatch at index %d: expected %v, got %v", i, expected.Type, tok.Type)
		}
		if tok.Literal != expected.Literal {
			t.Fatalf("Token literal mismatch at index %d: expected %q, got %q", i, expected.Literal, tok.Literal)
		}

		val := fmt.Sprintf("type: %d and literal: %s", tok.Type, tok.Literal)
		fmt.Println(val)
	}
	//  this will run all the tokens if you want to comment out the test and uncomment this
<<<<<<< Updated upstream
	 for {
     tok := l.NextToken()
     if tok.Type == lexer.EOF {
         break
     }
     val := fmt.Sprintf("type: %d literal: %s", tok.Type, tok.Literal)
     fmt.Println(val)
  }
}

=======
	// for {
  //   tok := l.NextToken()
  //   if tok.Type == lexer.EOF {
  //       break
  //   }
  //   val := fmt.Sprintf("type: %d literal: %s", tok.Type, tok.Literal)
  //   fmt.Println(val)
	// }
}
<<<<<<< Updated upstream
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
