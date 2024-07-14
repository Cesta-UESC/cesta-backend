package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;
import jakarta.persistence.Id;

import java.util.Date;

@Entity
@Table(name = "tabela_estabelecimentos")
public class Organization {
    @Id
    @GeneratedValue
    @Column(name = "estabelecimento_id", columnDefinition = "UNSIGNED INT(11)")
    private Long id;

    // TODO: add neighbourhood relationship
    @Column(name = "bairro_id")
    private Integer neighborhoodId;

    @Column(name = "estabelecimento_nome", nullable = false, length = 45)
    private String name;
    
    @Column(name = "estabelecimento_endereco", nullable = true, length = 255)
    private String address;

    @Column(name = "estabelecimento_contato", nullable = true, length = 45)
    private String contact;

    @Column(name = "estabelecimento_telefone", nullable = true, length = 20)
    private String phone;

    @Column(name = "estabelecimento_referencial", nullable = true, length = 255)
    private String reference;

    @Column(name = "estabelecimento_data", nullable = false)
    private Date date;

    @Column(name = "estabelecimento_ativo", nullable = true, columnDefinition = "TINYINT(1)", length = 1)
    private Boolean enabled;
}