package org.uesc.cestaBasica.entities;

import java.util.Set;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.OneToMany;
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
@Table(name = "tabela_cronograma_coletas")
public class CollectCronogram {
    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    @Column(name = "cronograma_id", columnDefinition = "UNSIGNED INT(11)")
    private Long id;

    @Column(name = "ano", columnDefinition = "UNSIGNED INT(11)", nullable = false)
    private Long year;

    @OneToMany(cascade = CascadeType.ALL, mappedBy = "collectCronogram")
    private Set<CollectCronogramInterval> collectCronogramIntervals;
}