package org.uesc.cestaBasica.entities;

import java.math.BigDecimal;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
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
@Table(name = "tabela_precos")
public class Price {
    
    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    @Column(name = "precos_id")
    private Integer id;

    @Column(name = "medida_id", nullable = false)
    private Integer measureId;

    @Column(name = "produto_id", nullable = false)
    private Integer productId;

    @Column(name = "coleta_id", nullable = false)
    private Integer collectId;

    @Column(name = "precos_media_observado", nullable = true, precision = 16, scale = 2)
    private BigDecimal observedAverage;

    @Column(name = "precos_media", nullable = true, precision = 16, scale = 2)
    private BigDecimal average;

    @Column(name = "precos_total", nullable = true, precision = 16, scale = 2)
    private BigDecimal totalPrice;

}
