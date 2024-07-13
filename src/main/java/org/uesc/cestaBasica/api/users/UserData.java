package org.uesc.cestaBasica.api.users;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@AllArgsConstructor
@NoArgsConstructor
@Data
public class UserData {
    
    private int id;

    private String name;

    @Deprecated
    private String clean_password;

    private String email;

}