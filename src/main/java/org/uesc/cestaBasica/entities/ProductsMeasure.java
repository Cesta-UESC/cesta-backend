package org.uesc.cestaBasica.entities;

import java.math.BigDecimal;

import org.uesc.cestaBasica.entities.compositeForeingKeys.ProductsMeasurePK;

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
@Table(name = "tabela_produtos_medidas")
@IdClass(ProductsMeasurePK.class)
public class ProductsMeasure {

    @Id
    @Column(name = "produto_id")
    private Integer productId;

    @Id
    @Column(name = "medida_id")
    private Integer measureId;

    @Column(name = "medida_pesquisada", nullable = false, precision = 16, scale = 3)
    private BigDecimal searchedMeasure;

}
