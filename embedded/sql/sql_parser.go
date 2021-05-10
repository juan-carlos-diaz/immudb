// Code generated by goyacc -l -o sql_parser.go sql_grammar.y. DO NOT EDIT.
package sql

import __yyfmt__ "fmt"

func setResult(l yyLexer, stmts []SQLStmt) {
	l.(*lexer).result = stmts
}

type yySymType struct {
	yys      int
	stmts    []SQLStmt
	stmt     SQLStmt
	colsSpec []*ColSpec
	colSpec  *ColSpec
	cols     []*ColSelector
	rows     []*RowSpec
	row      *RowSpec
	values   []ValueExp
	value    ValueExp
	id       string
	number   uint64
	str      string
	boolean  bool
	blob     []byte
	sqlType  SQLValueType
	aggFn    AggregateFn
	ids      []string
	col      *ColSelector
	sel      Selector
	sels     []Selector
	distinct bool
	ds       DataSource
	tableRef *TableRef
	joins    []*JoinSpec
	join     *JoinSpec
	joinType JoinType
	boolExp  ValueExp
	err      error
	ordcols  []*OrdCol
	opt_ord  Comparison
	logicOp  LogicOperator
	cmpOp    CmpOperator
	numOp    NumOperator
}

const CREATE = 57346
const USE = 57347
const DATABASE = 57348
const SNAPSHOT = 57349
const SINCE = 57350
const UP = 57351
const TO = 57352
const TABLE = 57353
const INDEX = 57354
const ON = 57355
const ALTER = 57356
const ADD = 57357
const COLUMN = 57358
const PRIMARY = 57359
const KEY = 57360
const BEGIN = 57361
const TRANSACTION = 57362
const COMMIT = 57363
const INSERT = 57364
const UPSERT = 57365
const INTO = 57366
const VALUES = 57367
const SELECT = 57368
const DISTINCT = 57369
const FROM = 57370
const BEFORE = 57371
const TX = 57372
const JOIN = 57373
const HAVING = 57374
const WHERE = 57375
const GROUP = 57376
const BY = 57377
const LIMIT = 57378
const ORDER = 57379
const ASC = 57380
const DESC = 57381
const AS = 57382
const NOT = 57383
const LIKE = 57384
const IF = 57385
const EXISTS = 57386
const NULL = 57387
const JOINTYPE = 57388
const NUMOP = 57389
const LOP = 57390
const CMPOP = 57391
const IDENTIFIER = 57392
const TYPE = 57393
const NUMBER = 57394
const VARCHAR = 57395
const BOOLEAN = 57396
const BLOB = 57397
const AGGREGATE_FUNC = 57398
const ERROR = 57399
const STMT_SEPARATOR = 57400

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"CREATE",
	"USE",
	"DATABASE",
	"SNAPSHOT",
	"SINCE",
	"UP",
	"TO",
	"TABLE",
	"INDEX",
	"ON",
	"ALTER",
	"ADD",
	"COLUMN",
	"PRIMARY",
	"KEY",
	"BEGIN",
	"TRANSACTION",
	"COMMIT",
	"INSERT",
	"UPSERT",
	"INTO",
	"VALUES",
	"SELECT",
	"DISTINCT",
	"FROM",
	"BEFORE",
	"TX",
	"JOIN",
	"HAVING",
	"WHERE",
	"GROUP",
	"BY",
	"LIMIT",
	"ORDER",
	"ASC",
	"DESC",
	"AS",
	"NOT",
	"LIKE",
	"IF",
	"EXISTS",
	"NULL",
	"JOINTYPE",
	"NUMOP",
	"LOP",
	"CMPOP",
	"IDENTIFIER",
	"TYPE",
	"NUMBER",
	"VARCHAR",
	"BOOLEAN",
	"BLOB",
	"AGGREGATE_FUNC",
	"ERROR",
	"','",
	"'.'",
	"STMT_SEPARATOR",
	"'('",
	"')'",
	"'@'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 229

