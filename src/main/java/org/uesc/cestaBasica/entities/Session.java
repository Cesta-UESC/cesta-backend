package org.uesc.cestaBasica.entities;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Table;
import jakarta.persistence.Id;

@Entity
@Table(name = "tabela_sessao")
public class Session {
    @Id
    @Column(name = "sessao_id", length = 128, nullable = false)
    private String id;

    @Column(name = "sessao_ip", length = 16, nullable = false)
    private String ip;

    @Column(name = "sessao_tempo")
    private long time;
}