package org.uesc.cestaBasica.entities;

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
@Table(name = "tabela_usuarios")
public class User {
    
    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    @Column(name = "usuario_id")
    private int id;

    @Column(name = "usuario_nome", nullable = false, length = 64)
    private String name;

    // TODO: encript passwd
    @Deprecated
    @Column(name = "usuario_senha", nullable = false, length = 128)
    private String clean_password;

    @Column(name = "usuario_email", nullable = false, length = 255)
    private String email;
}
