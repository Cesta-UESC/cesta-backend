package org.uesc.cestaBasica.entities;

import java.math.BigDecimal;

import org.uesc.cestaBasica.entities.compositeForeingKeys.SearchCityPK;

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
@Table(name = "tabela_pesquisas_cidades")
@IdClass(SearchCityPK.class)
public class SearchCity {

    @Id
    @Column(name = "cidade_id")
    private Integer cityId;

    @Id
    @Column(name = "pesquisa_id")
    private Integer searchId;

    @Column(name = "variacao_mensal", nullable = true, precision = 16, scale = 2)
    private BigDecimal monthlyVariation;

    @Column(name = "variacao_semestral", nullable = true, precision = 16, scale = 2)
    private BigDecimal semiannualVariation;

    @Column(name = "variacao_anual", nullable = true, precision = 16, scale = 2)
    private BigDecimal anualVariation;

    @Column(name = "gasto_mensal_cesta", nullable = true, precision = 16, scale = 2)
    private BigDecimal monthlyExpense;

    @Column(name = "tempo_trabalho", nullable = true, precision = 16, scale = 2)
    private BigDecimal workTime;

}
