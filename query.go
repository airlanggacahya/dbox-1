package dbox

type QueryPartType string

const (
	QueryPartSelect  = "SELECT"
	QueryPartFrom    = "FROM"
	QueryPartWhere   = "WHERE"
	QueryPartGroup   = "GROUP BY"
	QueryPartOrder   = "ORDER BY"
	QueryPartInsert  = "INSERT"
	QueryPartUpdate  = "UPDATE"
	QueryPartDelete  = "DELETE"
	QueryPartSave    = "SAVE"
	QueryPartCommand = "COMMAND"

	QueryPartJoin      = "JOIN"
	QueryPartLeftJoin  = "LEFT JOIN"
	QueryPartRightJoin = "RIGHT JOIN"
)

type IQuery interface {
	Cursor() (*Cursor, error)
	Connection() IConnection

	SetConnection(IConnection) IQuery
	SetThis(IQuery) IQuery

	Select(...string) IQuery
	From(string) IQuery
	Where(...*Filter) IQuery
}

type QueryPart struct {
	PartType string
	Value    interface{}
}

type Query struct {
	thisQuery IQuery
	conn      IConnection

	Parts []*QueryPart
}

func (q *Query) this() IQuery {
	if q.thisQuery == nil {
		return q
	} else {
		return q.thisQuery
	}
}

func (q *Query) addPart(qp *QueryPart) IQuery {
	if q.Parts == nil {
		q.Parts = []*QueryPart{}
	}
	q.Parts = append(q.Parts, qp)
	return q.this()
}

func (q *Query) SetConnection(c IConnection) IQuery {
	q.conn = c
	return q.this()
}

func (q *Query) SetThis(t IQuery) IQuery {
	q.thisQuery = t
	return t
}

func (q *Query) Connection() IConnection {
	return q.conn
}

func (q *Query) Cursor() (*Cursor, error) {
	return nil, nil
}

func (q *Query) Select(ss ...string) IQuery {
	q.addPart(&QueryPart{QueryPartSelect, ss})
	return q.this()
}

func (q *Query) From(objname string) IQuery {
	q.addPart(&QueryPart{QueryPartFrom, objname})
	return q.this()
}

func (q *Query) Where(fs ...*Filter) IQuery {
	q.addPart(&QueryPart{QueryPartWhere, fs})
	return q.this()
}

func (q *Query) OrderBy(ords ...string) IQuery {
	q.addPart(&QueryPart{QueryPartOrder, ords})
	return q.this()
}
