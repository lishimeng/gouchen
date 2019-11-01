package connrepo

import (
	"fmt"
	"github.com/lishimeng/gouchen/internal/connector"
)

type Repository struct {
	connectors map[string]*connector.Connector
	name2id    map[string]string
}

func New() *Repository {
	r := Repository{
		connectors: make(map[string]*connector.Connector),
		name2id:    make(map[string]string),
	}

	return &r
}

func (r Repository) Register(c *connector.Connector) {
	id := (*c).GetID()
	name := (*c).GetName()
	r.connectors[id] = c
	r.name2id[name] = id
}

func (r Repository) GetByID(id string) (c *connector.Connector, err error) {
	c, ok := r.connectors[id]
	if !ok {
		err = fmt.Errorf("no such connector id[%s]", id)
	}
	return c, err
}

func (r Repository) GetByName(name string) (c *connector.Connector, err error) {
	id, ok := r.name2id[name]
	if !ok {
		err = fmt.Errorf("no such connector name[%s]", name)
	} else {
		c, err = r.GetByID(id)
		if err != nil {
			err = fmt.Errorf("no such connector name[%s]", name)
		}
	}
	return c, err
}