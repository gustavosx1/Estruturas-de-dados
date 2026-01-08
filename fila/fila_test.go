package main

import "testing"

func TestAddAtendePessoa(t *testing.T) {
	fila := AbreFila()
	cli1 := &Cliente{"João", 30, false, "Meningite"}
	cli2 := &Cliente{"Claudio", 30, false, "Sarmapo"}
	cli3 := &Cliente{"Gusta", 30, false, "Dengue"}
	cli4 := &Cliente{"Maria", 60, true, "Virose"}
	at := Contrata("Mariazinha")
	_ = fila.AddCliente(cli1)
	_ = fila.AddCliente(cli2)
	_ = fila.AddCliente(cli3)
	_ = fila.AddCliente(cli4)

	if c := fila.data.Back(); c.Value != cli4 || c.Value == nil {
		t.Errorf("Erro, o último da fila não existe ou está errado, ACTUAL: %v", c.Value)
	}

	_ = fila.Atende(at)
	_ = fila.Atende(at)
	_ = fila.Atende(at)
	_ = fila.Atende(at)

	if c := fila.data.Front(); c != nil {
		t.Errorf("erro, não deveria ter nenhum cliente na fila: %v", c.Value)
	}
}

func TestFuraFila(t *testing.T) {
	fila := AbreFila()
	cli1 := &Cliente{"João", 30, false, "Meningite"}
	cli2 := &Cliente{"Claudio", 30, false, "Sarmapo"}
	cli3 := &Cliente{"Gusta", 30, false, "Dengue"}
	cli4 := &Cliente{"Maria", 60, true, "Virose"}
	_ = fila.AddCliente(cli1)
	_ = fila.AddCliente(cli2)
	_ = fila.AddCliente(cli3)
	_ = fila.AddCliente(cli4)

	err := fila.FuraFila(cli4)
	if err != nil {
		t.Errorf("erro ao furar a fila por algum motivo: %v", err)
	}

	if fila.data.Front().Value != cli4 {
		t.Errorf("erro, primeiro da fila actual: %v", fila.data.Front().Value)
	}

	err = fila.FuraFila(cli3)
	if err == nil {
		t.Errorf("erro esperado aqui ao tentar furar a fila com cliente jovem: %v", err)
	}
}
