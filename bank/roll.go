package bank
import (
	"time"
)

type RollDef struct {
	Name string
	Type string
	InitialSize int
	Copy Copier
}

type Roll struct {
	name string
	typename string
	initialSize int
	bank *Bank
	id string
	copy Copier
	items []interface{}
	updatedOn time.Time
	createdOn time.Time
}

func insert(offset int, items []interface{}, item interface{}) []interface{} {
	s := append(items, 0)
	copy(s[offset+1:], s[offset:])
	s[offset] = item
	return s
}

func (r *Roll) clone(items []interface{}) *Roll {
	return &Roll{
		name: r.name,
		typename: r.typename,
		initialSize: r.initialSize,
		bank: r.bank,
		id: r.id,
		copy: r.copy,
		items: items,
		updatedOn: time.Now(),
		createdOn: r.createdOn,
	}
}

func (r *Roll) Where(fn Pred) *Roll {
	res := make([]interface{}, 0)
	for _,e := range r.items {
		if fn(e) {
			res = append(res, e)
		}
	}
	return r.clone(res)
}

func (r *Roll) Find(fn Pred) (interface{},bool) {
	for _,e := range r.items {
		if fn(e) {
			return e, true
		}
	}
	return nil, false
}

func (r *Roll) Map(fn Mapper) *Roll {
	res := make([]interface{}, len(r.items))
	for i,e := range r.items {
		res[i] = fn(i,e)
	}
	return r.clone(res)
}

func (r *Roll) IsEmpty() bool {
	return r.items != nil && len(r.items) == 0
}

func (r *Roll) All(fn Pred) bool {
	hasAll := true
	for _,e := range r.items {
		hasAll = hasAll && fn(e)
		if !hasAll {
			return hasAll
		}
	}
	return hasAll
}

func (r *Roll) Any(fn Pred) bool {
	for _,e := range r.items {
		yes := fn(e)
		if yes {
			return yes
		}
	}
	return false
}

func (r *Roll) Clear() *Roll {
	return r.clone(make([]interface{}, 0))
}

func (r *Roll) Insert(index int, item interface{}) *Roll {
	return r.clone(insert(index, r.items, item))
}

func (r *Roll) Add(more ...interface{}) *Roll {
	return r.clone(append(r.items, more...))
}

func (r *Roll) Length() int { return len(r.items) }
func (r *Roll) Name() string { return r.name }
func (r *Roll) Type() string { return r.typename }
func (r *Roll) Size() int { return r.initialSize }
func (r *Roll) Bank() *Bank { return r.bank }

