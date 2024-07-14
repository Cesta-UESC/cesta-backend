package org.uesc.cestaBasica.services;

import org.springframework.stereotype.Service;
import org.uesc.cestaBasica.api.users.UserModel;
import org.uesc.cestaBasica.entities.User;

@Service
public class Converter {
    public UserModel parse(User u) {
        return UserModel
                .builder()
                .id(u.getId())
                .email(u.getEmail())
                .name(u.getName())
                .build();
    }
}