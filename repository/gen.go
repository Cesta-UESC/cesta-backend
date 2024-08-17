package repository

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                            = new(Query)
	AuxiliarCronograma           *auxiliarCronograma
	AuxiliarPrecos               *auxiliarPrecos
	Bairros                      *bairros
	Boletim                      *boletim
	Cidades                      *cidades
	ColetaEstSec                 *coletaEstSec
	Coletas                      *coletas
	CronogramaColetas            *cronogramaColetas
	DelimitadorRacao             *delimitadorRacao
	Equipe                       *equipe
	EquipeFuncoes                *equipeFuncoes
	EstabelecimentoHasSecundario *estabelecimentoHasSecundario
	Estabelecimentos             *estabelecimentos
	EstabelecimentosSecundarios  *estabelecimentosSecundarios
	Mes                          *mes
	PesquisaResultadosProdutos   *pesquisaResultadosProdutos
	Pesquisas                    *pesquisas
	PesquisasCidades             *pesquisasCidades
	Precos                       *precos
	Produtos                     *produtos
	ProdutosMedidas              *produtosMedidas
	RacaoMinima                  *racaoMinima
	Salarios                     *salarios
	Sessao                       *sessao
	TiposProdutos                *tiposProdutos
	UnidadeMedidas               *unidadeMedidas
	Usuarios                     *usuarios
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AuxiliarCronograma = &Q.AuxiliarCronograma
	AuxiliarPrecos = &Q.AuxiliarPrecos
	Bairros = &Q.Bairros
	Boletim = &Q.Boletim
	Cidades = &Q.Cidades
	ColetaEstSec = &Q.ColetaEstSec
	Coletas = &Q.Coletas
	CronogramaColetas = &Q.CronogramaColetas
	DelimitadorRacao = &Q.DelimitadorRacao
	Equipe = &Q.Equipe
	EquipeFuncoes = &Q.EquipeFuncoes
	EstabelecimentoHasSecundario = &Q.EstabelecimentoHasSecundario
	Estabelecimentos = &Q.Estabelecimentos
	EstabelecimentosSecundarios = &Q.EstabelecimentosSecundarios
	Mes = &Q.Mes
	PesquisaResultadosProdutos = &Q.PesquisaResultadosProdutos
	Pesquisas = &Q.Pesquisas
	PesquisasCidades = &Q.PesquisasCidades
	Precos = &Q.Precos
	Produtos = &Q.Produtos
	ProdutosMedidas = &Q.ProdutosMedidas
	RacaoMinima = &Q.RacaoMinima
	Salarios = &Q.Salarios
	Sessao = &Q.Sessao
	TiposProdutos = &Q.TiposProdutos
	UnidadeMedidas = &Q.UnidadeMedidas
	Usuarios = &Q.Usuarios
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                           db,
		AuxiliarCronograma:           newAuxiliarCronograma(db, opts...),
		AuxiliarPrecos:               newAuxiliarPrecos(db, opts...),
		Bairros:                      newBairros(db, opts...),
		Boletim:                      newBoletim(db, opts...),
		Cidades:                      newCidades(db, opts...),
		ColetaEstSec:                 newColetaEstSec(db, opts...),
		Coletas:                      newColetas(db, opts...),
		CronogramaColetas:            newCronogramaColetas(db, opts...),
		DelimitadorRacao:             newDelimitadorRacao(db, opts...),
		Equipe:                       newEquipe(db, opts...),
		EquipeFuncoes:                newEquipeFuncoes(db, opts...),
		EstabelecimentoHasSecundario: newEstabelecimentoHasSecundario(db, opts...),
		Estabelecimentos:             newEstabelecimentos(db, opts...),
		EstabelecimentosSecundarios:  newEstabelecimentosSecundarios(db, opts...),
		Mes:                          newMes(db, opts...),
		PesquisaResultadosProdutos:   newPesquisaResultadosProdutos(db, opts...),
		Pesquisas:                    newPesquisas(db, opts...),
		PesquisasCidades:             newPesquisasCidades(db, opts...),
		Precos:                       newPrecos(db, opts...),
		Produtos:                     newProdutos(db, opts...),
		ProdutosMedidas:              newProdutosMedidas(db, opts...),
		RacaoMinima:                  newRacaoMinima(db, opts...),
		Salarios:                     newSalarios(db, opts...),
		Sessao:                       newSessao(db, opts...),
		TiposProdutos:                newTiposProdutos(db, opts...),
		UnidadeMedidas:               newUnidadeMedidas(db, opts...),
		Usuarios:                     newUsuarios(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AuxiliarCronograma           auxiliarCronograma
	AuxiliarPrecos               auxiliarPrecos
	Bairros                      bairros
	Boletim                      boletim
	Cidades                      cidades
	ColetaEstSec                 coletaEstSec
	Coletas                      coletas
	CronogramaColetas            cronogramaColetas
	DelimitadorRacao             delimitadorRacao
	Equipe                       equipe
	EquipeFuncoes                equipeFuncoes
	EstabelecimentoHasSecundario estabelecimentoHasSecundario
	Estabelecimentos             estabelecimentos
	EstabelecimentosSecundarios  estabelecimentosSecundarios
	Mes                          mes
	PesquisaResultadosProdutos   pesquisaResultadosProdutos
	Pesquisas                    pesquisas
	PesquisasCidades             pesquisasCidades
	Precos                       precos
	Produtos                     produtos
	ProdutosMedidas              produtosMedidas
	RacaoMinima                  racaoMinima
	Salarios                     salarios
	Sessao                       sessao
	TiposProdutos                tiposProdutos
	UnidadeMedidas               unidadeMedidas
	Usuarios                     usuarios
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                           db,
		AuxiliarCronograma:           q.AuxiliarCronograma.clone(db),
		AuxiliarPrecos:               q.AuxiliarPrecos.clone(db),
		Bairros:                      q.Bairros.clone(db),
		Boletim:                      q.Boletim.clone(db),
		Cidades:                      q.Cidades.clone(db),
		ColetaEstSec:                 q.ColetaEstSec.clone(db),
		Coletas:                      q.Coletas.clone(db),
		CronogramaColetas:            q.CronogramaColetas.clone(db),
		DelimitadorRacao:             q.DelimitadorRacao.clone(db),
		Equipe:                       q.Equipe.clone(db),
		EquipeFuncoes:                q.EquipeFuncoes.clone(db),
		EstabelecimentoHasSecundario: q.EstabelecimentoHasSecundario.clone(db),
		Estabelecimentos:             q.Estabelecimentos.clone(db),
		EstabelecimentosSecundarios:  q.EstabelecimentosSecundarios.clone(db),
		Mes:                          q.Mes.clone(db),
		PesquisaResultadosProdutos:   q.PesquisaResultadosProdutos.clone(db),
		Pesquisas:                    q.Pesquisas.clone(db),
		PesquisasCidades:             q.PesquisasCidades.clone(db),
		Precos:                       q.Precos.clone(db),
		Produtos:                     q.Produtos.clone(db),
		ProdutosMedidas:              q.ProdutosMedidas.clone(db),
		RacaoMinima:                  q.RacaoMinima.clone(db),
		Salarios:                     q.Salarios.clone(db),
		Sessao:                       q.Sessao.clone(db),
		TiposProdutos:                q.TiposProdutos.clone(db),
		UnidadeMedidas:               q.UnidadeMedidas.clone(db),
		Usuarios:                     q.Usuarios.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                           db,
		AuxiliarCronograma:           q.AuxiliarCronograma.replaceDB(db),
		AuxiliarPrecos:               q.AuxiliarPrecos.replaceDB(db),
		Bairros:                      q.Bairros.replaceDB(db),
		Boletim:                      q.Boletim.replaceDB(db),
		Cidades:                      q.Cidades.replaceDB(db),
		ColetaEstSec:                 q.ColetaEstSec.replaceDB(db),
		Coletas:                      q.Coletas.replaceDB(db),
		CronogramaColetas:            q.CronogramaColetas.replaceDB(db),
		DelimitadorRacao:             q.DelimitadorRacao.replaceDB(db),
		Equipe:                       q.Equipe.replaceDB(db),
		EquipeFuncoes:                q.EquipeFuncoes.replaceDB(db),
		EstabelecimentoHasSecundario: q.EstabelecimentoHasSecundario.replaceDB(db),
		Estabelecimentos:             q.Estabelecimentos.replaceDB(db),
		EstabelecimentosSecundarios:  q.EstabelecimentosSecundarios.replaceDB(db),
		Mes:                          q.Mes.replaceDB(db),
		PesquisaResultadosProdutos:   q.PesquisaResultadosProdutos.replaceDB(db),
		Pesquisas:                    q.Pesquisas.replaceDB(db),
		PesquisasCidades:             q.PesquisasCidades.replaceDB(db),
		Precos:                       q.Precos.replaceDB(db),
		Produtos:                     q.Produtos.replaceDB(db),
		ProdutosMedidas:              q.ProdutosMedidas.replaceDB(db),
		RacaoMinima:                  q.RacaoMinima.replaceDB(db),
		Salarios:                     q.Salarios.replaceDB(db),
		Sessao:                       q.Sessao.replaceDB(db),
		TiposProdutos:                q.TiposProdutos.replaceDB(db),
		UnidadeMedidas:               q.UnidadeMedidas.replaceDB(db),
		Usuarios:                     q.Usuarios.replaceDB(db),
	}
}

