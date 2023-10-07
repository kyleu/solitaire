// Package menu - Content managed by Project Forge, see [projectforge.md] for details.
package menu

import "github.com/samber/lo"

var Separator = &Item{}

type Item struct {
	Key         string `json:"key"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Badge       string `json:"badge,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Route       string `json:"route,omitempty"`
	Children    Items  `json:"children,omitempty"`
}

func (i *Item) AddChild(child *Item) {
	i.Children = append(i.Children, child)
}

func (i *Item) Desc() string {
	if i.Description != "" {
		return i.Description
	}
	return i.Title
}

type Items []*Item

func (i Items) Get(key string) *Item {
	return lo.FindOrElse(i, nil, func(item *Item) bool {
		return item.Key == key
	})
}

func (i Items) GetByPath(path []string) *Item {
	if len(path) == 0 {
		return nil
	}
	ret := i.Get(path[0])
	if ret == nil {
		return nil
	}
	if len(path) > 1 {
		return ret.Children.GetByPath(path[1:])
	}
	return ret
}

func (i Items) Keys() []string {
	return lo.Map(i, func(x *Item, _ int) string {
		return x.Key
	})
}
