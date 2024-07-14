package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;
import jakarta.persistence.Id;

import java.util.Date;

@Entity
@Table(name = "tabela_cidades")
public class City {
    @Id
    @GeneratedValue
    @Column(name = "cidade_id", columnDefinition = "UNSIGNED INT(11)")
    private Long id;

    @Column(name = "cidade_nome", nullable = false, length = 45, unique = true)
    private String name;

    @Column(name = "cidade_data", nullable = false)
    private Date registrationDate;
}