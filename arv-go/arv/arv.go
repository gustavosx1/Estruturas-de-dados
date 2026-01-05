package arv

import "fmt"

type No struct {
	Data int
	Dir  *No
	Esq  *No
}

func Create(data int) *No {
	return &No{Data: data}
}

func (n *No) Print() {
	if n == nil {
		return
	}

	fila := []*No{n}

	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]

		fmt.Println(atual.Data)
		if atual.Dir != nil {
			fila = append(fila, atual.Dir)
		}

		if atual.Esq != nil {
			fila = append(fila, atual.Esq)
		}
	}
}

func (n *No) Altura() int {
	if n == nil {
		return -1
	}
	dir := n.Dir.Altura()
	esq := n.Esq.Altura()
	if dir > esq {
		return 1 + dir
	}
	return 1 + esq
}

func (n *No) Busca(key int) (*No, error) {
	if n == nil {
		return nil, fmt.Errorf("erro ao buscar valor na árvore")
	}
	if n.Data > key {
		return n.Esq.Busca(key)
	} else if n.Data < key {
		return n.Dir.Busca(key)
	}
	return n, nil
}

func (n *No) Insert(d int) error {
	if n.Data > d {
		if n.Esq == nil {
			n.Esq = &No{Data: d}
			return nil
		} else {
			return n.Esq.Insert(d)
		}
	} else if n.Data < d {
		if n.Dir == nil {
			n.Dir = &No{Data: d}
			return nil
		} else {
			return n.Dir.Insert(d)
		}
	}
	return fmt.Errorf("erro, Valor provavelmente já está na árvore")
}