type queryCtx struct {
	AuxiliarCronograma           IAuxiliarCronogramaDo
	AuxiliarPrecos               IAuxiliarPrecosDo
	Bairros                      IBairrosDo
	Boletim                      IBoletimDo
	Cidades                      ICidadesDo
	ColetaEstSec                 IColetaEstSecDo
	Coletas                      IColetasDo
	CronogramaColetas            ICronogramaColetasDo
	DelimitadorRacao             IDelimitadorRacaoDo
	Equipe                       IEquipeDo
	EquipeFuncoes                IEquipeFuncoesDo
	EstabelecimentoHasSecundario IEstabelecimentoHasSecundarioDo
	Estabelecimentos             IEstabelecimentosDo
	EstabelecimentosSecundarios  IEstabelecimentosSecundariosDo
	Mes                          IMesDo
	PesquisaResultadosProdutos   IPesquisaResultadosProdutosDo
	Pesquisas                    IPesquisasDo
	PesquisasCidades             IPesquisasCidadesDo
	Precos                       IPrecosDo
	Produtos                     IProdutosDo
	ProdutosMedidas              IProdutosMedidasDo
	RacaoMinima                  IRacaoMinimaDo
	Salarios                     ISalariosDo
	Sessao                       ISessaoDo
	TiposProdutos                ITiposProdutosDo
	UnidadeMedidas               IUnidadeMedidasDo
	Usuarios                     IUsuariosDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AuxiliarCronograma:           q.AuxiliarCronograma.WithContext(ctx),
		AuxiliarPrecos:               q.AuxiliarPrecos.WithContext(ctx),
		Bairros:                      q.Bairros.WithContext(ctx),
		Boletim:                      q.Boletim.WithContext(ctx),
		Cidades:                      q.Cidades.WithContext(ctx),
		ColetaEstSec:                 q.ColetaEstSec.WithContext(ctx),
		Coletas:                      q.Coletas.WithContext(ctx),
		CronogramaColetas:            q.CronogramaColetas.WithContext(ctx),
		DelimitadorRacao:             q.DelimitadorRacao.WithContext(ctx),
		Equipe:                       q.Equipe.WithContext(ctx),
		EquipeFuncoes:                q.EquipeFuncoes.WithContext(ctx),
		EstabelecimentoHasSecundario: q.EstabelecimentoHasSecundario.WithContext(ctx),
		Estabelecimentos:             q.Estabelecimentos.WithContext(ctx),
		EstabelecimentosSecundarios:  q.EstabelecimentosSecundarios.WithContext(ctx),
		Mes:                          q.Mes.WithContext(ctx),
		PesquisaResultadosProdutos:   q.PesquisaResultadosProdutos.WithContext(ctx),
		Pesquisas:                    q.Pesquisas.WithContext(ctx),
		PesquisasCidades:             q.PesquisasCidades.WithContext(ctx),
		Precos:                       q.Precos.WithContext(ctx),
		Produtos:                     q.Produtos.WithContext(ctx),
		ProdutosMedidas:              q.ProdutosMedidas.WithContext(ctx),
		RacaoMinima:                  q.RacaoMinima.WithContext(ctx),
		Salarios:                     q.Salarios.WithContext(ctx),
		Sessao:                       q.Sessao.WithContext(ctx),
		TiposProdutos:                q.TiposProdutos.WithContext(ctx),
		UnidadeMedidas:               q.UnidadeMedidas.WithContext(ctx),
		Usuarios:                     q.Usuarios.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
