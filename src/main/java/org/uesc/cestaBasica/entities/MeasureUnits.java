package org.uesc.cestaBasica.entities;

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
@Table(name = "tabela_unidade_medidas")
public class MeasureUnits {
    
    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    @Column(name = "medida_id")
    private Integer id;

    @Column(name = "medida_descricao", nullable = false, length = 20)
    private String description;

    @Column(name = "medida_simbolo", nullable = false, length = 20)
    private String simbol;
}
