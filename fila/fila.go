package main

import (
	"container/list"
	"fmt"
	"sync"
)

type Cliente struct {
	Nome   string
	Idade  int
	Idoso  bool
	Doenca string
}

type Clientes struct {
	data *list.List
	mu   sync.Mutex
}

type Atendente struct {
	Nome         string
	Atendimentos int
}

func Contrata(nome string) *Atendente {
	return &Atendente{Nome: nome}
}

func AbreFila() *Clientes {
	return &Clientes{data: list.New()}
}

func (c *Clientes) AddCliente(cli *Cliente) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data.PushBack(cli)
	return nil
}

func (c *Clientes) Atende(a *Atendente) (*Cliente, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	first := c.data.Front()
	if first == nil {
		return nil, fmt.Errorf("fila vazia")
	}

	cli := first.Value.(*Cliente)
	c.data.Remove(first)
	a.Atendimentos++
	return cli, nil
}

func (c *Clientes) Primeiro() *Cliente {
	c.mu.Lock()
	defer c.mu.Unlock()

	first := c.data.Front()
	if first == nil {
		return nil
	}
	return first.Value.(*Cliente)
}

func main() {
	fila := AbreFila()
	at := Contrata("Mariazinha")

	cli1 := &Cliente{"João", 30, false, "Meningite"}
	cli2 := &Cliente{"Claudio", 30, false, "Sarampo"}
	cli3 := &Cliente{"Gusta", 30, false, "Dengue"}
	cli4 := &Cliente{"Maria", 60, true, "Virose"}

	var wg sync.WaitGroup
	wg.Add(2)

	// produtor concorrente
	go func() {
		defer wg.Done()
		fila.AddCliente(cli1)
		fila.AddCliente(cli2)
		fila.AddCliente(cli3)
		fila.AddCliente(cli4)
	}()

	// consumidor concorrente
	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			cli, _ := fila.Atende(at)
			fmt.Println("Atendido:", cli.Nome)
		}
	}()

	wg.Wait()
	fmt.Println("Total atendimentos:", at.Atendimentos)
}
