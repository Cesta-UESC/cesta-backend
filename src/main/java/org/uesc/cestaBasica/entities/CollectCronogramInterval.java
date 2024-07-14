package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;

import java.util.Date;

@Entity
@Table(name = "tabela_auxiliar_cronograma")
public class CollectCronogramInterval {
    @Id
    @GeneratedValue
    @Column(name = "id", nullable = true)
    // TODO: add this column
    private Integer id;

    // TODO: ADD CRONOGRAMA RELATIONSHIP
    @Column(name = "cronograma_id", columnDefinition = "UNSIGNED INT(11)", nullable = false)
    private Long cronogramId;

    // TODO: ADD MONTH RELATIONSHIP
    @Column(name = "mes_id", nullable = false)
    private Integer monthId;

    @Column(name = "inicio_coleta", nullable = false)
    private Date startDate;

    @Column(name = "fim_coleta", nullable = false)
    private Date endDate;
}