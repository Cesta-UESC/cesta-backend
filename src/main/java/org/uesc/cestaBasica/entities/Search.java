package org.uesc.cestaBasica.entities;

import java.util.Date;

import com.fasterxml.jackson.annotation.JsonFormat;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Entity
@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Table(name = "tabela_pesquisas")
public class Search {

    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    @Column(name = "pesquisa_id")
    private Integer searchId;

    @Column(name = "delimitador_id", nullable = true)
    private Integer delimiterId;

    @Column(name = "salario_id", nullable = true)
    private Integer salaryId;

    @Column(name = "pesquisa_fechada", nullable = false)
    private Integer closed;

    @Column(name = "pesquisa_data", nullable = false)
    @JsonFormat(shape = JsonFormat.Shape.STRING, pattern="yyyy-MM-dd")
    private Date date;

    @Column(name = "pesquisa_detalhada", nullable = true)
    private Integer detailed;
}
