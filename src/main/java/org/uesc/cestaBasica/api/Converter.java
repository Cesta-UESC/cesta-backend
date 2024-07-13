package org.uesc.cestaBasica.api;

import org.uesc.cestaBasica.api.users.UserModel;
import org.uesc.cestaBasica.entities.User;

public class Converter {
    public static UserModel parse(User u) {
        return UserModel
            .builder()
            .id(u.getId())
            .email(u.getEmail())
            .name(u.getName())
            .build();
    }
}