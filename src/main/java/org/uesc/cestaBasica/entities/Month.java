package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;
import jakarta.persistence.Id;

@Entity
@Table(name = "tabela_mes")
public class Month {
    @Id
    @GeneratedValue
    @Column(name = "mes_id")
    private Integer id;

    @Column(name = "mes_nome", nullable = false, length = 45)
    private String name;
}