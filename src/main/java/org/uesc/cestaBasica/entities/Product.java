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
@Table(name = "tabela_produtos")
public class Product {

    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    @Column(name = "produto_id")
    private Integer id;

    @Column(name = "produto_nome", nullable = false, length = 45)
    private String name;

    @Column(name = "produto_cesta", nullable = false)
    private Integer onBasket;

    @Column(name = "produto_nome_visualizacao", nullable = false, length = 45)
    private String viewName;

    @Column(name = "produto_tipo", nullable = false)
    private Integer type;

}
