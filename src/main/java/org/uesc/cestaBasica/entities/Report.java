package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;
import jakarta.persistence.Id;

@Entity
@Table(name = "tabela_boletim")
public class Report {
    @Id
    @GeneratedValue
    @Column(name = "boletim_id", columnDefinition = "UNSIGNED INT(11)")
    private Long id;

    @Column(name = "boletim_nome", nullable = false, length = 45, unique = true)
    private String name;
}