package bank

import (
	"encoding/json"
	"time"
	"io/ioutil"
	"os"
	"errors"
	"fmt"
	"github.com/nu7hatch/gouuid"
)

const rwx_w__w_ = 0644

func newId() string {
	u4, err := uuid.NewV4()
	if err == nil {
		return fmt.Sprintf("id-", u4)
	} else {
		return "BAD GUID CREATION!!!"
	}
}

type Pred func(item interface{}) bool
type Mapper func(index int, item interface{}) interface{}
type Copier func(a interface{}) interface{}

type Bank struct {
	name string
	filename string
	rolls map[string][]interface{}
	createdOn time.Time
	updatedOn time.Time
}

type unexportedBank struct {
	Name string							`json:"name"`
	Filename string						`json:"filename"`
	Rolls map[string][]interface{}		`json:"rolls"`
	CreatedOn time.Time					`json:"createdOn"`
	UpdatedOn time.Time					`json:"updatedOn"`
}

func (b *Bank) Length(id string) int {
	return len(b.rolls)
}

func (b *Bank) preMarshal() *unexportedBank {
	return &unexportedBank{
		Name:b.name,
		Filename:b.filename,
		Rolls:b.rolls,
		CreatedOn:b.createdOn,
		UpdatedOn:b.updatedOn,
	}
}

func (b *unexportedBank) postUnmarshal() *Bank {
	return &Bank{
		name:b.Name,
		filename:b.Filename,
		rolls:b.Rolls,
		createdOn:b.CreatedOn,
		updatedOn:b.UpdatedOn,
	}
}

func (b *Bank) Register(roll RollDef) (*Roll, error) {
 	if roll.Name == "" {
		return nil, errors.New(
			"Cannot register a roll definition without a 'Name'")
	}
	r := &Roll{
		name: roll.Name,
		initialSize: roll.InitialSize,
		bank:b,
		id:newId(),
		items:make([]interface{}, 0),
	}
	return r, nil
}

func NewBank(name string) *Bank {
	return &Bank{
		name:name,
		rolls: make(map[string][]interface{}, 0),
		createdOn:time.Now(),
		updatedOn:time.Now(),
	}
}

func LoadBank(file string) (*Bank, error) {
	bytes,err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	bank := &unexportedBank{}
	err = json.Unmarshal(bytes, bank)
	if err != nil {
		return nil, err
	}

	return bank.postUnmarshal(), nil
}

func (b *Bank) SaveToFile(file string) error {
	b.filename = file
	pre := b.preMarshal()
	bytes, err := json.MarshalIndent(pre, "", "  ")
	if err == nil {
		return ioutil.WriteFile(file, bytes, rwx_w__w_)
	} else {
		return err
	}
}

func (b *Bank) Name() string { return b.name }
func (b *Bank) CreatedOn() time.Time { return b.createdOn }
func (b *Bank) UpdatedOn() time.Time { return b.updatedOn }

func fileExists(file string) (bool, error) {
	_, err := os.Stat(file);

	if os.IsNotExist(err) {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

