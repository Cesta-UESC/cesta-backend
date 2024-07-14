package org.uesc.cestaBasica.entities;

import java.math.BigDecimal;

import org.uesc.cestaBasica.entities.compositeForeingKeys.MinimalRationPK;

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
@Table(name = "tabela_racao_minima")
@IdClass(MinimalRationPK.class)
public class MinimalRation {

    @Id
    private Integer delimiterId;

    @Id
    private Integer productId;

    @Id
    private Integer measureId;

    @Column(name = "racao_minima_quantidade", nullable = false, precision = 10, scale = 2)
    private BigDecimal minimumQuantity;

    @Column(name = "racao_minima_transformador", nullable = false, precision = 10, scale = 2)
    private BigDecimal minimumTransformer;

    @Column(name = "racao_minima_medida", nullable = false)
    private Integer minimumMeasure;

}
