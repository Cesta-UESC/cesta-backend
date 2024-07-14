package org.uesc.cestaBasica.controllers;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.uesc.cestaBasica.api.users.UserModel;
import org.uesc.cestaBasica.services.Converter;
import org.uesc.cestaBasica.services.UserService;

@RestController
@RequestMapping("/api/users")
public class UserController extends BaseController {

    @Autowired
    private UserService userService;

    @Autowired
    private Converter converter;

    @GetMapping()
    public List<UserModel> list() {
        return userService
                .list()
                .stream()
                .map(u -> converter.parse(u))
                .toList();
    }
}
