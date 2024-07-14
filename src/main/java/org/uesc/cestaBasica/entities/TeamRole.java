package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;

@Entity
@Table(name = "tabela_equipe_funcoes")
public class TeamRole {
    @Id
    @GeneratedValue
    @Column(name = "id")
    private Integer id;

    @Column(name = "funcao", length = 100, nullable = false, unique = true)
    private String nome;
}