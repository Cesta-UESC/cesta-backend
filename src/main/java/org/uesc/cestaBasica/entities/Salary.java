package org.uesc.cestaBasica.entities;

import java.math.BigDecimal;
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
@Table(name = "tabela_tipos_produtos")
public class Salary {
    
    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    @Column(name = "salario_id")
    private Integer id;

    @Column(name = "salario_valor_bruto", nullable = false, precision = 16, scale = 2)
    private BigDecimal rawValue;

    @Column(name = "salario_valor_liquido", nullable = false, precision = 16, scale = 2)
    private BigDecimal value;

    @Column(name = "salario_em_uso", nullable = false)
    private Integer inUse;

    @Column(name = "salario_data_registro", nullable = false)
    @JsonFormat(shape = JsonFormat.Shape.STRING, pattern="yyyy-MM-dd")
    private Date registrationDate;

    @Column(name = "salario_nome", nullable = false, length = 20)
    private String name;

    @Column(name = "salario_simbolo", nullable = false, length = 20)
    private String symbol;

}
