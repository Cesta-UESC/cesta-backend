package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Table;
import jakarta.persistence.UniqueConstraint;

@Entity
@Table(name = "tabela_coleta_est_sec",
    uniqueConstraints=@UniqueConstraint(columnNames={"coleta_id", "produto_id", "est_has_sec_id"})
)
public class CollectSecondaryOrganization {
    @Id
    @GeneratedValue
    @Column(name = "coleta_est_sec_id", columnDefinition = "UNSIGNED INT(11)")
    private Long id;

    // TODO: ADD  RELATIONSHIP
    @Column(name = "coleta_id", nullable = false)
    private Integer collectId;

    // TODO: ADD  RELATIONSHIP
    @Column(name = "produto_id", columnDefinition = "UNSIGNED INT(11)", nullable = false, unique = true)
    private Long produtoId;

    // TODO: ADD MONTH RELATIONSHIP
    @Column(name = "est_has_sec_id", columnDefinition = "UNSIGNED INT(11)", nullable = false, unique = true)
    private Long organizationHasSecundaryId;
}