package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func main() {
	config := &Config{}

	_, err := toml.DecodeFile("./conf/fruit.toml", config)
	if err != nil {
		fmt.Printf("get config failed. err : %s ", err)
		return
	}

	fmt.Printf("%+v \n", config)
}

type Config struct {
	Fruits []*Fruit `toml:"fruit"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%+v", *c)
}

type Fruit struct {
	Name      string     `toml:"name"`
	Physical  *Physical  `toml:"physical"`
	Varieties []*Variety `toml:"variety"`
}

func (c *Fruit) String() string {
	return fmt.Sprintf("%+v", *c)
}

type Physical struct {
	Color string `toml:"color"`
	Shape string `toml:"shape"`
}

func (c *Physical) String() string {
	return fmt.Sprintf("%+v", *c)
}

type Variety struct {
	Name string `toml:"name"`
}

func (c *Variety) String() string {
	return fmt.Sprintf("%+v", *c)
}
