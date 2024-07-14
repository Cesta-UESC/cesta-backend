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
@Table(name = "tabela_tipos_produtos")
public class ProductsType {
    
    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    @Column(name = "tipo_id")
    private Integer id;

    @Column(name = "tipo_nome", nullable = false, length = 45)
    private String name;

}
