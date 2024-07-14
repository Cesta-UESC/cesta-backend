package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;
import jakarta.persistence.Id;

import java.util.Date;

@Entity
@Table(name = "tabela_delimitador_racao")
public class PortionDelimiter {
    @Id
    @GeneratedValue
    @Column(name = "delimitador_id", columnDefinition = "UNSIGNED INT(11)")
    private Long id;

    @Column(name = "delimitador_descricao", length = 255, nullable = false)
    private String description;

    @Column(name = "delimitador_em_uso", nullable = false, columnDefinition = "INT")
    private Boolean used;

    @Column(name = "delimitador_data_registro", nullable = false)
    private Date date;

    @Column(name = "delimitador_oficial", nullable = false, columnDefinition = "SMALLINT")
    private Boolean oficial;
}