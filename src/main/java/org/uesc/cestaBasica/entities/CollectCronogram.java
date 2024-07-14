package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;
import jakarta.persistence.Id;

@Entity
@Table(name = "tabela_cronograma_coletas")
public class CollectCronogram {
    @Id
    @GeneratedValue
    @Column(name = "cronograma_id", columnDefinition = "UNSIGNED INT(11)")
    private Long id;

    @Column(name = "ano", columnDefinition = "UNSIGNED INT(11)", nullable = false)
    private Long year;
}