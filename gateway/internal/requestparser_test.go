package internal

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/qkbyte/go-zero/rest/pathvar"
	"github.com/stretchr/testify/assert"
)

func TestNewRequestParserNoVar(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	parser, err := NewRequestParser(req, nil)
	assert.Nil(t, err)
	assert.NotNil(t, parser)
}

func TestNewRequestParserWithVars(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req = pathvar.WithVars(req, map[string]string{"a": "b"})
	parser, err := NewRequestParser(req, nil)
	assert.Nil(t, err)
	assert.NotNil(t, parser)
}

func TestNewRequestParserNoVarWithBody(t *testing.T) {
	req := httptest.NewRequest("GET", "/", strings.NewReader(`{"a": "b"}`))
	parser, err := NewRequestParser(req, nil)
	assert.Nil(t, err)
	assert.NotNil(t, parser)
}

func TestNewRequestParserWithVarsWithBody(t *testing.T) {
	req := httptest.NewRequest("GET", "/", strings.NewReader(`{"a": "b"}`))
	req = pathvar.WithVars(req, map[string]string{"c": "d"})
	parser, err := NewRequestParser(req, nil)
	assert.Nil(t, err)
	assert.NotNil(t, parser)
}

func TestNewRequestParserWithVarsWithWrongBody(t *testing.T) {
	req := httptest.NewRequest("GET", "/", strings.NewReader(`{"a": "b"`))
	req = pathvar.WithVars(req, map[string]string{"c": "d"})
	parser, err := NewRequestParser(req, nil)
	assert.NotNil(t, err)
	assert.Nil(t, parser)
}

func TestNewRequestParserWithForm(t *testing.T) {
	req := httptest.NewRequest("GET", "/val?a=b", nil)
	parser, err := NewRequestParser(req, nil)
	assert.Nil(t, err)
	assert.NotNil(t, parser)
}

func TestNewRequestParserWithBadForm(t *testing.T) {
	req := httptest.NewRequest("GET", "/val?a%1=b", nil)
	parser, err := NewRequestParser(req, nil)
	assert.NotNil(t, err)
	assert.Nil(t, parser)
}

func TestRequestParser_buildJsonRequestParser(t *testing.T) {
	parser, err := buildJsonRequestParser(map[string]interface{}{"a": make(chan int)}, nil)
	assert.NotNil(t, err)
	assert.Nil(t, parser)
}
