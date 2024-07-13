package org.uesc.cestaBasica.controllers;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import org.uesc.cestaBasica.api.CestaData;
import org.uesc.cestaBasica.entities.Cesta;
import org.uesc.cestaBasica.services.CestaService;

import jakarta.validation.Valid;

@RestController
@RequestMapping("/api/v1/cesta")
public class CestaController extends BaseController {

    @Autowired
    private CestaService cestaService;

    @GetMapping
    public List<Cesta> listCesta() {
        return cestaService.list();
    }

    @PostMapping
    public Cesta createCesta(@Valid @RequestBody CestaData cestaData) {
        return cestaService.create(cestaData);
    }
}
