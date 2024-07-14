package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;

@Entity
@Table(name = "tabela_equipe")
public class Team {
    @Id
    @GeneratedValue
    @Column(name = "id")
    private Integer id;

    @Column(name = "nome_completo", length = 200, nullable = false)
    private String nome;

    @Column(name = "email", length = 50, nullable = false)
    private String email;

    @Column(name = "funcao_id", nullable = false)
    private Integer roleId;

    @Column(name = "mostrar_home", nullable = true, columnDefinition = "TINYINT(1)", length = 1)
    private Boolean showHome;

    @Column(name = "mostrar_contatos", nullable = true, columnDefinition = "TINYINT(1)", length = 1)
    private Boolean showContact;
}
