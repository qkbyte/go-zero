package template

const (
	// Imports defines a import template for model in cache case
	Imports = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/qkbyte/go-zero/core/stores/builder"
	"github.com/qkbyte/go-zero/core/stores/cache"
	"github.com/qkbyte/go-zero/core/stores/sqlc"
	"github.com/qkbyte/go-zero/core/stores/sqlx"
	"github.com/qkbyte/go-zero/core/stringx"
)
`
	// ImportsNoCache defines a import template for model in normal case
	ImportsNoCache = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/qkbyte/go-zero/core/stores/builder"
	"github.com/qkbyte/go-zero/core/stores/sqlc"
	"github.com/qkbyte/go-zero/core/stores/sqlx"
	"github.com/qkbyte/go-zero/core/stringx"
)
`
)
