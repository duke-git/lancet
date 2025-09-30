// Copyright 2025 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

package enum

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

type Status int

const (
	Unknown Status = iota
	Active
	Inactive
)

func TestNewItem(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestNewItem")

	items := NewItems[Status](
		Active, "Active",
		Inactive, "Inactive",
	)

	assert.Equal(2, len(items))
	assert.Equal(Active, items[0].Value())
	assert.Equal("Active", items[0].Name())
	assert.Equal(Inactive, items[1].Value())
	assert.Equal("Inactive", items[1].Name())
}

func TestItem_Valid(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestItem_Valid")

	item := NewItem(Active, "Active")
	assert.Equal(true, item.Valid())

	invalidItem := NewItem(Unknown, "")
	assert.Equal(false, invalidItem.Valid())
}

func TestItem_MarshalJSON(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestItem_MarshalJSON")

	item := NewItem(Active, "Active")
	data, err := item.MarshalJSON()
	assert.IsNil(err)
	assert.Equal("{\"name\":\"Active\",\"value\":1}", string(data))

	var unmarshaledItem Item[Status]
	err = unmarshaledItem.UnmarshalJSON(data)
	assert.IsNil(err)
	assert.Equal(item.Value(), unmarshaledItem.Value())
	assert.Equal(item.Name(), unmarshaledItem.Name())
}

func TestRegistry_AddAndGet(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRegistry_AddAndGet")

	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")

	registry.Add(item1, item2)

	assert.Equal(2, registry.Size())

	item, ok := registry.GetByValue(Active)
	assert.Equal(true, ok)
	assert.Equal("Active", item.Name())

	item, ok = registry.GetByName("Inactive")
	assert.Equal(true, ok)
	assert.Equal(Inactive, item.Value())
}

func TestRegistry_Remove(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRegistry_Remove")

	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")

	registry.Add(item1, item2)
	assert.Equal(2, registry.Size())

	removed := registry.Remove(Active)
	assert.Equal(true, removed)
	assert.Equal(1, registry.Size())

	_, ok := registry.GetByValue(Active)
	assert.Equal(false, ok)
}

func TestRegistry_Update(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRegistry_Update")

	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	registry.Add(item1)

	updated := registry.Update(Active, "Activated")
	assert.Equal(true, updated)

	item, ok := registry.GetByValue(Active)
	assert.Equal(true, ok)
	assert.Equal("Activated", item.Name())
}

func TestRegistry_Contains(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRegistry_Contains")

	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	registry.Add(item1)

	assert.Equal(true, registry.Contains(Active))
	assert.Equal(false, registry.Contains(Inactive))
}

func TestRegistry_Validate(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRegistry_Validate")

	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")
	registry.Add(item1, item2)

	err := registry.Validate(Active)
	assert.IsNil(err)
	err = registry.Validate(Inactive)
	assert.IsNil(err)

	err = registry.Validate(Unknown)
	assert.IsNotNil(err)
}

func TestRegistry_ValidateAll(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRegistry_ValidateAll")

	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")
	registry.Add(item1, item2)

	err := registry.ValidateAll(Active, Inactive)
	assert.IsNil(err)

	err = registry.ValidateAll(Active, Unknown)
	assert.IsNotNil(err)
}
