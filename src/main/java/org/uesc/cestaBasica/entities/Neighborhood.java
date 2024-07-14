package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;
import jakarta.persistence.Id;

@Entity
@Table(name = "tabela_bairros")
public class Neighborhood {
    @Id
    @GeneratedValue
    @Column(name = "bairro_id", columnDefinition = "UNSIGNED INT(11)")
    private Long id;

    @Column(name = "bairro_nome", length = 45, nullable = false)
    private String nome;

    // TODO: add cidade relationship
    @Column(name = "cidade_id", columnDefinition = "UNSIGNED INT(11)")
    private Long cidadeId;
}