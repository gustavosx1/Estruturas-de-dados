package arv

import "testing"

func TestCria(t *testing.T) {
	d := 10
	arv := Create(d)
	if arv == nil || arv.Data != d {
		t.Error("Erro ao criar a árvore")
	}
}

func TestInsere(t *testing.T) {
	d := 10
	no1 := 15
	no2 := 5
	no3 := 20
	no4 := 15
	arv := Create(d)
	err := arv.Insert(no1)
	if err != nil {
		t.Errorf("Erro Inserindo o elemento 1: %e", err)
	}
	_ = arv.Insert(no2)
	_ = arv.Insert(no3)
	err = arv.Insert(no4)
	if err == nil {
		t.Error("Erro esperado ao inserir valor igual!")
	}
}

func TestBusca(t *testing.T) {
	d := 10
	no1 := 15
	no2 := 5
	no3 := 20
	no4 := 7
	arv := Create(d)
	_ = arv.Insert(no1)
	_ = arv.Insert(no2)
	_ = arv.Insert(no3)

	busca, err := arv.Busca(no3)
	if err != nil || busca.Data != no3 {
		t.Errorf("erro ao buscar elemento")
	}

	b, err := arv.Busca(no4)
	if err == nil || b != nil {
		t.Error("erro esperado aqui: ")
	}
}

func TestAltura(t *testing.T) {
	d := 10
	no1 := 15
	no2 := 5
	no3 := 20
	no5 := 25
	no6 := 30
	no7 := 35
	arv := Create(d)
	err := arv.Insert(no1)
	if err != nil {
		t.Errorf("Erro Inserindo o elemento 1: %e", err)
	}
	_ = arv.Insert(no2)
	_ = arv.Insert(no3)
	_ = arv.Insert(no5)
	_ = arv.Insert(no6)
	_ = arv.Insert(no7)

	if alt := arv.Altura(); alt != 5 {
		t.Error("Altura deveria ser igual a 5 ")
	}
}