var yyAct = [...]int{

	196, 36, 55, 143, 123, 6, 125, 142, 100, 71,
	62, 91, 72, 126, 7, 85, 128, 135, 51, 182,
	188, 187, 133, 181, 129, 130, 131, 132, 37, 135,
	149, 150, 151, 127, 165, 134, 129, 130, 131, 132,
	38, 177, 159, 47, 49, 175, 106, 134, 52, 106,
	107, 48, 76, 105, 58, 117, 156, 3, 113, 77,
	97, 156, 73, 144, 155, 81, 79, 68, 66, 57,
	19, 98, 67, 58, 195, 186, 53, 96, 162, 95,
	115, 38, 174, 138, 88, 52, 94, 37, 2, 89,
	124, 191, 17, 104, 18, 20, 103, 83, 149, 150,
	151, 116, 9, 38, 110, 112, 54, 149, 32, 151,
	33, 179, 157, 35, 119, 137, 101, 114, 101, 102,
	86, 136, 87, 78, 139, 75, 48, 61, 145, 59,
	48, 153, 154, 46, 43, 39, 149, 93, 161, 80,
	41, 70, 152, 141, 60, 74, 197, 198, 167, 56,
	170, 164, 168, 184, 171, 172, 173, 185, 148, 122,
	109, 176, 4, 178, 147, 111, 180, 82, 64, 63,
	23, 9, 12, 13, 120, 118, 31, 30, 69, 21,
	160, 84, 14, 12, 13, 65, 190, 193, 194, 189,
	15, 16, 45, 14, 158, 50, 42, 199, 8, 29,
	200, 15, 16, 24, 140, 9, 27, 28, 25, 26,
	40, 166, 192, 183, 121, 146, 108, 92, 90, 44,
	22, 34, 163, 169, 99, 11, 10, 5, 1,
}
var yyPact = [...]int{

	-3, -1000, 179, -3, -1000, 10, -3, -1000, 159, 143,
	-1000, -1000, 197, 200, 188, 153, 152, -1000, -1000, -3,
	-1000, -3, 31, -1000, 85, 97, 183, 84, 184, 83,
	80, 80, 179, 168, 48, 109, -1000, 8, 14, -1000,
	79, 103, 77, -1000, 140, 138, 170, 7, 13, 6,
	-1000, 157, -3, 1, 31, -1000, 75, -10, 73, 5,
	95, 4, -1000, 137, 45, 165, 70, 72, 70, -1000,
	168, 91, -1000, 76, 109, -1000, -1000, -2, 12, 68,
	-1000, 69, 44, -1000, 68, -9, -1000, -1000, -12, -1000,
	127, -1000, 91, 134, 140, -4, -1000, -1000, 67, 22,
	-1000, 50, -7, -1000, -1000, 150, 64, 149, 125, -28,
	-1000, 1, 109, -1000, -1000, 66, 102, -1000, 2, -1000,
	2, 132, 123, 51, 100, -1000, -28, -28, 3, -1000,
	-1000, -1000, -1000, -5, 62, -1000, 181, -20, 162, -1000,
	-1000, 93, 20, -1000, -16, 20, 111, -28, 53, -28,
	-28, -28, 29, 60, -17, 145, -21, -1000, -28, -1000,
	61, -1000, 2, -39, -1000, 0, 117, 122, 51, 17,
	-1000, -1000, 60, 89, -1000, -1000, -41, -1000, 51, -42,
	-1000, -1000, -16, 109, 39, 53, 53, -1000, -1000, -1000,
	-1000, -1000, 16, 108, -1000, 53, -1000, -1000, -1000, 108,
	-1000,
}
var yyPgo = [...]int{

	0, 228, 162, 18, 227, 14, 226, 225, 5, 224,
	8, 15, 223, 7, 3, 222, 6, 90, 221, 1,
	220, 9, 12, 219, 10, 218, 11, 217, 4, 216,
	215, 214, 213, 2, 212, 211, 0, 210, 204, 88,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 39, 39, 4, 4, 5,
	5, 3, 3, 6, 6, 6, 6, 6, 6, 23,
	23, 37, 37, 7, 7, 13, 13, 14, 11, 11,
	12, 12, 15, 15, 16, 16, 16, 16, 16, 16,
	16, 9, 9, 10, 38, 38, 8, 20, 20, 18,
	18, 17, 17, 17, 19, 19, 19, 21, 21, 21,
	22, 22, 24, 24, 25, 25, 26, 26, 27, 29,
	29, 31, 31, 30, 30, 32, 32, 35, 35, 34,
	34, 36, 36, 36, 33, 33, 28, 28, 28, 28,
	28, 28, 28, 28, 28,
}
var yyR2 = [...]int{

	0, 2, 2, 2, 4, 0, 2, 1, 5, 1,
	1, 2, 3, 3, 3, 4, 11, 7, 6, 0,
	3, 0, 3, 8, 8, 1, 3, 3, 1, 3,
	1, 3, 1, 3, 1, 1, 1, 1, 3, 2,
	1, 1, 3, 3, 0, 2, 12, 0, 1, 2,
	4, 1, 3, 4, 1, 3, 5, 1, 5, 3,
	1, 3, 0, 3, 0, 1, 1, 2, 5, 0,
	2, 0, 3, 0, 2, 0, 2, 0, 3, 2,
	4, 0, 1, 1, 0, 2, 1, 1, 3, 2,
	3, 3, 3, 3, 4,
}
var yyChk = [...]int{

	-1000, -1, -39, 60, -2, -4, -8, -5, 19, 26,
	-6, -7, 4, 5, 14, 22, 23, -39, -39, 60,
	-39, 20, -20, 27, 6, 11, 12, 6, 7, 11,
	24, 24, -39, -39, -18, -17, -19, 56, 50, 50,
	-37, 43, 13, 50, -23, 8, 50, -22, 50, -22,
	-2, -3, -5, 28, 58, -33, 40, 61, 59, 50,
	41, 50, -24, 29, 30, 15, 61, 59, 61, 21,
	-39, -21, -22, 61, -17, 50, 62, -19, 50, 61,
	44, 61, 30, 52, 16, -11, 50, 50, -11, -3,
	-25, -26, -27, 46, -22, -8, -33, 62, 59, -9,
	-10, 50, 50, 52, -10, 62, 58, 62, -29, 33,
	-26, 31, -24, 62, 50, 58, 51, 62, 25, 50,
	25, -31, 34, -28, -17, -16, 41, 61, 44, 52,
	53, 54, 55, 50, 63, 45, -21, -33, 17, -10,
	-38, 41, -13, -14, 61, -13, -30, 32, 35, 47,
	48, 49, 42, -28, -28, 61, 61, 50, 13, 62,
	18, 45, 58, -15, -16, 50, -35, 37, -28, -12,
	-19, -28, -28, -28, 53, 62, -8, 62, -28, 50,
	-14, 62, 58, -32, 36, 35, 58, 62, 62, -16,
	-33, 52, -34, -19, -19, 58, -36, 38, 39, -19,
	-36,
}
var yyDef = [...]int{

	5, -2, 0, 5, 1, 5, 5, 7, 0, 47,
	9, 10, 0, 0, 0, 0, 0, 6, 2, 5,
	3, 5, 0, 48, 0, 21, 0, 0, 19, 0,
	0, 0, 6, 0, 0, 84, 51, 0, 54, 13,
	0, 0, 0, 14, 62, 0, 0, 0, 60, 0,
	4, 0, 5, 0, 0, 49, 0, 0, 0, 0,
	0, 0, 15, 0, 0, 0, 0, 0, 0, 8,
	11, 64, 57, 0, 84, 85, 52, 0, 55, 0,
	22, 0, 0, 20, 0, 0, 28, 61, 0, 12,
	69, 65, 66, 0, 62, 0, 50, 53, 0, 0,
	41, 0, 0, 63, 18, 0, 0, 0, 71, 0,
	67, 0, 84, 59, 56, 0, 44, 17, 0, 29,
	0, 73, 0, 70, 86, 87, 0, 0, 0, 34,
	35, 36, 37, 54, 0, 40, 0, 0, 0, 42,
	43, 0, 23, 25, 0, 24, 77, 0, 0, 0,
	0, 0, 0, 89, 0, 0, 0, 39, 0, 58,
	0, 45, 0, 0, 32, 0, 75, 0, 74, 72,
	30, 88, 92, 93, 91, 90, 0, 38, 68, 0,
	26, 27, 0, 84, 0, 0, 0, 94, 16, 33,
	46, 76, 78, 81, 31, 0, 79, 82, 83, 81,
	80,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	61, 62, 3, 3, 58, 3, 59, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 63,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 60,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = yyDollar[2].stmts
			setResult(yylex, yyDollar[2].stmts)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 4:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmts = append([]SQLStmt{yyDollar[1].stmt}, yyDollar[4].stmts...)
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
		}
	case 8:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &TxStmt{stmts: yyDollar[4].stmts}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmts = append([]SQLStmt{yyDollar[1].stmt}, yyDollar[3].stmts...)
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &CreateDatabaseStmt{DB: yyDollar[3].id}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &UseDatabaseStmt{DB: yyDollar[3].id}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &UseSnapshotStmt{sinceTx: yyDollar[3].number, asBefore: yyDollar[4].number}
		}
	case 16:
		yyDollar = yyS[yypt-11 : yypt+1]
		{
			yyVAL.stmt = &CreateTableStmt{ifNotExists: yyDollar[3].boolean, table: yyDollar[4].id, colsSpec: yyDollar[6].colsSpec, pk: yyDollar[10].id}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{table: yyDollar[4].id, col: yyDollar[6].id}
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &AddColumnStmt{table: yyDollar[3].id, colSpec: yyDollar[6].colSpec}
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[3].number
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{isInsert: true, tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.rows = []*RowSpec{yyDollar[1].row}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.rows = append(yyDollar[1].rows, yyDollar[3].row)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.row = &RowSpec{Values: yyDollar[2].values}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = append(yyDollar[1].ids, yyDollar[3].id)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.cols = []*ColSelector{yyDollar[1].col}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = append(yyDollar[1].cols, yyDollar[3].col)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = []ValueExp{yyDollar[1].value}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].value)
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Number{val: yyDollar[1].number}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Varchar{val: yyDollar[1].str}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Bool{val: yyDollar[1].boolean}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Blob{val: yyDollar[1].blob}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.value = &SysFn{fn: yyDollar[1].id}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.value = &Param{id: yyDollar[2].id}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &NullValue{}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.colsSpec = []*ColSpec{yyDollar[1].colSpec}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.colsSpec = append(yyDollar[1].colsSpec, yyDollar[3].colSpec)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.colSpec = &ColSpec{colName: yyDollar[1].id, colType: yyDollar[2].sqlType, notNull: yyDollar[3].boolean}
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 46:
		yyDollar = yyS[yypt-12 : yypt+1]
		{
			yyVAL.stmt = &SelectStmt{
				distinct:  yyDollar[2].distinct,
				selectors: yyDollar[3].sels,
				ds:        yyDollar[5].ds,
				joins:     yyDollar[6].joins,
				where:     yyDollar[7].boolExp,
				groupBy:   yyDollar[8].cols,
				having:    yyDollar[9].boolExp,
				orderBy:   yyDollar[10].ordcols,
				limit:     yyDollar[11].number,
				as:        yyDollar[12].id,
			}
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.distinct = false
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.distinct = true
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[1].sel.setAlias(yyDollar[2].id)
			yyVAL.sels = []Selector{yyDollar[1].sel}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[3].sel.setAlias(yyDollar[4].id)
			yyVAL.sels = append(yyDollar[1].sels, yyDollar[3].sel)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sel = yyDollar[1].col
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, col: "*"}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, db: yyDollar[3].col.db, table: yyDollar[3].col.table, col: yyDollar[3].col.col}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.col = &ColSelector{col: yyDollar[1].id}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.col = &ColSelector{table: yyDollar[1].id, col: yyDollar[3].id}
		}
	case 56:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.col = &ColSelector{db: yyDollar[1].id, table: yyDollar[3].id, col: yyDollar[5].id}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ds = yyDollar[1].tableRef
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyDollar[2].tableRef.asBefore = yyDollar[3].number
			yyDollar[2].tableRef.as = yyDollar[4].id
			yyVAL.ds = yyDollar[2].tableRef
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ds = yyDollar[2].stmt.(*SelectStmt)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.tableRef = &TableRef{table: yyDollar[1].id}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.tableRef = &TableRef{db: yyDollar[1].id, table: yyDollar[3].id}
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[3].number
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joins = nil
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = yyDollar[1].joins
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = []*JoinSpec{yyDollar[1].join}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.joins = append([]*JoinSpec{yyDollar[1].join}, yyDollar[2].joins...)
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.join = &JoinSpec{joinType: yyDollar[1].joinType, ds: yyDollar[3].ds, cond: yyDollar[5].boolExp}
		}
	case 69:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolExp = nil
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[2].boolExp
		}
	case 71:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.cols = nil
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = yyDollar[3].cols
		}
	case 73:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolExp = nil
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[2].boolExp
		}
	case 75:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.number = yyDollar[2].number
		}
	case 77:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ordcols = nil
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ordcols = yyDollar[3].ordcols
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.ordcols = []*OrdCol{{sel: yyDollar[1].col, cmp: yyDollar[2].opt_ord}}
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ordcols = append(yyDollar[1].ordcols, &OrdCol{sel: yyDollar[3].col, cmp: yyDollar[4].opt_ord})
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.opt_ord = GreaterOrEqualTo
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = GreaterOrEqualTo
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = LowerOrEqualTo
		}
	case 84:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.id = ""
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.id = yyDollar[2].id
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[1].sel
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[1].value
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = &NumExp{left: yyDollar[1].boolExp, op: yyDollar[2].numOp, right: yyDollar[3].boolExp}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolExp = &NotBoolExp{exp: yyDollar[2].boolExp}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[2].boolExp
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = &LikeBoolExp{sel: yyDollar[1].sel, pattern: yyDollar[3].str}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = &BinBoolExp{op: yyDollar[2].logicOp, left: yyDollar[1].boolExp, right: yyDollar[3].boolExp}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = &CmpBoolExp{op: yyDollar[2].cmpOp, left: yyDollar[1].boolExp, right: yyDollar[3].boolExp}
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.boolExp = &ExistsBoolExp{q: (yyDollar[3].stmt).(*SelectStmt)}
		}
	}
	goto yystack /* stack new state and value */
}