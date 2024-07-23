package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.Date;

@Entity
@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Table(name = "tabela_auxiliar_cronograma")
public class CollectCronogramInterval {
    @Id
    @GeneratedValue
    @Column(name = "id", nullable = true)
    // TODO: add this column
    private Integer id;

    @ManyToOne
    @JoinColumn(name = "cronograma_id", nullable = false)
    private CollectCronogram collectCronogram;

    @ManyToOne
    @JoinColumn(name = "mes_id", nullable = false)
    private Month month;

    @Column(name = "inicio_coleta", nullable = false)
    private Date startDate;

    @Column(name = "fim_coleta", nullable = false)
    private Date endDate;
}