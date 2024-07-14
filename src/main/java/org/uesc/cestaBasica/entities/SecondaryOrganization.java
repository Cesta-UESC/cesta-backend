package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "tabela_estabelecimentos_secundarios")
public class SecondaryOrganization {

    @Id
    @GeneratedValue
    @Column(name = "estabelecimento_sec_id")
    private Integer id;

    @Column(name = "estabelecimento_sec_nome", length = 260)
    private String name;

}