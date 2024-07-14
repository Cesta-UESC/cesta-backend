package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;
import jakarta.persistence.Id;

import java.math.BigDecimal;
import java.util.Date;

@Entity
@Table(name = "tabela_coletas")
public class Collect {
    @Id
    @GeneratedValue
    @Column(name = "coleta_id")
    private Integer id;

    @Column(name = "coleta_preco_cesta", nullable = true, scale = 16, precision = 2)
    private BigDecimal value;

    @Column(name = "coleta_data", nullable = false)
    private Date date;

    @Column(name = "coleta_fechada", nullable = true)
    private Integer closed;

    // TODO: add relationship
    @Column(name = "estabelecimento_id", nullable = true, columnDefinition = "UNSIGNED INT(11)")
    private Long estabelecimentoId;

    // TODO: add relationship
    @Column(name = "pesquisa_id", nullable = false, columnDefinition = "UNSIGNED INT(11)")
    private Long pesquisaId;
}