package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "tabela_estabelecimento_has_secundario")
public class OrganizationHasSecond {

    @Id
    @GeneratedValue
    @Column(name = "est_has_sec_id")
    private Integer id;

    @Column(name = "estabelecimento_id")
    private Integer organizationId;

    @Column(name = "estabelecimento_sec_id")
    private Integer secondaryOrganizationId;

}