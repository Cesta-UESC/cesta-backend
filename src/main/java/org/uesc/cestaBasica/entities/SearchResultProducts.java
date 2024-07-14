package org.uesc.cestaBasica.entities;

import java.math.BigDecimal;

import org.uesc.cestaBasica.entities.compositeForeingKeys.SearchResultProductsPK;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.IdClass;
import jakarta.persistence.Table;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Entity
@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Table(name = "tabela_pesquisa_resultados_produtos")
@IdClass(SearchResultProductsPK.class)
public class SearchResultProducts {

    @Id
    @Column(name = "produto_id")
    private Integer productId;

    @Id
    @Column(name = "cidade_id")
    private Integer cityId;

    @Id
    @Column(name = "pesquisa_id")
    private Integer searchId;

    @Column(name = "produto_variacao_mensal", nullable = true, precision = 16, scale = 2)
    private BigDecimal monthlyVariation;

    @Column(name = "produto_variacao_semestral", nullable = true, precision = 16, scale = 2)
    private BigDecimal semiannualVariation;

    @Column(name = "produto_variacao_anual", nullable = true, precision = 16, scale = 2)
    private BigDecimal anualVariation;

    @Column(name = "produto_preco_medio", nullable = true, precision = 16, scale = 2)
    private BigDecimal averagePrice;

    @Column(name = "produto_preco_total", nullable = true, precision = 16, scale = 2)
    private BigDecimal totalPrice;

    @Column(name = "produto_tempo_trabalho", nullable = true, precision = 16, scale = 2)
    private BigDecimal workTime;

}
