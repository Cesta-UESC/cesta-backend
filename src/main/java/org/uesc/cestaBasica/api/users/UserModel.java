package org.uesc.cestaBasica.api.users;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class UserModel {
    private int id;
    private String name;
    private String email;
}