package enum

import "fmt"

func ExampleNewItem() {
	items := NewItems(
		Pair[Status]{Value: Active, Name: "Active"},
		Pair[Status]{Value: Inactive, Name: "Inactive"},
	)

	fmt.Println(items[0].Name(), items[0].Value())
	fmt.Println(items[1].Name(), items[1].Value())

	// Output:
	// Active 1
	// Inactive 2
}

func ExampleItem_Valid() {
	item := NewItem(Active, "Active")
	fmt.Println(item.Valid())

	invalidItem := NewItem(Unknown, "")
	fmt.Println(invalidItem.Valid())

	// Output:
	// true
	// false
}

func ExampleItem_MarshalJSON() {
	item := NewItem(Active, "Active")
	data, _ := item.MarshalJSON()
	fmt.Println(string(data))

	var unmarshaledItem Item[Status]
	_ = unmarshaledItem.UnmarshalJSON(data)
	fmt.Println(unmarshaledItem.Name(), unmarshaledItem.Value())

	// Output:
	// {"name":"Active","value":1}
	// Active 1
}

func ExampleRegistry_Add() {
	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")

	registry.Add(item1, item2)

	if item, found := registry.GetByValue(Active); found {
		fmt.Println("Found by value:", item.Name())
	}

	if item, found := registry.GetByName("Inactive"); found {
		fmt.Println("Found by name:", item.Value())
	}

	// Output:
	// Found by value: Active
	// Found by name: 2
}

func ExampleRegistry_Remove() {
	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")

	registry.Add(item1)
	fmt.Println("Size before removal:", registry.Size())

	removed := registry.Remove(Active)
	fmt.Println("Removed:", removed)
	fmt.Println("Size after removal:", registry.Size())

	// Output:
	// Size before removal: 1
	// Removed: true
	// Size after removal: 0
}

func ExampleRegistry_Update() {
	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")

	registry.Add(item1)
	updated := registry.Update(Active, "Activated")
	fmt.Println("Updated:", updated)

	if item, found := registry.GetByValue(Active); found {
		fmt.Println("New name:", item.Name())
	}

	// Output:
	// Updated: true
	// New name: Activated
}

func ExampleRegistry_Items() {
	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")

	registry.Add(item1, item2)

	for _, item := range registry.Items() {
		fmt.Println(item.Name(), item.Value())
	}

	// Output:
	// Active 1
	// Inactive 2
}

func ExampleRegistry_Contains() {
	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	registry.Add(item1)

	fmt.Println("Contains Active:", registry.Contains(Active))
	fmt.Println("Contains Inactive:", registry.Contains(Inactive))
}

func ExampleRegistry_Validate() {
	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")
	registry.Add(item1, item2)

	fmt.Println("Validate Active:", registry.Validate(Active))
	fmt.Println("Validate Inactive:", registry.Validate(Inactive))
	fmt.Println("Validate Unknown:", registry.Validate(Unknown))

	// Output:
	// Validate Active: <nil>
	// Validate Inactive: <nil>
	// Validate Unknown: invalid enum value: 0
}

func ExampleRegistry_ValidateAll() {
	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")
	registry.Add(item1, item2)

	fmt.Println("ValidateAll Active, Inactive:", registry.ValidateAll(Active, Inactive))
	fmt.Println("ValidateAll Active, Unknown:", registry.ValidateAll(Active, Unknown))

	// Output:
	// ValidateAll Active, Inactive: <nil>
	// ValidateAll Active, Unknown: invalid enum value: 0
}

func ExampleRegistry_Size() {
	registry := NewRegistry[Status]()
	fmt.Println("Initial size:", registry.Size())

	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")
	registry.Add(item1, item2)

	fmt.Println("Size after adding items:", registry.Size())

	registry.Remove(Active)
	fmt.Println("Size after removing an item:", registry.Size())

	// Output:
	// Initial size: 0
	// Size after adding items: 2
	// Size after removing an item: 1
}

func ExampleRegistry_Range() {
	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")
	registry.Add(item1, item2)

	registry.Range(func(item *Item[Status]) bool {
		fmt.Println(item.Name(), item.Value())
		return true // continue iteration
	})

	// Output:
	// Active 1
	// Inactive 2
}

func ExampleRegistry_SortedItems() {
	registry := NewRegistry[Status]()
	item1 := NewItem(Inactive, "Inactive")
	item2 := NewItem(Active, "Active")
	registry.Add(item1, item2)

	for _, item := range registry.SortedItems(func(i1, i2 *Item[Status]) bool {
		return i1.value < i2.value
	}) {
		fmt.Println(item.Name(), item.Value())
	}

	// Output:
	// Active 1
	// Inactive 2
}

func ExampleRegistry_Filter() {
	registry := NewRegistry[Status]()
	item1 := NewItem(Active, "Active")
	item2 := NewItem(Inactive, "Inactive")
	registry.Add(item1, item2)

	activeItems := registry.Filter(func(item *Item[Status]) bool {
		return item.Value() == Active
	})

	for _, item := range activeItems {
		fmt.Println(item.Name(), item.Value())
	}

	// Output:
	// Active 1
}
