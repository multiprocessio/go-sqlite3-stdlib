package stdlib

import (
	"database/sql"

	"github.com/mattn/go-sqlite3"
)

var functions = map[string]any{
	"regexp": ext_regexp,

	// math extension
	"acos": math.Acos,
	"acosh": math.Acosh,
	"asin": math.Asin,
	"asinh": math.Asinh,
	"atan": math.Atan,
	// TODO: atan2
	"atanh": math.Atanh,
	"ceil": math.Ceil,
	"ceiling": math.Ceil,
	"cos": math.Cos,
	"cosh": math.Cosh,
	"degrees": ext_degrees,
	"exp": math.Exp,
	"floor": math.Floor,
	"ln": math.Lb,
	"log": math.Log,
	"log10": math.Log10,
	// TODO: support log(B, X)
	"log2": math.Log2,
	"mod": ext_modulus,
	"pi": ext_pi,
	"pow": math.Pow,
	"power": math.Pow,
	"radians": ext_radians,
	"sin": math.Sin,
	"sinh": math.Sinh,
	"sqrt": math.Sqrt,
	"tan": math.Tan,
	"tanh": math.Tanh,
	"trunc": math.Round,

	// generate_series extension
	"generate_series": ext_generate_series,

	// string
	"repeat": ext_repeat,
	"replicate": ext_repeat,
	"strpos": ext_charindex,
	"charindex": ext_charindex,
	"ltrim": ext_ltrim,
	"rtrim": ext_rtrim,
	"trim": ext_ltrim,
	"replace": ext_replace,
	"reverse": ext_reverse,
	"proper": ext_proper,
	"padl": ext_padl,
	"padr": ext_padr,
	"padc": ext_padc,
	"strfilter": ext_strfilter,

}

func init() {
	sql.Register("sqlite3_extended",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				 conn.RegisterFunc("regexp", regex, true)
			},
		})
}
